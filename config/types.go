package config

type HttpTest struct {
	URL                  string            `yaml:"URL"`
	Headers              map[string]string `yaml:"Headers"`
	Body                 string            `yaml:"Body"`
	SuccessResponseCodes []int             `yaml:"SuccessResponseCodes"`
}

type InputConfig struct {
	MinParallel      int       `yaml:"MinParallel"`
	RequestPerMinute int       `yaml:"RequestPerMinute"`
	Intervals        int       `yaml:"Intervals"`
	Loop             int       `yaml:"Loop"`
	HttpTest         *HttpTest `yaml:"HttpTest"`
}
