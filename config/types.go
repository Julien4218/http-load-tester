package config

type HttpTest struct {
	URL                  string            `yaml:"URL"`
	Method               string            `yaml:"Method"`
	Headers              map[string]string `yaml:"Headers"`
	Body                 string            `yaml:"Body"`
	SingleLineBody       bool              `yaml:"SingleLineBody"`
	SuccessResponseCodes []int             `yaml:"SuccessResponseCodes"`
	SuccessJqQuery       string            `yaml:"SuccessJqQuery"`
}

type InputConfig struct {
	MinParallel      int       `yaml:"MinParallel"`
	RequestPerMinute int       `yaml:"RequestPerMinute"`
	Loop             int       `yaml:"Loop"`
	HttpTest         *HttpTest `yaml:"HttpTest"`
}
