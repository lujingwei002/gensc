package gen_const

import (
	"fmt"
	"log"

	"github.com/lujingwei002/gensc"
	gen "github.com/lujingwei002/gensc/gen/gen_const"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	respositoryVersion string
)

var Cmd = &cobra.Command{
	Use:   "gen_const",
	Short: "gen const code",
	RunE: func(cmd *cobra.Command, args []string) error {
		var c gensc.Config
		err := viper.Unmarshal(&c)
		if err != nil {
			log.Fatalf("unable to decode into struct, %v", err)
		}
		fmt.Printf("read config %#v\n", c.GenConst)
		if err := gen.Gen(c.GenConst); err != nil {
			return err
		}
		return nil
	},
}

func init() {

}
