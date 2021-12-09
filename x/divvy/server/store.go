package server

import (
	"encoding/binary"
	"fmt"

	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/address"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"github.com/regen-network/regen-ledger/orm"
	"github.com/regen-network/regen-ledger/v2/x/divvy"
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

type storeKey []byte

type Sequence interface {
	NextVal(ctx orm.HasKVStore) uint64
}

func nextAllocatorAddress(seq Sequence, ctx orm.HasKVStore, parentAddr sdk.Address) sdk.AccAddress {
	id := seq.NextVal(ctx)
	var idbz = make([]byte, 8)
	binary.BigEndian.PutUint64(idbz, id)

	return address.Derive(parentAddr.Bytes(), idbz)
}

func (s serverImpl) allocatorStore(ctx orm.HasKVStore) sdk.KVStore {
	d := ctx.KVStore(s.key)
	return prefix.NewStore(d, allocatorTablePrefix)
}

func (s serverImpl) streamStore(ctx orm.HasKVStore) sdk.KVStore {
	d := ctx.KVStore(s.key)
	return prefix.NewStore(d, streamTablePrefix)
}

func (s serverImpl) getAllocator(ctx orm.HasKVStore, address string) (sdk.AccAddress, storeKey, *divvy.StoreAllocator, error) {
	key, err := getAllocatorKey(address)
	if err != nil {
		return nil, nil, nil, err
	}
	bz := ctx.KVStore(s.key).Get(key)
	if bz == nil {
		return nil, nil, nil, sdkerrors.ErrNotFound.Wrapf("key not found: %q", key)
	}
	var a divvy.StoreAllocator
	return key, key, &a, s.cdc.Unmarshal(bz, &a)
}

// Creates a fully qualified (with table prefix) store key based on allocator bech32 address
func getAllocatorKey(allocatorAddr string) ([]byte, error) {
	return subKeyFromAddr(allocatorTablePrefix, allocatorAddr)
}

func subKeyFromAddr(prefix []byte, addr string) ([]byte, error) {
	a, err := sdk.AccAddressFromBech32(addr)
	if err != nil {
		return nil, err
	}
	key := make([]byte, 0, len(prefix)+len(a))
	key = append(key, prefix...)
	return append(key, a...), nil
}

func save(db sdk.KVStore, key storeKey, o codec.ProtoMarshaler, cdc codec.BinaryCodec) error {
	bz, err := cdc.Marshal(o)
	if err != nil {
		return err
	}
	db.Set(key, bz)
	return nil
}
