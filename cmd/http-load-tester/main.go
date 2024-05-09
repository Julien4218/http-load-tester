package main

import (
	"flag"

	"github.com/Julien4218/http-load-tester/config"
	"github.com/Julien4218/http-load-tester/observability"
	"github.com/Julien4218/http-load-tester/process"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var (
	InputConfig         *config.InputConfig
	InputConfigFilepath string
)

func main() {
	if err := Command.Execute(); err != nil {
		if err != flag.ErrHelp {
			log.Fatal(err)
		}
	}
}

func init() {
	Command.Flags().StringVar(&InputConfigFilepath, "config", "", "Input config file")
	_ = Command.MarkFlagRequired("config")
}

func globalInit(cmd *cobra.Command, args []string) {
	go observability.Init()

	var err error
	InputConfig, err = config.Init(InputConfigFilepath)
	if err != nil {
		log.Error(err)
		log.Exit(1)
	}

	errors := InputConfig.Validate()
	if len(errors) > 0 {
		for _, err = range errors {
			log.Error(err)
		}
		log.Exit(1)
	}
}

var Command = &cobra.Command{
	Use:              "http-load-tester",
	Short:            "Http Load Tester",
	PersistentPreRun: globalInit,
	Long:             `Execute a test on an http endpoint at a specified target RPM frequency (request per minute). Environment variable replacement is allowed with the syntax [env:MY_VAR_NAME] anywhere in the input config file`,
	Run: func(cmd *cobra.Command, args []string) {
		process.Execute(InputConfig)
	},
}
