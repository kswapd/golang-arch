package srccpp

// #include "library-bridge.h"
import "C"
import (
	"fmt"
	"unsafe"
)

// #cgo LDFLAGS: -L. -llibrary
// #include "library-bridge.h"
// If used above comment, should generate library first:
// clang++ -o liblibrary.so library.cpp library-bridge.cpp -std=c++17 -O3 -Wall -Wextra -fPIC -shared
// reference: https://stackoverflow.com/questions/1713214/how-to-use-c-in-go
type Foo struct {
	ptr unsafe.Pointer
}

func NewFoo(value int) Foo {
	var foo Foo
	foo.ptr = C.LIB_NewFoo(C.int(value))
	return foo
}

func (foo Foo) Free() {
	C.LIB_DestroyFoo(foo.ptr)
}

func (foo Foo) value() int {
	return int(C.LIB_FooValue(foo.ptr))
}

func CallCppClass() {
	foo := NewFoo(888)
	defer foo.Free() // The Go analog to C++'s RAII
	fmt.Println("[go]", foo.value())
}
