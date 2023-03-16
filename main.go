package main

// #cgo LDFLAGS: -L. -lgpt
// #include "gpt.h"
import "C"
import (
	"fmt"
	"unsafe"
)

//export goCallback
func goCallback(embeddings *C.float, size C.int) {

	floatArr := (*[1 << 30]float32)(unsafe.Pointer(embeddings))[0:int(size)]

	for i := 0; i < int(size); i++ {
		fmt.Println(floatArr[i])
	}
}

func main() {
	C.get_embeddings(C.CString("Hello, World!"), C.GoCallback(C.goCallback))

	select {}
}
