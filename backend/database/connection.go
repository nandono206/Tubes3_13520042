package database

import (
	"gorm.io/gorm"
	"gorm.io/driver/mysql"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/nandono206/FinalSeasonPart3/backend/models"
)

var DB *gorm.DB
// var rawDB *sql.DB

func Connect(){

	sqlDB, err1 := sql.Open("mysql", "root:@/datatubes3stima")
	if err1 != nil{
		panic("panic")
	}
	// rawDB = sqlDB
	gormDB, err := gorm.Open(mysql.New(mysql.Config{
		Conn: sqlDB,
	  }), &gorm.Config{DisableForeignKeyConstraintWhenMigrating: true,})

	if err != nil{
		panic("panic")
	}

	DB = gormDB

	gormDB.Migrator().CreateTable(&models.Disease{})
	gormDB.Migrator().CreateTable(&models.Record{})
	

}