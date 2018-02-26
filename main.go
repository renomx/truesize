package main

import "log"

func main() {
	// Initialize App
	a := App{}
	a.SetConfig()

	log.Println(a.Config)

    a.Initialize(
    	a.Config.Local.Host, 
    	a.Config.Local.User, 
    	a.Config.Local.Password, 
    	a.Config.Local.DbName,
    	a.Config.Local.DbPort)

    a.Run(a.Config.Local.ApiPort)
}