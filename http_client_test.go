package gosteem

import (
	json "encoding/json"
	"testing"
	"time"
)

var (
	testAPIAddress = "https://api.steemit.com"
)

func TestHTTPClient_call(t *testing.T) {
	client := NewHTTPClient(testAPIAddress, 10*time.Second)

	var count int
	err := client.Call("get_account_count", nil, &count)
	if err != nil {
		t.Fatal(err)
	}

	if count == 0 {
		t.Fatalf("expected account_count > 0 but got 0")
	}
	t.Logf("get_account_count: %d", count)

	var result json.RawMessage
	err = client.Call("get_dynamic_global_properties", nil, &result)
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("get_dynamic_global_properties: %s", result)
}
