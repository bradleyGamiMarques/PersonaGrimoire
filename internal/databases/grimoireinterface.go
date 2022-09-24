package databases

import (
	"context"

	"github.com/google/uuid"
)

type PersonaGrimoire interface {
	CheckIfArcanaExistsByUUID(ctx context.Context, arcanaID uuid.UUID) (exists bool, err error)
	CheckIfArcanaExistsByName(ctx context.Context, arcanaName string) (exists bool, err error)
}
