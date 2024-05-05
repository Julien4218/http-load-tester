package main

import (
	"flag"
	"fmt"
	"io/ioutil"

	"github.com/Julien4218/http-load-tester/common"
	"github.com/Julien4218/http-load-tester/process"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"gopkg.in/yaml.v3"
)

func main() {
	if err := Command.Execute(); err != nil {
		if err != flag.ErrHelp {
			log.Fatal(err)
		}
	}
}

func init() {
	Ctx = &common.CommandContext{
		InputConfig: &common.InputConfig{},
	}
}

func globalInit(cmd *cobra.Command, args []string) {
	file, err := ioutil.ReadFile(InputConfigFilepath)
	if err != nil {
		log.Error(fmt.Sprintf("A valid config filepath is required, please specify an input config file like `--config my-config.yaml`, detail:%s", err))
		log.Exit(1)
	}

	fmt.Println(string(file))
	err = yaml.Unmarshal(file, Ctx.InputConfig)
	if err != nil {
		log.Error(fmt.Sprintf("A valid YAML config filepath is required, please specify an input config file like `--config my-config.yaml`, detail:%s", err))
		log.Exit(1)
	}

	errors := Validate(Ctx.InputConfig)
	if len(errors) > 0 {
		for _, error := range errors {
			log.Error(error)
		}
		log.Exit(1)
	}
}

var (
	InputConfigFilepath string
	Ctx                 *common.CommandContext
)

var Command = &cobra.Command{
	Use:              "http-load-tester",
	Short:            "Http Load Tester",
	PersistentPreRun: globalInit,
	Long:             `Execute a test on an http endpoint at a specified target RPM frequency (request per minute). Environment variable replacement is allowed with the syntax [env:MY_VAR_NAME] anywhere in the input config file`,
	Run: func(cmd *cobra.Command, args []string) {
		process.Execute(Ctx)

	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func init() {
	Command.Flags().IntVar(&Ctx.RPM, "rpm", 60, "RPM as number of http Request Per Minute")
	Command.Flags().IntVar(&Ctx.Loop, "loop", 0, "Loop (0 for continuously looping)")
	Command.Flags().StringVar(&InputConfigFilepath, "config", "", "Input config file")
	_ = Command.MarkFlagRequired("config")
}

func Validate(inputConfig *common.InputConfig) []string {
	errors := []string{}
	if inputConfig.MinParallel <= 0 {
		errors = append(errors, fmt.Sprintf("MinParallel must be greater than 0, received:%d", inputConfig.MinParallel))
	}

	if inputConfig.HttpTest == nil {
		errors = append(errors, "HttpTest is missing in the input configuration")
	}

	if len(Ctx.InputConfig.HttpTest.URL) == 0 {
		errors = append(errors, "HttpTest is missing a URL")
	}
	return errors
}
