package main

import "fmt"

/*
Pointer
A variable whose range of values consists of memory cell addresses
or a special value — the zero address (nil).

Т* <---> unsafe. Pointer <---> unitptr
Pointer represents a pointer to an arbitrary type. There
are four special operations available for type Pointer
that are not available for other types
- A pointer value of any type can be converted to a Pointer
- A Pointer can be converted to a pointer value of any type
- A uintptr can be converted to a Pointer
- A Pointer can be converted to a uintptr.
uintptr
This is an unsigned integer type which is large enough
to hold any pointer address. Therefore its size is
platform dependent. It is just an integer
representation of an address.

*/
//How to get the address of the variable
// 1 way
// pointer := new(int)

//2 way
// var value int
//pointer := &value

//3 way - only for structs
//pointer := &int{}

func main() {
	var value int32 = 100
	var pointer *int32 = &value

	fmt.Println("address: ", pointer)
	fmt.Println("value: ", *pointer)

	*pointer = 500

	fmt.Println("address: ", pointer)
	fmt.Println("value: ", *pointer)
}
