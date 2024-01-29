package ratelimiter

import (
	"github.com/rs/zerolog"

	ratelimiter_datastore "github.com/mathcale/goexpert-rate-limiter-challenge/internal/infra/database/ratelimiter"
	"github.com/mathcale/goexpert-rate-limiter-challenge/internal/pkg/logger"
)

type RateLimiterInterface interface {
	Check() bool
}

type RateLimiter struct {
	Logger              zerolog.Logger
	Datastore           ratelimiter_datastore.RateLimiterDatastoreInterface
	MaxRequestsPerIP    int
	MaxRequestsPerToken int
	TimeWindowMillis    int
	CooldownTimeMillis  int
}

func NewRateLimiter(
	logger logger.LoggerInterface,
	datastore ratelimiter_datastore.RateLimiterDatastoreInterface,
	ipMaxReqs int,
	tokenMaxReqs int,
	timeWindow int,
	cooldownTime int,
) *RateLimiter {
	return &RateLimiter{
		Logger:              logger.GetLogger(),
		Datastore:           datastore,
		MaxRequestsPerIP:    ipMaxReqs,
		MaxRequestsPerToken: tokenMaxReqs,
		TimeWindowMillis:    timeWindow,
		CooldownTimeMillis:  cooldownTime,
	}
}

func (rl *RateLimiter) Check() bool {
	return true
}
