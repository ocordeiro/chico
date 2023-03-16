package main

// #cgo LDFLAGS: -L. -lgpt
// #include "gpt.h"
import "C"
import (
	"database/sql"
	"fmt"
	"log"
	"unsafe"

	"github.com/mattn/go-sqlite3"
)

//export goCallback
func goCallback(embeddings *C.float, size C.int) {

	floatArr := (*[1 << 30]float32)(unsafe.Pointer(embeddings))[0:size]

	s := int(size)

	for i := 0; i < s; i++ {
		fmt.Println(floatArr[i])
	}
}

func main() {

	sql.Register("sqlite3_with_extensions",
		&sqlite3.SQLiteDriver{
			Extensions: []string{
				"vector0",
				"vss0",
			},
		})

	db, err := sql.Open("sqlite3_with_extensions", "vector.db")
	if err != nil {
		log.Fatal(err)
	}

	sqlStmt := `select vss_version();`
	data, err := db.Exec(sqlStmt)
	if err != nil {
		log.Printf("%q: %s\n", err, sqlStmt)
		return
	}
	fmt.Println(data)

	defer db.Close()

	C.get_embeddings(C.CString("Hello, World!"), C.GoCallback(C.goCallback))
}
