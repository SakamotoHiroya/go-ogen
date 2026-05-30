package service

import (
	"context"
	"math/rand/v2"

	"github.com/SakamotoHiroya/go-ogen/internal/gen/greetingapi"
)

type GreetingService struct {
	greetings []string
}

func NewGreetingService() *GreetingService {
	return &GreetingService{
		greetings: []string{
			"Hello",
			"Hi",
			"Hey",
			"Good morning",
			"Good afternoon",
			"Good evening",
		},
	}
}

func (s *GreetingService) CreateGreeting(ctx context.Context, req *greetingapi.CreateGreetingRequest) (*greetingapi.GreetingResponse, error) {
	return &greetingapi.GreetingResponse{
		Message: "Hello " + req.Name,
	}, nil
}

func (s *GreetingService) GetRandomGreeting(ctx context.Context) (*greetingapi.GreetingResponse, error) {
	return &greetingapi.GreetingResponse{
		Message: s.greetings[rand.IntN(len(s.greetings))],
	}, nil
}
