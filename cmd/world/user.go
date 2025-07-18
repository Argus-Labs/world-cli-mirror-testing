package main

import (
	"context"

	cmdsetup "github.com/argus-labs/world-cli-mirror-testing/v2/internal/app/world-cli/controllers/cmd_setup"
	"github.com/argus-labs/world-cli-mirror-testing/v2/internal/app/world-cli/models"
)

//nolint:gochecknoglobals // standard kong plugin struct
var UserCmdPlugin struct {
	User *UserCmd `cmd:""`
}

//nolint:lll // needed to put all the help text in the same line
type UserCmd struct {
	Invite *InviteUserToOrganizationCmd     `cmd:"" group:"User Commands:" optional:"" help:"Invite a user to an organization"`
	Role   *ChangeUserRoleInOrganizationCmd `cmd:"" group:"User Commands:" optional:"" help:"Change a user's role in an organization"`
	Update *UpdateUserCmd                   `cmd:"" group:"User Commands:" optional:"" help:"Update a user"`
}

type InviteUserToOrganizationCmd struct {
	Context      context.Context       `kong:"-"`
	Dependencies cmdsetup.Dependencies `kong:"-"`
	Email        string                `         flag:"" help:"The email of the user to invite"`
	Role         string                `         flag:"" help:"The role of the user to invite"`
}

func (c *InviteUserToOrganizationCmd) Run() error {
	req := models.SetupRequest{
		LoginRequired:        models.NeedLogin,
		OrganizationRequired: models.NeedExistingData,
		ProjectRequired:      models.Ignore,
	}

	flags := models.InviteUserToOrganizationFlags{
		Email: c.Email,
		Role:  c.Role,
	}

	return cmdsetup.WithSetup(c.Context, c.Dependencies, req, func(state models.CommandState) error {
		return c.Dependencies.UserHandler.InviteToOrganization(c.Context, *state.Organization, flags)
	})
}

type ChangeUserRoleInOrganizationCmd struct {
	Context      context.Context       `kong:"-"`
	Dependencies cmdsetup.Dependencies `kong:"-"`
	Email        string                `         flag:"" help:"The email of the user to change the role of"`
	Role         string                `         flag:"" help:"The new role of the user"`
}

func (c *ChangeUserRoleInOrganizationCmd) Run() error {
	req := models.SetupRequest{
		LoginRequired:        models.NeedLogin,
		OrganizationRequired: models.NeedExistingData,
		ProjectRequired:      models.Ignore,
	}

	flags := models.ChangeUserRoleInOrganizationFlags{
		Email: c.Email,
		Role:  c.Role,
	}

	return cmdsetup.WithSetup(c.Context, c.Dependencies, req, func(state models.CommandState) error {
		return c.Dependencies.UserHandler.ChangeRoleInOrganization(c.Context, *state.Organization, flags)
	})
}

type UpdateUserCmd struct {
	Context      context.Context       `kong:"-"`
	Dependencies cmdsetup.Dependencies `kong:"-"`
	Name         string                `         flag:"" help:"The new name of the user"`
}

func (c *UpdateUserCmd) Run() error {
	req := models.SetupRequest{
		LoginRequired:        models.NeedLogin,
		OrganizationRequired: models.NeedExistingData,
		ProjectRequired:      models.Ignore,
	}

	flags := models.UpdateUserFlags{
		Name: c.Name,
	}

	return cmdsetup.WithSetup(c.Context, c.Dependencies, req, func(_ models.CommandState) error {
		err := c.Dependencies.UserHandler.Update(c.Context, flags)
		return err
	})
}
