Blever
=====================

### Usage

> `➜ go get github.com/alash3al/bbadger` .

```go
package main

import (
	"fmt"

	"github.com/alash3al/bbadger"
	"github.com/blevesearch/bleve"
)

func main() {
	// create/open bleveIndex
	index, err := bbadger.BleveIndex("/tmp/badger/indexName", bleve.NewIndexMapping())

    // index some data
    err = index.Index(identifier, your_data)

    // search for some text
    query := bleve.NewMatchQuery("text")
    search := bleve.NewSearchRequest(query)
    searchResults, err := index.Search(search)
}

```
