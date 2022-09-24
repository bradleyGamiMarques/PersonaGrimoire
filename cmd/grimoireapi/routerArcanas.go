package main

import (
	"context"
	"fmt"

	"github.com/bradleyGamiMarques/PersonaGrimoire/api"
)

func (c *Container) GetPersona5ArcanaByName(ctx context.Context, request api.GetPersona5ArcanaByNameRequestObject) (api.GetPersona5ArcanaByNameResponseObject, error) {
	// Check if Arcana exists.
	exists, err := c.PersonaGrimoire.CheckIfArcanaExistsByName(ctx, request.ArcanaName)
	if err != nil {
		c.Logger.Errorf("Error attempting to check if Arcana exists. Arcana Name: %s Error: %s", request.ArcanaName, err.Error())
		return api.ServerErrorJSONResponse{Code: 500, Error: true, Message: "Error attempting to check if Arcana exists."}, nil
	}
	if !exists {
		c.Logger.Warnf("Attempted to get Arcana by name that does not exist. Arcana Name: %s", request.ArcanaName)
		return api.NotFoundJSONResponse{Code: 404, Error: true, Message: fmt.Sprintf("Attempted to get Arcana by Name that does not exist. Arcana Name: %s", request.ArcanaName)}, nil
	}
	arcana, err := c.PersonaGrimoire.GetPersona5ArcanaByName(ctx, request.ArcanaName)
	if err != nil {
		c.Logger.Errorf("Error attempting to get Arcana by name. Arcana Name: %s Error %s", request.ArcanaName, err.Error())
		return api.ServerErrorJSONResponse{Code: 500, Error: true, Message: fmt.Sprintf("Error attempting to get Arcana by Name. Arcana Name: %s", request.ArcanaName)}, nil
	}
	return api.GetPersona5ArcanaByName200JSONResponse(arcana), nil
}

func (c *Container) GetPersona5ArcanaByUUID(ctx context.Context, request api.GetPersona5ArcanaByUUIDRequestObject) (api.GetPersona5ArcanaByUUIDResponseObject, error) {
	// Check if Arcana exists.
	exists, err := c.PersonaGrimoire.CheckIfArcanaExistsByUUID(ctx, request.ArcanaUUID)
	if err != nil {
		c.Logger.Errorf("Error attempting to check if Arcana exists. Arcana ID: %s Error: %s", request.ArcanaUUID, err.Error())
		return api.ServerErrorJSONResponse{Code: 500, Error: true, Message: "Error attempting to check if Arcana exists."}, nil
	}
	if !exists {
		c.Logger.Warnf("Attempted to get Arcana by ID that does not exist. Arcana ID: %s", request.ArcanaUUID)
		return api.NotFoundJSONResponse{Code: 404, Error: true, Message: fmt.Sprintf("Attempted to get Arcana by ID that does not exist. Arcana ID: %s", request.ArcanaUUID)}, nil
	}
	arcana, err := c.PersonaGrimoire.GetPersona5ArcanaByUUID(ctx, request.ArcanaUUID)
	if err != nil {
		c.Logger.Errorf("Error attempting to get Arcana by ID. Arcana ID: %s Error %s", request.ArcanaUUID, err.Error())
		return api.ServerErrorJSONResponse{Code: 500, Error: true, Message: fmt.Sprintf("Error attempting to get Arcana by ID. Arcana ID: %s", request.ArcanaUUID)}, nil
	}
	return api.GetPersona5ArcanaByName200JSONResponse(arcana), nil
}
