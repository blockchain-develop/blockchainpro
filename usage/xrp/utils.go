package xrp

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/rubblelabs/ripple/data"
	"io/ioutil"
	"net/http"
)

const MainNetUrl = "https://s2.ripple.com:51234"

type accountInfoResp struct {
	Result struct {
		AccountData struct {
			Account           string `json:"Account"`
			Balance           string `json:"Balance"`
			Flags             int    `json:"Flags"`
			LedgerEntryType   string `json:"LedgerEntryType"`
			OwnerCount        int    `json:"OwnerCount"`
			PreviousTxnID     string `json:"PreviousTxnID"`
			PreviousTxnLgrSeq int    `json:"PreviousTxnLgrSeq"`
			Sequence          int    `json:"Sequence"`
			Index             string `json:"index"`
		} `json:"account_data"`
	} `json:"result"`
}

func getAccount(url string, addr string) (*accountInfoResp, error) {
	data := make(map[string]interface{})
	data["method"] = "account_info"
	data["params"] = []map[string]interface{} {
		{
			"account": addr,
			"strict": true,
			"ledger_index": "current",
			"queue": true,
		},
	}
	bytesData, _ := json.Marshal(data)

	req, err := http.NewRequest("POST", url, bytes.NewReader(bytesData))
	if err != nil {
		return nil, err
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("response status code: %d", resp.StatusCode)
	}
	respBody, _ := ioutil.ReadAll(resp.Body)
	fmt.Printf("resp body: %s\n", string(respBody))
	var p accountInfoResp
	err = json.Unmarshal(respBody, &p)
	if err != nil {
		return nil, err
	}
	return &p, nil
}

type FeeResp struct {
	Result struct {
		Drops struct {
			BaseFee       string `json:"base_fee"`
			MedianFee     string `json:"median_fee"`
			MinimumFee    string `json:"minimum_fee"`
			OpenLedgerFee string `json:"open_ledger_fee"`
		} `json:"drops"`
	} `json:"result"`
}

func getFee(url string) (*FeeResp, error) {
	data := make(map[string]interface{})
	data["method"] = "fee"
	data["params"] = []map[string]interface{} {
	}
	bytesData, _ := json.Marshal(data)

	req, err := http.NewRequest("POST", url, bytes.NewReader(bytesData))
	if err != nil {
		return nil, err
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("response status code: %d", resp.StatusCode)
	}
	respBody, _ := ioutil.ReadAll(resp.Body)
	fmt.Printf("resp body: %s\n", string(respBody))
	var p FeeResp
	err = json.Unmarshal(respBody, &p)
	if err != nil {
		return nil, err
	}
	return &p, nil
}

type submitParam struct {
	TxBlob string `json:"tx_blob"`
}

type submitResp struct {
	Result struct {
		Error          string `json:"error"`
		ErrorException string `json:"error_exception"`
		ErrorMessage   string `json:"error_message"`

		Accepted            bool                   `json:"accepted"`
		EngineResult        data.TransactionResult `json:"engine_result"`
		EngineResultCode    int                    `json:"engine_result_code"`
		EngineResultMessage string                 `json:"engine_result_message"`
		TxBlob              string                 `json:"tx_blob"`
		Tx                  struct {
			Hash string `json:"hash"`
		} `json:"tx_json"`
	} `json:"result"`
}

func submitTransaction(url string, subReq *submitParam) (*submitResp, error) {
	data := make(map[string]interface{})
	data["method"] = "submit"
	data["params"] = []submitParam{*subReq}

	bytesData, _ := json.Marshal(data)

	req, err := http.NewRequest("POST", url, bytes.NewReader(bytesData))
	if err != nil {
		return nil, err
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("response status code: %d", resp.StatusCode)
	}
	respBody, _ := ioutil.ReadAll(resp.Body)
	fmt.Printf("resp body: %s\n", string(respBody))
	var p submitResp
	err = json.Unmarshal(respBody, &p)
	if err != nil {
		return nil, err
	}
	return &p, nil
}
