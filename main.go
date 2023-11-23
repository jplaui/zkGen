package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"

	// "log"
	cmds "transpiler/commands"
)

func newVersionCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "version",
		Short: "Print the version number of zkGen CMD toolkit.",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("zkGen version 0.1")
		},
	}

	return cmd
}

func Commands() *cobra.Command {

	// create new cobra command
	cmd := &cobra.Command{
		Use:   "zkGen",
		Short: "\nWelcome,\n\nzkGen is a command-line tool to transpile the zkGen json DSL into secure computation circuits.\n",
	}

	// version command
	cmd.AddCommand(newVersionCommand())

	// transpiler
	cmd.AddCommand(cmds.PolicyTranspileCommand())
	cmd.AddCommand(cmds.PolicyGetCommand())
	cmd.AddCommand(cmds.PolicyListCommand())

	return cmd
}

func main() {

	// start command-line toolkit
	cmd := Commands()
	if err := cmd.Execute(); err != nil {
		os.Exit(0)
	}
}
