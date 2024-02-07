package configuration

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func DatabaseConnection(appConfig AppConfig, secrets Secrets) *gorm.DB {

	sqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		appConfig.Database.Host,
		appConfig.Database.Port,
		secrets.Database.Username,
		secrets.Database.Password,
		appConfig.Database.Schema,
	)

	db, err := gorm.Open(postgres.Open(sqlInfo), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	return db
}
