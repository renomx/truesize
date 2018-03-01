package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	_ "github.com/lib/pq"
)

type App struct {
	Config *Config
	Router *mux.Router
	DB     *gorm.DB
	Model  *Shoe
	View   *SimpleShoe
}

func (a *App) SetConfig() {
	a.Config = GetConfig()
}

func (a *App) Initialize(host, port, user, password, dbname string) {
	connectionString :=
		fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
			host, port, user, password, dbname)

	var err error
	a.DB, err = gorm.Open("postgres", connectionString)
	if err != nil {
		log.Fatal(err)
	}

	defer a.DB.Close()

	a.InitializeRoutes()

	a.Run(a.Config.Local.ApiPort)
}

func (a *App) InitializeRoutes() {

	a.Router = mux.NewRouter()

	a.Router.HandleFunc("/", a.sayHello).Methods("GET")
	a.Router.HandleFunc("/shoe", a.getShoes).Methods("GET")
	a.Router.HandleFunc("/shoe/{shoename}", a.getShoe).Methods("GET")
	a.Router.HandleFunc("/shoe", a.createShoe).Methods("POST")
	a.Router.HandleFunc("/shoe/truetosize/{shoename}", a.addTrueToSize).Methods("PUT")

	a.Router.HandleFunc("/error", a.dbNotReady).Methods("GET")

	log.Println("Initializing routes")
}

func (a *App) Run(port string) {
	log.Printf("Listening on %s", port)
	log.Fatal(http.ListenAndServe(port, a.Router))
}

func respondWithError(w http.ResponseWriter, code int, message string) {
	respondWithJSON(w, code, map[string]string{"error": message})
}

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

/***** HANDLERS ****/

func (a *App) dbNotReady(w http.ResponseWriter, r *http.Request) {
	text := "Seems like our database is not ready yet, please try again in a few moments"
	anonymousStruct := struct {
		Message string
	}{
		text,
	}

	a.Initialize(
		a.Config.Local.Host,
		a.Config.Local.DbPort,
		a.Config.Local.User,
		a.Config.Local.Password,
		a.Config.Local.DbName)

	respondWithJSON(w, http.StatusOK, anonymousStruct)
}

func (a *App) sayHello(w http.ResponseWriter, r *http.Request) {

	text := "Welcome to True to Size Value"

	anonymousStruct := struct {
		Message            string
		ListOfShoes        string
		SpecificShoe       string
		CreateShoe         string
		AddTrueToSizeValue string
		ToDos              []string
	}{
		text,
		"/shoe - GET",
		"/shoe/{shoename} - GET",
		"/shoe - POST",
		"/shoe/truetosize/{shoename} - PUT",
		[]string{
			"Complete tests configuration",
			"Add Swagger or some sort of documentation",
			"Implement delete shoe",
		},
	}
	respondWithJSON(w, http.StatusOK, anonymousStruct)
}

func (a *App) getShoes(w http.ResponseWriter, r *http.Request) {

	shoes, err := a.Model.GetShoes(a.DB)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJSON(w, http.StatusOK, shoes)

}

func (a *App) getShoe(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	shoeName := vars["shoename"]
	s := Shoe{}

	shoe, err := s.GetShoe(a.DB, shoeName)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
	}

	respondWithJSON(w, http.StatusOK, shoe)

}

func (a *App) createShoe(w http.ResponseWriter, r *http.Request) {

	var shoe = SimpleShoe{}
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&shoe); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	defer r.Body.Close()

	err := validateSize(shoe.Sizes)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	s, err := shoe.CreateShoe(a.DB)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJSON(w, http.StatusOK, s)

}

func (a *App) addTrueToSize(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	shoeName := vars["shoename"]

	var shoe = Shoe{}

	bodyBytes, err := ioutil.ReadAll(r.Body)
	m := make(map[string]int)
	err = json.Unmarshal(bodyBytes, &m)

	log.Println(m["size"])

	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	defer r.Body.Close()

	err = validateSize([]int{m["size"]})
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	s, err := shoe.AddTrueToSize(a.DB, shoeName, m["size"])
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJSON(w, http.StatusOK, s)
}

// Validations

func validateSize(sizes []int) error {
	for _, size := range sizes {
		if size < 1 || size > 5 {
			log.Println("Size out of range")
			return errors.New("Size must be between 1 and 5")
		}
	}
	return nil
}
