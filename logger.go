package logger

import (
	"io"
	logger "log"
	"os"
	"os/signal"
	"strings"
	"sync"
	"syscall"

	"github.com/hashicorp/logutils"
)

const (
	DEFAULT_LOG_LEVEL = "INFO"
)

var (
	log *logger.Logger
)

func init() {
	SetLogger(DEFAULT_LOG_LEVEL, "")
}

func initLogFilter(logLevel string, logFile string) io.Writer {
	writer := os.Stderr
	if len(logFile) > 0 {
		var err error
		writer, err = os.OpenFile(logFile, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0644)
		if err != nil {
			panic(err)
		}
	}
	filter := &logutils.LevelFilter{
		Levels:   []logutils.LogLevel{"DEBUG", "INFO", "WARN", "ERROR", "NONE"},
		MinLevel: logutils.LogLevel(logLevel),
		Writer:   writer,
	}
	return filter
}

func GetLogger() *logger.Logger {
	return log
}

func SetLogger(logLevel string, logFile string) {
	logLevel = strings.ToUpper(logLevel)
	mutex := &sync.Mutex{}
	log = &logger.Logger{}
	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGUSR1)
	go func() {
		for sig := range c {
			mutex.Lock()
			log.Printf("[WARN] Go a %v Signal! Reopen logs", sig)
			log.SetOutput(initLogFilter(logLevel, logFile))
			mutex.Unlock()
		}
	}()
	log.SetFlags(logger.LstdFlags | logger.Lshortfile)
	log.SetOutput(initLogFilter(logLevel, logFile))
}