package client

import (
	"fmt"
	"net/http"
	"os"
	"regexp"
	"strings"

	"github.com/Julien4218/http-load-tester/config"
	log "github.com/sirupsen/logrus"
)

func Execute(test config.HttpTest) bool {
	var err error
	body := test.Body
	body, err = replaceEnvVar(body)
	if err != nil {
		log.Errorf("error while setting body, detail:%s", err)
		return false
	}
	req, err := http.NewRequest("POST", test.URL, strings.NewReader(body))
	if err != nil {
		log.Errorf("error while creating request with url:%s, detail:%s", test.URL, err)
		return false
	}
	for k, v := range test.Headers {
		v, err = replaceEnvVar(v)
		if err != nil {
			log.Errorf("error while setting headers, detail:%s", err)
			return false
		}
		req.Header.Set(k, v)
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Errorf("error while executing request with url:%s, detail:%s", test.URL, err)
		return false
	}
	defer resp.Body.Close()

	for _, status := range test.SuccessResponseCodes {
		if resp.Status == fmt.Sprintf("%d", status) {
			return true
		}
	}

	return false
}

func replaceEnvVar(value string) (string, error) {
	if len(value) == 0 {
		return "", nil
	}
	result := value
	re := regexp.MustCompile(`env::(\w+)`)
	for found := re.Find([]byte(result)); found != nil; found = re.Find([]byte(result)) {
		match := string(found)
		key := strings.ReplaceAll(match, "env::", "")
		replaced := os.Getenv(key)
		if replaced == "" {
			return "", fmt.Errorf("environment variable %s is defined but not found, cannot replace", match)
		}
		result = strings.ReplaceAll(result, match, replaced)
	}
	return result, nil
}
