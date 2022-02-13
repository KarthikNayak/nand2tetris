package asm

import (
	"fmt"
	"translator/internal/asm/math"
	"translator/internal/asm/mem"
	"translator/internal/lexer"
)

type ASMer interface {
	Convert() ([]string, error)
}

type Data struct {
	CMD      lexer.CMD
	Stack    lexer.Stack
	Filename string
	Val      int
}

func Factory(d Data) (ASMer, error) {
	switch d.CMD {
	case lexer.POP:
		fallthrough
	case lexer.PUSH:
		switch d.Stack {
		case lexer.Static:
			return mem.NewStatic(d.CMD, d.Filename, d.Val), nil
		case lexer.Constant:
			return mem.NewConstant(d.CMD, d.Val), nil
		case lexer.Pointer:
			return mem.NewPointer(d.CMD, d.Filename, d.Val), nil
		case lexer.Temp:
			return mem.NewTemp(d.CMD, d.Val), nil
		case lexer.Local:
			return mem.NewLocation(d.CMD, "LCL", d.Val), nil
		case lexer.That:
			return mem.NewLocation(d.CMD, "THAT", d.Val), nil
		case lexer.This:
			return mem.NewLocation(d.CMD, "THIS", d.Val), nil
		case lexer.Argument:
			return mem.NewLocation(d.CMD, "ARG", d.Val), nil
		default:
			return nil, fmt.Errorf("not known stack type: %s", d.Stack)
		}
	case lexer.ADD:
		return math.NewAdd(), nil
	case lexer.SUB:
		return math.NewSubtract(), nil
	}
	return nil, fmt.Errorf("non supported cmd: %s", d.CMD)
}
