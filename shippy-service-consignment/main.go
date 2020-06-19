package main

import (
	"context"
	"log"

	"github.com/micro/go-micro/v2"
	pb "github.com/ryanyogan/shippy/shippy-service-consignment/proto/consignment"
)

type repository interface {
	Create(*pb.Consignment) (*pb.Consignment, error)
	GetAll() []*pb.Consignment
}

// Repository - Memory repo, simulats a datastore
type Repository struct {
	consignments []*pb.Consignment
}

// Create will create a consignment and return it, or return an error
func (repo *Repository) Create(consignment *pb.Consignment) (*pb.Consignment, error) {
	updated := append(repo.consignments, consignment)
	repo.consignments = updated
	return consignment, nil
}

// GetAll consignments
func (repo *Repository) GetAll() []*pb.Consignment {
	return repo.consignments
}

// Service should implement all methods to satisfy the service.
type consignmentService struct {
	repo repository
}

func (s *consignmentService) CreateConsignment(ctx context.Context, req *pb.Consignment, res *pb.Response) error {
	consignment, err := s.repo.Create(req)
	if err != nil {
		return err
	}

	res.Created = true
	res.Consignment = consignment
	return nil
}

func (s *consignmentService) GetConsignments(ctx context.Context, req *pb.GetRequest, res *pb.Response) error {
	consignments := s.repo.GetAll()
	res.Consignments = consignments
	return nil
}

func main() {
	repo := &Repository{}

	svc := micro.NewService(
		micro.Name("shippy.service.consignment"),
	)
	svc.Init()

	if err := pb.RegisterShippingServiceHandler(svc.Server(), &consignmentService{repo}); err != nil {
		log.Panic(err)
	}

	if err := svc.Run(); err != nil {
		log.Panic(err)
	}
}
