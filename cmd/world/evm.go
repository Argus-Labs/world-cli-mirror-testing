package main

import (
	"context"

	cmdsetup "github.com/argus-labs/world-cli-mirror-testing/v2/internal/app/world-cli/controllers/cmd_setup"
	"github.com/argus-labs/world-cli-mirror-testing/v2/internal/app/world-cli/models"
)

//nolint:gochecknoglobals // standard kong plugin struct
var EvmCmdPlugin struct {
	Evm *EvmCmd `cmd:"" group:"EVM Commands:" help:"Manage your EVM blockchain environment"`
}

type EvmCmd struct {
	Config string `flag:"" help:"A TOML config file"`

	Start *StartEVMCmd `cmd:"" group:"EVM Commands:" help:"Launch your EVM blockchain environment"`
	Stop  *StopEVMCmd  `cmd:"" group:"EVM Commands:" help:"Shut down your EVM blockchain environment"`
}

//nolint:lll // needed to put all the help text in the same line
type StartEVMCmd struct {
	Parent       *EvmCmd               `kong:"-"`
	Context      context.Context       `kong:"-"`
	Dependencies cmdsetup.Dependencies `kong:"-"`
	DAAuthToken  string                `         flag:"" optional:"" help:"The DA Auth Token that allows the rollup to communicate with the Celestia client."`
	UseDevDA     bool                  `         flag:"" optional:"" help:"Use a locally running DA layer"                                                    name:"dev"`
}

func (c *StartEVMCmd) Run() error {
	flags := models.StartEVMFlags{
		Config:      c.Parent.Config,
		DAAuthToken: c.DAAuthToken,
		UseDevDA:    c.UseDevDA,
	}
	return c.Dependencies.EVMHandler.Start(c.Context, flags)
}

type StopEVMCmd struct {
	Parent       *EvmCmd               `kong:"-"`
	Context      context.Context       `kong:"-"`
	Dependencies cmdsetup.Dependencies `kong:"-"`
}

func (c *StopEVMCmd) Run() error {
	flags := models.StopEVMFlags{
		Config: c.Parent.Config,
	}
	return c.Dependencies.EVMHandler.Stop(c.Context, flags)
}
