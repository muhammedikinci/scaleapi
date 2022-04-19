package main

import (
	"flag"
	"log"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/muhammedikinci/scaleapi/cmd/server/custom_middleware"
	"github.com/muhammedikinci/scaleapi/cmd/server/handler"
	"github.com/muhammedikinci/scaleapi/pkg/api"
	"github.com/muhammedikinci/scaleapi/pkg/repository/gormrepo"
)

func main() {
	dsn := flag.String("dsn", "host=localhost user=postgres password=postgres dbname=scaleflix", "Database Connection String")

	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Llongfile)

	repository, err := gormrepo.NewRepository(errorLog, infoLog, dsn)

	if err != nil {
		panic(err)
	}

	userApi := api.UserAPI{
		Repository: repository.User,
	}

	movieHandler := handler.NewMovieHandler(repository.Movie)
	serieHandler := handler.NewSerieHandler(repository.Serie)
	userHandler := handler.NewUserHandler(userApi)

	e := echo.New()

	e.GET("/movies", movieHandler.GetAllMovies, custom_middleware.UserCheck(userApi))
	e.GET("/movies/:id", movieHandler.FindById, custom_middleware.UserCheck(userApi))
	e.POST("/movies", movieHandler.AddMovie, custom_middleware.AdminCheck(userApi))

	e.GET("/series", serieHandler.GetAllSeries, custom_middleware.UserCheck(userApi))
	e.GET("/series/:id", serieHandler.FindById, custom_middleware.UserCheck(userApi))
	e.POST("/series", serieHandler.AddSerie, custom_middleware.AdminCheck(userApi))

	e.POST("/login", userHandler.Login)
	e.POST("/register", userHandler.Register)

	e.Logger.Fatal(e.Start(":8080"))
}
