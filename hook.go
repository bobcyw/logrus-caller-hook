package logrus_caller_hook

import (
	"path/filepath"
	"runtime"
	"strconv"
	"strings"

	"github.com/bobcyw/logrus"
)

var (
	//DefaultFieldName is used by New() to create default filed name
	DefaultFieldName = "caller"
)

type normalCallerHook struct{}

func (hook *normalCallerHook) Fire(entry *logrus.Entry) error {
	entry.Data[DefaultFieldName] = hook.caller(entry)
	return nil
}

func (hook *normalCallerHook) Levels() []logrus.Level {
	return []logrus.Level{
		logrus.PanicLevel,
		logrus.FatalLevel,
		logrus.ErrorLevel,
		logrus.WarnLevel,
		logrus.InfoLevel,
		logrus.DebugLevel,
	}
}

func (hook *normalCallerHook) caller(entry *logrus.Entry) string {
	if _, file, line, ok := runtime.Caller(5); ok {
		return strings.Join([]string{filepath.Base(file), strconv.Itoa(line)}, ":")
	} else {
		return ""
	}
}

/*
New create a default hook. Default field name is caller.
*/
func New() logrus.Hook {
	return &normalCallerHook{}
}

type customCallerHook struct {
	normalCallerHook
	fieldName string
	fullPath  bool
}

func (hook *customCallerHook) Fire(entry *logrus.Entry) error {
	entry.Data[hook.fieldName] = hook.caller(entry)
	return nil
}

func (hook *customCallerHook) caller(entry *logrus.Entry) string {
	if _, file, line, ok := runtime.Caller(5); ok {
		if hook.fullPath == false {
			return strings.Join([]string{filepath.Base(file), strconv.Itoa(line)}, ":")
		} else {
			return strings.Join([]string{file, strconv.Itoa(line)}, ":")
		}
	} else {
		return ""
	}
}

/*
NewCustom create a custom hook. User can define the field name and define whether use full file path
*/
func NewCustom(fieldName string, fullPath bool) logrus.Hook {
	return &customCallerHook{fieldName: fieldName, fullPath: fullPath}
}
