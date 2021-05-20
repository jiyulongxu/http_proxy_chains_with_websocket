package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/n454149301/http_proxy/client"
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

	var tmpClient client.Client
	if err = json.Unmarshal(configStr, &tmpClient); err != nil {
		panic(err.Error())
	}

	(&tmpClient).Start()
}
