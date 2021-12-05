package module

import (
	// "context"
	"encoding/json"

	sdkclient "github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/gorilla/mux"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"github.com/spf13/cobra"

	climodule "github.com/regen-network/regen-ledger/types/module/client/cli"
	restmodule "github.com/regen-network/regen-ledger/types/module/client/grpc_gateway"
	servermodule "github.com/regen-network/regen-ledger/types/module/server"
	"github.com/regen-network/regen-ledger/v2/x/divvy"
	"github.com/regen-network/regen-ledger/v2/x/divvy/client"
	"github.com/regen-network/regen-ledger/v2/x/divvy/server"
	// "github.com/regen-network/regen-ledger/v2/x/divvy/simulation"
)

type Module struct {
	bank divvy.BankKeeper
}

// NewModule returns a new Module object.
func NewModule(b divvy.BankKeeper) Module {
	return Module{
		bank: b,
	}
}

var _ module.AppModuleBasic = Module{}
var _ servermodule.Module = Module{}
var _ restmodule.Module = Module{}
var _ climodule.Module = Module{}

// TODO:
//var _ module.AppModuleSimulation = Module{}

func (a Module) Name() string {
	return divvy.ModuleName
}

func (a Module) RegisterInterfaces(registry types.InterfaceRegistry) {
	divvy.RegisterTypes(registry)
}

func (a Module) RegisterServices(configurator servermodule.Configurator) {
	if err := server.RegisterServices(configurator, a.bank); err != nil {
		panic(err)
	}
}

func (a Module) RegisterGRPCGatewayRoutes(clientCtx sdkclient.Context, mux *runtime.ServeMux) {
	// TODO:
	// divvy.RegisterQueryHandlerClient(context.Background(), mux, divvy.NewQueryClient(clientCtx))
}

func (a Module) DefaultGenesis(cdc codec.JSONCodec) json.RawMessage {
	return nil
	// return cdc.MustMarshalJSON(divvy.DefaultGenesisState())
}

func (a Module) ValidateGenesis(cdc codec.JSONCodec, _ sdkclient.TxEncodingConfig, bz json.RawMessage) error {
	return nil
	// var data divvy.GenesisState
	// if err := cdc.UnmarshalJSON(bz, &data); err != nil {
	// 	return fmt.Errorf("failed to unmarshal %s genesis state: %w", divvy.ModuleName, err)
	// }

	// return data.Validate()
}

func (a Module) GetQueryCmd() *cobra.Command {
	return client.QueryCmd(a.Name())
}

func (a Module) GetTxCmd() *cobra.Command {
	return client.TxCmd(a.Name())
}

// ConsensusVersion implements AppModule/ConsensusVersion.
func (Module) ConsensusVersion() uint64 { return 1 }

/**** DEPRECATED ****/
func (a Module) RegisterRESTRoutes(sdkclient.Context, *mux.Router) {}
func (a Module) RegisterLegacyAminoCodec(cdc *codec.LegacyAmino)   {}

// AppModuleSimulation functions

// GenerateGenesisState creates a randomized GenesisState of the divvy module.
// func (Module) GenerateGenesisState(simState *module.SimulationState) {
// 	simulation.RandomizedGenState(simState)
// }

// ProposalContents returns all the divvy content functions used to
// simulate proposals.
func (Module) ProposalContents(simState module.SimulationState) []simtypes.WeightedProposalContent {
	return nil
}

// // RandomizedParams creates randomized divvy param changes for the simulator.
// func (Module) RandomizedParams(r *rand.Rand) []simtypes.ParamChange {
// 	return simulation.ParamChanges(r)
// }

// RegisterStoreDecoder registers a decoder for divvy module's types
func (Module) RegisterStoreDecoder(sdr sdk.StoreDecoderRegistry) {
}

// WeightedOperations returns all the divvy module operations with their respective weights.
// NOTE: This is no longer needed for the modules which uses ADR-33, divvy module `WeightedOperations`
// registered in the `x/divvy/server` package.
func (Module) WeightedOperations(simState module.SimulationState) []simtypes.WeightedOperation {
	return nil
}
