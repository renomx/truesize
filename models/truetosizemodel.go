package model

import "log"

type Shoe struct {
	Id     int     
	Name   string  
	Sizes  []int
}


func InitModel(db *sql.DB) error {
    
    qs := []string {
    "CREATE DATABASE truesize",
    "CREATE TABLE IF NOT EXISTS shoes (id serial, name text, sizes integer[])"
    }

    for _, q := range qs {
        _, err := db.Exec(q)
        if err != nil {
            panic(err)
            return err
        }
    }
    return nil
}


func (s *Shoe) getShoes(db *sql.DB) ([]Shoe, error) {
	err := db.Model(&Shoe)
    Column("shoe.*", "Sizes")
    Relation("Sizes", func(q *orm.Query) (*orm.Query, error) {
        return q, nil
    })

    if(err != nil) {
        panic(err)
    }
}

func (s *Shoe) getShoe(db *sql.DB) error {

}

func (s *Shoe) createShoe(db *sql.DB) error {

    log.Println(s)


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

func (s *Show) addTrueToSize(db *sql.DB) error {
	return errors.New("not implemented")
}


func (s *Show) deleteShoe(db *sql.DB) error {

}
