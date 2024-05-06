package config

import (
	"fmt"
)

func (c *InputConfig) Validate() []error {
	all := []error{}
	if c == nil {
		return []error{fmt.Errorf("the InputConfig has not been initialized")}
	}
	if c.MinParallel <= 0 {
		all = append(all, fmt.Errorf("attribute MinParallel must be greater than 0, received:%d", c.MinParallel))
	}
	if c.RequestPerMinute <= 0 {
		all = append(all, fmt.Errorf("attribute RequestPerMinute must be greater than 0, received:%d", c.RequestPerMinute))
	}
	if c.Loop < 0 {
		all = append(all, fmt.Errorf("attribute Loop must either be 0 or a number of minutes:%d", c.Loop))
	}

	if c.HttpTest == nil {
		all = append(all, fmt.Errorf("HttpTest is missing in the input configuration"))
	} else {
		if len(c.HttpTest.URL) == 0 {
			all = append(all, fmt.Errorf("HttpTest is missing a URL"))
		}
	}

	return all
}
