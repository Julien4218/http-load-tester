package process

import (
	"math"
	"time"
)

type CalculatorResult struct {
	TargetParallel int
	WaitTime       time.Duration
}

func getTiming(rpm int, duration_ms int, parallel int) *CalculatorResult {
	result := &CalculatorResult{}

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
	result.WaitTime = time.Duration(actual_ms) * time.Millisecond

	return result
}
