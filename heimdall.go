package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/gojektech/heimdall/v6/httpclient"
	"github.com/gojektech/heimdall/v6/hystrix"
	"github.com/gojektech/heimdall/v6/plugins"
)

//func main() {
//	hystrixFunc()
//}

func simple() {
	timeout := 1000 * time.Millisecond
	client := httpclient.NewClient(httpclient.WithHTTPTimeout(timeout))

	requestLogger := plugins.NewRequestLogger(nil, nil)
	client.AddPlugin(requestLogger)

	req, _ := http.NewRequest(http.MethodGet, "http://www.baidu.com", nil)
	res, err := client.Do(req)
	if err != nil {
		panic(err)
	}

	_, err = ioutil.ReadAll(res.Body)
	// fmt.Println(string(body))
}

func hystrixFunc() {
	max := 5
	timeOut := 10000 * time.Millisecond

	client := hystrix.NewClient(
		hystrix.WithHTTPTimeout(timeOut),
		hystrix.WithHystrixTimeout(timeOut),

		hystrix.WithCommandName("google_get_request"),
		hystrix.WithMaxConcurrentRequests(30),
		hystrix.WithErrorPercentThreshold(max),
		//hystrix.WithStatsDCollector("localhost:8125", "myapp.hystrix"),
	)

	for i := 0; i < max*10; i++ {
		res, err := client.Get(fmt.Sprintf("http://localhost:10094/v1/apibusserv/order/query"), nil)
		if err != nil {
			fmt.Println(err)
			continue
		}

		body, err := ioutil.ReadAll(res.Body)
		fmt.Println(string(body))
	}

}
