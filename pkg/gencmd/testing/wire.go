//+build wireinject

package testing

import (
	"github.com/google/wire"

	"github.com/lehajam/grapi/pkg/cli"
	"github.com/lehajam/grapi/pkg/gencmd"
)

func NewTestApp(*gencmd.Command, cli.UI) (*gencmd.App, error) {
	wire.Build(
		gencmd.Set,
	)
	return nil, nil
}
