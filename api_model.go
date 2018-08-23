package fastseer

import (
	"github.com/ezeev/fastseer/rules"
)

type MessageResponse struct {
	Message string `json:"message"`
}

type ErrorResponse struct {
	Message string `json:"message"`
	Error   error  `json:"error"`
}

type SearchRuleList []rules.Rule
