package handlers

import (
	"github.com/globalsign/mgo"
	pb "github.com/jessequinn/microgqlserver/srv/proto/vessel"
	rs "github.com/jessequinn/microgqlserver/srv/repositories"
	"golang.org/x/net/context"
)

// Our gRPC Service handler
type Service struct {
	Session *mgo.Session
}

func (s *Service) GetRepo() rs.Repository {
	return &rs.VesselRepository{s.Session.Clone()}
}

func (s *Service) FindAvailable(ctx context.Context, req *pb.Specification, res *pb.Response) error {
	repo := s.GetRepo()
	defer repo.Close()
	// Find the next available vessel
	vessel, err := repo.FindAvailable(req)
	if err != nil {
		return err
	}
	// Set the vessel as part of the response message type
	res.Vessel = vessel
	return nil
}

func (s *Service) Create(ctx context.Context, req *pb.Vessel, res *pb.Response) error {
	repo := s.GetRepo()
	defer repo.Close()
	if err := repo.Create(req); err != nil {
		return err
	}
	res.Vessel = req
	res.Created = true
	return nil
}
