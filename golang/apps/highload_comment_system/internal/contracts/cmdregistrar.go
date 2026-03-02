package contracts

import "github.com/spf13/cobra"

type CmdRegistrar interface {
	Register(cmd *cobra.Command)
	// RegisterGroup adds commands, likely first command is parent of the rest commands
	RegisterGroup(cmds ...*cobra.Command)
}
