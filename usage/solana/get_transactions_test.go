package solana

import (
	"bufio"
	"context"
	"fmt"
	"github.com/gagliardetto/solana-go"
	"github.com/gagliardetto/solana-go/rpc"
	"os"
	"testing"
)

func TestTransaction(t *testing.T) {
	client := rpc.New(rpc.MainNetBeta_RPC)
	ctx := context.Background()
	txids := []string{
		"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",
	}
	//
	filePath := "./solana.txt"
	file, err := os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		panic(err)
	}
	//及时关闭file句柄
	defer file.Close()
	//写入文件时，使用带缓存的 *Writer
	write := bufio.NewWriter(file)

	for _, txid := range txids {
		sig := solana.MustSignatureFromBase58(txid)
		tx, err := client.GetTransaction(ctx, sig, &rpc.GetTransactionOpts{
			Encoding:   solana.EncodingBase64,
			Commitment: rpc.CommitmentFinalized,
		})
		if err != nil {
			panic(err)
		}
		write.WriteString(fmt.Sprintf("# %s\n", txid))
		write.WriteString(fmt.Sprintf(`curl --location --request POST 'http://localhost:8087/api/v1/console/admin/syncer/rescan?start=%d&end=%d&chain=solana' \
--header 'Content-Type: application/json' \
--data-raw '{}'`, tx.Slot, tx.Slot+1))
		write.WriteString(fmt.Sprintf("\n"))
		write.WriteString(fmt.Sprintf("sleep 3"))
		write.WriteString(fmt.Sprintf("\n\n"))
	}

	write.Flush()
}
