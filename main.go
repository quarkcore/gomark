package main

import (
	"fmt"
	"os"
	"quarkcore/gomark/cmd"
)

func main() {
    if len(os.Args) < 2 {
        fmt.Println("Usage: gomark [command] [args]")
        os.Exit(1)
    }

    switch os.Args[1] {
    case "add":
        cmd.Add()
    case "open":
        cmd.Open()
    case "list":
        cmd.List()
    default:
        fmt.Println("gomark: unknown command")
    }
}
