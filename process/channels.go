package process

type Channels struct {
	jobs    chan int
	success chan bool
	fail    chan bool
}

func (c *Channels) Poll() *int {
	if len(c.jobs) > 0 {
		number := <-c.jobs
		return &number
	} else {
		return nil
	}
}

func (c *Channels) Complete(success bool) {
	if success {
		c.success <- true
	} else {
		c.fail <- true
	}
}
