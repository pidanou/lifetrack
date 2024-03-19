package api

import (
	"net/http"

	"github.com/pidanou/lifetrack/internal/datastore"
	"github.com/pidanou/lifetrack/internal/handlers"
)

type Server struct {
	port      string
	datastore datastore.Datastore
}

func NewServer(datastore datastore.Datastore, port string) *Server {
	return &Server{port: port, datastore: datastore}
}

func (s *Server) Start() error {
	routes := s.InitRoutes()
	return http.ListenAndServe(s.port, routes)
}

func (s *Server) InitRoutes() *http.ServeMux {
	mux := http.NewServeMux()

	hydrationHandler := handlers.HydrationHandler{Datastore: s.datastore}
	mux.HandleFunc("POST /hydration", hydrationHandler.HandlePostHydration)
	mux.HandleFunc("GET /hydration", hydrationHandler.HandleGetHydration)

	sleepHandler := handlers.SleepHandler{Datastore: s.datastore}
	mux.HandleFunc("POST /sleep", sleepHandler.HandlePostSleep)

	return mux
}
