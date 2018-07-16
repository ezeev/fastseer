package util

import (
	"fmt"
)

func ValidateOptionsMap(themap map[string]string, requiredKeys ...string) error {
	for _, v := range requiredKeys {
		if themap[v] == "" {
			return fmt.Errorf("Missing required options: %s", v)
		}
	}
	return nil
}
