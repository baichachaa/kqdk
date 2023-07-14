package app

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	ormlogger "gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"log"
	"moul.io/zapgorm2"
	"os"
	"time"
)

type Record struct {
	RecordId       string    `gorm:"column:-"`                                 // 记录号
	Name           string    `gorm:"column:employee_name;type:varchar(50)"`    // 姓名--
	IdentityNo     string    `gorm:"column:card_no;type:varchar(50)"`          // 人资编码 身份证号码格式的去除--
	DepartMentName string    `gorm:"column:department_name;type:varchar(200)"` // -- 部门名称--
	DeviceInout    string    `gorm:"column:device_name;type:varchar(200)"`     // 出入标识 1 入0 出--
	AuthTime       time.Time `gorm:"column:snap_time;type:datetime"`           // 出入时间--
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

	dsn := "%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local&tls=preferred"
	dsn = fmt.Sprintf(dsn, settings.DataBase.Username, settings.DataBase.Password, settings.DataBase.IP, settings.DataBase.Port, settings.DataBase.Database)
	tdb, err := gorm.Open(mysql.Open(dsn), &gormConfig)

	if err != nil {
		fmt.Println("mysql 连接失败")
		fmt.Println(err.Error())
		appLogger.Error("mysql 连接失败")
		appLogger.Error(err.Error())
		fmt.Println(err)
		if *isTest {
			os.Exit(0)
		}
	}

	// 检查数据库连接情况
	rs := Record{}
	err = tdb.Select("employee_name").Table("aitcp_employee_t").Limit(1).Scan(&rs).Error
	if err != nil {
		fmt.Println("mysql 连接失败")
		fmt.Println(err.Error())
		appLogger.Error("mssql 连接失败")
		appLogger.Error(err.Error())
		fmt.Println(err)
		if *isTest {
			os.Exit(0)
		}
	}

	appMysql = tdb
	return tdb
}
