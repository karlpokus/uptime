package uptime

import (
	"net/http"
	"context"
	"time"
	"fmt"
	"strings"
)

var httpClient = &http.Client{
	//Timeout: 3 * time.Second,
	Transport: &http.Transport{},
}

func httpCall(s *Service) error {
	req, err := http.NewRequest(s.Method, s.Url, nil)
	if err != nil {
		return err
	}
	ctx, cancel := context.WithTimeout(context.Background(), 3 * time.Second)
	defer cancel()
	if s.Auth != "" {
		auth := strings.Split(s.Auth, ":")
		req.SetBasicAuth(auth[0], auth[1])
	}
	status, err := httpDo(req.WithContext(ctx))
	if err != nil {
		return err
	}
	if status != s.Expect {
		return fmt.Errorf("%d", status)
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
