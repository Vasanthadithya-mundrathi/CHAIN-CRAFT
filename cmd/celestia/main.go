// Author: Vasanthadithya Mundrathi from CBIT college
// ChainCraft - A blockchain node implementation for data availability sampling
package main

import (
	"context"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/pflag"

	cmdnode "github.com/celestiaorg/celestia-node/cmd"
)

func WithSubcommands() func(*cobra.Command, []*pflag.FlagSet) {
	return func(c *cobra.Command, flags []*pflag.FlagSet) {
		c.AddCommand(
			cmdnode.Init(flags...),
			cmdnode.Start(cmdnode.WithFlagSet(flags)),
			cmdnode.AuthCmd(flags...),
			cmdnode.ResetStore(flags...),
			cmdnode.RemoveConfigCmd(flags...),
			cmdnode.UpdateConfigCmd(flags...),
		)
	}
}

func init() {
	bridgeCmd := cmdnode.NewBridge(WithSubcommands())
	lightCmd := cmdnode.NewLight(WithSubcommands())
	fullCmd := cmdnode.NewFull(WithSubcommands())
	rootCmd.AddCommand(
		bridgeCmd,
		lightCmd,
		fullCmd,
		docgenCmd,
		versionCmd,
	)
	rootCmd.SetHelpCommand(&cobra.Command{})
}

func main() {
	err := run()
	if err != nil {
		os.Exit(1)
	}
}

func run() error {
	return rootCmd.ExecuteContext(context.Background())
}

var rootCmd = &cobra.Command{
	Use: "chaincraft [  bridge  ||  full ||  light  ] [subcommand]",
	Short: `
	 ██████   ██████ 
	██      ██
	██      ██
	██      ██
	 ██████   ██████

	 

	ChainCraft - Blockchain Infrastructure & Data Availability
	`,
	Args: cobra.NoArgs,
	CompletionOptions: cobra.CompletionOptions{
		DisableDefaultCmd: false,
	},
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		author, _ := cmd.Flags().GetBool("author")
		if author {
			cmd.Printf("ChainCraft v1.0.0\n")
			cmd.Printf("Author: Vasanthadithya Mundrathi\n")
			cmd.Printf("Institution: CBIT college\n")
			cmd.Printf("Description: Blockchain node implementation for data availability sampling\n")
			cmd.Printf("Contact: vasanthadithya@example.com\n")
			os.Exit(0)
		}
	},
}

func init() {
	rootCmd.PersistentFlags().BoolP("author", "a", false, "Display author information")
}
