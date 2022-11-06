package main

import (
    "bufio"
    "fmt"
    "log"
    "os"

    "./lisp"
)

func main() {
    var parser lisp.Parser
    var context lisp.Context

    logger := log.New(os.Stderr, "Error: ", 0)
    reader := bufio.NewReader(os.Stdin)

    for {
        fmt.Print("> ")
        input, err := reader.ReadString('\n')

        if err != nil {
            break
        }

        expr, err := parser.ParseExpr(input)

        if err != nil {
            logger.Println(err)
            continue
        }

        expr, err = context.Eval(expr)

        if err != nil {
            logger.Println(err)
            continue
        }

        fmt.Println(expr)
    }
}
