package internal

import (
	"errors"
	"fmt"
	"io"
	"strings"
	"unicode"
)

// TokenKind indicates the lexical unit of a token.
type TokenKind int

const (
	TUnknown TokenKind = iota
	TLeftParen
	TRightParen
	TString
	TNumber
	TBool
	TIdent
)

var tokenKinds = [...]string{
	TLeftParen:  "(",
	TRightParen: ")",
	TString:     "STRING",
	TNumber:     "NUMBER",
	TBool:       "BOOL",
	TIdent:      "IDENT",
}

func (t TokenKind) String() string {
	if t >= 0 && int(t) < len(tokenKinds) {
		return tokenKinds[t]
	}
	return tokenKinds[0]
}

// Token represents a token of the language.
type Token struct {
	Kind TokenKind
}

// Tokenize converts a string into a list of tokens.
func Tokenize(s string) ([]Token, error) {
	lexer := NewLexer(strings.NewReader(s))
	tokens := []Token{}

	for {
		token, err := lexer.Next()
		if err != nil {
			if err == io.EOF {
				break
			}
			return nil, err
		}
		tokens = append(tokens, *token)
	}

	return tokens, nil
}

// Lexer holds the state required to tokenize a source and generates
// a stream of tokens until the source is consumed.
type Lexer struct {
	scanner io.RuneScanner
}

// NewLexer creates a new Lexer initalized with the given io.RuneScanner.
func NewLexer(scanner io.RuneScanner) *Lexer {
	return &Lexer{scanner}
}

// Next returns the next token or io.EOF when the end
// of the source is reached.
func (l *Lexer) Next() (*Token, error) {
	if err := l.skipWhitespace(); err != nil {
		return nil, err
	}

	r, err := l.read()
	if err != nil {
		return nil, err
	}

	switch {
	case r == '(':
		return &Token{TLeftParen}, nil
	case r == ')':
		return &Token{TRightParen}, nil
	case r == '"':
		return l.readString()
	case unicode.IsDigit(r):
		return l.readNumber(r)
	case unicode.IsLetter(r):
		return l.readIdent(r)
	}

	return nil, fmt.Errorf("unexpected rune '%c'", r)
}

func (l *Lexer) read() (rune, error) {
	r, _, err := l.scanner.ReadRune()
	return r, err
}

func (l *Lexer) unread() error {
	return l.scanner.UnreadRune()
}

func (l *Lexer) skipWhitespace() error {
	for {
		r, err := l.read()
		if err != nil {
			return err
		}
		if !unicode.IsSpace(r) {
			if err := l.unread(); err != nil {
				return err
			}
			return nil
		}
	}
}

func (l *Lexer) readString() (*Token, error) {
	for {
		r, err := l.read()
		if err != nil {
			if err == io.EOF {
				return nil, errors.New("unterminated string literal")
			}
			return nil, err
		}
		if r == '"' {
			return &Token{TString}, nil
		}
	}
}

func (l *Lexer) readNumber(r rune) (*Token, error) {
	hasDot := false

	for {
		r, err := l.read()
		if err != nil {
			if err == io.EOF {
				return &Token{TNumber}, nil
			}
			return nil, err
		}

		if r == '.' && !hasDot {
			hasDot = true
			continue
		}

		if !unicode.IsDigit(r) {
			if err := l.unread(); err != nil {
				return nil, err
			}
			return &Token{TNumber}, nil
		}
	}
}

// keywords contains all the reserved keywords of the language.
var keywords = map[string]TokenKind{
	"true":  TBool,
	"false": TBool,
}

func (l *Lexer) readIdent(r rune) (*Token, error) {
	lexeme := []rune{r}

	for {
		r, err := l.read()
		if err != nil {
			if err == io.EOF {
				break
			}
			return nil, err
		}

		if !isAlphaNumeric(r) {
			if err := l.unread(); err != nil {
				return nil, err
			}
			break
		}

		lexeme = append(lexeme, r)
	}

	kind, ok := keywords[string(lexeme)]
	if !ok {
		return &Token{TIdent}, nil
	}

	return &Token{kind}, nil
}

// isAlphaNumeric returns whether the rule is an alphanumeric character.
func isAlphaNumeric(r rune) bool {
	return unicode.IsLetter(r) || unicode.IsDigit(r) || r == '_'
}
