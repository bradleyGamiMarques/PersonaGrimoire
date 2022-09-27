package main

import (
	"context"
	"fmt"

	"github.com/bradleyGamiMarques/PersonaGrimoire/api"
)

const GET_ALL_PERSONA_5_ARCANAS_PAGINATION_LIMIT_DEFAULT = 22
const GET_ALL_PERSONA_5_ARCANAS_PAGINATION_LIMIT_MAXIMUM = 22

func (c *Container) GetAllPersona5Arcanas(ctx context.Context, request api.GetAllPersona5ArcanasRequestObject) (api.GetAllPersona5ArcanasResponseObject, error) {
	limit, offset := GET_ALL_PERSONA_5_ARCANAS_PAGINATION_LIMIT_DEFAULT, 0
	// Check if paginated data was request and is in bounds.
	if request.Params.Limit != nil {
		if *request.Params.Limit <= 0 || *request.Params.Limit > GET_ALL_PERSONA_5_ARCANAS_PAGINATION_LIMIT_MAXIMUM {
			c.Logger.Warnf("User attempted to get data out of bounds. Limit must be greater than zero and less than %d. Limit: %s", GET_ALL_PERSONA_5_ARCANAS_PAGINATION_LIMIT_MAXIMUM, *request.Params.Limit)
			return api.BadRequestJSONResponse{
				Code:    400,
				Error:   true,
				Message: fmt.Sprintf("Limit must be between 0 and %d", GET_ALL_PERSONA_5_ARCANAS_PAGINATION_LIMIT_MAXIMUM),
			}, nil
		}
		limit = *request.Params.Limit
		c.Logger.Infof("Limit: %d", limit)
	}
	if request.Params.Offset != nil {
		if *request.Params.Offset < 0 {
			c.Logger.Warnf("User attempted to get data out of bounds. Offset must be greater than or equal to zero. Offset: %s", *request.Params.Offset)
			return api.BadRequestJSONResponse{
				Code:    400,
				Error:   true,
				Message: "Offset must be 0 or larger ",
			}, nil
		}
		offset = *request.Params.Offset
		c.Logger.Infof("Limit: %d", offset)
	}
	// Get paginated list of Persona 5 arcanas.
	data, err := c.PersonaGrimoire.GetAllPersona5Arcanas(ctx, limit, offset)
	if err != nil {
		c.Logger.Errorf("Failed to get Persona 5 Arcanas. Limit: %d Offset: %d Error: %s", limit, offset, err.Error())
		return api.ServerErrorJSONResponse{
			Code:    500,
			Error:   true,
			Message: "Failed to get Persona 5 Arcanas.",
		}, nil
	}
	// Return a 200 response with an array of Persona 5 Arcanas.
	return api.GetAllPersona5Arcanas200JSONResponse(data), nil
}
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
