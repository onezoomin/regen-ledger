package server

import (
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/regen-network/regen-ledger/orm"
	servermodule "github.com/regen-network/regen-ledger/types/module/server"
	"github.com/regen-network/regen-ledger/v2/x/divvy"
	"github.com/regen-network/regen-ledger/x/group/exported"
)

// storage key prefixes
const (
	allocatorTablePrefix    byte = 0x0
	allocatorTableSeqPrefix byte = 0x1

	streamTablePrefix    byte = 0x2
	streamTableSeqPrefix byte = 0x3
)

type serverImpl struct {
	key        sdk.StoreKey
	bankKeeper exported.BankKeeper

	allocatorTable orm.AutoUInt64Table
	streamTable    orm.AutoUInt64Table
}

func newServer(storeKey servermodule.RootModuleKey, bankKeeper exported.BankKeeper, cdc codec.Codec) (serverImpl, error) {
	s := serverImpl{key: storeKey, bankKeeper: bankKeeper}
	allocatorTable, err := orm.NewAutoUInt64TableBuilder(allocatorTablePrefix, allocatorTableSeqPrefix, storeKey, &divvy.Allocator{}, cdc)
	if err != nil {
		return s, err
	}
	s.allocatorTable = allocatorTable.Build()

	streamTable, err := orm.NewAutoUInt64TableBuilder(streamTablePrefix, streamTableSeqPrefix, storeKey, &divvy.SlowReleaseStream{}, cdc)
	if err != nil {
		return s, err
	}
	s.streamTable = streamTable.Build()

	return s, nil
}
