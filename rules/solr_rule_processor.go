package rules

import (
	"log"
	"strings"
	"time"

	"github.com/ezeev/solrg"
)

func ProcessSolrRules(userQuery *solrg.SolrParams) {
	// Begin Solr as Rules Engine
	rulesQuery := &solrg.SolrParams{Sort: "order_i desc"}
	var rq string
	rq = rq + "matchQueryTriggers_ss:\"" + userQuery.Q + "\" "
	rq = rq + "containsAnyQueryTriggers_txt:" + userQuery.Q + " "
	for _, fq := range userQuery.Fq {
		rq = rq + "containsFqs_ss:\"" + fq + "\""
	}

	//execute query
	sc, err := solrg.NewDirectSolrClient("172.104.9.135:8983/solr")
	if err != nil {
		log.Println(err)
	}
	rulesQuery.Q = rq
	resp, err := sc.Query("fastseer-staging.myshopify.com_rules", "select", rulesQuery, time.Second*10)
	if err != nil {
		log.Println(err)
	}

	log.Println(rulesQuery.Q)
	log.Printf("Matched %d rules", resp.Response.NumFound)

	// take actions
	for _, doc := range resp.Response.Docs {
		// query replacement
		if doc.HasField("actReplaceQuery_s") {
			userQuery.Q = doc.String("actReplaceQuery_s")
		}

		// add fqs
		if doc.HasField("actAddFqs_ss") {
			fqs, _ := doc.StringSlice("actAddFqs_ss")
			userQuery.Fq = append(userQuery.Fq, fqs...)
		}

		// add bqs
		if doc.HasField("actAddBqs_ss") {
			bqs, _ := doc.StringSlice("actAddBqs_ss")
			bqStr := strings.Join(bqs, " ")
			userQuery.Bq = userQuery.Bq + " " + bqStr
		}

	}

}
