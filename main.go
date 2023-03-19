package main

// #cgo LDFLAGS: -L. -lgpt
// #include "gpt.h"
import "C"
import (
	"database/sql"
	"log"
	"unsafe"

	"github.com/mattn/go-sqlite3"
)

//export goCallback
func goCallback(embeddings *C.float, size C.int) {

	floatArr := (*[1 << 30]float32)(unsafe.Pointer(embeddings))[0:size]
	//to BLOB

	blob := make([]byte, len(floatArr)*4)
	for i, v := range floatArr {
		blob[i*4] = byte(v)
	}

	_, err := db.Exec(`INSERT INTO documents(embeddings) VALUES (?)`, blob)

	if err != nil {
		log.Fatal(err)
	}

	_, err = db.Exec(`INSERT INTO vss_documents(rowid,embeddings) 
		SELECT rowid,embeddings FROM documents
		`)
	if err != nil {
		log.Fatal(err)
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

	createDb := `
		CREATE TABLE IF NOT EXISTS documents (
			rowid INTEGER PRIMARY KEY,
			content TEXT,
			embeddings BLOB
		);
		CREATE VIRTUAL TABLE IF NOT EXISTS 
		vss_documents USING vss0(
			embeddings(768)
		);`
	_, err = db.Exec(createDb)
	if err != nil {
		log.Printf("%q: %s\n", err, createDb)
		return
	}

	defer db.Close()

	C.get_embeddings(C.CString("Joa tilapia"), C.GoCallback(C.goCallback))
}
