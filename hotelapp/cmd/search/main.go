package main

import (
	"flag"

	"github.com/nmikal98/cs499-hw3/hotelapp/internal/search"
	"github.com/nmikal98/cs499-hw3/hotelapp/pkg/tracing"
	log "github.com/sirupsen/logrus"
)

var (
	port       = flag.Int("port", 8082, "The service port")
	addr       = flag.String("addr", "0.0.0.0", "Address of the service")
	jaegeraddr = flag.String("jaeger", "jaeger:6831", "Jaeger address")
	geoaddr    = flag.String("geoaddr", "geo:8083", "Address of the geo service")
	rateaddr   = flag.String("rateaddr", "rate:8084", "Address of the rate service")
)

func main() {
	flag.Parse()

	tracer, err := tracing.NewTracer("geo", *jaegeraddr)
	if err != nil {
		log.Panicf("Got error while initializing jaeger agent: %v", err)
	}

	srv := search.NewSearch(*addr, *port, *geoaddr, *rateaddr, tracer)

	if err := srv.Run(); err != nil {
		log.Fatalf("run main error: %v", err)
	}
}
