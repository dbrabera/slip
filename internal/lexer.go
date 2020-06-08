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
	TInt
	TBool
	TSymbol
)

var tokenKinds = [...]string{
	TLeftParen:  "(",
	TRightParen: ")",
	TString:     "STRING",
	TInt:        "INT",
	TBool:       "BOOL",
	TSymbol:     "SYMBOL",
}

func (t TokenKind) String() string {
	if t >= 0 && int(t) < len(tokenKinds) {
		return tokenKinds[t]
	}
	return tokenKinds[0]
}

// Token represents a token of the language.
type Token struct {
	Kind   TokenKind
	Lexeme string
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
	scanner   io.RuneScanner
	lookahead *lookahead
}

type lookahead struct {
	token *Token
	err   error
}

// NewLexer creates a new Lexer initalized with the given io.RuneScanner.
func NewLexer(scanner io.RuneScanner) *Lexer {
	return &Lexer{scanner, nil}
}

// Next returns the next token or io.EOF when the end
// of the source is reached.
func (l *Lexer) Next() (*Token, error) {
	if l.lookahead != nil {
		token, err := l.lookahead.token, l.lookahead.err
		l.lookahead = nil
		return token, err
	}

	if err := l.skipWhitespace(); err != nil {
		return nil, err
	}

	r, err := l.read()
	if err != nil {
		return nil, err
	}

	switch {
	case r == '(':
		return &Token{TLeftParen, ""}, nil
	case r == ')':
		return &Token{TRightParen, ""}, nil
	case r == '"':
		return l.readString()
	case unicode.IsDigit(r):
		return l.readInt(r)
	case isIdent(r):
		return l.readIdent(r)
	}

	return nil, fmt.Errorf("unexpected rune '%c'", r)
}

func (l *Lexer) Peek() (*Token, error) {
	if l.lookahead != nil {
		return l.lookahead.token, l.lookahead.err
	}

	token, err := l.Next()
	if err != nil && err != io.EOF {
		return nil, err
	}

	l.lookahead = &lookahead{
		token: token,
		err:   err,
	}
	return token, err
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
	buf := []rune{}

	for {
		r, err := l.read()
		if err != nil {
			if err == io.EOF {
				return nil, errors.New("unterminated string literal")
			}
			return nil, err
		}
		if r == '"' {
			return &Token{TString, string(buf)}, nil
		}
		buf = append(buf, r)
	}
}

func (l *Lexer) readInt(r rune) (*Token, error) {
	buf := []rune{r}

	for {
		r, err := l.read()
		if err != nil {
			if err == io.EOF {
				return &Token{TInt, string(buf)}, nil
			}
			return nil, err
		}

		if !unicode.IsDigit(r) {
			if err := l.unread(); err != nil {
				return nil, err
			}
			return &Token{TInt, string(buf)}, nil
		}

		buf = append(buf, r)
	}
}

// keywords contains all the reserved keywords of the language.
var keywords = map[string]TokenKind{
	"true":  TBool,
	"false": TBool,
}

func (l *Lexer) readIdent(r rune) (*Token, error) {
	buf := []rune{r}

	for {
		r, err := l.read()
		if err != nil {
			if err == io.EOF {
				break
			}
			return nil, err
		}

		if !isIdent(r) {
			if err := l.unread(); err != nil {
				return nil, err
			}
			break
		}

		buf = append(buf, r)
	}

	lexeme := string(buf)

	kind, ok := keywords[lexeme]
	if !ok {
		return &Token{TSymbol, lexeme}, nil
	}

	return &Token{kind, lexeme}, nil
}

// isIdent returns whether the rune can belong to an identifier.
func isIdent(r rune) bool {
	switch r {
	case '!', '@', '$', '%', '^', '&', '*', '-', '_', '+', '=', '|',
		'~', ':', '<', '>', '.', '?', '\\', '/', ',':
		return true
	default:
		return unicode.IsLetter(r) || unicode.IsDigit(r)
	}
}
