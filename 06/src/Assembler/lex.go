package main

import (
    "regexp"
    "bufio"
    "os"
    "errors"
)

const (
    NUM = 1
    ID = 2
    OP = 3
    ERROR = 4
)

const (
    num_re = "\\d+"
    id_start = "\\w_\\.$:"
    id_re = "[" + id_start + "]" + "[" + id_start + "\\d]*"
    op_re = "[+=;()@\\\\&|!-]"
    word_re = num_re + "|" + id_re + "|" + op_re
    comment_re = "//.*$"
    not_whitespace_re = "\\S"
)

var commentReg *regexp.Regexp = nil
var wordReg *regexp.Regexp = nil
var numReg *regexp.Regexp = nil
var idReg *regexp.Regexp = nil
var opReg *regexp.Regexp = nil
var notWhiteSpaceReg *regexp.Regexp = nil

func init() {
    commentReg = regexp.MustCompile(comment_re)
    wordReg = regexp.MustCompile(word_re)
    numReg = regexp.MustCompile(num_re)
    idReg = regexp.MustCompile(id_re)
    opReg = regexp.MustCompile(op_re)
    notWhiteSpaceReg = regexp.MustCompile(not_whitespace_re)
}

type Token struct {
    Type int
    Word string
}

type Lex struct {
    *os.File
    *bufio.Scanner
    CurrentLine string
    CurrentCommand []Token
    CurrentToken Token
}

func NewLex(file string) (lex *Lex, err error) {
    var f *os.File
    f, err = os.Open(file)
    if err != nil {
        return
    }

    scanner := bufio.NewScanner(f)
    lex = &Lex {
        File: f,
        Scanner: scanner,
        CurrentLine: "",
        CurrentCommand: make([]Token, 0),
        CurrentToken: Token {
            Type: ERROR,
            Word: "",
        },
    }
    return
}

func (lex *Lex) Close() {
    lex.File.Close()
}

func (lex *Lex) HasMoreCommand() (yes bool, err error) {
    var line string
    var line_no_comment string
    for lex.Scan() {
        line = lex.Text()
        err = lex.Err()
        if err != nil {
            return
        }
        line_no_comment = commentReg.ReplaceAllString(line, "")
        if !notWhiteSpaceReg.MatchString(line_no_comment) {
            continue
        }
        lex.CurrentLine = line_no_comment
        yes = true
        return
    }
    return
}

func (lex *Lex) NextCommand() (tokens []Token, err error) {
     lex.CurrentCommand, err = tokenize(lex.CurrentLine)
     if err != nil {
        return
     }
     lex.NextToken()
     tokens = lex.CurrentCommand
     return
}

func (lex *Lex) HasNextToken() bool {
    return len(lex.CurrentCommand) > 0
}

func (lex *Lex) NextToken() (token Token) {
    if lex.HasNextToken() {
        lex.CurrentToken = lex.CurrentCommand[0]
        lex.CurrentCommand = lex.CurrentCommand[1:]
        token = lex.CurrentToken
    } else {
        token.Type = ERROR
        token.Word = ""
    }
    return
}

func (lex *Lex) PeekToken() (token Token) {
    if lex.HasNextToken() {
        token = lex.CurrentCommand[0]
    } else {
        token.Type = ERROR
        token.Word = ""
    }
    return
}

func tokenize(line string) (tokens []Token, err error) {
    tokens = make([]Token, 0)
    words := wordReg.FindAllString(line, -1)
    if len(words) == 0 {
        err = errors.New("malformed command")
    }
    for _, word := range words {
        token := token_word(word)
        tokens = append(tokens, token)
    }
    return
}

func token_word(word string) (token Token) {
    token.Type = ERROR
    token.Word = word
    if is_num(word) {
        token.Type = NUM
        return
    } else if is_id(word) {
        token.Type = ID
        return
    } else if is_op(word) {
        token.Type = OP
        return
    }
    return
}

func is_num(word string) bool {
    return numReg.MatchString(word)
}

func is_op(word string) bool {
    return opReg.MatchString(word)
}

func is_id(word string) bool {
    return idReg.MatchString(word)
}
