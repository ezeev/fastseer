package fastseer

type MessageResponse struct {
	Message string `json:"message"`
}

type ErrorResponse struct {
	Message string `json:"message"`
	Error   error  `json:"error"`
}
