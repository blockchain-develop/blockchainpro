package ethereum

import (
	"encoding/hex"
	"encoding/json"
	"fmt"
	"github.com/blockchainpro/usage/utils"
	"github.com/ethereum/go-ethereum/core/types"
	"math/big"
	"strconv"
)

type jsonError struct {
	Code    int              `json:"code"`
	Message string           `json:"message"`
	Data    interface{}     `json:"data,omitempty"`
}

type heightReq struct {
	JsonRpc string          `json:"jsonrpc"`
	Method  string          `json:"method"`
	Params  []string        `json:"params"`
	Id      uint            `json:"id"`
}

type heightRsp struct {
	JsonRpc string          `json:"jsonrpc"`
	Result  string          `json:"result,omitempty"`
	Error   *jsonError      `json:"error,omitempty"`
	Id      uint            `json:"id"`
}

type proofReq struct {
	JsonRPC string          `json:"jsonrpc"`
	Method  string          `json:"method"`
	Params  []interface{}  `json:"params"`
	Id      uint            `json:"id"`
}

type proofRsp struct {
	JsonRPC string           `json:"jsonrpc"`
	Result  ETHProof         `json:"result,omitempty"`
	Error   *jsonError       `json:"error,omitempty"`
	Id      uint             `json:"id"`
}

type ETHProof struct {
	Address       string         `json:"address"`
	Balance       string         `json:"balance"`
	CodeHash      string         `json:"codeHash"`
	Nonce         string         `json:"nonce"`
	StorageHash   string         `json:"storageHash"`
	AccountProof  []string       `json:"accountProof"`
	StorageProofs []StorageProof `json:"storageProof"`
}

type StorageProof struct {
	Key   string       `json:"key"`
	Value string       `json:"value"`
	Proof []string     `json:"proof"`
}

type blockReq struct {
	JsonRpc string               `json:"jsonrpc"`
	Method  string               `json:"method"`
	Params  []interface{}       `json:"params"`
	Id      uint                 `json:"id"`
}

type blockRsp struct {
	JsonRPC string               `json:"jsonrpc"`
	Result  *types.Header        `json:"result,omitempty"`
	Error   *jsonError           `json:"error,omitempty"`
	Id      uint                 `json:"id"`
}


type getBalanceReq struct {
	JsonRpc string          `json:"jsonrpc"`
	Method  string          `json:"method"`
	Params  []string        `json:"params"`
	Id      uint            `json:"id"`
}

type getBalanceRsp struct {
	JsonRpc string          `json:"jsonrpc"`
	Result  string          `json:"result,omitempty"`
	Error   *jsonError      `json:"error,omitempty"`
	Id      uint            `json:"id"`
}

type getTransactionReq struct {
	JsonRpc string          `json:"jsonrpc"`
	Method  string          `json:"method"`
	Params  []string        `json:"params"`
	Id      uint            `json:"id"`
}

type getTransactionRsp struct {
	JsonRpc string          `json:"jsonrpc"`
	Result  string          `json:"result,omitempty"`
	Error   *jsonError      `json:"error,omitempty"`
	Id      uint            `json:"id"`
}

type EthClient struct {
	client     *utils.RestClient
}

func NewEthClient(url string) *EthClient {
	ethclient := &EthClient{
		client: utils.NewRestClient(url),
	}
	return ethclient
}

func (client *EthClient) GetNodeHeader(height uint64) (*types.Header, error) {
	params := []interface{} {fmt.Sprintf("0x%x", height), true}
	req := &blockReq{
		JsonRpc: "2.0",
		Method:  "eth_getBlockByNumber",
		Params:  params,
		Id:      1,
	}
	reqdata, err := json.Marshal(req)
	if err != nil {
		return nil, fmt.Errorf("GetNodeHeight: marshal req err: %s", err)
	}
	rspdata, err := client.client.SendRestRequest(reqdata)
	if err != nil {
		return nil, fmt.Errorf("GetNodeHeight err: %s", err)
	}
	rsp := &blockRsp{}
	err = json.Unmarshal(rspdata, rsp)
	if err != nil {
		return nil, fmt.Errorf("GetNodeHeight, unmarshal resp err: %s", err)
	}
	if rsp.Error != nil {
		return nil, fmt.Errorf("GetNodeHeight, unmarshal resp err: %s", rsp.Error.Message)
	}
	return rsp.Result, nil
}

func (client *EthClient) GetNodeHeight() (uint64, error) {
	req := &heightReq{
		JsonRpc: "2.0",
		Method:  "eth_blockNumber",
		Params:  make([]string, 0),
		Id:      1,
	}
	reqData, err := json.Marshal(req)
	if err != nil {
		return 0, fmt.Errorf("GetNodeHeight: marshal req err: %s", err)
	}
	rspData, err := client.client.SendRestRequest(reqData)
	if err != nil {
		return 0, fmt.Errorf("GetNodeHeight err: %s", err)
	}
	rsp := &heightRsp{}
	err = json.Unmarshal(rspData, rsp)
	if err != nil {
		return 0, fmt.Errorf("GetNodeHeight, unmarshal resp err: %s", err)
	}
	if rsp.Error != nil {
		return 0, fmt.Errorf("GetNodeHeight, unmarshal resp err: %s", rsp.Error.Message)
	}
	height, err := strconv.ParseUint(rsp.Result, 0, 64)
	if err != nil {
		return 0, fmt.Errorf("GetNodeHeight, parse resp height %s failed", rsp.Result)
	} else {
		return height, nil
	}
}

func (client *EthClient) GetProof(contractAddress string, key string, blockheight string) ([]byte, error) {
	req := &proofReq{
		JsonRPC: "2.0",
		Method:  "eth_getProof",
		Params:  []interface{}{contractAddress, []string{key}, blockheight},
		Id:      1,
	}
	reqdata, err := json.Marshal(req)
	if err != nil {
		return nil, fmt.Errorf("get_ethproof: marshal req err: %s", err)
	}
	fmt.Printf("proof req is:%s\n", string(reqdata))
	rspdata, err := client.client.SendRestRequest(reqdata)
	if err != nil {
		return nil, fmt.Errorf("GetProof: send request err: %s", err)
	}
	rsp := &proofRsp{}
	err = json.Unmarshal(rspdata, rsp)
	if err != nil {
		return nil, fmt.Errorf("GetProof, unmarshal resp err: %s", err)
	}
	if rsp.Error != nil {
		return nil, fmt.Errorf("GetNodeHeight, unmarshal resp err: %s", rsp.Error.Message)
	}
	result, err := json.Marshal(rsp.Result)
	if err != nil {
		return nil, fmt.Errorf("GetProof, Marshal result err: %s", err)
	}
	//fmt.Printf("proof res is:%s\n", string(result))
	return result, nil
}

func (client *EthClient) GetBalance(height uint64) (uint64, error) {
	params := []string {"0xc94770007dda54cF92009BFF0dE90c06F603a09f", "latest"}
	req := &getBalanceReq{
		JsonRpc: "2.0",
		Method:  "eth_getBalance",
		Params:  params,
		Id:      1,
	}
	reqdata, err := json.Marshal(req)
	if err != nil {
		return 0, fmt.Errorf("GetNodeHeight: marshal req err: %s", err)
	}
	rspdata, err := client.client.SendRestRequest(reqdata)
	if err != nil {
		return 0, fmt.Errorf("GetNodeHeight err: %s", err)
	}
	rsp := &getBalanceRsp{}
	err = json.Unmarshal(rspdata, rsp)
	if err != nil {
		return 0, fmt.Errorf("GetNodeHeight, unmarshal resp err: %s", err)
	}
	if rsp.Error != nil {
		return 0, fmt.Errorf("GetNodeHeight, unmarshal resp err: %s", rsp.Error.Message)
	}
	balance, err := strconv.ParseUint(rsp.Result, 0, 64)
	if err != nil {
		return 0, fmt.Errorf("GetNodeHeight, parse resp height %s failed", rsp.Result)
	} else {
		return balance, nil
	}
}

func (client *EthClient) GetTransaction(hash string) (uint64, error) {
	params := []string {"0xfd40b4f7ebf6bebe5f7cab25ed2ff4938c6ba68d055fecaff78470aaf170f09e"}
	req := &getTransactionReq{
		JsonRpc: "2.0",
		Method:  "eth_getTransactionByHash",
		Params:  params,
		Id:      1,
	}
	reqdata, err := json.Marshal(req)
	if err != nil {
		return 0, fmt.Errorf("GetNodeHeight: marshal req err: %s", err)
	}
	rspdata, err := client.client.SendRestRequest(reqdata)
	if err != nil {
		return 0, fmt.Errorf("GetNodeHeight err: %s", err)
	}
	rsp := &getTransactionRsp{}
	err = json.Unmarshal(rspdata, rsp)
	if err != nil {
		return 0, fmt.Errorf("GetNodeHeight, unmarshal resp err: %s", err)
	}
	if rsp.Error != nil {
		return 0, fmt.Errorf("GetNodeHeight, unmarshal resp err: %s", rsp.Error.Message)
	}
	balance, err := strconv.ParseUint(rsp.Result, 0, 64)
	if err != nil {
		return 0, fmt.Errorf("GetNodeHeight, parse resp height %s failed", rsp.Result)
	} else {
		return balance, nil
	}
}

func EncodeBigInt(b *big.Int) string {
	if b.Uint64() == 0 {
		return "00"
	}
	return hex.EncodeToString(b.Bytes())
}


func (client *EthClient) GetNodeHeaderByHash(hash string) (*types.Header, error) {
	params := []interface{} {fmt.Sprintf(hash), true}
	req := &blockReq{
		JsonRpc: "2.0",
		Method:  "eth_getBlockByHash",
		Params:  params,
		Id:      1,
	}
	reqdata, err := json.Marshal(req)
	if err != nil {
		return nil, fmt.Errorf("GetNodeHeight: marshal req err: %s", err)
	}
	rspdata, err := client.client.SendRestRequest(reqdata)
	if err != nil {
		return nil, fmt.Errorf("GetNodeHeight err: %s", err)
	}
	rsp := &blockRsp{}
	err = json.Unmarshal(rspdata, rsp)
	if err != nil {
		return nil, fmt.Errorf("GetNodeHeight, unmarshal resp err: %s", err)
	}
	if rsp.Error != nil {
		return nil, fmt.Errorf("GetNodeHeight, unmarshal resp err: %s", rsp.Error.Message)
	}
	return rsp.Result, nil
}
