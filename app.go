package main

import (
    "database/sql"
    "log"
    "net/http"
    "fmt"

    "gitlab.com/reneamontes/truesize/config"

    "github.com/gorilla/mux"
    _ "github.com/lib/pq"
)

type App struct {
	Config    *config.Config
    Router *mux.Router
    DB     *sql.DB
}

func (a *App) SetConfig() {
	a.Config = config.GetConfig()	
}

func (a *App) Initialize(host, user, password, dbname string) {
    connectionString :=
        fmt.Sprintf("host=%s user=%s password=%s dbname=%s", host, user, password, dbname)

    var err error
    a.DB, err = sql.Open("postgres", connectionString)
    if err != nil {
        log.Fatal(err)
    }

    a.Router = mux.NewRouter()
    //a.initializeRoutes()
}

func (a *App) initializeRoutes() {
    /*a.Router.HandleFunc("/products", a.getProducts).Methods("GET")
    a.Router.HandleFunc("/product", a.createProduct).Methods("POST")
    a.Router.HandleFunc("/product/{id:[0-9]+}", a.getProduct).Methods("GET")
    a.Router.HandleFunc("/product/{id:[0-9]+}", a.updateProduct).Methods("PUT")
    a.Router.HandleFunc("/product/{id:[0-9]+}", a.deleteProduct).Methods("DELETE")
    */

    a.Router.HandleFunc("/", a.sayHello).Methods("GET")
}

func (a *App) sayHello(w http.ResponseWriter, r *http.Request) {
    fmt.Println("Hola")
}

func (a *App) Run(addr string) {
    log.Fatal(http.ListenAndServe(":8000", a.Router))
}