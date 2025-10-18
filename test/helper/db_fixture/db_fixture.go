package db_fixture

import (
	"getfund-api-v2/internal/infra/db"

	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func NewMemoryDB(tables ...any) *db.GetFund {
	gorm_db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})
	if err != nil {
		panic("Erro ao conectar ao banco em memória: " + err.Error())
	}

	for _, table := range tables {
		if err := gorm_db.AutoMigrate(table); err != nil {
			panic("Erro ao migrar tabela: " + err.Error())
		}
	}

	return &db.GetFund{DB: *gorm_db}
}
