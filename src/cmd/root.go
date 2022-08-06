package cmd

import "github.com/spf13/cobra"

func GetRoot() *cobra.Command {
	cmdRoot := &cobra.Command{}

	cmdRoot.AddCommand(server())

	return cmdRoot
}
