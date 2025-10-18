package config_sqlite

import (
	logger "getfund-api-v2/internal/shared/log"

	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
)

func New() *gorm.DB {
	logger := logger.New("Sqlite config")
	db, err := gorm.Open(sqlite.Open("getfund.db"), &gorm.Config{})
	if err != nil {
		logger.Errorf("Erro ao conectar ao banco de dados:", err)
		return nil
	}

	logger.Info("Database connected")

	return db
}
