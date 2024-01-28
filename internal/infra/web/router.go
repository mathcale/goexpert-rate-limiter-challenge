package web

import "github.com/mathcale/goexpert-rate-limiter-challenge/internal/infra/web/handlers"

type WebRouterInterface interface {
	Build() []RouteHandler
}

type WebRouter struct {
	WebClimateHandler handlers.HelloWebHandlerInterface
}

func NewWebRouter(webClimateHandler handlers.HelloWebHandlerInterface) *WebRouter {
	return &WebRouter{
		WebClimateHandler: webClimateHandler,
	}
}

func (wr *WebRouter) Build() []RouteHandler {
	return []RouteHandler{
		{
			Path:        "/",
			Method:      "GET",
			HandlerFunc: wr.WebClimateHandler.SayHello,
		},
	}
}
