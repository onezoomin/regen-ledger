package parse

import (
	"testing"
)

func TestParseTime(t *testing.T) {
	testCases := []string{
		"2021-12-08T01:00:00Z",
		"2021-12-08T01:00:00+02:00",
	}
	for _, tc := range testCases {
		if _, e := Time(tc, "startTime", nil); e != nil {
			t.Error(e)
		}
	}
}

func TestParseRecipients(t *testing.T) {
	testCases := []string{
		`[]`,
		`[{"address": "regen1ah7qh4af2fa8d0h4aef9sxqqec8wkmqaxtxa0r", "share": 1000000, "name": "a1"}]`,
	}
	for i, tc := range testCases {
		if _, e := Recipient(tc, nil); e != nil {
			t.Error("test case:", i, e)
		}
	}
}
