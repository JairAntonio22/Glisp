package lisp

import "strings"

type Expr struct {
    atom    *string
    car     *Expr
    cdr     *Expr
}

func (e *Expr) String() string {
    if e == nil {
        return "<nil>"
    }

    var b strings.Builder

    if e.atom != nil {
        b.WriteString(*e.atom)
    } else {
        b.WriteRune('(')
        b.WriteString(e.car.String())

        for it := e.cdr; it != nil; it = it.cdr {
            b.WriteRune(' ')
            b.WriteString(it.car.String())
        }

        b.WriteRune(')')
    }

    return b.String()
}
