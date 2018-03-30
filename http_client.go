package gosteem

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"sync/atomic"
	"time"
)

// test
// curl --silent https://api.steemit.com --data '{"jsonrpc":"2.0","id":0,"method":"get_account_count","params":[]}' | python -m json.tool

const jsonContextType = "application/json"

var requestID uint64

// HTTPClient is the client to communicate with steemd via http.
type HTTPClient struct {
	c        *http.Client
	nodeAddr string
}

func NewHTTPClient(nodeAddr string, timeout time.Duration) *HTTPClient {
	client := &HTTPClient{}
	client.c = &http.Client{
		Timeout: timeout,
	}

	client.nodeAddr = nodeAddr

	return client
}

func (c *HTTPClient) Call(method string, params, result interface{}) error {
	var rawParams json.RawMessage
	var err error
	if params != nil {
		rawParams, err = json.Marshal(params)
	}

	if err != nil {
		return err
	}

	r := Request{
		Version: "2.0",
		Method:  method,
		Params:  rawParams,
		ID:      atomic.AddUint64(&requestID, 1),
	}

	data, err := r.MarshalJSON()
	if err != nil {
		return err
	}
	resp, err := c.c.Post(c.nodeAddr, jsonContextType, bytes.NewReader(data))
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	data, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	res := &Response{}
	err = json.Unmarshal(data, res)
	if err != nil {
		return err
	}

	if res.Error != nil {
		return res.Error
	}

	if string(res.Result) == "[]" || string(res.Result) == "[[]]" {
		return nil
	}

	return json.Unmarshal(res.Result, result)
}
