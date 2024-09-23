package main

import (
	"backend/common/config"
	"backend/common/fiber"
	"backend/common/grpc"
	"backend/endpoint"
	"backend/endpoint/get"
	"backend/endpoint/system"
	"backend/procedure/driver/metadata"
	"backend/util/signature"
	"go.uber.org/fx"
)

func main() {
	fx.New(
		fx.Provide(
			config.Init,
			fiber.Init,
			grpc.Init,
			signature.Init,
			system.NewHandler,
			get.NewHandler,
		),
		fx.Invoke(
			metadata.Init,
			endpoint.Bind,
		),
	).Run()
}
