package config

import (
	"io/ioutil"
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestInit(t *testing.T) {
	fileContent := `
MinParallel: 1
RequestPerMinute: 60

`
	filepath, err := createTempYAMLFile(fileContent)
	if err != nil {
		t.Fatalf("Error creating temporary YAML file: %v", err)
	}
	defer os.Remove(filepath)

	config, err := Init(filepath)
	if err != nil {
		t.Errorf("Expected no error, but got: %v", err)
	}
	require.Equal(t, config.MinParallel, 1)
}

func TestShouldFailMissingFile(t *testing.T) {
	_, err := Init("")
	require.Error(t, err)
	require.Contains(t, err.Error(), "valid config filepath")
}

func TestShouldFailInvalidYaml(t *testing.T) {
	fileContent := `
				this is not a valid yaml file
		really
	not
`
	filepath, err := createTempYAMLFile(fileContent)
	if err != nil {
		t.Fatalf("Error creating temporary YAML file: %v", err)
	}
	defer os.Remove(filepath)

	_, err = Init(filepath)
	require.Error(t, err)
	require.Contains(t, err.Error(), "YAML")
}

func createTempYAMLFile(content string) (string, error) {
	tmpfile, err := ioutil.TempFile("", "test-config-*.yaml")
	if err != nil {
		return "", err
	}
	defer tmpfile.Close()

	_, err = tmpfile.WriteString(content)
	if err != nil {
		return "", err
	}

	return tmpfile.Name(), nil
}
