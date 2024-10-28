package main

import (
	"go.uber.org/fx"
	"strawhouse-backend/common/config"
	"strawhouse-backend/common/fiber"
	"strawhouse-backend/common/grpc"
	"strawhouse-backend/common/logger"
	"strawhouse-backend/common/pogreb"
	"strawhouse-backend/endpoint"
	"strawhouse-backend/endpoint/get"
	"strawhouse-backend/endpoint/system"
	"strawhouse-backend/procedure/driver/feed"
	"strawhouse-backend/procedure/driver/metadata"
	"strawhouse-backend/procedure/driver/transfer"
	"strawhouse-backend/service/file"
	"strawhouse-backend/service/plugin"
	"strawhouse-backend/util/eventfeed"
	"strawhouse-backend/util/fileflag"
	"strawhouse-backend/util/filepath"
	"strawhouse-backend/util/signature"
)

func main() {
	fx.New(
		logger.Init(),
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
			plugin.Serve,
			system.NewHandler,
			get.NewHandler,
		),
		fx.Invoke(
			invoke,
			endpoint.Bind,
			metadata.Register,
			feed.Register,
			transfer.Register,
		),
	).Run()
}

func invoke(plugin *plugin.Service) {
	_ = plugin
}
