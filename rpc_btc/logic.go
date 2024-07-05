package rpc_btc

import (
	"encoding/json"
	"fmt"

	"github.com/btcsuite/btcd/rpcclient"
)

// creates new client instance.
func CreateClient() (*rpcclient.Client, error){
	client, err := rpcclient.New(&rpcclient.ConnConfig{
		HTTPPostMode: true,
		DisableTLS:   true,
		Host:         "127.0.0.1:38332",
		User:         "dexterhv",
		Pass:         "verma",
	}, nil)
	
	return client,err
}


// Generalized function to make a raw JSON-RPC request to bitcoind server
func CreateRawRequest(client *rpcclient.Client, method string, params []json.RawMessage, result interface{}) error {
	// Call RawRequest
	rawResp, err := client.RawRequest(method, params)
	if err != nil {
		return fmt.Errorf("RawRequest error: %v", err)
	}

	// Print the raw response
	fmt.Printf("Raw response: %s\n", rawResp)

	// Unmarshal the response into the provided result type
	err = json.Unmarshal(rawResp, result)
	if err != nil {
		return fmt.Errorf("error unmarshaling response: %v", err)
	}

	return nil
}
