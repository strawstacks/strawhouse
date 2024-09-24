package main

import (
	"backend/common/config"
	"backend/common/fiber"
	"backend/common/grpc"
	"backend/common/pogreb"
	"backend/endpoint"
	"backend/endpoint/get"
	"backend/endpoint/system"
	"backend/procedure/driver/metadata"
	"backend/util/name"
	"backend/util/signature"
	"go.uber.org/fx"
)

func main() {
	fx.New(
		fx.Provide(
			config.Init,
			pogreb.Init,
			fiber.Init,
			grpc.Init,
			name.Init,
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
