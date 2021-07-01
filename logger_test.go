package logger

import (
	"testing"
)

func init() {
	InitLogrus(TraceLevel)
}

func TestLogger_print(t *testing.T) {
	//SetCtxId(uuid.New())
	Trace("sss")
}
