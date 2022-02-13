package internal

import (
	"fmt"
	"translator/internal/asm"
	"translator/internal/lexer"
)

func translate(input []string, filename string) ([]string, error) {
	final := []string{}

	for i, line := range input {
		cmd, stack, value, err := lexer.Lex(line)
		if err != nil {
			return final, fmt.Errorf("couldn't do lexical analysis: %w", err)
		}

		asmer, err := asm.Factory(asm.Data{
			CMD:      cmd,
			Stack:    stack,
			Filename: filename,
			Val:      value,
			Index:    i,
		})
		if err != nil {
			return final, fmt.Errorf("couldn't get the asm factory: %w", err)
		}

		asm, err := asmer.Convert()
		if err != nil {
			return final, fmt.Errorf("couldn't produce asm code: %w", err)
		}
		final = append(final, asm...)
	}

	return final, nil
}
