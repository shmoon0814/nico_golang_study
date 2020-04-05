package main

import (
	"github.com/labstack/echo"
	"github.com/shmoon0814/learngo/scrapper"
)

func handleHome(c echo.Context) error {
	return c.File("home.html")
}
func handleScapper(c echo.Context) error {
	term := c.FormValue("term")
	scrapper.Scrape(term)
	return nil
}
func main() {
	e := echo.New()
	e.GET("/", handleHome)
	e.POST("/scrape", handleScapper)
	e.Logger.Fatal(e.Start(":8080"))
	//scrapper.Scrape("python")
}
