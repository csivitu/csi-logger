package initializers

import (
	"log"
	"os"
	"path/filepath"
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var Logger *zap.SugaredLogger

const (
	InfoLevel  = zapcore.InfoLevel  // InfoLevel logs are Debug priority logs.
	WarnLevel  = zapcore.WarnLevel  // WarnLevel logs are more important than Info, but don't need individual human review.
	ErrorLevel = zapcore.ErrorLevel // ErrorLevel logs are high-priority.
)

const (
	InfoLogFilePath  = "logs/info.log"
	WarnLogFilePath  = "logs/warn.log"
	ErrorLogFilePath = "logs/error.log"
)

var logFiles struct {
	InfoLogFile  *os.File
	WarnLogFile  *os.File
	ErrorLogFile *os.File
}

func AddLogger() {
	openLogFiles()

	infoCore := newCore(logFiles.InfoLogFile, InfoLevel)
	warnCore := newCore(logFiles.WarnLogFile, WarnLevel)
	errorCore := newCore(logFiles.ErrorLogFile, ErrorLevel)

	Logger = zap.New(zapcore.NewTee(infoCore, warnCore, errorCore)).Sugar()
}

func newCore(LogFile *os.File, LoggerLevel zapcore.Level) zapcore.Core {
	encoderCfg := zap.NewProductionEncoderConfig()
	encoderCfg.EncodeTime = zapcore.TimeEncoderOfLayout(time.RFC1123)
	fileCore := zapcore.NewCore(
		zapcore.NewConsoleEncoder(encoderCfg),
		zapcore.AddSync(LogFile),
		LoggerLevel,
	)
	return fileCore
}

func openLogFiles() {
	logFiles.InfoLogFile = openFile(InfoLogFilePath)
	logFiles.WarnLogFile = openFile(WarnLogFilePath)
	logFiles.ErrorLogFile = openFile(ErrorLogFilePath)
}

func openFile(LogFilePath string) *os.File {
	dir := filepath.Dir(LogFilePath)

	if _, err := os.Stat(dir); os.IsNotExist(err) {
		err := os.MkdirAll(dir, 0755)
		if err != nil {
			log.Printf("Failed to create directory: " + err.Error())
		}
	}

	logFile, err := os.OpenFile(LogFilePath, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		log.Printf("Failed to open log file: " + err.Error())
	}
	return logFile
}

func LoggerCleanUp() {
	logFiles.InfoLogFile.Close()
	logFiles.WarnLogFile.Close()
	logFiles.ErrorLogFile.Close()
	Logger.Sync()
}
