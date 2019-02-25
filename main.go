package main

import (
	"os"
	"github.com/karlpokus/uptime/pkg/uptime"
)

func main() {
	uptime, err := uptime.New(os.Args[1])
	if err != nil {
		panic(err)
	}
	uptime.Run()
}
