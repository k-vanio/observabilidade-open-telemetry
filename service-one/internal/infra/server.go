package infra

import (
	"net/http"
	"time"

	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"
	"github.com/k-vanio/observabilidade-open-telemetry/shared"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/propagation"
)

type Server struct {
	Config *shared.Config
}

func NewServer(config *shared.Config) *Server {
	return &Server{
		Config: config,
	}
}

func (s *Server) CreateServer() *chi.Mux {
	router := chi.NewRouter()

	router.Use(middleware.RequestID)
	router.Use(middleware.RealIP)
	router.Use(middleware.Recoverer)
	router.Use(middleware.Logger)
	router.Use(middleware.Timeout(60 * time.Second))

	// promhttp
	router.Handle("/metrics", promhttp.Handler())

	// request handler
	router.Get("/", s.HandleRequest)

	return router
}

func (s *Server) HandleRequest(w http.ResponseWriter, r *http.Request) {
	carrier := propagation.HeaderCarrier(r.Header)
	ctx := r.Context()
	ctx = otel.GetTextMapPropagator().Extract(ctx, carrier)
	ctx, span := s.Config.OTELTracer.Start(ctx, s.Config.RequestNameOTEL)
	defer span.End()

	req, _ := http.NewRequest(s.Config.ExternalCallMethod, s.Config.ExternalCallURL, nil)

	otel.GetTextMapPropagator().Inject(ctx, propagation.HeaderCarrier(req.Header))

	http.DefaultClient.Do(req)

	http.ResponseWriter(w).Write([]byte(s.Config.Content))
}
