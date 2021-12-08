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
