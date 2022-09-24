package databases

import (
	"context"
	"fmt"

	"github.com/bradleyGamiMarques/PersonaGrimoire/api"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type PersonaGrimoireImpl struct {
	Gorm   *gorm.DB
	Logger *logrus.Logger
}

// Check if something exists.
func (p *PersonaGrimoireImpl) CheckIfArcanaExistsByUUID(ctx context.Context, arcanaUUID api.ArcanaID) (exists bool, err error) {
	var count int64
	err = p.Gorm.WithContext(ctx).Model(&api.P5Arcana{}).Where(&api.P5Arcana{ArcanaID: arcanaUUID}).Count(&count).Error
	if err != nil {
		return false, fmt.Errorf("failed to count arcanas Error: %w", err)
	}
	return count > 0, fmt.Errorf("NOT IMPLEMENTED")
}
func (p *PersonaGrimoireImpl) CheckIfArcanaExistsByName(ctx context.Context, arcanaName api.ArcanaName) (exists bool, err error) {
	var count int64
	err = p.Gorm.WithContext(ctx).Model(&api.P5Arcana{}).Where(&api.P5Arcana{ArcanaName: arcanaName}).Count(&count).Error
	if err != nil {
		return false, fmt.Errorf("failed to count arcanas Error: %w", err)
	}
	return count > 0, fmt.Errorf("NOT IMPLEMENTED")
}

// CRUDs relating to Persona 5 Arcanas
func (p *PersonaGrimoireImpl) GetPersona5ArcanaByName(ctx context.Context, arcanaName api.ArcanaName) (arcana api.P5Arcana, err error) {
	return api.P5Arcana{}, fmt.Errorf("NOT IMPLEMENTED")
}

func (p *PersonaGrimoireImpl) GetPersona5ArcanaByUUID(ctx context.Context, arcanaUUID api.ArcanaID) (arcana api.P5Arcana, err error) {
	return api.P5Arcana{}, fmt.Errorf("NOT IMPLEMENTED")
}
