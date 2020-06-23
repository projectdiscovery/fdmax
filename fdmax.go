package fdmax

import (
	"syscall"
)

const (
	Max uint64 = 999999
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
	return syscall.Setrlimit(syscall.RLIMIT_NOFILE, &rLimit)
}
