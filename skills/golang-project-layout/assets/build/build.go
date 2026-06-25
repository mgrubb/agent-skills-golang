package main

import (
	"os"

	"github.com/goyek/goyek/v3"
	"github.com/goyek/x/cmd"
)

var clean = goyek.Define(goyek.Task{
	Name:  "clean",
	Usage: "remove generated artifacts",
	Action: func(a *goyek.A) {
		if err := os.RemoveAll("dist"); err != nil {
			a.Fatal(err)
		}
	},
})

var build = goyek.Define(goyek.Task{
	Name:  "build",
	Usage: "build all commands into dist/bin",
	Action: func(a *goyek.A) {
		mustMkdir("dist/bin")
		cmd.Exec(a, "go build -o dist/bin/ ./cmd/...")
	},
})

var all = goyek.Define(goyek.Task{
	Name:  "all",
	Usage: "run the standard local pipeline",
	Deps:  goyek.Deps{check, build},
})
