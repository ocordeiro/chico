package main

// #cgo LDFLAGS: -L. -lgpt
// #include "gpt.h"
import "C"

//export goCallback
func goCallback(output *C.char) {
	println(C.GoString(output))
}

func main() {
	C.get_embeddings(C.CString("Hello, World!"), C.GoCallback(C.goCallback))

	select {}
}
