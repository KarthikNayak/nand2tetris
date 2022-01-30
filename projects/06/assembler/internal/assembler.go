package internal

import (
	"fmt"
	"strconv"
	"strings"
)

var jmpSymbols = map[string]string{
	"JGT": "001",
	"JEQ": "010",
	"JGE": "011",
	"JLT": "100",
	"JNE": "101",
	"JLE": "110",
	"JMP": "111",
}

var cmpSymbols = map[string]string{
	"0":   "101010",
	"1":   "111111",
	"-1":  "111010",
	"D":   "001100",
	"A":   "110000",
	"!D":  "001101",
	"!A":  "110001",
	"-D":  "001111",
	"-A":  "110011",
	"D+1": "011111",
	"A+1": "110111",
	"D-1": "001110",
	"A-1": "110010",
	"D+A": "000010",
	"D-A": "010011",
	"A-D": "000111",
	"D&A": "000000",
	"D|A": "010101",
}

func assembler(s []string) []string {
	s = replaceSymbols(s)

	var machineCode []string

	for _, val := range s {
		if val[0] == '@' {
			i, _ := strconv.Atoi(val[1:])
			opCode := fmt.Sprintf("0%015b", i)
			machineCode = append(machineCode, opCode)
		} else {
			opCode := "111"
			jmpS, compS, destS := "", "", ""

			containsJmp := strings.Contains(val, ";")
			if containsJmp {
				split := strings.Split(val, ";")
				val = split[0]
				jmpS = split[1]
			}

			containsDest := strings.Contains(val, "=")
			if containsDest {
				split := strings.Split(val, "=")
				destS = split[0]
				compS = split[1]
			} else {
				compS = val
			}

			a := "0"
			if strings.Contains(compS, "M") {
				a = "1"
			}

			// comp
			comp := cmpSymbols[strings.ReplaceAll(compS, "M", "A")]

			// dest
			dest := "000"
			if containsDest {
				d := ""
				if strings.Contains(destS, "A") {
					d = d + "1"
				} else {
					d = d + "0"
				}
				if strings.Contains(destS, "D") {
					d = d + "1"
				} else {
					d = d + "0"
				}
				if strings.Contains(destS, "M") {
					d = d + "1"
				} else {
					d = d + "0"
				}
				dest = d
			}

			// jmp
			jmp := "000"
			if containsJmp {
				jmp = jmpSymbols[jmpS]
			}

			opCode = opCode + a + comp + dest + jmp
			machineCode = append(machineCode, opCode)
		}
	}

	return machineCode
}
