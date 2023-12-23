package main

import (
	"fmt"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
	"os"
	"time"
)

var logger *zap.Logger

func init() {
	// 采用默认zap提供的日志打印设置
	encoderConfig := zap.NewProductionEncoderConfig()
	// 设置日志记录中时间的格式
	encoderConfig.EncodeTime = func(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
		enc.AppendString(t.Format("2006-01-02 15:04:05"))
	}
	// 设置
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	// 生成打印到日志文件中的encoder
	fileEncoder := zapcore.NewJSONEncoder(encoderConfig)

	// 将日志等级标识设置为大写并且有颜色
	encoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	// 返回完整调用路径
	encoderConfig.EncodeCaller = zapcore.FullCallerEncoder
	// 生成打印到console的encoder
	consoleEncoder := zapcore.NewConsoleEncoder(encoderConfig)

	lumberJackLogger := &lumberjack.Logger{
		Filename:   "./test.log",
		MaxSize:    10,
		MaxBackups: 5,
		MaxAge:     30,
		Compress:   false,
	}
	core := zapcore.NewTee(
		// 同时向控制台和文件写日志， 生产环境记得把控制台写入去掉，日志记录的基本是Debug 及以上，生产环境记得改成Info
		zapcore.NewCore(consoleEncoder, zapcore.AddSync(os.Stdout), zapcore.DebugLevel),
		zapcore.NewCore(fileEncoder, zapcore.AddSync(lumberJackLogger), zapcore.DebugLevel),
	)
	// 返回调用栈
	logger = zap.New(core, zap.AddCaller(), zap.AddCallerSkip(2))
}
func main() {

	// gin 使用 zap https://github.com/gin-contrib/zap

	fmt.Println("hello")

	logger.Info("ff")
}
