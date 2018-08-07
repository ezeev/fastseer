package search

import (
	"fmt"
)

type SearchEngine interface {
	CreateIndex(name string, address string, options map[string]string) error
	DeleteIndex(name string, address string) error
	IndexDocuments(index string, address string, docs interface{}) error
	IndexStruct(index string, address string, data interface{}) error
	DeleteDocuments(index string, address string, query string) error
	Query(index string, address string, query interface{}) (interface{}, error)
}

func NewSearchEngine(impl string) (SearchEngine, error) {
	if impl == "solr" {
		return &SolrSearch{}, nil
	}
	return nil, fmt.Errorf("Search engine implementation not found!")
}
