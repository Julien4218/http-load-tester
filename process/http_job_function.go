package process

import (
	"github.com/Julien4218/http-load-tester/client"
	"github.com/Julien4218/http-load-tester/config"
)

type HttpJobFunction struct {
}

func NewHttpJobFunction() *HttpJobFunction {
	return &HttpJobFunction{}
}

func (d *HttpJobFunction) Execute(test *config.HttpTest) bool {
	return client.Execute(test)
}
