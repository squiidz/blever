package main

import (
	"fmt"
	"log"

	"github.com/blevesearch/bleve"
	bbadger "github.com/squiidz/blever"
)

type User struct {
	ID       uint64
	Username string
	Password string
}

func main() {
	u := User{
		ID:       1,
		Username: "john Doe",
		Password: "a1b2c3",
	}
	// create/open bleveIndex
	index, err := bbadger.BleveIndex("badger/userIndex", bleve.NewIndexMapping())
	if err != nil {
		log.Fatal(err)
	}
	// index some data
	err = index.Index(u.Username, u)
	if err != nil {
		log.Fatal(err)
	}
	// search for some text
	query := bleve.NewMatchQuery("a1b2c3")
	search := bleve.NewSearchRequest(query)
	searchResult, err := index.Search(search)
	if err != nil {
		log.Fatal(err)
	}
	id := searchResult.Hits[0].ID
	doc, _ := index.Document(id)
	for _, f := range doc.Fields {
		fmt.Println(f.Name(), ":", string(f.Value()))
	}
}
