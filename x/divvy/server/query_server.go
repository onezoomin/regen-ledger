package server

import (
	"context"
	"math"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/regen-network/regen-ledger/orm"
	"github.com/regen-network/regen-ledger/types"
	"github.com/regen-network/regen-ledger/v2/x/divvy"
)

func (s serverImpl) Allocators(goCtx context.Context, req *divvy.QueryAllocators) (*divvy.QueryAllocatorsResp, error) {
	if req == nil {
		return nil, errEmptyRequest()
	}
	ctx := types.UnwrapSDKContext(goCtx)
	iter, err := s.allocatorTable.PrefixScan(ctx, 0, math.MaxUint64)
	if err != nil {
		return nil, ormError(err)
	}

	var as []divvy.Allocator
	pageResp, err := orm.Paginate(iter, req.Pagination, &as)

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
