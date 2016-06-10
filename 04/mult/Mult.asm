// This file is part of www.nand2tetris.org
// and the book "The Elements of Computing Systems"
// by Nisan and Schocken, MIT Press.
// File name: projects/04/Mult.asm

// Multiplies R0 and R1 and stores the result in R2.
// (R0, R1, R2 refer to RAM[0], RAM[1], and RAM[2], respectively.)

// Put your code here.
    @2
    M=0      //R2=0
(LOOP)
    @1
    D=M      //D=R1
    @0
    D=D-A    //D=R1-0
    @END
    D;JLE    //if (R1-0) <= 0 goto END
    @0
    D=M      //D=R0
    @2
    M=M+D    //R2=R2+R0
    @1
    M=M-1    //R1=R1-1
    @LOOP
    0;JMP
(END)
    @END
    0;JMP
