//+build wireinject

package testing

import (
	"github.com/google/wire"

	"github.com/lehajam/grapi/pkg/cli"
	"github.com/lehajam/grapi/pkg/gencmd"
	"github.com/lehajam/grapi/pkg/protoc"
	"github.com/lehajam/grapi/pkg/svcgen"
)

func NewTestApp(*gencmd.Command, protoc.Wrapper, cli.UI) (*svcgen.App, error) {
	wire.Build(
		gencmd.Set,
		svcgen.ProvideParamsBuilder,
		svcgen.App{},
	)
	return nil, nil
}
