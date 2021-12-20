package runtime

import (
	"bytes"
	"runtime"
	"strconv"
)

func GetGoroutineId() int64 {
	//g := getg()
	g := sigFetchG(nil)
	return g.goid
}

func GetGID() uint64 {
	b := make([]byte, 64)
	b = b[:runtime.Stack(b, false)]
	b = bytes.TrimPrefix(b, []byte("goroutine "))
	b = b[:bytes.IndexByte(b, ' ')]
	n, _ := strconv.ParseUint(string(b), 10, 64)
	return n
}
