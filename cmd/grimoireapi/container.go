package main

import (
	"context"
	"os"

	"github.com/bradleyGamiMarques/PersonaGrimoire/internal/databases"
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

const DEFAULT_LOG_LEVEL = logrus.InfoLevel

type Container struct {
	Router          *echo.Echo
	Logger          *logrus.Logger
	Context         context.Context
	Gorm            *gorm.DB
	PersonaGrimoire databases.PersonaGrimoire
}

func Initialize() *Container {
	c := &Container{
		Context: context.Background(),
		Logger:  initalizeLogger(),
		Router:  initializeRouter(),
		Gorm:    initializeGorm(),
	}
	c.PersonaGrimoire = &databases.PersonaGrimoireImpl{
		Gorm:   c.Gorm,
		Logger: c.Logger,
	}
	return c
}

func initalizeLogger() *logrus.Logger {
	logger := logrus.New()
	logger.SetReportCaller(true)
	logger.SetLevel(DEFAULT_LOG_LEVEL)
	logger.SetFormatter(&logrus.TextFormatter{FullTimestamp: true})
	logger.SetOutput(os.Stdout)
	return logger
}
func initializeGorm() *gorm.DB {
	dsn := "host=localhost user=postgres password=postgres dbname=persona_grimoire_db port=5432 sslmode=disable TimeZone=America/Los_Angeles"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		PrepareStmt: true,
	})
	if err != nil {
		panic("Failed to connect to the Postgres database" + err.Error())
	}
	return db
}
func initializeRouter() *echo.Echo {
	router := echo.New()
	router.HideBanner = true
	router.HidePort = true
	return router
}
