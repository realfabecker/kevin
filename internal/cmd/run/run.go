package run

import (
	"fmt"

	"github.com/realfabecker/kevin/internal/adapters/logger"
	"github.com/realfabecker/kevin/internal/adapters/render"
	"github.com/realfabecker/kevin/internal/adapters/runner"
	"github.com/realfabecker/kevin/internal/core/domain"
	"github.com/spf13/cobra"
)

var DryRun bool

func subCmdRunE(c domain.Cmd) func(cmd *cobra.Command, args []string) error {
	return func(cmd *cobra.Command, args []string) error {
		for _, f := range c.Flags {
			v, _ := cmd.Flags().GetString(f.Name)
			if err := c.SetFlag(f.Name, v); err != nil {
				return err
			}
		}
		if c.Type == "proxy" && len(args) > 0 {
			for x, y := range args {
				c.AddArg(fmt.Sprintf("arg_%d", x), y)
			}
		} else if len(args) > 0 && len(args) == len(c.Args) {
			for i, a := range args {
				if err := c.SetArg(i, a); err != nil {
					return err
				}
			}
		}
		if c.GetNofRequiredArgs() > 0 && len(args) < c.GetNofRequiredArgs() {
			return fmt.Errorf("you must supply at least %d arguments for this command", c.GetNofRequiredArgs())
		}
		rn := runner.New(runner.NewCliOpts{
			Logger: logger.NewConsoleLogger(),
			Render: render.NewScriptRender(),
		})
		return rn.Run(&c, DryRun)
	}
}

func newSubCmd(c domain.Cmd) *cobra.Command {
	render := render.NewScriptRender()
	if va, _ := render.Render(&c, c.Short); va != "" {
		c.Short = va
	}
	cmd := &cobra.Command{
		Use:   c.Name,
		Short: c.Short,
		RunE:  subCmdRunE(c),
	}
	for _, f := range c.Flags {
		if f.Short == "" {
			cmd.Flags().String(f.Name, f.Value, f.Usage)
		} else {
			cmd.Flags().StringP(f.Name, f.Short, f.Value, f.Usage)
		}
		if f.Required {
			_ = cmd.MarkFlagRequired(f.Name)
		}
	}
	return cmd
}

func pushUnqCmd(root *cobra.Command, cmd *cobra.Command) {
	for _, v := range root.Commands() {
		if v.Name() == cmd.Name() {
			root.RemoveCommand(v)
			break
		}
	}
	root.AddCommand(cmd)
}

func pushMatrixCmd(c domain.Cmd, root *cobra.Command) {
	if c.Matrix != nil && len(c.Matrix.Name) > 0 {
		for _, v := range c.Matrix.Name {
			c.Name = v
			pushUnqCmd(root, newSubCmd(c))
		}
	} else {
		pushUnqCmd(root, newSubCmd(c))
	}
}

func pushGroupCmd(c domain.Cmd, root *cobra.Command) {
	groupCmd := &cobra.Command{
		Use:   c.Name,
		Short: c.Short,
	}
	AttachCmd(groupCmd, c.Commands)
	pushUnqCmd(root, groupCmd)
}

func AttachCmd(root *cobra.Command, cmds []domain.Cmd) {
	for _, v := range cmds {
		func(c domain.Cmd) {
			if c.Commands != nil {
				pushGroupCmd(c, root)
			} else {
				pushMatrixCmd(c, root)
			}
		}(v)
	}
	root.PersistentFlags().BoolVar(&DryRun, "dry-run", false, "run in dry run mode")
}
