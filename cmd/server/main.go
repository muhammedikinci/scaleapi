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

	movieApi := api.MovieAPI{
		Repository: repository.Movie,
	}

	userApi := api.UserAPI{
		Repository: repository.User,
	}

	serieApi := api.SerieAPI{
		Repository: repository.Serie,
	}

	seasonApi := api.SeasonAPI{
		Repository: repository.Season,
	}

	movieHandler := handler.NewMovieHandler(movieApi)
	serieHandler := handler.NewSerieHandler(serieApi)
	userHandler := handler.NewUserHandler(userApi)
	seasonHandler := handler.NewSeasonHandler(seasonApi)

	e := echo.New()

	e.GET("/movies", movieHandler.GetAllMovies, custom_middleware.UserCheck(userApi))
	e.GET("/movies/filter", movieHandler.Filter, custom_middleware.UserCheck(userApi))
	e.GET("/movies/:id", movieHandler.FindById, custom_middleware.UserCheck(userApi))
	e.POST("/movies", movieHandler.AddMovie, custom_middleware.AdminCheck(userApi))

	e.GET("/series", serieHandler.GetAllSeries, custom_middleware.UserCheck(userApi))
	e.GET("/series/filter", serieHandler.Filter, custom_middleware.UserCheck(userApi))
	e.GET("/series/:id", serieHandler.FindById, custom_middleware.UserCheck(userApi))
	e.POST("/series", serieHandler.AddSerie, custom_middleware.AdminCheck(userApi))

	e.GET("/seasons/:id", seasonHandler.FindById, custom_middleware.UserCheck(userApi))
	e.GET("/series/:serieId/seasons/:id", seasonHandler.FindById, custom_middleware.UserCheck(userApi))
	e.GET("/series/:serieId/seasons", seasonHandler.FindAllSeasonsInSerie, custom_middleware.UserCheck(userApi))
	e.POST("/seasons", seasonHandler.AddSeason, custom_middleware.AdminCheck(userApi))

	e.POST("/login", userHandler.Login)
	e.POST("/register", userHandler.Register)

	e.Logger.Fatal(e.Start(":8080"))
}
