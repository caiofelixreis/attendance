package database

import (
	"fmt"

	"gorm.io/gorm"
)

type ORMConfig struct {
	Name      string
	Dialector gorm.Dialector
	Config    *gorm.Config
}

type DatabaseConnection struct {
	db        *gorm.DB
	ORMConfig *ORMConfig
}

func (dbConn *DatabaseConnection) Init() (*gorm.DB, error) {
	gormDb, err := gorm.Open(dbConn.ORMConfig.Dialector, dbConn.ORMConfig.Config)

	if err != nil {
		return nil, err
	}

	fmt.Println("✅ Connected to database:", dbConn.ORMConfig.Name)

	dbConn.db = gormDb

	return dbConn.db, nil
}

func (dbConn *DatabaseConnection) Close() error {
	sqlDb, err := dbConn.db.DB()

	if err != nil {
		return err
	}

	sqlDb.Close()

	fmt.Println("❌ Closed connection to database:", dbConn.ORMConfig.Name)

	return nil
}

func (dbConn *DatabaseConnection) PrintSchema() {
	// var tables []struct {
	// 	TableName string `gorm:"column:table_name"`
	// }

	// dbConn.db.Raw("SELECT table_name FROM information_schema.tables WHERE table_schema = 'public'").Scan(&tables)
	// for _, table := range tables {
	// 	fmt.Printf("Table: %s\n", table.TableName)

	// 	var columns []struct {
	// 		ColumnName string `gorm:"column:column_name"`
	// 		DataType   string `gorm:"column:data_type"`
	// 	}

	// 	dbConn.db.Raw("SELECT column_name, data_type FROM information_schema.columns WHERE table_name = ?", table.TableName).Scan(&columns)
	// 	for _, column := range columns {
	// 		fmt.Printf("\tColumn: %s - %s\n", column.ColumnName, column.DataType)
	// 	}
	// }

	var students []struct {
		ID string `gorm:"column:id"`
	}

	dbConn.db.Raw("SELECT * FROM student_models").Scan(&students)

	for _, student := range students {
		fmt.Printf("Student: %s\n", student.ID)
	}

}

func (dbConn *DatabaseConnection) GetConn() *gorm.DB {
	return dbConn.db
}

var Postgres = &DatabaseConnection{
	ORMConfig: &PostgresOrmConfig.ORMConfig,
}
