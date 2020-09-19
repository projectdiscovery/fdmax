package fdmax

import (
	"runtime"
	"syscall"
)

const (
	UnixMax uint64 = 999999
	OSXMax  uint64 = 24576
)

type Limits struct {
	Current uint64
	Max     uint64
}

func Get() (*Limits, error) {
	var rLimit syscall.Rlimit
	err := syscall.Getrlimit(syscall.RLIMIT_NOFILE, &rLimit)
	if err != nil {
		return nil, err
	}

	return &Limits{Current: rLimit.Cur, Max: rLimit.Max}, nil
}

func Set(maxLimit uint64) error {
	var rLimit syscall.Rlimit
	rLimit.Max = maxLimit

	rLimit.Cur = maxLimit
	// https://github.com/golang/go/issues/30401
	if runtime.GOOS == "darwin" && rLimit.Cur > OSXMax {
		rLimit.Cur = OSXMax
	}

	return syscall.Setrlimit(syscall.RLIMIT_NOFILE, &rLimit)
}
