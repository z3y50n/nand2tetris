// This file is part of www.nand2tetris.org
// and the book "The Elements of Computing Systems"
// by Nisan and Schocken, MIT Press.
// File name: projects/4/Fill.asm

// Runs an infinite loop that listens to the keyboard input. 
// When a key is pressed (any key), the program blackens the screen,
// i.e. writes "black" in every pixel. When no key is pressed, 
// the screen should be cleared.
	
	@status
	M=-1
	@newstatus
	M=0 
	@CHECKCHANGE
	0;JMP
(KBDLOOP)
	@KBD
	D=M
	@SETWHITE
	D;JEQ
	@SETBLACK
	0;JMP
(SETBLACK)
	@newstatus
	M=-1
	@CHECKCHANGE
	0;JMP
(SETWHITE)
	@newstatus
	M=0
	@CHECKCHANGE
	0;JMP
(CHECKCHANGE)
	@status
	D=M
	@newstatus
	D=D-M
	@KBDLOOP
	D;JEQ
(SETSCREEN)
	@newstatus
	D=M
	@status
	M=D

	@SCREEN
	D=A
	@8191
	D=D+A
	@position
	M=D
(PAINTLOOP)
	@position
	D=M
	@SCREEN
	D=D-A
	@KBDLOOP
	D;JLT

	@status
	D=M
	@position
	A=M
	M=D
	@position
	M=M-1
	@PAINTLOOP
	0;JMP
	

	

