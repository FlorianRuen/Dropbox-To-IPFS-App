package database

import (
	"github.com/FlorianRuen/Dropbox-To-IPFS-App/backend/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectToDatabase(dbConfig model.DatabaseConfig) *gorm.DB {
	db, err := gorm.Open(postgres.Open(
		"host=" + dbConfig.Host +
			" port=" + dbConfig.Port +
			" user=" + dbConfig.User +
			" dbname=" + dbConfig.Name +
			" password=" + dbConfig.Password +
			" sslmode=" + dbConfig.SslMode))

	if err != nil {
		panic(err)
	}

	return db
}
