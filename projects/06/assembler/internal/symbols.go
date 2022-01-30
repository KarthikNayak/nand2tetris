package internal

import (
	"fmt"
	"strconv"
)

var Symbols map[string]int = map[string]int{
	"R0":     0,
	"R1":     1,
	"R2":     2,
	"R3":     3,
	"R4":     4,
	"R5":     5,
	"R6":     6,
	"R7":     7,
	"R8":     8,
	"R9":     9,
	"R10":    10,
	"R11":    11,
	"R12":    12,
	"R13":    13,
	"R14":    14,
	"R15":    15,
	"SCREEN": 16384,
	"KBD":    24576,
	"SP":     0,
	"LCL":    1,
	"ARG":    2,
	"THIS":   3,
	"THAT":   4,
}

func replaceSymbols(s []string) []string {
	var newS []string

	count := 0
	for _, val := range s {
		if val[0] == '(' {
			Symbols[val[1:len(val)-1]] = count
		} else {
			count += 1
			newS = append(newS, val)
		}
	}

	var finalS []string

	cur := 16
	for _, val := range newS {
		_, err := strconv.Atoi(val[1:])
		if val[0] == '@' && err != nil {
			if i, ok := Symbols[val[1:]]; ok {
				finalS = append(finalS, fmt.Sprintf("@%d", i))
			} else {
				Symbols[val[1:]] = cur
				finalS = append(finalS, fmt.Sprintf("@%d", cur))
				cur += 1
			}
		} else {
			finalS = append(finalS, val)
		}
	}

	return finalS
}
