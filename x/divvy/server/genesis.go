package server

import (
	"encoding/json"

	"github.com/cosmos/cosmos-sdk/codec"
	// sdk "github.com/cosmos/cosmos-sdk/types"
	// sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	abci "github.com/tendermint/tendermint/abci/types"

	"github.com/regen-network/regen-ledger/types"
)

func (s serverImpl) InitGenesis(ctx types.Context, cdc codec.Codec, data json.RawMessage) ([]abci.ValidatorUpdate, error) {
	// TODO: no need for genesis?

	return []abci.ValidatorUpdate{}, nil
}

// ExportGenesis will dump the divvy module state into a serializable GenesisState.
func (s serverImpl) ExportGenesis(ctx types.Context, cdc codec.Codec) (json.RawMessage, error) {
	// TODO: no export genensis
	return nil, nil
}
