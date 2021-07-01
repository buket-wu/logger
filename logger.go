package logger

import (
	"github.com/sirupsen/logrus"
	"os"
	"sync"
)

// Level type
type Level uint32

const (
	// PanicLevel level, highest level of severity. Logs and then calls panic with the
	// message passed to Debug, Info, ...
	PanicLevel Level = iota
	// FatalLevel level. Logs and then calls `logger.Exit(1)`. It will exit even if the
	// logging level is set to Panic.
	FatalLevel
	// ErrorLevel level. Logs. Used for errors that should definitely be noted.
	// Commonly used for hooks to send errors to an error tracking service.
	ErrorLevel
	// WarnLevel level. Non-critical entries that deserve eyes.
	WarnLevel
	// InfoLevel level. General operational entries about what's going on inside the
	// application.
	InfoLevel
	// DebugLevel level. Usually only enabled when debugging. Very verbose logging.
	DebugLevel
	// TraceLevel level. Designates finer-grained informational events than the Debug.
	TraceLevel
)

func init() {
	ctxMap = &CtxMap{
		m:      make(map[int64]interface{}),
		locker: new(sync.RWMutex),
	}

	disableGoId = false
}

func InitLogrus(lev Level) {
	logrus.SetLevel(logrus.Level(lev))

	logrus.SetReportCaller(true)

	logrus.SetFormatter(&Formatter{
		HideKeys: true,
	})

	logrus.SetOutput(os.Stdout)
}
