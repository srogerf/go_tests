package main

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	_ "github.com/srogerf/go_tests/test1/errors"
	_ "github.com/srogerf/go_tests/test1/files"
	"github.com/srogerf/go_tests/test1/loops"
	"github.com/srogerf/go_tests/test1/structs"
	"github.com/srogerf/meta_fix/file_scan"
	"log"
	"net/http"
	"os"
)

func init() {
	log.SetOutput(os.Stdout)
}

func main() {
	log.Println("Test programs")

	loops.Run()
	structs.Run()

	file_scan.List("x:/video/downloaded/television/")
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Routes
	e.GET("/", hello)

	// Start server
	//	e.Logger.Fatal(e.Start(":1323"))
}

// Handler
func hello(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}
