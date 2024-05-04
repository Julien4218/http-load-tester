//go:build tools
// +build tools

package main

import (
	// build/test.mk
	_ "github.com/stretchr/testify/assert"

	// build/test.mk
	_ "gotest.tools/gotestsum"
)
