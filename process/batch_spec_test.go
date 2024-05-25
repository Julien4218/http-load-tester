package process

import (
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func Test60RPM_1(t *testing.T) {
	rpm := 60
	duration := 100 * time.Millisecond
	parallel := 1
	result := GetBatchSpec(rpm, duration, parallel)

	require.Equal(t, int64(1000), result.MaxWaitTime.Milliseconds())
	require.Equal(t, 1, result.TargetParallel)
}

func Test60RPM_4(t *testing.T) {
	rpm := 60
	duration := 100 * time.Millisecond
	parallel := 4
	result := GetBatchSpec(rpm, duration, parallel)

	require.Equal(t, int64(4000), result.MaxWaitTime.Milliseconds())
	require.Equal(t, 4, result.TargetParallel)
}

func Test60RPM_5(t *testing.T) {
	rpm := 60
	duration := 100 * time.Millisecond
	parallel := 5
	result := GetBatchSpec(rpm, duration, parallel)

	require.Equal(t, int64(5000), result.MaxWaitTime.Milliseconds())
	require.Equal(t, 5, result.TargetParallel)
}

func Test6000RPM_5(t *testing.T) {
	rpm := 6000
	duration := 100 * time.Millisecond
	parallel := 5
	result := GetBatchSpec(rpm, duration, parallel)

	require.Equal(t, int64(110), result.MaxWaitTime.Milliseconds())
	require.Equal(t, 12, result.TargetParallel)
}

func Test60000RPM_5(t *testing.T) {
	rpm := 60000
	duration := 100 * time.Millisecond
	parallel := 5
	result := GetBatchSpec(rpm, duration, parallel)

	require.Equal(t, int64(110), result.MaxWaitTime.Milliseconds())
	require.Equal(t, 121, result.TargetParallel)
}

func Test60000RPM_back(t *testing.T) {
	rpm := 60000
	duration := 300 * time.Millisecond
	parallel := 5
	result := GetBatchSpec(rpm, duration, parallel)

	require.Equal(t, int64(330), result.MaxWaitTime.Milliseconds())
	require.Equal(t, 363, result.TargetParallel)
}
