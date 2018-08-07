package rules

import (
	"strings"

	"github.com/ezeev/solrg"
)

type SolrQueryRule struct {
	MatchQueryTriggers       []string          `json:"matchQueryTriggers"`
	ContainsAnyQueryTriggers []string          `json:"containsQueryTriggers"`
	MatchFilterQueryTriggers []string          `json:"matchFilterQueryTriggers"`
	ActReplaceQuery          string            `json:"actReplaceQuery"`
	ActAddFqs                []string          `json:"actAddFqs"`
	ActAddBqs                []string          `json:"actAddBqs"`
	Rule                     Rule              `json:"-"`
	SolrParams               *solrg.SolrParams `json:"-"`
	RulesSolrQuery           *solrg.SolrParams `json:"-"`
}

func (q *SolrQueryRule) Prepare(params interface{}) {
	q.Rule = NewRule("solr")
	// triggers
	q.SolrParams = params.(*solrg.SolrParams)
	for _, match := range q.MatchQueryTriggers {
		trigger := func(interface{}) bool {
			return (match == q.SolrParams.Q)
		}
		q.Rule.AddTrigger(trigger)
	}

	for _, subQ := range q.ContainsAnyQueryTriggers {
		trigger := func(interface{}) bool {
			return (strings.Contains(q.SolrParams.Q, subQ))
		}
		q.Rule.AddTrigger(trigger)
	}

	for _, match := range q.MatchFilterQueryTriggers {
		trigger := func(interface{}) bool {
			for _, fq := range q.SolrParams.Fq {
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
		f := func(params interface{}) error {
			q.SolrParams.Q = q.ActReplaceQuery
			return nil
		}
		q.Rule.AddAction(f)
	}
	if len(q.ActAddFqs) > 0 {
		f := func(params interface{}) error {
			for _, fq := range q.ActAddFqs {
				q.SolrParams.Fq = append(q.SolrParams.Fq, fq)
			}
			return nil
		}
		q.Rule.AddAction(f)
	}
	if len(q.ActAddBqs) > 0 {
		f := func(params interface{}) error {
			for _, bq := range q.ActAddBqs {
				q.SolrParams.Bq = q.SolrParams.Bq + " " + bq
			}
			return nil
		}
		q.Rule.AddAction(f)
	}
}

func (q *SolrQueryRule) Execute() {
	q.Rule.Execute(q.SolrParams)
}
