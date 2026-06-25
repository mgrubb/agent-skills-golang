package main

import (
	"github.com/goyek/goyek/v3"
	"github.com/goyek/x/cmd"
)

var format = goyek.Define(goyek.Task{
	Name:  "format",
	Usage: "format Go code",
	Action: func(a *goyek.A) {
		cmd.Exec(a, "go fmt ./...")
	},
})

var lint = goyek.Define(goyek.Task{
	Name:  "lint",
	Usage: "run golangci-lint",
	Action: func(a *goyek.A) {
		cmd.Exec(a, "golangci-lint run ./...")
	},
})

var test = goyek.Define(goyek.Task{
	Name:  "test",
	Usage: "run tests and write artifacts to dist/test",
	Action: func(a *goyek.A) {
		mustMkdir("dist/test")
		cmd.Exec(a, "go test -race -coverprofile=dist/test/coverage.out ./...")
		cmd.Exec(a, "go tool cover -html=dist/test/coverage.out -o dist/test/coverage.html")
	},
})

var check = goyek.Define(goyek.Task{
	Name:  "check",
	Usage: "run format, lint, and tests",
	Deps:  goyek.Deps{format, lint, test},
})
