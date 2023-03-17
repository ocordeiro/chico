package main

// #cgo LDFLAGS: -L. -lgpt
// #include "gpt.h"
import "C"
import (
	"database/sql"
	"encoding/binary"
	"fmt"
	"log"
	"math"
	"unsafe"

	"github.com/mattn/go-sqlite3"
)

func float32SliceToBytes(floats []float32) []byte {
	buf := make([]byte, len(floats)*4)
	for i, f := range floats {
		binary.LittleEndian.PutUint32(buf[i*4:], math.Float32bits(f))
	}
	return buf
}

func bytesToFloat32Slice(bytes []byte) []float32 {
	floats := make([]float32, len(bytes)/4)
	for i := range floats {
		floats[i] = math.Float32frombits(binary.LittleEndian.Uint32(bytes[i*4:]))
	}
	return floats
}

//export goCallback
func goCallback(embeddings *C.float, size C.int) {

	floatArr := (*[1 << 30]float32)(unsafe.Pointer(embeddings))[0:size]

	//s := int(size)
	//for i := 0; i < s; i++ {
	//fmt.Println(floatArr[i])
	//}

	byteArr := float32SliceToBytes(floatArr)

	sqlStmt := `insert into vss_docs (content_embedding) values (?);`
	_, err := db.Exec(sqlStmt, byteArr)
	if err != nil {
		log.Printf("%q: %s\n", err, sqlStmt)
		return
	}

}

var db *sql.DB

func main() {

	sql.Register("sqlite3_with_extensions",
		&sqlite3.SQLiteDriver{
			Extensions: []string{
				"vector0",
				"vss0",
			},
		})

	var err error
	db, err = sql.Open("sqlite3_with_extensions", "vector.db")
	if err != nil {
		log.Fatal(err)
	}

	createDb := `CREATE VIRTUAL TABLE IF NOT EXISTS 
		vss_docs using vss0(content_embedding(768)
		);`
	_, err = db.Exec(createDb)
	if err != nil {
		log.Printf("%q: %s\n", err, createDb)
		return
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
