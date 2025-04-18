package assembler

import (
	"06_assembler/code"
	"06_assembler/parser"
	"fmt"
	"strconv"
	"strings"
)

var symbolTable = map[string]int{
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
	"SP":     0,
	"LCL":    1,
	"ARG":    2,
	"THIS":   3,
	"THAT":   4,
	"SCREEN": 16384,
	"KBD":    24576,
}
var varAddress = 16

func Compile(asm string) string {
	p := parser.NewParser(asm)
	firstPass(p)
	p.Reset()
	return secondPass(p)
}

func firstPass(p *parser.Parser) {
	foundLs := 0
	for p.HasMoreLines() {
		p.Advance()
		t := p.GetInstructionType()
		line := p.GetCurrentLine()
		if t == parser.L_INSTR {
			symbol := p.Symbol()
			symbolTable[symbol] = line - foundLs
			foundLs++
		}
	}
}

func secondPass(p *parser.Parser) string {
	var hackCode strings.Builder
	for p.HasMoreLines() {
		p.Advance()
		t := p.GetInstructionType()
		var bin string
		if t == parser.C_INSTR {
			bin = getCInstruction(p.Comp(), p.Dest(), p.Jump())
		} else if t == parser.A_INSTR {
			bin = getAInstruction(p.Symbol())
		} else {
			continue
		}
		hackCode.WriteString(bin + "\n")
	}
	return hackCode.String()
}

func getCInstruction(comp, dest, jump string) string {
	return "111" + code.Comp(comp) + code.Dest(dest) + code.Jump(jump)
}

func getAInstruction(symbol string) string {
	num, err := strconv.Atoi(symbol)
	if err != nil {
		address, found := symbolTable[symbol]
		if !found {
			address = varAddress
			symbolTable[symbol] = address
			varAddress++
		}
		return fmt.Sprintf("0%015b", address)
	}
	return fmt.Sprintf("0%015b", num)
}
