package common

type CommandContext struct {
	RPM         int
	Loop        int
	InputConfig *InputConfig
}

type InputConfig struct {
	MinParallel int       `yaml:"MinParallel"`
	HttpTest    *HttpTest `yaml:"HttpTest"`
}

type HttpTest struct {
	URL                  string            `yaml:"URL"`
	Headers              map[string]string `yaml:"Headers"`
	Body                 string            `yaml:"Body"`
	SuccessResponseCodes []int             `yaml:"SuccessResponseCodes"`
}
