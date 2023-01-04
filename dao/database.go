package dao

import (
	"findings/model"
	"fmt"
	"sync"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var (
	db   *gorm.DB
	err  error
	lock = sync.Mutex{}
)

type IDatabase interface {
	GetConnection() *gorm.DB
}

type database struct {
}

func NewDatabaseInstance() *database {
	return &database{}
}

func (dbConfig *database) GetConnection() *gorm.DB {
	if db == nil {
		lock.Lock()
		defer lock.Unlock()
		if db == nil {
			fmt.Println("Creating db connection !!")
			// dsn := "root:admin123@tcp(127.0.0.1:3306)/findings?charset=utf8mb4&parseTime=True&loc=Local"
			// db, err = gorm.Open(mysql.New(mysql.Config{
			// 	DSN:                       dsn,   // data source name
			// 	DefaultStringSize:         256,   // default size for string fields
			// 	DisableDatetimePrecision:  true,  // disable datetime precision, which not supported before MySQL 5.6
			// 	DontSupportRenameIndex:    true,  // drop & create when rename index, rename index not supported before MySQL 5.7, MariaDB
			// 	DontSupportRenameColumn:   true,  // `change` when rename column, rename column not supported before MySQL 8, MariaDB
			// 	SkipInitializeWithVersion: false, // auto configure based on currently MySQL version
			// }), &gorm.Config{})
			db, err = gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
			if err != nil {
				panic("failed to connect database")
			}
			db.AutoMigrate(&model.Repository{}, &model.ScanDetail{}, &model.Finding{})
		}
	} else {
		fmt.Println("Using already created db connection !!")
	}
	return db
}
