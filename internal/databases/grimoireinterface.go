package databases

import (
	"context"

	"github.com/google/uuid"
)

type PersonaGrimoire interface {
	CheckIfArcanaExistsByUUID(ctx context.Context, arcanaID uuid.UUID) (exists bool, err error)
}
