package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/google/uuid"
	goahttp "goa.design/goa/v3/http"

	// Use gen prefix for generated packages
	concerts "goa.design/examples/concerts"
	genconcerts "goa.design/examples/concerts/gen/concerts"
	genhttp "goa.design/examples/concerts/gen/http/concerts/server"
)

// ConcertsService implements the genconcerts.Service interface
type ConcertsService struct {
	concerts []*genconcerts.Concert // In-memory storage
}

// List upcoming concerts with optional pagination.
func (m *ConcertsService) List(ctx context.Context, p *genconcerts.ListPayload) ([]*genconcerts.Concert, error) {
	start := (p.Page - 1) * p.Limit
	end := start + p.Limit
	if end > len(m.concerts) {
		end = len(m.concerts)
	}
	return m.concerts[start:end], nil
}

// Create a new concerts entry.
func (m *ConcertsService) Create(ctx context.Context, p *genconcerts.ConcertPayload) (*genconcerts.Concert, error) {
	newConcert := &genconcerts.Concert{
		ID:     uuid.New().String(),
		Artist: p.Artist,
		Date:   p.Date,
		Venue:  p.Venue,
		Price:  p.Price,
	}
	m.concerts = append(m.concerts, newConcert)
	return newConcert, nil
}

// Get a single concert by ID.
func (m *ConcertsService) Show(ctx context.Context, p *genconcerts.ShowPayload) (*genconcerts.Concert, error) {
	for _, concert := range m.concerts {
		if concert.ID == p.ConcertID {
			return concert, nil
		}
	}
	// Use designed error
	return nil, genconcerts.MakeNotFound(fmt.Errorf("concert not found: %s", p.ConcertID))
}

// Update an existing concert by ID.
func (m *ConcertsService) Update(ctx context.Context, p *genconcerts.UpdatePayload) (*genconcerts.Concert, error) {
	for i, concert := range m.concerts {
		if concert.ID == p.ConcertID {
			if p.Artist != nil {
				concert.Artist = *p.Artist
			}
			if p.Date != nil {
				concert.Date = *p.Date
			}
			if p.Venue != nil {
				concert.Venue = *p.Venue
			}
			if p.Price != nil {
				concert.Price = *p.Price
			}
			m.concerts[i] = concert
			return concert, nil
		}
	}
	return nil, genconcerts.MakeNotFound(fmt.Errorf("concert not found: %s", p.ConcertID))
}

// Remove a concert from the system by ID.
func (m *ConcertsService) Delete(ctx context.Context, p *genconcerts.DeletePayload) error {
	for i, concert := range m.concerts {
		if concert.ID == p.ConcertID {
			m.concerts = append(m.concerts[:i], m.concerts[i+1:]...)
			return nil
		}
	}
	return genconcerts.MakeNotFound(fmt.Errorf("concert not found: %s", p.ConcertID))
}

// main instantiates the service and starts the HTTP server.
func main() {
	// Instantiate the service
	svc := &ConcertsService{}

	// Wrap it in the generated endpoints
	endpoints := genconcerts.NewEndpoints(svc)

	// Build an HTTP handler
	mux := goahttp.NewMuxer()
	requestDecoder := goahttp.RequestDecoder
	responseEncoder := goahttp.ResponseEncoder

	// Use the embedded OpenAPI files from the top-level package
	openAPIFS := concerts.OpenAPIFileSystem()

	handler := genhttp.New(endpoints, mux, requestDecoder, responseEncoder, nil, nil, openAPIFS, openAPIFS)

	// Mount the handler on the mux
	genhttp.Mount(mux, handler)

	// Create a new HTTP server
	port := "8080"
	server := &http.Server{Addr: ":" + port, Handler: mux}

	// Log the supported routes
	for _, mount := range handler.Mounts {
		log.Printf("%q mounted on %s %s", mount.Method, mount.Verb, mount.Pattern)
	}

	// Start the server (this will block the execution)
	log.Printf("Starting concerts service on :%s", port)
	if err := server.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}
