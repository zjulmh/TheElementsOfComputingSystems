package main

import (
    "os"
    "fmt"
)

func main() {
    if len(os.Args) != 2 {
        fmt.Fprintln(os.Stderr, "Usage: Assembler file.asm")
        return
    }

    asm, err := NewAssembler(os.Args[1])
    if err != nil {
        fmt.Fprintln(os.Stderr, err.Error())
    }
    defer asm.Close()

    err = asm.Assemble()
    if err != nil {
        fmt.Fprintln(os.Stderr, err.Error())
        return
    }
    return
}
