package models

import (
    "database/sql"
    "errors"
    "log"

    _ "github.com/lib/pq"
)

type Shoe struct {
	Id     int     
	Name   string  
	Sizes  []int
}


func (s *Shoe) InitModel(db *sql.DB) error {
    
    // Add initialization scripts as necesary
    qs := []string {
    "CREATE TABLE IF NOT EXISTS shoes (id serial, name text, sizes integer[])",
    "INSERT INTO shoes ('Adidas 1', '{1,4,3,2,5,2}')",
    }

    for _, q := range qs {
        _, err := db.Exec(q)
        if err != nil {
            panic(err)
            return err
        }
    }

    log.Println("DB Initialized")
    return nil
}


func (s *Shoe) GetShoes(db *sql.DB, start, count int) ([]Shoe, error) {
	  rows, err := db.Query(
        "SELECT id, name, sizes FROM shoes LIMIT $1 OFFSET $2",
        count, start)

    if err != nil {
        return nil, err
    }

    defer rows.Close()

    shoes := []Shoe{}

    for rows.Next() {
        var s Shoe
        if err := rows.Scan(&s.Id, &s.Name, &s.Sizes); err != nil {
            return nil, err
        }
        shoes = append(shoes, s)
    }

    return shoes, nil
}

func (s *Shoe) getShoe(db *sql.DB) error {
    return errors.New("not implemented")
}

func (s *Shoe) CreateShoe(db *sql.DB) error {
 
    


	/*err := db.QueryRow(
        "INSERT INTO shoes(name) VALUES($1) RETURNING id",
        s.Name).Scan(&s.Id)

    if err != nil {
        return err
    }

    i := 0
    for range s.TrueToSizeData {
    	err := db.QueryRow(
    		"INSERT INTO sizes(shoe_id, size) VALUES($1,$2) RETURNING id",
    		s.Id, s.TrueToSizeData[i])
    }*/


    return nil
}

func (s *Shoe) addTrueToSize(db *sql.DB) error {
	return errors.New("not implemented")
}


func (s *Shoe) deleteShoe(db *sql.DB) error {
    return errors.New("not implemented")
}
