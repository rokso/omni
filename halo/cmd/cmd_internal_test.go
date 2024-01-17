package cmd

import (
	"bytes"
	"context"
	"testing"

	"github.com/omni-network/omni/halo/app"
	libcmd "github.com/omni-network/omni/lib/cmd"
	"github.com/omni-network/omni/test/tutil"

	"github.com/stretchr/testify/require"
)

//go:generate go test . -update -clean

func TestRunCmd(t *testing.T) { //nolint:paralleltest,tparallel // RunCmd modifies global state via setMonikerForT
	setMonikerForT(t)

	tests := []struct {
		Name  string
		Args  []string
		Files map[string][]byte
	}{
		{
			Name: "defaults",
			Args: slice("run"),
		},
		{
			Name: "flags",
			Args: slice("run", "--home=foo", "--engine-jwt-file=bar"),
		},
		{
			Name: "toml files",
			Args: slice("run", "--home=testinput/input1"),
		},
		{
			Name: "json files",
			Args: slice("run", "--home=testinput/input2"),
		},
	}

	for _, test := range tests {
		test := test      // Pin
		args := test.Args // Pin
		t.Run(test.Name, func(t *testing.T) {
			t.Parallel()

			cmd := newRunCmd(func(_ context.Context, actual app.Config) error {
				tutil.RequireGoldenJSON(t, actual)

				return nil
			})

			rootCmd := libcmd.NewRootCmd("halo", "", cmd)
			rootCmd.SetArgs(args)
			require.NoError(t, rootCmd.Execute())
		})
	}
}

func TestCLIReference(t *testing.T) {
	t.Parallel()
	const root = "halo" // Use to identify root command (vs subcommands).

	tests := []struct {
		Command string
	}{
		{root},
		{"run"},
	}

	for _, test := range tests {
		test := test // Pin
		t.Run(test.Command, func(t *testing.T) {
			t.Parallel()

			var args []string
			if test.Command != root {
				args = append(args, test.Command)
			}
			args = append(args, "--help")

			cmd := New()
			cmd.SetArgs(args)

			var bz bytes.Buffer
			cmd.SetOut(&bz)

			require.NoError(t, cmd.Execute())

			tutil.RequireGoldenBytes(t, bz.Bytes())
		})
	}
}

// slice is a convenience function for creating string slice literals.
func slice(strs ...string) []string {
	return strs
}