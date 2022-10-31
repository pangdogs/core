package timer

import (
	"github.com/pangdogs/galaxy/runtime"
	"time"
)

type Timer interface {
	AfterFunc(dur time.Duration, fun func()) Handle
	TickFunc(interval time.Duration, count int, fun func()) Handle
}

type Handle struct {
}

func newTimer(...any) Timer {

}

type _Timer struct {
}

func (t *_Timer) Init(ctx runtime.Context) {

}
