package process

import (
	"time"

	"math/rand"

	"github.com/Julien4218/http-load-tester/config"
	log "github.com/sirupsen/logrus"
)

func Listen(channels *Channels, processor int, test *config.HttpTest) {
	defer channels.Done()

	for {
		job := channels.Poll()
		if job != nil {
			result := executeJob(processor, *job, test)
			channels.Complete(result)
		} else {
			return
		}
	}
}

func executeJob(processor int, index int, test *config.HttpTest) bool {
	log.Infof("execute job processor:%d index:%d", processor, index)
	time.Sleep(time.Millisecond * 100)
	val := rand.Intn(100)
	return val < 50
}
