package client

import (
	"time"

	sdkclient "github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/regen-network/regen-ledger/v2/x/divvy"
	"github.com/spf13/cobra"
)

// TxCmd returns a root CLI command handler for all x/divvy transaction commands.
func TxCmd(name string) *cobra.Command {
	cmd := &cobra.Command{
		SuggestionsMinimumDistance: 2,
		DisableFlagParsing:         true,

		Use:   name,
		Short: "Ecocredit module transactions",
		RunE:  sdkclient.ValidateCmd,
	}
	cmd.AddCommand(
		TxCreateAllocator(),
		TxClaimAllocator(),
		TxSetAllocatorRecipients(),
	)
	return cmd
}

func txflags(cmd *cobra.Command) *cobra.Command {
	flags.AddTxFlagsToCmd(cmd)
	cmd.MarkFlagRequired(flags.FlagFrom)
	return cmd
}

// TxCreateAllocator returns a transaction command that creates a new allocator.
func TxCreateAllocator() *cobra.Command {
	return txflags(&cobra.Command{
		Use:   "create-allocator [name] [url] [startTime] [endTime] [interval] [recipient-map]",
		Short: "Creates a new allocator for a set of recipients, time is expected in RFC3339 format (eg '2006-01-02T15:04:05.999999-07:00'), and interval is a number in seconds. `recipient-map` is a JSON map of address to a share, where 1000000 share = 100%",
		Args:  cobra.ExactArgs(6),
		RunE: func(cmd *cobra.Command, args []string) error {
			cctx, err := sdkclient.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			var errmsgs []string
			startTime, errmsgs := parseTime(args[2], "startTime", errmsgs)
			endTime, errmsgs := parseTime(args[3], "endTime", errmsgs)
			interval, errmsgs := parseUint(args[4], "interval", errmsgs)
			recipients, errmsgs := parseRecipient(args[5], errmsgs)
			if err := divvy.ErrorStringsToError(errmsgs); err != nil {
				return err
			}

			msg := divvy.MsgCreateAllocator{
				Admin:      cctx.GetFromAddress().String(), // --from flagadmin.String(),
				Start:      startTime,
				End:        endTime,
				Interval:   time.Duration(interval) * time.Second,
				Name:       args[0],
				Url:        args[1],
				Recipients: recipients,
			}
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(cctx, cmd.Flags(), &msg)
		},
	})
}

// TxClaimAllocator returns a transaction command that distribute allocator balances to
// recipients.
func TxClaimAllocator() *cobra.Command {
	return txflags(&cobra.Command{
		Use:   "claim [allocator_address]",
		Short: "distribute allocator balances to its recipients",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			cctx, err := sdkclient.GetClientTxContext(cmd)
			if err != nil {
				return err
			}
			msg := divvy.MsgClaimAllocations{
				Sender:    cctx.GetFromAddress().String(),
				Allocator: args[0]}
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(cctx, cmd.Flags(), &msg)
		},
	})
}

// TxSetAllocatorRecipients updates allocator recipient list
func TxSetAllocatorRecipients() *cobra.Command {
	return txflags(&cobra.Command{
		Use:   "set-allocator-recipients [allocator_address] [recipient-map]",
		Short: "TxSetAllocatorRecipients updates allocator recipient list. See create-allocator for inforamtion about the recipient-map format",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			cctx, err := sdkclient.GetClientTxContext(cmd)
			if err != nil {
				return err
			}
			recipients, errmsgs := parseRecipient(args[1], nil)
			if err = divvy.ErrorStringsToError(errmsgs); err != nil {
				return err
			}
			msg := divvy.MsgSetAllocatorRecipients{
				Sender:     cctx.GetFromAddress().String(),
				Address:    args[0],
				Recipients: recipients}
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(cctx, cmd.Flags(), &msg)
		},
	})
}
