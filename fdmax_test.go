package fdmax

import (
	"runtime"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestFileDescriptors(t *testing.T) {
	switch runtime.GOOS {
	case "windows":
		res, err := Get()
		require.ErrorIs(t, err, UnsupportedPlatformError)
		require.Nil(t, res)
		err = Set(1)
		require.ErrorIs(t, err, UnsupportedPlatformError)
	}
}
