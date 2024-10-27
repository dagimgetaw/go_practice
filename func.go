package main
import "fmt"

func greet(name string) string {
	return "Hello " + name + " how are u."
}

func main() {
	fmt.Println(greet("John"))
}