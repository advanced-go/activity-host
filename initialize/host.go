package initialize

import (
	"github.com/advanced-go/stdlib/core"
	"github.com/advanced-go/stdlib/host"
	"net/http"
	"time"
)

func Host(cmdLine []string) error {
	// Initialize host proxy for all HTTP handlers,and add intermediaries
	host.SetHostTimeout(time.Second * 3)
	host.SetAuthExchange(AuthHandler, nil)
	//registerExchanges()
	//err := host.RegisterExchange(module.Authority, host.NewAccessLogIntermediary(http2.Exchange))
	return nil
}

/*
func registerExchanges() error {
	err := host.RegisterExchange(module.Authority, host.NewAccessLogIntermediary(http2.Exchange))
	if err != nil {
		return err
	}
	err = host.RegisterExchange(module.Authority, host.NewAccessLogIntermediary(-search", http2.Exchange))
	if err != nil {
		return err
	}
	return nil
}

*/

func AuthHandler(r *http.Request) (*http.Response, *core.Status) {
	/*
		if r != nil {
			tokenString := r.Header.Get(host.Authorization)
			if tokenString == "" {
				status := core.NewStatus(http.StatusUnauthorized)
				return &http.Response{StatusCode: status.HttpCode()}, status
				//w.WriteHeader(http.StatusUnauthorized)
				//fmt.Fprint(w, "Missing authorization header")
			}
		}


	*/
	return &http.Response{StatusCode: http.StatusOK}, core.StatusOK()

}
