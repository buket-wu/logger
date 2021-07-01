package logger

import (
	"github.com/sirupsen/logrus"
	"testing"
)

func init() {
	InitLogrus(logrus.TraceLevel)
}

func TestLogger_print(t *testing.T) {
	//SetCtxId(uuid.New())
	Trace("sss")
}
