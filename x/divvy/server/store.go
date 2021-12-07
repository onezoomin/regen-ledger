package server

import (
	"encoding/binary"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/address"

	"github.com/regen-network/regen-ledger/orm"
)

// storage key prefixes
const (
	allocatorTableSeqPrefix byte = 0x0
	streamTableSeqPrefix    byte = 0x1
)

var (
	allocatorTablePrefix = []byte{0x2}
	streamTablePrefix    = []byte{0x3}
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

func getAllocatorKey(allocator string) ([]byte, error) {
	return getKey(allocatorTablePrefix, allocator)
}

func getKey(prefix []byte, addr string) ([]byte, error) {
	a, err := sdk.AccAddressFromBech32(addr)
	if err != nil {
		return nil, err
	}
	key := make([]byte, 0, len(prefix)+len(a))
	key = append(key, prefix...)
	return append(key, a...), nil
}
