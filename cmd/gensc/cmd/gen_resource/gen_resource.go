package gen_resource

import (
	"fmt"
	"log"

	"github.com/lujingwei002/gensc"
	gen "github.com/lujingwei002/gensc/gen/gen_resource"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	respositoryVersion string
)

var Cmd = &cobra.Command{
	Use:   "gen-resource",
	Short: "gen resource code",
	RunE: func(cmd *cobra.Command, args []string) error {
		var c gensc.Config
		err := viper.Unmarshal(&c)
		if err != nil {
			log.Fatalf("unable to decode into struct, %v", err)
		}
		fmt.Printf("read config %#v\n", c.GenResource)
		if err := gen.Gen(c.GenResource); err != nil {
			return err
		}
		return nil
	},
}

func init() {
	Cmd.Flags().StringVar(&respositoryVersion, "version", "", "respository  version")
	viper.BindPFlag("GenResource.RespositoryVersion", Cmd.Flags().Lookup("version"))
}
