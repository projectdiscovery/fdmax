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

func TestSetAboveHardLimitNeverLowersHard(t *testing.T) {
	var before unix.Rlimit
	require.NoError(t, unix.Getrlimit(unix.RLIMIT_NOFILE, &before))
	if before.Max > 1<<30 {
		t.Skip("hard limit effectively unlimited")
	}
	t.Cleanup(func() { _ = Set(uint64(before.Cur)) })

	require.NoError(t, Set(uint64(before.Max)+1000))

	var after unix.Rlimit
	require.NoError(t, unix.Getrlimit(unix.RLIMIT_NOFILE, &after))
	require.GreaterOrEqual(t, uint64(after.Max), uint64(before.Max), "hard limit was lowered")
	require.LessOrEqual(t, uint64(after.Cur), uint64(after.Max))
	require.GreaterOrEqual(t, uint64(after.Cur), uint64(before.Cur), "soft limit shrank")
}
