package util

import (
	"bufio"
	"fmt"
	"io"
	"os"

	"github.com/linnv/logx"
)

type NetState struct {
	InuseTCP int64 `json:"InuseTCP"`
	InuseUDP int64 `json:"InuseUDP"`
}

//LoadNetState implements load file  `/proc/$pid/net/sockstat`
//linux only
func LoadNetState(pid int) NetState {
	netstatPath := fmt.Sprintf("/proc/%d/net/sockstat", pid)
	return loadNetState(netstatPath)
	// readline
	// bs,err:=ioutils.ReadAll(netstatPath)
	// logx.CheckErr(err)
}

func loadNetState(p string) NetState {
	file, err := os.Open(p)
	logx.CheckErr(err)
	defer file.Close()

	reader := bufio.NewReader(file)

	for {
		line, bs, err := reader.ReadLine()

		if err == io.EOF {
			break
		}

		logx.Debugf("bs: %+v line: %s\n", bs, line)
	}
	ns := NetState{}
	return ns
}
