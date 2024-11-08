package response

import (
	"encoding/json"
	"net/http"
)

func Response(w http.ResponseWriter) *response {
	return &response{
		rw:         w,
		statusCode: 200,
	}
}

func (resp *response) StatusCode(code int) *response {
	resp.statusCode = code
	return resp
}

func (resp *response) Json(data any) {
	resp.rw.Header().Set("Content-Type", "Application/json")
	resp.rw.WriteHeader(resp.statusCode)
	json.NewEncoder(resp.rw).Encode(data)
}

func (resp *response) Success(msg string, data any) {
	resp.Json(Map{
		"message": msg,
		"data":    data,
		"status":  "success",
	})
}

func (resp *response) Error(msg string, errors any) {
	resp.Json(Map{
		"message": msg,
		"errors":  errors,
		"status":  "error",
	})
}
