// User Input- Scanf-Multiple input

package main
import "fmt"

func main() {
    var name string
    var is_muggle bool

    fmt.Print("Entre names: ")
    fmt.Scanf("%s %t", &name, &is_muggle)
    fmt.Println(name, is_muggle)
}