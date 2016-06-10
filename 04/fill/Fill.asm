// This file is part of www.nand2tetris.org
// and the book "The Elements of Computing Systems"
// by Nisan and Schocken, MIT Press.
// File name: projects/04/Fill.asm

// Runs an infinite loop that listens to the keyboard input.
// When a key is pressed (any key), the program blackens the screen,
// i.e. writes "black" in every pixel. When no key is pressed, the
// program clears the screen, i.e. writes "white" in every pixel.

// Put your code here.
    @status
    M=0
    @SETSCREEN                   //initial screen to white
    0;JMP

(LOOP)
    @new_status                   //new_status  = 0xffff
    M=-1
    @KBD
    D=M                           //D=M[KBD]
    @CLEAN
    D;JEQ                         //if no key pressed, set new_status = 0x0000
    @COMPARE
    0;JMP

(CLEAN)
    @new_status
    M=0

(COMPARE)                          //new_status == status, goto LOOP
    @status
    D=M
    @new_status
    D=D-M
    @LOOP
    D;JEQ
    @new_status                    //new_status != status, status = new_status
    D=M
    @status
    M=D

(SETSCREEN)
    @SCREEN
    D=A
    @8192
    D=D+A
    @i
    M=D

(SETLOOP)
    @i
    D=M-1
    M=D                            //i=i-1
    @SCREEN
    D=D-A
    @LOOP
    D;JLT
    @status
    D=M
    @i
    A=M
    M=D
    @SETLOOP
    0;JMP
