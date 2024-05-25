package process

import (
	"fmt"
	"time"

	"github.com/Julien4218/http-load-tester/config"
	"github.com/Julien4218/http-load-tester/observability"
	log "github.com/sirupsen/logrus"
)

func Execute(config *config.InputConfig, dryRun bool) {
	log.Info(fmt.Sprintf("Start execution with rpm:%d, loop:%d, parallel:%d on URL:%s", config.RequestPerMinute, config.Loop, config.MinParallel, config.HttpTest.URL))

	var function JobFunction
	if dryRun {
		function = NewDryRunJobFunction(time.Millisecond * 100)
	} else {
		function = NewHttpJobFunction()
	}

	pool := NewJobPool(function)
	for processor := 0; processor < config.MinParallel; processor++ {
		pool.CreateProcessor()
	}

	spec := NewBatchSpec(config.MinParallel)
	b := NewBatch(spec, config.HttpTest)
	count := 1
	for {
		lastStart := time.Now()
		log.Infof("executing batch loop:%d", count)
		result := b.Execute(pool)

		batchDuration := time.Since(lastStart)
		delay := b.spec.MaxWaitTime - batchDuration
		if delay > 0 {
			log.Infof("pacing to match desired rpm, waiting %dms", delay.Milliseconds())
			time.Sleep(delay)
		} else {
			log.Infof("no pacing required, batch duration of %dms is longer than maximum tolerated of %dms", batchDuration.Milliseconds(), b.spec.MaxWaitTime.Milliseconds())
		}
		finalDuration := time.Since(lastStart)
		rpmPace := int64(spec.TargetParallel) * (60000 / finalDuration.Milliseconds())
		log.Infof("completed:%d in %dms actualPace:%d", spec.TargetParallel, finalDuration.Milliseconds(), rpmPace)
		observability.GetMetrics().RpmPace.Set(float64(rpmPace))
		observability.HarvestNow()

		// prepare next
		spec = GetBatchSpec(config.RequestPerMinute, result.Duration, spec.TargetParallel)
		b = NewBatch(spec, config.HttpTest)
		pool.AdjustSize(spec, config.MinParallel, config.RequestPerMinute)

		// process next
		if config.Loop > 0 && count == config.Loop {
			break
		}
		count++
	}

}
