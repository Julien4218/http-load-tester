package process

import (
	"sync"
	"time"

	"github.com/Julien4218/http-load-tester/config"
	log "github.com/sirupsen/logrus"
)

type JobFunction interface {
	Execute(test *config.HttpTest) bool
}

type JobProcessor struct {
	id       int
	wg       *sync.WaitGroup
	function JobFunction
}

func NewJobProcessor(id int, wg *sync.WaitGroup, function JobFunction) *JobProcessor {
	return &JobProcessor{
		id,
		wg,
		function,
	}
}

func (p *JobProcessor) ListenJob(channels *Channels, test *config.HttpTest) {
	defer p.wg.Done()
	for {
		job := channels.Poll()
		if job != nil {
			start := time.Now()
			log.Infof("execute job processor:%d index:%d", p.id, *job)
			result := p.function.Execute(test)
			duration := time.Since(start)
			channels.Complete(result, duration)
			continue
		} else {
			return
		}
	}
}
