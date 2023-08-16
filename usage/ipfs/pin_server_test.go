package ipfs

import (
	"bytes"
	"context"
	"fmt"
	"github.com/ipfs/go-cid"
	api "github.com/ipfs/go-ipfs-api"
	client "github.com/ipfs/go-pinning-service-http-client"
	"net/http"
	"testing"
)

func TestListPinObjects(t *testing.T) {
    c := client.NewClient("https://api.pinata.cloud/psa", "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VySW5mb3JtYXRpb24iOnsiaWQiOiIzNWRlMzM3MC0wODEzLTQ3MWYtOTNiMS01N2MzOTJhZGY4ZTEiLCJlbWFpbCI6InZnYW90YW5AZ21haWwuY29tIiwiZW1haWxfdmVyaWZpZWQiOnRydWUsInBpbl9wb2xpY3kiOnsicmVnaW9ucyI6W3siaWQiOiJGUkExIiwiZGVzaXJlZFJlcGxpY2F0aW9uQ291bnQiOjF9LHsiaWQiOiJOWUMxIiwiZGVzaXJlZFJlcGxpY2F0aW9uQ291bnQiOjF9XSwidmVyc2lvbiI6MX0sIm1mYV9lbmFibGVkIjpmYWxzZSwic3RhdHVzIjoiQUNUSVZFIn0sImF1dGhlbnRpY2F0aW9uVHlwZSI6InNjb3BlZEtleSIsInNjb3BlZEtleUtleSI6IjcwNTE4NTVjYmQ1MGMxM2JlMGY4Iiwic2NvcGVkS2V5U2VjcmV0IjoiOGYzOGI3MGVlNThlYWVhNTRhNWY0YzJiZDIxYWUxNDliOWQ3OGExNWMzM2UyYTA2ZDAwNWFjODUxN2Y2NTkxNSIsImlhdCI6MTY3Njg5NjMyMn0.wiU8Vh7nDxNfBwrTV1mna3GdYo9dTK53as_PQIYz6mE")
    rs, err := c.LsSync(context.Background())
    if err != nil {
    	panic(err)
	}
	for _, r := range rs {
		fmt.Printf("%s\n", r.String())
	}
}

func TestPins(t *testing.T) {
	url := "https://ipfs.phemex.com/"
	token := "AUnivT2vzsQHp9wjMtunQddGFUWo7CAzA6d6ygRYrfUpCzGiicNwvFKu7Kc9KQK4"
	c := &http.Client{
		Transport: &authTransport{
			RoundTripper: http.DefaultTransport,
			Token:        token,
		},
	}

	ipfsClient := api.NewShellWithClient(url, c)

	repo := "/soulpass-metadata-dev/"
	tokenId := 100003
	nftMetaJson := "{\"key:\":\"hello, ipfs\"}"
	path := fmt.Sprintf("%s%d.json", repo, tokenId)
	err := ipfsClient.FilesWrite(context.Background(), path, bytes.NewBufferString(string(nftMetaJson)),
		api.FilesWrite.Create(true),
		api.FilesWrite.CidVersion(1))
	if err != nil {
		panic(err)
	}
	//
	stat, err := ipfsClient.FilesStat(context.Background(), path)
	if err != nil {
		panic(err)
	}
	fmt.Printf("add nft meta into nft, cid: %s\n", stat.Hash)
	//fileHash := stat.Hash
	//
	stat, err = ipfsClient.FilesStat(context.Background(), repo)
	if err != nil {
		panic(err)
	}
	fmt.Printf("new sbt directory, cid: %s\n", stat.Hash)

	//
	pinClient := client.NewClient("https://api.pinata.cloud/psa", "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VySW5mb3JtYXRpb24iOnsiaWQiOiIzNWRlMzM3MC0wODEzLTQ3MWYtOTNiMS01N2MzOTJhZGY4ZTEiLCJlbWFpbCI6InZnYW90YW5AZ21haWwuY29tIiwiZW1haWxfdmVyaWZpZWQiOnRydWUsInBpbl9wb2xpY3kiOnsicmVnaW9ucyI6W3siaWQiOiJGUkExIiwiZGVzaXJlZFJlcGxpY2F0aW9uQ291bnQiOjF9LHsiaWQiOiJOWUMxIiwiZGVzaXJlZFJlcGxpY2F0aW9uQ291bnQiOjF9XSwidmVyc2lvbiI6MX0sIm1mYV9lbmFibGVkIjpmYWxzZSwic3RhdHVzIjoiQUNUSVZFIn0sImF1dGhlbnRpY2F0aW9uVHlwZSI6InNjb3BlZEtleSIsInNjb3BlZEtleUtleSI6IjcwNTE4NTVjYmQ1MGMxM2JlMGY4Iiwic2NvcGVkS2V5U2VjcmV0IjoiOGYzOGI3MGVlNThlYWVhNTRhNWY0YzJiZDIxYWUxNDliOWQ3OGExNWMzM2UyYTA2ZDAwNWFjODUxN2Y2NTkxNSIsImlhdCI6MTY3Njg5NjMyMn0.wiU8Vh7nDxNfBwrTV1mna3GdYo9dTK53as_PQIYz6mE")
	cid, err := cid.Decode(stat.Hash)
	if err != nil {
		panic(err)
	}
	rs, err := pinClient.Add(context.Background(), cid, client.PinOpts.WithName("/phemex_dev/"))
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s\n", rs.String())
}
