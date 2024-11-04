package parser

type TokenType string

type Token struct {
	Type    TokenType
	Literal string
}

const (
	OPTION = "OPTION"
	VALUE  = "VALUE"
	EOF    = "EOF"
)
