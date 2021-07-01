package logger

import (
	"github.com/petermattis/goid"
	"sync"
)

var (
	ctxMap      *CtxMap
	disableGoId bool
)

type CtxMap struct {
	m      map[int64]interface{}
	locker *sync.RWMutex
}

func getGoId() int64 {
	return goid.Get()
}

func getCtxId() interface{} {
	goId := getGoId()
	ctxMap.locker.RLock()
	ctxId, ok := ctxMap.m[goId]
	ctxMap.locker.RUnlock()
	if !ok {
		ctxId = goId
	}

	return ctxId
}

func SetCtxId(ctxId interface{}) {
	goId := getGoId()
	ctxMap.locker.Lock()
	ctxMap.m[goId] = ctxId
	ctxMap.locker.Unlock()
}

func CleatCtxId() {
	goId := getGoId()
	ctxMap.locker.Lock()
	delete(ctxMap.m, goId)
	ctxMap.locker.Unlock()
}

func SetDisableGoId(v bool) {
	disableGoId = v
}
