package main

import "fmt"
import "time"

func main() {
    fmt.Printf("one \n")
    go testFunction()
    fmt.Printf("two\n")
    time.Sleep(3 * time.Second)
}

func testFunction() {
    fmt.Printf("checking\n")
}