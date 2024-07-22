package logx

import (
	"davinci-chat/config"
	"davinci-chat/consts"
	"fmt"
	"log"
	"os"
	"sync"
)

type Logger struct {
	*log.Logger
	Level consts.LogLevel
	File  *os.File
}

func (l *Logger) Close() {
	err := l.File.Close()
	if err != nil {
		log.Fatal(err)
	}
}

func (l *Logger) Info(format string, v ...interface{}) {
	if l.Level >= consts.LevelInfo {
		l.Printf("[INFO] "+format, v...)
	}
}

func (l *Logger) Debug(format string, v ...interface{}) {
	if l.Level >= consts.LevelDebug {
		l.Printf("[DEBUG] "+format, v...)
	}
}

func (l *Logger) Fatal(v ...interface{}) {
	l.Printf("[FATAL] " + fmt.Sprint(v...))
	os.Exit(1)
}

var (
	once   sync.Once
	logger *Logger
)

func OpenLogFile(logPath string) *os.File {
	file, err := os.OpenFile(logPath, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("error opening file: %v", err)
	}
	return file
}

func GetLogPath() string {
	env := config.GetRunEnv()
	switch env {
	case consts.Production:
		return "/var/log/chat/app.log"
	case consts.Development:
		return "./app.log"
	default:
		return ""
	}
}

func GetLogger() *Logger {
	once.Do(func() {
		logPath := GetLogPath()
		file := OpenLogFile(logPath)
		log.SetOutput(file)
		logger = &Logger{
			Logger: log.New(file, "", log.Ldate|log.Ltime),
			File:   file,
			Level:  consts.LevelDebug,
		}
	})
	return logger
}
