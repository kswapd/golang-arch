package srcobjc

// #cgo CFLAGS: -x objective-c
// #cgo LDFLAGS: -framework Foundation
// #include "bridge.h"
import "C"

func MyObjectiveCMethod() {
	C.callObjectiveCMethod()
}
