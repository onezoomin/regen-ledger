package server

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/regen-network/regen-ledger/types"
	"github.com/regen-network/regen-ledger/v2/x/divvy"
)

func (s serverImpl) Allocators(goCtx context.Context, req *divvy.QueryAllocators) (*divvy.QueryAllocatorsResp, error) {
	ctx, err := unwrapAndCheck(goCtx, req)
	if err != nil {
		return nil, err
	}

	var as []divvy.Allocator
	pageResp, err := query.Paginate(s.getAllocatorStore(ctx), req.Pagination, func(key, val []byte) error {
		var a divvy.StoreAllocator
		if err := s.cdc.Unmarshal(val, &a); err != nil {
			return err
		}
		as = append(as, divvy.Allocator{Address: sdk.AccAddress(key).String(), A: a})
		return nil
	})

	return &divvy.QueryAllocatorsResp{
		Allocator:  as,
		Pagination: pageResp,
	}, ormError(err)
}

func (s serverImpl) AllocatorsByOwner(goCtx context.Context, req *divvy.QueryAllocatorsByOwner) (*divvy.QueryAllocatorsResp, error) {
	panic("not implemented") // TODO: Implement
}

func errEmptyRequest() error {
	return status.Errorf(codes.InvalidArgument, "empty request")
}

func unwrapAndCheck(ctx context.Context, req interface{}) (types.Context, error) {
	if req == nil {
		return types.Context{}, errEmptyRequest()
	}
	return types.UnwrapSDKContext(ctx), nil
}
