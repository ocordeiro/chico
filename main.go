package main

// #cgo LDFLAGS: -L. -lgpt
// #include "gpt.h"
import "C"
import (
	"bytes"
	"database/sql"
	"encoding/binary"
	"encoding/json"
	"fmt"
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

	document := "Quero comprar um carros"
	embeddings := jsonEmbeddings(document)
	insertDocument(document, embeddings)

	document = "Quero comprar uma moto"
	embeddings = jsonEmbeddings(document)
	insertDocument(document, embeddings)

	indexDocuments()

	document = "Quero comprar uma toyota"
	embeddings = jsonEmbeddings(document)

	search()
	//list()
	listIndex()
	countIndexes()
}

func jsonEmbeddings(text string) string {
	embeddings := C.get_embeddings(C.CString(text))

	floatArr := (*[768]float32)(unsafe.Pointer(embeddings))[:]

	jsonData, _ := json.Marshal(floatArr)

	return string(jsonData)
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
			embeddings(768) with "Flat,IDMap2"
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

func countIndexes() {
	get := `select count(*) from vss_documents`
	var count int

	err := db.QueryRow(get).Scan(&count)
	if err != nil {
		log.Printf("%q: %s\n", err, get)
		return
	}
	log.Printf("Count: %v", count)

}

func search() {

	fmt.Println("Search")

	stml := `select rowid, distance
	from vss_documents
	where vss_search(
		embeddings,
		vss_search_params(
			vector_from_blob((select embeddings from documents where rowid = 1)),
			128
		)
	)`

	rows, err := db.Query(stml)
	if err != nil {
		log.Printf("%q: %s\n", err, stml)
		return
	}
	defer rows.Close()

	for rows.Next() {
		var rowid int
		var distance float64
		err = rows.Scan(&rowid, &distance)
		if err != nil {
			log.Fatal(err)
		}
		log.Printf("Rowid: %v, Distance: %v", rowid, distance)
	}

}

func list() {

	fmt.Println("List")

	stml := `select rowid, content, embeddings from documents`

	rows, err := db.Query(stml)
	if err != nil {
		log.Printf("%q: %s\n", err, stml)
		return
	}
	defer rows.Close()

	for rows.Next() {
		var rowid int
		var content string
		var embeddings []byte
		err = rows.Scan(&rowid, &content, &embeddings)
		if err != nil {
			log.Fatal(err)
		}
		log.Printf("Rowid: %v, Content: %v, Embeddings: %v", rowid, content, embeddings)
	}

}

func listIndex() {

	fmt.Println("List Index")

	stmt := `select rowid, embeddings from vss_documents`

	rows, err := db.Query(stmt)
	if err != nil {
		log.Printf("%q: %s\n", err, stmt)
		return
	}
	defer rows.Close()

	for rows.Next() {
		var rowid int
		var embeddings []byte
		err = rows.Scan(&rowid, &embeddings)
		if err != nil {
			log.Fatal(err)
		}
		log.Printf("Rowid: %v, Embeddings: %v", rowid, embeddings)
	}

}
