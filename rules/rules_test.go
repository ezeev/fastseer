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

	q := SolrQueryRule{}
	q.MatchQueryTriggers = append(q.MatchQueryTriggers, "test rule")
	q.ActReplaceQuery = "replaced!"
	q.ActAddFqs = []string{"cat:001", "cat:002"}
	q.ActAddBqs = []string{"sku:001^0.5", "sku:002^6.5"}

	q.Prepare(params)

	q.Execute()

	t.Log(params)
}

func TestSolrQueryRule(t *testing.T) {
	params := &solrg.SolrParams{
		Q:  "this won't match",
		Fq: []string{"cat:001"},
	}
	ProcessSolrRules(params)
	t.Log(params)
}

func BenchmarkSolrRuleQuery(b *testing.B) {
	for n := 0; n < b.N; n++ {
		params := &solrg.SolrParams{
			Q: "test rule",
		}
		ProcessSolrRules(params)
	}
}

func TestContainsAnyRule(t *testing.T) {
	params := &solrg.SolrParams{
		Q: "test rule",
	}

	q := SolrQueryRule{}
	q.ContainsAnyQueryTriggers = []string{"test"}
	q.ActAddBqs = []string{"sku:001^0.9"}
	q.Prepare(params)
	q.Execute()
	t.Log(params)
}

func BenchmarkAnyRulePrepare(b *testing.B) {

	q := SolrQueryRule{}
	//241 ns/op

	q.ContainsAnyQueryTriggers = []string{"test", "six", "seven", "eight"}
	//634 ns/op

	q.MatchQueryTriggers = []string{"test rule one two three four five"}
	//781 ns/op

	q.ActAddBqs = []string{"sku:001^0.9"}
	//925 ns/op
	q.ActAddFqs = []string{"fq1:123", "fq2:234", "fq:345"}
	//1278 ns/op

	for n := 0; n < b.N; n++ {
		params := &solrg.SolrParams{
			Q: "test rule one two three four five",
		}

		q.Prepare(params)
		q.Execute()
		//t.Log(params)
	}
}
