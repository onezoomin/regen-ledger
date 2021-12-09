package client

import (
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/spf13/cobra"

	// "github.com/regen-network/regen-ledger/v2/x/divvy"

	"github.com/regen-network/regen-ledger/v2/x/divvy"
	"github.com/regen-network/regen-ledger/v2/x/divvy/parse"
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
		queryAllocatorCmd(),

		queryStreamCmd(),
		queryStreamsCmd(),
	)
	return cmd
}

func qflags(cmd *cobra.Command) *cobra.Command {
	flags.AddQueryFlagsToCmd(cmd)
	return cmd
}

// queryAllocatorCmd returns a query command that selects allocator by address.
func queryAllocatorCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "allocator [address]",
		Short: "selects allocator by address",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			c, ctx, err := mkQueryClient(cmd)
			if err != nil {
				return err
			}
			addr := args[0]
			if err := parse.CheckAddress(addr, "allocator"); err != nil {
				return err
			}
			res, err := c.AllocatorByAddress(cmd.Context(), &divvy.QueryAllocator{
				Address: addr,
			})
			return print(ctx, res, err)
		},
	}
	return qflags(cmd)
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
	flags.AddPaginationFlagsToCmd(cmd, "allocators")
	return qflags(cmd)
}

//
// STRAEMS
//

// queryStreamsCmd returns a query command that selects allocator by address.
func queryStreamCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "stream [address]",
		Short: "selects allocator by address",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			c, ctx, err := mkQueryClient(cmd)
			if err != nil {
				return err
			}
			addr := args[0]
			if err := parse.CheckAddress(addr, "allocator"); err != nil {
				return err
			}
			res, err := c.StreamByAddress(cmd.Context(), &divvy.QueryStream{
				Address: addr,
			})
			return print(ctx, res, err)
		},
	}
	return qflags(cmd)
}

// queryStreamsCmd returns a query command that lists all slow release streams.
func queryStreamsCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "streams",
		Short: "List all slow release streams with pagination flags",
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
			res, err := c.Streams(cmd.Context(), &divvy.QueryStreams{
				Pagination: pagination,
			})
			return print(ctx, res, err)
		},
	}
	flags.AddPaginationFlagsToCmd(cmd, "streams")
	return qflags(cmd)
}
