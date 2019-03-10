package endpoints

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/gilmoreg/seibiki/internal/service"
	"github.com/go-kit/kit/endpoint"
	httptransport "github.com/go-kit/kit/transport/http"
	"go.uber.org/zap"
)

// Handler - new http.Handler
func Handler(svc service.LookupService, logger *zap.SugaredLogger) *httptransport.Server {
	return httptransport.NewServer(
		createEndpoint(svc),
		decodeQueryRequest,
		encodeResponse,
	)
}

func createEndpoint(svc service.LookupService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(queryRequest).Query
		return svc.Lookup(req), nil
	}
}

func decodeQueryRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var query queryRequest
	body, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(body, &query)
	if err != nil {
		return nil, err
	}
	return query, nil
}

func encodeResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "OPTIONS, POST")
	return json.NewEncoder(w).Encode(response)
}

type queryRequest struct {
	Query string `json:"query"`
}
