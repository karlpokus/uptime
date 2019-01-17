package service

import (
	"time"
	"fmt"
	"sync"
)

type Service struct {
	Name, Type, Url, Method string
	Expect int
}

type Checker func(Service) error

var checkers = make(map[string]Checker)

func (s Service) Check(wg *sync.WaitGroup, c chan string) {
	defer wg.Done()
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
