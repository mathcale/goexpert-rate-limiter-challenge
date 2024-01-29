package ratelimiter

import (
	"errors"
	"fmt"

	"github.com/rs/zerolog"

	"github.com/mathcale/goexpert-rate-limiter-challenge/config"
)

var datastoreFactories = make(map[string]RateLimiterDatastore)

type RateLimiterDatastore func(cfg config.Conf, logger zerolog.Logger) (RateLimiterDatastoreInterface, error)

type RateLimiterDatastoreFactory struct {
	Config config.Conf
	Logger zerolog.Logger
}

func NewRateLimiterDatastoreFactory(cfg config.Conf, logger zerolog.Logger) *RateLimiterDatastoreFactory {
	register("redis", NewRedisRateLimiterDatastore, logger)

	return &RateLimiterDatastoreFactory{
		Config: cfg,
		Logger: logger,
	}
}

func (r *RateLimiterDatastoreFactory) Create() (RateLimiterDatastoreInterface, error) {
	factory, ok := datastoreFactories[r.Config.RateLimiterDatastore]
	if !ok {
		return nil, errors.New(fmt.Sprintf("invalid rate limiter datastore [%s]", r.Config.RateLimiterDatastore))
	}

	return factory(r.Config, r.Logger)
}

func register(name string, factory RateLimiterDatastore, logger zerolog.Logger) {
	if factory == nil {
		logger.Debug().Msgf("[RateLimiterFactory]: factory [%s] does not exist", name)
	}

	if _, registered := datastoreFactories[name]; registered {
		logger.Debug().Msgf("[RateLimiterFactory]: factory [%s] already registered and will be ignored", name)
	}

	datastoreFactories[name] = factory
}
