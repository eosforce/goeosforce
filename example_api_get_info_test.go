package eos_test

import (
	"context"
	"encoding/json"
	"fmt"

	eos "github.com/eosforce/goeosforce"
)

func ExampleAPI_GetInfo() {
	api := eos.New(getAPIURL())

	info, err := api.GetInfo(context.Background())
	if err != nil {
		panic(fmt.Errorf("get info: %s", err))
	}

	bytes, err := json.Marshal(info)
	if err != nil {
		panic(fmt.Errorf("json marshal response: %s", err))
	}

	fmt.Println(string(bytes))
}
