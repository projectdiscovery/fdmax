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
		require.ErrorIs(t, err, ErrUnsupportedPlatform)
		require.Nil(t, res)
		err = Set(1)
		require.ErrorIs(t, err, ErrUnsupportedPlatform)
	case "darwin":
		currentUlimit, err := GetWithUlimit()
		require.Nil(t, err)
		current, err := Get()
		require.Nil(t, err)
		require.Equal(t, currentUlimit.Current, current.Current)
		wanted := uint64(444)
		err = Set(wanted)
		require.Nil(t, err)
		newUlimitWithCLI, err := GetWithUlimit()
		require.Nil(t, err)
		newUlimitWithAPI, err := Get()
		require.Nil(t, err)
		require.Equal(t, wanted, newUlimitWithCLI.Current)
		require.Equal(t, wanted, newUlimitWithAPI.Current)
		require.True(t, currentUlimit.Current != newUlimitWithAPI.Current)
	case "linux":
		currentUlimit, err := Get()
		require.Nil(t, err)
		require.Positive(t, currentUlimit.Current)
		wanted := uint64(444)
		err = Set(wanted)
		require.Nil(t, err)
		newUlimitWithAPI, err := Get()
		require.Nil(t, err)
		require.Equal(t, wanted, newUlimitWithAPI.Current)
		require.True(t, currentUlimit.Current != newUlimitWithAPI.Current)
	}
}
