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
	case "darwin", "linux":
		currentUlimit, err := GetWithUlimit()
		require.Nil(t, err)
		current, err := Get()
		require.Nil(t, err)
		require.Equal(t, currentUlimit.Current, current.Current)
		wanted := uint64(512)
		err = Set(wanted)
		require.Nil(t, err)
		newUlimitWithCLI, err := GetWithUlimit()
		require.Nil(t, err)
		newUlimitWithAPI, err := Get()
		require.Nil(t, err)
		require.Equal(t, wanted, newUlimitWithCLI.Current)
		require.Equal(t, wanted, newUlimitWithAPI.Current)
	}
}
