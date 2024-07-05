package main

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/babylonchain/babylon/btcstaking"
	"github.com/btcsuite/btcd/btcutil"
	"github.com/btcsuite/btcd/rpcclient"
	"github.com/decred/dcrd/dcrec/secp256k1"
)

// creates new client instance.
func bitcoindCreateClient() (*rpcclient.Client, error){
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
func bitcoindCreateRawRequest(client *rpcclient.Client, method string, params []json.RawMessage, result interface{}) error {
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


// https://github.com/babylonchain/babylon/blob/dev/btcstaking/staking_test.go

/*
func bitcoindDoRawStakingTransaction(client *rpcclient.Client, staker_key *secp256k1.PublicKey, fpKey *secp256k1.PublicKey, stakingAmount btcutil.Amount, stakingTime uint16, allowHighFees bool ) (string, error) {
	// 1) Building the unfunded and not signed staking transaction
	// BuildV0IdentifiableStakingOutputsAndTx's Wrapper
	stakingInfo, err := btcstaking.BuildStakingInfo() 


	// 2) funding the raw transaction 
	funded_staking_transaction, err := client.FundRawTransaction(stakingInfoTxnHash)
	if err != nil {
		log.Println("error: ", err)
	}

	// 3) sign the funded staking transaction
	funded_staking_transaction_hex := funded_staking_transaction.Transaction
	signed_funded_staking_transaction, ok, err := client.SignRawTransactionWithWallet(funded_staking_transaction_hex)
	if err != nil {
		log.Println("error: ", err)
	}


	// 4) send the raw signed and funded transaction
	sent_raw_transaction, err := client.SendRawTransaction(signed_funded_staking_transaction, allowHighFees)
	if err != nil {
		log.Println("error: ", err)
	}

	// print the Staking Transaction Hex
	log.Println("Staking Transaction: ", sent_raw_transaction)
}
*/