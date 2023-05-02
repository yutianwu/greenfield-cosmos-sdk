package nft

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

const (
	// TypeMsgSend nft message types
	TypeMsgSend = "send"
)

var _ sdk.Msg = &MsgSend{}

// GetSigners returns the expected signers for MsgSend.
func (m MsgSend) GetSigners() []sdk.AccAddress {
	signer, _ := sdk.AccAddressFromHexUnsafe(m.Sender)
	return []sdk.AccAddress{signer}
}
