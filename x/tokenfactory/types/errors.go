package types

// DONTCOVER

import (
	sdkerrors "cosmossdk.io/errors"
)

// x/tokenfactory module sentinel errors
var (
	ErrInvalidSigner    = sdkerrors.Register(ModuleName, 1100, "expected gov account as only signer for proposal message")
	ErrSample           = sdkerrors.Register(ModuleName, 1101, "sample error")
	ErrDenomExists      = sdkerrors.Register(ModuleName, 1102, "denom already exists")
	ErrDenomNotExists   = sdkerrors.Register(ModuleName, 1103, "denom not exists")
	ErrExceedsMaxSupply = sdkerrors.Register(ModuleName, 1104, "exceeds max supply")
)
