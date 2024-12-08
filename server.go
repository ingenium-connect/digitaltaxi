package main

import (
	"context"
	"strconv"

	"github.com/ingenium-connect/digitaltaxi/pkg/digitaltaxi/application/common"
	"github.com/ingenium-connect/digitaltaxi/pkg/digitaltaxi/application/common/helpers"
	"github.com/ingenium-connect/digitaltaxi/pkg/digitaltaxi/presentation"
)

const waitSeconds = 30

func main() {
	ctx := context.Background()

	port, err := strconv.Atoi(helpers.MustGetEnvVar(common.PortEnvVarName))
	if err != nil {
		helpers.LogStartupError(ctx, err)
	}

	presentation.PrepareServer(ctx, port)
}
