package fastseer

import (
	"bytes"
	"encoding/json"
	"net/http"
	"testing"

	"github.com/ezeev/fastseer/rules"
)

func TestRulesCRUD(t *testing.T) {

	server, _ := NewServer("config-stage.yaml")
	go server.Start()
	// wait until the server is ready
	server.ServerReady()

	// create a rule
	rule := rules.SearchRule{}
	rule.ID = "test01"
	rule.MatchQueryTriggersSs = []string{"test"}
	rule.ActReplaceQueryS = "i was the replacement"

	url := "http://localhost:8082" + apiV1HandlePostRules + "?" + testAuthParams

	rulesList := []rules.SearchRule{rule}
	b, _ := json.Marshal(rulesList)
	resp, err := http.Post(url, "application/json", bytes.NewBuffer(b))
	if err != nil {
		t.Error(err)
	}
	if resp.StatusCode != 200 {
		t.Fail()
	}
	resp.Body.Close()
	//now get it
	geturl := url + "&q=id:test01"
	resp, err = http.Get(geturl)
	if err != nil {
		t.Error(err)
	}
	if resp.StatusCode != 200 {
		t.Fail()
	}

	var newRules []rules.SearchRule

	decoder := json.NewDecoder(resp.Body)
	decoder.Decode(&newRules)
	if newRules[0].ID != "test01" {
		t.Error("Did not get back expected ID!")
	}
	t.Log(newRules)
	resp.Body.Close()

	//delete
	delurl := url + "&id=test01"
	req, _ := http.NewRequest("DELETE", delurl, nil)
	cli := http.Client{}
	resp, err = cli.Do(req)
	if err != nil {
		t.Error(err)
	}
	if resp.StatusCode != 200 {
		t.Fail()
	}
	t.Log(resp)
	resp.Body.Close()

	server.Shutdown()

}
