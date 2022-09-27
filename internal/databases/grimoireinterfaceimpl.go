package databases

import (
	"context"
	"errors"
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
	return count > 0, nil
}
func (p *PersonaGrimoireImpl) CheckIfArcanaExistsByName(ctx context.Context, arcanaName api.ArcanaName) (exists bool, err error) {
	var count int64
	err = p.Gorm.WithContext(ctx).Model(&api.P5Arcana{}).Where(&api.P5Arcana{ArcanaName: arcanaName}).Count(&count).Error
	if err != nil {
		return false, fmt.Errorf("failed to count arcanas Error: %w", err)
	}
	return count > 0, nil
}

// CRUDs relating to Persona 5 Arcanas
func (p *PersonaGrimoireImpl) GetAllPersona5Arcanas(ctx context.Context, limit, offset int) (arcanas []api.P5ArcanaResponse, err error) {
	p.Gorm.Model(api.P5Arcana{}).WithContext(ctx).Limit(limit).Offset(offset).Select("ArcanaName", "ArcanaNumber", "ArcanaNumeral").Find(&arcanas)
	return arcanas, nil
}
func (p *PersonaGrimoireImpl) GetPersona5ArcanaByName(ctx context.Context, arcanaName api.ArcanaName) (arcana api.P5ArcanaResponse, err error) {
	err = p.Gorm.WithContext(ctx).Model(&api.P5Arcana{}).Where(&api.P5Arcana{ArcanaName: arcanaName}).Select("ArcanaName", "ArcanaNumber", "ArcanaNumeral").First(&arcana).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			p.Logger.Warnf("Attempted to get Persona 5 Arcana by name that does not exist. Error: %s", err.Error())
			return api.P5ArcanaResponse{}, fmt.Errorf("attempted to get Persona 5 Arcana by name that does not exist Error: %w", err)
		}
	}
	return arcana, nil
}

func (p *PersonaGrimoireImpl) GetPersona5ArcanaByUUID(ctx context.Context, arcanaUUID api.ArcanaID) (arcana api.P5ArcanaResponse, err error) {
	err = p.Gorm.WithContext(ctx).Model(&api.P5Arcana{}).Where(&api.P5Arcana{ArcanaID: arcanaUUID}).Select("ArcanaName", "ArcanaNumber", "ArcanaNumeral").First(&arcana).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			p.Logger.Warnf("Attempted to get Persona 5 Arcana by ID that does not exist. Error: %s", err.Error())
			return api.P5ArcanaResponse{}, fmt.Errorf("attempted to get Persona 5 Arcana by ID that does not exist Error: %w", err)
		}
	}
	return arcana, nil
}
