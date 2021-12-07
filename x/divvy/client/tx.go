package client

import (
	"time"

	"fmt"
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
		Use:   "create-allocator [name] [url] [startTime] [endTime] [interval] [recipientMap]",
		Short: "Creates a new allocator for a set of recipients, time is expected in RFC3339 format (eg '2006-01-02T15:04:05.999999-07:00'), and interval is a number in seconds. `recipientMap` is a JSON map of address to a share, where 1000000 share = 100%",
		Args:  cobra.ExactArgs(6),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := sdkclient.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			var errmsgs []string
			admin := clientCtx.GetFromAddress() // --from flag
			startTime, errmsgs := parseTime(args[2], "startTime", errmsgs)
			endTime, errmsgs := parseTime(args[3], "endTime", errmsgs)
			interval, errmsgs := parseUint(args[4], "interval", errmsgs)
			recipients, errmsgs := parseRecipient(args[5], errmsgs)
			if err := divvy.ErrorStringsToError(errmsgs); err != nil {
				return err
			}

			fmt.Println(">>> end after start", endTime.After(startTime))

			msg := divvy.MsgCreateAllocator{
				Admin:      admin.String(),
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
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), &msg)
		},
	})
}
