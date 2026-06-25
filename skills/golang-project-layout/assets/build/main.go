package main

import (
	"errors"
	"os"

	"github.com/goyek/goyek/v3"
	"github.com/goyek/x/boot"
)

func main() {
	if err := os.Chdir(".."); err != nil {
		panic(err)
	}

	mustMkdir("dist/bin")
	mustMkdir("dist/test")

	goyek.SetDefault(all)
	boot.Main()
}

func mustMkdir(path string) {
	err := os.MkdirAll(path, 0o755)
	if err != nil && !errors.Is(err, os.ErrExist) {
		panic(err)
	}
}
