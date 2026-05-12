//go:build darwin

package fdmax

import (
	"testing"

	"github.com/stretchr/testify/require"
)

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
