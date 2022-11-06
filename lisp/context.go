package lisp

import (
    "errors"
    "os"
)

type Context struct {
    symbols map[string]*Expr
}

func (c *Context) Eval(expr *Expr) (*Expr, error) {
    if expr.atom != nil {
        return expr, nil
    } else {
        return c.apply(expr.car, expr.cdr)
    }
}

func (c *Context) apply(fn, x *Expr) (*Expr, error) {
    if fn == nil {
        return nil, errors.New("nil is not a procedure")
    }

    if fn.atom != nil {
        switch *fn.atom {
        case "car":
            return x.car.car, nil

        case "cdr":
            return x.car.cdr, nil

        case "cons":
            return &Expr{car: x.car, cdr: x.cdr}, nil

        case "atom?":
            if x.car.atom != nil {
                return trueAtom, nil
            } else {
                return falseAtom, nil
            }

        case "eq?":
            arg1 := x.car.atom
            arg2 := x.cdr.car.atom

            if arg1 != nil && arg2 != nil {
                if *arg1 == *arg2 {
                    return trueAtom, nil
                } else {
                    return falseAtom, nil
                }
            } else {
                return nil, errors.New("eq arguments must be atoms")
            }

        case "apply":
            return c.apply(x.car, x.cdr)

        case "exit":
            os.Exit(0)

        default:
        }
    } else {
    }

    return nil, nil
}
