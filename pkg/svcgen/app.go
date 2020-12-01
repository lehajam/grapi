package svcgen

import (
	"github.com/lehajam/grapi/pkg/gencmd"
	"github.com/lehajam/grapi/pkg/protoc"
	"github.com/lehajam/grapi/pkg/svcgen/params"
)

type CreateAppFunc func(*gencmd.Command) (*App, error)

type App struct {
	ProtocWrapper protoc.Wrapper
	ParamsBuilder params.Builder
}
