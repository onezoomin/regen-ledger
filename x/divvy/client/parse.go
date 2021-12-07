package client

import (
	"encoding/json"
	"fmt"
	"strconv"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/errors"

	"github.com/regen-network/regen-ledger/v2/x/divvy"
)

func parseTime(s, paramName string, errmsgs []string) (time.Time, []string) {
	t, err := time.Parse(time.RFC3339, s)
	if err != nil {
		errmsgs = append(errmsgs, fmt.Sprintf("wrong %s format, RFC3339 is expected [%v]", paramName, err))
	}
	return t, errmsgs
}

func parseUint(s, paramName string, errmsgs []string) (uint64, []string) {
	i, err := strconv.ParseUint(s, 10, 64)
	if err != nil {
		errmsgs = append(errmsgs, fmt.Sprintf("wrong %s, a positive integer is expected", paramName))
	}
	return uint64(i), errmsgs
}

func parseRecipient(s string, errmsg []string) ([]divvy.Recipient, []string) {
	recipients := []divvy.Recipient{}
	if err := json.Unmarshal([]byte(s), &recipients); err != nil {
		errmsg = append(errmsg, fmt.Sprintf("invalid recipient map: %v", err))
	}
	return recipients, errmsg
}

func parseAddress(s, field string) error {
	if _, err := sdk.AccAddressFromBech32(s); err != nil {
		return errors.ErrInvalidAddress.Wrapf("%q is not a valid %s address", s, field)
	}
	return nil
}
