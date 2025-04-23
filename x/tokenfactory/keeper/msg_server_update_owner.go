package keeper

import (
	"context"

	"tokenfactory/x/tokenfactory/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) UpdateOwner(goCtx context.Context, msg *types.MsgUpdateOwner) (*types.MsgUpdateOwnerResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	valFound, isFound := k.GetDenom(ctx, msg.Denom)
	if !isFound {
		return nil, types.ErrDenomNotExists
	}
	if valFound.Owner != msg.Owner {
		return nil, types.ErrInvalidSigner
	}
	// addressOwerNew, err := sdk.AccAddressFromBech32(msg.NewOwner)
	// if err != nil {
	// 	return nil, err
	// }
	var denom = types.Denom{
		Owner:              msg.NewOwner,
		Denom:              msg.Denom,
		Description:        valFound.Description,
		MaxSupply:          valFound.MaxSupply,
		Supply:             valFound.Supply,
		Precision:          valFound.Precision,
		Ticker:             valFound.Ticker,
		Url:                valFound.Url,
		CanChangeMaxSupply: valFound.CanChangeMaxSupply,
	}
	k.SetDenom(ctx, denom)

	return &types.MsgUpdateOwnerResponse{}, nil
}
