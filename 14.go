// command line arguments with conditionals

package main

import (
    "fmt"
    "os"
)

func main(){
    args:= os.Args

    if len(args) < 2 {
        fmt.Println("Error: not enough arguments")
        os.Exit(1)
    }


    fmt.Printf("hello world \nos.Args: %v\n", args, args[1:])
}