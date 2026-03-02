package cmd

import (
	"bytes"
	"context"
	"fmt"

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

func (k *Kernel) Execute(ctx context.Context) error {
	buf := bytes.NewBuffer(make([]byte, 0))

	k.command.SetOut(buf)
	defer fmt.Print(buf.String())

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
