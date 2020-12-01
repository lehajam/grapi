package svcgen

import (
	"github.com/google/wire"

	"github.com/lehajam/grapi/pkg/cli"
	"github.com/lehajam/grapi/pkg/grapicmd"
	"github.com/lehajam/grapi/pkg/protoc"
	"github.com/lehajam/grapi/pkg/svcgen/params"
	_ "github.com/lehajam/grapi/pkg/svcgen/template"
)

func ProvideParamsBuilder(rootDir cli.RootDir, protocCfg *protoc.Config, grapiCfg *grapicmd.Config) params.Builder {
	return params.NewBuilder(
		rootDir,
		protocCfg.ProtosDir,
		protocCfg.OutDir,
		grapiCfg.Grapi.ServerDir,
		grapiCfg.Package,
	)
}

var Set = wire.NewSet(
	ProvideParamsBuilder,
	App{},
)
