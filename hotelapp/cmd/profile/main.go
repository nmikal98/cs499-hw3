package main

import (
	"flag"

	"github.com/nmikal98/cs499-hw3/hotelapp/internal/profile"
	"github.com/nmikal98/cs499-hw3/hotelapp/pkg/tracing"
	log "github.com/sirupsen/logrus"
)

var (
	port        = flag.Int("port", 8081, "The service port")
	addr        = flag.String("addr", "0.0.0.0", "Address of the service")
	jaegeraddr  = flag.String("jaeger", "jaeger:6831", "Jaeger address")
	jsonDataDir = flag.String("jsondata", "data/medium", "Directory containing json data files")
)

func main() {
	flag.Parse()

	tracer, err := tracing.NewTracer("profile", *jaegeraddr)
	if err != nil {
		log.Panicf("Got error while initializing jaeger agent: %v", err)
	}

	// Initialize Database
	pdb := initializeProfileDatabase()

	srv := profile.NewProfile(*addr, *port, pdb, tracer)

	if err := srv.Run(); err != nil {
		log.Fatalf("run main error: %v", err)
	}
}
