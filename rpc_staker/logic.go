package rpc_staker

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"math/rand"
	"net/http"
)

// makes a call to the specified URL and returns an error if the response status is 500
func GetFinalityProvidersList() (string, error) {
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

			var errorResponse FinalityProviderErrorResponse
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

func StakingApiGetFinalityProvidersList() (*Response, error) {
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
func GetRandomFinalityProviderBtcPk(providers []FinalityProvider) (string, error) {
	if len(providers) == 0 {
		return "", fmt.Errorf("no finality providers available")
	}
	randomIndex := rand.Intn(len(providers))
	return providers[randomIndex].BtcPk, nil
}

// performs a POST request to stake and returns the response or an error
func PerformStakeTransaction(stakerAddr string, stakingAmt int, fpBtcPk string, stakingTime int) (*StakingResponse, error) {
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
