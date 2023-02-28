package hook

import (
	config "github.com/kairos-io/kairos/pkg/config"
	schema "github.com/kairos-io/kairos/pkg/config/schemas"
	"github.com/kairos-io/kairos/pkg/utils"

	events "github.com/kairos-io/kairos/sdk/bus"
)

type RunStage struct{}

func (r RunStage) Run(c config.Config) error {
	utils.SH("elemental run-stage kairos-install.after")             //nolint:errcheck
	events.RunHookScript("/usr/bin/kairos-agent.install.after.hook") //nolint:errcheck
	return nil
}

// KRun is a temporary function that does the same as Run. It will be removed as soon as the transition from config.Config to schema.KConfig is finished.
func (r RunStage) KRun(kc schema.KConfig) error {
	utils.SH("elemental run-stage kairos-install.after")             //nolint:errcheck
	events.RunHookScript("/usr/bin/kairos-agent.install.after.hook") //nolint:errcheck
	return nil
}
