//go:build !windows

package fdmax

import (
	"testing"

	"github.com/stretchr/testify/require"
	"golang.org/x/sys/unix"
)

func TestSetPreservesHardLimit(t *testing.T) {
	var before unix.Rlimit
	require.NoError(t, unix.Getrlimit(unix.RLIMIT_NOFILE, &before))

	require.NoError(t, Set(1024))
	t.Cleanup(func() { _ = Set(uint64(before.Cur)) })

	var after unix.Rlimit
	require.NoError(t, unix.Getrlimit(unix.RLIMIT_NOFILE, &after))

	require.Equalf(t, before.Max, after.Max, "Set lowered hard limit: before=%d after=%d", before.Max, after.Max)
	require.Equalf(t, uint64(1024), uint64(after.Cur), "Set did not apply soft: want 1024, got %d", after.Cur)
}
