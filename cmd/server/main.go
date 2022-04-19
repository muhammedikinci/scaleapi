package main

import (
	"flag"
	"log"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/muhammedikinci/scaleapi/cmd/server/handler"
	"github.com/muhammedikinci/scaleapi/pkg/api"
)

func main() {
	dsn := flag.String("dsn", "host=localhost user=postgres password=postgres dbname=scaleflix", "Database Connection String")

	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Llongfile)

	api, err := api.NewApi(errorLog, infoLog, dsn)

	if err != nil {
		panic(err)
	}

	movieHandler := handler.NewMovieHandler(api)

	e := echo.New()

	e.GET("/movies", movieHandler.GetAllMovies)
	e.GET("/movies/:id", movieHandler.FindById)
	e.POST("/movies", movieHandler.AddMovie)

	e.Logger.Fatal(e.Start(":1323"))
}
