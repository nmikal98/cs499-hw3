package profile

import (
	"context"
	"fmt"
	"log"
	"net"
	"time"

	"github.com/grpc-ecosystem/grpc-opentracing/go/otgrpc"
	pb "github.com/nmikal98/cs499-hw3/hotelapp/internal/profile/proto"
	"github.com/opentracing/opentracing-go"
	"google.golang.org/grpc"
	"google.golang.org/grpc/keepalive"
)

// Profile implements the profile service
type Profile struct {
	port      int
	addr      string
	dbsession *DatabaseSession
	tracer    opentracing.Tracer
}

// NewProfile returns a new Profile service
func NewProfile(a string, p int, db *DatabaseSession, tr opentracing.Tracer) *Profile {
	return &Profile{
		addr:      a,
		port:      p,
		dbsession: db,
		tracer:    tr,
	}
}

// Run starts the server
func (s *Profile) Run() error {
	// TODO: Implement me

	if s.port == 0 {
		return fmt.Errorf("server port must be set")
	}

	opts := []grpc.ServerOption{
		grpc.KeepaliveParams(keepalive.ServerParameters{
			Timeout: 120 * time.Second,
		}),
		grpc.KeepaliveEnforcementPolicy(keepalive.EnforcementPolicy{
			PermitWithoutStream: true,
		}),
		grpc.UnaryInterceptor(
			otgrpc.OpenTracingServerInterceptor(s.tracer),
		),
	}

	srv := grpc.NewServer(opts...)
	pb.RegisterProfileServer(srv, s)

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", s.port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	return srv.Serve(lis)
}

// GetProfiles returns hotel profiles for requested IDs
func (s *Profile) GetProfiles(ctx context.Context, req *pb.Request) (*pb.Result, error) {
	// TODO: Implement me
	var err error
	res := new(*pb.Result)
	res.Hotels, err = s.dbsession.GetProfiles(req.HotelIds)
	return res, err
}
