package root

import (
	"github.com/argus-labs/world-cli/v2/internal/app/world-cli/shared/dependency"
	"github.com/argus-labs/world-cli/v2/internal/app/world-cli/shared/teacmd"
	"github.com/argus-labs/world-cli/v2/internal/pkg/tea/component/program"
	"github.com/argus-labs/world-cli/v2/internal/pkg/tea/style"
	tea "github.com/charmbracelet/bubbletea"
)

func (h *Handler) Doctor() error {
	p := program.NewTeaProgram(NewWorldDoctorModel())
	_, err := p.Run()
	if err != nil {
		return err
	}
	return nil
}

func doctorDeps() []dependency.Dependency {
	return []dependency.Dependency{
		dependency.Git,
		dependency.Go,
		dependency.Docker,
		dependency.DockerDaemon,
	}
}

//////////////////////
// Bubble Tea Model //
//////////////////////

type WorldDoctorModel struct {
	DepStatus    []teacmd.DependencyStatus
	DepStatusErr error
}

func NewWorldDoctorModel() WorldDoctorModel {
	return WorldDoctorModel{}
}

//////////////////////////
// Bubble Tea Lifecycle //
//////////////////////////

// Init returns an initial command for the application to run.
func (m WorldDoctorModel) Init() tea.Cmd {
	return teacmd.CheckDependenciesCmd(doctorDeps())
}

// Update handles incoming events and updates the model accordingly.
func (m WorldDoctorModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.Type { //nolint:gocritic,exhaustive // cleaner with switch
		case tea.KeyCtrlC:
			return m, tea.Quit
		}
	case teacmd.CheckDependenciesMsg:
		m.DepStatus = msg.DepStatus
		m.DepStatusErr = msg.Err
		return m, tea.Quit
	}
	return m, nil
}

// View renders the model to the screen.
func (m WorldDoctorModel) View() string {
	depList, help := teacmd.PrintDependencyStatus(m.DepStatus)
	out := style.Container.Render("--- World CLI Doctor ---") + "\n\n"
	out += "Checking dependencies...\n"
	out += depList + "\n" + help + "\n"
	return out
}
