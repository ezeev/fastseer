package search

import (
	"fmt"
	"log"
	"strconv"
	"time"

	"github.com/ezeev/fastseer/util"
	"github.com/ezeev/solrg"
)

type SolrSearch struct {
}

func (s *SolrSearch) CreateIndex(name string, address string, options map[string]string) error {

	err := util.ValidateOptionsMap(options, "numShards", "replicationFactor")
	if err != nil {
		return err
	}

	shards, err := strconv.Atoi(options["numShards"])
	if err != nil {
		return err
	}
	replicas, err := strconv.Atoi(options["replicationFactor"])
	if err != nil {
		return err
	}

	sc, err := solrg.NewDirectSolrClient(address)
	if err != nil {
		return err
	}

	createErr := sc.CreateCollection(name, shards, replicas, time.Second*10)
	if _, ok := createErr.(*solrg.SolrCollectionExistsError); ok {
		// collection already exists, log but don't throw an error
		log.Println("Collection already exists - not throwing error!")
	} else {
		return err
	}
	return nil
}

func (s *SolrSearch) DeleteIndex(name string, address string) error {

	sc, err := solrg.NewDirectSolrClient(address)
	if err != nil {
		return err
	}
	err = sc.DeleteCollection(name)
	return err
}

func (s *SolrSearch) IndexDocuments(indexName string, solrUrl string, docs interface{}) error {

	if solrDocs, ok := docs.(*solrg.SolrDocumentCollection); ok {

		cli, err := solrg.NewDirectSolrClient(solrUrl)
		if err != nil {
			return err
		}
		err = cli.PostDocs(solrDocs, indexName)
		if err != nil {
			log.Printf("Error indexing documents: %s", err.Error())
			return err
		}
		cli.Commit(indexName)

	} else {
		return fmt.Errorf("docs needs to be a *solrg.SolrDocumentCollection")
	}

	return nil
}

func (s *SolrSearch) DeleteDocuments(indexName string, solrUrl string, query string) error {

	cli, err := solrg.NewDirectSolrClient(solrUrl)
	if err != nil {
		return err
	}
	err = cli.DeleteByQuery(indexName, query)
	if err != nil {
		return err
	}
	return cli.Commit(indexName)

}

func (s *SolrSearch) Query(index string, solrUrl string, query interface{}) (interface{}, error) {

	if solrQuery, ok := query.(*solrg.SolrParams); ok {

		cli, err := solrg.NewDirectSolrClient(solrUrl)
		if err != nil {
			return nil, err
		}
		return cli.Query(index, "select", solrQuery, time.Second*5)

	} else {
		return nil, fmt.Errorf("query must be of type *solrg.SolrParams for Solr implementation")
	}
	return nil, nil
}
