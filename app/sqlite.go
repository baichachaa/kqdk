package app

import (
	"fmt"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	ormlogger "gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"log"
	"os"
	"time"
)

func SqliteService() {

	// 开发环境
	log := ormlogger.New(log.New(os.Stdout, "\r\n", log.LstdFlags), ormlogger.Config{
		SlowThreshold:             200 * time.Millisecond,
		LogLevel:                  ormlogger.Info,
		Colorful:                  true,
		IgnoreRecordNotFoundError: false,
	})

	db, err := gorm.Open(sqlite.Open(`cache.db`), &gorm.Config{
		Logger:                                   log,
		PrepareStmt:                              true,
		NamingStrategy:                           schema.NamingStrategy{SingularTable: true},
		QueryFields:                              true,
		DisableForeignKeyConstraintWhenMigrating: true,
		CreateBatchSize:                          1000,
	})
	if err != nil {
		appLogger.Error("sqlite cache.db打开失败")
		appLogger.Error(err.Error())
		fmt.Println(err)
		os.Exit(0)
	}
	_ = db.AutoMigrate(&Record{})
	appSqlite = db

}
