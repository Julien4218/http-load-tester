package client

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestShouldReplaceEnv(t *testing.T) {
	input := "data with env::MY_KEY to replace"
	t.Setenv("MY_KEY", "my-secret-value")
	output, err := replaceEnvVar(input)
	assert.Nil(t, err)
	assert.Equal(t, "data with my-secret-value to replace", output)
}

func TestShouldReplaceAllMatchingEnv(t *testing.T) {
	input := "data with env::MY_KEY to replace and another env::MY_KEY to replace too"
	t.Setenv("MY_KEY", "my-secret-value")
	output, err := replaceEnvVar(input)
	assert.Nil(t, err)
	assert.Equal(t, "data with my-secret-value to replace and another my-secret-value to replace too", output)
}

func TestShouldNotReplaceWithErrorWhenNotFound(t *testing.T) {
	input := "data with env::MY_KEY_NOT_FOUND to replace"
	_, err := replaceEnvVar(input)
	assert.NotNil(t, err)
}
