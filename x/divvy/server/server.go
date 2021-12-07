package server

import (
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/address"

	"github.com/regen-network/regen-ledger/orm"
	"github.com/regen-network/regen-ledger/types/module/server"
	"github.com/regen-network/regen-ledger/v2/x/divvy"
	"github.com/regen-network/regen-ledger/x/group/exported"
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

type serverImpl struct {
	key        sdk.StoreKey
	cdc        codec.Codec
	bankKeeper exported.BankKeeper

	allocatorSeq orm.Sequence
	streamSeq    orm.Sequence

	// module addresses for deriving entity addresses
	allocatorAddr sdk.Address
	streamAddr    sdk.Address
}

func newServer(storeKey server.RootModuleKey, bank exported.BankKeeper, cdc codec.Codec) (serverImpl, error) {
	s := serverImpl{key: storeKey, cdc: cdc, bankKeeper: bank}
	s.allocatorSeq = orm.NewSequence(storeKey, allocatorTableSeqPrefix)
	s.streamSeq = orm.NewSequence(storeKey, streamTableSeqPrefix)

	s.allocatorAddr = sdk.AccAddress(address.Module(divvy.ModuleName, allocatorTablePrefix))
	s.streamAddr = sdk.AccAddress(address.Module(divvy.ModuleName, streamTablePrefix))

	return s, nil
}

func RegisterServices(configurator server.Configurator, bank divvy.BankKeeper) error {
	impl, err := newServer(configurator.ModuleKey(), bank, configurator.Marshaler())
	if err != nil {
		return err
	}
	divvy.RegisterMsgServer(configurator.MsgServer(), impl)
	divvy.RegisterQueryServer(configurator.QueryServer(), impl)
	configurator.RegisterGenesisHandlers(impl.InitGenesis, impl.ExportGenesis)
	// TODO:
	// configurator.RegisterWeightedOperationsHandler(impl.WeightedOperations)
	// configurator.RegisterInvariantsHandler(impl.RegisterInvariants)
	return nil
}
