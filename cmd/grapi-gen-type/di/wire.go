//+build wireinject

package di

import (
	"github.com/google/wire"

	"github.com/lehajam/grapi/pkg/cli"
	"github.com/lehajam/grapi/pkg/gencmd"
	"github.com/lehajam/grapi/pkg/protoc"
)

func NewApp(*gencmd.Command) (*App, error) {
	wire.Build(
		App{},
		gencmd.Set,
		cli.UIInstance,
		protoc.WrapperSet,
	)
	return nil, nil
}
