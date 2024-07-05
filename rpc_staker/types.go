package rpc_staker

// ErrorResponse represents the structure of the error response from the server
type FinalityProviderErrorResponse struct {
	JSONRPC string `json:"jsonrpc"`
	ID      int    `json:"id"`
	Error   struct {
		Code    int    `json:"code"`
		Message string `json:"message"`
		Data    string `json:"data"`
	} `json:"error"`
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