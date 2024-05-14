package process

import (
	"sync"
	"time"

	"math/rand"

	"github.com/Julien4218/http-load-tester/config"
	log "github.com/sirupsen/logrus"
)

type JobProcessor struct {
	id int
	wg *sync.WaitGroup
}

func NewJobProcessor(id int, wg *sync.WaitGroup) *JobProcessor {
	return &JobProcessor{
		id,
		wg,
	}
}

func (p *JobProcessor) ListenJob(channels *Channels, test *config.HttpTest) {
	defer p.wg.Done()
	for {
		job := channels.Poll()
		if job != nil {
			start := time.Now()
			result := p.executeJob(*job, test)
			duration := time.Since(start)
			channels.Complete(result, duration)
			continue
		} else {
			return
		}
	}
}

func (p *JobProcessor) executeJob(index int, test *config.HttpTest) bool {
	log.Infof("execute job processor:%d index:%d", p.id, index)
	time.Sleep(time.Millisecond * 100)
	val := rand.Intn(100)
	return val < 50
}
