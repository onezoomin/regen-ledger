package divvy

import (
	"errors"
	"fmt"
	"time"

	"github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	proto "github.com/gogo/protobuf/proto"
)

var _ sdk.Msg = &MsgCreateAllocator{}

// var _ sdk.Msg = &MsgUpdateAllocatorSetting{}}
// var _ sdk.Msg = &MsgSetAllocationMap{}
// var _ sdk.Msg = &MsgRemoveAllocator{}
// var _ sdk.Msg = &MsgCreateSlowReleaseStream{}
// var _ sdk.Msg = &MsgPauseSlowReleaseStream{}
// var _ sdk.Msg = &MsgEditSlowReleaseStream{}

// GetSigners returns the expected signers for a MsgCreateGroup.
func (m MsgCreateAllocator) GetSigners() []sdk.AccAddress {
	admin, err := sdk.AccAddressFromBech32(m.Admin)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{admin}
}

// ValidateBasic does a sanity check on the provided data
func (m MsgCreateAllocator) ValidateBasic() error {
	var errmsgs []string
	_, err := sdk.AccAddressFromBech32(m.Admin)
	if err != nil {
		errmsgs = append(errmsgs, fmt.Sprintf("Malformed admin address [%s]", err.Error()))
	}
	if !m.End.After(m.Start) {
		errmsgs = append(errmsgs, "`end` must be after start")
	}
	if m.Interval < time.Second {
		errmsgs = append(errmsgs, "`interval` must be at least 1s")
	}
	if len(m.Name) == 0 {
		errmsgs = append(errmsgs, "`name` must be defined")
	}
	if err := ValidateEntries(m.Entries); err != nil {
		errmsgs = append(errmsgs, err.Error())
	}
	return errorStringsToError(errmsgs)
}

func (m MsgCreateAllocator) ValidateExtra(ctx sdk.Context) error {
	t := ctx.BlockTime()
	if m.End.Before(t) {
		return fmt.Errorf("`end` must be after current block time (%v)", t)
	}
}

func ValidateEntries(entries []Recipient) error {
	const expected = 1_000_000
	var sum uint32 = 0
	for i := range entries {
		sum += entries[i].Share
	}
	if sum != expected {
		return fmt.Errorf("sum of shares in entries must be %d, got %d", expected, sum)
	}
	return nil
}

// if  {
// 	errmsgs = append(errmsgs, )
// }
