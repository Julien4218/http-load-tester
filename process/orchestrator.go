package process

import (
	"fmt"

	"github.com/Julien4218/http-load-tester/config"
	log "github.com/sirupsen/logrus"
)

func Execute(config *config.InputConfig) {
	log.Info(fmt.Sprintf("Start execution with rpm:%d, loop:%d, parallel:%d on URL:%s", config.RequestPerMinute, config.Loop, config.MinParallel, config.HttpTest.URL))

	pool := NewJobPool()
	for processor := 0; processor < config.MinParallel; processor++ {
		pool.CreateProcessor()
	}

	b := NewBatch(3, config.HttpTest)
	b.Execute(pool)

}
