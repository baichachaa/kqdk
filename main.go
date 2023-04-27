package main

import (
	"kqdk/app"
	"log"
	"os"
	"path/filepath"
)

func chdir() {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		log.Panicln(err)
	}

	err = os.Chdir(dir)
	if err != nil {
		log.Panicln(err)
	}
}

var debug string

func main() {
	if debug != "true" {
		chdir()
	}
	app.Run()

	app.SystemService()
}
