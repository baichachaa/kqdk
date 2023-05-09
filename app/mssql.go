package app

import (
	"fmt"
	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
	ormlogger "gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"log"
	"moul.io/zapgorm2"
	"os"
	"time"
)

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

// NewService returns a new (SQL) base service for common operations.
func MssqlService() *gorm.DB {

	// 开发环境
	logger := ormlogger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		ormlogger.Config{
			SlowThreshold:             200 * time.Millisecond, // 慢 SQL 阈值
			LogLevel:                  ormlogger.Info,         // Log level
			Colorful:                  true,                   // 彩色打印
			IgnoreRecordNotFoundError: false,                  // 忽略查询为0行的错误
		},
	)

	// 生产环境
	logger2 := zapgorm2.New(appLogger)
	logger2.SetAsDefault() // optional: configure gorm to use this zapgorm.Logger for callbacks
	logger2.SlowThreshold = 200 * time.Millisecond
	logger2.LogLevel = ormlogger.Warn
	logger2.SkipCallerLookup = true
	logger2.IgnoreRecordNotFoundError = true

	gormConfig := gorm.Config{
		Logger:                                   logger2,
		PrepareStmt:                              true,
		NamingStrategy:                           schema.NamingStrategy{SingularTable: true},
		QueryFields:                              true,
		DisableForeignKeyConstraintWhenMigrating: true,
		CreateBatchSize:                          1000,
	}

	if settings.Debug {
		gormConfig.Logger = logger
	}

	dsn := "sqlserver://%s:%s@%s:%d?database=%s&encrypt=disable"
	dsn = fmt.Sprintf(dsn, settings.Mssql.Username, settings.Mssql.Password, settings.Mssql.IP, settings.Mssql.Port, settings.Mssql.Database)
	tdb, err := gorm.Open(sqlserver.Open(dsn), &gormConfig)
	if err != nil {
		fmt.Println("mssql 连接失败")
		fmt.Println(err.Error())
		appLogger.Error("mssql 连接失败")
		appLogger.Error(err.Error())
		fmt.Println(err)
		os.Exit(0)
	}

	// 检查数据库连接情况
	rs := Record{}
	err = tdb.Select("top 1 Record_ID").Find(&rs).Error
	if err != nil {
		fmt.Println("mssql 连接失败")
		fmt.Println(err.Error())
		appLogger.Error("mssql 连接失败")
		appLogger.Error(err.Error())
		fmt.Println(err)
		os.Exit(0)
	}

	appMssql = tdb
	return tdb
}
