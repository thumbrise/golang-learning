package core

import (
	"bytes"
	"context"

	"github.com/spf13/cobra"
)

type Kernel struct {
	command *cobra.Command
}

func NewKernel() *Kernel {
	cmdRoot := &Kernel{
		command: &cobra.Command{
			Use: "demo",
		},
	}

	return cmdRoot
}

func (k *Kernel) Execute(ctx context.Context, bufout *bytes.Buffer, buferr *bytes.Buffer) error {
	k.command.SetOut(bufout)
	k.command.SetErr(buferr)

	return k.command.ExecuteContext(ctx)
}

func (k *Kernel) Add(cmd *cobra.Command) {
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
