package process

import (
	"math"
	"time"

	log "github.com/sirupsen/logrus"
)

const (
	MULTIPLIER_BUFFER float64 = 1.1
)

type BatchSpec struct {
	TargetParallel int
	MaxWaitTime    time.Duration
}

func NewBatchSpec(rpm int) *BatchSpec {
	return &BatchSpec{
		TargetParallel: rpm,
		MaxWaitTime:    time.Duration(0),
	}
}

func GetBatchSpec(rpm int, duration time.Duration, parallel int) *BatchSpec {
	result := NewBatchSpec(rpm)

	rpms := float64(rpm) / 60000
	mspr := 1 / rpms
	buffer_duration_ms := float64(duration.Milliseconds()) * MULTIPLIER_BUFFER
	min_duration := float64(mspr)
	if buffer_duration_ms > min_duration {
		min_duration = buffer_duration_ms
	}
	min_parallel := int(math.Round(min_duration * float64(rpms)))
	result.TargetParallel = parallel
	if min_parallel > parallel {
		result.TargetParallel = min_parallel
	}
	actual_ms := float64(result.TargetParallel) / rpms
	result.MaxWaitTime = time.Duration(actual_ms) * time.Millisecond

	previous_mspr := float64(duration.Milliseconds()) / float64(parallel)
	if previous_mspr >= mspr {
		result.TargetParallel = int(float64(result.TargetParallel) * MULTIPLIER_BUFFER)
	}

	log.Infof("calculated next batch target:%d maxWaitTime:%dms", result.TargetParallel, result.MaxWaitTime.Milliseconds())

	return result
}
