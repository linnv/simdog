package simdog

import (
	"os"
	"sync"

	"github.com/linnv/logx"
	"github.com/shirou/gopsutil/process"
)

var once sync.Once
var pid int32
var CurProc *process.Process

func init() {
	once.Do(func() {
		pid = int32(os.Getpid())
		//@TODO
		p, err := process.NewProcess(pid)
		if err != nil {
			logx.Debugf("Cannot read process info: %v", err)
		}
		CurProc = p
	})
}
