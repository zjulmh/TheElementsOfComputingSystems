package main

import (
    "fmt"
    "errors"
)

type Code struct {
}

func (code *Code) Gen_ACommand(address uint16) (command string, err error){
    command = "0" + fmt.Sprintf("%015b", address)
    return
}

var dest_code map[string]int = map[string]int {
    "": 0,
    "M": 1,
    "D": 2,
    "MD": 3,
    "A": 4,
    "AM": 5,
    "AD": 6,
    "AMD": 7,
}

func dest_cmd(dest string) (command string, err error) {
    value, ok := dest_code[dest]
    if !ok {
        err = errors.New("No such dest")
        return
    }
    command = fmt.Sprintf("%03b", value)
    return
}

var comp_codes map[string]string = map[string]string {
    "0":"0101010",
    "1":"0111111",
    "-1":"0111010",
    "D":"0001100",
    "A":"0110000",
    "!D":"0001101",
    "!A":"0110001",
    "-D":"0001111",
    "-A":"0110011",
    "D+1":"0011111",
    "A+1":"0110111",
    "D-1":"0001110",
    "A-1":"0110010",
    "D+A":"0000010",
    "D-A":"0010011",
    "A-D":"0000111",
    "D&A":"0000000",
    "D|A":"0010101",
    "M":"1110000",
    "!M":"1110001",
    "-M":"1110011",
    "M+1":"1110111",
    "M-1":"1110010",
    "D+M":"1000010",
    "D-M":"1010011",
    "M-D":"1000111",
    "D&M":"1000000",
    "D|M":"1010101",
}

func comp_cmd(comp string) (command string, err error) {
    var ok bool
    command, ok = comp_codes[comp]
    if !ok {
        err = errors.New("No such comp command")
        return
    }
    return
}

var jump_code map[string]int = map[string]int {
    "": 0,
    "JGT": 1,
    "JEQ": 2,
    "JGE": 3,
    "JLT": 4,
    "JNE": 5,
    "JLE": 6,
    "JMP": 7,
}

func jump_cmd(jump string) (command string, err error) {
    value, ok := jump_code[jump]
    if !ok {
        err = errors.New("No such jump command")
        return
    }
    command = fmt.Sprintf("%03b", value)
    return
}

func (code *Code) Gen_CCommand (dest, comp, jump string) (command string, err error) {
    var compcmd string
    var destcmd string
    var jumpcmd string
    compcmd, err = comp_cmd(comp)
    if err != nil {
        return
    }
    destcmd, err = dest_cmd(dest)
    if err != nil {
        return
    }
    jumpcmd, err = jump_cmd(jump)
    if err != nil {
        return
    }
    command = "111" + compcmd + destcmd + jumpcmd
    return
}
