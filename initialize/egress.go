package initialize

import (
	"github.com/advanced-go/stdlib/httpx"
)

func EgressProxies(cmdLine []string) error {

	err := httpx.RegisterExchange("", nil)
	return err
}
