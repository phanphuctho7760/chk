package database

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"time"
)

var dbBuiltin *sql.DB

var DbGorm *gorm.DB

func Init(dbInfo string) *gorm.DB {
	var err error

	dbBuiltin, err = sql.Open("mysql", dbInfo)
	if err != nil {
		panic("failed to connect mysql database")
	}
	dbBuiltin.SetConnMaxLifetime(3 * time.Minute)
	dbBuiltin.SetMaxOpenConns(100)
	dbBuiltin.SetMaxIdleConns(100)

	DbGorm, err = gorm.Open(mysql.New(mysql.Config{
		Conn: dbBuiltin,
	}), &gorm.Config{})
	if err != nil {
		panic("failed to connect gorm database")
	}

	return DbGorm
}
