package server

import (
	"context"

	// "github.com/cosmos/cosmos-sdk/types/errors"

	"github.com/regen-network/regen-ledger/types"
	"github.com/regen-network/regen-ledger/v2/x/divvy"
)

// Creates a new stream to feed an address
// User creates a stream. Parameters:
// * % of total amount to be streamed to allocator every second.
// * destination address where the stream will go (ideally allocator)
func (s serverImpl) CreateSlowReleaseStream(goCtx context.Context, msg *divvy.MsgCreateSlowReleaseStream) (*divvy.MsgCreateSlowReleaseStreamResp, error) {
	ctx := types.UnwrapSDKContext(goCtx)
	addr := nextSeqBasedAddr(s.streamSeq, ctx, s.allocatorAddr)
	x := divvy.SlowReleaseStream{
		Admin:       msg.Admin,
		Start:       msg.Start,
		Interval:    msg.Interval,
		Destination: msg.Destination,
		Name:        msg.Name,
		Paused:      false,
		Strategy:    msg.Strategy,
	}
	err := save(s.streamStore(ctx), storeKey(addr), &x, s.cdc)
	if err != nil {
		return nil, err
	}

	addrStr := addr.String()
	err = ctx.EventManager().EmitTypedEvent(&divvy.EventCreateStream{
		Address: addrStr,
	})
	if err != nil {
		return nil, eventError(err)
	}

	return &divvy.MsgCreateSlowReleaseStreamResp{Address: addrStr}, nil
}

func (s serverImpl) PauseSlowReleaseStream(goCtx context.Context, msg *divvy.MsgPauseSlowReleaseStream) (*divvy.MsgEmptyResp, error) {
	panic("not implemented") // TODO: Implement
}

func (s serverImpl) EditSlowReleaseStream(goCtx context.Context, msg *divvy.MsgEditSlowReleaseStream) (*divvy.MsgEmptyResp, error) {
	panic("not implemented") // TODO: Implement
}
