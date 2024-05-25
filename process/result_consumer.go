package process

import (
	"time"

	"github.com/Julien4218/http-load-tester/observability"
	log "github.com/sirupsen/logrus"
)

func ListenResult(success chan bool, fail chan bool, done chan bool) {
	for {
		if len(success) > 0 {
			<-success
			observability.GetMetrics().TestCount.Inc()
			observability.GetMetrics().SuccessCount.Inc()
			log.Debugf("SUCCESS")
			continue
		}
		if len(fail) > 0 {
			<-fail
			observability.GetMetrics().TestCount.Inc()
			observability.GetMetrics().FailCount.Inc()
			log.Errorf("FAIL")
			continue
		}
		if len(done) > 0 {
			isDone := <-done
			if isDone {
				close(done)
				return
			}
		}
		time.Sleep(time.Millisecond * 10)
	}
}
