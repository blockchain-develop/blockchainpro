package binance

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type BinanceSdk struct {
	client    *http.Client
	url       string
}

func DefaultBinanceSdk() *BinanceSdk {
	//return NewCoinMarketCapSdk("https://api.coinmarketcap.com/v2")
	return NewBinanceSdk("https://api1.binance.com/")
}

func NewBinanceSdk(url string) *BinanceSdk {
	client := &http.Client{}
	sdk := &BinanceSdk{
		client: client,
		url: url,
	}
	return sdk
}

func (sdk  *BinanceSdk) QuotesLatest() ([]*Ticker, error) {
	req, err := http.NewRequest("GET", sdk.url + "api/v3/ticker/price", nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Accepts", "application/json")
	req.Header.Add("X-CMC_PRO_API_KEY", "8efe5156-8b37-4c77-8e1d-a140c97bf466")

	resp, err := sdk.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("response status code: %d", resp.StatusCode)
	}
	respBody, _ := ioutil.ReadAll(resp.Body)
	fmt.Printf("resp body: %s\n", string(respBody))
	tickers := make([]*Ticker, 0)
	err = json.Unmarshal(respBody, &tickers)
	if err != nil {
		return nil, err
	}
	return tickers, nil
}
