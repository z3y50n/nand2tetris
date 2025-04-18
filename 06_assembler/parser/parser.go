package parser

import (
	"strings"
)

type Parser struct {
	lines       []string
	currentLine int
}

type InstructionType int

const (
	A_INSTR InstructionType = iota
	C_INSTR
	L_INSTR
)

func NewParser(code string) *Parser {
	var lines []string
	instructions := strings.Split(code, "\n")
	for _, instruction := range instructions {
		if !isWhitespaceOrComment(instruction) {
			lines = append(lines, cleanupInstruction(instruction))
		}
	}

	return &Parser{
		lines:       lines,
		currentLine: -1,
	}
}

func (p *Parser) HasMoreLines() bool {
	return p.currentLine < len(p.lines)-1
}

func (p *Parser) GetCurrentLine() int {
	return p.currentLine
}

func (p *Parser) Advance() {
	if p.HasMoreLines() {
		p.currentLine++
	}
}

func (p *Parser) Reset() {
	p.currentLine = -1
}

func (p *Parser) GetInstructionType() InstructionType {
	instruction := p.lines[p.currentLine]
	if strings.HasPrefix(instruction, "(") && strings.HasSuffix(instruction, ")") {
		return L_INSTR
	}
	if strings.HasPrefix(instruction, "@") {
		return A_INSTR
	}
	// We assume instructions are error free
	return C_INSTR
}

func (p *Parser) Symbol() string {
	instruction := p.lines[p.currentLine]
	instructionType := p.GetInstructionType()
	if instructionType == A_INSTR {
		return instruction[1:]
	}
	if instructionType == L_INSTR {
		return instruction[1 : len(instruction)-1]
	}
	return ""
}

func (p *Parser) Dest() string {
	instruction := p.lines[p.currentLine]
	if p.GetInstructionType() != C_INSTR {
		return ""
	}
	if strings.Contains(instruction, "=") {
		return strings.Split(instruction, "=")[0]
	}
	return "null"
}

func (p *Parser) Comp() string {
	instruction := p.lines[p.currentLine]
	if p.GetInstructionType() != C_INSTR {
		return ""
	}
	equalIndex := strings.Index(instruction, "=")
	semicolonIndex := strings.Index(instruction, ";")

	start := 0
	if equalIndex != -1 {
		start = equalIndex + 1
	}

	end := len(instruction)
	if semicolonIndex != -1 {
		end = semicolonIndex
	}
	return instruction[start:end]
}

func (p *Parser) Jump() string {
	instruction := p.lines[p.currentLine]
	if p.GetInstructionType() != C_INSTR {
		return ""
	}
	if strings.Contains(instruction, ";") {
		return strings.Split(instruction, ";")[1]
	}
	return "null"
}

func isWhitespaceOrComment(instruction string) bool {
	return cleanupInstruction(instruction) == ""
}

/*
Remove whitespace and trailing comments from instruction
*/
func cleanupInstruction(instruction string) string {
	return strings.Split(strings.Trim(instruction, " "), "//")[0]
}
