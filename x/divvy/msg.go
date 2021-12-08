package divvy

import (
	"fmt"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/regen-network/regen-ledger/types"
)

var _ sdk.Msg = &MsgCreateAllocator{}
var _ sdk.Msg = &MsgUpdateAllocatorSetting{}
var _ sdk.Msg = &MsgSetAllocatorRecipients{}
var _ sdk.Msg = &MsgRemoveAllocator{}
var _ sdk.Msg = &MsgClaimAllocations{}

var _ sdk.Msg = &MsgCreateSlowReleaseStream{}
var _ sdk.Msg = &MsgPauseSlowReleaseStream{}
var _ sdk.Msg = &MsgEditSlowReleaseStream{}

// GetSigners returns the expected signers for a MsgCreateGroup.
func (msg MsgCreateAllocator) GetSigners() []sdk.AccAddress {
	return getSingers(msg.Admin)
}

// ValidateBasic does a sanity check on the provided data
func (msg MsgCreateAllocator) ValidateBasic() error {
	errs := checkAllocatorTimestamps(msg.Start, msg.End, msg.Interval, msg.Name)
	if _, err := sdk.AccAddressFromBech32(msg.Admin); err != nil {
		errs = append(errs, fmt.Sprintf("Malformed admin address [%s]", err.Error()))
	}
	if err := validateRecipients(msg.Recipients); err != nil {
		errs = append(errs, err.Error())
	}
	return ErrorStringsToError(errs)
}

// Validate makes all additional validation (not present in ValidateBasic)
func (msg MsgCreateAllocator) Validate(ctx types.Context) error {
	t := ctx.BlockTime()
	if msg.End.Before(t) {
		return fmt.Errorf("`end` must be after current block time (%v)", t)
	}
	return nil
}

/****************
  MsgUpdateAllocatorSetting
  /**************/

func (msg MsgUpdateAllocatorSetting) GetSigners() []sdk.AccAddress {
	return getSingers(msg.Sender)
}

func (msg MsgUpdateAllocatorSetting) ValidateBasic() error {
	errs := checkAllocatorTimestamps(msg.Start, msg.End, msg.Interval, msg.Name)
	if _, err := sdk.AccAddressFromBech32(msg.Sender); err != nil {
		errs = append(errs, fmt.Sprintf("Malformed admin address [%s]", err.Error()))
	}
	return ErrorStringsToError(errs)
}

/****************
  MsgSetAllocationMap
  /**************/

func (msg MsgSetAllocatorRecipients) GetSigners() []sdk.AccAddress {
	return getSingers(msg.Sender)
}

func (msg MsgSetAllocatorRecipients) ValidateBasic() error {
	return validateRecipients(msg.Recipients)
}

/****************
  MsgRemoveAllocator
  /**************/

func (msg MsgRemoveAllocator) GetSigners() []sdk.AccAddress {
	return getSingers(msg.Sender)
}

func (msg MsgRemoveAllocator) ValidateBasic() error { return nil }

/****************
  MsgClaimAllocations
  /**************/

func (msg MsgClaimAllocations) GetSigners() []sdk.AccAddress {
	return getSingers(msg.Sender)
}

func (msg MsgClaimAllocations) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Allocator)
	return errors.ErrInvalidAddress.Wrapf("allocator_address [%v]", err)
}

/****************
  MsgCreateSlowReleaseStream
  /**************/

func (msg MsgCreateSlowReleaseStream) GetSigners() []sdk.AccAddress {
	return getSingers(msg.Admin)
}

func (msg MsgCreateSlowReleaseStream) ValidateBasic() error {
	var errs []string
	if msg.Interval < time.Second {
		errs = append(errs, "`interval` must be at least 1s")
	}
	_, err := sdk.AccAddressFromBech32(msg.Destination)
	if err != nil {
		errs = append(errs, fmt.Sprintf("`destination` address is malformed [%v]", err))
	}
	return ErrorStringsToError(errs)
}

/****************
  MsgPauseSlowReleaseStream
  /**************/

func (msg MsgPauseSlowReleaseStream) GetSigners() []sdk.AccAddress {
	return getSingers(msg.Sender)
}

func (msg MsgPauseSlowReleaseStream) ValidateBasic() error {
	return nil
}

/****************
  MsgEditSlowReleaseStream
  /**************/

func (msg MsgEditSlowReleaseStream) GetSigners() []sdk.AccAddress {
	return getSingers(msg.Sender)
}

func (msg MsgEditSlowReleaseStream) ValidateBasic() error {
	var errs []string
	if msg.Interval < time.Second {
		errs = append(errs, "`interval` must be at least 1s")
	}
	_, err := sdk.AccAddressFromBech32(msg.Destination)
	if err != nil {
		errs = append(errs, fmt.Sprintf("`destination` address is malformed [%v]", err))
	}
	return ErrorStringsToError(errs)
}

/*
func (msg ) GetSigners() []sdk.AccAddress {
	return getSingers(msg.Sender)
}

func (msg ) ValidateBasic() error {

}
*/

func checkAllocatorTimestamps(start, end time.Time, interval time.Duration, name string) []string {
	var errs []string
	if !end.After(start) {
		errs = append(errs, "`end` must be after start")
		fmt.Println("start:", start, "..... end", end)
	}
	if interval < time.Second {
		errs = append(errs, "`interval` must be at least 1s")
	}
	if len(name) == 0 {
		errs = append(errs, "`name` must be defined")
	}
	return errs
}

// TotalShares is a number for 100% shares
const TotalShares = 1_000_000

func validateRecipients(entries []Recipient) error {
	var sum uint32 = 0
	for i := range entries {
		if _, err := sdk.AccAddressFromBech32(entries[i].Address); err != nil {
			return fmt.Errorf("Wrong recipient %d, %w", i, err)
		}
		sum += entries[i].Share
	}
	if sum != TotalShares {
		return fmt.Errorf("sum of shares in entries must be %d, got %d", TotalShares, sum)
	}
	return nil
}

func getSingers(signer string) []sdk.AccAddress {
	a, err := sdk.AccAddressFromBech32(signer)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{a}
}

// if  {
// 	errmsgs = append(errmsgs, )
// }
