package main

import (
	"testing"

	"github.com/bradleyjkemp/cupaloy/v2"
	"github.com/izumin5210/clig/pkg/clib"
	"github.com/lehajam/grapi/pkg/cli"
	"github.com/lehajam/grapi/pkg/gencmd"
	gencmdtesting "github.com/lehajam/grapi/pkg/gencmd/testing"
	"github.com/lehajam/grapi/pkg/grapicmd"
	"github.com/spf13/afero"
)

func TestCommand(t *testing.T) {
	cases := []struct {
		test  string
		args  []string
		files []string
	}{
		{
			test:  "simple",
			args:  []string{"foo"},
			files: []string{"cmd/foo/run.go"},
		},
	}

	rootDir := cli.RootDir{clib.Path("/home/src/testapp")}

	createGenApp := func(cmd *gencmd.Command) (*gencmd.App, error) {
		return gencmdtesting.NewTestApp(cmd, cli.NopUI)
	}
	createCmd := func(t *testing.T, fs afero.Fs) gencmd.Executor {
		ctx := &grapicmd.Ctx{
			FS:      fs,
			RootDir: rootDir,
		}
		return buildCommand(gencmd.WithGrapiCtx(ctx), gencmd.WithCreateAppFunc(createGenApp))
	}

	for _, tc := range cases {
		t.Run(tc.test, func(t *testing.T) {
			fs := afero.NewMemMapFs()
			afero.WriteFile(fs, rootDir.Join("grapi.toml").String(), []byte{}, 0755)

			t.Run("generate", func(t *testing.T) {
				cmd := createCmd(t, fs)
				cmd.Command().SetArgs(append([]string{"generate"}, tc.args...))
				err := cmd.Execute()

				if err != nil {
					t.Errorf("returned an error: %+v", err)
				}

				for _, file := range tc.files {
					t.Run(file, func(t *testing.T) {
						data, err := afero.ReadFile(fs, rootDir.Join(file).String())

						if err != nil {
							t.Errorf("returned an error: %v", err)
						}

						cupaloy.SnapshotT(t, string(data))
					})
				}
			})

			t.Run("destroy", func(t *testing.T) {
				cmd := createCmd(t, fs)
				cmd.Command().SetArgs(append([]string{"destroy"}, tc.args...))
				err := cmd.Execute()

				if err != nil {
					t.Errorf("returned an error: %+v", err)
				}

				for _, file := range tc.files {
					t.Run(file, func(t *testing.T) {
						ok, err := afero.Exists(fs, rootDir.Join(file).String())

						if err != nil {
							t.Errorf("Exists(fs, %q) returned an error: %v", file, err)
						}

						if ok {
							t.Errorf("%q should not exist", file)
						}
					})
				}
			})
		})
	}
}
