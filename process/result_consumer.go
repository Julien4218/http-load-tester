package process

import (
	"time"

	log "github.com/sirupsen/logrus"
)

func ListenResult(success chan bool, fail chan bool) {
	for {
		if len(success) > 0 {
			<-success
			log.Infof("SUCCESS")
		}
		if len(fail) > 0 {
			<-fail
			log.Infof("FAIL")
		}
		time.Sleep(time.Millisecond * 10)
	}
}
