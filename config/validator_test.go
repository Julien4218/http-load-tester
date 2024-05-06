package config

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestShouldValidate(t *testing.T) {
	config := &InputConfig{
		MinParallel:      4,
		RequestPerMinute: 120,
		Loop:             0,
		HttpTest: &HttpTest{
			URL: "http://localhost",
		},
	}
	errors := config.Validate()
	require.Equal(t, len(errors), 0)
}

func TestShouldValidateMinParallel(t *testing.T) {
	config := &InputConfig{
		MinParallel:      0,
		RequestPerMinute: 120,
		HttpTest: &HttpTest{
			URL: "http://localhost",
		},
	}
	errors := config.Validate()
	require.NotEqual(t, len(errors), 0)
	for _, err := range errors {
		if strings.Contains(err.Error(), "MinParallel") {
			return
		}
	}
	require.Fail(t, "no validation error for MinParallel")
}

func TestShouldValidateRequestPerMinute(t *testing.T) {
	config := &InputConfig{
		MinParallel:      4,
		RequestPerMinute: 0,
		HttpTest: &HttpTest{
			URL: "http://localhost",
		},
	}
	errors := config.Validate()
	require.NotEqual(t, len(errors), 0)
	for _, err := range errors {
		if strings.Contains(err.Error(), "RequestPerMinute") {
			return
		}
	}
	require.Fail(t, "no validation error for RequestPerMinute")
}

func TestShouldValidateLoop(t *testing.T) {
	config := &InputConfig{
		MinParallel:      4,
		RequestPerMinute: 120,
		Loop:             -1,
		HttpTest: &HttpTest{
			URL: "http://localhost",
		},
	}
	errors := config.Validate()
	require.NotEqual(t, len(errors), 0)
	for _, err := range errors {
		if strings.Contains(err.Error(), "Loop") {
			return
		}
	}
	require.Fail(t, "no validation error for Loop")
}

func TestShouldValidateHttpTestExist(t *testing.T) {
	config := &InputConfig{
		MinParallel:      4,
		RequestPerMinute: 120,
	}
	errors := config.Validate()
	require.NotEqual(t, len(errors), 0)
	for _, err := range errors {
		if strings.Contains(err.Error(), "HttpTest") {
			return
		}
	}
	require.Fail(t, "no validation error for HttpTest")
}

func TestShouldValidateHttpTestUrlExist(t *testing.T) {
	config := &InputConfig{
		MinParallel:      4,
		RequestPerMinute: 120,
		HttpTest:         &HttpTest{},
	}
	errors := config.Validate()
	require.NotEqual(t, len(errors), 0)
	for _, err := range errors {
		if strings.Contains(err.Error(), "URL") {
			return
		}
	}
	require.Fail(t, "no validation error for HttpTest URL")
}
