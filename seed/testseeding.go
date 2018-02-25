package seed

import "github.com/renomx/truesize/main"

const databaseCreation = "CREATE DATABASE truesize_test"

const shoesTableCreation = `
	CREATE TABLE IF NOT EXISTS shoes
	(
		id SERIAL NOT NULL,
		name TEXT NOT NULL,
	)`

const sizesTableCreation = `
	CREATE TABLE IF NOT EXISTS sizes
	(
		id SERIAL NOT NULL
		shoe_id INTEGER NOT NULL
		size INTEGER NOT NULL
	)
`
var a main.App

func SeedDataBase() {
	a = main.App{}
	a.SetConfig()

	a.Initialize(
		a.Config.Tests.Host, 
		a.Config.Tests.User, 
		a.Config.Tests.Password, 
		a.Config.Tests.DbName
	)

	ensureDbExists()

}

func ensureDbExists() {
	if _, err := a.DB.Exec(databaseCreation); err != nil {
		log.Fata(err)
	}

	if _, err := a.DB.Exec(shoesTableCreation); err != nil {
		log.Fata(err)
	}

	if _, err := a.DB.Exec(sizesTableCreation); err != nil {
		log.Fata(err)
	}
}