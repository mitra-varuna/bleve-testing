package main

import (
	"fmt"

	"github.com/blevesearch/bleve"
)

type School struct {
	ID          int
	Name        string
	Description string
}

func main() {

	mapping := bleve.NewIndexMapping()
	index, err := bleve.New("example.bleve", mapping)

	if err == bleve.ErrorIndexPathExists {
		index, err = bleve.Open("example.bleve")
	}

	if err != nil {
		fmt.Println("Error in creating index")
		fmt.Println(err)
		return
	}

	newSchool := School{ID: 1, Name: "Test School", Description: "Test Description"}

	err = index.Index("1", newSchool)

	if err != nil {
		fmt.Println(err)
		return
	}

	//Search
	query := bleve.NewMatchQuery("test")
	search := bleve.NewSearchRequest(query)
	results, err := index.Search(search)

	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(results)

}
