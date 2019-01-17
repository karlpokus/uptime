package service

import (
	"net/http"
	"time"
	"errors"
)

var httpClient = &http.Client{
	Timeout: 3 * time.Second, // try this before adding context
	Transport: &http.Transport{},
}

func httpCall(s Service) error {
	req, err := http.NewRequest(s.Method, s.Url, nil)
	if err != nil {
		return err
	}
	status, err := httpDo(req)
	if err != nil {
		return err
	}
	if status != s.Expect {
		return errors.New(s.Name + " failed")
	}
	return nil
}

func httpDo(req *http.Request) (int, error) {
	res, err := httpClient.Do(req)
	if err != nil {
		return 0, err
	}
	defer res.Body.Close()
	return res.StatusCode, nil
}
