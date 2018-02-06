package main

import (
    "fmt"
    "os"
    "io/ioutil"
    "github.com/Petelliott/pb/lexer"
    "github.com/Petelliott/pb/parser"
    _ "github.com/Petelliott/pb/optimizer"
    "github.com/Petelliott/pb/coder"
)

func main() {
    bytes, _ := ioutil.ReadAll(os.Stdin)

    ti := lexer.NewTokenIterator(lexer.Lex(string(bytes)))
    ast := parser.ParseProgram(&ti)

    // TODO: default optimizer levels
    //ast = optimizer.OptSimpleExpr()

    fmt.Print(coder.GenMipsProgram(ast))

    // TODO: different targets
}
