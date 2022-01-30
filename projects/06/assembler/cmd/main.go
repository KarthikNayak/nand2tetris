package main

import (
	"assembler/internal"
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	if len(os.Args) != 2 {
		handleError("Usage: ./assembler <file_to_translate>", nil)
	}

	programName := os.Args[1]
	f, err := os.Open(programName)
	if err != nil {
		handleError("error opening assembly file:", err)
	}

	m := internal.Convert(f)

	filename := strings.Split(programName, ".")[0]
	file, err := os.Create(fmt.Sprintf("./%s.hack", filename))
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
