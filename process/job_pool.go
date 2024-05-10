package process

import (
	"sync"

	"github.com/Julien4218/http-load-tester/config"
	log "github.com/sirupsen/logrus"
)

type JobPool struct {
	processors []*JobProcessor
	waitGroup  *sync.WaitGroup
}

func NewJobPool() *JobPool {
	return &JobPool{
		processors: []*JobProcessor{},
		waitGroup:  &sync.WaitGroup{},
	}
}

func (p *JobPool) CreateProcessor() {
	id := len(p.processors) + 1
	processor := NewJobProcessor(id, p.waitGroup)
	p.processors = append(p.processors, processor)
	log.Infof("Adding processor:%d", id)
}

func (p *JobPool) Start(channels *Channels, httpTest *config.HttpTest) {
	// Finish any work before starting new ones
	p.WaitForCompletion()
	// Start processors
	for _, processor := range p.processors {
		p.waitGroup.Add(1)
		go processor.ListenJob(channels, httpTest)
	}
}

func (p *JobPool) WaitForCompletion() {
	p.waitGroup.Wait()
}

func (p *JobPool) Size() int {
	return len(p.processors)
}
