package main

import (
	"context"
	"log"
	"strconv"

	"github.com/ingenium-connect/digitaltaxi/pkg/digitaltaxi/application/common"
	"github.com/ingenium-connect/digitaltaxi/pkg/digitaltaxi/application/common/helpers"
	"github.com/ingenium-connect/digitaltaxi/pkg/digitaltaxi/presentation"
)

func main() {
	ctx := context.Background()

	port, err := strconv.Atoi(helpers.MustGetEnvVar(common.PortEnvVarName))
	if err != nil {
		log.Panicf("Could not get environment variable %s", helpers.MustGetEnvVar(common.PortEnvVarName))
	}

	presentation.PrepareServer(ctx, port)
}
