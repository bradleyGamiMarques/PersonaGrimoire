package databases

import (
	"github.com/bradleyGamiMarques/PersonaGrimoire/api"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

func RegisterDatabaseComponents(gorm *gorm.DB, logger *logrus.Logger) {
	err := gorm.Exec(`CREATE EXTENSION IF NOT EXISTS "uuid-ossp"`)
	if err.Error != nil {
		logger.Panicf("Error adding extension: %v", err.Error.Error())
	}
	migrationError := gorm.AutoMigrate(&api.P5Arcana{})
	if migrationError != nil {
		logger.Panicf("gorm Migration Error for Arcana: %v", migrationError)
	}

	migrationError = gorm.AutoMigrate(&api.P5Persona{})
	if migrationError != nil {
		logger.Panicf("gorm Migration Error for P5Persona: %v", migrationError)
	}

	migrationError = gorm.AutoMigrate(&api.P5PersonaSkills{})
	if migrationError != nil {
		logger.Panicf("gorm Migration Error for P5PersonaSkill: %v", migrationError)
	}
	migrationError = gorm.AutoMigrate(&api.P5PersonaSkillJunction{})
	if migrationError != nil {
		logger.Panicf("gorm Migration Error for P5PersonaSkillJunction: %v", migrationError)
	}
}
