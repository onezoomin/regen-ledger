package server

import (
	"context"

	"github.com/regen-network/regen-ledger/types"
	"github.com/regen-network/regen-ledger/v2/x/divvy"
)

// Allocator is a distribution engine, which distributes everything which is
// comming in configurable interval periods to registered entries. Each
// allocator has only one owner. Ideally this can be managed by a group
// module.
func (s serverImpl) CreateAllocator(goCtx context.Context, msg *divvy.MsgCreateAllocator) (*divvy.MsgCreateAllocatorResp, error) {
	ctx := types.UnwrapSDKContext(goCtx)
	if err := msg.Validate(ctx); err != nil {
		return nil, err
	}
	addr := nextAddress(s.allocatorSeq, ctx, s.allocatorAddr)
	db := s.getAllocatorStore(ctx)

	bz, err := s.cdc.Marshal(&divvy.StoreAllocator{
		Admin:      msg.Admin,
		Start:      msg.Start,
		End:        msg.End,
		Interval:   msg.Interval,
		Name:       msg.Name,
		Url:        msg.Url,
		Paused:     false,
		Recipients: msg.Recipients,
	})
	if err != nil {
		return nil, err
	}
	db.Set(addr, bz)

	addrStr := addr.String()
	err = ctx.EventManager().EmitTypedEvent(&divvy.EventCreateAllocator{
		Address: addrStr,
	})
	if err != nil {
		return nil, eventError(err)
	}

	return &divvy.MsgCreateAllocatorResp{Address: addrStr}, nil
}

// Updates all allocator settings except admin and entry map.
func (s serverImpl) UpdateAllocatorSetting(goCtx context.Context, msg *divvy.MsgUpdateAllocatorSetting) (*divvy.MsgEmptyResp, error) {
	panic("not implemented") // TODO: Implement
}

// Allocator owner can update the recipient list by setting a new
// allocation map.
func (s serverImpl) SetAllocationMap(goCtx context.Context, msg *divvy.MsgSetAllocationMap) (*divvy.MsgEmptyResp, error) {
	panic("not implemented") // TODO: Implement
}

// Removes allocator and disables all streamers!
func (s serverImpl) RemoveAllocator(goCtx context.Context, msg *divvy.MsgRemoveAllocator) (*divvy.MsgCreateAllocatorResp, error) {
	panic("not implemented") // TODO: Implement
}

// Creates a new stream to feed an address
// User creates a stream. Parameters:
// * % of total amount to be streamed to allocator every second.
// * destination address where the stream will go (ideally allocator)
func (s serverImpl) CreateSlowReleaseStream(goCtx context.Context, msg *divvy.MsgCreateSlowReleaseStream) (*divvy.MsgCreateSlowReleaseStreamResp, error) {
	panic("not implemented") // TODO: Implement
}

func (s serverImpl) PauseSlowReleaseStream(goCtx context.Context, msg *divvy.MsgPauseSlowReleaseStream) (*divvy.MsgEmptyResp, error) {
	panic("not implemented") // TODO: Implement
}

func (s serverImpl) EditSlowReleaseStream(goCtx context.Context, msg *divvy.MsgEditSlowReleaseStream) (*divvy.MsgEmptyResp, error) {
	panic("not implemented") // TODO: Implement
}
