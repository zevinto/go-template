package logger

import (
	"go-template/internal/config"
	"go-template/internal/global"
	"go-template/internal/tool"
	"os"
	"strings"
	"time"

	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var (
	options []zap.Option
	conf    config.Logger
)

func New() (logger *zap.Logger, err error) {
	// 获取配置文件
	conf = global.Config.Logger

	// 创建日志存放目录
	rootDir, _ := tool.GetRootDir()
	logDir := rootDir + conf.FilePath
	err = os.MkdirAll(logDir, os.ModePerm)
	if err != nil {
		return
	}

	loggerConf := genConfig()
	loggerConf.EncoderConfig = genEncodeConfig()

	writer, err := genWriteSyncer()
	if err != nil {
		return
	}

	core := zapcore.NewCore(
		zapcore.NewJSONEncoder(loggerConf.EncoderConfig),
		writer,
		loggerConf.Level,
	)

	// 添加调用位置
	options = append(options, zap.AddCaller())

	logger = zap.New(core, options...)

	return
}

// 生成WriteSyncer
func genWriteSyncer() (writeSyncer zapcore.WriteSyncer, err error) {
	// 创建日志存放目录
	rootDir, _ := tool.GetRootDir()
	logDir := rootDir + conf.FilePath
	err = os.MkdirAll(logDir, os.ModePerm)
	if err != nil {
		return
	}

	lumberJack := &lumberjack.Logger{
		Filename:   logDir + conf.FileName, // 文件名称
		MaxSize:    conf.MaxSize,           // 切割的文件大小 单位MB
		MaxBackups: conf.MaxBackups,        // 保留的最大文件数量
		MaxAge:     conf.MaxAge,            // 保留的文件最大天数
	}

	writeSyncer = zapcore.AddSync(lumberJack)
	return
}

// 生成配置
func genConfig() (config zap.Config) {
	config = zap.NewProductionConfig()

	config.EncoderConfig = genEncodeConfig()
	config.Level = zap.NewAtomicLevelAt(getLevel())
	return
}

// 生成编码配置
func genEncodeConfig() (c zapcore.EncoderConfig) {
	c = zap.NewProductionEncoderConfig()

	c.EncodeTime = func(time time.Time, encoder zapcore.PrimitiveArrayEncoder) {
		encoder.AppendString(time.Format("2006-01-02 15:04:05"))
	}

	c.EncodeLevel = func(l zapcore.Level, encoder zapcore.PrimitiveArrayEncoder) {
		encoder.AppendString(strings.ToUpper(l.String()))
	}

	c.TimeKey = "time"

	return
}

// 配置文件的 level 转换为 zapcore 的 level
func getLevel() (level zapcore.Level) {
	switch conf.Level {
	case "debug":
		level = zap.DebugLevel
	case "info":
		level = zap.InfoLevel
	case "warn":
		level = zap.WarnLevel
	case "error":
		level = zap.ErrorLevel
	case "dpanic":
		level = zap.DPanicLevel
	case "panic":
		level = zap.PanicLevel
	case "fatal":
		level = zap.FatalLevel
	default:
		level = zap.InfoLevel
	}

	return
}
