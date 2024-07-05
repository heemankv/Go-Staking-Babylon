package main

import (
	"fmt"
	"strconv"
	"strings"
)

// Due to unupdated btcd rpc client,
// Assumption: we will have to assume that the bitcoind is setup with only 1 wallet, whose name is preknown.

// AddressInfo represents the structure of address information
type AddressInfo struct {
	Purpose string `json:"purpose"`
	// Add other fields as needed
}

func parseBalance(balanceStr string) (float64, error) {
	// Trim any leading or trailing whitespace
	balanceStr = strings.TrimSpace(balanceStr)

	// Remove " BTC" suffix
	if strings.HasSuffix(balanceStr, " BTC") {
		balanceStr = balanceStr[:len(balanceStr)-4]
	}

	// Parse balance as float64
	balanceFloat, err := strconv.ParseFloat(balanceStr, 64)
	if err != nil {
		return 0.0, fmt.Errorf("Error parsing balance: %v", err)
	}

	return balanceFloat, nil
}


// func main() {
// 	// 1) Create an RPC Client.
// 	client, err := bitcoindCreateClient()
// 	if err != nil {
// 		log.Fatalf("error creating new btc client: %v", err)
// 	}
// 	defer client.Shutdown()


// 	// 2) Check available wallets, it should return the wallets that were setup.

// 	// Prepare the method and parameters
// 	method := "listwallets"
// 	params := []json.RawMessage{}

// 	// Define a variable to hold the result
// 	var result []string

// 	// Call the generalized function
// 	err = bitcoindCreateRawRequest(client, method, params, &result)
// 	if err != nil {
// 		log.Fatalf("Error calling RPC method: %v", err)
// 	}

// 	// Print the result
// 	fmt.Printf("Result: %+v\n", result)

// 	walletToTrack := result[0]

// 	// 3) Get the addresses of the given label, by default use the first address.

// 	// Prepare the method and parameters
// 	method = "getaddressesbylabel"
// 	params = []json.RawMessage{json.RawMessage(`"btcstaker"`)}

// 	// Define a variable to hold the result
// 	var result2 map[string]AddressInfo

// 	// Call the generalized function
// 	err = bitcoindCreateRawRequest(client, method, params, &result2)
// 	if err != nil {
// 		log.Fatalf("Error calling RPC method: %v", err)
// 	}

// 	// Print the result
// 	for address, info := range result2 {
// 		fmt.Printf("Address: %s, Purpose: %s\n", address, info.Purpose)
// 	}

// 	// Get the first address to track
// 	var addressToTrack string
// 	for address := range result2 {
// 		addressToTrack = address
// 		break
// 	}

// 	log.Println("Tracking address: ", addressToTrack, "and wallet :", walletToTrack)

// 	// 4) in Loop : break if balance > 0.0005
// 	// Loop to check balance until it's greater than 0.0005 BTC
// 	for {
// 		balanceResult, err := client.GetBalance("*")
// 		if err != nil {
// 			log.Fatalf("Error getting balance: %v", err)
// 		}

// 		balanceFloat, err := parseBalance(balanceResult.String())
// 		if err != nil {
// 			log.Fatalf("Error parsing balance: %v", err)
// 		}

// 		fmt.Printf("Balance of %s: %f BTC\n", addressToTrack, balanceFloat)

// 		if balanceFloat > 0.0005 {
// 			break
// 		}

// 		time.Sleep(5 * time.Second)
// 	}

// 	fmt.Printf("Balance of %s is now greater than 0.0005 BTC\n", addressToTrack)


























































// 	//  Staking Transaction Code helps
// 	// accounts, err := client.FundRawTransaction()
// 	// accounts, err := client.SignRawTransactionWithWallet()
// 	// accounts, err := client.SendRawTransaction()


// }


func main(){

	// 1) Call the Stakerd Finality Provider function
	// response, err := stakerdGetFinalityProvidersList()
	// if err != nil {
	// 	fmt.Printf("Request failed: %v\n", err)
	// 	return
	// }

	// fmt.Printf("Response: %s\n", response)



	// 2) Call the Staking-Api Finality Provider function

	// response2, err2 := stakingApiGetFinalityProvidersList()
	// if err2 != nil {
	// 	fmt.Printf("Error: %v\n", err2)
	// 	return
	// }

	// // Print the response
	// fmt.Printf("Response: %+v\n", response2)

	// btcPk, err := getRandomFinalityProviderBtcPk(response2.Data)
	// if err != nil {
	// 	log.Fatalf("Failed to get random finality provider btc_pk: %v", err)
	// }

	// fmt.Printf("Random Finality Provider BTC PK: %s\n", btcPk)

	// stakerAddr := "tb1p64nx6k9d57l57f386mzl9fm9fkpeftet45ejzzq3cyvglqpl9wdqm5y788"
	// stakingAmt := 1000000
	// stakingTime := 1000

	// response, err := stakerdDoStakeTransaction(stakerAddr, stakingAmt, btcPk, stakingTime)
	// if err != nil {
	// 	fmt.Printf("Request failed: %v\n", err)
	// 	return
	// }

	// fmt.Printf("Response: %+v\n", response)

	stakerDoStakeTransaction2()


}