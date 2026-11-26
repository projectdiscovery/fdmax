//go:build !windows

package fdmax

import (
	"bytes"
	"os/exec"
	"runtime"
	"strconv"

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

func GetWithUlimit() (*Limits, error) {
	if runtime.GOOS != "darwin" {
		return nil, ErrUnsupportedPlatform
	}
	cmd := exec.Command("ulimit", "-n")
	ulimitCurrent, err := cmd.Output()
	if err != nil {
		return nil, err
	}
	ulimitCurrent = bytes.ReplaceAll(ulimitCurrent, []byte("\n"), []byte(""))
	ulimitCurrentInt, err := strconv.ParseUint(string(ulimitCurrent), 10, 64)
	if err != nil {
		return nil, err
	}
	return &Limits{Current: ulimitCurrentInt, Max: ulimitCurrentInt}, nil
}

// Set raises or lowers the soft RLIMIT_NOFILE. The hard limit is never
// lowered. When the request exceeds it, privileged callers get the hard limit
// raised to match; everyone else is clamped to the existing hard limit.
func Set(maxLimit uint64) error {
	var rLimit unix.Rlimit
	if err := unix.Getrlimit(unix.RLIMIT_NOFILE, &rLimit); err != nil {
		return err
	}
	cur := getMaxLimit(maxLimit)
	if cur > rLimit.Max {
		raised := unix.Rlimit{Cur: cur, Max: cur}
		if err := unix.Setrlimit(unix.RLIMIT_NOFILE, &raised); err == nil {
			return nil
		}
		cur = rLimit.Max
	}
	rLimit.Cur = cur
	return unix.Setrlimit(unix.RLIMIT_NOFILE, &rLimit)
}

func SetWithUlimit(maxLimit uint64) error {
	if runtime.GOOS != "darwin" {
		return ErrUnsupportedPlatform
	}
	maxLimit = getMaxLimit(maxLimit)
	cmd := exec.Command("ulimit", "-n", strconv.FormatUint(maxLimit, 10))
	_, err := cmd.Output()
	return err
}

func getMaxLimit(maxLimit uint64) uint64 {
	if runtime.GOOS == "darwin" && maxLimit > OSXMax {
		return OSXMax
	}
	return maxLimit
}
