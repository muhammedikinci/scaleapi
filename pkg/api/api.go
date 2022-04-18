package api

import (
	"database/sql"
	"log"

	"github.com/muhammedikinci/scaleapi/pkg/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type api struct {
	errorLog *log.Logger
	infoLog  *log.Logger
	db       *gorm.DB
}

func NewApi(errorLog *log.Logger, infoLog *log.Logger, dsn *string) (*api, error) {
	db, err := openDB(*dsn)

	if err != nil {
		return nil, err
	}

	db.AutoMigrate(&models.Genre{}, &models.Movie{}, &models.Serie{}, &models.Season{}, &models.Episode{}, &models.User{})

	return &api{
		errorLog: errorLog,
		infoLog:  infoLog,
		db:       db,
	}, nil
}

func openDB(dsn string) (*gorm.DB, error) {
	db, err := sql.Open("pgx", dsn)

	if err != nil {
		return nil, err
	}

	if err = db.Ping(); err != nil {
		return nil, err
	}

	gormDB, err := gorm.Open(postgres.New(postgres.Config{
		Conn: db,
	}), &gorm.Config{})

	if err != nil {
		return nil, err
	}

	return gormDB, nil
}
