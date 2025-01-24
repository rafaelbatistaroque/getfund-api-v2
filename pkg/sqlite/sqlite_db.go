package sqlitedb

import (
	logger "getfund-api-v2/pkg/log"

	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
)

func New() *gorm.DB {
	logger := logger.New("Sqlite config")
	db, err := gorm.Open(sqlite.Open("getfund.db"), &gorm.Config{})
	if err != nil {
		logger.Errorf("Erro ao conectar ao banco de dados:", err)
	}

	logger.Info("Database connected")

	return db
}
