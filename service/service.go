package service

import (
	"time"
	"fmt"
)

type Service struct {
	Name, Type, Url, Method, Auth string
	Expect int
}

type Checker func(Service) error

var checkers = make(map[string]Checker)

func (s Service) Check(c chan string) {
	start := time.Now()
	err := checkers[s.Type](s)
	if err != nil {
		c <- fmt.Sprintf("%s failed %s", s.Name, err.Error())
		return
	}
	end := time.Now()
	elapsed := end.Sub(start)
	c <- fmt.Sprintf("%s ok %v", s.Name, elapsed)
}

func init() {
	checkers["http"] = httpCall
	checkers["mongodb"] = mongoCall
}
