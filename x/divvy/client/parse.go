package client

import (
	"encoding/json"
	"fmt"
	"strconv"
	"time"

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
