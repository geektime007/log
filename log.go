// Package log implements the logging supports.
//
//	log.Info("Everything is working great.")
//	log.Infof("Task %v is created successfully", id)
//	log.Warnf("Something is working not properly")
//	log.Warnf("Something is working not properly")
//	log.Smart(err, "Smartly choosing the logging level by error level")
package log

import (
	"fmt"

	"github.com/geektime007/util/log"
	"github.com/geektime007/util/orderedmap"
)

// Level
const (
	DEBUG int = iota
	INFO
	WARN
	ERROR
)

// Default logger prefix length.
const DefaultLoggerName = "Geek"

type loggerWrapper log.Logger

// Global logger.
var globalLogger = (*loggerWrapper)(log.Get(DefaultLoggerName))

// Global logging metas.
var globalLoggingMeta = orderedmap.NewOrderedMap()

// Inits global logger on this package's first import.
func init() { globalLogger.Init() }

// Unwraps to *log.Logger.
func (logger *loggerWrapper) Unwrap() *log.Logger { return (*log.Logger)(logger) }

// Inits this logger.
func (logger *loggerWrapper) Init() {
	logger.Unwrap().SetCallerDepth(3)
	logger.Unwrap().DisableCallerSourceLogging()
	logger.Unwrap().SetNameLength(len(DefaultLoggerName))
}

// Set log level
func (logger *loggerWrapper) SetLogLevel(level int) { logger.Unwrap().SetLevel(level) }

// SetMeta sets logging meta for all logger.
func SetMeta(name string, value interface{}) {
	globalLoggingMeta.Set(name, value)
}

// Prefix returns a new logger with name `{NAME} :: {s}`
func Prefix(s string) *loggerWrapper {
	name := fmt.Sprintf("%v %v", DefaultLoggerName, s)
	logger := (*loggerWrapper)(log.Get(name))
	logger.Init()
	return logger
}

func P(s string) *loggerWrapper { return Prefix(s) }

// Debug logs message with level DEBUG.
func (logger *loggerWrapper) Debug(a ...interface{}) error { return logger.Unwrap().Debug(a...) }

// Info logs message with level INFO.
func (logger *loggerWrapper) Info(a ...interface{}) error { return logger.Unwrap().Info(a...) }

// Warn logs message with level WARN.
func (logger *loggerWrapper) Warn(a ...interface{}) error { return logger.Unwrap().Warn(a...) }

// Error logs message with level ERROR.
func (logger *loggerWrapper) Error(a ...interface{}) error { return logger.Unwrap().Error(a...) }

// Fatal and logs message with level FATAL.
func (logger *loggerWrapper) Fatal(a ...interface{}) {
	logger.Unwrap().Error(a...)
	logger.Unwrap().Fatal("[ fatal ] program exiting.")
}

// Debugf formats and logs message with level DEBUG.
func (logger *loggerWrapper) Debugf(format string, a ...interface{}) error {
	return logger.Unwrap().Debugf(format, a...)
}

// Infof formats and logs message with level INFO.
func (logger *loggerWrapper) Infof(format string, a ...interface{}) error {
	return logger.Unwrap().Infof(format, a...)
}

// Warnf formats and logs message with level WARN.
func (logger *loggerWrapper) Warnf(format string, a ...interface{}) error {
	return logger.Unwrap().Warnf(format, a...)
}

// Errorf formats and logs message with level ERROR.
func (logger *loggerWrapper) Errorf(format string, a ...interface{}) error {
	return logger.Unwrap().Errorf(format, a...)
}

// Fatalf formats and logs message with level FATAL.
func (logger *loggerWrapper) Fatalf(format string, a ...interface{}) {
	logger.Unwrap().Errorf(format, a...)
	logger.Unwrap().Fatalf("[ fatal ] program exiting.")
}

//// Smart formats and logs message with level FATAL.
//// The logging level is automatically choosed by error.
//// If the given err is not nil or Nil, logs an error message.
//// Otherwise, logs an info message.
//func (logger *loggerWrapper) Smart(err error, format string, a ...interface{}) {
//	// Gets a *Error for `err`.
//	e := errors.Unknown
//
//	if err == nil {
//		// The original err is nil, then e is Nil.
//		e = errors.Nil
//	} else {
//		// ok is true if err is an *Error.
//		var ok bool
//		e, ok = err.(*errors.Error)
//		if !ok {
//			// The original err is not an *Error, then e is an unknown error.
//			e = errors.Unknown.CloneWithOriginError(err)
//		}
//	}
//
//	// Logging
//	if e.IsNil() {
//		logger.Infof(format, a...)
//	} else {
//		logger.Errorf(format, a...)
//	}
//}
//
//// Smart formats and logs message with level FATAL.
//// The logging level is automatically choosed by error.
//// If the given err is not nil or Nil, logs an warn message.
//// Otherwise, logs an debug message.
//func (logger *loggerWrapper) SmartDebug(err error, format string, a ...interface{}) {
//	// Gets a *Error for `err`.
//	e := errors.Unknown
//
//	if err == nil {
//		// The original err is nil, then e is Nil.
//		e = errors.Nil
//	} else {
//		// ok is true if err is an *Error.
//		var ok bool
//		e, ok = err.(*errors.Error)
//		if !ok {
//			// The original err is not an *Error, then e is an unknown error.
//			e = errors.Unknown.CloneWithOriginError(err)
//		}
//	}
//
//	// Logging
//	if e.IsNil() {
//		logger.Debugf(format, a...)
//	} else {
//		logger.Warnf(format, a...)
//	}
//}

// Set all logger level
func SetAllLoggerLevel(level int) {
	loggers := log.GetRegistry()
	for k, _ := range loggers {
		loggers[k].SetLevel(level)
	}
}

// Set Log Level
func SetLogLevel(level int) { globalLogger.SetLogLevel(level) }

// Debug logs message with level DEBUG.
func Debug(a ...interface{}) error { return globalLogger.Debug(a...) }

// Info logs message with level INFO.
func Info(a ...interface{}) error { return globalLogger.Info(a...) }

// Warn logs message with level WARN.
func Warn(a ...interface{}) error { return globalLogger.Warn(a...) }

// Error logs message with level ERROR.
func Error(a ...interface{}) error { return globalLogger.Error(a...) }

// Fatal and logs message with level FATAL.
func Fatal(a ...interface{}) { globalLogger.Fatal(a...) }

// Debugf formats and logs message with level DEBUG.
func Debugf(format string, a ...interface{}) error { return globalLogger.Debugf(format, a...) }

// Infof formats and logs message with level INFO.
func Infof(format string, a ...interface{}) error { return globalLogger.Infof(format, a...) }

// Warnf formats and logs message with level WARN.
func Warnf(format string, a ...interface{}) error { return globalLogger.Warnf(format, a...) }

// Errorf formats and logs message with level ERROR.
func Errorf(format string, a ...interface{}) error { return globalLogger.Errorf(format, a...) }

// Fatalf formats and logs message with level FATAL.
func Fatalf(format string, a ...interface{}) { globalLogger.Fatalf(format, a...) }

//func Smart(err error, format string, a ...interface{}) { globalLogger.Smart(err, format, a...) }
//func SmartDebug(err error, format string, a ...interface{}) {
//	globalLogger.SmartDebug(err, format, a...)
//}
