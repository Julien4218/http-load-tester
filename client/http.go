package main

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/Julien4218/http-load-tester/config"
	log "github.com/sirupsen/logrus"
)

func Execute(test config.HttpTest) bool {
	req, err := http.NewRequest("POST", test.URL, strings.NewReader(test.Body))
	if err != nil {
		log.Errorf("error while creating request with url:%s, detail:%s", test.URL, err)
		return false
	}
	for k, v := range test.Headers {
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
