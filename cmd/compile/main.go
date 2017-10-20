package main

import (
	"os"
	"github.com/z-ray/log"
	"github.com/stephenlyu/goformula/easylang"
	"path/filepath"
	"strings"
)

func Usage() {
	log.Println(`Usage: compile input-files`)
	os.Exit(1)
}

func compile(inputFile string) {
	dir := filepath.Dir(inputFile)
	parts := strings.Split(filepath.Base(inputFile), ".")

	if len(parts) > 1 {
		parts = parts[:len(parts) - 1]
	}
	parts = append(parts, "lua")

	outputFile := filepath.Join(dir, strings.ToLower(strings.Join(parts, ".")))

	err := easylang.Compile(inputFile, outputFile)
	if err != nil {
		log.Errorf("compile %s fail, error: %v", inputFile, err)
	}
}

func main() {
	if len(os.Args) < 2 {
		Usage()
	}

	for _, file := range os.Args[1:] {
		compile(file)
	}
}
