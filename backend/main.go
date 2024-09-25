package main

import (
	"github.com/strawstacks/strawhouse/backend/common/config"
	"github.com/strawstacks/strawhouse/backend/common/fiber"
	"github.com/strawstacks/strawhouse/backend/common/grpc"
	"github.com/strawstacks/strawhouse/backend/common/pogreb"
	"github.com/strawstacks/strawhouse/backend/endpoint"
	"github.com/strawstacks/strawhouse/backend/endpoint/get"
	"github.com/strawstacks/strawhouse/backend/endpoint/system"
	"github.com/strawstacks/strawhouse/backend/procedure/driver/metadata"
	"github.com/strawstacks/strawhouse/backend/util/name"
	"github.com/strawstacks/strawhouse/backend/util/signature"
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
			endpoint.Bind,
			metadata.Init,
		),
	).Run()
}
