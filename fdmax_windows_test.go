//go:build windows

package fdmax

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestFileDescriptors(t *testing.T) {
	res, err := Get()
	require.ErrorIs(t, err, ErrUnsupportedPlatform)
	require.Nil(t, res)
	err = Set(1)
	require.ErrorIs(t, err, ErrUnsupportedPlatform)
}
