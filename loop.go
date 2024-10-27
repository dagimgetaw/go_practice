package main
import "fmt"

func main() {
    var i int = 0

	for {
		if i >= 10 {
			break
		}
		fmt.Println(i)
		i = i + 1
	}

	for n := 0; n < 10; n++ {
		fmt.Println("loop")
	}
}