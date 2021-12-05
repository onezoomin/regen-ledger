package client

import (
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/spf13/cobra"
	// "github.com/regen-network/regen-ledger/v2/x/divvy"
)

// QueryCmd returns the parent command for all x/divvy query commands.
func QueryCmd(name string) *cobra.Command {
	cmd := &cobra.Command{
		SuggestionsMinimumDistance: 2,
		DisableFlagParsing:         true,

		Args:  cobra.ExactArgs(1),
		Use:   name,
		Short: "Query commands for the divvy module",
		RunE:  client.ValidateCmd,
	}
	// cmd.AddCommand(
	// )
	return cmd
}

func qflags(cmd *cobra.Command) *cobra.Command {
	flags.AddQueryFlagsToCmd(cmd)
	return cmd
}
