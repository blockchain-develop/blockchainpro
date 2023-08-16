package cardano

import (
	"context"
	"encoding/hex"
	json2 "encoding/json"
	"fmt"

	//"github.com/btcsuite/btcutil/bech32"
	"github.com/coinbase/rosetta-sdk-go/keys"
	"github.com/coinbase/rosetta-sdk-go/types"
	"github.com/echovl/bech32"
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

	//address := ""
	//publickey := make([]byte, 0)
	// 399b0e5af500166bee42594787f1a8faad99062f0a34621f9aa3c27b033d31c8
	{
		privateKey, err := keys.ImportPrivateKey("3b596c94bdc0aab357f0aab5193dadffe7c2a0c9bf8f6a6c45969763c4eca6bd", types.Edwards25519)
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
		//address = response.AccountIdentifier.Address
		//publickey = privateKey.PublicKey.Bytes
		fmt.Printf("address: %s\n", response.AccountIdentifier.Address)
		fmt.Printf("pubkey: %s\n", hex.EncodeToString(request.PublicKey.Bytes))
	}
	/*
	{
		addressBytes := make([]byte, 29)
		header := 0x60 | (byte(1) & 0xFF)
		hash, err := blake2b.New(224/8, nil)
		if err != nil {
			panic(err)
		}

		hash.Write(publickey[:32])
		paymentHash := hash.Sum(nil)

		addressBytes[0] = header
		copy(addressBytes[1:], paymentHash)

		hrp := "addr"
		address1, err := bech32.EncodeFromBase256(hrp, addressBytes)
		if err != nil {
			panic(err)
		}
		fmt.Printf("address: %s\n", address1)
	}

	 */

	/*
	{
		hrp, xxx, err := bech32.Decode(address)
		if err != nil {
			panic(err)
		}
		decoded, err := bech32.ConvertBits(xxx, 5, 8, false)
		if err != nil {
			panic(err)
		}
	 */

	/*
	hrp, decoded, err := bech32.DecodeToBase256("addr_xvk135hqmkaqydnxnq6wmjkkhasvwjprpnqnzsrwwes6mql45enlcsqs0vz4xasvja4qsvrw93v3gc7mmep76cf67a4yffdrsvuxqm9qhwqlvnay5")

	//hrp, decoded, err := bech32.DecodeToBase256(address)
	if err != nil {
		panic(err)
	}
	fmt.Printf("hrp: %s, data: %s\n", hrp, hex.EncodeToString(decoded))


	if len(decoded) != 29 && len(decoded) != 64{
		panic(err)
	}
	//
	header := 0x60 | (byte(0) & 0xFF)
	if decoded[0] != header {
		panic(err)
	}
	if hrp != "addr" {
		panic(err)
	}
*/
}


func TestDeriveAddressxx(t *testing.T) {
	addressBytes := make([]byte, 29)
	header := 0x60 | (byte(1) & 0xFF)

	/*
	hash, err := blake2b.New(224/8, nil)
	if err != nil {
		panic(err)
	}


	hash.Write(publickey[:32])
	paymentHash := hash.Sum(nil)

	 */
	paymentHash, _ := hex.DecodeString("12ee66f8356d94944c8a8fb3cb663d9648e2ee1522f532ab622cfa61")

	addressBytes[0] = header
	copy(addressBytes[1:], paymentHash)

	hrp := "addr"
	address1, err := bech32.EncodeFromBase256(hrp, addressBytes)
	if err != nil {
		panic(err)
	}
	fmt.Printf("address: %s\n", address1)
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

func TestGenerateAccount(t *testing.T) {
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

	var key *keys.KeyPair
	{
		var keypair *keys.KeyPair
		var err error
		if true {
			keypair, err = keys.GenerateKeypair(types.Edwards25519)
			if err != nil {
				panic(err)
			}
		}
		prikey := hex.EncodeToString(keypair.PrivateKey)
		pubkey := hex.EncodeToString(keypair.PublicKey.Bytes)
		fmt.Printf("private key: %s, public key: %s\n", prikey, pubkey)
		key = keypair
	}
	{
		request := &types.ConstructionDeriveRequest{
			NetworkIdentifier: primaryNetwork,
			PublicKey:         key.PublicKey,
			Metadata:          nil,
		}
		derive, rerr, err := client.ConstructionAPI.ConstructionDerive(ctx, request)
		if err != nil {
			panic(err)
		}
		if rerr != nil {
			panic(err)
		}
		fmt.Printf("address: %s\n", derive.AccountIdentifier.Address)
	}
}

func TestBuildTransaction(t *testing.T) {
	client := NewClient()
	ctx := context.Background()
	var rosettaErr *types.Error
	var err error
	var primaryNetwork *types.NetworkIdentifier
	{
		request := &types.MetadataRequest{}
		var networkList *types.NetworkListResponse
		networkList, rosettaErr, err = client.NetworkAPI.NetworkList(ctx, request)
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
		// inputs
		adaCurrency := &types.Currency{
			Symbol:   "ADA",
			Decimals: 6,
			Metadata: nil,
		}
		inputOperationIdentifier := &types.OperationIdentifier{
			Index:        0,
		}
		inputAccount := &types.AccountIdentifier{
			Address:    "addr_test1vqc9yc3n6arjq4fs9lmpqcjzn9x87z98nxshw7c9ymhequsc4egp4",
		}
		coinIdentifier := &types.CoinIdentifier{
			Identifier: "8de70f0addead859cef16b8d3da6fc06ee21a51508e2384c10cbfc152aa8e4b7:1",
		}
		coinChange := &types.CoinChange{
			CoinIdentifier: coinIdentifier,
			CoinAction: "coin_spent",
		}
		amount := &types.Amount{
			Value:    "-8834000",
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
		//
		relatedOperations := make([]*types.OperationIdentifier, 0)
		relatedOperations = append(relatedOperations, inputOperationIdentifier)
		// outputs
		outputOperationIdentifier := &types.OperationIdentifier{
			Index:        1,
		}
		outputAccount := &types.AccountIdentifier{
			Address:    "addr_test1vppe5dy9j5clrxgc38p7x9gc7zxx43et4cg46m6uulppk5c2uhh2k",
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
		// charge
		chargeOperationIdentifier := &types.OperationIdentifier{
			Index:        2,
		}
		chargeAccount := &types.AccountIdentifier{
			Address:    "addr_test1vqc9yc3n6arjq4fs9lmpqcjzn9x87z98nxshw7c9ymhequsc4egp4",
		}
		chargeAmount := &types.Amount{
			Value:    "7668000",
			Currency: adaCurrency,
			Metadata: nil,
		}
		operations = append(operations, &types.Operation{
			OperationIdentifier: chargeOperationIdentifier,
			RelatedOperations:   relatedOperations,
			Type:                "output",
			Status:              nil,
			Account:             chargeAccount,
			Amount:              chargeAmount,
			CoinChange:          nil,
			Metadata:            nil,
		})
		request := &types.ConstructionPreprocessRequest{
			NetworkIdentifier: primaryNetwork,
			Operations:        operations,
		}
		preProcess, rosettaErr, err = client.ConstructionAPI.ConstructionPreprocess(ctx, request)
		if err != nil {
			panic(err)
		}
		if rosettaErr != nil {
			panic(rosettaErr)
		}
		fmt.Printf("require keys: %d\n", len(preProcess.RequiredPublicKeys))
	}
	var metadata *types.ConstructionMetadataResponse
	{
		request := &types.ConstructionMetadataRequest{
			NetworkIdentifier: primaryNetwork,
			Options:           preProcess.Options,
		}
		metadata, rosettaErr, err = client.ConstructionAPI.ConstructionMetadata(ctx, request)
		if err != nil {
			panic(err)
		}
		if rosettaErr != nil {
			panic(rosettaErr)
		}
		fmt.Printf("fee: %s\n", metadata.SuggestedFee[0].Value)
	}
	var payloadRsp *types.ConstructionPayloadsResponse
	{
		request := &types.ConstructionPayloadsRequest{
			NetworkIdentifier: primaryNetwork,
			Operations:        operations,
			Metadata:          metadata.Metadata,
		}
		payloadRsp, rosettaErr, err = client.ConstructionAPI.ConstructionPayloads(ctx, request)
		if err != nil {
			panic(err)
		}
		if rosettaErr != nil {
			panic(rosettaErr)
		}
		fmt.Printf("payloads: %d, raw tx: %s\n", len(payloadRsp.Payloads), payloadRsp.UnsignedTransaction)
	}
	// signature
	privateKey := "3b596c94bdc0aab357f0aab5193dadffe7c2a0c9bf8f6a6c45969763c4eca6bd"
	keypair, err := keys.ImportPrivateKey(privateKey, types.Edwards25519)
	if err != nil {
		panic(err)
	}
	signer, err := keypair.Signer()
	if err != nil {
		panic(err)
	}
	signature, err := signer.Sign(payloadRsp.Payloads[0], types.Ed25519)
	if err != nil {
		panic(err)
	}
	err = signer.Verify(signature)
	if err != nil {
		fmt.Printf("err: %s\n", err.Error())
		panic(err)
	}
	//
	var combine *types.ConstructionCombineResponse
	{
		request := &types.ConstructionCombineRequest{
			NetworkIdentifier:   primaryNetwork,
			UnsignedTransaction: payloadRsp.UnsignedTransaction,
			Signatures:          []*types.Signature{signature},
		}
		combine, rosettaErr, err = client.ConstructionAPI.ConstructionCombine(ctx, request)
		if err != nil {
			panic(err)
		}
		if rosettaErr != nil {
			panic(rosettaErr)
		}
		fmt.Printf("signed tx: %s\n", combine.SignedTransaction)
	}
	{
		request := &types.ConstructionSubmitRequest{
			NetworkIdentifier: primaryNetwork,
			SignedTransaction: combine.SignedTransaction,
		}
		var submit *types.TransactionIdentifierResponse
		submit, rosettaErr, err = client.ConstructionAPI.ConstructionSubmit(ctx, request)
		if err != nil {
			panic(err)
		}
		if rosettaErr != nil {
			panic(rosettaErr)
		}
		fmt.Printf("signed tx hash: %s\n", submit.TransactionIdentifier.Hash)
	}
}



func TestBuildTransaction_SendMyCoin(t *testing.T) {
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

func TestDecodeTransaction(t *testing.T) {
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
		transaction := "8279062238336134303038353832353832306130613039633236343536303864653933343132396635623162316566636434643864613364613834633861336137616461333666313465326539366332303030303832353832303933636166316436386334393433333364346237666464653437313139313537376539323262306434653533663135666161353938383564353538656362663930303832353832303132663837386266323130643363396137653337613036343532363462666266303264306466396137373535363364343639313936636439666633313837343730303832353832306363656462616631396132626636653236663937386361356539646264336131363962616538333466323261613631333235303038626534393665626230383530303832353832303366643538396331623666376163376438346135623536653138353633623235663637336561323566613036336239326330616631353132613830316430373730313031383238323538316436316439386637396636633035336230333464646530393038636631666337316533643263393734303338333335303332393532666636346131316132613333306563663832353831643631656637323239653830393533623766623435656362653234353339333565643264363830383565646139636134633862316138636234373331623030303030303031663339643132616330323161303030333063343130333161303338346334613361313030383538323538323035636638383535383863386632393562323163646234633533396261636132313933306334623365333239643430333132616262653763663635343539356166353834303932663734363733666431633661326365633463353462323664613563346535333630393063656163393239383131326635653062363431393661653332353461373934656334323335656136333366326635366663306532353164666439373665613237373835643661643561653565306332333663656566626531323066383235383230343337353233363461303239373031363030636230616461376232333634366336623139363930353631363338376332343337393531306139373432653765313538343038303932666161363338623338353265626364663137356662366436633036383235353439333832613836393961643862363533626431333738653334383939383466353766366336386162643838353730383966326131663735663831393338333766653139613434623233313364366665333833383133313335623630613832353832303039353932633833663633383363343061323562616332346563373538346561663332626135653764303935363536343562383863613262366337626230336135383430306439613831333433653731323161613964313537376132663766646534653336313238646536326164333236323565323234303265396638666266393439643332393165373137326461663734623562303639343032316139383534623138636331653330323362373964326665343130343763666661306232623536306638323538323032363130613837326333643030326236373361356631356636373632643664633062303430643637323532666639353535613431313061396538313435663837353834303734623630613062313538643830623865343532373861333161323765313034623833316636316138326438333534333530383362636236656161363062653732353838636335653163646532646163646465343161393739626636346464656230653132633865363430623035636539393964636566653162613433353034383235383230663233613639343734303536343161643537323564366438366133653233633538323064616431336261323563363533373834653065333665336236356237373538343032383461633764633939323834303164306566393437396133303330653233373437346234323032316666333631616139656133353334613136636436353132666437383562376464326234383965386634613631366365383232323535656366363330363863613836343463366232313639343739626362663536323530306636a16a6f7065726174696f6e7385a5746f7065726174696f6e5f6964656e746966696572a165696e64657800647479706565696e707574676163636f756e74a16761646472657373783a6164647231763966787a38356132676b3779357667726361706b77707273716e616e70677a67347a6e323679656e7273733861637a6c6e6a6e6e66616d6f756e74a26576616c7565692d39303136373038356863757272656e6379a26673796d626f6c6341444168646563696d616c73066b636f696e5f6368616e6765a26f636f696e5f6964656e746966696572a16a6964656e7469666965727842613061303963323634353630386465393334313239663562316231656663643464386461336461383463386133613761646133366631346532653936633230303a306b636f696e5f616374696f6e6a636f696e5f7370656e74a5746f7065726174696f6e5f6964656e746966696572a165696e64657801647479706565696e707574676163636f756e74a16761646472657373783a616464723176386364377932713435366e35796467687839327539656168637167676b716a6a356376667161643677733666726764673271757066616d6f756e74a26576616c7565692d35313036373039386863757272656e6379a26673796d626f6c6341444168646563696d616c73066b636f696e5f6368616e6765a26f636f696e5f6964656e746966696572a16a6964656e7469666965727842393363616631643638633439343333336434623766646465343731313931353737653932326230643465353366313566616135393838356435353865636266393a306b636f696e5f616374696f6e6a636f696e5f7370656e74a5746f7065726174696f6e5f6964656e746966696572a165696e64657802647479706565696e707574676163636f756e74a16761646472657373783a6164647231767868306474677574766d73797a32646163353632776b657132336d6d7a61677733706e37347a373530663368797130783767327966616d6f756e74a26576616c7565682d343036313737376863757272656e6379a26673796d626f6c6341444168646563696d616c73066b636f696e5f6368616e6765a26f636f696e5f6964656e746966696572a16a6964656e7469666965727842313266383738626632313064336339613765333761303634353236346266626630326430646639613737353536336434363931393663643966663331383734373a306b636f696e5f616374696f6e6a636f696e5f7370656e74a5746f7065726174696f6e5f6964656e746966696572a165696e64657803647479706565696e707574676163636f756e74a16761646472657373783a6164647231767970726e373778766537766a636c3663796730757975707377783936797639646c686d396d346e63767470716e7177676466367a66616d6f756e74a26576616c7565692d32353334313530396863757272656e6379a26673796d626f6c6341444168646563696d616c73066b636f696e5f6368616e6765a26f636f696e5f6964656e746966696572a16a6964656e7469666965727842636365646261663139613262663665323666393738636135653964626433613136396261653833346632326161363133323530303862653439366562623038353a306b636f696e5f616374696f6e6a636f696e5f7370656e74a5746f7065726174696f6e5f6964656e746966696572a165696e64657804647479706565696e707574676163636f756e74a16761646472657373783a616464723176386868793230677039666d30373639616a6c7a6735756e746d666464717939616b3575356e797472327874677563686a6e77703066616d6f756e74a26576616c75656b2d383931393637363139316863757272656e6379a26673796d626f6c6341444168646563696d616c73066b636f696e5f6368616e6765a26f636f696e5f6964656e746966696572a16a6964656e7469666965727842336664353839633162366637616337643834613562353665313835363362323566363733656132356661303633623932633061663135313261383031643037373a316b636f696e5f616374696f6e6a636f696e5f7370656e74"
		request := &types.ConstructionParseRequest{
			NetworkIdentifier: primaryNetwork,
			Signed: true,
			Transaction: transaction,
		}
		submit, rosettaErr, err := client.ConstructionAPI.ConstructionParse(ctx, request)
		if err != nil {
			panic(err)
		}
		if rosettaErr != nil {
			panic(rosettaErr)
		}
		txjson, _ := json2.MarshalIndent(submit, "", "    ")
		fmt.Printf("signed tx: %s", string(txjson))
	}
}
