package server

import (
	"encoding/binary"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/address"

	"github.com/regen-network/regen-ledger/orm"
)

type Sequence interface {
	NextVal(ctx orm.HasKVStore) uint64
}

func nextAddress(seq Sequence, ctx orm.HasKVStore, parentAddr sdk.Address) sdk.AccAddress {
	id := seq.NextVal(ctx)
	var idbz = make([]byte, 8)
	binary.BigEndian.PutUint64(idbz, id)

	return address.Derive(parentAddr.Bytes(), idbz)
}

func (s serverImpl) getAllocatorStore(ctx orm.HasKVStore) sdk.KVStore {
	d := ctx.KVStore(s.key)
	return prefix.NewStore(d, allocatorTablePrefix)
}

func (s serverImpl) getStreamStore(ctx orm.HasKVStore) sdk.KVStore {
	d := ctx.KVStore(s.key)
	return prefix.NewStore(d, streamTablePrefix)
}
