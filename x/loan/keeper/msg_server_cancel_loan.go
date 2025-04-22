package keeper

import (
	"context"
	"loan/x/loan/types"

	errorsmod "cosmossdk.io/errors"

	// errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (k msgServer) CancelLoan(goCtx context.Context, msg *types.MsgCancelLoan) (*types.MsgCancelLoanResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	loan, found := k.GetLoan(ctx, msg.Id)
	if !found {
		return nil, errorsmod.Wrapf(sdkerrors.ErrKeyNotFound, "loan with id %d not found", msg.Id)
	}
	if loan.Borrower != msg.Creator {
		return nil, errorsmod.Wrapf(sdkerrors.ErrUnauthorized, "incorrect borrower address: %s", msg.Creator)
	}
	if loan.State != "requested" {
		return nil, errorsmod.Wrapf(types.ErrWrongLoanState, "loan with id %v is not in requested state", loan.State)
	}
	borrower, _ := sdk.AccAddressFromBech32(loan.Borrower)
	collateral, _ := sdk.ParseCoinsNormalized(loan.Collateral)
	err := k.bankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, borrower, collateral)
	if err != nil {
		return nil, err
	}
	loan.State = "cancelled"
	k.SetLoan(ctx, loan)

	return &types.MsgCancelLoanResponse{}, nil
}
