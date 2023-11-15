package cmd

import (
	"fmt"
	"log"
	"os"

	"github.com/lujingwei002/gensc"
	"github.com/lujingwei002/gensc/cmd/gensc/cmd/gen"
	"github.com/lujingwei002/gensc/cmd/gensc/cmd/gen_behavior"
	"github.com/lujingwei002/gensc/cmd/gensc/cmd/gen_const"
	"github.com/lujingwei002/gensc/cmd/gensc/cmd/gen_model"
	"github.com/lujingwei002/gensc/cmd/gensc/cmd/gen_resource"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	configFile string
	rootCmd    = &cobra.Command{
		Use:   "gensc",
		Short: "Short",
		Long: `L
O
N
G`,
		Run: func(cmd *cobra.Command, args []string) {
			// Do Stuff Here
		},
	}
)

func init() {
	cobra.OnInitialize(initConfig)
	rootCmd.AddCommand(gen.Cmd)
	rootCmd.AddCommand(gen_resource.Cmd)
	rootCmd.AddCommand(gen_const.Cmd)
	rootCmd.AddCommand(gen_model.Cmd)
	rootCmd.AddCommand(gen_behavior.Cmd)
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func initConfig() {
	if configFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(configFile)
	} else {
		// Search config in home directory with name ".gensc" (without extension).
		viper.AddConfigPath(".")
		viper.SetConfigType("yaml")
		viper.SetConfigName(".gensc.yaml")
	}
	viper.SetEnvPrefix("gensc")
	viper.AutomaticEnv()
	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}
	var c gensc.Config
	err := viper.Unmarshal(&c)
	if err != nil {
		log.Fatalf("unable to decode into struct, %v", err)
	}
}
