package main

import (
	"fmt"
	"os"
	"io/ioutil"
	"errors"
	"time"
	"gopkg.in/yaml.v2"
	"github.com/karlpokus/uptime/service"
)

type Conf struct {
	Services []service.Service
}

func (c *Conf) serviceTypesAreValid() error {
	isValid := func(str string) bool {
		switch str {
		case "http", "mongodb", "redis":
			return true
		}
		return false
	}
	for _, service := range c.Services {
		if !isValid(service.Type) {
			return errors.New(service.Type + " is invalid")
		}
	}
	return nil
}

func (c *Conf) ReadFile(path string) error {
	f, err := ioutil.ReadFile(path)
	if err != nil {
		return err
	}
	if err := yaml.Unmarshal(f, c); err != nil {
		return err
	}
	if err := c.serviceTypesAreValid(); err != nil {
		return err
	}
	return nil
}

func runChecks(conf *Conf) (chan string, int) {
	sCount := len(conf.Services)
	c := make(chan string, sCount)
	for _, s := range conf.Services {
		go s.Check(c)
	}
	return c, sCount
}

func output(c chan string, n int) {
	fmt.Printf("\nUPTIME CHECK \t %s\n\n", time.Now().Format(time.RFC3339))
	for i := 0; i < n; i++ {
		fmt.Println(<-c)
	}
	fmt.Println()
}

func main() {
	if len(os.Args) < 2 {
		panic(errors.New("Missing config path arg"))
	}
	conf := new(Conf)
	if err := conf.ReadFile(os.Args[1]); err != nil {
		panic(err)
	}
	output(runChecks(conf))
}
