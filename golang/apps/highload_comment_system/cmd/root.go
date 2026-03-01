package cmd

import (
	"context"
	"io"

	"github.com/spf13/cobra"
)

type Kernel struct {
	command *cobra.Command
}

func NewKernel() *Kernel {
	cmdRoot := &Kernel{
		command: &cobra.Command{
			Use:     "demo",
			Aliases: []string{"go run ."},
			Short:   "CLI",
		},
	}

	return cmdRoot
}

func (k *Kernel) Execute(ctx context.Context, buf io.Writer) error {
	if k == nil {
		panic("nil command")
	}

	if k.command == nil {
		panic("command is nil")
	}

	k.command.SetOut(buf)

	return k.command.ExecuteContext(ctx)
}

func (k *Kernel) AddCommand(cmd *cobra.Command) {
	k.command.AddCommand(cmd)
}

// AddGroup adds commands, likely first command is parent of the rest commands
func (k *Kernel) AddGroup(cmds ...*cobra.Command) {
	var groupHead *cobra.Command
	for _, cmd := range cmds {
		if groupHead == nil {
			k.command.AddCommand(cmd)
			groupHead = cmd

			continue
		}

		groupHead.AddCommand(cmd)
	}
}
