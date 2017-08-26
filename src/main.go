package main

import (
	"fmt"
	"github.com/ybbus/jsonrpc"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

func keepLines(s string, n int) string {
	result := strings.Join(strings.Split(s, "\n")[:n], "\n")
	return strings.Replace(result, "\r", "", -1)
}

func getConfig() {
	return
}

func main() {
	queryMetric()
	postData()
}

func postData() {
	resp, err := http.PostForm("http://localhost:3000/miner",
		url.Values{"q": {"github"}})
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	fmt.Println("post:\n", keepLines(string(body), 3))
}

func queryMetric() {

	type ClaymoreMiner struct {
		Result string `json:"result"`
	}

	// create client
	rpcClient := jsonrpc.NewRPCClient("http://localhost:4343")

	// execute rpc to service
	response, _ := rpcClient.Call("getMinetStat1")

	// parse result into struct
	var claymore ClaymoreMiner
	response.GetObject(&claymore)

}
