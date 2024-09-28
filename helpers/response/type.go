package response

import "net/http"

type Map map[string]interface{}

type response struct {
	rw         http.ResponseWriter
	statusCode int
}
