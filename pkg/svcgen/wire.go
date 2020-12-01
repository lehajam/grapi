//+build wireinject

package svcgen

import (
	"github.com/google/wire"

	"github.com/lehajam/grapi/pkg/cli"
	"github.com/lehajam/grapi/pkg/gencmd"
	"github.com/lehajam/grapi/pkg/protoc"
)

func NewApp(*gencmd.Command) (*App, error) {
	wire.Build(
		Set,
		gencmd.Set,
		cli.UIInstance,
		protoc.WrapperSet,
	)
	return nil, nil
}
