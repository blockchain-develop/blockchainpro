package bitcoin

import (
	"bytes"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"github.com/blockchainpro/usage/utils"
	"github.com/btcsuite/btcd/btcjson"
	"github.com/btcsuite/btcd/wire"
)

type BTCTools struct {
	restclient *utils.RestClient
}

type Request struct {
	Jsonrpc string        `json:"jsonrpc"`
	Method  string        `json:"method"`
	Params  []interface{} `json:"params"`
	Id      int           `json:"id"`
}

type Response struct {
	Result json.RawMessage   `json:"result"`
	Error  *btcjson.RPCError `json:"error"` //maybe wrong
	Id     int               `json:"id"`
}

type BlockHeader struct {
	Hash   string `json:"hash"`
	Time   uint32 `json:"time"`
	Height uint32 `json:"height"`
}

func NewBtcTools(url string, user string, password string) *BTCTools {
	restclient := utils.NewRestClient(url)
	restclient.SetAuth(user, password)
	tool := &BTCTools{
		restclient: restclient,
	}
	return tool
}

func (self *BTCTools) GetCurrentHeight() (uint32, error) {
	req := Request{
		Jsonrpc: "1.0",
		Method:  "getblockcount",
		Params:  nil,
		Id:      1,
	}
	reqdata, err := json.Marshal(req)
	if err != nil {
		return 0, fmt.Errorf("GetCurrentHeight - failed to marshal request: %v", err)
	}

	repdata, err := self.restclient.SendRestRequestWithAuth(reqdata)
	if err != nil {
		return 0, fmt.Errorf("GetCurrentHeight - send request failed: %v", err)
	}

	rep := &Response{}
	err = json.Unmarshal(repdata, &rep)
	if err != nil {
		return 0, fmt.Errorf("GetCurrentHeight - failed to unmarshal response: %v", err)
	}
	if rep.Error != nil {
		return 0, fmt.Errorf("GetCurrentHeight - response shows failure: %v", rep.Error.Message)
	}
	var blockCount uint32
	err = json.Unmarshal(rep.Result, &blockCount)
	if err != nil {
		return 0, fmt.Errorf("GetCurrentHeight - failed to parse height: %v", err)
	}
	return blockCount, nil
}

func (self *BTCTools) GetBlockHash(height uint32) (string, error) {
	req := Request{
		Jsonrpc: "1.0",
		Method:  "getblockhash",
		Params:  []interface{}{height},
		Id:      1,
	}
	reqdata, err := json.Marshal(req)
	if err != nil {
		return "", fmt.Errorf("GetBlockHash - failed to marshal request: %v", err)
	}

	repdata, err := self.restclient.SendRestRequestWithAuth(reqdata)
	if err != nil {
		return "", fmt.Errorf("GetBlockHash - send request failed: %v", err)
	}

	rep := &Response{}
	err = json.Unmarshal(repdata, &rep)
	if err != nil {
		return "", fmt.Errorf("GetBlockHash - failed to unmarshal response: %v", err)
	}

	if rep.Error != nil {
		return "", fmt.Errorf("GetBlockHash - response shows failure: %v", rep.Error.Message)
	}
	var hash string
	err = json.Unmarshal(rep.Result, &hash)
	if err != nil {
		return "", fmt.Errorf("GetCurrentHeight - failed to parse height: %v", err)
	}
	return hash, nil
}

func (self *BTCTools) GetBlockHeader(hash string) (*wire.BlockHeader, error) {
	req := Request{
		Jsonrpc: "1.0",
		Method:  "getblockheader",
		Params:  []interface{}{hash},
		Id:      1,
	}
	reqdata, err := json.Marshal(req)
	if err != nil {
		return nil, fmt.Errorf("GetBlockHeader - failed to marshal request: %v", err)
	}

	repdata, err := self.restclient.SendRestRequestWithAuth(reqdata)
	if err != nil {
		return nil, fmt.Errorf("GetBlockHeader - send request failed: %v", err)
	}

	rep := &Response{}
	err = json.Unmarshal(repdata, &rep)
	if err != nil {
		return nil, fmt.Errorf("GetBlockHeader - failed to unmarshal response: %v", err)
	}

	if rep.Error != nil {
		return nil, fmt.Errorf("GetBlockHeader - response shows failure: %v", rep.Error.Message)
	}

	blockheader := wire.BlockHeader{}
	err = blockheader.BtcDecode(bytes.NewBuffer(rep.Result), wire.ProtocolVersion, wire.LatestEncoding)
	//err = json.Unmarshal(rep.Result, &blockheader)
	if err != nil {
		return nil, fmt.Errorf("GetCurrentHeight - failed to parse height: %v", err)
	}
	return &blockheader, nil
}

func (self *BTCTools) GetTx(hash string) (*btcjson.TxRawResult, error) {
	req := Request{
		Jsonrpc: "1.0",
		Method:  "getrawtransaction",
		Params:  []interface{}{hash, true},
		Id:      1,
	}
	reqdata, err := json.Marshal(req)
	if err != nil {
		return nil, fmt.Errorf("GetTx - failed to marshal request: %v", err)
	}

	repdata, err := self.restclient.SendRestRequestWithAuth(reqdata)
	if err != nil {
		return nil, fmt.Errorf("GetTx - send request failed: %v", err)
	}

	rep := &Response{}
	err = json.Unmarshal(repdata, &rep)
	if err != nil {
		return nil, fmt.Errorf("GetTx - failed to unmarshal response: %v", err)
	}
	if rep.Error != nil {
		return nil, fmt.Errorf("GetTx - response shows failure: %v", rep.Error.Message)
	}

	txRawResult := &btcjson.TxRawResult{}
	err = json.Unmarshal(rep.Result, txRawResult)
	if err != nil {
		return nil, fmt.Errorf("GetTx - Unmarshal Result: %v", err)
	}
	return txRawResult, nil
}

func (self *BTCTools) GetTxOut(hash string, vout uint32) (*btcjson.GetTxOutResult, error) {
	req := Request{
		Jsonrpc: "1.0",
		Method:  "gettxout",
		Params:  []interface{}{hash, vout, false},
		Id:      1,
	}
	reqdata, err := json.Marshal(req)
	if err != nil {
		return nil, fmt.Errorf("GetTxOut - failed to marshal request: %v", err)
	}
	repdata, err := self.restclient.SendRestRequestWithAuth(reqdata)
	if err != nil {
		return nil, fmt.Errorf("GetTxOut - send request failed: %v", err)
	}

	rep := &Response{}
	err = json.Unmarshal(repdata, &rep)
	if err != nil {
		return nil, fmt.Errorf("GetTxOut - failed to unmarshal response: %v", err)
	}

	txOutResult := &btcjson.GetTxOutResult{}
	err = json.Unmarshal(rep.Result, txOutResult)
	if err != nil {
		return nil, fmt.Errorf("GetTxOut - Unmarshal Result: %v", err)
	}
	return txOutResult, nil
}

//func (self *BTCTools) GetTxInfo(hash string) (*btcjson.GetTransactionResult, error) {
//	req := Request{
//		Jsonrpc: "1.0",
//		Method:  "gettransaction",
//		Params:  []interface{}{hash},
//		Id:      1,
//	}
//	reqdata, err := json.Marshal(req)
//	if err != nil {
//		return nil, fmt.Errorf("GetTx - failed to marshal request: %v", err)
//	}
//
//	repdata, err := self.restclient.SendRestRequestWithAuth(reqdata)
//	if err != nil {
//		return nil, fmt.Errorf("GetTx - send request failed: %v", err)
//	}
//
//	rep := &Response{}
//	err = json.Unmarshal(repdata, &rep)
//	if err != nil {
//		return nil, fmt.Errorf("GetTx - failed to unmarshal response: %v", err)
//	}
//	if rep.Error != nil {
//		return nil, fmt.Errorf("GetTx - response shows failure: %v", rep.Error.Message)
//	}
//
//	txInfo := &btcjson.GetTransactionResult{}
//	json.Unmarshal(rep.Result, txInfo)
//	return txInfo, nil
//}

func (self *BTCTools) GetBlock(hash string) (*btcjson.BlockDetails, error) {
	req := Request{
		Jsonrpc: "1.0",
		Method:  "getblock",
		Params:  []interface{}{hash, true},
		Id:      1,
	}
	reqdata, err := json.Marshal(req)
	if err != nil {
		return nil, fmt.Errorf("GetBlock - failed to marshal request: %v", err)
	}

	repdata, err := self.restclient.SendRestRequestWithAuth(reqdata)
	if err != nil {
		return nil, fmt.Errorf("GetBlock - send request failed: %v", err)
	}

	rep := &Response{}
	err = json.Unmarshal(repdata, &rep)
	if err != nil {
		return nil, fmt.Errorf("GetBlock - failed to unmarshal response: %v", err)
	}
	if rep.Error != nil {
		return nil, fmt.Errorf("GetBlock - response shows failure: %v", rep.Error.Message)
	}

	block := &btcjson.BlockDetails{}
	json.Unmarshal(rep.Result, block)
	return block, nil
}

func (self *BTCTools) GetBlockHeaderByHeight(height uint32) (*wire.BlockHeader, error) {
	hash, err := self.GetBlockHash(height)
	if err != nil {
		return nil, err
	}
	header, err := self.GetBlockHeader(hash)
	if err != nil {
		return nil, err
	}
	return header, nil
}

func (self *BTCTools) PaserRawTx(rawTx []byte) (*wire.MsgTx, string, error) {
	tx := wire.MsgTx{}
	err := tx.BtcDecode(bytes.NewBuffer(rawTx), wire.ProtocolVersion, wire.LatestEncoding)
	if err != nil {
		return nil, "", fmt.Errorf("PaserRawTx - failed: %v", err)
	}
	return &tx, tx.TxHash().String(), nil
}

func (self *BTCTools) TestRawTx() {
	rawStr := "01000000015a920177b4840c020857f841f5c44f25d7d42b79740fbcab6f167bb8af94ffba000000006a47304402202b81bee8c38bacdbceecd54f633a06ea36d49b454fd967945a2a076ad53dfcbb022061a0d407e92a8d89d9729559568657a26adeb320c649a2b9ae71c4fd4d8f0b19012103128a2c4525179e47f38cf3fefca37a61548ca4610255b3fb4ee86de2d3e80c0fffffffff02583e0f000000000017a91487a9652e9b396545598c0fc72cb5a98848bf93d38700000000000000003d6a3b660200000000000000000000000000000014680170e3b0e05dc5faed76d85ad2d4f67b6399fc14f3b8a17f1f957f60c88f105e32ebff3f022e56a400000000"
	rawTx, _ := hex.DecodeString(rawStr)
	tx := wire.MsgTx{}
	err := tx.BtcDecode(bytes.NewBuffer(rawTx), wire.ProtocolVersion, wire.LatestEncoding)
	if err != nil {
		fmt.Printf("paser failed!")
	} else {
		fmt.Printf("paser successful!")
	}
}
