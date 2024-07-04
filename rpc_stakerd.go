package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
	"time"
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
			body, readErr := ioutil.ReadAll(resp.Body)
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

	body, readErr := ioutil.ReadAll(resp.Body)
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

	body, err := ioutil.ReadAll(resp.Body)
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
	rand.Seed(time.Now().UnixNano())
	randomIndex := rand.Intn(len(providers))
	return providers[randomIndex].BtcPk, nil
}