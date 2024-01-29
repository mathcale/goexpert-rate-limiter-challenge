package dependencyinjector

import (
	"github.com/mathcale/goexpert-rate-limiter-challenge/config"
	"github.com/mathcale/goexpert-rate-limiter-challenge/internal/infra/web"
	"github.com/mathcale/goexpert-rate-limiter-challenge/internal/infra/web/handlers"
	"github.com/mathcale/goexpert-rate-limiter-challenge/internal/infra/web/middlewares"
	"github.com/mathcale/goexpert-rate-limiter-challenge/internal/pkg/logger"
	"github.com/mathcale/goexpert-rate-limiter-challenge/internal/pkg/ratelimiter"
	"github.com/mathcale/goexpert-rate-limiter-challenge/internal/pkg/responsehandler"
)

type DependencyInjectorInterface interface {
	Inject() *Dependencies
}

type DependencyInjector struct {
	Config *config.Conf
}

type Dependencies struct {
	Logger          logger.LoggerInterface
	ResponseHandler responsehandler.WebResponseHandlerInterface
	HelloWebHandler handlers.HelloWebHandlerInterface
	WebServer       web.WebServerInterface
}

func NewDependencyInjector(c *config.Conf) *DependencyInjector {
	return &DependencyInjector{
		Config: c,
	}
}

func (di *DependencyInjector) Inject() *Dependencies {
	logger := logger.NewLogger(di.Config.LogLevel)
	responseHandler := responsehandler.NewWebResponseHandler()

	limiter := ratelimiter.NewRateLimiter(
		logger,
		di.Config.RateLimiterIPMaxRequests,
		di.Config.RateLimiterTokenMaxRequests,
		di.Config.RateLimiterTimeWindowMilliseconds,
		di.Config.RateLimiterCooldownTimeMilliseconds,
	)

	helloWebHandler := handlers.NewHelloWebHandler(responseHandler)
	rateLimiterMiddleware := middlewares.NewRateLimiterMiddleware(logger, limiter)

	webRouter := web.NewWebRouter(helloWebHandler, rateLimiterMiddleware)
	webServer := web.NewWebServer(
		di.Config.WebServerPort,
		logger.GetLogger(),
		webRouter.Build(),
		webRouter.BuildMiddlewares(),
	)

	return &Dependencies{
		Logger:          logger,
		ResponseHandler: responseHandler,
		HelloWebHandler: helloWebHandler,
		WebServer:       webServer,
	}
}
