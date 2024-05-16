package client

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"regexp"
	"strconv"
	"strings"

	"github.com/Julien4218/http-load-tester/config"
	"github.com/itchyny/gojq"
	log "github.com/sirupsen/logrus"
)

func Execute(test *config.HttpTest) bool {
	var err error
	body := test.Body
	body, err = replaceEnvVar(body)
	if err != nil {
		log.Errorf("error while setting body, detail:%s", err)
		return false
	}
	if test.SingleLineBody {
		body = strings.ReplaceAll(body, "\n", "")
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

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Errorf("error while processing http response with ReadAll(), detail:%s", err)
		return false
	}

	if !isResponseSuccessful(resp, test) {
		log.Errorf("http response without success code matching, code:%d response:`%s`", resp.StatusCode, string(respBody))
		return false
	}

	if len(test.SuccessJqQuery) > 0 {
		if isResponseJqValid(respBody, test) {
			return true
		}
		log.Errorf("http response without success JQ query, query:`%s` response:`%s`", test.SuccessJqQuery, string(respBody))
	}

	return false
}

func isResponseSuccessful(resp *http.Response, test *config.HttpTest) bool {
	for _, status := range test.SuccessResponseCodes {
		if resp.StatusCode == status {
			return true
		}
	}
	return false
}

func isResponseJqValid(respBody []byte, test *config.HttpTest) bool {
	query, err := gojq.Parse(test.SuccessJqQuery)
	if err != nil {
		log.Error(err)
		return false
	}

	var obj interface{}
	err = json.Unmarshal(respBody, &obj)
	if err != nil {
		log.Errorf("cannot unmarshal response, detail:%s", err)
		return false
	}

	iter := query.Run(obj)
	v, ok := iter.Next()
	if !ok {
		log.Errorf("cannot Next() result, detail:%s", err)
		return false
	}

	if err, ok := v.(error); ok {
		log.Errorf("jq evaluation resulted in error, detail:%s", err)
		return false
	}

	s, err := json.Marshal(v)
	if err != nil {
		log.Errorf("cannot marshal jq response value, detail:%s", err)
		return false
	}

	valid, err := strconv.ParseBool(string(s))
	if err != nil {
		log.Errorf("cannot evaluate response `%s` as a boolean with query `%s`, detail:%s", string(s), test.SuccessJqQuery, err)
		return false
	}
	return valid
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
