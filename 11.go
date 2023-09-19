// User Input- Scanf


package main
import "fmt"

func main() {
    var name string
    fmt.Print("Enter name: ")
    fmt.Scanf("%s", &name)
    fmt.Println("Hey there, ",name)
}