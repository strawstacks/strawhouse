package main

import (
	"github.com/strawstacks/strawhouse/strawhouse-backend/common/config"
	"github.com/strawstacks/strawhouse/strawhouse-backend/common/fiber"
	"github.com/strawstacks/strawhouse/strawhouse-backend/common/grpc"
	"github.com/strawstacks/strawhouse/strawhouse-backend/common/pogreb"
	"github.com/strawstacks/strawhouse/strawhouse-backend/endpoint"
	"github.com/strawstacks/strawhouse/strawhouse-backend/endpoint/get"
	"github.com/strawstacks/strawhouse/strawhouse-backend/endpoint/system"
	"github.com/strawstacks/strawhouse/strawhouse-backend/procedure/driver/feed"
	"github.com/strawstacks/strawhouse/strawhouse-backend/procedure/driver/metadata"
	"github.com/strawstacks/strawhouse/strawhouse-backend/service/file"
	"github.com/strawstacks/strawhouse/strawhouse-backend/util/eventfeed"
	"github.com/strawstacks/strawhouse/strawhouse-backend/util/fileflag"
	"github.com/strawstacks/strawhouse/strawhouse-backend/util/filepath"
	"github.com/strawstacks/strawhouse/strawhouse-backend/util/signature"
	"go.uber.org/fx"
)

func main() {
	fx.New(
		fx.Provide(
			config.Init,
			pogreb.Init,
			fiber.Init,
			grpc.Init,
			filepath.Init,
			fileflag.Init,
			eventfeed.Init,
			signature.Init,
			file.Serve,
			system.NewHandler,
			get.NewHandler,
		),
		fx.Invoke(
			endpoint.Bind,
			metadata.Register,
			feed.Register,
		),
	).Run()
}
