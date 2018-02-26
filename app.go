package main

import (
    "fmt"
    "log"
    "errors"
    "strconv"
    "io/ioutil"
    "net/http" 
    "encoding/json"

    "github.com/renomx/truesize/config"
    "github.com/renomx/truesize/models"

    "github.com/gorilla/mux"
    _ "github.com/lib/pq"
    "github.com/jinzhu/gorm"
    _ "github.com/jinzhu/gorm/dialects/postgres"
)

type App struct {
	Config *config.Config
    Router *mux.Router
    DB     *gorm.DB
    Model  *models.Shoe
    View   *models.SimpleShoe
}

func (a *App) SetConfig() {
	a.Config = config.GetConfig()	
}

func (a *App) Initialize(host, port, user, password, dbname string) {
    connectionString :=
        fmt.Sprintf("host=%s user=%s password=%s dbname=%s sslmode=disable", 
            host, port, user, password, dbname)   

    var err error

    if a.DB, err = gorm.Open("postgres", connectionString); err != nil {
        log.Fatal(err)
    }

    a.Model.InitModel(a.DB)

    
    a.InitializeRoutes()    
}

func (a *App) InitializeRoutes() {

    a.Router = mux.NewRouter()
    a.Router.HandleFunc("/", a.sayHello).Methods("GET")
    a.Router.HandleFunc("/shoe", a.getShoes).Methods("GET")
    a.Router.HandleFunc("/shoe/{id}", a.getShoe).Methods("GET")
    a.Router.HandleFunc("/shoe", a.createShoe).Methods("POST")
    a.Router.HandleFunc("/shoe/truetosize/{shoename}", a.addTrueToSize).Methods("PUT")
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


func (a *App) sayHello(w http.ResponseWriter, r *http.Request) {
    
    text := "Hola"

    anonymousStruct := struct {
        Message    string
    }{
        text,
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
    id, err := strconv.Atoi(vars["id"])
    if err != nil {
        respondWithError(w, http.StatusBadRequest, "Invalid shoe ID")
        return
    }
    s := models.Shoe{}
    
    shoe, err := s.GetShoe(a.DB, id)
    if err != nil {        
        respondWithError(w, http.StatusInternalServerError, err.Error())
    }

    respondWithJSON(w, http.StatusOK, shoe)

}

func (a *App) createShoe(w http.ResponseWriter, r *http.Request) {
    
    var shoe = models.SimpleShoe{}
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

    var shoe = models.Shoe{}

    bodyBytes, err := ioutil.ReadAll(r.Body)
    m := make(map[string]int)
    err = json.Unmarshal(bodyBytes, &m)

    log.Println(m["size"])
    
    if err != nil {
        respondWithError(w, http.StatusBadRequest, "Invalid request payload")
        return
    }
    defer r.Body.Close()

    err = validateSize([]int {m["size"]})
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
    for _,size := range sizes {
        if(size < 1 || size > 5) {
            log.Println("Size out of range")
            return errors.New("Size must be between 1 and 5")
        }
    }    
    return nil 
}