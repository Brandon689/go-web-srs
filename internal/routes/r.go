package routes

import (
	"net/http"
	
	"github.com/labstack/echo/v4"
	//"github.com/labstack/echo/v4/middleware"
	"github.com/Brandon689/animetrack/internal/db"
)

// type Route struct {
// 	Database *db.Db
// }

// func NewRoute(database *db.Db) *Route {
// 	return &Route{Database: database}
// }

// func (r *Route) LLL() string {
// 	return r.Database.Goal()
// }

type Route struct {
}

func NewRoute() *Route {
	return &Route{}
}

func (r *Route) LLL() string {
	return "s"
}

func (r *Route) RegisterRoutes(e *echo.Echo) {
	// Define a route that responds with "hello world"
	e.GET("/", r.helloWorldHandler)
	e.GET("/u", r.a)
	e.GET("/c", r.ct)
}

func (r *Route) helloWorldHandler(c echo.Context) error {
	return c.String(http.StatusOK, "ok")
}

func (r *Route) a(c echo.Context) error {
	return c.JSON(http.StatusOK, db.GetAllProducts())
}

func (r *Route) ct(c echo.Context) error {
	e2 := db.CreateProduct()
	return c.JSON(http.StatusOK, e2)
}

//return c.JSON(http.StatusOK, products)

// type Db struct {
// 	route *routes.Route
// }

// func NewDb(route *routes.Route) *Db {
//     return &Db{route: route}
// }

// func (d *Db) Goal() string {
//     return d.route.LLL()
// }