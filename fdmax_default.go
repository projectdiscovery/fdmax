//go:build !windows

package fdmax

import (
	"runtime"

	"golang.org/x/sys/unix"
)

// Get the current limits
func Get() (*Limits, error) {
	var rLimit unix.Rlimit
	err := unix.Getrlimit(unix.RLIMIT_NOFILE, &rLimit)
	if err != nil {
		return nil, err
	}

	return &Limits{Current: uint64(rLimit.Cur), Max: uint64(rLimit.Max)}, nil
}

func Set(maxLimit uint64) error {
	var rLimit unix.Rlimit
	rLimit.Max = maxLimit

	rLimit.Cur = maxLimit
	// https://github.com/golang/go/issues/30401
	if runtime.GOOS == "darwin" && rLimit.Cur > OSXMax {
		rLimit.Cur = OSXMax
	}

	return unix.Setrlimit(unix.RLIMIT_NOFILE, &rLimit)
}
