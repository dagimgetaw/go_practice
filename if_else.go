package main
import "fmt"

func main() {
	var name string = "Joe"
	var num int = 3
	var result string = result(num)

	fmt.Printf("%v ur number %v is %v", name, num, result)
}

func result(num int) string {
	if num == 1 {
		return "one"
	} else if num == 2 {
		return "two"
	} else {
		return  "three"
	}
}