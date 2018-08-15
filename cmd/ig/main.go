package main

import (
	"net"
	stdhttp "net/http"
	"os"

	"github.com/go-kit/kit/log"
	"github.com/rossmcf/gogogadgetgokit/endpoints"
	"github.com/rossmcf/gogogadgetgokit/http"
	"github.com/rossmcf/gogogadgetgokit/service"
)

func main() {
	// Service
	logger := log.NewLogfmtLogger(os.Stdout)
	svc := service.Service{}

	// Endpoint
	inspect := endpoints.MakeInspectEndpoint(svc)
	inspect = endpoints.LoggingMiddleware(logger)(inspect)
	ep := endpoints.Endpoints{Inspect: inspect}

	// Transport
	httpHandler := http.NewHTTPHandler(ep)
	httpListener, err := net.Listen("tcp", ":8080")
	if err != nil {
		panic(err)
	}
	stdhttp.Serve(httpListener, httpHandler)
}
