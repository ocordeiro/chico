package main

// #cgo LDFLAGS: -L. -lgpt
// #include "gpt.h"
import "C"
import (
	"database/sql"
	"log"
	"path"

	"github.com/mattn/go-sqlite3"
)

//export goCallback
func goCallback(embeddings *C.float, size C.int) {

	//floatArr := (*[1 << 30]float32)(unsafe.Pointer(embeddings))[0:size]

	println(size)
}

var db *sql.DB

func main() {

	db = initDatabase()
	defer db.Close()

	C.get_embeddings(C.CString("Joa tilapia"), C.GoCallback(C.goCallback))
}

func initDatabase() *sql.DB {

	sqlitePath := "sqlite-vss"

	sql.Register("sqlite3_with_extensions",
		&sqlite3.SQLiteDriver{
			Extensions: []string{
				path.Join(sqlitePath, "vector0"),
				path.Join(sqlitePath, "vss0"),
			},
		})

	var err error

	db, err = sql.Open("sqlite3_with_extensions", path.Join(sqlitePath, "vector.db"))
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
		return nil
	}

	return db
}

func insertDocument(content string, embeddings []float32) {
	insert := `
		INSERT INTO documents (content, embeddings) 
		VALUES (?, ?);`
	_, err := db.Exec(insert, content, embeddings)
	if err != nil {
		log.Printf("%q: %s\n", err, insert)
		return
	}
}
