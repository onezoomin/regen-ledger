package server

import (
	"context"

	"github.com/cosmos/cosmos-sdk/types/errors"

	"github.com/regen-network/regen-ledger/types"
	"github.com/regen-network/regen-ledger/v2/x/divvy"
)

// Allocator is a distribution engine, which distributes everything which is
// comming in configurable interval periods to registered entries. Each
// allocator has only one owner. Ideally this can be managed by a group
// module.
func (s serverImpl) CreateAllocator(goCtx context.Context, msg *divvy.MsgCreateAllocator) (*divvy.MsgCreateAllocatorResp, error) {
	ctx := types.UnwrapSDKContext(goCtx)
	err := msg.Validate(ctx)
	if err != nil {
		return nil, err
	}
	addr := nextAllocatorAddress(s.allocatorSeq, ctx, s.allocatorAddr)
	a := divvy.StoreAllocator{
		Admin:    msg.Admin,
		Start:    msg.Start,
		End:      msg.End,
		Interval: msg.Interval,
		Name:     msg.Name,
		Url:      msg.Url,
		Paused:   false,
	}
	a.Recipients, err = recipientsToStoreRecipients(msg.Recipients)
	if err != nil {
		return nil, err
	}
	err = save(s.allocatorStore(ctx), storeKey(addr), &a, s.cdc)
	if err != nil {
		return nil, err
	}

	addrStr := addr.String()
	err = ctx.EventManager().EmitTypedEvent(&divvy.EventCreateAllocator{
		Address: addrStr,
	})
	if err != nil {
		return nil, eventError(err)
	}

	return &divvy.MsgCreateAllocatorResp{Address: addrStr}, nil
}

// anyone can claim allocatinos for registered recipient
func (s serverImpl) ClaimAllocations(goCtx context.Context, msg *divvy.MsgClaimAllocations) (*divvy.MsgClaimAllocationsResp, error) {
	ctx, err := unwrapAndCheck(goCtx, msg)
	if err != nil {
		return nil, err
	}
	addr, key, a, err := s.getAllocator(ctx, msg.Allocator)
	now := ctx.BlockTime()
	if now.Before(a.NextClaim) {
		return nil, errors.ErrInvalidRequest.Wrapf("Claim only possible after %v", a.NextClaim)
	}
	a.NextClaim = now.Add(a.Interval)
	coins, err := distributeBalance(addr, a, s.bank, &ctx.Context)
	if err != nil {
		return nil, err
	}
	if err = save(ctx.KVStore(s.key), key, a, s.cdc); err != nil {
		return nil, err
	}

	return &divvy.MsgClaimAllocationsResp{Coins: coins}, err
}

// Updates all allocator settings except admin and entry map.
func (s serverImpl) UpdateAllocatorSettings(goCtx context.Context, msg *divvy.MsgUpdateAllocatorSettings) (*divvy.MsgEmptyResp, error) {
	ctx, err := unwrapAndCheck(goCtx, msg)
	if err != nil {
		return nil, err
	}
	if err = msg.Validate(ctx); err != nil {
		return nil, err
	}

	_, key, a, err := s.getAllocator(ctx, msg.Address)
	if err != nil {
		return nil, err
	}
	if err = assertAllocatorAdmin(msg.Sender, a); err != nil {
		return nil, err
	}
	a.Start = msg.Start
	a.End = msg.End
	a.Interval = msg.Interval
	a.Name = msg.Name
	a.Url = msg.Url
	err = save(ctx.KVStore(s.key), key, a, s.cdc)
	return &divvy.MsgEmptyResp{}, err
}

// Allocator owner can update the recipient list by setting a new
// allocation map.
func (s serverImpl) SetAllocatorRecipients(goCtx context.Context, msg *divvy.MsgSetAllocatorRecipients) (*divvy.MsgEmptyResp, error) {
	ctx, err := unwrapAndCheck(goCtx, msg)
	if err != nil {
		return nil, err
	}
	_, key, a, err := s.getAllocator(ctx, msg.Address)
	if err != nil {
		return nil, err
	}
	a.Recipients, err = recipientsToStoreRecipients(msg.Recipients)
	if err != nil {
		return nil, err
	}
	return &divvy.MsgEmptyResp{}, save(ctx.KVStore(s.key), key, a, s.cdc)
}

// Removes allocator and disables all streamers!
func (s serverImpl) RemoveAllocator(goCtx context.Context, msg *divvy.MsgRemoveAllocator) (*divvy.MsgEmptyResp, error) {
	ctx, err := unwrapAndCheck(goCtx, msg)
	if err != nil {
		return nil, err
	}
	_, key, a, err := s.getAllocator(ctx, msg.Address)
	if err != nil {
		return nil, err
	}
	if err = assertAllocatorAdmin(msg.Sender, a); err != nil {
		return nil, err
	}
	ctx.KVStore(s.key).Delete(key)
	return &divvy.MsgEmptyResp{}, nil

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
