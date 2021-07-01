package logger

import (
	"github.com/sirupsen/logrus"
	"os"
	"sync"
)

func init() {
	ctxMap = &CtxMap{
		m:      make(map[int64]interface{}),
		locker: new(sync.RWMutex),
	}

	disableGoId = false
}

func InitLogrus(lev logrus.Level) {
	logrus.SetLevel(lev)

	logrus.SetReportCaller(true)

	logrus.SetFormatter(&Formatter{
		HideKeys: true,
	})

	logrus.SetOutput(os.Stdout)
}
