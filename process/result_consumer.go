package process

import (
	"time"

	"github.com/Julien4218/http-load-tester/observability"
	log "github.com/sirupsen/logrus"
)

func ListenResult(success chan bool, fail chan bool) {
	for {
		if len(success) > 0 {
			<-success
			observability.GetMetrics().TestCount.Inc()
			observability.GetMetrics().SuccessCount.Inc()
			log.Infof("SUCCESS")
		}
		if len(fail) > 0 {
			<-fail
			observability.GetMetrics().TestCount.Inc()
			observability.GetMetrics().FailCount.Inc()
			log.Infof("FAIL")
		}
		time.Sleep(time.Millisecond * 10)
	}
}
