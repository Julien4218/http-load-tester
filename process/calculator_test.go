package process

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test60RPM_1(t *testing.T) {
	rpm := 60
	duration_ms := 100
	parallel := 1
	result := getTiming(rpm, duration_ms, parallel)

	require.Equal(t, int64(1000), result.WaitTime.Milliseconds())
	require.Equal(t, 1, result.TargetParallel)
}

func Test60RPM_5(t *testing.T) {
	rpm := 60
	duration_ms := 100
	parallel := 5
	result := getTiming(rpm, duration_ms, parallel)

	require.Equal(t, int64(5000), result.WaitTime.Milliseconds())
	require.Equal(t, 5, result.TargetParallel)
}

func Test6000RPM_5(t *testing.T) {
	rpm := 6000
	duration_ms := 100
	parallel := 5
	result := getTiming(rpm, duration_ms, parallel)

	require.Equal(t, int64(200), result.WaitTime.Milliseconds())
	require.Equal(t, 20, result.TargetParallel)
}
