package commands

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/anchore/clio"
	"DIDTrustCore/cmd/grype/cli/options"
	"DIDTrustCore/grype/db/v6/distribution"
	"DIDTrustCore/grype/db/v6/installation"
)

func DBDelete(app clio.Application) *cobra.Command {
	opts := options.DefaultDatabaseCommand(app.ID())

	cmd := &cobra.Command{
		Use:     "delete",
		Short:   "Delete the vulnerability database",
		Args:    cobra.ExactArgs(0),
		PreRunE: disableUI(app),
		RunE: func(_ *cobra.Command, _ []string) error {
			return runDBDelete(*opts)
		},
	}

	// prevent from being shown in the grype config
	type configWrapper struct {
		*options.DatabaseCommand `yaml:",inline" mapstructure:",squash"`
	}

	return app.SetupCommand(cmd, &configWrapper{opts})
}

func runDBDelete(opts options.DatabaseCommand) error {
	client, err := distribution.NewClient(opts.ToClientConfig())
	if err != nil {
		return fmt.Errorf("unable to create distribution client: %w", err)
	}
	c, err := installation.NewCurator(opts.ToCuratorConfig(), client)
	if err != nil {
		return fmt.Errorf("unable to create curator: %w", err)
	}

	if err := c.Delete(); err != nil {
		return fmt.Errorf("unable to delete vulnerability database: %+v", err)
	}

	return stderrPrintLnf("Vulnerability database deleted")
}
