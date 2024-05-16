package process

import (
	"fmt"
	"time"

	"github.com/Julien4218/http-load-tester/config"
	log "github.com/sirupsen/logrus"
)

func Execute(config *config.InputConfig) {
	log.Info(fmt.Sprintf("Start execution with rpm:%d, loop:%d, parallel:%d on URL:%s", config.RequestPerMinute, config.Loop, config.MinParallel, config.HttpTest.URL))

	pool := NewJobPool(NewDryRunJobFunction(time.Millisecond * 100))
	for processor := 0; processor < config.MinParallel; processor++ {
		pool.CreateProcessor()
	}

	initSpec := NewBatchSpec(config.MinParallel)
	b := NewBatch(initSpec, config.HttpTest)
	count := 0
	for {
		result := b.Execute(pool)
		duration := result.Duration
		spec := GetBatchSpec(config.RequestPerMinute, int(duration.Milliseconds()), config.MinParallel)
		b = NewBatch(spec, config.HttpTest)
		pool.AdjustSize(spec, config.MinParallel, config.RequestPerMinute)
		if config.Loop > 0 && count == config.Loop {
			break
		}
		count++
		log.Infof("executing batch loop:%d", count)
	}

}
