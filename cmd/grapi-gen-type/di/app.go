package di

import (
	"github.com/lehajam/grapi/pkg/gencmd"
	"github.com/lehajam/grapi/pkg/protoc"
)

type CreateAppFunc func(*gencmd.Command) (*App, error)

type App struct {
	Protoc protoc.Wrapper
}
