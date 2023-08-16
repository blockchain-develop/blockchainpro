package xrp

import (
	"encoding/json"
	"fmt"
	"github.com/rubblelabs/ripple/crypto"
	"github.com/rubblelabs/ripple/data"
	"strconv"
	"testing"
)

func TestSendTransaction(t *testing.T) {
	priv, err := crypto.NewEd25519Key([]byte{1,2,3,4,5,6})
	if err != nil {
		panic(err)
	}
	account, err := crypto.AccountId(priv, nil)
	if err != nil {
		panic(err)
	}
	fromAddr := account.String()
	fmt.Printf("addr: %s\n", fromAddr)
	aaa, err := crypto.AccountPublicKey(priv, nil)
	if err != nil {
		panic(err)
	}
	pubkey := data.PublicKey{}
	copy(pubkey[:], aaa.Payload())
	//
	from, err := data.NewAccountFromAddress(fromAddr)
	if err != nil {
		panic(err)
	}
	//
	toAddr := "rPCSYgr539f81tDrVSKCyLnBW4yBqdCpAa"
	to, err := data.NewAccountFromAddress(toAddr)
	if err != nil {
		panic(err)
	}
	amount, err := data.NewAmount(int64(1000000000))
	if err != nil {
		panic(err)
	}
	feeResp, err := getFee(MainNetUrl)
	if err != nil {
		panic(err)
	}
	feeAmt, _ := strconv.Atoi(feeResp.Result.Drops.MedianFee)
	fee, err := data.NewNativeValue(int64(feeAmt))
	if err != nil {
		panic(err)
	}
	destTag := uint32(0)

	acctInfo, err := getAccount(MainNetUrl, from.String())
	if err != nil {
		panic(err)
	}
	nonce := uint32(acctInfo.Result.AccountData.Sequence)

	flag := data.TxPartialPayment

	pay := data.Payment{
		TxBase: data.TxBase{
			TransactionType: data.PAYMENT,
			Account:         *from,
			Sequence:        nonce,
			Fee:             *fee,
			Flags: &flag,
		},
		Destination:    *to,
		DestinationTag: &destTag,
		Amount:         *amount,
	}

	pay.InitialiseForSigning()
	pay.SigningPubKey = &pubkey
	//copy(pay.GetPublicKey().Bytes(), pay.GetPublicKey().Bytes())
	raw, msg, err := data.SigningHash(&pay)
	if err != nil {
		panic(err)
	}
	//
	sig, err := crypto.Sign(priv.Private(nil), raw.Bytes(), append(pay.SigningPrefix().Bytes(), msg...))
	if err != nil {
		panic(err)
	}
	*pay.TxnSignature = sig
	//
	hash, signedRaw, err := data.Raw(&pay)
	if err != nil {
		panic(err)
	}
	copy(pay.GetHash().Bytes(), hash.Bytes())
	//
	xxxx, _ := json.MarshalIndent(pay, "", "    ")
	fmt.Printf("xxxx: %s\n", xxxx)
	//
	req := submitParam{
		TxBlob: fmt.Sprintf("%X", signedRaw),
	}
	response, err := submitTransaction(MainNetUrl, &req)
	if err != nil {
		panic(err)
	}
	fmt.Printf("response: %v", response)
}
