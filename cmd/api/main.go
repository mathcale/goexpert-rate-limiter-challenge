package main

import (
	"github.com/mathcale/goexpert-rate-limiter-challenge/config"
	"github.com/mathcale/goexpert-rate-limiter-challenge/internal/pkg/dependencyinjector"
)

func main() {
	configs, err := config.Load(".")
	if err != nil {
		panic(err)
	}

	di := dependencyinjector.NewDependencyInjector(configs)
	deps := di.Inject()

	deps.WebServer.Start()
}
