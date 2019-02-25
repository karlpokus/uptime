package uptime

import (
	"time"
	"fmt"
)

type Service struct {
	Name, Type, Url, Method, Auth, Pwd string
	Expect int
}

func (s *Service) run(c chan string) {
	var err error
	start := time.Now()
	switch s.Type {
	case "http":
		err = httpCall(s)
	case "mongodb":
		err = mongoCall(s)
	case "redis":
		err = redisCall(s)
	}
	if err != nil {
		c <- fmt.Sprintf("%s \t fail \t %s", s.Name, err.Error())
		return
	}
	end := time.Now()
	elapsed := end.Sub(start)
	c <- fmt.Sprintf("%s \t ok \t %v", s.Name, elapsed)
}
