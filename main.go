package main

import (
	"kqdk/app"
	"log"
	"os"
	"path/filepath"
	"strconv"
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

	isDebug, _ := strconv.ParseBool(debug)

	// 修改运行目录
	if isDebug != true {
		chdir()
	}

	app.Run(isDebug)
}
