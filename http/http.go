package http

import "context"
import "encoding/json"

import "net/http"

import httptransport "github.com/go-kit/kit/transport/http"
import "github.com/rossmcf/gogogadgetgokit/endpoints"

func NewHTTPHandler(endpoints endpoints.Endpoints) http.Handler {
	m := http.NewServeMux()
	m.Handle("/inspect", httptransport.NewServer(endpoints.Inspect, DecodeInspectRequest, EncodeInspectResponse))
	return m
}
func DecodeInspectRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var req endpoints.InspectRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	return req, err
}
func EncodeInspectResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	return json.NewEncoder(w).Encode(response)
}
