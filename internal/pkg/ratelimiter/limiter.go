package ratelimiter

import (
	"github.com/mathcale/goexpert-rate-limiter-challenge/internal/pkg/logger"
	"github.com/rs/zerolog"
)

type RateLimiterInterface interface {
	Check() bool
}

type RateLimiter struct {
	Logger              zerolog.Logger
	MaxRequestsPerIP    int
	MaxRequestsPerToken int
	TimeWindowMillis    int
	CooldownTimeMillis  int
}

func NewRateLimiter(
	logger logger.LoggerInterface,
	ipMaxReqs int,
	tokenMaxReqs int,
	timeWindow int,
	cooldownTime int,
) *RateLimiter {
	return &RateLimiter{
		Logger:              logger.GetLogger(),
		MaxRequestsPerIP:    ipMaxReqs,
		MaxRequestsPerToken: tokenMaxReqs,
		TimeWindowMillis:    timeWindow,
		CooldownTimeMillis:  cooldownTime,
	}
}

func (rl *RateLimiter) Check() bool {
	return true
}
