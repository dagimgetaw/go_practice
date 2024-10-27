package main
import "fmt"

func main() {
	var arr1 [3]int32
	var arr2 [3]int32 = [3]int32{1,2,3}

	fmt.Println(arr1)
	fmt.Println(arr2)

	fmt.Println(arr2[1])

	arr2[0] = 8

	fmt.Println(arr2)

	arr3 := [5]int32{1,2,3,4,5}
	arr4 := [...]int32{1,3,5,7,9}

	fmt.Println(arr3)
}