package service

import (
	"net/http"
	"time"
	"errors"
	"strings"
)

var httpClient = &http.Client{
	Timeout: 3 * time.Second,
	Transport: &http.Transport{},
}

func httpCall(s Service) error {
	req, err := http.NewRequest(s.Method, s.Url, nil)
	if err != nil {
		return err
	}
	if s.Auth != "" {
		auth := strings.Split(s.Auth, ":")
		req.SetBasicAuth(auth[0], auth[1])
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
