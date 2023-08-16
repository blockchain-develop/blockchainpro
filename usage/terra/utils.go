package terra

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/shopspring/decimal"
	"io/ioutil"
	"net/http"
	"net/url"
)

var (
	// mainnet
	//MainNet = "https://terra-rpc.easy2stake.com"
	MainNet = "https://lcd.terra.dev"
	// testnet
	TestNet = "https://bombay.stakesystems.io:2053"
)

type NodeInfo struct {
}

func GetNodeInfo(endpoint string) (*NodeInfo, error) {
	req, err := http.NewRequest("GET", endpoint+"/node_info", nil)
	if err != nil {
		return nil, err
	}
	client := &http.Client{}
	rsp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer rsp.Body.Close()
	if rsp.StatusCode != 200 {
		return nil, fmt.Errorf("response status code: %d", rsp.StatusCode)
	}
	rspData, _ := ioutil.ReadAll(rsp.Body)
	var nodeInfo NodeInfo
	err = json.Unmarshal(rspData, &nodeInfo)
	if err != nil {
		return nil, err
	}
	return &nodeInfo, nil
}

type ExchangeRate struct {
	Height uint64          `json:"height,string"`
	Result decimal.Decimal `json:"result, string"`
}

func GetExchangeRate(endpoint string, coin string) (*ExchangeRate, error) {
	req, err := http.NewRequest("GET", endpoint+fmt.Sprintf("/oracle/denoms/%s/exchange_rate", coin), nil)
	if err != nil {
		return nil, err
	}
	client := &http.Client{}
	rsp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer rsp.Body.Close()
	if rsp.StatusCode != 200 {
		return nil, fmt.Errorf("response status code: %d", rsp.StatusCode)
	}
	rspData, _ := ioutil.ReadAll(rsp.Body)
	var exchangeRate ExchangeRate
	err = json.Unmarshal(rspData, &exchangeRate)
	if err != nil {
		return nil, err
	}
	return &exchangeRate, nil
}

type Contract struct {
	ContractInfo *ContractInfo `json:"contract_info"`
}

type ContractInfo struct {
	Address string `json:"address"`
}

func GetContractInfo(endpoint string, contractAddr string) (*Contract, error) {
	req, err := http.NewRequest("GET", endpoint+fmt.Sprintf("/terra/wasm/v1beta1/contracts/%s", contractAddr), nil)
	if err != nil {
		return nil, err
	}
	client := &http.Client{}
	rsp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer rsp.Body.Close()
	if rsp.StatusCode != 200 {
		return nil, fmt.Errorf("response status code: %d", rsp.StatusCode)
	}
	rspData, _ := ioutil.ReadAll(rsp.Body)
	var contract Contract
	err = json.Unmarshal(rspData, &contract)
	if err != nil {
		return nil, err
	}
	return &contract, nil
}

type Pool struct {
}
type QueryMsg struct {
	Pool Pool `json:"pool"`
}

type AssetToken struct {
	Denom string `json:"denom"`
}

type AssetInfo struct {
	AssetToken *AssetToken `json:"native_token"`
}

type Asset struct {
	AssetInfo *AssetInfo      `json:"info"`
	Amount    decimal.Decimal `json:"amount, string"`
}

type QueryResult struct {
	Assets     []*Asset        `json:"assets"`
	TotalShare decimal.Decimal `json:"total_share, string"`
}

type QueryResponse struct {
	QueryResult *QueryResult `json:"query_result"`
}

func QueryContract(endpoint string, contractAddr string) (*QueryResult, error) {
	req, err := http.NewRequest("GET", endpoint+fmt.Sprintf("/terra/wasm/v1beta1/contracts/%s/store", contractAddr), nil)
	if err != nil {
		return nil, err
	}

	var queryMsg QueryMsg
	msg, _ := json.Marshal(queryMsg)

	//q := req.URL.Query()
	q := url.Values{}
	q.Add("query_msg", base64.StdEncoding.EncodeToString(msg))
	req.URL.RawQuery = q.Encode()

	fmt.Printf("raw url: %s\n", req.URL.String())

	client := &http.Client{}
	rsp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer rsp.Body.Close()
	if rsp.StatusCode != 200 {
		return nil, fmt.Errorf("response status code: %d", rsp.StatusCode)
	}
	rspData, _ := ioutil.ReadAll(rsp.Body)
	var queryResponse QueryResponse
	err = json.Unmarshal(rspData, &queryResponse)
	if err != nil {
		return nil, err
	}
	return queryResponse.QueryResult, nil
}
