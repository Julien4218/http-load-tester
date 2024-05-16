package process

import (
	"math/rand"
	"time"

	"github.com/Julien4218/http-load-tester/config"
)

type DryRunJobFunction struct {
	sleepDuration time.Duration
}

func NewDryRunJobFunction(sleepDuration time.Duration) *DryRunJobFunction {
	return &DryRunJobFunction{sleepDuration}
}

func (d *DryRunJobFunction) Execute(test *config.HttpTest) bool {
	time.Sleep(time.Millisecond * 100)
	val := rand.Intn(100)
	return val < 50
}
