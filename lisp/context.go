package lisp

import (
    "errors"
    "fmt"
)

type Context struct {
    atoms map[string]*Expr
}

func NewContext() *Context {
    return &Context{atoms: make(map[string]*Expr)}
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

    if fn.IsAtom() {
        switch *fn.atom {
        case "car":
            return x.car.car, nil

        case "cdr":
            return x.car.cdr, nil

        case "cons":
            return &Expr{car: x.car, cdr: x.cdr}, nil

        case "atom?":
            if x.car.IsAtom() {
                return trueAtom, nil
            } else {
                return falseAtom, nil
            }

        case "eq?":
            arg1 := x.car.atom
            arg2 := x.cdr.car.atom

            if arg1 == nil || arg2 == nil {
                return nil, errors.New("eq? arguments must be atoms")
            }

            if *arg1 == *arg2 {
                return trueAtom, nil
            } else {
                return falseAtom, nil
            }

        default:
            eval_fn, err := c.eval(fn)

            if err != nil {
                return nil, err
            }

            return c.apply(eval_fn, x)
        }
    } else {
        switch *fn.car.atom {
        case "lambda":
            return nil, errors.New("lambda not implemented")

        case "define":
            return nil, errors.New("define not implemented")

        default:
            return nil, fmt.Errorf("%s not implemented", *fn.car.atom)
        }
    }
}

func (c *Context) eval(expr *Expr) (*Expr, error) {
    if expr.IsAtom() {
        expr_def, exists := c.atoms[*expr.atom]

        if !exists {
            return nil, fmt.Errorf("%s is not defined", *expr.atom)
        }

        return expr_def.cdr, nil

    } else if expr.car.IsAtom() {
        switch *expr.car.atom {
        case "quote":
            return expr.cdr.car, nil

        case "cond":

        default:
        }
    } else {
    }

    return nil, nil
}
