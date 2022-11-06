package lisp

var trueLiteral string = "true"
var falseLiteral string = "false"

var trueAtom *Expr = &Expr{atom: &trueLiteral}
var falseAtom *Expr = &Expr{atom: &falseLiteral}
