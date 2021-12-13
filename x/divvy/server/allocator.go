package server

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/errors"

	// "github.com/regen-network/regen-ledger/types"
	"github.com/regen-network/regen-ledger/v2/x/divvy"
)

// s is the root store
func distributeBalance(addr sdk.AccAddress, a *divvy.StoreAllocator, bank divvy.BankKeeper, ctx *sdk.Context) (sdk.Coins, error) {
	// TODO: charge extra gas
	totalShares := sdk.NewIntFromUint64(divvy.TotalShares)
	coins := bank.GetAllBalances(*ctx, addr)
	for _, r := range a.Recipients {
		out := make(sdk.Coins, len(coins))
		// TODO: maybe reorder the loop to optimize the `amount` computation
		for i, c := range coins {
			out[i].Denom = c.Denom
			out[i].Amount = c.Amount.Mul(sdk.NewIntFromUint64(uint64(r.Share))).Quo(totalShares)
		}
		if err := bank.SendCoins(*ctx, addr, r.Address, out); err != nil {
			return nil, err
		}
	}
	return coins, nil
}

// returns error if sender is not allocator admin
func assertAllocatorAdmin(sender string, a *divvy.StoreAllocator) error {
	if a.Admin != sender {
		return errors.ErrUnauthorized
	}
	return nil
}

func storeAllocatorToAllocator(as *divvy.StoreAllocator, addr string) *divvy.Allocator {
	rs := make([]divvy.Recipient, len(as.Recipients))
	for i := range as.Recipients {
		rs[i] = storeRecipientToRecipient(&as.Recipients[i])
	}
	return &divvy.Allocator{
		Admin:      as.Admin,
		Start:      as.Start,
		End:        as.End,
		Interval:   as.Interval,
		Name:       as.Name,
		Url:        as.Url,
		Paused:     as.Paused,
		Address:    addr,
		Recipients: rs,
		NextClaim:  as.NextClaim,
	}
}
