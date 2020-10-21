package coinmarketcap

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

type CoinMarketCapSdk struct {
	client    *http.Client
	url       string
}

func DefaultCoinMarketCapSdk() *CoinMarketCapSdk {
	//return NewCoinMarketCapSdk("https://api.coinmarketcap.com/v2")
	return NewCoinMarketCapSdk("https://pro-api.coinmarketcap.com/v1/cryptocurrency/")
}

func NewCoinMarketCapSdk(url string) *CoinMarketCapSdk {
	client := &http.Client{}
	sdk := &CoinMarketCapSdk{
		client: client,
		url: url,
	}
	return sdk
}

type ListingsMedia struct {
	Data []*Listing `json:"data"`
}

func (sdk  *CoinMarketCapSdk) ListingsLatest() ([]*Listing, error) {
	req, err := http.NewRequest("GET", sdk.url + "listings/latest", nil)
	if err != nil {
		return nil, err
	}

	q := url.Values{}
	q.Add("start", "1")
	q.Add("limit", "5000")
	q.Add("convert", "USD")

	req.Header.Set("Accepts", "application/json")
	req.Header.Add("X-CMC_PRO_API_KEY", "8efe5156-8b37-4c77-8e1d-a140c97bf466")
	req.URL.RawQuery = q.Encode()

	resp, err := sdk.client.Do(req);
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("response status code: %d", resp.StatusCode)
	}
	respBody, _ := ioutil.ReadAll(resp.Body)
	var body ListingsMedia
	err = json.Unmarshal(respBody, &body)
	if err != nil {
		return nil, err
	}
	return body.Data, nil
}

type QuotesLatestMedia struct {
	Data  map[string]*Ticker `json:"data"`
}

func (sdk  *CoinMarketCapSdk) QuotesLatest() (map[string]*Ticker, error) {
	req, err := http.NewRequest("GET", sdk.url + "quotes/latest", nil)
	if err != nil {
		return nil, err
	}

	q := url.Values{}
	q.Add("id", "1,2620,2316,7083")
	q.Add("convert", "USD")

	req.Header.Set("Accepts", "application/json")
	req.Header.Add("X-CMC_PRO_API_KEY", "8efe5156-8b37-4c77-8e1d-a140c97bf466")
	req.URL.RawQuery = q.Encode()

	resp, err := sdk.client.Do(req);
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("response status code: %d", resp.StatusCode)
	}
	respBody, _ := ioutil.ReadAll(resp.Body)
	var body QuotesLatestMedia
	err = json.Unmarshal(respBody, &body)
	if err != nil {
		return nil, err
	}
	return body.Data, nil
}
