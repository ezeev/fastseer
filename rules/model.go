package rules

type SearchRule struct {
	ID                          string   `json:"id"`
	MatchQueryTriggersSs        []string `json:"matchQueryTriggers_ss"`
	ContainsAnyQueryTriggersTxt []string `json:"containsAnyQueryTriggers_txt"`
	ContainsFqsSs               []string `json:"containsFqs_ss"`
	ActReplaceQueryS            string   `json:"actReplaceQuery_s"`
	Name                        string   `json:"name_s"`
	ActAddFqsSs                 []string `json:"actAddFqs_ss"`
	ActAddBqsSs                 []string `json:"actAddBqs_ss"`
	ActAddFacetFieldsSs         []string `json:"actAddFacetFields_ss"`
	TagsSs                      []string `json:"tags_ss"`
	OrderI                      int      `json:"order_i"`
}
