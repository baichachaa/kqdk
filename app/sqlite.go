package app

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	ormlogger "gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"log"
	"os"
	"time"
)

// github.com/mattn/go-sqlite3

type Record struct {
	RecordId       int       `gorm:"column:Record_ID;primaryKey"`              // 记录号
	Name           string    `gorm:"column:name;type:varchar(50)"`             // 姓名--
	Sex            string    `gorm:"column:Sex;type:varchar(10)"`              // 性别--
	IdentityNo     string    `gorm:"column:IdentityNo;type:varchar(50)"`       // 人资编码 身份证号码格式的去除--
	DepartMentName string    `gorm:"column:DepartMentName;type:varchar(200)"`  // -- 部门名称--
	DeviceInout    int       `gorm:"column:Device_InOut;type:int"`             // 出入标识 1 入0 出--
	DeviceLocation string    `gorm:"column:Device_Location;type:varchar(100)"` // 出入位置--
	AuthTime       time.Time `gorm:"column:AuthTime;type:datetime"`            // 出入时间--
}

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
		panic(err)
	}
	_ = db.AutoMigrate(&Record{})
	appSqlite = db

}
