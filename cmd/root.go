package cmd

import (
	"github.com/spf13/cobra"
)

var (
	cfgFile string
)

var RootCmd = &cobra.Command{
	Use:   "todo",
	Short: "命令行todo",
	Long:  "命令行todo",
}

func init() {
	RootCmd.AddCommand(CreateCMD())
	RootCmd.AddCommand(DescribeCMD())
	RootCmd.AddCommand(QueryCMD())
	RootCmd.AddCommand(UpdateCMD())
	RootCmd.AddCommand(StatusCMD())
}
