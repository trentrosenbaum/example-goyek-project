package main

import (
	"github.com/goyek/goyek/v2"
	"github.com/goyek/x/cmd"
)

var fmt = goyek.Define(goyek.Task{
	Name:  "fmt",
	Usage: "go fmt",
	Action: func(a *goyek.A) {
		cmd.Exec(a, `go fmt ./...`)
	},
})
