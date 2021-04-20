package fjrpc

import (
	gj "github.com/filecoin-project/go-jsonrpc"
)

// export things from go-jsonrpc
type ClientCloser = gj.ClientCloser
type Config = gj.Config
type Option = gj.Option
type ServerOption = gj.ServerOption

var (
	WithReconnectBackoff = gj.WithNoReconnect
	WithPingInterval     = gj.WithPingInterval
	WithTimeout          = gj.WithTimeout
	WithNoReconnect      = gj.WithNoReconnect
	WithParamEncoder     = gj.WithParamEncoder
	WithMaxRequestSize   = gj.WithMaxRequestSize
)
