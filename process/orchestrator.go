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
	testChan := make(chan int, config.RequestPerMinute)
	var wg sync.WaitGroup

	for i := 0; i < config.RequestPerMinute; i++ {
		testChan <- i
	}

	start := time.Now()
	for i := 0; i < config.MinParallel; i++ {
		log.Infof("Adding processor:%d", i)
		wg.Add(1)
		go Listen(i, testChan, &wg, config.HttpTest)
	}
	log.Info("Waiting for completion")
	wg.Wait()
	duration := time.Since(start)
	each := duration.Milliseconds() / int64(config.RequestPerMinute)
	log.Infof("Total duration:%dMs each:%dMs", duration.Milliseconds(), each)
}
