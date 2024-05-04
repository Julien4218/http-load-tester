package main

import (
	"os"

	log "github.com/sirupsen/logrus"
)

func main() {
	if len(os.Args) < 2 {
		displaySyntax()
		os.Exit(1)
	}
}

func displaySyntax() {
	log.Info("execute a test on an http endpoint at a specified target RPM frequency (request per minute).")
	log.Info("syntax: http-load-tester 120")
}
