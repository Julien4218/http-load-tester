package process

import (
	"fmt"
	"sync"
	"time"

	"github.com/Julien4218/http-load-tester/config"
	log "github.com/sirupsen/logrus"
)

func Execute(config *config.InputConfig) {
	log.Info(fmt.Sprintf("Start execution with rpm:%d, loop:%d, parallel:%d on URL:%s", config.RequestPerMinute, config.Loop, config.MinParallel, config.HttpTest.URL))

	channels := &Channels{
		jobs:    make(chan int, config.RequestPerMinute),
		success: make(chan bool, config.RequestPerMinute),
		fail:    make(chan bool, config.RequestPerMinute),
		wg:      &sync.WaitGroup{},
	}

	go ListenResult(channels.success, channels.fail)

	for i := 0; i < config.RequestPerMinute; i++ {
		channels.jobs <- i
	}

	start := time.Now()
	for processor := 0; processor < config.MinParallel; processor++ {
		log.Infof("Adding processor:%d", processor)
		channels.wg.Add(1)
		go Listen(channels, processor, config.HttpTest)
	}
	log.Info("Waiting for completion")
	channels.wg.Wait()
	duration := time.Since(start)
	each := duration.Milliseconds() / int64(config.RequestPerMinute)
	log.Infof("Total duration:%dMs each:%dMs", duration.Milliseconds(), each)
}
