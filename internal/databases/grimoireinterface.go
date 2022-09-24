package databases

import (
	"context"

	"github.com/bradleyGamiMarques/PersonaGrimoire/api"
)

type PersonaGrimoire interface {
	// Check for specific things existing.
	CheckIfArcanaExistsByUUID(ctx context.Context, arcanaUUID api.ArcanaID) (exists bool, err error)
	CheckIfArcanaExistsByName(ctx context.Context, arcanaName api.ArcanaName) (exists bool, err error)

	// CRUDs relating to P5 Arcanas
	// GET
	GetPersona5ArcanaByName(ctx context.Context, arcanaName api.ArcanaName) (arcana api.P5Arcana, err error)
	GetPersona5ArcanaByUUID(ctx context.Context, arcanaUUID api.ArcanaID) (arcana api.P5Arcana, err error)
}
