package main

import (
	"fmt"
	"net/http"

	"github.com/bradleyGamiMarques/PersonaGrimoire/api"
	"github.com/bradleyGamiMarques/PersonaGrimoire/internal/databases"
)

const DEFAULT_SERVER_PORT = 5000

func main() {
	c := Initialize()

	c.RegisterComponents()

	handler := api.NewStrictHandler(c, nil)
	api.RegisterHandlers(c.Router, handler)
	c.Logger.Infof("Starting server on port:%d", DEFAULT_SERVER_PORT)
	err := c.Router.Start(fmt.Sprintf(":%d", DEFAULT_SERVER_PORT))
	if err == http.ErrServerClosed {
		c.Logger.Info("Server closed")
	} else if err != nil {
		c.Logger.Errorf("internal server error Error: %s", err.Error())
	}

}

func (c *Container) RegisterComponents() {
	databases.RegisterDatabaseComponents(c.Gorm, c.Logger)
}
