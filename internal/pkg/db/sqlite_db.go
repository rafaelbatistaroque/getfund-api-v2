package db

import (
	applog "getfund-api-v2/internal/log"

	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
)

func New() *gorm.DB {
	db, err := gorm.Open(sqlite.Open("../../getfund.db"), &gorm.Config{})
	if err != nil {
		applog.Error.Fatal("Erro ao conectar ao banco de dados:", err)
	}

	applog.Info.Print("Database connected")

	return db
}
