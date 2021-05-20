package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/n454149301/http_proxy/server"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("args len not 2")
		return
	}
	configStr, err := ioutil.ReadFile(os.Args[1])
	if err != nil {
		panic(err.Error())
	}

	var tmpServer server.Server
	if err = json.Unmarshal(configStr, &tmpServer); err != nil {
		panic(err.Error())
	}
	(&tmpServer).Start()
}
