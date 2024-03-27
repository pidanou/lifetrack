package api

import (
	"net/http"

	"github.com/labstack/echo/v4"

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
	e := echo.New()

	hydrationHandler := handlers.HydrationHandler{Datastore: s.datastore}
	e.GET("/hydration", hydrationHandler.HandleGetHydration)
	e.POST("/hydration", hydrationHandler.HandlePostHydration)

	sleepHandler := handlers.SleepHandler{Datastore: s.datastore}
	e.POST("/sleep", sleepHandler.HandlePostSleep)

	return http.ListenAndServe(s.port, e)
}
