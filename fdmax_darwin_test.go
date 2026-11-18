//go:build darwin

package fdmax

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestOSXMaxFromSysctl(t *testing.T) {
	require.GreaterOrEqual(t, OSXMax, uint64(24576))
	require.LessOrEqual(t, OSXMax, uint64(1<<22))
}

func TestFileDescriptors(t *testing.T) {
	before, err := Get()
	require.Nil(t, err)

	wanted := uint64(444)
	require.Nil(t, Set(wanted))
	t.Cleanup(func() { _ = Set(before.Current) })

	after, err := Get()
	require.Nil(t, err)
	require.Equal(t, wanted, after.Current)
	require.Equal(t, before.Max, after.Max)
	require.NotEqual(t, before.Current, after.Current)
}
