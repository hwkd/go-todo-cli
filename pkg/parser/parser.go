package parser

type Parser struct {
	scanner *Scanner
}

func NewParser(input string) *Parser {
	return &Parser{scanner: NewScanner(input)}
}

func (p *Parser) ParseArgs() [][]string {
	return [][]string{}
}
