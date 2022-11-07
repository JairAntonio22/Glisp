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

    reader := bufio.NewReader(os.Stdin)
    logger := log.New(os.Stdout, "Error: ", 0)
    context := lisp.NewContext()

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
