package cardano

import (
	"context"
	"encoding/hex"
	"fmt"
	"github.com/coinbase/rosetta-sdk-go/keys"
	"github.com/coinbase/rosetta-sdk-go/types"
	"testing"
)

func TestDeriveAddress(t *testing.T) {
	client := NewClient()
	ctx := context.Background()
	var primaryNetwork *types.NetworkIdentifier
	{
		request := &types.MetadataRequest{}
		networkList, rosettaErr, err := client.NetworkAPI.NetworkList(ctx, request)
		if err != nil {
			panic(err)
		}
		if rosettaErr != nil {
			panic(err)
		}
		primaryNetwork = networkList.NetworkIdentifiers[0]
	}

	{
		privateKey, err := keys.ImportPrivateKey("0858d5d6bb5ec1d25e4f719ca44190db2e9d18cbe8a803d06231cb3628ef56b8", types.Edwards25519)
		if err != nil {
			panic(err)
		}
		request := &types.ConstructionDeriveRequest{
			NetworkIdentifier: primaryNetwork,
			PublicKey:         privateKey.PublicKey,
			Metadata:          nil,
		}
		response, rerr, err := client.ConstructionAPI.ConstructionDerive(ctx, request)
		if err != nil {
			panic(err)
		}
		if rerr != nil {
			panic(err)
		}
		fmt.Printf("address: %s\n", response.AccountIdentifier.Address)
	}
}

func TestDeriveAddress1(t *testing.T) {
	client := NewClient()
	ctx := context.Background()
	var primaryNetwork *types.NetworkIdentifier
	{
		request := &types.MetadataRequest{}
		networkList, rosettaErr, err := client.NetworkAPI.NetworkList(ctx, request)
		if err != nil {
			panic(err)
		}
		if rosettaErr != nil {
			panic(err)
		}
		primaryNetwork = networkList.NetworkIdentifiers[0]
	}

	{
		xx, _ := keys.ImportPrivateKey("dee4abe45a948a9888ba5c21b53aaf72e1ee91f17f9c821718ae7665daf1cc5d", types.Edwards25519)
		data, _ := hex.DecodeString("e43adee176b10c252a0d00713b807c4ea0a030193cb561af3eb9915109bcbb96")
		pubkey := &types.PublicKey{
			Bytes:     data,
			CurveType: types.Edwards25519,
		}
		//pubkey = xx.PublicKey
		request := &types.ConstructionDeriveRequest{
			NetworkIdentifier: primaryNetwork,
			PublicKey:         pubkey,
			Metadata:          nil,
		}
		pubkey = xx.PublicKey
		response, rerr, err := client.ConstructionAPI.ConstructionDerive(ctx, request)
		if err != nil {
			panic(err)
		}
		if rerr != nil {
			panic(err)
		}
		fmt.Printf("address: %s\n", response.AccountIdentifier.Address)
	}
}


func TestDeriveAddress2(t *testing.T) {
	client := NewClient()
	ctx := context.Background()
	var primaryNetwork *types.NetworkIdentifier
	{
		request := &types.MetadataRequest{}
		networkList, rosettaErr, err := client.NetworkAPI.NetworkList(ctx, request)
		if err != nil {
			panic(err)
		}
		if rosettaErr != nil {
			panic(err)
		}
		primaryNetwork = networkList.NetworkIdentifiers[0]
	}

	for i := 0;i < 10;i ++ {
		keypair, err := keys.GenerateKeypair(types.Edwards25519)
		if err != nil {
			panic(err)
		}
		prikey := hex.EncodeToString(keypair.PrivateKey)
		pubkey := hex.EncodeToString(keypair.PublicKey.Bytes)

		request := &types.ConstructionDeriveRequest{
			NetworkIdentifier: primaryNetwork,
			PublicKey:         keypair.PublicKey,
			Metadata:          nil,
		}
		response, rerr, err := client.ConstructionAPI.ConstructionDerive(ctx, request)
		if err != nil {
			panic(err)
		}
		if rerr != nil {
			panic(err)
		}
		fmt.Printf("prikey: %s pubkey: %s address: %s\n", prikey, pubkey, response.AccountIdentifier.Address)
	}
}

func TestBuildTransaction(t *testing.T) {
	client := NewClient()
	ctx := context.Background()
	var primaryNetwork *types.NetworkIdentifier
	{
		request := &types.MetadataRequest{}
		networkList, rosettaErr, err := client.NetworkAPI.NetworkList(ctx, request)
		if err != nil {
			panic(err)
		}
		if rosettaErr != nil {
			panic(err)
		}
		primaryNetwork = networkList.NetworkIdentifiers[0]
	}
	var preProcess *types.ConstructionPreprocessResponse
	operations := make([]*types.Operation, 0)
	{
		adaCurrency := &types.Currency{
			Symbol:   "ADA",
			Decimals: 6,
			Metadata: nil,
		}
		inputOperationIdentifier := &types.OperationIdentifier{
			Index:        0,
			NetworkIndex: nil,
		}
		inputAccount := &types.AccountIdentifier{
			Address:    "Ae2tdPwUPEZ3DdaWu8jn553npu6jwEPAJiahruj3xQjPXxgoxfYDWusJz7x",
			SubAccount: nil,
			Metadata:   nil,
		}
		amount := &types.Amount{
			Value:    "1000000",
			Currency: adaCurrency,
			Metadata: nil,
		}
		coinIdentifier := &types.CoinIdentifier{Identifier: ""}
		coinChange := &types.CoinChange{
			CoinIdentifier: coinIdentifier,
			CoinAction:     "coin_spent",
		}
		operations = append(operations, &types.Operation{
			OperationIdentifier: inputOperationIdentifier,
			RelatedOperations:   nil,
			Type:                "input",
			Status:              nil,
			Account:             inputAccount,
			Amount:              amount,
			CoinChange:          coinChange,
			Metadata:            nil,
		})
		outputOperationIdentifier := &types.OperationIdentifier{
			Index:        1,
			NetworkIndex: nil,
		}
		relatedOperations := make([]*types.OperationIdentifier, 0)
		relatedOperations = append(relatedOperations, inputOperationIdentifier)
		outputAccount := &types.AccountIdentifier{
			Address:    "Ae2tdPwUPEZ3DdaWu8jn553npu6jwEPAJiahruj3xQjPXxgoxfYDWusJz7x",
			SubAccount: nil,
			Metadata:   nil,
		}
		operations = append(operations, &types.Operation{
			OperationIdentifier: outputOperationIdentifier,
			RelatedOperations:   relatedOperations,
			Type:                "output",
			Status:              nil,
			Account:             outputAccount,
			Amount:              amount,
			CoinChange:          coinChange,
			Metadata:            nil,
		})
		request := &types.ConstructionPreprocessRequest{
			NetworkIdentifier: primaryNetwork,
			Operations:        operations,
		}
		preProcess, rosettaErr, err := client.ConstructionAPI.ConstructionPreprocess(ctx, request)
		if err != nil {
			panic(err)
		}
		if rosettaErr != nil {
			panic(rosettaErr)
		}
		fmt.Printf("%d\n", len(preProcess.RequiredPublicKeys))
	}
	var metadata *types.ConstructionMetadataResponse
	{
		request := &types.ConstructionMetadataRequest{
			NetworkIdentifier: primaryNetwork,
			Options:           preProcess.Options,
		}
		metadata, rosettaErr, err := client.ConstructionAPI.ConstructionMetadata(ctx, request)
		if err != nil {
			panic(err)
		}
		if rosettaErr != nil {
			panic(rosettaErr)
		}
		fmt.Printf("fee: %s", metadata.SuggestedFee[0].Value)
	}
	var payloadRsp *types.ConstructionPayloadsResponse
	{
		request := &types.ConstructionPayloadsRequest{
			NetworkIdentifier: primaryNetwork,
			Operations:        operations,
			Metadata:          metadata.Metadata,
		}
		payloadRsp, rosettaErr, err := client.ConstructionAPI.ConstructionPayloads(ctx, request)
		if err != nil {
			panic(err)
		}
		if rosettaErr != nil {
			panic(rosettaErr)
		}
		fmt.Printf("payloads: %d, tx hash: %s", len(payloadRsp.Payloads), payloadRsp.UnsignedTransaction)
	}
	var combine *types.ConstructionCombineResponse
	{
		request := &types.ConstructionCombineRequest{
			NetworkIdentifier:   primaryNetwork,
			UnsignedTransaction: payloadRsp.UnsignedTransaction,
			Signatures:          nil,
		}
		combine, rosettaErr, err := client.ConstructionAPI.ConstructionCombine(ctx, request)
		if err != nil {
			panic(err)
		}
		if rosettaErr != nil {
			panic(rosettaErr)
		}
		fmt.Printf("signed tx: %s", combine.SignedTransaction)
	}
	{
		request := &types.ConstructionSubmitRequest{
			NetworkIdentifier: primaryNetwork,
			SignedTransaction: combine.SignedTransaction,
		}
		submit, rosettaErr, err := client.ConstructionAPI.ConstructionSubmit(ctx, request)
		if err != nil {
			panic(err)
		}
		if rosettaErr != nil {
			panic(rosettaErr)
		}
		fmt.Printf("signed tx: %s", submit.TransactionIdentifier.Hash)
	}
}


func TestBuildTransaction1(t *testing.T) {
	client := NewClient()
	ctx := context.Background()
	var primaryNetwork *types.NetworkIdentifier
	{
		request := &types.MetadataRequest{}
		networkList, rosettaErr, err := client.NetworkAPI.NetworkList(ctx, request)
		if err != nil {
			panic(err)
		}
		if rosettaErr != nil {
			panic(err)
		}
		primaryNetwork = networkList.NetworkIdentifiers[0]
	}
	var preProcess *types.ConstructionPreprocessResponse
	operations := make([]*types.Operation, 0)
	{
		adaCurrency := &types.Currency{
			Symbol:   "ADA",
			Decimals: 6,
			Metadata: nil,
		}
		inputOperationIdentifier := &types.OperationIdentifier{
			Index:        0,
		}
		inputAccount := &types.AccountIdentifier{
			Address:    "addr_test1qzgd37t7vwufes5xmfltc2ccy8eu5nrrtfg33434f8qp0kdzql220qjjr7klj3z0zxh3dufhg43ne29znkcjxqtrarmqkzazsk",
		}
		coinIdentifier := &types.CoinIdentifier{Identifier: "9c543d15330d3b8710c1d26b61de999fac74944a6997582bbca9ba48981e0336:0"}
		coinChange := &types.CoinChange{
			CoinIdentifier: coinIdentifier,
			CoinAction: "coin_spent",
		}
		amount := &types.Amount{
			Value:    "-1000000",
			Currency: adaCurrency,
			Metadata: nil,
		}
		operations = append(operations, &types.Operation{
			OperationIdentifier: inputOperationIdentifier,
			RelatedOperations:   nil,
			Type:                "input",
			Status:              nil,
			Account:             inputAccount,
			Amount:              amount,
			CoinChange:          coinChange,
			Metadata:            nil,
		})
		outputOperationIdentifier := &types.OperationIdentifier{
			Index:        1,
		}
		relatedOperations := make([]*types.OperationIdentifier, 0)
		relatedOperations = append(relatedOperations, inputOperationIdentifier)
		outputAccount := &types.AccountIdentifier{
			Address:    "addr_test1qp8pcm3klc8ne2tw2nrfvx48jexdt9897vxtxwxtyhvqtfazql220qjjr7klj3z0zxh3dufhg43ne29znkcjxqtrarmq3jmqal",
		}
		outAmount := &types.Amount{
			Value:    "1000000",
			Currency: adaCurrency,
			Metadata: nil,
		}
		operations = append(operations, &types.Operation{
			OperationIdentifier: outputOperationIdentifier,
			RelatedOperations:   relatedOperations,
			Type:                "output",
			Status:              nil,
			Account:             outputAccount,
			Amount:              outAmount,
			CoinChange:          nil,
			Metadata:            nil,
		})
		request := &types.ConstructionPreprocessRequest{
			NetworkIdentifier: primaryNetwork,
			Operations:        operations,
		}
		preProcess, rosettaErr, err := client.ConstructionAPI.ConstructionPreprocess(ctx, request)
		if err != nil {
			panic(err)
		}
		if rosettaErr != nil {
			panic(rosettaErr)
		}
		fmt.Printf("%d\n", len(preProcess.RequiredPublicKeys))
	}
	var metadata *types.ConstructionMetadataResponse
	{
		request := &types.ConstructionMetadataRequest{
			NetworkIdentifier: primaryNetwork,
			Options:           preProcess.Options,
		}
		metadata, rosettaErr, err := client.ConstructionAPI.ConstructionMetadata(ctx, request)
		if err != nil {
			panic(err)
		}
		if rosettaErr != nil {
			panic(rosettaErr)
		}
		fmt.Printf("fee: %s", metadata.SuggestedFee[0].Value)
	}
	var payloadRsp *types.ConstructionPayloadsResponse
	{
		request := &types.ConstructionPayloadsRequest{
			NetworkIdentifier: primaryNetwork,
			Operations:        operations,
			Metadata:          metadata.Metadata,
		}
		payloadRsp, rosettaErr, err := client.ConstructionAPI.ConstructionPayloads(ctx, request)
		if err != nil {
			panic(err)
		}
		if rosettaErr != nil {
			panic(rosettaErr)
		}
		fmt.Printf("payloads: %d, tx hash: %s", len(payloadRsp.Payloads), payloadRsp.UnsignedTransaction)
	}
	var combine *types.ConstructionCombineResponse
	{
		request := &types.ConstructionCombineRequest{
			NetworkIdentifier:   primaryNetwork,
			UnsignedTransaction: payloadRsp.UnsignedTransaction,
			Signatures:          nil,
		}
		combine, rosettaErr, err := client.ConstructionAPI.ConstructionCombine(ctx, request)
		if err != nil {
			panic(err)
		}
		if rosettaErr != nil {
			panic(rosettaErr)
		}
		fmt.Printf("signed tx: %s", combine.SignedTransaction)
	}
	{
		request := &types.ConstructionSubmitRequest{
			NetworkIdentifier: primaryNetwork,
			SignedTransaction: combine.SignedTransaction,
		}
		submit, rosettaErr, err := client.ConstructionAPI.ConstructionSubmit(ctx, request)
		if err != nil {
			panic(err)
		}
		if rosettaErr != nil {
			panic(rosettaErr)
		}
		fmt.Printf("signed tx: %s", submit.TransactionIdentifier.Hash)
	}
}
