package xlog

import (
	"bufio"
	"bytes"
	"os"
	"time"

	"github.com/lestrrat/go-file-rotatelogs"
	"github.com/pkg/errors"
	"github.com/rifflock/lfshook"
	log "github.com/sirupsen/logrus"
)

type Fields log.Fields

var lPath = ""
var lLevel = ""
var lRotateInteval = uint(24)

// Init: Init log save path and log level
func Init(logPath string, logLevel string, logLifeCycle, logRotateInteval uint) {

	if "" != logPath {
		lPath = logPath
	}
	if "" != logLevel {
		lLevel = logLevel
	}
	if 0 != logRotateInteval {
		lRotateInteval = logRotateInteval
	}

	// lowwer only
	lv := string(bytes.ToLower([]byte(lLevel)))

	// log level
	switch lv {
	case "debug":
		log.SetLevel(log.DebugLevel)
	case "info":
		log.SetLevel(log.InfoLevel)
	case "warn":
		log.SetLevel(log.WarnLevel)
	case "error":
		log.SetLevel(log.ErrorLevel)
	default:
		log.SetLevel(log.DebugLevel)
	}

	//// log cut
	//if logPath == "" {
	//	// print to console
	//	log.SetOutput(os.Stdout)
	//	return
	//}
	writer, err := rotatelogs.New(
		lPath+".%Y%m%d%H%M",
		rotatelogs.WithLinkName(lPath), // 生成软链，指向最新日志文件
				rotatelogs.WithMaxAge(time.Duration(logLifeCycle)*time.Hour), // 文件最大保存时间
		rotatelogs.WithRotationTime(time.Duration(lRotateInteval)*time.Hour), // 日志切割时间间隔
	)
	if err != nil {
		log.Errorf("config local file system logger error. %v", errors.WithStack(err))
	}

	// shield console output
	setNull()

	// set different dir for different level
	lfHook := lfshook.NewHook(lfshook.WriterMap{
		log.DebugLevel: writer, // 为不同级别设置不同的输出目的
		log.InfoLevel:  writer,
		log.WarnLevel:  writer,
		log.ErrorLevel: writer,
	}, log.Formatter(&log.JSONFormatter{TimestampFormat: time.RFC3339Nano}))
	log.AddHook(lfHook)
}

// setNull: don't output to console
// In:
// Out:
func setNull() {
	src, err := os.OpenFile(os.DevNull, os.O_APPEND|os.O_WRONLY, os.ModeAppend)
	if err != nil {
		log.Errorf("open dev null error. %v", errors.WithStack(err))
	}
	writer := bufio.NewWriter(src)
	log.SetOutput(writer)
}

// buildEntry：add seesion to fields and return a new entry
// In: 	-string, session
//		-*log.Fields
// Out:	-*log.Entry
func buildEntry(session string, fields *Fields) (entry *log.Entry) {
	(*fields)["session"] = session
	logFields := log.Fields{}
	for k, v := range *fields {
		logFields[k] = v
	}
	entry = log.WithFields(logFields)
	return
}

// Debug: log of debug level
func Debug(session string, message string, fields Fields) {
	buildEntry(session, &fields).Debug(message)
}

// DebugSimple: simple log of debug level
func DebugSimple(message string, fields Fields) {
	buildEntry("", &fields).Debug(message)
}

// Info: log of info level
func Info(session string, message string, fields Fields) {
	buildEntry(session, &fields).Info(message)
}

// InfoSimple: simple log of info level
func InfoSimple(message string, fields Fields) {
	buildEntry("", &fields).Info(message)
}

// Warn: log of warn level
func Warn(session string, message string, fields Fields) {
	buildEntry(session, &fields).Warn(message)
}

// WarnSimple: simple log of warn level
func WarnSimple(message string, fields Fields) {
	buildEntry("", &fields).Warn(message)
}

// Error: log of error level
func Error(session string, message string, fields Fields) {
	buildEntry(session, &fields).Error(message)
}

// ErrorSimple: simple log of error level
func ErrorSimple(message string, fields Fields) {
	buildEntry("", &fields).Error(message)
}
