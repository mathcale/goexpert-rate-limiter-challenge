package middlewares

import (
	"net/http"

	"github.com/mathcale/goexpert-rate-limiter-challenge/internal/pkg/logger"
	"github.com/mathcale/goexpert-rate-limiter-challenge/internal/pkg/ratelimiter"
	"github.com/rs/zerolog"
)

type RateLimiterMiddlewareInterface interface {
	Handle(next http.Handler) http.Handler
}

type RateLimiterMiddleware struct {
	Logger  zerolog.Logger
	Limiter ratelimiter.RateLimiterInterface
}

func NewRateLimiterMiddleware(
	logger logger.LoggerInterface,
	limiter ratelimiter.RateLimiterInterface,
) *RateLimiterMiddleware {
	return &RateLimiterMiddleware{
		Logger:  logger.GetLogger(),
		Limiter: limiter,
	}
}

func (rlm *RateLimiterMiddleware) Handle(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		rlm.Logger.Info().Msg("Hello from RateLimiterMiddleware")

		next.ServeHTTP(w, r)
	})
}
