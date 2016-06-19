package main

import (
    "fmt"
    "os"
    "strconv"
    "strings"
)

type Assembler struct {
    FileName string
    OutFile *os.File
    *SymbolTable
    SymbolAddress uint16
}

func NewAssembler(file string) (asm *Assembler, err error) {
    var symbol_table *SymbolTable
    symbol_table, err = NewSymbolTable()
    if err != nil {
        return
    }
    asm = &Assembler {
        FileName: file,
        SymbolTable: symbol_table,
        SymbolAddress: 16,
    }
    output_name := asm.outputfile()
    asm.OutFile, err = os.OpenFile(output_name, os.O_WRONLY|os.O_CREATE, os.ModePerm)
    if err != nil {
        return
    }
    return
}

func (asm *Assembler) pass_0() (error){
    parser, err := NewParser(asm.FileName)
    if err != nil {
        return err
    }
    defer parser.Close()
    var cur_address uint16 = 0
    for {
        yes, err := parser.HasMoreCommand()
        if err != nil {
            return err
        }
        if yes {
            parser.Advance()
            cmd := parser.CmdType
            if cmd == A_COMMAND || cmd == C_COMMAND {
                cur_address += 1
            } else if cmd == L_COMMAND {
                asm.SymbolTable.Add_Entry(parser.Symbol, cur_address)
            }
        } else {
            break
        }
    }
    return nil
}

func (asm *Assembler) pass_1() (error) {
    parser, err := NewParser(asm.FileName)
    if err != nil {
        return err
    }
    defer parser.Close()
    code := &Code{}
    for {
        yes, err := parser.HasMoreCommand()
        if err != nil {
            return err
        }
        if yes {
            parser.Advance()
            cmd := parser.CmdType
            if cmd == A_COMMAND {
                address, err := asm.get_address(parser.Symbol)
                if err != nil {
                    return err
                }
                out, err := code.Gen_ACommand(address)
                if err != nil {
                    return err
                }
                fmt.Fprintln(asm.OutFile, out)
            } else if cmd == C_COMMAND {
                fmt.Sprintln(os.Stdout, parser.Dest, parser.Comp, parser.Jump)
                out, err := code.Gen_CCommand(parser.Dest, parser.Comp, parser.Jump)
                if err != nil {
                    return err
                }
                fmt.Fprintln(asm.OutFile, out)
            }
        } else {
            break
        }
    }
    return nil
}

func (asm *Assembler) get_address(symbol string) (uint16, error) {
    address, err := strconv.ParseUint(symbol, 10, 16)
    if err == nil {
        return uint16(address), nil
    }
    if !asm.SymbolTable.Contains(symbol) {
        asm.SymbolTable.Add_Entry(symbol, asm.SymbolAddress)
        asm.SymbolAddress += 1
    }
    return asm.SymbolTable.Get_address(symbol)
}

func (asm *Assembler) Assemble() error {
    err := asm.pass_0()
    if err != nil {
        return err
    }
    err = asm.pass_1()
    return err
}

func (asm *Assembler) Close() {
    asm.OutFile.Close()
}

func (asm *Assembler) outputfile() string {
    if strings.HasSuffix(asm.FileName, ".asm") {
        return strings.Replace(asm.FileName, ".asm", ".hack", -1)
    } else {
        return asm.FileName + ".hack"
    }
}
