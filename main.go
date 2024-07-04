package main

import (
	"encoding/json"
	"fmt"
	"log"
)

// Due to unupdated btcd rpc client,
// Assumption: we will have to assume that the bitcoind is setup with only 1 wallet, whose name is preknown.


func main() {
	client, err := bitcoindCreateClient()
	if err != nil {
		log.Fatalf("error creating new btc client: %v", err)
	}
	defer client.Shutdown()

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

		// // iterate over accounts (map[string]btcutil.Amount) and write to stdout
	// for label, amount := range accounts {
	// 	log.Printf("%s: %s", label, amount)
	// }

	// // prepare a sendMany transaction
	// receiver1, err := btcutil.DecodeAddress("1someAddressThatIsActuallyReal", &chaincfg.MainNetParams)
	// if err != nil {
	// 	log.Fatalf("address receiver1 seems to be invalid: %v", err)
	// }
	// receiver2, err := btcutil.DecodeAddress("1anotherAddressThatsPrettyReal", &chaincfg.MainNetParams)
	// if err != nil {
	// 	log.Fatalf("address receiver2 seems to be invalid: %v", err)
	// }
	// receivers := map[btcutil.Address]btcutil.Amount{
	// 	receiver1: 42,  // 42 satoshi
	// 	receiver2: 100, // 100 satoshi
	// }

	// // create and send the sendMany tx
	// txSha, err := client.SendMany("some-account-label-from-which-to-send", receivers)
	// if err != nil {
	// 	log.Fatalf("error sendMany: %v", err)
	// }
	// log.Printf("sendMany completed! tx sha is: %s", txSha.String())

	//  Staking Transaction Code helps
	// accounts, err := client.FundRawTransaction()
	// accounts, err := client.SignRawTransactionWithWallet()
	// accounts, err := client.SendRawTransaction()


}
