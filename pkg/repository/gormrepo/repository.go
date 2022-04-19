package gormrepo

import (
	"database/sql"
	"log"

	"github.com/muhammedikinci/scaleapi/pkg/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type repository struct {
	Movie   movieRepository
	Serie   serieRepository
	User    userRepository
	Season  seasonRepository
	Episode episodeRepository
}

func NewRepository(errorLog *log.Logger, infoLog *log.Logger, dsn *string) (*repository, error) {
	db, err := openDB(*dsn)

	if err != nil {
		return nil, err
	}

	db.AutoMigrate(&models.Movie{}, &models.Serie{}, &models.Season{}, &models.Episode{}, &models.User{})

	db.Session(&gorm.Session{FullSaveAssociations: true})

	return &repository{
		Movie: movieRepository{
			errorLog: errorLog,
			db:       db,
		},
		Serie: serieRepository{
			errorLog: errorLog,
			db:       db,
		},
		User: userRepository{
			errorLog: errorLog,
			db:       db,
		},
		Season: seasonRepository{
			errorLog: errorLog,
			db:       db,
		},
		Episode: episodeRepository{
			errorLog: errorLog,
			db:       db,
		},
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
