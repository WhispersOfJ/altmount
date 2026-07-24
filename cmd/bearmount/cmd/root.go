package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var (
	configFile string
)

var rootCmd = &cobra.Command{
	Use:   "bearmount",
	Short: "BearMount WebDAV server backed by NZB/Usenet",
}

func init() {
	rootCmd.PersistentFlags().StringVar(&configFile, "config", "./config.yaml", "config file (default is ./config.yaml)")
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
