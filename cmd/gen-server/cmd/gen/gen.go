package gen

import (
	"github.com/spf13/cobra"
)

var ()

var Cmd = &cobra.Command{
	Use:   "gen",
	Short: "gen code",
	RunE: func(cmd *cobra.Command, args []string) error {
		return nil
	},
}

func init() {
}
