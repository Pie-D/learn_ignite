package keeper

import (
	"context"

	"loan/x/loan/types"

	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (k msgServer) RepayLoan(goCtx context.Context, msg *types.MsgRepayLoan) (*types.MsgRepayLoanResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	loan, found := k.GetLoan(ctx, msg.Id)
	if !found {
		return nil, errorsmod.Wrapf(sdkerrors.ErrNotFound, "Loan %d not found", msg.Id)
	}
	if loan.State != "approved" {
		return nil, errorsmod.Wrapf(types.ErrWrongLoanState, "Sate loan is %v", loan.State)
	}
	lender, _ := sdk.AccAddressFromBech32(loan.Lender)
	borrower, _ := sdk.AccAddressFromBech32(loan.Borrower)
	if loan.Borrower != msg.Creator {
		return nil, errorsmod.Wrapf(sdkerrors.ErrUnauthorized, "Incorrect borrower address %s", msg.Creator)
	}
	amount, _ := sdk.ParseCoinsNormalized(loan.Amount)
	fee, _ := sdk.ParseCoinsNormalized(loan.Fee)
	collateral, _ := sdk.ParseCoinsNormalized(loan.Collateral)
	err := k.bankKeeper.SendCoins(ctx, borrower, lender, amount)
	if err != nil {
		return nil, errorsmod.Wrapf(sdkerrors.ErrInsufficientFunds, "Failed to send coins from %s to %s: %v", borrower.String(), lender.String(), err)
	}
	err = k.bankKeeper.SendCoins(ctx, borrower, lender, fee)
	if err != nil {
		return nil, errorsmod.Wrapf(sdkerrors.ErrInsufficientFunds, "Failed to send coins from %s to %s: %v", borrower.String(), lender.String(), err)
	}
	err = k.bankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, borrower, collateral)
	if err != nil {
		return nil, errorsmod.Wrapf(sdkerrors.ErrInsufficientFunds, "Failed to send coins from %v to %s: %v", types.ModuleName, borrower.String(), err)
	}
	return &types.MsgRepayLoanResponse{}, nil
}
