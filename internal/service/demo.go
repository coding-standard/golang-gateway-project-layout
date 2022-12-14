package service

import (
	"context"
	generalv1 "github.com/coding-standard/golang-project-layout/api/general/v1"
	projectv1 "github.com/coding-standard/golang-project-layout/api/golang-project-layout/v1"
)

type DemoService struct {
	// This is generated by protoc
	projectv1.UnimplementedDemoServer
}

func NewDemoService() *DemoService {
	return &DemoService{}
}

func (s *DemoService) Demo(ctx context.Context, req *generalv1.DemoRequest) (*generalv1.DemoResponse, error) {
	return &generalv1.DemoResponse{
		Demo: &generalv1.Demo{
			Demo: req.Demo,
		},
	}, nil
}
