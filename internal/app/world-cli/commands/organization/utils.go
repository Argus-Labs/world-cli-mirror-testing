package organization

import (
	"github.com/argus-labs/world-cli-mirror-testing/v2/internal/app/world-cli/models"
	"github.com/argus-labs/world-cli-mirror-testing/v2/internal/pkg/printer"
	"github.com/rotisserie/eris"
)

func (h *Handler) saveOrganization(org models.Organization) error {
	h.configService.GetConfig().OrganizationID = org.ID
	err := h.configService.Save()
	if err != nil {
		return eris.Wrap(err, "Failed to save organization: "+org.Name)
	}
	return nil
}

func (h *Handler) showOrganizationList(org models.Organization, orgs []models.Organization) {
	printer.NewLine(1)
	printer.Headerln("  Organization Information  ")
	for _, organization := range orgs {
		if organization.ID == org.ID {
			printer.Infof("• %s (%s) [SELECTED]\n", organization.Name, organization.Slug)
		} else {
			printer.Infof("  %s (%s)\n", organization.Name, organization.Slug)
		}
	}
}

func (h *Handler) PrintNoOrganizations() {
	printer.NewLine(1)
	printer.Headerln("   No Organizations Found   ")
	printer.Info("1. Use ")
	printer.Notification("'world organization create'")
	printer.Infoln(" to create an organization.")
	printer.Info("2. Have a member send invite using ")
	printer.Notification("'world organization invite'")
	printer.Infoln(".")
}
