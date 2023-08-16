package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

var (
	FaucetUrl = "https://api.bitaps.com/eth/testnet/v1/faucet/send/payment"
)

type Request struct {
	Address string `json:"address"`
	Amount uint64 `json:"amount"`
}

type Response struct {
	TxHash string `json:"tx_hash"`
	Result string `json:"result"`
}

func getETH() (string, error) {
	//
	reqOri := &Request{
		Address: "0x5F0d8D94513465118D70B953601AE1CcC2249766",
		Amount:  1000000000000000000,
	}
	reqData, _ := json.Marshal(reqOri)
	reqReader := bytes.NewReader(reqData)
	req, err := http.NewRequest("POST", FaucetUrl, reqReader)
	if err != nil {
		return "", err
	}
	//req.Header.Set("Accepts", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		return "", fmt.Errorf("response status code: %d", resp.StatusCode)
	}
	rspData, _ := ioutil.ReadAll(resp.Body)
	var rsp Response
	err = json.Unmarshal(rspData, &rsp)
	if err != nil {
		return "", err
	}
	if rsp.Result != "success" {
		return "", fmt.Errorf(rsp.Result)
	}
	return rsp.TxHash, nil
}

func main() {
	timer2 := time.NewTicker(time.Minute * 2)
	for {
		select {
		case <-timer2.C:
			hash, err := getETH()
			if err != nil {
				log.Printf("failed!, %s\n", err)
			} else {
				log.Printf("success, %s\n", hash)
			}
		}
	}
}
