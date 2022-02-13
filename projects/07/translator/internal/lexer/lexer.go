package lexer

import (
	"fmt"
	"strconv"
	"strings"
)

type CMD string

var (
	POP  CMD = "POP"
	PUSH CMD = "PUSH"
	ADD  CMD = "ADD"
	SUB  CMD = "SUB"
	NEG  CMD = "NEG"
	EQ   CMD = "EQ"
	GT   CMD = "GT"
	LT   CMD = "LT"
	AND  CMD = "AND"
	OR   CMD = "OR"
	NOT  CMD = "NOT"
)

type Stack string

var (
	Local    Stack = "Local"
	Argument Stack = "Argument"
	This     Stack = "This"
	That     Stack = "That"
	Constant Stack = "Constant"
	Static   Stack = "Static"
	Temp     Stack = "Temp"
	Pointer  Stack = "Pointer"
)

func Lex(line string) (CMD, Stack, int, error) {
	s := strings.Split(line, " ")

	var cmd CMD
	var stack Stack
	var value int
	var err error

	cmd, err = parseCMD(s[0])
	if err != nil {
		return cmd, stack, value, fmt.Errorf("couldn't lex line: %w", err)
	}

	if len(s) == 3 {
		stack, err = parseStack(s[1])
		if err != nil {
			return cmd, stack, value, fmt.Errorf("couldn't lex line: %w", err)
		}

		value, err = strconv.Atoi(s[2])
		if err != nil {
			return cmd, stack, value, fmt.Errorf("couldn't lex line: %w", err)
		}
	}

	return cmd, stack, value, nil
}

func parseCMD(cmd string) (CMD, error) {
	switch cmd {
	case "pop":
		return POP, nil
	case "push":
		return PUSH, nil
	case "add":
		return ADD, nil
	case "sub":
		return SUB, nil
	case "neg":
		return NEG, nil
	case "eq":
		return EQ, nil
	case "gt":
		return GT, nil
	case "lt":
		return LT, nil
	case "and":
		return AND, nil
	case "or":
		return OR, nil
	case "not":
		return NOT, nil
	}

	return "", fmt.Errorf("unknown command type: %s", cmd)
}

func parseStack(stack string) (Stack, error) {
	switch stack {
	case "local":
		return Local, nil
	case "argument":
		return Argument, nil
	case "this":
		return This, nil
	case "that":
		return That, nil
	case "constant":
		return Constant, nil
	case "static":
		return Static, nil
	case "temp":
		return Temp, nil
	case "pointer":
		return Pointer, nil
	}

	return "", fmt.Errorf("unknown stack type: %s", stack)
}
