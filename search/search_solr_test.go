package search

import (
	"testing"
	"time"

	"github.com/ezeev/solrg"
)

const solrAddr = "172.104.9.135:8983/solr"

func TestSolrImpl(t *testing.T) {

	solr, err := NewSearchEngine("solr")
	if err != nil {
		t.Error(err)
	}

	//create a collection
	err = solr.CreateIndex("test2", solrAddr, map[string]string{"numShards": "1", "replicationFactor": "1"})
	if err != nil {
		t.Error(err)
	}
	time.Sleep(time.Second * 2)

	//try adding a couple of docs
	docs := solrg.NewSolrDocumentCollection()
	doc1 := solrg.NewSolrDocument("1")
	doc1.SetField("name_s", []string{"1"})
	doc2 := solrg.NewSolrDocument("2")
	doc2.SetField("name_s", []string{"2"})

	docs.AddDoc(doc1)
	docs.AddDoc(doc2)

	//index
	err = solr.IndexDocuments("test2", solrAddr, &docs)
	if err != nil {
		t.Error(err)
	}

	//query
	query := solrg.SolrParams{
		Q: "*:*",
	}
	resp, err := solr.Query("test2", solrAddr, &query)
	if err != nil {
		t.Error(err)
	}
	solrResp := resp.(*solrg.SolrSearchResponse)
	if len(solrResp.Response.Docs) != 2 {
		t.Error("Expected 2 results!")
	}

	//delete
	err = solr.DeleteDocuments("test2", solrAddr, "*:*")
	if err != nil {
		t.Error(err)
	}

	//delete it
	err = solr.DeleteIndex("test2", solrAddr)
	if err != nil {
		t.Error(err)
	}

}

func TestSolrG(t *testing.T) {
	sc, err := solrg.NewDirectSolrClient(solrAddr)
	if err != nil {
		t.Error(err)
	}

	// create test collection
	err = sc.CreateCollection("test", 1, 1, time.Second*10)
	if err != nil {
		t.Error(err)
	}
	t.Log("Collection created")

	// try creating it again (we expect an error)
	err = sc.CreateCollection("test", 1, 1, time.Second*10)
	if ae, ok := err.(*solrg.SolrCollectionExistsError); ok {
		t.Logf("Received SolrCollectionExistsError as expected: %s", ae.Error())
	}

	time.Sleep(time.Second * 3)

	// now delete it
	err = sc.DeleteCollection("test")
	if err != nil {
		t.Error(err)
	}
}
