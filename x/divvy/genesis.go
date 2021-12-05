package divvy

type GenesisState struct{}

// DefaultGenesisState returns a default divvy module genesis state.
func DefaultGenesisState() *GenesisState {
	return &GenesisState{}
}
