package main

import (
    "bufio"
    "fmt"
    "os"

    "./lisp"
)

func main() {
    var parser lisp.Parser

    reader := bufio.NewReader(os.Stdin)

    for {
        fmt.Print("> ")
        input, err := reader.ReadString('\n')

        if err != nil {
            break
        }

        expr, err := parser.ParseExpr(input)

        if err != nil {
            fmt.Fprintf(os.Stderr, "Error: %v\n", err)
            continue
        }

        fmt.Println(expr)
    }
}
