package db

import (
	"gorm.io/gorm"
	"gorm.io/driver/sqlite"
)


type Product struct {
	gorm.Model
	//ID     string `gorm:"type:uuid;primary_key" json:"id"`
	Code  string	`json:"text"`
	Price uint	`json:"integer"`
}

var database *gorm.DB

func InitDB() error {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		return err
	}
	db.AutoMigrate(&Product{})
	database = db
	return nil
}

func GetProduct() Product {
	var product Product
	database.First(&product, 1)
	return product
}

func GetAllProducts() []Product {
	var products []Product
	database.Find(&products)
	return products
}

func CreateProduct() Product {
	e1 := &Product{Code: "D42", Price: 100}
	database.Create(e1)
	return *e1
}
