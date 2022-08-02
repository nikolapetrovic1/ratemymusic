package endpoint

import (
	"context"
	endpoint "github.com/go-kit/kit/endpoint"
	model "github.com/nikolapetrovic1/ratemymusic/songs/pkg/model"
	service "github.com/nikolapetrovic1/ratemymusic/songs/pkg/service"
)

// GetRequest collects the request parameters for the Get method.
type GetRequest struct{}

// GetResponse collects the response parameters for the Get method.
type GetResponse struct {
	M0 model.Song `json:"m0"`
	E1 error      `json:"e1"`
}

// MakeGetEndpoint returns an endpoint that invokes Get on the service.
func MakeGetEndpoint(s service.SongsService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		m0, e1 := s.Get(ctx)
		return GetResponse{
			E1: e1,
			M0: m0,
		}, nil
	}
}

// Failed implements Failer.
func (r GetResponse) Failed() error {
	return r.E1
}

// GetAllRequest collects the request parameters for the GetAll method.
type GetAllRequest struct{}

// GetAllResponse collects the response parameters for the GetAll method.
type GetAllResponse struct {
	M0 []model.Song `json:"m0"`
	E1 error        `json:"e1"`
}

// MakeGetAllEndpoint returns an endpoint that invokes GetAll on the service.
func MakeGetAllEndpoint(s service.SongsService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		m0, e1 := s.GetAll(ctx)
		return GetAllResponse{
			E1: e1,
			M0: m0,
		}, nil
	}
}

// Failed implements Failer.
func (r GetAllResponse) Failed() error {
	return r.E1
}

// AddRequest collects the request parameters for the Add method.
type AddRequest struct {
	Song model.Song `json:"song"`
}

// AddResponse collects the response parameters for the Add method.
type AddResponse struct {
	M0 model.Song `json:"m0"`
	E1 error      `json:"e1"`
}

// MakeAddEndpoint returns an endpoint that invokes Add on the service.
func MakeAddEndpoint(s service.SongsService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(AddRequest)
		m0, e1 := s.Add(ctx, req.Song)
		return AddResponse{
			E1: e1,
			M0: m0,
		}, nil
	}
}

// Failed implements Failer.
func (r AddResponse) Failed() error {
	return r.E1
}

// DeleteRequest collects the request parameters for the Delete method.
type DeleteRequest struct {
	Id string `json:"id"`
}

// DeleteResponse collects the response parameters for the Delete method.
type DeleteResponse struct {
	E0 error `json:"e0"`
}

// MakeDeleteEndpoint returns an endpoint that invokes Delete on the service.
func MakeDeleteEndpoint(s service.SongsService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(DeleteRequest)
		e0 := s.Delete(ctx, req.Id)
		return DeleteResponse{E0: e0}, nil
	}
}

// Failed implements Failer.
func (r DeleteResponse) Failed() error {
	return r.E0
}

// Failure is an interface that should be implemented by response types.
// Response encoders can check if responses are Failer, and if so they've
// failed, and if so encode them using a separate write path based on the error.
type Failure interface {
	Failed() error
}

// Get implements Service. Primarily useful in a client.
func (e Endpoints) Get(ctx context.Context) (m0 model.Song, e1 error) {
	request := GetRequest{}
	response, err := e.GetEndpoint(ctx, request)
	if err != nil {
		return
	}
	return response.(GetResponse).M0, response.(GetResponse).E1
}

// GetAll implements Service. Primarily useful in a client.
func (e Endpoints) GetAll(ctx context.Context) (m0 []model.Song, e1 error) {
	request := GetAllRequest{}
	response, err := e.GetAllEndpoint(ctx, request)
	if err != nil {
		return
	}
	return response.(GetAllResponse).M0, response.(GetAllResponse).E1
}

// Add implements Service. Primarily useful in a client.
func (e Endpoints) Add(ctx context.Context, song model.Song) (m0 model.Song, e1 error) {
	request := AddRequest{Song: song}
	response, err := e.AddEndpoint(ctx, request)
	if err != nil {
		return
	}
	return response.(AddResponse).M0, response.(AddResponse).E1
}

// Delete implements Service. Primarily useful in a client.
func (e Endpoints) Delete(ctx context.Context, id string) (e0 error) {
	request := DeleteRequest{Id: id}
	response, err := e.DeleteEndpoint(ctx, request)
	if err != nil {
		return
	}
	return response.(DeleteResponse).E0
}
