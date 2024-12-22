package helper

import (
	"context"
	"log"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm/logger"
)

type FileLogger struct {
	logLevel logger.LogLevel
	logger   *log.Logger
}

func NewFileGormLogger(logerLevel logger.LogLevel, filePath string) (*FileLogger, error) {
	file, err := os.OpenFile(filePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		return nil, err
	}

	return &FileLogger{
		logLevel: logerLevel,
		logger:   log.New(file, "", log.LstdFlags),
	}, nil
}

func (f *FileLogger) LogMode(level logger.LogLevel) logger.Interface {
	newLogger := *f
	newLogger.logLevel = level
	return &newLogger
}

func (f *FileLogger) Info(ctx context.Context, msg string, data ...interface{}) {
	if f.logLevel >= logger.Info {
		f.logger.Printf("[INFO] "+msg, data...)
	}
}

func (f *FileLogger) Warn(ctx context.Context, msg string, data ...interface{}) {
	if f.logLevel >= logger.Warn {
		f.logger.Printf("[WARN] "+msg, data...)
	}
}

func (f *FileLogger) Error(ctx context.Context, msg string, data ...interface{}) {
	if f.logLevel >= logger.Error {
		f.logger.Printf("[ERROR] "+msg, data...)
	}
}

func (f *FileLogger) Trace(ctx context.Context, begin time.Time, fc func() (string, int64), err error) {
	elapsed := time.Since(begin)
	sql, rows := fc()

	if err != nil {
		f.logger.Printf("[TRACE][ERROR] %s | %s | %d rows | %v", elapsed, sql, rows, err)
	} else {
		f.logger.Printf("[TRACE] %s | %s | %d rows", elapsed, sql, rows)
	}
}

func NewFileGinLog() gin.HandlerFunc {
	file, err := os.OpenFile("pkg/log/gin.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		panic(err)
	}

	logger := log.New(file, "", log.LstdFlags)
	return func(ctx *gin.Context) {
		startTime := time.Now()
		ctx.Next()

		latency := time.Since(startTime)
		status := ctx.Writer.Status()
		clientIP := ctx.ClientIP()
		method := ctx.Request.Method
		path := ctx.Request.URL.Path
		userAgent := ctx.Request.UserAgent()

		logger.Printf("| %d | %v | %s | %s | %s | %s |\n",
			status, latency, clientIP, method, path, userAgent)
	}

}

func RefreshLog(filePaths []string) {
	for _, filePath := range filePaths {
		err := os.Remove(filePath)
		if err != nil {
			panic(err)
		}
	}
}
