package internal

import (
	"errors"
	"fmt"
	"io"
	"strconv"
	"strings"
)

// Parse converts a string into a list of values.
func Parse(s string) ([]Value, error) {
	parser := NewParser(NewLexer(strings.NewReader(s)))
	return parser.Parse()
}

// Parser implements a recursive descent parser.
//
// It accepts the following grammar:
//
// root  = value { value }
// value = list | INT | BOOL | STRING | SYMBOL
// list  = '(' { value } ')'
type Parser struct {
	lexer *Lexer
}

func NewParser(lexer *Lexer) *Parser {
	return &Parser{lexer}
}

func (p *Parser) Parse() ([]Value, error) {
	values := []Value{}

	for {
		value, err := p.parseValue()
		if err != nil {
			if err == io.EOF {
				break
			}
			return nil, err
		}
		values = append(values, value)
	}

	return values, nil
}

func (p *Parser) parseValue() (Value, error) {
	token, err := p.lexer.Peek()
	if err != nil {
		return nil, err
	}

	switch token.Kind {
	case TLeftParen:
		return p.parseList()
	case TInt:
		return p.parseInt()
	case TBool:
		return p.parseBool()
	case TString:
		return p.parseString()
	case TSymbol:
		return p.parseSymbol()
	default:
		return nil, fmt.Errorf("unexpected token '%s'", token.Kind)
	}
}

func (p *Parser) parseList() (Value, error) {
	if err := p.advance(); err != nil {
		return nil, err
	}

	list := NewList()

	for {
		token, err := p.lexer.Peek()
		if err != nil {
			if err == io.EOF {
				return nil, errors.New("unterminated list")
			}
			return nil, err
		}

		if token.Kind == TRightParen {
			if err := p.advance(); err != nil {
				return nil, err
			}
			break
		}

		value, err := p.parseValue()
		if err != nil {
			return nil, err
		}

		list = append(list, value)
	}

	return list, nil
}

func (p *Parser) parseInt() (Value, error) {
	token, err := p.match(TInt)
	if err != nil {
		return nil, err
	}

	val, err := strconv.ParseInt(token.Lexeme, 10, 64)
	if err != nil {
		return nil, err
	}

	return NewInt(val), nil
}

func (p *Parser) parseBool() (Value, error) {
	token, err := p.match(TBool)
	if err != nil {
		return nil, err
	}

	val, err := strconv.ParseBool(token.Lexeme)
	if err != nil {
		return nil, err
	}

	return NewBool(val), nil
}

func (p *Parser) parseString() (Value, error) {
	token, err := p.match(TString)
	if err != nil {
		return nil, err
	}
	return NewString(token.Lexeme), nil
}

func (p *Parser) parseSymbol() (Value, error) {
	token, err := p.match(TSymbol)
	if err != nil {
		return nil, err
	}
	return NewSymbol(token.Lexeme), nil
}

func (p *Parser) advance() error {
	_, err := p.lexer.Next()
	return err
}

func (p *Parser) match(kind TokenKind) (*Token, error) {
	token, err := p.lexer.Next()
	if err != nil {
		return nil, err
	}
	if token.Kind != kind {
		return nil, fmt.Errorf("unexpected token '%s'", token.Kind)
	}
	return token, nil
}
