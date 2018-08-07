package rules

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

