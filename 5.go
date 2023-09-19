// Mtutux


package main
import "fmt"

type mytype struct {
    counter int
}


func main() {
    mytypeInstance := mytype{}

    for i := 0; i < 5; i++ {
        go func(mytypeInstance mytype) {
            mytypeInstance.counter++
            }(mytypeInstance)
    }
    fmt.Printf("Counter: %d\n", mytypeInstance.counter)
}