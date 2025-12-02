package main

import (
	"bufio"
	"log"
	"os"
)

func readDataFile(pathName string) []string {

	file, err := os.Open(pathName)

	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	filescanner := bufio.NewScanner(file)

	var fileLines []string

	for filescanner.Scan() {
		fileLines = append(fileLines, filescanner.Text())
	}
	if err := filescanner.Err(); err != nil {
		log.Fatal(err)
	}

	return fileLines

}
