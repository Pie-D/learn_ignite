package simulation

import (
	"context"

	"tokenfactory/x/tokenfactory/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	// authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	// simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
)

type BankKeeperAdapter struct {
	types.BankKeeper
}

func (b BankKeeperAdapter) SpendableCoins(ctx context.Context, addr sdk.AccAddress) sdk.Coins {
	// Chuyển đổi context.Context thành sdk.Context
	sdkCtx := ctx.(sdk.Context)
	return b.BankKeeper.SpendableCoins(sdkCtx, addr)
}

// Adapter để khớp interface với simulation.AccountKeeper
type AccountKeeperAdapter struct {
	AK types.AccountKeeper
}

// Convert std context về sdk.Context (ngược với mong muốn, nhưng vì signature yêu cầu)
func (a AccountKeeperAdapter) GetAccount(ctx context.Context, addr sdk.AccAddress) sdk.AccountI {
	sdkCtx := sdk.UnwrapSDKContext(ctx)
	return a.AK.GetAccount(sdkCtx, addr)

}
