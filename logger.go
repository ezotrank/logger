package logger

import (
	"flag"
	"log"
	"os"
)

const (
	FATALLV = iota
	ERRORLV
	WARNLV
	INFOLV
	DEBUGLV
	VERBOSELV
)

var logger Logger

func init() {
	flag.IntVar(&logger.level, "v", 3, "log levels: Fatal,Error,Warn,Info,Debug,Verbose - from 0 to 5")
	logger.init()
}

type Logger struct {
	loggers map[uint]*log.Logger
	level   int
}

func (self *Logger) init() {
	self.loggers = map[uint]*log.Logger{
		FATALLV:   log.New(os.Stderr, "fatal: ", log.Ldate|log.Ltime),
		ERRORLV:   log.New(os.Stderr, "error: ", log.Ldate|log.Ltime),
		WARNLV:    log.New(os.Stderr, "warning: ", log.Ldate|log.Ltime),
		INFOLV:    log.New(os.Stderr, "info: ", log.Ldate|log.Ltime),
		DEBUGLV:   log.New(os.Stderr, "debug: ", log.Ldate|log.Ltime),
		VERBOSELV: log.New(os.Stderr, "verbose: ", log.Ldate|log.Ltime),
	}
}

func AvailableForLevel(lvl int) bool {
    return lvl <= logger.level
}

func (self *Logger) Println(level uint, v ...interface{}) {
	if level <= uint(self.level) {
		self.loggers[level].Println(v...)
	}
}

func (self *Logger) Printf(level uint, t string, v ...interface{}) {
	if level <= uint(self.level) {
		self.loggers[level].Printf(t, v...)
	}
}

func Fatalln(v ...interface{}) {
	logger.loggers[FATALLV].Fatalln(v...)
}

func Fatalf(tmpl string, v ...interface{}) {
	logger.loggers[FATALLV].Fatalf(tmpl, v...)
}

func Infoln(v ...interface{}) {
	logger.Println(INFOLV, v...)
}

func Infof(t string, v ...interface{}) {
	logger.Printf(INFOLV, t, v...)
}

func Debugln(v ...interface{}) {
	logger.Println(DEBUGLV, v...)
}

func Debugf(t string, v ...interface{}) {
	logger.Printf(DEBUGLV, t, v...)
}

func Errorln(v ...interface{}) {
	logger.Println(ERRORLV, v...)
}

func Errorf(t string, v ...interface{}) {
	logger.Printf(ERRORLV, t, v...)
}

func Warnln(v ...interface{}) {
	logger.Println(WARNLV, v...)
}

func Warnf(t string, v ...interface{}) {
	logger.Printf(WARNLV, t, v...)
}

func Verboseln(v ...interface{}) {
	logger.Println(VERBOSELV, v...)
}

func Verbosef(t string, v ...interface{}) {
	logger.Printf(VERBOSELV, t, v...)
}
