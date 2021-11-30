package server

import (
	"context"

	// "github.com/regen-network/regen-ledger/v2/x/divvy"
	"github.com/regen-network/regen-ledger/v2/x/divvy"
)

var s divvy.MsgServer

// Allocator is a distribution engine, which distributes everything which is
// comming in configurable interval periods to registered entries. Each
// allocator has only one owner. Ideally this can be managed by a group
// module.
func (s serverImpl) CreateAllocator(ctx context.Context, msg *divvy.MsgCreateAllocator) (*divvy.MsgCreateAllocatorResp, error) {
	panic("not implemented") // TODO: Implement
}

// Updates all allocator settings except admin and entry map.
func (s serverImpl) UpdateAllocatorSetting(ctx context.Context, msg *divvy.MsgUpdateAllocatorSetting) (*divvy.MsgEmptyResp, error) {
	panic("not implemented") // TODO: Implement
}

// Allocator owner can update the recipient list by setting a new
// allocation map.
func (s serverImpl) SetAllocationMap(ctx context.Context, msg *divvy.MsgSetAllocationMap) (*divvy.MsgEmptyResp, error) {
	panic("not implemented") // TODO: Implement
}

// Removes allocator and disables all streamers!
func (s serverImpl) RemoveAllocator(ctx context.Context, msg *divvy.MsgRemoveAllocator) (*divvy.MsgCreateAllocatorResp, error) {
	panic("not implemented") // TODO: Implement
}

// Creates a new stream to feed an address
// User creates a stream. Parameters:
// * % of total amount to be streamed to allocator every second.
// * destination address where the stream will go (ideally allocator)
func (s serverImpl) CreateSlowReleaseStream(ctx context.Context, msg *divvy.MsgCreateSlowReleaseStream) (*divvy.MsgEmptyResp, error) {
	panic("not implemented") // TODO: Implement
}

func (s serverImpl) PauseSlowReleaseStream(ctx context.Context, msg *divvy.MsgPauseSlowReleaseStream) (*divvy.MsgEmptyResp, error) {
	panic("not implemented") // TODO: Implement
}

func (s serverImpl) EditSlowReleaseStream(ctx context.Context, msg *divvy.MsgEditSlowReleaseStream) (*divvy.MsgEmptyResp, error) {
	panic("not implemented") // TODO: Implement
}
