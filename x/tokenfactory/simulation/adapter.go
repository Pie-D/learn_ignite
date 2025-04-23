package simulation

import (
	"context"

	"tokenfactory/x/tokenfactory/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

type BankKeeperAdapter struct {
	types.BankKeeper
}

func (b BankKeeperAdapter) SpendableCoins(ctx context.Context, addr sdk.AccAddress) sdk.Coins {
	// Chuyển đổi context.Context thành sdk.Context
	sdkCtx := ctx.(sdk.Context)
	return b.BankKeeper.SpendableCoins(sdkCtx, addr)
}
