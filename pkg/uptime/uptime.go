package uptime

import (
	"fmt"
	"io/ioutil"
	"time"
	"gopkg.in/yaml.v2"
)

type Uptime struct {
	Services []*Service
}

func New(path string) (*Uptime, error) {
	u := new(Uptime)
	f, err := ioutil.ReadFile(path)
	if err != nil {
		return u, err
	}
	if err := yaml.Unmarshal(f, u); err != nil {
		return u, err
	}
	// validation
	for _, s := range u.Services {
		if !validType(s.Type) {
			return u, fmt.Errorf("%s is of invalid type: %s", s.Name, s.Type)
		}
		addDefaults(s)
	}
	return u, nil
}

func validType(t string) bool {
	switch t {
	case "http", "mongodb", "redis":
		return true
	}
	return false
}

func addDefaults(s *Service) {
	if s.Type == "http" && s.Expect == 0 {
		s.Expect = 200
	}
}

func output(c chan string, n int) {
	fmt.Printf("UPTIME CHECK %s\n", time.Now().Format(time.RFC3339))
	for i := 0; i < n; i++ {
		fmt.Println(<-c)
	}
}

func (u *Uptime) Run() {
	n := len(u.Services)
	c := make(chan string, n)
	for _, s := range u.Services {
		go s.run(c)
	}
	output(c, n)
}
