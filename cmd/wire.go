//go:build wireinject
// +build wireinject

package cmd

import (
	"github.com/google/wire"
	"github.com/spf13/cobra"
	"github.com/xuender/fairy/meta"
	"github.com/xuender/fairy/pb"
	"github.com/xuender/fairy/ui"
	"github.com/xuender/fairy/watch"
)

func InitMeta(cmd *cobra.Command) *meta.Service {
	wire.Build(
		meta.NewService,
	)

	return &meta.Service{}
}

func InitUI(cmd *cobra.Command) *ui.Service {
	wire.Build(
		ui.NewService,
		pb.NewConfig,
	)

	return &ui.Service{}
}

func InitWatch(cmd *cobra.Command) *watch.Service {
	wire.Build(
		meta.NewService,
		pb.NewConfig,
		watch.NewService,
	)

	return &watch.Service{}
}
