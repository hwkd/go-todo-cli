package parser

type Scanner struct {
	input        string
	char         byte
	position     int
	readPosition int
}

func NewScanner(input string) *Scanner {
	scanner := &Scanner{input: input}
	scanner.readChar()
	return scanner
}

func (s *Scanner) NextToken() *Token {
	var token *Token

	s.skipWhiteSpaces()

	switch s.char {
	case '-':
		s.readChar()
		// Assume option character exists after '-'.
		token = &Token{
			Type:    OPTION,
			Literal: string(s.char),
		}
	case '"':
		s.readChar()
		token = &Token{
			Type:    VALUE,
			Literal: s.readString(),
		}
	default:
		token = &Token{}
	}

	s.readChar()
	return token
}

func (s *Scanner) readChar() {
	if s.readPosition >= len(s.input) {
		s.char = 0
	} else {
		s.char = s.input[s.readPosition]
	}
	s.position = s.readPosition
	s.readPosition++
}

func (s *Scanner) peekChar() byte {
	if s.readPosition >= len(s.input) {
		return 0
	} else {
		return s.input[s.readPosition]
	}
}

func (s *Scanner) skipWhiteSpaces() {
	for s.char == ' ' || s.char == '\n' || s.char == '\t' || s.char == '\r' {
		s.readChar()
	}
}

func (s *Scanner) readString() string {
	position := s.position
	for s.char != '"' {
		s.readChar()
	}
	str := s.input[position:s.position]
	s.readChar()
	return str
}
