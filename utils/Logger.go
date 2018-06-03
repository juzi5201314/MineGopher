package utils

import (
	"fmt"
	"log"
	"os"
	"strings"
)

const (
	Debug      = "[Debug]"
	Info       = "[Info]"
	Notice     = "[Notice]"
	Alert      = "[Alert]"
	Error      = "[Error]"
	Warning    = "[Warning]"
	Critical   = "[Critical]"
	Chat       = "[Chat]"
	StackTrace = "[Stack Trace]"
)

type Logger struct {
	messageQueue chan string
	file         *os.File
	closed       bool
}

var static_logger *Logger = nil

func NewLogger(file *os.File) *Logger {
	logger := new(Logger)
	logger.messageQueue = make(chan string, 10)
	logger.file = file
	logger.closed = false
	go logger.process()
	static_logger = logger
	return logger
}

func (logger *Logger) process() {
	for !logger.closed {
		logger.write()
	}
}

func (logger *Logger) write() {
	d := ColorString(<-logger.messageQueue)
	logger.file.WriteString(d.StripAll() + "\n")
	log.Println(d.ToANSI() + AnsiReset)
}

func (logger *Logger) Close() {
	logger.closed = true
	defer logger.file.Close()
	logger.Info("Logger closed!")
	for len(logger.messageQueue) > 0 {
		logger.write()
	}
}

func (logger *Logger) Notice(messages ...interface{}) {
	logger.messageQueue <- Yellow + Notice + " " + strings.Trim(fmt.Sprint(messages), "[]")
}

func (logger *Logger) Debug(messages ...interface{}) {
	logger.messageQueue <- (Orange + Debug + " " + strings.Trim(fmt.Sprint(messages), "[]"))
}

func (logger *Logger) Info(messages ...interface{}) {
	logger.messageQueue <- (BrightCyan + Info + " " + strings.Trim(fmt.Sprint(messages), "[]"))
}

func (logger *Logger) Alert(messages ...interface{}) {
	logger.messageQueue <- (BrightRed + Alert + " " + strings.Trim(fmt.Sprint(messages), "[]"))
}

func (logger *Logger) Warning(messages ...interface{}) {
	logger.messageQueue <- (BrightRed + Bold + Warning + " " + strings.Trim(fmt.Sprint(messages), "[]"))
}

func (logger *Logger) Critical(messages ...interface{}) {
	logger.messageQueue <- (BrightRed + Underlined + Bold + Critical + " " + strings.Trim(fmt.Sprint(messages), "[]"))
}

func (logger *Logger) Error(messages ...interface{}) {
	logger.messageQueue <- (Red + Error + " " + strings.Trim(fmt.Sprint(messages), "[]"))
}

func (logger *Logger) Println(messages ...interface{}) {
	logger.messageQueue <- (BrightCyan + Chat + " " + strings.Trim(fmt.Sprint(messages), "[]"))
}

func (logger *Logger) PacicError(err error) {
	if err == nil {
		return
	}
	logger.Error(err.Error())
}

func GetLogger() *Logger {
	return static_logger
}
