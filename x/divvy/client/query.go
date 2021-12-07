package client

import (
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/regen-network/regen-ledger/v2/x/divvy"
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
	cmd.AddCommand(
		queryAllocatorsCmd(),
	)
	return cmd
}

func qflags(cmd *cobra.Command) *cobra.Command {
	flags.AddQueryFlagsToCmd(cmd)
	return cmd
}

// queryAllocatorsCmd returns a query command that lists all allocators.
func queryAllocatorsCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "allocators",
		Short: "List all allocators with pagination flags",
		Args:  cobra.ExactArgs(0),
		RunE: func(cmd *cobra.Command, args []string) error {
			c, ctx, err := mkQueryClient(cmd)
			if err != nil {
				return err
			}

			pagination, err := client.ReadPageRequest(cmd.Flags())
			if err != nil {
				return err
			}

			res, err := c.Allocators(cmd.Context(), &divvy.QueryAllocators{
				Pagination: pagination,
			})
			return print(ctx, res, err)
		},
	}
	flags.AddPaginationFlagsToCmd(cmd, "classes")
	return qflags(cmd)
}
