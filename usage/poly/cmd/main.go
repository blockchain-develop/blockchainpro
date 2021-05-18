package main

import (
	"encoding/hex"
	"fmt"
	"github.com/ontio/ontology-crypto/signature"
	"github.com/polynetwork/poly-go-sdk"
	"github.com/polynetwork/poly/common"
	"github.com/urfave/cli"
	"os"
	"runtime"
	"strings"
)

var (
	walletPathFlag = cli.StringFlag{
		Name:  "wallet",
		Usage: "wallet file `<path>`",
		Value: "",
	}

	walletPasswordFlag = cli.StringFlag{
		Name:  "password",
		Usage: "wallet password",
		Value: "123456",
	}

	hashFlag = cli.StringFlag{
		Name:  "hash",
		Usage: "hash",
		Value: "5bc73c6794f63d6508ea447f5373026299f5dd8cafdd6ed73c3519d4c029ba0b",
	}
)

//getFlagName deal with short flag, and return the flag name whether flag name have short name
func getFlagName(flag cli.Flag) string {
	name := flag.GetName()
	if name == "" {
		return ""
	}
	return strings.TrimSpace(strings.Split(name, ",")[0])
}

func setupApp() *cli.App {
	app := cli.NewApp()
	app.Usage = "sign hash"
	app.Action = StartServer
	app.Version = "1.0.0"
	app.Copyright = "Copyright in 2019 The Ontology Authors"
	app.Flags = []cli.Flag{
		walletPathFlag,
		hashFlag,
		walletPasswordFlag,
	}
	app.Commands = []cli.Command{}
	app.Before = func(context *cli.Context) error {
		runtime.GOMAXPROCS(runtime.NumCPU())
		return nil
	}
	return app
}

func StartServer(ctx *cli.Context) {
	walletPath := ctx.GlobalString(getFlagName(walletPathFlag))
	walletPassword := ctx.GlobalString(getFlagName(walletPasswordFlag))
	hash := ctx.GlobalString(getFlagName(hashFlag))

	var wallet *poly_go_sdk.Wallet
	var err error
	if !common.FileExisted(walletPath) {
		panic("wallet file is not exist!")
	}
	wallet, err = poly_go_sdk.OpenWallet(walletPath)
	if err != nil {
		panic(err)
	}

	signer, err := wallet.GetDefaultAccount([]byte(walletPassword))
	if err != nil || signer == nil {
		panic(err)
	}
	fmt.Printf("wallet address address: %s", signer.Address.ToBase58())

	//
	hash_hex, _ := hex.DecodeString(hash)
	sigData, err := signature.Sign(signature.SHA256withECDSA, signer.PrivateKey, hash_hex, nil)
	if err != nil {
		panic(err)
	}
	data, err := signature.Serialize(sigData)
	if err != nil {
		panic(err)
	}
	fmt.Printf("the sign data: %s\n", hex.EncodeToString(data))
}

func main() {
	if err := setupApp().Run(os.Args); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
