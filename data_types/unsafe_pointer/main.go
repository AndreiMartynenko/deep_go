package main

import (
	"fmt"
	"unsafe"
)

/*
We know that, for some cases, the unsafe mechanism can help
us write more efficient Go code. However, it is very easy to
introduce some subtle bugs which have very low possibilities
to produce when using the unsafe mechanism. A program with
these bugs may run well for a long time, but suddenly behave
abnormally and even crash at a later time. Such bugs are very
hard to detect and debug

Another big risk of using unsafe pointers comes from the
fact that the unsafe mechanism is not protected by the Go
compatibility guideline - code depending on unsafe
pointers works today could break since a later Go version
*/
func main() {
	var value uint32 = 0xFFFFFFFF

	pointer := unsafe.Pointer(&value)
	bytePointer := (*uint8)(pointer)

	fmt.Println("value1:", *bytePointer)

	pointer = unsafe.Add(pointer, 2)
	twoBytePointer := (*uint16)(pointer)

	fmt.Println("value2:", *twoBytePointer)
}
