package process

import (
	"math"
	"time"

	log "github.com/sirupsen/logrus"
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

func GetBatchSpec(rpm int, duration_ms int, parallel int) *BatchSpec {
	result := NewBatchSpec(rpm)

	rpms := float64(rpm) / 60000
	mspr := 1 / rpms
	min_duration_ms := float64(2 * duration_ms)
	math_min_duration := float64(mspr)
	if min_duration_ms > math_min_duration {
		math_min_duration = min_duration_ms
	}
	math_min_parallel := int(math.Round(math_min_duration * float64(rpms)))
	result.TargetParallel = parallel
	if math_min_parallel > parallel {
		result.TargetParallel = math_min_parallel
	}
	actual_ms := float64(result.TargetParallel) / rpms
	result.MaxWaitTime = time.Duration(actual_ms) * time.Millisecond

	log.Infof("calculated next batch target:%d maxWaitTime:%dms", result.TargetParallel, result.MaxWaitTime.Milliseconds())

	return result
}
