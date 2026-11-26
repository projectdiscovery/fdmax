//go:build darwin

package fdmax

import "golang.org/x/sys/unix"

func init() {
	if v, err := unix.SysctlUint32("kern.maxfilesperproc"); err == nil && v > 0 {
		OSXMax = uint64(v)
	}
}
