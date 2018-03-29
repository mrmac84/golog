package golog

import (
	"io"
	"log"
	"os"
)

//ErrorLog is a logger instance used for error logging
var errorLog *log.Logger

//InfoLog is a logger instance used for generic info logging
var infoLog *log.Logger

//DebugLog is a logger instance used for debugg logging
var debugLog *log.Logger

func init() {

	//init the log instances and set the out
	infoLog = log.New(os.Stdout, "INFO ", log.Ldate|log.Ltime|log.Lshortfile)
	errorLog = log.New(os.Stderr, "ERROR ", log.Ldate|log.Ltime|log.Lshortfile)
	debugLog = log.New(os.Stdout, "DEBUG ", log.Ldate|log.Ltime|log.Lshortfile)
}

//Error writes error messages to the established output
func Error(format string, v ...interface{}) {

	errorLog.Printf(format, v)
}

//Debug writes debug messages to the established output
func Debug(format string, v ...interface{}) {

	debugLog.Printf(format, v)
}

//Info writes info messages to the established output
func Info(format string, v ...interface{}) {

	infoLog.Printf(format, v)
}

//SetErrorPrefix updates the prefix used for error logs
func SetErrorPrefix(prefix string) {
	errorLog.SetPrefix(prefix)
}

//SetErrorOutput updates the  destination output for error logs
func SetErrorOutput(out io.Writer) {
	errorLog.SetOutput(out)
}

//SetInfoPrefix updates the prefix used for info logs
func SetInfoPrefix(prefix string) {
	infoLog.SetPrefix(prefix)
}

//SetInfoOutput updates the  destination output for info logs
func SetInfoOutput(out io.Writer) {
	infoLog.SetOutput(out)
}

//SetDebugPrefix updates the prefix used for debug logs
func SetDebugPrefix(prefix string) {
	debugLog.SetPrefix(prefix)
}

//SetDebugOutput updates the  destination output for debug logs
func SetDebugOutput(out io.Writer) {
	debugLog.SetOutput(out)
}
