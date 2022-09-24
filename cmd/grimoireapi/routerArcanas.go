package main

import (
	"context"

	"github.com/bradleyGamiMarques/PersonaGrimoire/api"
)

func (c *Container) GetPersona5ArcanaByName(ctx context.Context, request api.GetPersona5ArcanaByNameRequestObject) (api.GetPersona5ArcanaByNameResponseObject, error) {
	// Check if Arcana exists.
	exists, err := c.PersonaGrimoire.CheckIfArcanaExistsByName(ctx, request.ArcanaName)
	if err != nil {
		c.Logger.Errorf("Error attempting to check if Arcana exists Arcana Name: %s Error: %s", request.ArcanaName, err.Error())
		return api.ServerErrorJSONResponse{}, nil
	}
	if !exists {
		c.Logger.Warnf("Attempted to get Arcana by name that does not exist Arcana Name: %s", request.ArcanaName)
	}
	arcana, err := c.PersonaGrimoire.GetPersona5ArcanaByName(ctx, request.ArcanaName)
	c.Logger.Infof("Arcana: %v", arcana)
	if err != nil {
		c.Logger.Errorf("Error attempting to get Arcana by name Arcana Name: %s Error %s", request.ArcanaName, err.Error())
		return api.ServerErrorJSONResponse{}, nil
	}
	return api.GetPersona5ArcanaByName200JSONResponse(arcana), nil
}

func (c *Container) GetPersona5ArcanaByUUID(ctx context.Context, request api.GetPersona5ArcanaByUUIDRequestObject) (api.GetPersona5ArcanaByUUIDResponseObject, error) {
	return nil, nil
}
