package process

import (
	"sync"
	"time"

	"github.com/Julien4218/http-load-tester/config"
	log "github.com/sirupsen/logrus"
)

func Listen(processor int, testChan <-chan int, wg *sync.WaitGroup, test *config.HttpTest) {
	defer wg.Done()

	for {
		if len(testChan) > 0 {
			index := <-testChan
			executeJob(processor, index, test)
		} else {
			return
		}
	}
}

func executeJob(processor int, index int, test *config.HttpTest) {
	log.Infof("execute job processor:%d index:%d", processor, index)
	time.Sleep(time.Millisecond * 100)
}
