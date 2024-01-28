package dependencyinjector

import (
	"github.com/mathcale/goexpert-rate-limiter-challenge/config"
	"github.com/mathcale/goexpert-rate-limiter-challenge/internal/infra/web"
	"github.com/mathcale/goexpert-rate-limiter-challenge/internal/infra/web/handlers"
	"github.com/mathcale/goexpert-rate-limiter-challenge/internal/pkg/logger"
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

	helloWebHandler := handlers.NewHelloWebHandler(responseHandler)

	webRouter := web.NewWebRouter(helloWebHandler)
	webServer := web.NewWebServer(di.Config.WebServerPort, logger.GetLogger(), webRouter.Build())

	return &Dependencies{
		Logger:          logger,
		ResponseHandler: responseHandler,
		HelloWebHandler: helloWebHandler,
		WebServer:       webServer,
	}
}
