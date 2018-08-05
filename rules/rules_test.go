package rules

import (
	"testing"

	"github.com/ezeev/solrg"
)

func TestRule(t *testing.T) {

	params := &solrg.SolrParams{
		Q: "test rule",
	}

	//trigger := func(params *solrg.SolrParams) bool {
	trigger := func(params interface{}) bool {
		sparams := params.(*solrg.SolrParams)
		return (sparams.Q == "test rule")
	}
	action1 := func(params interface{}) error {
		sparams := params.(*solrg.SolrParams)
		sparams.Fq = append(sparams.Fq, "cat:001")
		return nil
	}

	action2 := func(params interface{}) error {
		sparams := params.(*solrg.SolrParams)
		sparams.Bq = "sku:0001^0.1"
		return nil
	}

	rule := NewRule("solr")
	rule.AddTrigger(trigger)
	rule.AddAction(action1)
	rule.AddAction(action2)

	rule.Execute(params)

	t.Log(params)
}

func TestQueryRule(t *testing.T) {

	params := &solrg.SolrParams{
		Q: "test rule",
	}

	q := QueryMatchRule{}
	q.MatchQueryTriggers = append(q.MatchQueryTriggers, "test rule")
	q.ActReplaceQuery = "replaced!"
	q.ActAddFqs = []string{"cat:001", "cat:002"}
	q.ActAddBqs = []string{"sku:001^0.5", "sku:002^6.5"}

	q.Execute(params)

	t.Log(params)

}
