package main

import (
	"log"
	"io/ioutil"
	"os"
	"errors"
	"gopkg.in/yaml.v2"
	"github.com/karlpokus/uptime/service"
)

var (
	stdout = log.New(os.Stdout, "uptime ", log.Ldate | log.Ltime)
	stderr = log.New(os.Stderr, "uptime ", log.Ldate | log.Ltime)
)

type Conf struct {
	Services []service.Service
}

func (c *Conf) serviceTypesAreValid() error {
	isValid := func(str string) bool {
		switch str {
		case "http", "mongodb":
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

func main() {
	if len(os.Args) < 2 {
		stderr.Fatal(errors.New("Missing config path arg"))
	}
	conf := new(Conf)
	if err := conf.ReadFile(os.Args[1]); err != nil {
		stderr.Fatal(err)
	}
	c, n := runChecks(conf)
	for i := 0; i < n; i++ {
		stdout.Println(<-c)
	}
}
