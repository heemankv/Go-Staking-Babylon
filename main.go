package main

import (
	"encoding/json"
	"fmt"
	"log"
)

// Due to unupdated btcd rpc client,
// Assumption: we will have to assume that the bitcoind is setup with only 1 wallet, whose name is preknown.

// AddressInfo represents the structure of address information
type AddressInfo struct {
	Purpose string `json:"purpose"`
	// Add other fields as needed
}


func main() {
	// 1) Create an RPC Client.
	client, err := bitcoindCreateClient()
	if err != nil {
		log.Fatalf("error creating new btc client: %v", err)
	}
	defer client.Shutdown()


	// 2) Check available wallets, it should return the wallets that were setup.

	// Prepare the method and parameters
	method := "listwallets"
	params := []json.RawMessage{}

	// Define a variable to hold the result
	var result []string

	// Call the generalized function
	err = bitcoindCreateRawRequest(client, method, params, &result)
	if err != nil {
		log.Fatalf("Error calling RPC method: %v", err)
	}

	// Print the result
	fmt.Printf("Result: %+v\n", result)

	// 3) Get the addresses of the given label, by default use the first address.

	// Prepare the method and parameters
	method = "getaddressesbylabel"
	params = []json.RawMessage{json.RawMessage(`"btcstaker"`)}

	// Define a variable to hold the result
	var result2 map[string]AddressInfo

	// Call the generalized function
	err = bitcoindCreateRawRequest(client, method, params, &result2)
	if err != nil {
		log.Fatalf("Error calling RPC method: %v", err)
	}

	// Print the result
	for address, info := range result2 {
		fmt.Printf("Address: %s, Purpose: %s\n", address, info.Purpose)
	}































































	//  Staking Transaction Code helps
	// accounts, err := client.FundRawTransaction()
	// accounts, err := client.SignRawTransactionWithWallet()
	// accounts, err := client.SendRawTransaction()


}
