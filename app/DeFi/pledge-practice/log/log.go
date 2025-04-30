package log

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
	"os"
	"path/filepath"
	"runtime"
	"time"
)

/*
*
1. 添加钩子函数，生成系统日志，备份以及文件名
*/
var Logger *zap.Logger

func init() {
	//1.
	//initlog1()
	//2.
	initLog2()
}

func initLog2() {
	config := zap.NewProductionConfig()
	config.Development = true
	config.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	hook := &lumberjack.Logger{
		Filename:   getCallerDir() + "/log/" + time.Now().Format("2006-01-02 11:55:22") + "access.log",
		MaxAge:     30,
		MaxBackups: 10,
		Compress:   true,
		LocalTime:  true,
		MaxSize:    50,
	}

	core := zapcore.NewCore(zapcore.NewJSONEncoder(config.EncoderConfig), zapcore.NewMultiWriteSyncer(zapcore.AddSync(hook), zapcore.AddSync(os.Stdout)), zapcore.InfoLevel)

	Logger = zap.New(core)
}

func initlog1() {
	//钩子函数
	hook := lumberjack.Logger{
		Filename:   getCallerDir() + "/log/access.log",
		MaxAge:     30,
		MaxBackups: 10,
		Compress:   true,
		LocalTime:  true,
		MaxSize:    50,
	}

	encoderConfig := zapcore.EncoderConfig{
		TimeKey:        "time",
		LevelKey:       "level",
		NameKey:        "logger",
		CallerKey:      "line",
		MessageKey:     "msg",
		StacktraceKey:  "stacktrace",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.LowercaseLevelEncoder,  // 小写编码器
		EncodeTime:     zapcore.ISO8601TimeEncoder,     // ISO8601 UTC 时间格式
		EncodeDuration: zapcore.SecondsDurationEncoder, //
		EncodeCaller:   zapcore.FullCallerEncoder,      // 全路径编码器
		EncodeName:     zapcore.FullNameEncoder,
	}
	core := zapcore.NewCore(zapcore.NewJSONEncoder(encoderConfig), zapcore.NewMultiWriteSyncer(zapcore.AddSync(os.Stderr), zapcore.AddSync(&hook)), zapcore.InfoLevel)

	Logger = zap.New(core, zap.AddCaller(), zap.Development(), zap.Fields(zap.String("service", "cy-defi")))
}

func getCallerDir() string {
	_, file, _, ok := runtime.Caller(0)
	if !ok {
		return ""
	}

	return filepath.Dir(file)
}
