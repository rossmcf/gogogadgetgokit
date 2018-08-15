package endpoints

import (
	"context"

	"github.com/go-kit/kit/endpoint"
	"github.com/rossmcf/gogogadgetgokit/service"
)

type InspectRequest struct {
	S string
}
type InspectResponse struct {
	S string
}

func MakeInspectEndpoint(s service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(InspectRequest)
		string1 := s.Inspect(req.S)
		return InspectResponse{S: string1}, nil
	}
}

type Endpoints struct {
	Inspect endpoint.Endpoint
}
