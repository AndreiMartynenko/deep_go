package main

import "fmt"

func process(temp **int32) {
	var value2 int32 = 200
	*temp = &value2

}

func main() {
	var value int32 = 100
	var pointer *int32 = &value

	fmt.Println("address: ", pointer)
	fmt.Println("value: ", *pointer)

	process(&pointer)

	//*pointer = 500

	fmt.Println("address: ", pointer)
	fmt.Println("value: ", *pointer)
}
