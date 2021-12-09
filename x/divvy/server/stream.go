package server

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/regen-network/regen-ledger/v2/x/divvy"
)

func storeStreamToStream(s *divvy.StoreSlowReleaseStream, addr string) *divvy.SlowReleaseStream {
	return &divvy.SlowReleaseStream{
		Admin:       sdk.AccAddress(s.Admin).String(),
		Start:       s.Start,
		Interval:    s.Interval,
		Destination: sdk.AccAddress(s.Destination).String(),
		Name:        s.Name,
		Paused:      false,
		Strategy:    s.Strategy,
		Address:     addr,
	}
}
