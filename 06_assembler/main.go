package main

import (
	"06_assembler/assembler"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func read_asm(filename string) string {
	b, err := os.ReadFile(filename)
	if err != nil {
		fmt.Println("Error reading", filename)
		os.Exit(1)
	}
	return string(b)
}

func write_hack(name, output_dir, hackCode string) {
	f, err := os.Create(filepath.Join(output_dir, name+".hack"))
	if err != nil {
		fmt.Println("Error creating output file")
		os.Exit(1)
	}
	f.WriteString(hackCode)
	f.Close()
}

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: ./assember <program.asm>")
		os.Exit(1)
	}

	filename := os.Args[1]
	file := filepath.Base(filename)
	name := strings.Split(file, ".")[0]

	asm := read_asm(filename)
	hackCode := assembler.Compile(asm)

	output_dir := filepath.Dir(filename)
	write_hack(name, output_dir, hackCode)
}
