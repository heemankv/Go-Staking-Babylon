package main

import (
	"encoding/json"
	"fmt"
	btcService "heemank_demo/rpc_btc"
	stakerService "heemank_demo/rpc_staker"
	"log"
	"strconv"
	"strings"
	"time"
)

// Due to unupdated btcd rpc client,
// Assumption: we will have to assume that the bitcoind is setup with only 1 wallet, whose name is preknown.



func parseBalance(balanceStr string) (float64, error) {
	// Trim any leading or trailing whitespace
	balanceStr = strings.TrimSpace(balanceStr)

	// Remove " BTC" suffix
	balanceStr = strings.TrimSuffix(balanceStr, " BTC")

	// Parse balance as float64
	balanceFloat, err := strconv.ParseFloat(balanceStr, 64)
	if err != nil {
		return 0.0, fmt.Errorf("error parsing balance: %v", err)
	}

	return balanceFloat, nil
}


func main() {

	// PART-1

	// 1) Create an RPC Client.
	client, err := btcService.CreateClient()
	if err != nil {
		log.Fatalf("error creating new btc client: %v", err)
	}
	defer client.Shutdown()

	// 2) Check available wallets, it should return the wallets that were setup.

	method := "listwallets"
	params := []json.RawMessage{}
	var result []string

	err = btcService.CreateRawRequest(client, method, params, &result)
	if err != nil {
		log.Fatalf("Error calling RPC method: %v", err)
	} else if !(len(result) > 0) {
		log.Fatalf("Error No Wallet Found: %v", result)
		return
	}

	log.Printf("Listwallets: %+v\n", result)

	// 3) Get the 0addresses of the given label, by default use the first address.
	
	walletToTrack := result[0]
	method = "getaddressesbylabel"
	params = []json.RawMessage{json.RawMessage(fmt.Sprintf(`"%s"`, walletToTrack))}
	var result2 map[string]btcService.AddressInfo

	err = btcService.CreateRawRequest(client, method, params, &result2)
	if err != nil {
		log.Fatalf("Error calling RPC method getaddressesbylabel : %v", err)
	}

	var addressToTrack string
	for address := range result2 {
		addressToTrack = address
		break
	}

	log.Println("Tracking address:", addressToTrack, "and wallet :", walletToTrack)

	// 4) in Loop : break if balance > 0.0005
	// Loop to check balance until it's greater than 0.0005 BTC

	// development 
	count := 0

	for {
		balanceResult, err := client.GetBalance("*")
		if err != nil {
			log.Fatalf("Error getting balance: %v", err)
		}

		balanceFloat, err := parseBalance(balanceResult.String())
		if err != nil {
			log.Fatalf("Error parsing balance: %v", err)
		}

		log.Printf("Balance of %s: %f BTC\n", addressToTrack, balanceFloat)

		if balanceFloat > 0.0005 || count > 5 {
			break
		}

		time.Sleep(5 * time.Second)
		count++
	}

	log.Printf("Balance of %s is now greater than 0.0005 BTC\n", addressToTrack)


	// PART-2

	// 1) Call the Stakerd Finality Provider function
	finalityProviders, err := stakerService.GetFinalityProvidersList()
	if err != nil {
		log.Printf("GetFinalityProvidersList failed: %v\n", err)
	} else{
		log.Println("GetFinalityProvidersList: ", finalityProviders)
	}
	
	apifinalityProviders, err2 := stakerService.StakingApiGetFinalityProvidersList()
	if err2 != nil {
		log.Printf("StakingApiGetFinalityProvidersList failed: %v\n", err2)
		return
	}

	// 2) Selecting Finality provider at random
	btcPk, err := stakerService.GetRandomFinalityProviderBtcPk(apifinalityProviders.Data)
	if err != nil {
		log.Fatalf("GetRandomFinalityProviderBtcPk failed: %v", err)
	}

	log.Printf("Random Finality Provider BTC PK: %s\n", btcPk)

	// 3) Performing Staking Transaction
	stakingAmt := 1000000
	stakingTime := 1000

	response, err := stakerService.PerformStakeTransaction(addressToTrack, stakingAmt, btcPk, stakingTime)
	if err != nil {
		log.Printf("PerformStakeTransaction failed: %v\n", err)
		return
	}

	log.Printf("PerformStakeTransaction Response: %+v\n", response)

}