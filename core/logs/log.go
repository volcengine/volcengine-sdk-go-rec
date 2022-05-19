package logs

import "log"

var Level = LevelWarn

type LevelEnum int

const (
	LevelError LevelEnum = iota
	LevelWarn
	LevelInfo
	LevelDebug
	LevelTrace
)

func Error(format string, v ...interface{}) {
	if Level < LevelError {
		return
	}
	log.Printf("[Error][VolcengineSDK] "+format, v...)
}

func Warn(format string, v ...interface{}) {
	if Level < LevelWarn {
		return
	}
	log.Printf("[Warn] [VolcengineSDK] "+format, v...)
}

func Info(format string, v ...interface{}) {
	if Level < LevelInfo {
		return
	}
	log.Printf("[Info] [VolcengineSDK] "+format, v...)
}

func Debug(format string, v ...interface{}) {
	if Level < LevelDebug {
		return
	}
	log.Printf("[Debug][VolcengineSDK] "+format, v...)
}

func Trace(format string, v ...interface{}) {
	if Level < LevelTrace {
		return
	}
	log.Printf("[Trace][VolcengineSDK] "+format, v...)
}
