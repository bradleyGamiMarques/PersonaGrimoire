package main

import (
	"context"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
)

const DEFAULT_LOG_LEVEL = logrus.InfoLevel

type Container struct {
	Router  *echo.Echo
	Logger  *logrus.Logger
	Context context.Context
}

func Initialize() *Container {
	c := &Container{
		Router:  echo.New(),
		Context: context.Background(),
		Logger:  initalizeLogger(),
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
