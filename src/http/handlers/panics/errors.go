package panics

import (
	"awesomeProject/src/interfaces/errors"
	"net/http"
)

func NewBadRequest() errors.PanicResponse {
	return &panicResponse{
		code: http.StatusBadRequest,
		body: httpMessage{false, "Bad request"},
	}
}

type panicResponse struct {
	code int
	body any
}

func (r *panicResponse) Code() int {
	return r.code
}

func (r *panicResponse) Object() any {
	return r.body
}

type httpMessage struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
}
