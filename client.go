package gosteem

type Client interface {
	Call(method string, params, result interface{}) error
}
