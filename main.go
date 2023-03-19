package main

// #cgo LDFLAGS: -L. -lgpt
// #include "gpt.h"
import "C"
import (
	"bytes"
	"database/sql"
	"encoding/binary"
	"encoding/json"
	"log"
	"path"
	"unsafe"

	"github.com/mattn/go-sqlite3"
)

var db *sql.DB

func float32SliceToByteSlice(floats []float32) []byte {
	buf := new(bytes.Buffer)
	for _, f := range floats {
		binary.Write(buf, binary.LittleEndian, f)
	}
	return buf.Bytes()
}

func main() {

	db = initDatabase()
	defer db.Close()

	document := "Quero comprar um carros."

	embeddings := C.get_embeddings(C.CString(document))

	floatArr := (*[768]float32)(unsafe.Pointer(embeddings))[:]

	jsonData, _ := json.Marshal(floatArr)

	insertDocument(document, string(jsonData))

	indexDocuments()

	getIndexes()
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

func insertDocument(content string, embeddings string) {

	insert := `
		INSERT INTO documents (content, embeddings) 
		VALUES (?, vector_to_blob(vector_from_json(?)))`
	_, err := db.Exec(insert, content, embeddings)
	if err != nil {
		log.Printf("%q: 2 %s\n", err, insert)
		return
	}
}

func indexDocuments() {
	index := `
		INSERT INTO vss_documents (rowid, embeddings) 
		SELECT rowid, embeddings FROM documents
		WHERE rowid NOT IN (SELECT rowid FROM vss_documents);`
	_, err := db.Exec(index)
	if err != nil {
		log.Printf("%q: %s\n", err, index)
		return
	}
}

func getIndexes() {
	get := `select count(*) from vss_documents`
	var count int

	err := db.QueryRow(get).Scan(&count)
	if err != nil {
		log.Printf("%q: %s\n", err, get)
		return
	}
	log.Printf("Count: %v", count)

}
