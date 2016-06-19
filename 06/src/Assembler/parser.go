package main

const (
    INVALID_COMMAND = 0
    A_COMMAND = 1
    C_COMMAND = 2
    L_COMMAND = 3
)

type Parser struct {
    FileName string
    lex *Lex
    CmdType int
    Symbol string
    Dest string
    Comp string
    Jump string
}

func NewParser(file string) (parser *Parser, err error) {
    var lex *Lex
    lex, err = NewLex(file)
    if err != nil {
        return
    }
    parser = &Parser {
        FileName: file,
        lex: lex,
        CmdType: INVALID_COMMAND,
        Symbol: "",
        Dest: "",
        Comp: "",
        Jump: "",
    }
    return
}

func (parser *Parser) Close() {
    parser.lex.Close()
    return
}

func (parser *Parser) initCmdInfo() {
    parser.Symbol = ""
    parser.Dest = ""
    parser.Comp = ""
    parser.Jump = ""
    parser.CmdType = INVALID_COMMAND
}

func (parser *Parser) HasMoreCommand() (yes bool, err error) {
    return parser.lex.HasMoreCommand()
}

func (parser *Parser) Advance() (err error){
    parser.initCmdInfo()
    _, err = parser.lex.NextCommand()
    if err != nil {
        return
    }

    token := parser.lex.CurrentToken
    if token.Type == OP && token.Word == "@" {
        parser.a_command()
    } else if token.Type == OP && token.Word == "(" {
        parser.l_command()
    } else {
        parser.c_command(token)
    }
    return
}

func (parser *Parser) a_command() {
    parser.CmdType = A_COMMAND
    token := parser.lex.NextToken()
    parser.Symbol = token.Word
    return
}

func (parser *Parser) l_command() {
    parser.CmdType = L_COMMAND
    token := parser.lex.NextToken()
    parser.Symbol = token.Word
    return
}

func (parser *Parser) c_command(token Token) {
    parser.CmdType = C_COMMAND
    comp_token := parser.get_dest(token)
    parser.get_comp(comp_token)
    parser.get_jump()
    return
}

func (parser *Parser) get_jump() {
    token := parser.lex.NextToken()
    if token.Type == OP && token.Word == ";" {
        jump_token := parser.lex.NextToken()
        parser.Jump = jump_token.Word
    }
    return
}

func (parser *Parser) get_dest(token Token) (comp_token Token){
    token_next := parser.lex.PeekToken()
    if token_next.Type == OP && token_next.Word == "=" {
        parser.lex.NextToken()
        parser.Dest = token.Word
        comp_token = parser.lex.NextToken()
    } else {
        comp_token = token
    }
    return
}

func (parser *Parser) get_comp(token Token) {
    if token.Type == OP && (token.Word == "-" || token.Word == "!") {
        token_next := parser.lex.NextToken()
        parser.Comp = token.Word + token_next.Word
    } else if token.Type == NUM || token.Type == ID {
        parser.Comp = token.Word
        token_next := parser.lex.PeekToken()
        if token_next.Type == OP && token_next.Word != ";" {
            parser.lex.NextToken()
            token_operand := parser.lex.NextToken()
            parser.Comp += token_next.Word + token_operand.Word
        }
    }
}

func (parser *Parser) CommandType() int {
    return parser.CmdType
}
