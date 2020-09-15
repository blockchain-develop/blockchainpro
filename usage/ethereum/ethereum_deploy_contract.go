package ethereum

import (
	"context"
	"fmt"
	"github.com/blockchainpro/usage/ethereum/contractabi/erc20_abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/crypto"
	"math/big"
	"testing"
)

func TestDeployContract(t *testing.T) {
	client := DefaultEthereumClient()
	privateKey := DefaultPrivateKey()
	address := crypto.PubkeyToAddress(privateKey.PublicKey)
	ctx := context.Background()

	auth := bind.NewKeyedTransactor(privateKey)
	auth.Nonce = big.NewInt(int64(client.GetNonceAt(ctx, address)))
	auth.Value = big.NewInt(0)      // in wei
	auth.GasLimit = uint64(5000000) // in units
	auth.GasPrice = big.NewInt(30000000000)
	addr, tx, contract, err := erc20_abi.DeployERC20(auth, client.Client)
	if err != nil {
		fmt.Printf("deploy contract - err: %s", err.Error())
	}
	fmt.Printf("deploy contract - the user address of the new contract is: %s", addr.String())
	fmt.Printf("deploy contract - tx: %s", tx.Hash().String())
	fmt.Printf("deploy contract - contract: %v", contract)
}
