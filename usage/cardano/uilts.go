package cardano

import (
	"github.com/coinbase/rosetta-sdk-go/client"
	"net/http"
	"time"
)

func NewClient() *client.APIClient {
	cfg := client.NewConfiguration(
		"http://localhost:8081",
		"cardano",
		&http.Client{Timeout: 10*time.Second})
	client := client.NewAPIClient(cfg)
	return client
}
