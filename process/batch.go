package process

import (
	"time"

	"github.com/Julien4218/http-load-tester/config"
	"github.com/Julien4218/http-load-tester/observability"
	log "github.com/sirupsen/logrus"
)

type Batch struct {
	size     int
	httpTest *config.HttpTest
}

type BatchResult struct {
	Duration time.Duration
}

func NewBatch(size int, httpTest *config.HttpTest) *Batch {
	return &Batch{
		size:     size,
		httpTest: httpTest,
	}
}

func (b *Batch) Execute(pool *JobPool) *BatchResult {
	result := &BatchResult{}

	channels := &Channels{
		jobs:         make(chan int, b.size),
		success:      make(chan bool, b.size),
		fail:         make(chan bool, b.size),
		job_duration: make(chan time.Duration, b.size),
	}

	for i := 0; i < b.size; i++ {
		channels.jobs <- i
	}

	done := make(chan bool, 1)
	go ListenResult(channels.success, channels.fail, done)

	log.Info("Waiting for completion")
	pool.Start(channels, b.httpTest)
	pool.WaitForCompletion()

	duration := getJobDuration(channels.job_duration)
	if duration != nil {
		result.Duration = *duration
		observability.GetMetrics().ElapsedTimeMs.Set(float64(duration.Milliseconds()))
	}

	log.Infof("Each test duration average:%dMs", duration.Milliseconds())

	observability.HarvestNow()
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
