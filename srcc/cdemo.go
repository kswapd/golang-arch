package srcc

/*
#cgo CFLAGS: -g -Wall
#include <stdio.h>
#include <stdlib.h>
#include "greeter.h"
int add(int a, int b) {
    return a + b;
}
*/
import "C"
import (
	"fmt"
	"unsafe"
)

func MyCallC() {
	x := 10
	y := 20
	sum := C.add(C.int(x), C.int(y))
	fmt.Printf("The sum of %d and %d is %d\n", x, y, sum)

	name := C.CString("Gopher")
	defer C.free(unsafe.Pointer(name))

	year := C.int(2018)

	g := C.struct_Greetee{
		name: name,
		year: year,
	}

	ptr := C.malloc(C.sizeof_char * 1024)
	defer C.free(unsafe.Pointer(ptr))

	size := C.greet(&g, (*C.char)(ptr))

	b := C.GoBytes(ptr, size)
	fmt.Println(string(b))
}
