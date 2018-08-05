package fastseer

import (
	"encoding/json"
	"net/http"
)

func SendErrorResponse(w http.ResponseWriter, r *http.Request, message string, err error) {
	w.WriteHeader(500)
	json.NewEncoder(w).Encode(ErrorResponse{Message: message, Error: err})
}
