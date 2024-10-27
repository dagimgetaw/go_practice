package main
import "fmt"

func main() {
	var age int = 5
	var name string = "name"
	var height float32 = 1.74

	n1, n2 := 3, 4
	text := "GO"
	const unchanged_text string = "This text can't be changed"

	fmt.Println(age)
	fmt.Println(name)
	fmt.Println(height)
	fmt.Println(n1, n2)
	fmt.Println(text)
	fmt.Println(unchanged_text)
}

