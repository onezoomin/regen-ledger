package client

import (
	"testing"
)

func TestParseTime(t *testing.T) {
	testCases := []string{
		"2021-12-08T01:00:00Z",
		"2021-12-08T01:00:00+02:00",
	}
	for _, tc := range testCases {
		if _, e := parseTime(tc, "startTime", nil); e != nil {
			t.Error(e)
		}
	}
}
