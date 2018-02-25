package main

import (
    "database/sql"
    "log"
    "net/http"
    "fmt"
    "encoding/json"

    "github.com/renomx/truesize/config"
    "github.com/renomx/truesize/models"

    "github.com/gorilla/mux"
    _ "github.com/lib/pq"
)

type App struct {
	Config *config.Config
    Router *mux.Router
    DB     *sql.DB
    Model  *models.Shoe
}

func (a *App) SetConfig() {
	a.Config = config.GetConfig()	
}

func (a *App) Initialize(host, user, password, dbname string) {
    connectionString :=
        fmt.Sprintf("host=%s user=%s password=%s dbname=%s  sslmode=disable", host, user, password, dbname)

    log.Println(connectionString)

    var err error
    a.DB, err = sql.Open("postgres", connectionString)
    if err != nil {
        log.Fatal(err)
    }
    defer a.DB.Close()

    a.Model.InitModel(a.DB)

    a.Router = mux.NewRouter()
    a.initializeRoutes()    
}

func (a *App) initializeRoutes() {
    /*a.Router.HandleFunc("/products", a.getProducts).Methods("GET")
    a.Router.HandleFunc("/product", a.createProduct).Methods("POST")
    a.Router.HandleFunc("/product/{id:[0-9]+}", a.getProduct).Methods("GET")
    a.Router.HandleFunc("/product/{id:[0-9]+}", a.updateProduct).Methods("PUT")
    a.Router.HandleFunc("/product/{id:[0-9]+}", a.deleteProduct).Methods("DELETE")
    */

    a.Router.HandleFunc("/", a.sayHello).Methods("GET")
    a.Router.HandleFunc("/createshoe", a.createShoe).Methods("POST")
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

func (a *App) createShoe(w http.ResponseWriter, r *http.Request) {
    

    /*shoe := model.Shoe


    if err := model.CreateShoe(a.DB); err != nil {
        respondWithError(w, http.StatusInternalServerError, err.Error())
        return
    }

    respondWithJSON(w, http.StatusCreated, shoe)*/
    
}