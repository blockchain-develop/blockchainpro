package ipfs

import (
    "io/ioutil"
    "testing"
)

import (
    "bytes"
    "context"
    "fmt"
    api "github.com/ipfs/go-ipfs-api"
    "net/http"
)

type authTransport struct {
    http.RoundTripper
    Token string
}

func (t authTransport) RoundTrip(r *http.Request) (*http.Response, error) {
    r.Header.Add("X-ipfs-Token", t.Token)
    return t.RoundTripper.RoundTrip(r)
}

func TestPin(t *testing.T) {
    url := "https://ipfs.phemex.com/"
    token := "AUnivT2vzsQHp9wjMtunQddGFUWo7CAzA6d6ygRYrfUpCzGiicNwvFKu7Kc9KQK4"
    c := &http.Client{
        Transport: &authTransport{
            RoundTripper: http.DefaultTransport,
            Token:        token,
        },
    }
    /*
    // add remote service
    {
        req, err := http.NewRequest("POST", "https://ipfs.phemex.com/api/v0/pin/remote/service/add", nil)
        if err != nil {
            panic(err)
        }

        q := req.URL.Query()
        q.Add("arg", "pindata1")
        q.Add("arg", "https://api.pinata.cloud/psa")
        q.Add("arg", "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VySW5mb3JtYXRpb24iOnsiaWQiOiIzNWRlMzM3MC0wODEzLTQ3MWYtOTNiMS01N2MzOTJhZGY4ZTEiLCJlbWFpbCI6InZnYW90YW5AZ21haWwuY29tIiwiZW1haWxfdmVyaWZpZWQiOnRydWUsInBpbl9wb2xpY3kiOnsicmVnaW9ucyI6W3siaWQiOiJGUkExIiwiZGVzaXJlZFJlcGxpY2F0aW9uQ291bnQiOjF9LHsiaWQiOiJOWUMxIiwiZGVzaXJlZFJlcGxpY2F0aW9uQ291bnQiOjF9XSwidmVyc2lvbiI6MX0sIm1mYV9lbmFibGVkIjpmYWxzZSwic3RhdHVzIjoiQUNUSVZFIn0sImF1dGhlbnRpY2F0aW9uVHlwZSI6InNjb3BlZEtleSIsInNjb3BlZEtleUtleSI6IjcwNTE4NTVjYmQ1MGMxM2JlMGY4Iiwic2NvcGVkS2V5U2VjcmV0IjoiOGYzOGI3MGVlNThlYWVhNTRhNWY0YzJiZDIxYWUxNDliOWQ3OGExNWMzM2UyYTA2ZDAwNWFjODUxN2Y2NTkxNSIsImlhdCI6MTY3Njg5NjMyMn0.wiU8Vh7nDxNfBwrTV1mna3GdYo9dTK53as_PQIYz6mE")
        req.URL.RawQuery = q.Encode()

        rsp, err := c.Do(req)
        if err != nil {
            panic(err)
        }

        defer rsp.Body.Close()
        body, err := ioutil.ReadAll(rsp.Body)
        if err != nil {
            panic(err)
        }

        fmt.Printf("%s\n", body)
    }

    // get remote service
    {
        req, err := http.NewRequest("POST", "https://ipfs.phemex.com/api/v0/pin/remote/service/ls", nil)
        if err != nil {
            panic(err)
        }

        q := req.URL.Query()
        q.Add("stat", "false")
        req.URL.RawQuery = q.Encode()

        rsp, err := c.Do(req)
        if err != nil {
            panic(err)
        }

        defer rsp.Body.Close()
        body, err := ioutil.ReadAll(rsp.Body)
        if err != nil {
            panic(err)
        }

        fmt.Printf("%s\n", body)
    }
*/
    ipfsClient := api.NewShellWithClient(url, c)

    repo := "/soulpass-metadata-dev/"
    tokenId := 100000
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
    fileHash := stat.Hash
    //
    stat, err = ipfsClient.FilesStat(context.Background(), repo)
    if err != nil {
        panic(err)
    }
    fmt.Printf("new sbt directory, cid: %s\n", stat.Hash)

    //
    {
        req, err := http.NewRequest("POST", "https://ipfs.phemex.com/api/v0/pin/remote/add", nil)
        if err != nil {
            panic(err)
        }

        q := req.URL.Query()
        q.Add("service", "pindata1")
        q.Add("name", "hello_ipfs.json")
        q.Add("arg", fileHash)
        req.URL.RawQuery = q.Encode()

        rsp, err := c.Do(req)
        if err != nil {
            panic(err)
        }

        defer rsp.Body.Close()
        body, err := ioutil.ReadAll(rsp.Body)
        if err != nil {
            panic(err)
        }

        fmt.Printf("%s\n", body)
    }
}

func TestLocalPin(t *testing.T) {
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
    tokenId := 110000
    nftMetaJson := "{\"key:\":\"hello, ipfs, 111\"}"
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
    ipfsPath := fmt.Sprintf("/ipfs/%s", stat.Hash)
    err = ipfsClient.Pin(ipfsPath)
    if err != nil {
        panic(err)
    }
}