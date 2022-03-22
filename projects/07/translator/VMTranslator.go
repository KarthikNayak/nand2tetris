package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"translator/internal"
)

func main() {
	if len(os.Args) != 2 {
		handleError("Usage: ./translator <file_to_translate>", nil)
	}

	programName := os.Args[1]
	f, err := os.Open(programName)
	if err != nil {
		handleError("error opening VM file:", err)
	}

	splits := strings.Split(programName, "/")
	filename := strings.Split(splits[len(splits)-1], ".")[0]
	m, err := internal.Convert(f, filename)
	if err != nil {
		handleError("error while converting VM code", err)
	}

	file, err := os.Create(fmt.Sprintf("./%s.asm", filename))
	if err != nil {
		log.Fatal(err)
	}

	writer := bufio.NewWriter(file)
	for _, line := range m {
		_, err := writer.WriteString(line + "\n")
		if err != nil {
			log.Fatalf("Got error while writing to a file. Err: %s", err.Error())
		}
	}
	writer.Flush()

}

func handleError(s string, err error) {
	fmt.Printf("%s: %s\n", s, err)
	os.Exit(-1)
}
