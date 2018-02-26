package models

import (
	"errors"
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
)

type SimpleShoe struct {
	Id         int     `json:"id"`
	Name       string  `json:"name"`
	Sizes      []int   `json:"sizes"`
	TrueToSize float64 `json:"truetosize"`
}

type Shoe struct {
	gorm.Model
	Name       string
	TrueToSize float64 `sql:"type:decimal(14,13);"`
	Sizes      []Size
}

type Size struct {
	gorm.Model
	Size   int
	ShoeID int
}

func (s *Shoe) InitModel(db *gorm.DB) error {

	db.DropTableIfExists(&Shoe{}, &Size{})

	db.AutoMigrate(&Shoe{}, &Size{})

	log.Println("DB Initialized")

	return nil
}

func (s *Shoe) GetShoes(db *gorm.DB) ([]SimpleShoe, error) {

	simpleShoes := []SimpleShoe{}
	shoes := []Shoe{}
	db.Find(&shoes)
	for _, s := range shoes {
		simpleShoes = append(simpleShoes, mapToSimpleShoe(db, int(s.ID)))

	}
	return simpleShoes, nil
}

func (s *Shoe) GetShoe(db *gorm.DB, shoeName string) (SimpleShoe, error) {

	shoe := Shoe{}
	db.Where("name=?", shoeName).Find(&shoe)
	return mapToSimpleShoe(db, int(shoe.ID)), nil
}

func (s *SimpleShoe) CreateShoe(db *gorm.DB) (SimpleShoe, error) {

	shoe := Shoe{
		Name: s.Name,
	}

	for _, element := range s.Sizes {
		size := Size{
			Size: element,
		}
		shoe.Sizes = append(shoe.Sizes, size)
	}

	db.Create(&shoe)

	return mapToSimpleShoe(db, int(shoe.ID)), nil
}

func (s *Shoe) AddTrueToSize(db *gorm.DB, shoeName string, size int) (SimpleShoe, error) {

	shoe := Shoe{}
	db.Where("name = ?", shoeName).Find(&shoe)

	newSize := Size{
		Size:   size,
		ShoeID: int(shoe.ID),
	}

	db.Create(&newSize)

	simpleShoe := mapToSimpleShoe(db, int(shoe.ID))

	shoe.TrueToSize = simpleShoe.TrueToSize

	db.Save(&shoe)

	return simpleShoe, nil
}

func (s *Shoe) deleteShoe(db *gorm.DB) error {
	return errors.New("not implemented")
}

func (s *Shoe) calculateTrueToSize(db *gorm.DB) error {
	return errors.New("not implemented")
}

func mapToSimpleShoe(db *gorm.DB, id int) SimpleShoe {

	simpleShoe := SimpleShoe{}
	shoe := Shoe{}

	db.Where("id = ?", id).Find(&shoe)

	sizes := []Size{}
	db.Find(&sizes).Where("shoe_id = ?", id)

	total := float64(len(sizes))
	count := 0

	for _, item := range sizes {
		if item.ShoeID == id {
			simpleShoe.Sizes = append(simpleShoe.Sizes, item.Size)
			count += item.Size
		}
	}

	simpleShoe.Id = id
	simpleShoe.Name = shoe.Name
	simpleShoe.TrueToSize = float64(count) / total

	return simpleShoe

}
