package contracts

import "github.com/spf13/cobra"

type CMDAdder interface {
	Add(cmd *cobra.Command)
	// AddGroup adds commands, likely first command is parent of the rest commands
	AddGroup(cmds ...*cobra.Command)
}
