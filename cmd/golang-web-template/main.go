package main

import (
	//"fmt"

	"github.com/labstack/echo/v4"
	"github.com/Brandon689/golang-web-template/internal/routes"
	"github.com/Brandon689/golang-web-template/internal/db"
)

func main() {


	err := db.InitDB()
	if err != nil {
		panic("failed to connect database")
	}

	//  sqlDB, _ := dbt.DB.DB()
	//  defer sqlDB.Close()




	db.CreateProduct()
	db.CreateProduct()
	db.CreateProduct()
	db.CreateProduct()





	
	// m := db.NewDb()
	// fmt.Println(m.Goal())

	//c := routes.NewRoute(m)
	c := routes.NewRoute()
	e := echo.New()

	// Register routes from another file
	c.RegisterRoutes(e)
	//fmt.Println(c.LLL())

	// Start the server on port 8080
	e.Start(":8080")

}
