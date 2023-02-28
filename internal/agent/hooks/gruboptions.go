package hook

import (
	"fmt"

	config "github.com/kairos-io/kairos/pkg/config"
	schema "github.com/kairos-io/kairos/pkg/config/schemas"
	"github.com/kairos-io/kairos/sdk/system"
)

type GrubOptions struct{}

func (b GrubOptions) Run(c config.Config) error {
	err := system.Apply(system.SetGRUBOptions(c.Install.GrubOptions))
	if err != nil {
		fmt.Println(err)
	}
	return nil
}

// KRun is a temporary function that does the same as Run. It will be removed as soon as the transition from config.Config to schema.KConfig is finished.
func (b GrubOptions) KRun(kc schema.KConfig) error {
	grubOptions, err := kc.GrubOptions()
	if err != nil {
		fmt.Println(err)
	}
	err = system.Apply(system.SetGRUBOptions(grubOptions))
	if err != nil {
		fmt.Println(err)
	}
	return nil
}

type GrubPostInstallOptions struct{}

func (b GrubPostInstallOptions) Run(c config.Config) error {
	err := system.Apply(system.SetGRUBOptions(c.GrubOptions))
	if err != nil {
		fmt.Println(err)
	}
	return nil
}

// KRun is a temporary function that does the same as Run. It will be removed as soon as the transition from config.Config to schema.KConfig is finished.
func (b GrubPostInstallOptions) KRun(kc schema.KConfig) error {
	grubOptions, err := kc.GrubOptions()
	if err != nil {
		fmt.Println(err)
	}
	err = system.Apply(system.SetGRUBOptions(grubOptions))
	if err != nil {
		fmt.Println(err)
	}
	return nil
}
