package main

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"math/rand"
	"net/http"

	dc "github.com/babylonchain/btc-staker/stakerservice/client"
)

// ErrorResponse represents the structure of the error response from the server
type StakerDFinalityProviderErrorResponse struct {
	JSONRPC string `json:"jsonrpc"`
	ID      int    `json:"id"`
	Error   struct {
		Code    int    `json:"code"`
		Message string `json:"message"`
		Data    string `json:"data"`
	} `json:"error"`
}

// makeRequest makes a call to the specified URL and returns an error if the response status is 500
func stakerdGetFinalityProvidersList() (string, error) {
	url := "http://127.0.0.1:15812/babylon_finality_providers"
	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		if resp.StatusCode == http.StatusInternalServerError {
			body, readErr := io.ReadAll(resp.Body)
			if readErr != nil {
				return "", readErr
			}

			var errorResponse StakerDFinalityProviderErrorResponse
			jsonErr := json.Unmarshal(body, &errorResponse)
			if jsonErr != nil {
				return "", jsonErr
			}

			return "", fmt.Errorf(fmt.Sprintf("Error: %s, Code: %d, Data: %s", errorResponse.Error.Message, errorResponse.Error.Code, errorResponse.Error.Data))
		}
		return "", fmt.Errorf(fmt.Sprintf("HTTP Error: %s", resp.Status))
	}

	body, readErr := io.ReadAll(resp.Body)
	if readErr != nil {
		return "", readErr
	}

	return string(body), nil
}

// description details of the finality provider
type Description struct {
	Moniker         string `json:"moniker"`
	Identity        string `json:"identity"`
	Website         string `json:"website"`
	SecurityContact string `json:"security_contact"`
	Details         string `json:"details"`
}

//  details of a finality provider
type FinalityProvider struct {
	Description       Description `json:"description"`
	Commission        string      `json:"commission"`
	BtcPk             string      `json:"btc_pk"`
	ActiveTVL         int64       `json:"active_tvl"`
	TotalTVL          int64       `json:"total_tvl"`
	ActiveDelegations int64       `json:"active_delegations"`
	TotalDelegations  int64       `json:"total_delegations"`
}

// structure of the API response
type Response struct {
	Data       []FinalityProvider `json:"data"`
	Pagination struct {
		NextKey string `json:"next_key"`
	} `json:"pagination"`
}

func stakingApiGetFinalityProvidersList() (*Response, error) {
	url := "https://staking-api.testnet.babylonchain.io/v1/finality-providers"

	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("failed to make GET request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %v", err)
	}

	var response Response
	err = json.Unmarshal(body, &response)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal response: %v", err)
	}

	return &response, nil
}

// selects a random finality provider and returns its btc_pk
func getRandomFinalityProviderBtcPk(providers []FinalityProvider) (string, error) {
	if len(providers) == 0 {
		return "", fmt.Errorf("no finality providers available")
	}
	randomIndex := rand.Intn(len(providers))
	return providers[randomIndex].BtcPk, nil
}



type StakingRequest struct {
	StakerAddress     string `json:"stakerAddress"`
	StakingAmount     int    `json:"stakingAmount"`
	FpBtcPks          string `json:"fpBtcPks"`
	StakingTimeBlocks int    `json:"stakingTimeBlocks"`
}

type StakingResponse struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}

// stakerdDoStakeTransaction performs a POST request to stake and returns the response or an error
func stakerdDoStakeTransaction(stakerAddr string, stakingAmt int, fpBtcPk string, stakingTime int) (*StakingResponse, error) {
	url := "http://127.0.0.1:15812/stake"

	requestBody := StakingRequest{
		StakerAddress:     stakerAddr,
		StakingAmount:     stakingAmt,
		FpBtcPks:          fpBtcPk,
		StakingTimeBlocks: stakingTime,
	}

	jsonData, err := json.Marshal(requestBody)
	if err != nil {
		return nil, err
	}

	resp, err := http.Post(url, "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, readErr := io.ReadAll(resp.Body)
		if readErr != nil {
			return nil, readErr
		}
		return nil, fmt.Errorf(fmt.Sprintf("HTTP Error: %s, Body: %s", resp.Status, string(body)))
	}

	body, readErr := io.ReadAll(resp.Body)
	if readErr != nil {
		return nil, readErr
	}

	var response StakingResponse
	jsonErr := json.Unmarshal(body, &response)
	if jsonErr != nil {
		return nil, jsonErr
	}

	return &response, nil
}


func stakerDoStakeTransaction2(){
	stakerClient, err := dc.NewStakerServiceJsonRpcClient("tcp://" + "127.0.0.1:15812")

	res,err := stakerClient.Health(context.Background())
	if err != nil {
		log.Println("error : ", err)
	}
	log.Println("res : ", res)
}






// type StakerServiceJsonRpcClient struct {
// 	client *jsonrpcclient.Client
// }

// // TODO Add some kind of timeout config
// func NewStakerServiceJsonRpcClient(remoteAddress string) (*StakerServiceJsonRpcClient, error) {
// 	client, err := jsonrpcclient.New(remoteAddress)
// 	if err != nil {
// 		return nil, err
// 	}

// 	return &StakerServiceJsonRpcClient{
// 		client: client,
// 	}, nil
// }


// func (c *StakerServiceJsonRpcClient) Health(ctx context.Context) (*service.ResultHealth, error) {
// 	result := new(service.ResultHealth)
// 	_, err := c.client.Call(ctx, "health", map[string]interface{}{}, result)
// 	if err != nil {
// 		return nil, err
// 	}
// 	return result, nil
// }


// func (c *StakerServiceJsonRpcClient) BabylonFinalityProviders(ctx context.Context, offset *int, limit *int) (*service.FinalityProvidersResponse, error) {
// 	result := new(service.FinalityProvidersResponse)

// 	params := make(map[string]interface{})

// 	if limit != nil {
// 		params["limit"] = limit
// 	}

// 	if offset != nil {
// 		params["offset"] = offset
// 	}

// 	_, err := c.client.Call(ctx, "babylon_finality_providers", params, result)
// 	if err != nil {
// 		return nil, err
// 	}
// 	return result, nil
// }

// func (c *StakerServiceJsonRpcClient) Stake(
// 	ctx context.Context,
// 	stakerAddress string,
// 	stakingAmount int64,
// 	fpPks []string,
// 	stakingTimeBlocks int64,
// ) (*service.ResultStake, error) {
// 	result := new(service.ResultStake)

// 	params := make(map[string]interface{})
// 	params["stakerAddress"] = stakerAddress
// 	params["stakingAmount"] = stakingAmount
// 	params["fpBtcPks"] = fpPks
// 	params["stakingTimeBlocks"] = stakingTimeBlocks

// 	_, err := c.client.Call(ctx, "stake", params, result)
// 	if err != nil {
// 		return nil, err
// 	}
// 	return result, nil
// }