package web

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	chizero "github.com/ironstar-io/chizerolog"
	"github.com/rs/zerolog"
)

type WebServerInterface interface {
	Start()
}

type RouteHandler struct {
	Path        string
	Method      string
	HandlerFunc http.HandlerFunc
}

type Middleware struct {
	Name    string
	Handler func(next http.Handler) http.Handler
}

type WebServer struct {
	Router        chi.Router
	Handlers      []RouteHandler
	Middlewares   []Middleware
	WebServerPort string
	Logger        zerolog.Logger
}

func NewWebServer(
	serverPort string,
	logger zerolog.Logger,
	handlers []RouteHandler,
	middlewares []Middleware,
) *WebServer {
	return &WebServer{
		Router:        chi.NewRouter(),
		Handlers:      handlers,
		Middlewares:   middlewares,
		WebServerPort: serverPort,
		Logger:        logger,
	}
}

func (s *WebServer) Start() {
	s.Router.Use(chizero.LoggerMiddleware(&s.Logger))

	for _, m := range s.Middlewares {
		s.Logger.Debug().Msgf("Registering middleware [%s]", m.Name)
		s.Router.Use(m.Handler)
	}

	for _, h := range s.Handlers {
		s.Logger.Debug().Msgf("Registering route %s %s", h.Method, h.Path)
		s.Router.MethodFunc(h.Method, h.Path, h.HandlerFunc)
	}

	s.Logger.Info().Msgf("Starting webserver on port %s", s.WebServerPort)

	http.ListenAndServe(s.WebServerPort, s.Router)
}
