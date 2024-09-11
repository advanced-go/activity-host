package initialize

import (
	"github.com/advanced-go/stdlib/httpx"
)

func EgressProxies() error {

	err := httpx.RegisterExchange("", nil)
	return err
}
