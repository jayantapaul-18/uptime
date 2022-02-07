package load

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

// type restime time.Time

var metrics = make(map[string]string)

//var metrics_response = make(map[string]restime)

func Loadrun(name string, url string, method string) {
	// url := "http://192.168.1.175:1881"
	// method := "GET"

	// client := &http.Client{}
	client := http.Client{Timeout: 5 * time.Second}
	req, err := http.NewRequest(method, url, nil)

	if err != nil {
		fmt.Println(err)
		return
	}
	start := time.Now()
	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer res.Body.Close()
	// fmt.Println("Response Time:", time.Since(start))
	responsetime := time.Since(start)
	fmt.Println("Response:", responsetime)

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return
	}

	metrics["API_URL"] = url
	metrics["API_METHOD"] = method
	metrics["NAME"] = name
	// metrics_response["RESPONSE_TIME"] = responsetime

	fmt.Println("Load-Request-Start ==================")
	fmt.Println("Load-Request:", req)
	fmt.Println("Load-Response:", res)
	fmt.Println(string("Load-Response-Body:"))
	fmt.Println(string(body))
	fmt.Println("Load-Request-End   ==================")

	// Metrics
	fmt.Println(metrics)
}
