package main

import (
	"backend/common/config"
	"backend/common/grpc"
	"backend/procedure/driver/metadata"
	"go.uber.org/fx"
)

func main() {
	fx.New(
		fx.Provide(
			config.Init,
			grpc.Init,
		),
		fx.Invoke(
			metadata.Init,
		),
	).Run()
}
