package rpc_btc

import (
	"encoding/json"
	"fmt"

	"github.com/btcsuite/btcd/rpcclient"
)

// creates new client instance.
func CreateClient(host string, user string, pass string) (*rpcclient.Client, error){
	client, err := rpcclient.New(&rpcclient.ConnConfig{
		HTTPPostMode: true,
		DisableTLS:   true,
		Host:         host,
		User:         user,
		Pass:         pass,
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

	// Unmarshal the response into the provided result type
	err = json.Unmarshal(rawResp, result)
	if err != nil {
		return fmt.Errorf("error unmarshaling response: %v", err)
	}

	return nil
}
