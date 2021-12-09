package server

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	// sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/cosmos/cosmos-sdk/types/query"

	"github.com/regen-network/regen-ledger/v2/x/divvy"
)

func (s serverImpl) StreamByAddress(goCtx context.Context, req *divvy.QueryStream) (*divvy.SlowReleaseStream, error) {
	ctx, err := unwrapAndCheck(goCtx, req)
	if err != nil {
		return nil, err
	}
	_, _, a, err := s.getSlowReleaseStream(ctx, req.Address)
	return storeStreamToStream(a, req.Address), err
}

func (s serverImpl) Streams(goCtx context.Context, req *divvy.QueryStreams) (*divvy.QueryStreamsResp, error) {
	ctx, err := unwrapAndCheck(goCtx, req)
	if err != nil {
		return nil, err
	}

	logger := ctx.Logger()
	var as []divvy.SlowReleaseStream
	pageResp, err := query.Paginate(s.streamStore(ctx), req.Pagination, func(key, val []byte) error {
		var a divvy.StoreSlowReleaseStream
		if err := s.cdc.Unmarshal(val, &a); err != nil {
			logger.Error("Can't unmarshal Stream", "err", err)
			return err
		}
		as = append(as, *storeStreamToStream(&a, sdk.AccAddress(key).String()))
		return nil
	})

	return &divvy.QueryStreamsResp{
		Streams:    as,
		Pagination: pageResp,
	}, ormError(err)
}
