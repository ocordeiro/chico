package main

// #cgo LDFLAGS: -L. -lgpt
// #include "gpt.h"
import "C"
import (
	"fmt"
	"unsafe"
)

//export goCallback
func goCallback(output *C.float) {

	floatArr := (*[1 << 30]float32)(unsafe.Pointer(output))[:10]

	for i := 0; i < int(10); i++ {
		fmt.Println(floatArr[i])
	}
}

func main() {
	C.get_embeddings(C.CString("Hello, World!"), C.GoCallback(C.goCallback))

	select {}
}
