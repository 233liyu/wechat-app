package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var (
	globalLogger *zap.Logger
)

// InitLogger 初始化全局 logger
func InitLogger() {
	config := zap.NewDevelopmentConfig()
	config.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	logger, err := config.Build()
	if err != nil {
		panic(err)
	}
	globalLogger = logger
	// 也可以配置输出到文件，这里暂时只输出到终端
}

// GetLogger 获取全局 logger 实例
func GetLogger() *zap.Logger {
	return globalLogger
}
