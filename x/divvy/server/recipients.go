package server

import (
	"github.com/cosmos/cosmos-sdk/types"
	"github.com/regen-network/regen-ledger/v2/x/divvy"
)

func recipientsToStoreRecipients(rs []divvy.Recipient) ([]divvy.StoreRecipient, error) {
	srs := make([]divvy.StoreRecipient, len(rs))
	var err error
	for i, r := range rs {
		srs[i].Address, err = types.AccAddressFromBech32(r.Address)
		if err != nil {
			return nil, err
		}
		srs[i].Name = r.Name
		srs[i].Share = r.Share
	}
	return srs, nil
}

func storeRecipientToRecipient(r *divvy.StoreRecipient) divvy.Recipient {
	return divvy.Recipient{
		Address: types.AccAddress(r.Address).String(),
		Share:   r.Share,
		Name:    r.Name,
	}
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
