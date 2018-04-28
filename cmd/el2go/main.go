package main

import (
	"os"
	"github.com/z-ray/log"
	"github.com/stephenlyu/goformula/easylang"
	"path/filepath"
	"strings"
	"flag"
)

func Usage() {
	log.Println(`Usage: compile [-package package-path] input-files`)
	os.Exit(1)
}

func compile(inputFile string, packagePath string) {
	parts := strings.Split(filepath.Base(inputFile), ".")

	if len(parts) > 1 {
		parts = parts[:len(parts) - 1]
	}
	if packagePath == "" {
		packagePath = strings.ToLower(parts[len(parts) - 1])
	}
	parts = append(parts, "go")
	outputFile := filepath.Join(strings.ToLower(strings.Join(parts, ".")))

	err := easylang.Compile2Go(inputFile, outputFile, nil, true, packagePath)
	if err != nil {
		log.Errorf("compile %s fail, error: %v", inputFile, err)
	}
}

func main() {
	packagePathPtr := flag.String("package", "", "Package full path")
	flag.Parse()

	if len(flag.Args()) <= 0 {
		Usage()
	}

	for _, file := range flag.Args() {
		compile(file, *packagePathPtr)
	}
}
