package log

import (
	"fmt"
	"io"
	"runtime"
	"strings"

	"github.com/sirupsen/logrus"
)

var (
	logger        = logrus.New()
	logFileSearch = "/pslytics-api/"
)

// Fields wraps logrus.Fields, which is a map[string]interface{}
type Fields logrus.Fields

func SetLogLevel(level logrus.Level) {
	logger.Level = level
}

func SetLogFormatter(formatter logrus.Formatter) {
	logger.Formatter = formatter
}

func SetOut(writer io.Writer) {
	logger.Out = writer
}

func formatMessageWithFileInfo(msg string) string {
	res := fmt.Sprintf("[%v] %v", fileInfo(3), msg)
	return res
}

// Debugf logs a message at level Debug on the standard logger.
func Debugf(format string, args ...interface{}) {
	entry := logger.WithFields(logrus.Fields{})
	format = formatMessageWithFileInfo(format)
	entry.Debugf(format, args...)
}

// Debugln logs a message with fields at level Debug on the standard logger.
func Debugln(args ...interface{}) {
	entry := logger.WithFields(logrus.Fields{})
	format := formatMessageWithFileInfo(sprintlnn(args...))
	entry.Debugln(format)
}

// Infof logs a message at level Info on the standard logger.
func Infof(format string, args ...interface{}) {
	entry := logger.WithFields(logrus.Fields{})
	format = formatMessageWithFileInfo(format)
	entry.Infof(format, args...)
}

// Infoln logs a message with fields at level Debug on the standard logger.
func Infoln(args ...interface{}) {
	entry := logger.WithFields(logrus.Fields{})
	format := formatMessageWithFileInfo(sprintlnn(args...))
	entry.Infoln(format)
}

// Info logs a message with fields at level Debug on the standard logger.
func Info(args ...interface{}) {
	entry := logger.WithFields(logrus.Fields{})
	format := formatMessageWithFileInfo(sprintlnn(args...))
	entry.Infoln(format)
}

// Warnf logs a message at level Warn on the standard logger.
func Warnf(format string, args ...interface{}) {
	entry := logger.WithFields(logrus.Fields{})
	format = formatMessageWithFileInfo(format)
	entry.Warnf(format, args...)
}

// Warnln logs a message with fields at level Debug on the standard logger.
func Warnln(args ...interface{}) {
	entry := logger.WithFields(logrus.Fields{})
	format := formatMessageWithFileInfo(sprintlnn(args...))
	entry.Warn(format)
}

// Error logs a message with fields at level Debug on the standard logger.
func Error(args ...interface{}) {
	entry := logger.WithFields(logrus.Fields{})
	format := formatMessageWithFileInfo(sprintlnn(args...))
	entry.Error(format)
}

// Errorf logs a message at level Error on the standard logger.
func Errorf(format string, args ...interface{}) {
	entry := logger.WithFields(logrus.Fields{})
	format = formatMessageWithFileInfo(format)
	entry.Errorf(format, args...)
}

// Errorln logs a message with fields at level Debug on the standard logger.
func Errorln(args ...interface{}) {
	entry := logger.WithFields(logrus.Fields{})
	format := formatMessageWithFileInfo(sprintlnn(args...))
	entry.Errorln(format)
}

// Fatalf logs a message at level Fatal on the standard logger.
func Fatalf(format string, args ...interface{}) {
	entry := logger.WithFields(logrus.Fields{})
	format = formatMessageWithFileInfo(format)
	entry.Fatalf(format, args...)
}

// Panic logs a message at level Panic on the standard logger.
func Panic(args ...interface{}) {
	entry := logger.WithFields(logrus.Fields{})
	format := formatMessageWithFileInfo(sprintlnn(args...))
	entry.Panic(format)
}

func sprintlnn(args ...interface{}) string {
	msg := fmt.Sprintln(args...)
	return msg[:len(msg)-1]
}

func fileInfo(skip int) string {
	_, file, line, ok := runtime.Caller(skip)
	if !ok {
		file = "<???>"
		line = 1
	} else {
		slash := strings.LastIndex(file, logFileSearch)
		if slash >= 0 {
			file = file[slash+1:]
		}
	}
	return fmt.Sprintf("%s:%d", file, line)
}
