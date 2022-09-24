package databases

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type PersonaGrimoireImpl struct {
	Gorm   *gorm.DB
	Logger *logrus.Logger
}

func (p *PersonaGrimoireImpl) CheckIfArcanaExists(ctx context.Context, arcanaID uuid.UUID) (exists bool, err error) {
	return false, fmt.Errorf("NOT IMPLEMENTED")
}
