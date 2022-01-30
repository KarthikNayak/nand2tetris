package internal

import (
	"bufio"
	"io"
	"strings"
)

type Source interface {
	io.Reader
}

func Convert(s Source) []string {
	var lines []string

	scanner := bufio.NewScanner(s)

	for scanner.Scan() {
		s := scanner.Text()

		if skipLine(s) {
			continue
		}

		s = cleanLine(s)

		lines = append(lines, s)
	}

	return assembler(lines)
}

func skipLine(s string) bool {
	s = strings.TrimSpace(s)

	if len(s) == 0 {
		return true
	}

	if s[0] == '/' && s[1] == '/' {
		return true
	}

	return false
}

func cleanLine(s string) string {
	split := strings.Split(s, "//")
	return strings.TrimSpace(split[0])
}
