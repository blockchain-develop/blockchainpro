package web3_storage

import (
	"context"
	"fmt"
	"github.com/web3-storage/go-w3s-client"
	"os"
	"path"
	"testing"
)

func TestAAAA(t *testing.T) {
	filename := "/Users/tangaoyuan/Documents/gopath/src/cmexpro.com/ipfs-refresh/hello.txt"
	token := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiJkaWQ6ZXRocjoweDI4NWNlRjcxOGE0OENEZDU0OERlMEEwYTExQzAwM2ExRjc1NEExZDIiLCJpc3MiOiJ3ZWIzLXN0b3JhZ2UiLCJpYXQiOjE2NzU3Njg3MTE5MTIsIm5hbWUiOiJlZ2FvdGFuIn0.j-CNkFITNUhrsDbmRwHpehEDK7vLWfYF8i5Pdlzs_dA"

	// Create a new web3.storage client using the token
	client, err := w3s.NewClient(w3s.WithToken(token))
	if err != nil {
		panic(err)
	}

	// Open the file for reading
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}

	basename := path.Base(filename)
	// Upload to web3.storage
	fmt.Printf("Storing %s ...\n", basename)
	cid, err := client.Put(context.Background(), file)
	if err != nil {
		panic(err)
	}

	gatewayURL := fmt.Sprintf("https://%s.ipfs.dweb.link/%s\n", cid.String(), basename)
	fmt.Printf("Stored %s with web3.storage! View it at: %s\n", basename, gatewayURL)
}
