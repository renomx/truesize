package main

func main() {
	// Initialize App
	a := App{}
	a.SetConfig()

    a.Initialize(
    	a.Config.Local.Host, 
    	a.Config.Local.User, 
    	a.Config.Local.Password, 
    	a.Config.Local.DbName)

    a.Run(a.Config.Local.ApiPort)
}