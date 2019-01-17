package service

import (
	"time"
	"fmt"
)

type Service struct {
	Name, Type, Url, Method string
	Expect int
}

type Checker func(Service) error

var checkers = make(map[string]Checker)

func (s Service) Check() string {
	start := time.Now()
	err := checkers[s.Type](s)
	if err != nil {
		return fmt.Sprintf("%s failed %s", s.Name, err.Error())
	}
	end := time.Now()
	elapsed := end.Sub(start)
	return fmt.Sprintf("%s ok %v", s.Name, elapsed)
}

func init() {
	checkers["http"] = httpCall
	checkers["mongodb"] = mongoCall
}
