package process

import "time"

type Channels struct {
	jobs         chan int
	success      chan bool
	fail         chan bool
	job_duration chan time.Duration
}

func (c *Channels) Poll() *int {
	if len(c.jobs) > 0 {
		number := <-c.jobs
		return &number
	} else {
		return nil
	}
}

func (c *Channels) Complete(success bool, job_duration time.Duration) {
	if success {
		c.success <- true
	} else {
		c.fail <- true
	}
	c.job_duration <- job_duration
}
