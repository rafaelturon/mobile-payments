package daemon

import (
	"github.com/rafaelturon/decred-pi-wallet/internal/structuredLog"
)

var (
	message = "starting..."
)

func init() {
	message = "inited"
}

func Start() {
	structuredLog.Debug(message)
}
