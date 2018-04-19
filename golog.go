package golog

import (
	"fmt"
	"io"
	"log"
	"path"
	"runtime"
	"strconv"
	"time"
)

func init() {

}

//Golog represents a golog instance with all the necessary options.
// infoPrefix: the prefix used for info logs, defaults to INFO
// debugPrefix: the prefix used for debug logs, defaults to DEBUG
//errorPrefix: the prefix used for error logs, defaults to ERROR
//showTimestmp: show timestamp in the logs; defults to true
//showPrefix: show prefix in the logs; defaults to true
//showCallerInfo: show info about the caller in the logs; defaults to true
//out: output destination for the logs; defaults to stdout
//gologer: the logger instance
type Golog struct {
	InfoPrefix     string
	DebugPrefix    string
	ErrorPrefix    string
	ShowTimestamp  bool
	ShowPrefix     bool
	ShowCallerInfo bool
	Out            io.Writer
	Gologger       *log.Logger
}

//New initializes a new Golog instance
func New(output io.Writer) *Golog {
	return &Golog{
		InfoPrefix:     "INFO",
		DebugPrefix:    "DEBUG",
		ErrorPrefix:    "ERROR",
		ShowTimestamp:  true,
		ShowPrefix:     true,
		ShowCallerInfo: true,
		Out:            output,
		Gologger:       log.New(output, "", 0),
	}
}

//getCallerInfo returns the info about the function calling golog
func getCallerInfo(skip int) string {
	var callerInfo string
	//var callingFuncName string

	_, fullFilePath, lineNumber, ok := runtime.Caller(skip)

	if ok {
		//callingFuncName = runtime.FuncForPC(pc).Name()

		// Split the path and use only the last 2 elements (package and file name)
		dirPath, fileName := path.Split(fullFilePath)
		var moduleName string
		if dirPath != "" {
			dirPath = dirPath[:len(dirPath)-1]
			_, moduleName = path.Split(dirPath)
		}
		callerInfo = moduleName + "/" + fileName + " " + strconv.Itoa(lineNumber)
	}

	return callerInfo
}

//Error writes error messages to the established output
func (g *Golog) Error(format string, v ...interface{}) {
	//add caller info to prefix
	prefix := fmt.Sprintf("ERROR %s %s ", time.Now().Format("02-01-2006 15:04:05"), getCallerInfo(2))

	g.Gologger.SetPrefix(prefix)

	g.Gologger.Printf(format, v...)
}

//Debug writes debug messages to the established output
func (g *Golog) Debug(format string, v ...interface{}) {
	//add caller info to prefix
	prefix := fmt.Sprintf("DEBUG %s %s ", time.Now().Format("02-01-2006 15:04:05"), getCallerInfo(2))

	g.Gologger.SetPrefix(prefix)

	g.Gologger.Printf(format, v...)
}

//Info writes info messages to the established output
func (g *Golog) Info(format string, v ...interface{}) {

	//add caller info to prefix
	prefix := fmt.Sprintf("INFO %s %s ", time.Now().Format("02-01-2006 15:04:05"), getCallerInfo(2))

	g.Gologger.SetPrefix(prefix)

	g.Gologger.Printf(format, v...)
}

//SetErrorPrefix updates the prefix used for error logs
func (g *Golog) SetErrorPrefix(prefix string) {
	g.Gologger.SetPrefix(prefix)
}

//SetErrorOutput updates the  destination output for error logs
func (g *Golog) SetErrorOutput(out io.Writer) {
	g.Gologger.SetOutput(out)
}

//SetInfoPrefix updates the prefix used for info logs
func (g *Golog) SetInfoPrefix(prefix string) {
	g.Gologger.SetPrefix(prefix)
}

//SetInfoOutput updates the  destination output for info logs
func (g *Golog) SetInfoOutput(out io.Writer) {
	g.Gologger.SetOutput(out)
}

//SetDebugPrefix updates the prefix used for debug logs
func (g *Golog) SetDebugPrefix(prefix string) {
	g.Gologger.SetPrefix(prefix)
}

//SetDebugOutput updates the  destination output for debug logs
func (g *Golog) SetDebugOutput(out io.Writer) {
	g.Gologger.SetOutput(out)
}
