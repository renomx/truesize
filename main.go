package main

import (
	"log"
)

func main() {
	// Initialize App
	a := App{}
	a.SetConfig()

	log.Println(a.Config.Local.Host)
    //a.Initialize(c.Host, c.User, c.Password, c.DbName)

    //a.Run(":8080")
}