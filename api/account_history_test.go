package api

import (
	"testing"
)

func TestGetGetAccountHistory(t *testing.T) {
	resp, err := GetAccountHistory(client, "steemit", 10, 10)
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("%s: %s", MethodGetAccountHistory, resp)
}
