package api

import (
	"time"

	"github.com/smallnest/gosteem"
)

var (
	testAPIAddress = "https://api.steemit.com"
	client         = gosteem.NewHTTPClient(testAPIAddress, 10*time.Second)
)
