package process

import (
	"time"

	"github.com/Julien4218/http-load-tester/config"
	"github.com/Julien4218/http-load-tester/observability"
	log "github.com/sirupsen/logrus"
)

type Batch struct {
	spec     *BatchSpec
	httpTest *config.HttpTest
}

type BatchResult struct {
	Duration time.Duration
}

func NewBatch(spec *BatchSpec, httpTest *config.HttpTest) *Batch {
	return &Batch{
		spec:     spec,
		httpTest: httpTest,
	}
}

func (b *Batch) Execute(pool *JobPool) *BatchResult {
	log.Infof("Executing batch size:%d pool:%d", b.spec.TargetParallel, pool.Size())

	channels := &Channels{
		jobs:         make(chan int, b.spec.TargetParallel),
		success:      make(chan bool, b.spec.TargetParallel),
		fail:         make(chan bool, b.spec.TargetParallel),
		job_duration: make(chan time.Duration, b.spec.TargetParallel),
	}

	for i := 0; i < b.spec.TargetParallel; i++ {
		channels.jobs <- i
	}

	done := make(chan bool, 1)
	go ListenResult(channels.success, channels.fail, done)

	log.Info("Waiting for completion")
	pool.Start(channels, b.httpTest)
	pool.WaitForCompletion()

	result := &BatchResult{}
	duration := getJobDuration(channels.job_duration)
	if duration != nil {
		result.Duration = *duration
		observability.GetMetrics().ElapsedTimeMs.Set(float64(duration.Milliseconds()))
	}

	log.Infof("Each test duration average:%dMs", duration.Milliseconds())

	done <- true

	return result
}

func getJobDuration(job_duration chan time.Duration) *time.Duration {
	var total time.Duration
	count := 0
	close(job_duration)
	for val := range job_duration {
		total += val
		count++
	}
	if count > 0 {
		total = total / time.Duration(count)
		return &total
	}
	return nil
}
