package database

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type PostgresORMConfig struct {
	ORMConfig
	dbUrl string
}

func (psg *PostgresORMConfig) initDialector() {
	psg.Dialector = postgres.Open(psg.dbUrl)
}

func (psg *PostgresORMConfig) initConfig() {
	psg.Config = &gorm.Config{
		// Logger: logger.Default.LogMode(logger.Info),
	}
}

func newPostgresORM() *PostgresORMConfig {
	postgresORM := &PostgresORMConfig{
		dbUrl: "postgres://postgres:supersafe@localhost:5432/capoeira_db",
		ORMConfig: ORMConfig{
			Name: "postgres",
		},
	}

	postgresORM.initConfig()
	postgresORM.initDialector()

	return postgresORM
}

var PostgresOrmConfig *PostgresORMConfig = newPostgresORM()
