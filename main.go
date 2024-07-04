package main

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/btcsuite/btcd/rpcclient"
)

// Due to unupdated btcd rpc client,
// Assumption: we will have to assume that the bitcoind is setup with only 1 wallet, whose name is preknown.


func main() {
	// 1) create new client instance.
	client, err := rpcclient.New(&rpcclient.ConnConfig{
		HTTPPostMode: true,
		DisableTLS:   true,
		Host:         "127.0.0.1:38332",
		User:         "dexterhv",
		Pass:         "verma",
	}, nil)
	if err != nil {
		log.Fatalf("error creating new btc client: %v", err)
	}

	// Prepare the method and parameters
	method := "listwallets"
	params := []json.RawMessage{}

	// Call RawRequest
	rawResp, err := client.RawRequest(method, params)
	if err != nil {
		log.Fatalf("RawRequest error: %v", err)
	}
	log.Println(rawResp, " are the currently rawResp ")

	// Print the raw response
	log.Printf("Raw response: %s\n", rawResp)

	// Unmarshal the response into a map
	var result []string
	err = json.Unmarshal(rawResp, &result)
	if err != nil {
		log.Fatalf("Error unmarshaling response: %v", err)
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
