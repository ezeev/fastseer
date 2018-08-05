package rules

import (
	"log"

	"github.com/ezeev/solrg"
)

type Rule interface {
	AddTrigger(func(interface{}) bool)
	AddAction(func(interface{}) error)
	Execute(interface{}) error
}

func NewRule(impl string) Rule {
	if impl == "solr" {
		return &SolrRule{}
	}
	return nil
}

type QueryMatchRule struct {
	MatchQueryTriggers       []string `json:"matchQueryTriggers"`
	MatchFilterQueryTriggers []string `json:"matchFilterQueryTriggers"`
	ActReplaceQuery          string   `json:"actReplaceQuery"`
	ActAddFqs                []string `json:"actAddFqs"`
	ActAddBqs                []string `json:"actAddBqs"`
	Rule                     Rule     `json:"-"`
}

func (q *QueryMatchRule) Execute(params interface{}) {

	q.Rule = NewRule("solr")

	// triggers
	sparams := params.(*solrg.SolrParams)
	for _, match := range q.MatchQueryTriggers {
		trigger := func(interface{}) bool {
			return (match == sparams.Q)
		}
		q.Rule.AddTrigger(trigger)
	}

	for _, match := range q.MatchFilterQueryTriggers {
		trigger := func(interface{}) bool {
			for _, fq := range sparams.Fq {
				if fq == match {
					return true
				}
			}
			return false
		}
		q.Rule.AddTrigger(trigger)
	}

	// actions
	if q.ActReplaceQuery != "" {
		log.Println("line 55")
		f := func(params interface{}) error {
			sparams := params.(*solrg.SolrParams)
			sparams.Q = q.ActReplaceQuery
			return nil
		}
		q.Rule.AddAction(f)
	}
	if len(q.ActAddFqs) > 0 {
		f := func(params interface{}) error {
			sparams := params.(*solrg.SolrParams)
			for _, fq := range q.ActAddFqs {
				sparams.Fq = append(sparams.Fq, fq)
			}
			return nil
		}
		q.Rule.AddAction(f)
	}
	if len(q.ActAddBqs) > 0 {
		f := func(params interface{}) error {
			sparams := params.(*solrg.SolrParams)
			for _, bq := range q.ActAddBqs {
				sparams.Bq = sparams.Bq + " " + bq
			}
			return nil
		}
		q.Rule.AddAction(f)
	}
	q.Rule.Execute(params)

}
