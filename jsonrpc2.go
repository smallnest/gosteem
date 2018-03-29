package gosteem

import (
	"encoding/json"
	"fmt"
)

// Request represents a JSON-RPC request or notification.
// See
// http://www.jsonrpc.org/specification#request_object and
// http://www.jsonrpc.org/specification#notification.
type Request struct {
	Version string          `json:"jsonrpc"`
	Method  string          `json:"method"`
	Params  json.RawMessage `json:"params,omitempty"`
	ID      uint64          `json:"id"`
	//Notification bool            `json:"-"`
}

type Response struct {
	Version string          `json:"jsonrpc"`
	ID      uint64          `json:"id"`
	Result  json.RawMessage `json:"result,omitempty"`
	Error   *Error          `json:"error,omitempty"`
}

// Error represents a JSON-RPC2 response error.
type Error struct {
	Code    int64            `json:"code"`
	Message string           `json:"message"`
	Data    *json.RawMessage `json:"data"`
}

func (e *Error) Error() string {
	return fmt.Sprintf("code: %d, message: %s", e.Code, e.Message)
}
