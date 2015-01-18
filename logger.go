package logger

import (
	"flag"
	"log"
	"os"
)

const (
	FatalLevel = iota
	ErrorLevel
	WarnLevel
	InfoLevel
	DebugLevel
	VerboseLevel
)

var logger Logger

func init() {
	flag.IntVar(&logger.level, "v", 3, "log levels Fatal,Error,Warn,Info,Debug,Verbose - 0..5")
	logger.init()
}

type Logger struct {
	loggers map[int]*log.Logger
	level   int
}

func (self *Logger) init() {
	self.loggers = map[int]*log.Logger{
		FatalLevel:   log.New(os.Stderr, "fatal: ", log.Ldate|log.Ltime|log.Lshortfile),
		ErrorLevel:   log.New(os.Stderr, "error: ", log.Ldate|log.Ltime|log.Lshortfile),
		WarnLevel:    log.New(os.Stderr, "warning: ", log.Ldate|log.Ltime|log.Lshortfile),
		InfoLevel:    log.New(os.Stderr, "info: ", log.Ldate|log.Ltime|log.Lshortfile),
		DebugLevel:   log.New(os.Stderr, "debug: ", log.Ldate|log.Ltime|log.Lshortfile),
		VerboseLevel: log.New(os.Stderr, "verbose: ", log.Ldate|log.Ltime|log.Lshortfile),
	}
}

func (self *Logger) Println(level int, v ...interface{}) {
	if self.level < level {
		return
	}
	self.loggers[level].Println(v)
}

func (self *Logger) Printf(level int, t string, v ...interface{}) {
	if self.level < level {
		return
	}
	self.loggers[level].Printf(t, v)
}

func Fatalln(v ...interface{}) {
	logger.loggers[FatalLevel].Fatalln(v)
}

func Fatalf(tmpl string, v ...interface{}) {
	logger.loggers[FatalLevel].Fatalf(tmpl, v)
}

func Infoln(v ...interface{}) {
	logger.Println(InfoLevel, v)
}

func Infof(t string, v ...interface{}) {
	logger.Printf(InfoLevel, t, v)
}

func Debugln(v ...interface{}) {
	logger.Println(DebugLevel, v)
}

func Debugf(t string, v ...interface{}) {
	logger.Printf(DebugLevel, t, v)
}

func Errorln(v ...interface{}) {
	logger.Println(ErrorLevel, v)
}

func Errorf(t string, v ...interface{}) {
	logger.Printf(ErrorLevel, t, v)
}

func Warnln(v ...interface{}) {
	logger.Println(WarnLevel, v)
}

func Warnf(t string, v ...interface{}) {
	logger.Printf(WarnLevel, t, v)
}

func Verboseln(v ...interface{}) {
	logger.Println(VerboseLevel, v)
}

func Verbosef(t string, v ...interface{}) {
	logger.Printf(VerboseLevel, t, v)
}
