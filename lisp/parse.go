package lisp

import (
    "errors"
    "strings"
    "unicode"
)

func isAtomRune(r rune) bool {
    return !unicode.IsSpace(r) && !(r == '(' || r == ')')
}

func tokenize(input string) []string {
    var tokens []string
    var b strings.Builder
    readingAtom := false

    for _, r := range input {
        if readingAtom {
            if isAtomRune(r) {
                b.WriteRune(r)
            } else {
                readingAtom = false
                tokens = append(tokens, b.String())
                b.Reset()

                if r == '(' {
                    tokens = append(tokens, "(")
                } else if r == ')' {
                    tokens = append(tokens, ")")
                }
            }
        } else {
            if isAtomRune(r) {
                readingAtom = true
                b.WriteRune(r)
            } else if r == '(' {
                tokens = append(tokens, "(")
            } else if r == ')' {
                tokens = append(tokens, ")")
            }
        }
    }

    return tokens
}

type Parser struct {
    tokens  []string
    pos     int
}

func (p *Parser) ParseExpr(input string) (*Expr, error) {
    if len(input) == 0 {
        return nil, errors.New("Empty input")
    }

    p.tokens = tokenize(input)
    p.pos = 0

    return p.parseExpr()
}

func (p *Parser) parseExpr() (*Expr, error) {
    if p.pos >= len(p.tokens) {
        return nil, errors.New("Missing closing parentesis")
    }

    switch p.tokens[p.pos] {
    case "(":
        p.pos++
        return p.parseList()

    case ")":
        return nil, errors.New("Closing parentesis before opening parentesis")

    default:
        atom := &Expr{atom: &p.tokens[p.pos]}
        p.pos++
        return atom, nil
    }
}

func (p *Parser) parseList() (*Expr, error) {
    if p.pos >= len(p.tokens) {
        return nil, errors.New("Missing closing parentesis")
    }

    if p.tokens[p.pos] == ")" {
        p.pos++
        return nil, nil
    } else {
        car, err := p.parseExpr()
        
        if err != nil {
            return nil, err
        }

        cdr, err := p.parseList()

        if err != nil {
            return nil, err
        }

        return &Expr{car: car, cdr: cdr}, nil
    }
}
