package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/cosmos/cosmos-sdk/x/auth/migrations/legacytx"
)

// verify interface at compile time
var (
	_ sdk.Msg = &MsgUnjail{}
	_ sdk.Msg = &MsgUpdateParams{}
	_ sdk.Msg = &MsgImpeach{}

	_ legacytx.LegacyMsg = &MsgUnjail{}
	_ legacytx.LegacyMsg = &MsgUpdateParams{}
)

// NewMsgUnjail creates a new MsgUnjail instance
func NewMsgUnjail(validatorAddr sdk.AccAddress) *MsgUnjail {
	return &MsgUnjail{
		ValidatorAddr: validatorAddr.String(),
	}
}

// GetSigners returns the expected signers for MsgUnjail.
func (msg MsgUnjail) GetSigners() []sdk.AccAddress {
	valAddr, _ := sdk.AccAddressFromHexUnsafe(msg.ValidatorAddr)
	return []sdk.AccAddress{valAddr}
}

// GetSignBytes gets the bytes for the message signer to sign on
func (msg MsgUnjail) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(&msg)
	return sdk.MustSortJSON(bz)
}

// GetSignBytes implements the LegacyMsg interface.
func (msg MsgUpdateParams) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(&msg))
}

// GetSigners returns the expected signers for a MsgUpdateParams message.
func (msg MsgUpdateParams) GetSigners() []sdk.AccAddress {
	addr, _ := sdk.AccAddressFromHexUnsafe(msg.Authority)
	return []sdk.AccAddress{addr}
}

// NewMsgImpeach creates a new MsgImpeach instance
func NewMsgImpeach(valAddr, from sdk.AccAddress) *MsgImpeach {
	return &MsgImpeach{
		ValidatorAddress: valAddr.String(),
		From:             from.String(),
	}
}

// GetSigners implements the sdk.Msg interface.
func (msg MsgImpeach) GetSigners() []sdk.AccAddress {
	fromAddr, _ := sdk.AccAddressFromHexUnsafe(msg.From)
	return []sdk.AccAddress{fromAddr}
}

// GetSignBytes implements the sdk.Msg interface.
func (msg MsgImpeach) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(&msg)
	return sdk.MustSortJSON(bz)
}

// ValidateBasic implements the sdk.Msg interface.
func (msg MsgImpeach) ValidateBasic() error {
	if _, err := sdk.AccAddressFromHexUnsafe(msg.From); err != nil {
		return sdkerrors.ErrInvalidAddress.Wrapf("invalid account address: %s", err)
	}

	if _, err := sdk.AccAddressFromHexUnsafe(msg.ValidatorAddress); err != nil {
		return sdkerrors.ErrInvalidAddress.Wrapf("invalid validator address: %s", err)
	}

	return nil
}
