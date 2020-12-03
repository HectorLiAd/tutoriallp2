package persona

import (
	"context"

	"github.com/go-kit/kit/endpoint"
)

type addPersonaRequest struct {
	PersonaID       int
	Nombre          string
	ApellidoPat     string
	ApellidoMat     string
	Genero          string
	Dni             string
	FechaNacimiento string
	Estado          string
}

func makeAddPersonEndpoint(s Service) endpoint.Endpoint {
	addPersonEndpoint := func(ctx context.Context, request interface{}) (interface{}, error) {
		
		return nil, nil
	}
	return addPersonEndpoint
}
