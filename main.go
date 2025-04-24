package main

import (
	"flag"
	"fmt"
	"os"
	"io"
)

func main() {
	fromPath := flag.String("from", "", "Path to the source file")
	toPath := flag.String("to", "", "Path to the destination file")

	flag.Parse()

	if *fromPath == "" || *toPath == "" {
		fmt.Println("Usage: mycp --from source.txt --to dest.txt")
		os.Exit(1)
	}

	fromFile, err := os.Open(*fromPath)
	if err != nil {
		fmt.Println("Error: could not open source file -", err)
		os.Exit(1)
	}
	defer fromFile.Close()

	toFile, err := os.Create(*toPath)
	if err != nil {
		fmt.Println("Error: could not create destination file -", err)
		os.Exit(1)
	}
	defer toFile.Close()

	_, err = io.Copy(toFile, fromFile)
	if err != nil {
		fmt.Println("Error: file copy failed -", err)
		os.Exit(1)
	}

	fmt.Println("Copy completed", *fromPath, "→" , *toPath)
}