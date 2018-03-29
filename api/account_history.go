package api

import (
	"encoding/json"

	"github.com/smallnest/gosteem"
)

var (
// (get_ops_in_block)
// (get_transaction)
// MethodGetAccountHistory = "get_account_history"
)

func GetAccountHistory(client gosteem.Client, account string, from, limit int) (json.RawMessage, error) {
	var resp json.RawMessage
	err := client.Call(MethodGetAccountHistory, []interface{}{account, from, limit}, &resp)
	return resp, err
}
