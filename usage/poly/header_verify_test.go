package poly

import (
	"bytes"
	"crypto"
	"crypto/elliptic"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"github.com/btcsuite/btcd/btcec"
	crypto1 "github.com/ethereum/go-ethereum/crypto"
	"github.com/ontio/ontology-crypto/ec"
	"github.com/ontio/ontology-crypto/keypair"
	xxxx "github.com/ontio/ontology-crypto/signature"
	"github.com/ontio/ontology-crypto/sm2"
	"github.com/polynetwork/poly-go-sdk"
	"github.com/polynetwork/poly/common"
	"github.com/polynetwork/poly/consensus/vbft/config"
	"github.com/polynetwork/poly/core/signature"
	"github.com/polynetwork/poly/core/types"
	"golang.org/x/crypto/ed25519"
	"golang.org/x/crypto/ripemd160"
	"strings"
	"testing"
)

func TestDecodeBlock(t *testing.T) {
	dataStr := "000000000000000000000000f2afea7a31891478041541d7021a670d5d9562e0a5cbdc87a419ca44c7db56790000000000000000000000000000000000000000000000000000000000000000fa77cfe7aa243e028342eb94bee47e496fc91e6e616c3865f5dcd11d530d52321966783675357246c516d1e6bf4ee30a696becafcfc2775c1bfb489ae47bda469e1ea26024a77900c5888e2724b7c7e9fd12017b226c6561646572223a312c227672665f76616c7565223a2242475a654e616c526d743657686f645977536176772f654b415a47563579784e5a6b513448424549733045684b4b444363624a427757584337415a716d5057444a4c456373546f456277504b64414830346b30386c75343d222c227672665f70726f6f66223a224531724173346d325a4b55564643762f2f6139657359574678332b4c6e2b596f5371533456672b2f523364786e45673173574d5979566c35574332453375695266316674422f755078596470724231436a2f376e76673d3d222c226c6173745f636f6e6669675f626c6f636b5f6e756d223a373932303030302c226e65775f636861696e5f636f6e666967223a6e756c6c7d0000000000000000000000000000000000000000032312050309c6475ce07577ab72a1f96c263e5030cb53a843b00ca1238a093d9dcb183e2f231205032bed55e8c4d9cbc50657ff5909ee51dc394a92aad911c36bace83c4d6354079423120502e68a6e54bdfa0af47bd18465f4352f5151dc729c61a7399909f1cd1c6d816c020342011bc8a68bf6425b067190ab85b5a6a53d369707a5860f029a08113f0b47f40177f779d4371990bc9e64fe2610dad3907c3eb718c9df1951d204703e4698057cca3642011cc0f9204241315c73d045edd6a4770bd5e3470e096ed640c4a18e93097144d7585cf56581b7d27b822bde423c36dfac591b03c2dfc66ac73475d048a452ee674642011c1b78382150b0fd161a39ddcd5f880de61f3bed21f0f72b5affc6782fef9b56191d6f84be5a893e77dac7aa54adb80d33a12db575c435344291341830659badd800000000"
	dataHex, _ := hex.DecodeString(dataStr)

	/*
	rpcRsp := &client.JsonRpcResponse{}
	err := json.Unmarshal(dataHex, rpcRsp)
	if err != nil {
		panic(err)
	}
	if rpcRsp.Error != 0 {
		panic(rpcRsp.Error)
	}
*/
	/*
	block, err := utils.GetBlock(dataHex)
	if err != nil {
		panic(err)
	}
	*/

	block, err := types.BlockFromRawBytes(dataHex)
	if err != nil {
		panic(err)
	}
	fmt.Printf("block header message: %s\n", hex.EncodeToString(block.Header.GetMessage()))
}

func TestHeaderVerify(t *testing.T) {
	sdk := poly_go_sdk.NewPolySdk()
	sdk.NewRpcClient().SetAddress("http://13.92.155.62:20336")

	block0, err := sdk.GetBlockByHeight(0)
	if err != nil {
		fmt.Printf("handleBlockHeader - GetNodeHeader on height :%d failed", block0.Header.Height)
		return
	} else {
		fmt.Printf("block height: %d\n", block0.Header.Height)
	}

	block1, err := sdk.GetBlockByHeight(7972644)
	if err != nil {
		fmt.Printf("handleBlockHeader - GetNodeHeader on height :%d failed", block1.Header.Height)
		return
	} else {
		fmt.Printf("block height: %d\n", block1.Header.Height)
	}

	fmt.Printf("block 0 header message: %s\n", hex.EncodeToString(block0.Header.GetMessage()))
	fmt.Printf("block 1 header message: %s\n", hex.EncodeToString(block1.Header.GetMessage()))

	// verify
	block := block1
	temp := sha256.Sum256(block.Header.GetMessage())
	headerhash := sha256.Sum256(temp[:])
	if headerhash != block.Header.Hash() {
		fmt.Printf("hash is not right\n")
	} else {
		fmt.Printf("hash %s is right\n", hex.EncodeToString(headerhash[:]))
	}

	//
	blkInfo := &vconfig.VbftBlockInfo{}
	if err := json.Unmarshal(block0.Header.ConsensusPayload, blkInfo); err != nil {
		fmt.Printf("unmarshal blockInfo error: %s", err)
	}

	var bookkeepers  []keypair.PublicKey
	if blkInfo.NewChainConfig != nil {
		for _, peer := range blkInfo.NewChainConfig.Peers {
			aaa,_ := hex.DecodeString(peer.ID)
			key, _ := keypair.DeserializePublicKey(aaa)
			bookkeepers = append(bookkeepers, key)
		}
	}
	bookkeepers = keypair.SortPublicKeys(bookkeepers)
	fmt.Printf("publickeys number: %d\n", len(bookkeepers))
	for _,key := range bookkeepers {
		add := types.AddressFromPubKey(key)
		fmt.Printf("ont address: %s\n", add.ToHexString())
		fmt.Printf("ont compress public key: %s\n", hex.EncodeToString(getontcompresskey(key)))
		//fmt.Printf("ont no compress public key: %s\n", hex.EncodeToString(getontnocompresskey(key)))
		fmt.Printf("eth public key: %s\n", hex.EncodeToString(getethnocompresskey(key)))
		//ethaddress := ethcommon.BytesToAddress(crypto1.Keccak256(getethnocompresskey(key)[1:])[12:])
		//fmt.Printf("eth address: %s\n", hex.EncodeToString(ethaddress[:]))
	}

	err = signature.VerifyMultiSignature(headerhash[:], bookkeepers, 3, block.Header.SigData)
	if err != nil {
		fmt.Printf("verify multi signature failed, %s\n", err.Error())
	} else {
		fmt.Printf("verify multi signature successful!\n")
	}

	//
	var publickeys []byte
	for _,key := range bookkeepers {
		publickeys = append(publickeys,getontnocompresskey(key)...)
	}
	fmt.Printf("publickeys: %s\n", hex.EncodeToString(publickeys))

	//
	var sigs []byte
	for _, sig := range block.Header.SigData {
		newsig, _ := xxxx.ConvertToEthCompatible(sig)
		sigs = append(sigs, newsig...)

		h := crypto.SHA256.New()
		h.Write(headerhash[:])
		digest := h.Sum(nil)

		sigKey, err := crypto1.Ecrecover(digest, newsig)
		if err != nil {
			panic(err)
		}
		fmt.Printf("sig: %s, key: %s\n", hex.EncodeToString(sig), hex.EncodeToString(sigKey))
	}

	for _, keeper := range block.Header.Bookkeepers {
		add := types.AddressFromPubKey(keeper)
		fmt.Printf("block keeper ont address: %s %s\n", add.ToHexString(), add.ToBase58())
	}
	//fmt.Printf("sigs: %s\n", hex.EncodeToString(sigs))

	//
	xxx, err := types.AddressFromBookkeepers(bookkeepers)
	fmt.Printf("block1 keeper: %s\n",hex.EncodeToString(xxx[:]))

	//
	fmt.Printf("block0 next keeper: %s\n", hex.EncodeToString(block0.Header.NextBookkeeper[:]))

	sink := common.NewZeroCopySink(nil)
	sink.WriteUint16(uint16(len(bookkeepers)))
	for _, publickey := range bookkeepers {
		key := keypair.SerializePublicKey(publickey)
		sink.WriteVarBytes(key)
	}
	m := len(bookkeepers) -(len(bookkeepers) - 1) / 3
	sink.WriteUint16(uint16(m))

	fmt.Printf("book keeper: %s\n",  hex.EncodeToString(sink.Bytes()))

	var addr common.Address
	temp = sha256.Sum256(sink.Bytes())
	md := ripemd160.New()
	md.Write(temp[:])
	md.Sum(addr[:0])
	fmt.Printf("my block1 keeper: %s\n", hex.EncodeToString(addr[:]))

	if addr != block0.Header.NextBookkeeper {
		fmt.Printf("keeper failed!\n")
	} else {
		fmt.Printf("keeper successful!\n")
	}
}


func getontcompresskey(key keypair.PublicKey) []byte {
	var buf bytes.Buffer
	switch t := key.(type) {
	case *ec.PublicKey:
		switch t.Algorithm {
		case ec.ECDSA:
			// Take P-256 as a special case
			if t.Params().Name == elliptic.P256().Params().Name {
				return ec.EncodePublicKey(t.PublicKey, true)
			}
			buf.WriteByte(byte(0x12))
		case ec.SM2:
			buf.WriteByte(byte(0x13))
		}
		label, err := GetCurveLabel(t.Curve.Params().Name)
		if err != nil {
			panic(err)
		}
		buf.WriteByte(label)
		buf.Write(ec.EncodePublicKey(t.PublicKey, true))
	case ed25519.PublicKey:
		panic("err")
	default:
		panic("err")
	}
	return buf.Bytes()
}

func getontnocompresskey(key keypair.PublicKey) []byte {
	var buf bytes.Buffer
	switch t := key.(type) {
	case *ec.PublicKey:
		switch t.Algorithm {
		case ec.ECDSA:
			// Take P-256 as a special case
			if t.Params().Name == elliptic.P256().Params().Name {
				return ec.EncodePublicKey(t.PublicKey, false)
			}
			buf.WriteByte(byte(0x12))
		case ec.SM2:
			buf.WriteByte(byte(0x13))
		}
		label, err := GetCurveLabel(t.Curve.Params().Name)
		if err != nil {
			panic(err)
		}
		buf.WriteByte(label)
		buf.Write(ec.EncodePublicKey(t.PublicKey, false))
	case ed25519.PublicKey:
		panic("err")
	default:
		panic("err")
	}
	return buf.Bytes()
}

func getethnocompresskey(key keypair.PublicKey) []byte {
	var buf bytes.Buffer
	switch t := key.(type) {
	case *ec.PublicKey:
		return crypto1.FromECDSAPub(t.PublicKey)
	case ed25519.PublicKey:
		panic("err")
	default:
		panic("err")
	}
	return buf.Bytes()
}

func GetCurveLabel(name string) (byte, error) {
	switch strings.ToUpper(name) {
	case strings.ToUpper(elliptic.P224().Params().Name):
		return 1, nil
	case strings.ToUpper(elliptic.P256().Params().Name):
		return 2, nil
	case strings.ToUpper(elliptic.P384().Params().Name):
		return 3, nil
	case strings.ToUpper(elliptic.P521().Params().Name):
		return 4, nil
	case strings.ToUpper(sm2.SM2P256V1().Params().Name):
		return 20, nil
	case strings.ToUpper(btcec.S256().Name):
		return 5, nil
	default:
		panic("err")
	}
}
