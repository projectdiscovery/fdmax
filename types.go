package fdmax

import "errors"

var (
	// ErrUnsupportedPlatform error if the platform doesn't support file descriptor increase via system api
	ErrUnsupportedPlatform = errors.New("unsupported platform")
)

// UnixMax on unix systems.
const UnixMax uint64 = 999999

// OSXMax on darwin. Populated from kern.maxfilesperproc at init on darwin;
// the literal here is a fallback for other platforms and when the sysctl fails.
var OSXMax uint64 = 24576

// Limits contains the file system descriptor limits
type Limits struct {
	Current uint64
	Max     uint64
}
