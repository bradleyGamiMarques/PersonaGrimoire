package databases

import (
	"github.com/bradleyGamiMarques/PersonaGrimoire/api"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

func RegisterDatabaseComponents(gorm *gorm.DB, logger *logrus.Logger) {
	migrationError := gorm.AutoMigrate(&api.P5Arcana{})
	if migrationError != nil {
		logger.Panicf("gorm Migration Error for Arcana: %v", migrationError)
	}

	migrationError = gorm.AutoMigrate(&api.P5Persona{})
	if migrationError != nil {
		logger.Panicf("gorm Migration Error for P5Persona: %v", migrationError)
	}

	migrationError = gorm.AutoMigrate(&api.P5PersonaSkill{})
	if migrationError != nil {
		logger.Panicf("gorm Migration Error for P5PersonaSkill: %v", migrationError)
	}

	migrationError = gorm.AutoMigrate(&api.P5PersonaStats{})
	if migrationError != nil {
		logger.Panicf("gorm Migration Error for P5PersonaSkill: %v", migrationError)
	}
}
