package model

type Shoe struct {
	Id             int     `json:"id"`
	Name           string  `json:"name"`
	TrueToSizeData []int   `json:"truetosizedata"`
}


func (s *Shoe) getShoes(db *sql.DB) ([]Shoe, error) {
	return errors.New("not implemented")
}

func (s *Shoe) getShoe(db *sql.DB) error {

}

func (s *Shoe) createShoe(db *sql.DB) error {
	return errors.New("not implemented")
}

func (s *Show) addTrueToSize(db *sql.DB) error {
	return errors.New("not implemented")
}


func (s *Show) deleteShoe(db *sql.DB) error {

}
