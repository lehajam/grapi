//+build wireinject

package di

import (
	"github.com/google/wire"
	"github.com/izumin5210/clig/pkg/clib"
	"github.com/izumin5210/gex/pkg/tool"

	"github.com/lehajam/grapi/pkg/cli"
	"github.com/lehajam/grapi/pkg/grapicmd"
	"github.com/lehajam/grapi/pkg/grapicmd/internal/module"
	"github.com/lehajam/grapi/pkg/grapicmd/internal/usecase"
	"github.com/lehajam/grapi/pkg/protoc"
)

func NewUI(*grapicmd.Ctx) cli.UI {
	wire.Build(Set)
	return nil
}

func NewScriptLoader(*grapicmd.Ctx) module.ScriptLoader {
	wire.Build(Set)
	return nil
}

func NewToolRepository(*grapicmd.Ctx) (tool.Repository, error) {
	wire.Build(Set)
	return nil, nil
}

func NewProtocWrapper(*grapicmd.Ctx) (protoc.Wrapper, error) {
	wire.Build(Set)
	return nil, nil
}

func NewInitializeProjectUsecase(*grapicmd.Ctx, clib.Path) (usecase.InitializeProjectUsecase, error) {
	wire.Build(Set)
	return nil, nil
}
