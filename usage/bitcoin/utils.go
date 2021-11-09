package bitcoin

import "github.com/btcsuite/btcd/rpcclient"

func NewBitcoinClient(url string, user string, passwd string) (*rpcclient.Client, error) {
	connCfg := &rpcclient.ConnConfig{
		Host:         url,
		DisableTLS:   true,
		User:         user,
		Pass:         passwd,
		HTTPPostMode: true,
	}
	client, err := rpcclient.New(connCfg, nil)
	if err != nil {
		return nil, err
	}
	return client, nil
}

func DefaultBitcoinClient() (*rpcclient.Client, error) {
	//url := "seed.tbtc.petertodd.org"
	url := "127.0.0.1:51001"
	return NewBitcoinClient(url, "test", "test")
}