package rules

type SolrRule struct {
	Triggered bool
	//Triggers  []func(*solrg.SolrParams) bool
	Triggers []func(interface{}) bool
	Actions  []func(interface{}) error
}

//func (s *SolrRule) AddTrigger(t func(*solrg.SolrParams) bool) {
func (s *SolrRule) AddTrigger(t func(interface{}) bool) {
	s.Triggers = append(s.Triggers, t)
}

//func (s *SolrRule) AddAction(a func(*solrg.SolrParams) error) {
func (s *SolrRule) AddAction(a func(interface{}) error) {
	s.Actions = append(s.Actions, a)
}

//func (s *SolrRule) Execute(params *solrg.SolrParams) error {

func (s *SolrRule) Execute(params interface{}) error {

	//trigger?
	for _, t := range s.Triggers {
		if t(params) {
			s.Triggered = true
		}
	}
	if s.Triggered {
		//run actions
		for _, a := range s.Actions {
			err := a(params)
			if err != nil {
				return err
			}
		}
	}
	return nil
}
