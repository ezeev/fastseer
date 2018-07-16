package util

import "testing"

func TestValidateOptionsMap(t *testing.T) {

	opts := map[string]string{"opt1": "val1", "opt2": "val2", "opt3": "val3"}
	err := ValidateOptionsMap(opts, "opt1", "opt2")
	if err != nil {
		t.Error(err)
		t.Fail()
	} else {
		t.Log("no error as expected")
	}

	err = ValidateOptionsMap(opts, "test")
	if err == nil {
		t.Error("expected error")
		t.Fail()
	} else {
		t.Log(err)
	}

}
