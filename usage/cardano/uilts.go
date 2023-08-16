package cardano

import (
	"github.com/coinbase/rosetta-sdk-go/client"
	"net/http"
	"time"
)

func NewClient() *client.APIClient {
	cfg := client.NewConfiguration(
		"http://127.0.0.1:8080",
		//"https://ada.getblock.io/mainnet",
		"cardano",
		&http.Client{Timeout: 10*time.Second})
	//cfg.AddDefaultHeader("x-api-key", "2fed06c5-78d9-400f-a7ae-c391e608939a")
	client := client.NewAPIClient(cfg)
	return client
}
