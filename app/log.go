package app

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
	"os"
)

func LogService() {
	hook := lumberjack.Logger{
		Filename:   "./logs/app.log",
		MaxSize:    10,  // 每个日志文件保存的最大尺寸 单位：M
		MaxBackups: 30,  // 日志文件最多保存多少个备份
		MaxAge:     365, // 文件最多保存多少天
		Compress:   false,
	}

	config := zap.NewDevelopmentEncoderConfig()
	config.EncodeTime = zapcore.TimeEncoderOfLayout("2006-01-02 15:04:05")
	encoder := zapcore.NewJSONEncoder(config)

	multiWriteSyncer := zapcore.NewMultiWriteSyncer(zapcore.AddSync(os.Stdout), zapcore.AddSync(&hook))

	core := zapcore.NewCore(encoder, multiWriteSyncer, zap.DebugLevel)
	appLogger = zap.New(core)

	//err := appLogger.Sync()
	//if err != nil {
	//	fmt.Println("日志打开失败")
	//	panic(err)
	//}

}
