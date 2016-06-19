package main

import (
    "errors"
)

type SymbolTable struct {
    Symbols map[string]uint16
}

func NewSymbolTable() (*SymbolTable, error) {
    table := &SymbolTable {
        Symbols: map[string]uint16 {
            "SP": 0,
            "LCL": 1,
            "ARG": 2,
            "THIS": 3,
            "THAT": 4,
            "R0": 0,
            "R1": 1,
            "R2": 2,
            "R3": 3,
            "R4": 4,
            "R5": 5,
            "R6": 6,
            "R7": 7,
            "R8": 8,
            "R9": 9,
            "R10": 10,
            "R11": 11,
            "R12": 12,
            "R13": 13,
            "R14": 14,
            "R15": 15,
            "SCREEN": 0x4000,
            "KBD": 0x6000,
        },
    }
    return table, nil
}

func (table *SymbolTable) Contains(symbol string) (ok bool) {
    _, ok = table.Symbols[symbol]
    return
}

func (table *SymbolTable) Get_address(symbol string) (uint16, error) {
    if table.Contains(symbol) {
        return table.Symbols[symbol], nil
    } else {
        return 0, errors.New("No such symbol")
    }
}

func (table *SymbolTable) Add_Entry(symbol string, address uint16) (error) {
    if table.Contains(symbol) {
        return errors.New("Already have such symbol")
    }
    table.Symbols[symbol] = address
    return nil
}
