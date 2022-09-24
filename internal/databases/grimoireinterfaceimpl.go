package databases

import (
	"context"
	"fmt"

	"github.com/bradleyGamiMarques/PersonaGrimoire/api"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type PersonaGrimoireImpl struct {
	Gorm   *gorm.DB
	Logger *logrus.Logger
}

// Check if something exists.
func (p *PersonaGrimoireImpl) CheckIfArcanaExistsByUUID(ctx context.Context, arcanaID uuid.UUID) (exists bool, err error) {
	var count int64
	err = p.Gorm.WithContext(ctx).Model(&api.Arcana{}).Where(&api.Arcana{ArcanaID: arcanaID}).Count(&count).Error
	if err != nil {
		return false, fmt.Errorf("failed to count arcanas Error: %w", err)
	}
	return count > 0, fmt.Errorf("NOT IMPLEMENTED")
}
func (p *PersonaGrimoireImpl) CheckIfArcanaExistsByName(ctx context.Context, arcanaName string) (exists bool, err error) {
	var count int64
	err = p.Gorm.WithContext(ctx).Model(&api.Arcana{}).Where(&api.Arcana{ArcanaName: arcanaName}).Count(&count).Error
	if err != nil {
		return false, fmt.Errorf("failed to count arcanas Error: %w", err)
	}
	return count > 0, fmt.Errorf("NOT IMPLEMENTED")
}
