package divvy

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// BankKeeper defines the expected interface needed to burn and send coins and to retrieve account balances.
type BankKeeper interface {
	// BurnCoins(ctx sdk.Context, moduleName string, amt sdk.Coins) error
	SpendableCoins(ctx sdk.Context, addr sdk.AccAddress) sdk.Coins
	SendCoinsFromAccountToModule(ctx sdk.Context, senderAddr sdk.AccAddress, recipientModule string, amt sdk.Coins) error
	SendCoins(ctx sdk.Context, fromAddr sdk.AccAddress, toAddr sdk.AccAddress, amt sdk.Coins) error

	// views

	GetAllBalances(ctx sdk.Context, addr sdk.AccAddress) sdk.Coins
}
