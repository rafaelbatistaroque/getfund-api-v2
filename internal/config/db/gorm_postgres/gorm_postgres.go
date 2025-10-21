package config_gorm_postgres

import (
	"database/sql"
	"fmt"
	"getfund-api-v2/internal/config/env"
	"getfund-api-v2/internal/infra/db"
	logger "getfund-api-v2/internal/shared/log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func New(env env.Variable) (*db.GetFund, *sql.DB) {
	logger := logger.New("Gorm Postgres config")

	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		env.GetDBHost(),
		env.GetDBPort(),
		env.GetDBUser(),
		env.GetDBPassword(),
		env.GetDBName(),
	)

	gorm_db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		logger.Errorf("Erro ao conectar ao banco de dados: %v", err)
		return nil, nil
	}

	logger.Info("Database connected")

	get_fund_db := &db.GetFund{DB: *gorm_db}

	get_fund_db.AutoMigrate()
	//get_fund_db.Seed()

	opened_db, _ := get_fund_db.DB.DB()

	return get_fund_db, opened_db
}
