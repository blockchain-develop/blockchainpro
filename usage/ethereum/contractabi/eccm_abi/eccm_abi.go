// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package eccm_abi

import (
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = abi.U256
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// ContextABI is the input ABI used to generate the binding from.
const ContextABI = "[{\"inputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"}]"

// Context is an auto generated Go binding around an Ethereum contract.
type Context struct {
	ContextCaller     // Read-only binding to the contract
	ContextTransactor // Write-only binding to the contract
	ContextFilterer   // Log filterer for contract events
}

// ContextCaller is an auto generated read-only Go binding around an Ethereum contract.
type ContextCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContextTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ContextTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContextFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ContextFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContextSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ContextSession struct {
	Contract     *Context          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ContextCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ContextCallerSession struct {
	Contract *ContextCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// ContextTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ContextTransactorSession struct {
	Contract     *ContextTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// ContextRaw is an auto generated low-level Go binding around an Ethereum contract.
type ContextRaw struct {
	Contract *Context // Generic contract binding to access the raw methods on
}

// ContextCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ContextCallerRaw struct {
	Contract *ContextCaller // Generic read-only contract binding to access the raw methods on
}

// ContextTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ContextTransactorRaw struct {
	Contract *ContextTransactor // Generic write-only contract binding to access the raw methods on
}

// NewContext creates a new instance of Context, bound to a specific deployed contract.
func NewContext(address common.Address, backend bind.ContractBackend) (*Context, error) {
	contract, err := bindContext(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Context{ContextCaller: ContextCaller{contract: contract}, ContextTransactor: ContextTransactor{contract: contract}, ContextFilterer: ContextFilterer{contract: contract}}, nil
}

// NewContextCaller creates a new read-only instance of Context, bound to a specific deployed contract.
func NewContextCaller(address common.Address, caller bind.ContractCaller) (*ContextCaller, error) {
	contract, err := bindContext(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ContextCaller{contract: contract}, nil
}

// NewContextTransactor creates a new write-only instance of Context, bound to a specific deployed contract.
func NewContextTransactor(address common.Address, transactor bind.ContractTransactor) (*ContextTransactor, error) {
	contract, err := bindContext(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ContextTransactor{contract: contract}, nil
}

// NewContextFilterer creates a new log filterer instance of Context, bound to a specific deployed contract.
func NewContextFilterer(address common.Address, filterer bind.ContractFilterer) (*ContextFilterer, error) {
	contract, err := bindContext(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ContextFilterer{contract: contract}, nil
}

// bindContext binds a generic wrapper to an already deployed contract.
func bindContext(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ContextABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Context *ContextRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Context.Contract.ContextCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Context *ContextRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Context.Contract.ContextTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Context *ContextRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Context.Contract.ContextTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Context *ContextCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Context.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Context *ContextTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Context.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Context *ContextTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Context.Contract.contract.Transact(opts, method, params...)
}

// ECCUtilsABI is the input ABI used to generate the binding from.
const ECCUtilsABI = "[]"

// ECCUtilsBin is the compiled bytecode used for deploying new contracts.
var ECCUtilsBin = "0x60556023600b82828239805160001a607314601657fe5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea265627a7a723158209f1b7746653c1995d82b37b72ee62ad02a890eb13138fc2037065e455b0d098764736f6c634300050f0032"

// DeployECCUtils deploys a new Ethereum contract, binding an instance of ECCUtils to it.
func DeployECCUtils(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *ECCUtils, error) {
	parsed, err := abi.JSON(strings.NewReader(ECCUtilsABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(ECCUtilsBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ECCUtils{ECCUtilsCaller: ECCUtilsCaller{contract: contract}, ECCUtilsTransactor: ECCUtilsTransactor{contract: contract}, ECCUtilsFilterer: ECCUtilsFilterer{contract: contract}}, nil
}

// ECCUtils is an auto generated Go binding around an Ethereum contract.
type ECCUtils struct {
	ECCUtilsCaller     // Read-only binding to the contract
	ECCUtilsTransactor // Write-only binding to the contract
	ECCUtilsFilterer   // Log filterer for contract events
}

// ECCUtilsCaller is an auto generated read-only Go binding around an Ethereum contract.
type ECCUtilsCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ECCUtilsTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ECCUtilsTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ECCUtilsFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ECCUtilsFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ECCUtilsSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ECCUtilsSession struct {
	Contract     *ECCUtils         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ECCUtilsCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ECCUtilsCallerSession struct {
	Contract *ECCUtilsCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// ECCUtilsTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ECCUtilsTransactorSession struct {
	Contract     *ECCUtilsTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// ECCUtilsRaw is an auto generated low-level Go binding around an Ethereum contract.
type ECCUtilsRaw struct {
	Contract *ECCUtils // Generic contract binding to access the raw methods on
}

// ECCUtilsCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ECCUtilsCallerRaw struct {
	Contract *ECCUtilsCaller // Generic read-only contract binding to access the raw methods on
}

// ECCUtilsTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ECCUtilsTransactorRaw struct {
	Contract *ECCUtilsTransactor // Generic write-only contract binding to access the raw methods on
}

// NewECCUtils creates a new instance of ECCUtils, bound to a specific deployed contract.
func NewECCUtils(address common.Address, backend bind.ContractBackend) (*ECCUtils, error) {
	contract, err := bindECCUtils(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ECCUtils{ECCUtilsCaller: ECCUtilsCaller{contract: contract}, ECCUtilsTransactor: ECCUtilsTransactor{contract: contract}, ECCUtilsFilterer: ECCUtilsFilterer{contract: contract}}, nil
}

// NewECCUtilsCaller creates a new read-only instance of ECCUtils, bound to a specific deployed contract.
func NewECCUtilsCaller(address common.Address, caller bind.ContractCaller) (*ECCUtilsCaller, error) {
	contract, err := bindECCUtils(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ECCUtilsCaller{contract: contract}, nil
}

// NewECCUtilsTransactor creates a new write-only instance of ECCUtils, bound to a specific deployed contract.
func NewECCUtilsTransactor(address common.Address, transactor bind.ContractTransactor) (*ECCUtilsTransactor, error) {
	contract, err := bindECCUtils(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ECCUtilsTransactor{contract: contract}, nil
}

// NewECCUtilsFilterer creates a new log filterer instance of ECCUtils, bound to a specific deployed contract.
func NewECCUtilsFilterer(address common.Address, filterer bind.ContractFilterer) (*ECCUtilsFilterer, error) {
	contract, err := bindECCUtils(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ECCUtilsFilterer{contract: contract}, nil
}

// bindECCUtils binds a generic wrapper to an already deployed contract.
func bindECCUtils(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ECCUtilsABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ECCUtils *ECCUtilsRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ECCUtils.Contract.ECCUtilsCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ECCUtils *ECCUtilsRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ECCUtils.Contract.ECCUtilsTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ECCUtils *ECCUtilsRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ECCUtils.Contract.ECCUtilsTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ECCUtils *ECCUtilsCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ECCUtils.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ECCUtils *ECCUtilsTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ECCUtils.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ECCUtils *ECCUtilsTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ECCUtils.Contract.contract.Transact(opts, method, params...)
}

// EthCrossChainManagerABI is the input ABI used to generate the binding from.
const EthCrossChainManagerABI = "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_eccd\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"height\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"rawHeader\",\"type\":\"bytes\"}],\"name\":\"ChangeBookKeeperEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"txId\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"proxyOrAssetContract\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"toChainId\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"toContract\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"rawdata\",\"type\":\"bytes\"}],\"name\":\"CrossChainEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"height\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"rawHeader\",\"type\":\"bytes\"}],\"name\":\"InitGenesisBlockEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Paused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"height\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"rawHeader\",\"type\":\"bytes\"}],\"name\":\"SyncBlockHeaderEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Unpaused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"fromChainID\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"toContract\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"crossChainTxHash\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"fromChainTxHash\",\"type\":\"bytes\"}],\"name\":\"VerifyAndExecuteTxEvent\",\"type\":\"event\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"_rawHeader\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"_pubKeyList\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"_sigList\",\"type\":\"bytes\"}],\"name\":\"ChangeBookKeeper\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"EthCrossChainDataAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"_rawHeader\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"_pubKeyList\",\"type\":\"bytes\"}],\"name\":\"InitGenesisBlock\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"rawHeader\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"sigList\",\"type\":\"bytes\"},{\"internalType\":\"bytes32[]\",\"name\":\"proof\",\"type\":\"bytes32[]\"},{\"internalType\":\"enumECCUtils.POS[]\",\"name\":\"position\",\"type\":\"uint8[]\"},{\"internalType\":\"bytes\",\"name\":\"toMerkleValueBs\",\"type\":\"bytes\"},{\"internalType\":\"uint64\",\"name\":\"blockHeight\",\"type\":\"uint64\"}],\"name\":\"SyncAndVerify\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"_rawHeader\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"_sigList\",\"type\":\"bytes\"}],\"name\":\"SyncBlockHeader\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"_toChainId\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"_toContract\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"_method\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"_txData\",\"type\":\"bytes\"}],\"name\":\"crossChain\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"isOwner\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"pause\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"unpause\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"newEthCrossChainManagerAddress\",\"type\":\"address\"}],\"name\":\"upgradeToNew\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"bytes32[]\",\"name\":\"_proof\",\"type\":\"bytes32[]\"},{\"internalType\":\"enumECCUtils.POS[]\",\"name\":\"_position\",\"type\":\"uint8[]\"},{\"internalType\":\"bytes\",\"name\":\"_toMerkleValueBs\",\"type\":\"bytes\"},{\"internalType\":\"uint64\",\"name\":\"_blockHeight\",\"type\":\"uint64\"}],\"name\":\"verifyAndExecuteTx\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// EthCrossChainManagerFuncSigs maps the 4-byte function signature to its string representation.
var EthCrossChainManagerFuncSigs = map[string]string{
	"21f966aa": "ChangeBookKeeper(bytes,bytes,bytes)",
	"00ba1694": "EthCrossChainDataAddress()",
	"adbccaf5": "InitGenesisBlock(bytes,bytes)",
	"48c2f675": "SyncAndVerify(bytes,bytes,bytes32[],uint8[],bytes,uint64)",
	"1e568e18": "SyncBlockHeader(bytes,bytes)",
	"bd5cf625": "crossChain(uint64,bytes,bytes,bytes)",
	"8f32d59b": "isOwner()",
	"8da5cb5b": "owner()",
	"8456cb59": "pause()",
	"5c975abb": "paused()",
	"715018a6": "renounceOwnership()",
	"f2fde38b": "transferOwnership(address)",
	"3f4ba83a": "unpause()",
	"7e724ff3": "upgradeToNew(address)",
	"d5ea92e3": "verifyAndExecuteTx(bytes32[],uint8[],bytes,uint64)",
}

// EthCrossChainManagerBin is the compiled bytecode used for deploying new contracts.
var EthCrossChainManagerBin = "0x60806040523480156200001157600080fd5b5060405162005a4a38038062005a4a833981810160405260208110156200003757600080fd5b50518060006200004f6001600160e01b03620000cd16565b600080546001600160a01b0319166001600160a01b0383169081178255604051929350917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0908290a3506000805460ff60a01b19169055600180546001600160a01b0319166001600160a01b039290921691909117905550620000d1565b3390565b61596980620000e16000396000f3fe608060405234801561001057600080fd5b50600436106100f45760003560e01c80637e724ff311610097578063adbccaf511610066578063adbccaf51461071d578063bd5cf62514610846578063d5ea92e314610964578063f2fde38b14610b17576100f4565b80637e724ff3146106df5780638456cb59146107055780638da5cb5b1461070d5780638f32d59b14610715576100f4565b80633f4ba83a116100d35780633f4ba83a1461040857806348c2f675146104105780635c975abb146106cd578063715018a6146106d5576100f4565b8062ba1694146100f95780631e568e181461011d57806321f966aa1461025a575b600080fd5b610101610b3d565b604080516001600160a01b039092168252519081900360200190f35b6102466004803603604081101561013357600080fd5b810190602081018135600160201b81111561014d57600080fd5b82018360208201111561015f57600080fd5b803590602001918460018302840111600160201b8311171561018057600080fd5b91908080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152509295949360208101935035915050600160201b8111156101d257600080fd5b8201836020820111156101e457600080fd5b803590602001918460018302840111600160201b8311171561020557600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250929550610b4c945050505050565b604080519115158252519081900360200190f35b6102466004803603606081101561027057600080fd5b810190602081018135600160201b81111561028a57600080fd5b82018360208201111561029c57600080fd5b803590602001918460018302840111600160201b831117156102bd57600080fd5b91908080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152509295949360208101935035915050600160201b81111561030f57600080fd5b82018360208201111561032157600080fd5b803590602001918460018302840111600160201b8311171561034257600080fd5b91908080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152509295949360208101935035915050600160201b81111561039457600080fd5b8201836020820111156103a657600080fd5b803590602001918460018302840111600160201b831117156103c757600080fd5b91908080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152509295506114b2945050505050565b610246611e37565b610246600480360360c081101561042657600080fd5b810190602081018135600160201b81111561044057600080fd5b82018360208201111561045257600080fd5b803590602001918460018302840111600160201b8311171561047357600080fd5b91908080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152509295949360208101935035915050600160201b8111156104c557600080fd5b8201836020820111156104d757600080fd5b803590602001918460018302840111600160201b831117156104f857600080fd5b91908080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152509295949360208101935035915050600160201b81111561054a57600080fd5b82018360208201111561055c57600080fd5b803590602001918460208302840111600160201b8311171561057d57600080fd5b9190808060200260200160405190810160405280939291908181526020018383602002808284376000920191909152509295949360208101935035915050600160201b8111156105cc57600080fd5b8201836020820111156105de57600080fd5b803590602001918460208302840111600160201b831117156105ff57600080fd5b9190808060200260200160405190810160405280939291908181526020018383602002808284376000920191909152509295949360208101935035915050600160201b81111561064e57600080fd5b82018360208201111561066057600080fd5b803590602001918460018302840111600160201b8311171561068157600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250929550505090356001600160401b03169150611fb69050565b6102466120ce565b6106dd6120de565b005b610246600480360360208110156106f557600080fd5b50356001600160a01b031661216f565b610246612280565b6101016123f5565b610246612404565b6102466004803603604081101561073357600080fd5b810190602081018135600160201b81111561074d57600080fd5b82018360208201111561075f57600080fd5b803590602001918460018302840111600160201b8311171561078057600080fd5b91908080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152509295949360208101935035915050600160201b8111156107d257600080fd5b8201836020820111156107e457600080fd5b803590602001918460018302840111600160201b8311171561080557600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250929550612428945050505050565b6102466004803603608081101561085c57600080fd5b6001600160401b038235169190810190604081016020820135600160201b81111561088657600080fd5b82018360208201111561089857600080fd5b803590602001918460018302840111600160201b831117156108b957600080fd5b919390929091602081019035600160201b8111156108d657600080fd5b8201836020820111156108e857600080fd5b803590602001918460018302840111600160201b8311171561090957600080fd5b919390929091602081019035600160201b81111561092657600080fd5b82018360208201111561093857600080fd5b803590602001918460018302840111600160201b8311171561095957600080fd5b509092509050612b9c565b6102466004803603608081101561097a57600080fd5b810190602081018135600160201b81111561099457600080fd5b8201836020820111156109a657600080fd5b803590602001918460208302840111600160201b831117156109c757600080fd5b9190808060200260200160405190810160405280939291908181526020018383602002808284376000920191909152509295949360208101935035915050600160201b811115610a1657600080fd5b820183602082011115610a2857600080fd5b803590602001918460208302840111600160201b83111715610a4957600080fd5b9190808060200260200160405190810160405280939291908181526020018383602002808284376000920191909152509295949360208101935035915050600160201b811115610a9857600080fd5b820183602082011115610aaa57600080fd5b803590602001918460018302840111600160201b83111715610acb57600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250929550505090356001600160401b031691506132ee9050565b6106dd60048036036020811015610b2d57600080fd5b50356001600160a01b031661396f565b6001546001600160a01b031681565b60008054600160a01b900460ff1615610b9f576040805162461bcd60e51b815260206004820152601060248201526f14185d5cd8589b194e881c185d5cd95960821b604482015290519081900360640190fd5b610ba7615524565b610bb0846139c2565b6101408101519091506001600160601b03191615610bff5760405162461bcd60e51b81526004018080602001828103825260298152602001806157b06029913960400191505060405180910390fd5b60008160e0015163ffffffff1611610c5e576040805162461bcd60e51b815260206004820152601760248201527f626c6f636b206865696768742073686f756c64203e2030000000000000000000604482015290519081900360640190fd5b60015460e082015160408051634ba1f4eb60e11b815263ffffffff9092166004830152516001600160a01b03909216916060918391639743e9d691602480820192600092909190829003018186803b158015610cb957600080fd5b505afa158015610ccd573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f191682016040526020811015610cf657600080fd5b8101908080516040519392919084600160201b821115610d1557600080fd5b908301906020820185811115610d2a57600080fd5b8251600160201b811182820188101715610d4357600080fd5b82525081516020918201929091019080838360005b83811015610d70578181015183820152602001610d58565b50505050905090810190601f168015610d9d5780820380516001836020036101000a031916815260200191505b50604052505081519192505015610dba57600193505050506114ac565b600080836001600160a01b03166371a777506040518163ffffffff1660e01b815260040160206040518083038186803b158015610df657600080fd5b505afa158015610e0a573d6000803e3d6000fd5b505050506040513d6020811015610e2057600080fd5b505160e08601519091506001600160401b03821663ffffffff90911610610e4957809150610fc9565b6060846001600160a01b031663f0b1fcfe6040518163ffffffff1660e01b815260040160006040518083038186803b158015610e8457600080fd5b505afa158015610e98573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f191682016040526020811015610ec157600080fd5b8101908080516040519392919084600160201b821115610ee057600080fd5b908301906020820185811115610ef557600080fd5b82518660208202830111600160201b82111715610f1157600080fd5b82525081516020918201928201910280838360005b83811015610f3e578181015183820152602001610f26565b505050509050016040525050509050600080610f668384518a60e0015163ffffffff16613adc565b9150915080610fa65760405162461bcd60e51b815260040180806020018281038252603481526020018061577c6034913960400191505060405180910390fd5b82826001600160401b031681518110610fbb57fe5b602002602001015194505050505b6060611114856001600160a01b0316637d737b8b856040518263ffffffff1660e01b815260040180826001600160401b03166001600160401b0316815260200191505060006040518083038186803b15801561102457600080fd5b505afa158015611038573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f19168201604052602081101561106157600080fd5b8101908080516040519392919084600160201b82111561108057600080fd5b90830190602082018581111561109557600080fd5b8251600160201b8111828201881017156110ae57600080fd5b82525081516020918201929091019080838360005b838110156110db5781810151838201526020016110c3565b50505050905090810190601f1680156111085780820380516001836020036101000a031916815260200191505b50604052505050613d29565b80519091506111328a8a84600360001960028702015b048503613ddc565b611183576040805162461bcd60e51b815260206004820152601f60248201527f56657269667920686561646572207369676e6174757265206661696c65642100604482015290519081900360640190fd5b856001600160a01b031663021d70ab8860e001518c6040518363ffffffff1660e01b8152600401808363ffffffff166001600160401b0316815260200180602001828103825283818151815260200191508051906020019080838360005b838110156111f95781810151838201526020016111e1565b50505050905090810190601f1680156112265780820380516001836020036101000a031916815260200191505b509350505050602060405180830381600087803b15801561124657600080fd5b505af115801561125a573d6000803e3d6000fd5b505050506040513d602081101561127057600080fd5b50516112ad5760405162461bcd60e51b815260040180806020018281038252602c815260200180615725602c913960400191505060405180910390fd5b6000866001600160a01b0316634ed1d8cc6040518163ffffffff1660e01b815260040160206040518083038186803b1580156112e857600080fd5b505afa1580156112fc573d6000803e3d6000fd5b505050506040513d602081101561131257600080fd5b505160e08901519091506001600160401b03821663ffffffff90911611156113f457866001600160a01b0316636d0444408960e001516040518263ffffffff1660e01b8152600401808263ffffffff166001600160401b03168152602001915050602060405180830381600087803b15801561138d57600080fd5b505af11580156113a1573d6000803e3d6000fd5b505050506040513d60208110156113b757600080fd5b50516113f45760405162461bcd60e51b815260040180806020018281038252602d8152602001806157f9602d913960400191505060405180910390fd5b7f1645cc3eff4fb98fef2ffd7806e7f39bfe18558f6656e69dfc9c8f26e958d4e78860e001518c604051808363ffffffff16815260200180602001828103825283818151815260200191508051906020019080838360005b8381101561146457818101518382015260200161144c565b50505050905090810190601f1680156114915780820380516001836020036101000a031916815260200191505b50935050505060405180910390a16001985050505050505050505b92915050565b60008054600160a01b900460ff1615611505576040805162461bcd60e51b815260206004820152601060248201526f14185d5cd8589b194e881c185d5cd95960821b604482015290519081900360640190fd5b61150d615524565b611516856139c2565b90506000600160009054906101000a90046001600160a01b031690506000816001600160a01b0316634ed1d8cc6040518163ffffffff1660e01b815260040160206040518083038186803b15801561156d57600080fd5b505afa158015611581573d6000803e3d6000fd5b505050506040513d602081101561159757600080fd5b505160e08401519091506001600160401b03821663ffffffff90911611611605576040805162461bcd60e51b815260206004820152601c60248201527f54686520686569676874206f662068656164657220696c6c6567616c00000000604482015290519081900360640190fd5b6101408301516001600160601b0319166116505760405162461bcd60e51b81526004018080602001828103825260258152602001806156116025913960400191505060405180910390fd5b6060826001600160a01b0316637d737b8b846001600160a01b03166371a777506040518163ffffffff1660e01b815260040160206040518083038186803b15801561169a57600080fd5b505afa1580156116ae573d6000803e3d6000fd5b505050506040513d60208110156116c457600080fd5b5051604080516001600160e01b031960e085901b1681526001600160401b039092166004830152516024808301926000929190829003018186803b15801561170b57600080fd5b505afa15801561171f573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f19168201604052602081101561174857600080fd5b8101908080516040519392919084600160201b82111561176757600080fd5b90830190602082018581111561177c57600080fd5b8251600160201b81118282018810171561179557600080fd5b82525081516020918201929091019080838360005b838110156117c25781810151838201526020016117aa565b50505050905090810190601f1680156117ef5780820380516001836020036101000a031916815260200191505b506040525050509050606061180382613d29565b805190915061181d8a89846003600019600287020161112a565b61186e576040805162461bcd60e51b815260206004820152601860248201527f566572696679207369676e6174757265206661696c6564210000000000000000604482015290519081900360640190fd5b6000606061187b8b61403f565b91509150816001600160601b0319168861014001516001600160601b031916146118e2576040805162461bcd60e51b815260206004820152601360248201527213995e1d109bdbdad95c9cc81a5b1b1959d85b606a1b604482015290519081900360640190fd5b866001600160a01b0316636d0444408960e001516040518263ffffffff1660e01b8152600401808263ffffffff166001600160401b03168152602001915050602060405180830381600087803b15801561193b57600080fd5b505af115801561194f573d6000803e3d6000fd5b505050506040513d602081101561196557600080fd5b50516119a25760405162461bcd60e51b815260040180806020018281038252602d8152602001806157f9602d913960400191505060405180910390fd5b866001600160a01b031663021d70ab8960e001518e6040518363ffffffff1660e01b8152600401808363ffffffff166001600160401b0316815260200180602001828103825283818151815260200191508051906020019080838360005b83811015611a18578181015183820152602001611a00565b50505050905090810190601f168015611a455780820380516001836020036101000a031916815260200191505b509350505050602060405180830381600087803b158015611a6557600080fd5b505af1158015611a79573d6000803e3d6000fd5b505050506040513d6020811015611a8f57600080fd5b5051611acc5760405162461bcd60e51b815260040180806020018281038252602c815260200180615725602c913960400191505060405180910390fd5b866001600160a01b031663a6c6fe7c8960e001516040518263ffffffff1660e01b8152600401808263ffffffff166001600160401b03168152602001915050602060405180830381600087803b158015611b2557600080fd5b505af1158015611b39573d6000803e3d6000fd5b505050506040513d6020811015611b4f57600080fd5b5051611b8c5760405162461bcd60e51b81526004018080602001828103825260378152602001806158546037913960400191505060405180910390fd5b866001600160a01b031663c6dbd0b78960e001516040518263ffffffff1660e01b8152600401808263ffffffff166001600160401b03168152602001915050602060405180830381600087803b158015611be557600080fd5b505af1158015611bf9573d6000803e3d6000fd5b505050506040513d6020811015611c0f57600080fd5b5051611c4c5760405162461bcd60e51b815260040180806020018281038252602c8152602001806156ce602c913960400191505060405180910390fd5b866001600160a01b031663191b0ab18960e00151611c69846140cf565b604080516001600160e01b031960e086901b16815263ffffffff8416600482019081526024820192835283516044830152835190929160640190602085019080838360005b83811015611cc6578181015183820152602001611cae565b50505050905090810190601f168015611cf35780820380516001836020036101000a031916815260200191505b509350505050602060405180830381600087803b158015611d1357600080fd5b505af1158015611d27573d6000803e3d6000fd5b505050506040513d6020811015611d3d57600080fd5b5051611d7a5760405162461bcd60e51b815260040180806020018281038252602b8152602001806156fa602b913960400191505060405180910390fd5b7fe60d33488cba3977bf65766cd2f8ac9617f64bf3b3198aff6240ce5c7d43b6908860e001518d604051808363ffffffff16815260200180602001828103825283818151815260200191508051906020019080838360005b83811015611dea578181015183820152602001611dd2565b50505050905090810190601f168015611e175780820380516001836020036101000a031916815260200191505b50935050505060405180910390a15060019b9a5050505050505050505050565b6000611e41612404565b611e80576040805162461bcd60e51b815260206004820181905260248201526000805160206157d9833981519152604482015290519081900360640190fd5b611e886120ce565b15611e9557611e956141d1565b60015460408051635c975abb60e01b815290516001600160a01b03909216918291635c975abb916004808301926020929190829003018186803b158015611edb57600080fd5b505afa158015611eef573d6000803e3d6000fd5b505050506040513d6020811015611f0557600080fd5b505115611fae57806001600160a01b0316633f4ba83a6040518163ffffffff1660e01b8152600401602060405180830381600087803b158015611f4757600080fd5b505af1158015611f5b573d6000803e3d6000fd5b505050506040513d6020811015611f7157600080fd5b5051611fae5760405162461bcd60e51b81526004018080602001828103825260298152602001806156366029913960400191505060405180910390fd5b600191505090565b60008054600160a01b900460ff1615612009576040805162461bcd60e51b815260206004820152601060248201526f14185d5cd8589b194e881c185d5cd95960821b604482015290519081900360640190fd5b6120138787610b4c565b612064576040805162461bcd60e51b815260206004820152601a60248201527f53796e53796e63426c6f636b486561646572206661696c656421000000000000604482015290519081900360640190fd5b612070858585856132ee565b6120c1576040805162461bcd60e51b815260206004820152601a60248201527f766572696679416e64457865637574655478206661696c656421000000000000604482015290519081900360640190fd5b5060019695505050505050565b600054600160a01b900460ff1690565b6120e6612404565b612125576040805162461bcd60e51b815260206004820181905260248201526000805160206157d9833981519152604482015290519081900360640190fd5b600080546040516001600160a01b03909116907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0908390a3600080546001600160a01b0319169055565b60008054600160a01b900460ff166121c5576040805162461bcd60e51b815260206004820152601460248201527314185d5cd8589b194e881b9bdd081c185d5cd95960621b604482015290519081900360640190fd5b6121cd612404565b61220c576040805162461bcd60e51b815260206004820181905260248201526000805160206157d9833981519152604482015290519081900360640190fd5b6001546040805163f2fde38b60e01b81526001600160a01b03858116600483015291519190921691829163f2fde38b9160248082019260009290919082900301818387803b15801561225d57600080fd5b505af1158015612271573d6000803e3d6000fd5b5050505060019150505b919050565b600061228a612404565b6122c9576040805162461bcd60e51b815260206004820181905260248201526000805160206157d9833981519152604482015290519081900360640190fd5b6122d16120ce565b6122dd576122dd614279565b60015460408051635c975abb60e01b815290516001600160a01b03909216918291635c975abb916004808301926020929190829003018186803b15801561232357600080fd5b505afa158015612337573d6000803e3d6000fd5b505050506040513d602081101561234d57600080fd5b5051611fae57806001600160a01b0316638456cb596040518163ffffffff1660e01b8152600401602060405180830381600087803b15801561238e57600080fd5b505af11580156123a2573d6000803e3d6000fd5b505050506040513d60208110156123b857600080fd5b5051611fae5760405162461bcd60e51b81526004018080602001828103825260278152602001806155ea6027913960400191505060405180910390fd5b6000546001600160a01b031690565b600080546001600160a01b0316612419614303565b6001600160a01b031614905090565b60008054600160a01b900460ff161561247b576040805162461bcd60e51b815260206004820152601060248201526f14185d5cd8589b194e881c185d5cd95960821b604482015290519081900360640190fd5b6001546040805163064f1acb60e41b815290516001600160a01b039092169182916364f1acb0916004808301926020929190829003018186803b1580156124c157600080fd5b505afa1580156124d5573d6000803e3d6000fd5b505050506040513d60208110156124eb57600080fd5b5051156125295760405162461bcd60e51b81526004018080602001828103825260388152602001806158ab6038913960400191505060405180910390fd5b612531615524565b61253a856139c2565b9050600060606125498661403f565b91509150816001600160601b0319168361014001516001600160601b031916146125b0576040805162461bcd60e51b815260206004820152601360248201527213995e1d109bdbdad95c9cc81a5b1b1959d85b606a1b604482015290519081900360640190fd5b60e08301516040805163021d70ab60e01b815263ffffffff831660048201908152602482019283528a5160448301528a516001600160a01b0389169463021d70ab9490938d9392606490910190602085019080838360005b83811015612620578181015183820152602001612608565b50505050905090810190601f16801561264d5780820380516001836020036101000a031916815260200191505b509350505050602060405180830381600087803b15801561266d57600080fd5b505af1158015612681573d6000803e3d6000fd5b505050506040513d602081101561269757600080fd5b50516126d45760405162461bcd60e51b815260040180806020018281038252602c815260200180615725602c913960400191505060405180910390fd5b836001600160a01b0316636d0444408460e001516040518263ffffffff1660e01b8152600401808263ffffffff166001600160401b03168152602001915050602060405180830381600087803b15801561272d57600080fd5b505af1158015612741573d6000803e3d6000fd5b505050506040513d602081101561275757600080fd5b50516127945760405162461bcd60e51b815260040180806020018281038252602d8152602001806157f9602d913960400191505060405180910390fd5b836001600160a01b031663a6c6fe7c8460e001516040518263ffffffff1660e01b8152600401808263ffffffff166001600160401b03168152602001915050602060405180830381600087803b1580156127ed57600080fd5b505af1158015612801573d6000803e3d6000fd5b505050506040513d602081101561281757600080fd5b50516128545760405162461bcd60e51b81526004018080602001828103825260378152602001806158546037913960400191505060405180910390fd5b836001600160a01b0316638b39aa256040518163ffffffff1660e01b8152600401602060405180830381600087803b15801561288f57600080fd5b505af11580156128a3573d6000803e3d6000fd5b505050506040513d60208110156128b957600080fd5b50516128f65760405162461bcd60e51b815260040180806020018281038252602e815260200180615826602e913960400191505060405180910390fd5b836001600160a01b031663c6dbd0b78460e001516040518263ffffffff1660e01b8152600401808263ffffffff166001600160401b03168152602001915050602060405180830381600087803b15801561294f57600080fd5b505af1158015612963573d6000803e3d6000fd5b505050506040513d602081101561297957600080fd5b50516129b65760405162461bcd60e51b815260040180806020018281038252602c8152602001806156ce602c913960400191505060405180910390fd5b836001600160a01b031663191b0ab18460e001516129d3846140cf565b604080516001600160e01b031960e086901b16815263ffffffff8416600482019081526024820192835283516044830152835190929160640190602085019080838360005b83811015612a30578181015183820152602001612a18565b50505050905090810190601f168015612a5d5780820380516001836020036101000a031916815260200191505b509350505050602060405180830381600087803b158015612a7d57600080fd5b505af1158015612a91573d6000803e3d6000fd5b505050506040513d6020811015612aa757600080fd5b5051612ae45760405162461bcd60e51b815260040180806020018281038252602b8152602001806156fa602b913960400191505060405180910390fd5b7ff01968fc3a2655cf1b5144cb32de6dc898f91b9239c103744e8457152ab2fbde8360e0015188604051808363ffffffff16815260200180602001828103825283818151815260200191508051906020019080838360005b83811015612b54578181015183820152602001612b3c565b50505050905090810190601f168015612b815780820380516001836020036101000a031916815260200191505b50935050505060405180910390a15060019695505050505050565b60008054600160a01b900460ff1615612bef576040805162461bcd60e51b815260206004820152601060248201526f14185d5cd8589b194e881c185d5cd95960821b604482015290519081900360640190fd5b60015460408051600162c2db5f60e01b0319815290516001600160a01b0390921691600091839163ff3d24a191600480820192602092909190829003018186803b158015612c3c57600080fd5b505afa158015612c50573d6000803e3d6000fd5b505050506040513d6020811015612c6657600080fd5b505160010190506060612c7882614307565b90506060612c8582614380565b612dbb600230612c9487614307565b60405160200180836001600160a01b03166001600160a01b031660601b815260140182805190602001908083835b60208310612ce15780518252601f199092019160209182019101612cc2565b6001836020036101000a038019825116818451168082178552505050505050905001925050506040516020818303038152906040526040518082805190602001908083835b60208310612d455780518252601f199092019160209182019101612d26565b51815160209384036101000a60001901801990921691161790526040519190930194509192505080830381855afa158015612d84573d6000803e3d6000fd5b5050506040513d6020811015612d9957600080fd5b5051604080516020818101939093528151808203909301835281019052614380565b612dcc612dc733614446565b614380565b612dd58f614461565b612e148f8f8080601f01602080910402602001604051908101604052809392919081815260200183838082843760009201919091525061438092505050565b612e538e8e8080601f01602080910402602001604051908101604052809392919081815260200183838082843760009201919091525061438092505050565b612e928d8d8080601f01602080910402602001604051908101604052809392919081815260200183838082843760009201919091525061438092505050565b6040516020018088805190602001908083835b60208310612ec45780518252601f199092019160209182019101612ea5565b51815160209384036101000a60001901801990921691161790528a5191909301928a0191508083835b60208310612f0c5780518252601f199092019160209182019101612eed565b51815160209384036101000a600019018019909216911617905289519190930192890191508083835b60208310612f545780518252601f199092019160209182019101612f35565b51815160209384036101000a600019018019909216911617905288519190930192880191508083835b60208310612f9c5780518252601f199092019160209182019101612f7d565b51815160209384036101000a600019018019909216911617905287519190930192870191508083835b60208310612fe45780518252601f199092019160209182019101612fc5565b51815160209384036101000a600019018019909216911617905286519190930192860191508083835b6020831061302c5780518252601f19909201916020918201910161300d565b51815160209384036101000a600019018019909216911617905285519190930192850191508083835b602083106130745780518252601f199092019160209182019101613055565b6001836020036101000a0380198251168184511680821785525050505050509050019750505050505050506040516020818303038152906040529050836001600160a01b0316639c59af4f8483805190602001206040518363ffffffff1660e01b81526004018083815260200182815260200192505050602060405180830381600087803b15801561310557600080fd5b505af1158015613119573d6000803e3d6000fd5b505050506040513d602081101561312f57600080fd5b505161316c5760405162461bcd60e51b81526004018080602001828103825260308152602001806159056030913960400191505060405180910390fd5b326001600160a01b03167f6ad3bf15c1988bc04bc153490cab16db8efb9a3990215bf1c64ea6e28be8848383338f8f8f876040518080602001876001600160a01b03166001600160a01b03168152602001866001600160401b03166001600160401b03168152602001806020018060200184810384528a818151815260200191508051906020019080838360005b838110156132125781810151838201526020016131fa565b50505050905090810190601f16801561323f5780820380516001836020036101000a031916815260200191505b5084810383528681526020018787808284376000838201819052601f909101601f191690920186810384528751815287516020918201939189019250908190849084905b8381101561329b578181015183820152602001613283565b50505050905090810190601f1680156132c85780820380516001836020036101000a031916815260200191505b50995050505050505050505060405180910390a25060019b9a5050505050505050505050565b60008054600160a01b900460ff1615613341576040805162461bcd60e51b815260206004820152601060248201526f14185d5cd8589b194e881c185d5cd95960821b604482015290519081900360640190fd5b600154604080516313b4763360e21b815290516001600160a01b03909216916000918391634ed1d8cc91600480820192602092909190829003018186803b15801561338b57600080fd5b505afa15801561339f573d6000803e3d6000fd5b505050506040513d60208110156133b557600080fd5b505190506001600160401b03808216908516111561341a576040805162461bcd60e51b815260206004820152601b60248201527f626c6f636b486569676874203e204c6174657374486569676874210000000000604482015290519081900360640190fd5b6060826001600160a01b0316639743e9d6866040518263ffffffff1660e01b815260040180826001600160401b03166001600160401b0316815260200191505060006040518083038186803b15801561347257600080fd5b505afa158015613486573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f1916820160405260208110156134af57600080fd5b8101908080516040519392919084600160201b8211156134ce57600080fd5b9083019060208201858111156134e357600080fd5b8251600160201b8111828201881017156134fc57600080fd5b82525081516020918201929091019080838360005b83811015613529578181015183820152602001613511565b50505050905090810190601f1680156135565780820380516001836020036101000a031916815260200191505b506040525050509050613567615524565b613570826139c2565b905061358a898983608001516135858b6144a4565b6145a8565b6135c55760405162461bcd60e51b815260040180806020018281038252602681526020018061565f6026913960400191505060405180910390fd5b6135cd61557f565b6135d688614659565b9050846001600160a01b031663cceb218d6135f48360000151614735565b6040518263ffffffff1660e01b81526004018082815260200191505060206040518083038186803b15801561362857600080fd5b505afa15801561363c573d6000803e3d6000fd5b505050506040513d602081101561365257600080fd5b5051156136905760405162461bcd60e51b81526004018080602001828103825260228152602001806158e36022913960400191505060405180910390fd5b846001600160a01b031663794b03006136ac8360000151614735565b6040518263ffffffff1660e01b815260040180828152602001915050602060405180830381600087803b1580156136e257600080fd5b505af11580156136f6573d6000803e3d6000fd5b505050506040513d602081101561370c57600080fd5b505161375f576040805162461bcd60e51b815260206004820181905260248201527f536176652063726f7373636861696e207478206578697374206661696c656421604482015290519081900360640190fd5b6000613772826040015160800151614795565b6040830151606001519091506001600160401b03166002146137c55760405162461bcd60e51b815260040180806020018281038252602b815260200180615751602b913960400191505060405180910390fd5b7f5944a0908ae763fa254df8e0da54a682cbc62115501336e18be52a7dd774b80b8260200151836040015160800151846000015185604001516000015160405180856001600160401b03166001600160401b03168152602001806020018060200180602001848103845287818151815260200191508051906020019080838360005b8381101561385f578181015183820152602001613847565b50505050905090810190601f16801561388c5780820380516001836020036101000a031916815260200191505b50848103835286518152865160209182019188019080838360005b838110156138bf5781810151838201526020016138a7565b50505050905090810190601f1680156138ec5780820380516001836020036101000a031916815260200191505b50848103825285518152855160209182019187019080838360005b8381101561391f578181015183820152602001613907565b50505050905090810190601f16801561394c5780820380516001836020036101000a031916815260200191505b5097505050505050505060405180910390a15060019a9950505050505050505050565b613977612404565b6139b6576040805162461bcd60e51b815260206004820181905260248201526000805160206157d9833981519152604482015290519081900360640190fd5b6139bf816147df565b50565b6139ca615524565b6139d2615524565b60006139de848261487f565b63ffffffff909116835290506139f4848261491d565b6001600160401b0390911660208401529050613a1084826149b9565b60408401919091529050613a2484826149b9565b60608401919091529050613a3884826149b9565b60808401919091529050613a4c84826149b9565b60a08401919091529050613a60848261487f565b63ffffffff90911660c08401529050613a79848261487f565b63ffffffff90911660e08401529050613a92848261491d565b6001600160401b039091166101008401529050613aaf8482614a13565b6101208401919091529050613ac48482614add565b506001600160601b0319166101408301525092915050565b6000806000846001600160401b031611613b3d576040805162461bcd60e51b815260206004820152601d60248201527f626f6f6b206b6565706572206c6973742063616e6e6f7420656d707479000000604482015290519081900360640190fd5b836001600160401b0316855114613b94576040805162461bcd60e51b815260206004820152601660248201527563616e6e6f74207061727469616c6c7920717565727960501b604482015290519081900360640190fd5b60008311613be9576040805162461bcd60e51b815260206004820152601d60248201527f626c6f636b20686569676874206d75737420626520706f736974697665000000604482015290519081900360640190fd5b600060001985016001600160401b03861660011415613c1257506000925060019150613d219050565b806001600160401b0316826001600160401b031611613cc157600060018383036001600160401b0316901c830190508588826001600160401b031681518110613c5757fe5b60200260200101516001600160401b03161415613c7c57935060019250613d21915050565b8588826001600160401b031681518110613c9257fe5b60200260200101516001600160401b03161015613cb457806001019250613cbb565b6001810391505b50613c12565b6001826001600160401b031610158015613d0257508487600184036001600160401b031681518110613cef57fe5b60200260200101516001600160401b0316105b15613d17575060001901915060019050613d21565b5060009250829150505b935093915050565b6060600080613d38848261491d565b80935081925050506060816001600160401b0316604051908082528060200260200182016040528015613d75578160200160208202803883390190505b509050606060005b836001600160401b0316811015613dd157613d988786614a13565b95509150613da582614795565b838281518110613db157fe5b6001600160a01b0390921660209283029190910190910152600101613d7d565b509095945050505050565b600080600280876040518082805190602001908083835b60208310613e125780518252601f199092019160209182019101613df3565b51815160209384036101000a60001901801990921691161790526040519190930194509192505080830381855afa158015613e51573d6000803e3d6000fd5b5050506040513d6020811015613e6657600080fd5b50516040805160208181019390935281518082038401815290820191829052805190928291908401908083835b60208310613eb25780518252601f199092019160209182019101613e93565b51815160209384036101000a60001901801990921691161790526040519190930194509192505080830381855afa158015613ef1573d6000803e3d6000fd5b5050506040513d6020811015613f0657600080fd5b5051855190915060009060419004815b81811015614030576000613f37613f328a604185026020614b3b565b614735565b90506000613f50613f328b604186026020016020614b3b565b905060008a6041850260400181518110613f6657fe5b602001015160f81c60f81b60f81c601b0190506000600188604051602001808281526020019150506040516020818303038152906040528051906020012083868660405160008152602001604052604051808581526020018460ff1660ff1681526020018381526020018281526020019450505050506020604051602081039080840390855afa158015613ffe573d6000803e3d6000fd5b5050506020604051035190506140148b82614bbb565b15614020576001870196505b505060019092019150613f169050565b50509092111595945050505050565b60006060604383518161404e57fe5b06156140a1576040805162461bcd60e51b815260206004820152601b60248201527f5f7075624b65794c697374206c656e67746820696c6c6567616c210000000000604482015290519081900360640190fd5b600060438451816140ae57fe5b0490506140c5816003600019820104830386614c0d565b9250925050915091565b8051606090816140de82614461565b905060005b828110156141c9578161410b612dc78784815181106140fe57fe5b6020026020010151614446565b6040516020018083805190602001908083835b6020831061413d5780518252601f19909201916020918201910161411e565b51815160209384036101000a600019018019909216911617905285519190930192850191508083835b602083106141855780518252601f199092019160209182019101614166565b6001836020036101000a03801982511681845116808217855250505050505090500192505050604051602081830303815290604052915080806001019150506140e3565b509392505050565b600054600160a01b900460ff16614226576040805162461bcd60e51b815260206004820152601460248201527314185d5cd8589b194e881b9bdd081c185d5cd95960621b604482015290519081900360640190fd5b6000805460ff60a01b191690557f5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa61425c614303565b604080516001600160a01b039092168252519081900360200190a1565b600054600160a01b900460ff16156142cb576040805162461bcd60e51b815260206004820152601060248201526f14185d5cd8589b194e881c185d5cd95960821b604482015290519081900360640190fd5b6000805460ff60a01b1916600160a01b1790557f62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a25861425c5b3390565b60606001600160ff1b03821115614365576040805162461bcd60e51b815260206004820152601760248201527f56616c75652065786365656473207468652072616e6765000000000000000000604482015290519081900360640190fd5b60405190506020815281602082015260408101604052919050565b805160609061438e81615013565b836040516020018083805190602001908083835b602083106143c15780518252601f1990920191602091820191016143a2565b51815160209384036101000a600019018019909216911617905285519190930192850191508083835b602083106144095780518252601f1990920191602091820191016143ea565b6001836020036101000a03801982511681845116808217855250505050505090500192505050604051602081830303815290604052915050919050565b604080516014815260609290921b6020830152818101905290565b6040516008808252606091906000601f5b828210156144945785811a826020860101536001919091019060001901614472565b5050506040818101905292915050565b6040516000602080830182815284519293600293859387939260210191908401908083835b602083106144e85780518252601f1990920191602091820191016144c9565b6001836020036101000a038019825116818451168082178552505050505050905001925050506040516020818303038152906040526040518082805190602001908083835b6020831061454c5780518252601f19909201916020918201910161452d565b51815160209384036101000a60001901801990921691161790526040519190930194509192505080830381855afa15801561458b573d6000803e3d6000fd5b5050506040513d60208110156145a057600080fd5b505192915050565b60008181805b875181101561464c578781815181106145c357fe5b60200260200101519150600060018111156145da57fe5b8782815181106145e657fe5b602002602001015160018111156145f957fe5b1415614610576146098284615159565b9250614644565b600187828151811061461e57fe5b6020026020010151600181111561463157fe5b1415614644576146418383615159565b92505b6001016145ae565b5050909214949350505050565b61466161557f565b61466961557f565b60006146758482614a13565b9083529050614684848261491d565b6001600160401b039091166020840152905061469e6155a3565b6146a88583614a13565b90825291506146b78583614a13565b602083019190915291506146cb8583614a13565b604083019190915291506146df858361491d565b6001600160401b03909116606083015291506146fb8583614a13565b6080830191909152915061470f8583614a13565b60a083019190915291506147238583614a13565b5060c082015260408301525092915050565b6000815160201461478d576040805162461bcd60e51b815260206004820152601760248201527f6279746573206c656e677468206973206e6f742033322e000000000000000000604482015290519081900360640190fd5b506020015190565b600081516014146147d75760405162461bcd60e51b81526004018080602001828103825260238152602001806156856023913960400191505060405180910390fd5b506014015190565b6001600160a01b0381166148245760405162461bcd60e51b81526004018080602001828103825260268152602001806156a86026913960400191505060405180910390fd5b600080546040516001600160a01b03808516939216917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e091a3600080546001600160a01b0319166001600160a01b0392909216919091179055565b60008083518360040111156148c9576040805162461bcd60e51b8152602060048201526016602482015260008051602061588b833981519152604482015290519081900360640190fd5b600060405160046000600182038760208a0101515b838310156148fe5780821a838601536001830192506001820391506148de565b50505080820160405260200390035192505050600482015b9250929050565b6000808351836008011115614967576040805162461bcd60e51b8152602060048201526016602482015260008051602061588b833981519152604482015290519081900360640190fd5b600060405160086000600182038760208a0101515b8383101561499c5780821a8386015360018301925060018203915061497c565b505050808201604052602003900351956008949094019450505050565b6000808351836020011115614a03576040805162461bcd60e51b8152602060048201526016602482015260008051602061588b833981519152604482015290519081900360640190fd5b5050602091810182015192910190565b6060600080614a22858561521d565b86519095509091508185011115614a6e576040805162461bcd60e51b8152602060048201526016602482015260008051602061588b833981519152604482015290519081900360640190fd5b606081158015614a8957604051915060208201604052614ad3565b6040519150601f8316801560200281840101848101888315602002848c0101015b81831015614ac2578051835260209283019201614aaa565b5050848452601f01601f1916604052505b5095930193505050565b6000808351836014011115614b27576040805162461bcd60e51b8152602060048201526016602482015260008051602061588b833981519152604482015290519081900360640190fd5b505081810160200151601482019250929050565b606081830184511015614b4d57600080fd5b606082158015614b6857604051915060208201604052614bb2565b6040519150601f8416801560200281840101858101878315602002848b0101015b81831015614ba1578051835260209283019201614b89565b5050858452601f01601f1916604052505b50949350505050565b6000805b8351811015614c0657838181518110614bd457fe5b60200260200101516001600160a01b0316836001600160a01b03161415614bfe5760019150614c06565b600101614bbf565b5092915050565b600060608080614c1c876152d3565b6040516020018083805190602001908083835b60208310614c4e5780518252601f199092019160209182019101614c2f565b51815160209384036101000a600019018019909216911617905285519190930192850191508083835b60208310614c965780518252601f199092019160209182019101614c77565b6001836020036101000a038019825116818451168082178552505050505050905001925050506040516020818303038152906040529050606086604051908082528060200260200182016040528015614cf9578160200160208202803883390190505b50905060005b87811015614e2a5782614d22612dc7614d1d89604386026043614b3b565b6152fc565b6040516020018083805190602001908083835b60208310614d545780518252601f199092019160209182019101614d35565b51815160209384036101000a600019018019909216911617905285519190930192850191508083835b60208310614d9c5780518252601f199092019160209182019101614d7d565b6001836020036101000a0380198251168184511680821785525050505050509050019250505060405160208183030381529060405292506000614df0614de788604385026043614b3b565b60036040614b3b565b8051906020012090508060001c838381518110614e0957fe5b6001600160a01b039092166020928302919091019091015250600101614cff565b5081614e35876152d3565b6040516020018083805190602001908083835b60208310614e675780518252601f199092019160209182019101614e48565b51815160209384036101000a600019018019909216911617905285519190930192850191508083835b60208310614eaf5780518252601f199092019160209182019101614e90565b6001836020036101000a038019825116818451168082178552505050505050905001925050506040516020818303038152906040529150600060036002846040518082805190602001908083835b60208310614f1c5780518252601f199092019160209182019101614efd565b51815160209384036101000a60001901801990921691161790526040519190930194509192505080830381855afa158015614f5b573d6000803e3d6000fd5b5050506040513d6020811015614f7057600080fd5b50516040805160208181019390935281518082038401815290820191829052805190928291908401908083835b60208310614fbc5780518252601f199092019160209182019101614f9d565b51815160209384036101000a60001901801990921691161790526040519190930194509192505080830381855afa158015614ffb573d6000803e3d6000fd5b50506040515160601b99929850919650505050505050565b606060fd826001600160401b0316101561503757615030826153ea565b905061227b565b61ffff826001600160401b0316116151155761505660fd60f81b615405565b61505f836152d3565b6040516020018083805190602001908083835b602083106150915780518252601f199092019160209182019101615072565b51815160209384036101000a600019018019909216911617905285519190930192850191508083835b602083106150d95780518252601f1990920191602091820191016150ba565b6001836020036101000a03801982511681845116808217855250505050505090500192505050604051602081830303815290604052905061227b565b63ffffffff826001600160401b03161161513f57615136607f60f91b615405565b61505f83615413565b6151506001600160f81b0319615405565b61505f83614461565b60408051600160f81b60208083019190915260218201859052604180830185905283518084039091018152606190920192839052815160009360029392909182918401908083835b602083106151c05780518252601f1990920191602091820191016151a1565b51815160209384036101000a60001901801990921691161790526040519190930194509192505080830381855afa1580156151ff573d6000803e3d6000fd5b5050506040513d602081101561521457600080fd5b50519392505050565b600080600061522c858561543c565b9450905060fd60f81b6001600160f81b0319821614156152615761525085856154a5565b8161ffff1691509250925050614916565b607f60f91b6001600160f81b03198216141561529457615281858561487f565b8163ffffffff1691509250925050614916565b6001600160f81b031980821614156152c6576152b0858561491d565b816001600160401b031691509250925050614916565b60f81c9150829050614916565b6040516002808252606091906000601f60ff861660208501536001919091019060001901614472565b6060602282511015615355576040805162461bcd60e51b815260206004820152601760248201527f6b6579206c656e67676820697320746f6f2073686f7274000000000000000000604482015290519081900360640190fd5b6153628260006023614b3b565b905060028260428151811061537357fe5b016020015160f81c8161538257fe5b0660ff16600014156153bc57600260f81b816002815181106153a057fe5b60200101906001600160f81b031916908160001a90535061227b565b600360f81b816002815181106153ce57fe5b60200101906001600160f81b031916908160001a905350919050565b604080516001815260f89290921b6020830152818101905290565b60606114ac8260f81c6153ea565b6040516004808252606091906000601f60ff861660208501536001919091019060001901614472565b6000808351836001011115615491576040805162461bcd60e51b81526020600482015260166024820152754f66667365742065786365656473206d6178696d756d60501b604482015290519081900360640190fd5b505081810160200151600182019250929050565b60008083518360020111156154ef576040805162461bcd60e51b8152602060048201526016602482015260008051602061588b833981519152604482015290519081900360640190fd5b6000604051846020870101518060011a82538060001a6001830153506002818101604052601d19909101519694019450505050565b604080516101608101825260008082526020820181905291810182905260608082018390526080820183905260a0820183905260c0820183905260e08201839052610100820183905261012082015261014081019190915290565b604080516060808201835281526000602082015290810161559e6155a3565b905290565b6040518060e0016040528060608152602001606081526020016060815260200160006001600160401b03168152602001606081526020016060815260200160608152509056fe70617573652045746843726f7373436861696e4461746120636f6e7472616374206661696c6564546865206e657874426f6f6b4b6565706572206f662068656164657220697320656d707479756e70617573652045746843726f7373436861696e4461746120636f6e7472616374206661696c65645665726966792063726f7373636861696e206d65726b6c652070726f6f66206661696c6564216279746573206c656e67746820646f6573206e6f74206d6174636820616464726573734f776e61626c653a206e6577206f776e657220697320746865207a65726f206164647265737353617665204d434b656570657248656967687420746f204461746120636f6e7472616374206661696c65642153617665204d434b6565706572427974657320746f204461746120636f6e7472616374206661696c65642153617665204d4348656164657220627974657320746f204461746120636f6e7472616374206661696c65642154686973205478206973206e6f742061696d696e672061742045746865726565756d206e6574776f726b2163616e206e6f742066696e6420626c6f636b2068656967687420696e20626f6f6b206b656570657220686569676874206c6973744865616465722e6e6578746e657874426f6f6b6b65657065722073686f756c6420626520656d7074794f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e657253617665204d43204c617465737448656967687420746f204461746120636f6e7472616374206661696c6564215361766520696e697447656e65736973426c6f636b20746f204461746120636f6e7472616374206661696c65642153617665204d43204c6174657374426f6f6b4b656570657248656967687420746f204461746120636f6e7472616374206661696c6564216f66667365742065786365656473206d6178696d756d0000000000000000000045746843726f7373436861696e4461746120636f6e74726163742068617320616c7265616479206265656e20696e697469616c697a656421746865207472616e73616374696f6e20686173206265656e20657865637574656421536176652065746854784861736820627920696e64657820746f204461746120636f6e7472616374206661696c656421a265627a7a7231582031f15093236ac88ed50d5234a0352cfc632d1da6523116629ef2f49e29430f8b64736f6c634300050f0032"

// DeployEthCrossChainManager deploys a new Ethereum contract, binding an instance of EthCrossChainManager to it.
func DeployEthCrossChainManager(auth *bind.TransactOpts, backend bind.ContractBackend, _eccd common.Address) (common.Address, *types.Transaction, *EthCrossChainManager, error) {
	parsed, err := abi.JSON(strings.NewReader(EthCrossChainManagerABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(EthCrossChainManagerBin), backend, _eccd)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &EthCrossChainManager{EthCrossChainManagerCaller: EthCrossChainManagerCaller{contract: contract}, EthCrossChainManagerTransactor: EthCrossChainManagerTransactor{contract: contract}, EthCrossChainManagerFilterer: EthCrossChainManagerFilterer{contract: contract}}, nil
}

// EthCrossChainManager is an auto generated Go binding around an Ethereum contract.
type EthCrossChainManager struct {
	EthCrossChainManagerCaller     // Read-only binding to the contract
	EthCrossChainManagerTransactor // Write-only binding to the contract
	EthCrossChainManagerFilterer   // Log filterer for contract events
}

// EthCrossChainManagerCaller is an auto generated read-only Go binding around an Ethereum contract.
type EthCrossChainManagerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EthCrossChainManagerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type EthCrossChainManagerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EthCrossChainManagerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type EthCrossChainManagerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EthCrossChainManagerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type EthCrossChainManagerSession struct {
	Contract     *EthCrossChainManager // Generic contract binding to set the session for
	CallOpts     bind.CallOpts         // Call options to use throughout this session
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// EthCrossChainManagerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type EthCrossChainManagerCallerSession struct {
	Contract *EthCrossChainManagerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts               // Call options to use throughout this session
}

// EthCrossChainManagerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type EthCrossChainManagerTransactorSession struct {
	Contract     *EthCrossChainManagerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts               // Transaction auth options to use throughout this session
}

// EthCrossChainManagerRaw is an auto generated low-level Go binding around an Ethereum contract.
type EthCrossChainManagerRaw struct {
	Contract *EthCrossChainManager // Generic contract binding to access the raw methods on
}

// EthCrossChainManagerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type EthCrossChainManagerCallerRaw struct {
	Contract *EthCrossChainManagerCaller // Generic read-only contract binding to access the raw methods on
}

// EthCrossChainManagerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type EthCrossChainManagerTransactorRaw struct {
	Contract *EthCrossChainManagerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewEthCrossChainManager creates a new instance of EthCrossChainManager, bound to a specific deployed contract.
func NewEthCrossChainManager(address common.Address, backend bind.ContractBackend) (*EthCrossChainManager, error) {
	contract, err := bindEthCrossChainManager(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &EthCrossChainManager{EthCrossChainManagerCaller: EthCrossChainManagerCaller{contract: contract}, EthCrossChainManagerTransactor: EthCrossChainManagerTransactor{contract: contract}, EthCrossChainManagerFilterer: EthCrossChainManagerFilterer{contract: contract}}, nil
}

// NewEthCrossChainManagerCaller creates a new read-only instance of EthCrossChainManager, bound to a specific deployed contract.
func NewEthCrossChainManagerCaller(address common.Address, caller bind.ContractCaller) (*EthCrossChainManagerCaller, error) {
	contract, err := bindEthCrossChainManager(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &EthCrossChainManagerCaller{contract: contract}, nil
}

// NewEthCrossChainManagerTransactor creates a new write-only instance of EthCrossChainManager, bound to a specific deployed contract.
func NewEthCrossChainManagerTransactor(address common.Address, transactor bind.ContractTransactor) (*EthCrossChainManagerTransactor, error) {
	contract, err := bindEthCrossChainManager(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &EthCrossChainManagerTransactor{contract: contract}, nil
}

// NewEthCrossChainManagerFilterer creates a new log filterer instance of EthCrossChainManager, bound to a specific deployed contract.
func NewEthCrossChainManagerFilterer(address common.Address, filterer bind.ContractFilterer) (*EthCrossChainManagerFilterer, error) {
	contract, err := bindEthCrossChainManager(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &EthCrossChainManagerFilterer{contract: contract}, nil
}

// bindEthCrossChainManager binds a generic wrapper to an already deployed contract.
func bindEthCrossChainManager(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(EthCrossChainManagerABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_EthCrossChainManager *EthCrossChainManagerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _EthCrossChainManager.Contract.EthCrossChainManagerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_EthCrossChainManager *EthCrossChainManagerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EthCrossChainManager.Contract.EthCrossChainManagerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_EthCrossChainManager *EthCrossChainManagerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _EthCrossChainManager.Contract.EthCrossChainManagerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_EthCrossChainManager *EthCrossChainManagerCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _EthCrossChainManager.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_EthCrossChainManager *EthCrossChainManagerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EthCrossChainManager.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_EthCrossChainManager *EthCrossChainManagerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _EthCrossChainManager.Contract.contract.Transact(opts, method, params...)
}

// EthCrossChainDataAddress is a free data retrieval call binding the contract method 0x00ba1694.
//
// Solidity: function EthCrossChainDataAddress() constant returns(address)
func (_EthCrossChainManager *EthCrossChainManagerCaller) EthCrossChainDataAddress(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _EthCrossChainManager.contract.Call(opts, out, "EthCrossChainDataAddress")
	return *ret0, err
}

// EthCrossChainDataAddress is a free data retrieval call binding the contract method 0x00ba1694.
//
// Solidity: function EthCrossChainDataAddress() constant returns(address)
func (_EthCrossChainManager *EthCrossChainManagerSession) EthCrossChainDataAddress() (common.Address, error) {
	return _EthCrossChainManager.Contract.EthCrossChainDataAddress(&_EthCrossChainManager.CallOpts)
}

// EthCrossChainDataAddress is a free data retrieval call binding the contract method 0x00ba1694.
//
// Solidity: function EthCrossChainDataAddress() constant returns(address)
func (_EthCrossChainManager *EthCrossChainManagerCallerSession) EthCrossChainDataAddress() (common.Address, error) {
	return _EthCrossChainManager.Contract.EthCrossChainDataAddress(&_EthCrossChainManager.CallOpts)
}

// IsOwner is a free data retrieval call binding the contract method 0x8f32d59b.
//
// Solidity: function isOwner() constant returns(bool)
func (_EthCrossChainManager *EthCrossChainManagerCaller) IsOwner(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _EthCrossChainManager.contract.Call(opts, out, "isOwner")
	return *ret0, err
}

// IsOwner is a free data retrieval call binding the contract method 0x8f32d59b.
//
// Solidity: function isOwner() constant returns(bool)
func (_EthCrossChainManager *EthCrossChainManagerSession) IsOwner() (bool, error) {
	return _EthCrossChainManager.Contract.IsOwner(&_EthCrossChainManager.CallOpts)
}

// IsOwner is a free data retrieval call binding the contract method 0x8f32d59b.
//
// Solidity: function isOwner() constant returns(bool)
func (_EthCrossChainManager *EthCrossChainManagerCallerSession) IsOwner() (bool, error) {
	return _EthCrossChainManager.Contract.IsOwner(&_EthCrossChainManager.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_EthCrossChainManager *EthCrossChainManagerCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _EthCrossChainManager.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_EthCrossChainManager *EthCrossChainManagerSession) Owner() (common.Address, error) {
	return _EthCrossChainManager.Contract.Owner(&_EthCrossChainManager.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_EthCrossChainManager *EthCrossChainManagerCallerSession) Owner() (common.Address, error) {
	return _EthCrossChainManager.Contract.Owner(&_EthCrossChainManager.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() constant returns(bool)
func (_EthCrossChainManager *EthCrossChainManagerCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _EthCrossChainManager.contract.Call(opts, out, "paused")
	return *ret0, err
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() constant returns(bool)
func (_EthCrossChainManager *EthCrossChainManagerSession) Paused() (bool, error) {
	return _EthCrossChainManager.Contract.Paused(&_EthCrossChainManager.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() constant returns(bool)
func (_EthCrossChainManager *EthCrossChainManagerCallerSession) Paused() (bool, error) {
	return _EthCrossChainManager.Contract.Paused(&_EthCrossChainManager.CallOpts)
}

// ChangeBookKeeper is a paid mutator transaction binding the contract method 0x21f966aa.
//
// Solidity: function ChangeBookKeeper(bytes _rawHeader, bytes _pubKeyList, bytes _sigList) returns(bool)
func (_EthCrossChainManager *EthCrossChainManagerTransactor) ChangeBookKeeper(opts *bind.TransactOpts, _rawHeader []byte, _pubKeyList []byte, _sigList []byte) (*types.Transaction, error) {
	return _EthCrossChainManager.contract.Transact(opts, "ChangeBookKeeper", _rawHeader, _pubKeyList, _sigList)
}

// ChangeBookKeeper is a paid mutator transaction binding the contract method 0x21f966aa.
//
// Solidity: function ChangeBookKeeper(bytes _rawHeader, bytes _pubKeyList, bytes _sigList) returns(bool)
func (_EthCrossChainManager *EthCrossChainManagerSession) ChangeBookKeeper(_rawHeader []byte, _pubKeyList []byte, _sigList []byte) (*types.Transaction, error) {
	return _EthCrossChainManager.Contract.ChangeBookKeeper(&_EthCrossChainManager.TransactOpts, _rawHeader, _pubKeyList, _sigList)
}

// ChangeBookKeeper is a paid mutator transaction binding the contract method 0x21f966aa.
//
// Solidity: function ChangeBookKeeper(bytes _rawHeader, bytes _pubKeyList, bytes _sigList) returns(bool)
func (_EthCrossChainManager *EthCrossChainManagerTransactorSession) ChangeBookKeeper(_rawHeader []byte, _pubKeyList []byte, _sigList []byte) (*types.Transaction, error) {
	return _EthCrossChainManager.Contract.ChangeBookKeeper(&_EthCrossChainManager.TransactOpts, _rawHeader, _pubKeyList, _sigList)
}

// InitGenesisBlock is a paid mutator transaction binding the contract method 0xadbccaf5.
//
// Solidity: function InitGenesisBlock(bytes _rawHeader, bytes _pubKeyList) returns(bool)
func (_EthCrossChainManager *EthCrossChainManagerTransactor) InitGenesisBlock(opts *bind.TransactOpts, _rawHeader []byte, _pubKeyList []byte) (*types.Transaction, error) {
	return _EthCrossChainManager.contract.Transact(opts, "InitGenesisBlock", _rawHeader, _pubKeyList)
}

// InitGenesisBlock is a paid mutator transaction binding the contract method 0xadbccaf5.
//
// Solidity: function InitGenesisBlock(bytes _rawHeader, bytes _pubKeyList) returns(bool)
func (_EthCrossChainManager *EthCrossChainManagerSession) InitGenesisBlock(_rawHeader []byte, _pubKeyList []byte) (*types.Transaction, error) {
	return _EthCrossChainManager.Contract.InitGenesisBlock(&_EthCrossChainManager.TransactOpts, _rawHeader, _pubKeyList)
}

// InitGenesisBlock is a paid mutator transaction binding the contract method 0xadbccaf5.
//
// Solidity: function InitGenesisBlock(bytes _rawHeader, bytes _pubKeyList) returns(bool)
func (_EthCrossChainManager *EthCrossChainManagerTransactorSession) InitGenesisBlock(_rawHeader []byte, _pubKeyList []byte) (*types.Transaction, error) {
	return _EthCrossChainManager.Contract.InitGenesisBlock(&_EthCrossChainManager.TransactOpts, _rawHeader, _pubKeyList)
}

// SyncAndVerify is a paid mutator transaction binding the contract method 0x48c2f675.
//
// Solidity: function SyncAndVerify(bytes rawHeader, bytes sigList, bytes32[] proof, uint8[] position, bytes toMerkleValueBs, uint64 blockHeight) returns(bool)
func (_EthCrossChainManager *EthCrossChainManagerTransactor) SyncAndVerify(opts *bind.TransactOpts, rawHeader []byte, sigList []byte, proof [][32]byte, position []uint8, toMerkleValueBs []byte, blockHeight uint64) (*types.Transaction, error) {
	return _EthCrossChainManager.contract.Transact(opts, "SyncAndVerify", rawHeader, sigList, proof, position, toMerkleValueBs, blockHeight)
}

// SyncAndVerify is a paid mutator transaction binding the contract method 0x48c2f675.
//
// Solidity: function SyncAndVerify(bytes rawHeader, bytes sigList, bytes32[] proof, uint8[] position, bytes toMerkleValueBs, uint64 blockHeight) returns(bool)
func (_EthCrossChainManager *EthCrossChainManagerSession) SyncAndVerify(rawHeader []byte, sigList []byte, proof [][32]byte, position []uint8, toMerkleValueBs []byte, blockHeight uint64) (*types.Transaction, error) {
	return _EthCrossChainManager.Contract.SyncAndVerify(&_EthCrossChainManager.TransactOpts, rawHeader, sigList, proof, position, toMerkleValueBs, blockHeight)
}

// SyncAndVerify is a paid mutator transaction binding the contract method 0x48c2f675.
//
// Solidity: function SyncAndVerify(bytes rawHeader, bytes sigList, bytes32[] proof, uint8[] position, bytes toMerkleValueBs, uint64 blockHeight) returns(bool)
func (_EthCrossChainManager *EthCrossChainManagerTransactorSession) SyncAndVerify(rawHeader []byte, sigList []byte, proof [][32]byte, position []uint8, toMerkleValueBs []byte, blockHeight uint64) (*types.Transaction, error) {
	return _EthCrossChainManager.Contract.SyncAndVerify(&_EthCrossChainManager.TransactOpts, rawHeader, sigList, proof, position, toMerkleValueBs, blockHeight)
}

// SyncBlockHeader is a paid mutator transaction binding the contract method 0x1e568e18.
//
// Solidity: function SyncBlockHeader(bytes _rawHeader, bytes _sigList) returns(bool)
func (_EthCrossChainManager *EthCrossChainManagerTransactor) SyncBlockHeader(opts *bind.TransactOpts, _rawHeader []byte, _sigList []byte) (*types.Transaction, error) {
	return _EthCrossChainManager.contract.Transact(opts, "SyncBlockHeader", _rawHeader, _sigList)
}

// SyncBlockHeader is a paid mutator transaction binding the contract method 0x1e568e18.
//
// Solidity: function SyncBlockHeader(bytes _rawHeader, bytes _sigList) returns(bool)
func (_EthCrossChainManager *EthCrossChainManagerSession) SyncBlockHeader(_rawHeader []byte, _sigList []byte) (*types.Transaction, error) {
	return _EthCrossChainManager.Contract.SyncBlockHeader(&_EthCrossChainManager.TransactOpts, _rawHeader, _sigList)
}

// SyncBlockHeader is a paid mutator transaction binding the contract method 0x1e568e18.
//
// Solidity: function SyncBlockHeader(bytes _rawHeader, bytes _sigList) returns(bool)
func (_EthCrossChainManager *EthCrossChainManagerTransactorSession) SyncBlockHeader(_rawHeader []byte, _sigList []byte) (*types.Transaction, error) {
	return _EthCrossChainManager.Contract.SyncBlockHeader(&_EthCrossChainManager.TransactOpts, _rawHeader, _sigList)
}

// CrossChain is a paid mutator transaction binding the contract method 0xbd5cf625.
//
// Solidity: function crossChain(uint64 _toChainId, bytes _toContract, bytes _method, bytes _txData) returns(bool)
func (_EthCrossChainManager *EthCrossChainManagerTransactor) CrossChain(opts *bind.TransactOpts, _toChainId uint64, _toContract []byte, _method []byte, _txData []byte) (*types.Transaction, error) {
	return _EthCrossChainManager.contract.Transact(opts, "crossChain", _toChainId, _toContract, _method, _txData)
}

// CrossChain is a paid mutator transaction binding the contract method 0xbd5cf625.
//
// Solidity: function crossChain(uint64 _toChainId, bytes _toContract, bytes _method, bytes _txData) returns(bool)
func (_EthCrossChainManager *EthCrossChainManagerSession) CrossChain(_toChainId uint64, _toContract []byte, _method []byte, _txData []byte) (*types.Transaction, error) {
	return _EthCrossChainManager.Contract.CrossChain(&_EthCrossChainManager.TransactOpts, _toChainId, _toContract, _method, _txData)
}

// CrossChain is a paid mutator transaction binding the contract method 0xbd5cf625.
//
// Solidity: function crossChain(uint64 _toChainId, bytes _toContract, bytes _method, bytes _txData) returns(bool)
func (_EthCrossChainManager *EthCrossChainManagerTransactorSession) CrossChain(_toChainId uint64, _toContract []byte, _method []byte, _txData []byte) (*types.Transaction, error) {
	return _EthCrossChainManager.Contract.CrossChain(&_EthCrossChainManager.TransactOpts, _toChainId, _toContract, _method, _txData)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns(bool)
func (_EthCrossChainManager *EthCrossChainManagerTransactor) Pause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EthCrossChainManager.contract.Transact(opts, "pause")
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns(bool)
func (_EthCrossChainManager *EthCrossChainManagerSession) Pause() (*types.Transaction, error) {
	return _EthCrossChainManager.Contract.Pause(&_EthCrossChainManager.TransactOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns(bool)
func (_EthCrossChainManager *EthCrossChainManagerTransactorSession) Pause() (*types.Transaction, error) {
	return _EthCrossChainManager.Contract.Pause(&_EthCrossChainManager.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_EthCrossChainManager *EthCrossChainManagerTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EthCrossChainManager.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_EthCrossChainManager *EthCrossChainManagerSession) RenounceOwnership() (*types.Transaction, error) {
	return _EthCrossChainManager.Contract.RenounceOwnership(&_EthCrossChainManager.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_EthCrossChainManager *EthCrossChainManagerTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _EthCrossChainManager.Contract.RenounceOwnership(&_EthCrossChainManager.TransactOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_EthCrossChainManager *EthCrossChainManagerTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _EthCrossChainManager.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_EthCrossChainManager *EthCrossChainManagerSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _EthCrossChainManager.Contract.TransferOwnership(&_EthCrossChainManager.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_EthCrossChainManager *EthCrossChainManagerTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _EthCrossChainManager.Contract.TransferOwnership(&_EthCrossChainManager.TransactOpts, newOwner)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns(bool)
func (_EthCrossChainManager *EthCrossChainManagerTransactor) Unpause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EthCrossChainManager.contract.Transact(opts, "unpause")
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns(bool)
func (_EthCrossChainManager *EthCrossChainManagerSession) Unpause() (*types.Transaction, error) {
	return _EthCrossChainManager.Contract.Unpause(&_EthCrossChainManager.TransactOpts)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns(bool)
func (_EthCrossChainManager *EthCrossChainManagerTransactorSession) Unpause() (*types.Transaction, error) {
	return _EthCrossChainManager.Contract.Unpause(&_EthCrossChainManager.TransactOpts)
}

// UpgradeToNew is a paid mutator transaction binding the contract method 0x7e724ff3.
//
// Solidity: function upgradeToNew(address newEthCrossChainManagerAddress) returns(bool)
func (_EthCrossChainManager *EthCrossChainManagerTransactor) UpgradeToNew(opts *bind.TransactOpts, newEthCrossChainManagerAddress common.Address) (*types.Transaction, error) {
	return _EthCrossChainManager.contract.Transact(opts, "upgradeToNew", newEthCrossChainManagerAddress)
}

// UpgradeToNew is a paid mutator transaction binding the contract method 0x7e724ff3.
//
// Solidity: function upgradeToNew(address newEthCrossChainManagerAddress) returns(bool)
func (_EthCrossChainManager *EthCrossChainManagerSession) UpgradeToNew(newEthCrossChainManagerAddress common.Address) (*types.Transaction, error) {
	return _EthCrossChainManager.Contract.UpgradeToNew(&_EthCrossChainManager.TransactOpts, newEthCrossChainManagerAddress)
}

// UpgradeToNew is a paid mutator transaction binding the contract method 0x7e724ff3.
//
// Solidity: function upgradeToNew(address newEthCrossChainManagerAddress) returns(bool)
func (_EthCrossChainManager *EthCrossChainManagerTransactorSession) UpgradeToNew(newEthCrossChainManagerAddress common.Address) (*types.Transaction, error) {
	return _EthCrossChainManager.Contract.UpgradeToNew(&_EthCrossChainManager.TransactOpts, newEthCrossChainManagerAddress)
}

// VerifyAndExecuteTx is a paid mutator transaction binding the contract method 0xd5ea92e3.
//
// Solidity: function verifyAndExecuteTx(bytes32[] _proof, uint8[] _position, bytes _toMerkleValueBs, uint64 _blockHeight) returns(bool)
func (_EthCrossChainManager *EthCrossChainManagerTransactor) VerifyAndExecuteTx(opts *bind.TransactOpts, _proof [][32]byte, _position []uint8, _toMerkleValueBs []byte, _blockHeight uint64) (*types.Transaction, error) {
	return _EthCrossChainManager.contract.Transact(opts, "verifyAndExecuteTx", _proof, _position, _toMerkleValueBs, _blockHeight)
}

// VerifyAndExecuteTx is a paid mutator transaction binding the contract method 0xd5ea92e3.
//
// Solidity: function verifyAndExecuteTx(bytes32[] _proof, uint8[] _position, bytes _toMerkleValueBs, uint64 _blockHeight) returns(bool)
func (_EthCrossChainManager *EthCrossChainManagerSession) VerifyAndExecuteTx(_proof [][32]byte, _position []uint8, _toMerkleValueBs []byte, _blockHeight uint64) (*types.Transaction, error) {
	return _EthCrossChainManager.Contract.VerifyAndExecuteTx(&_EthCrossChainManager.TransactOpts, _proof, _position, _toMerkleValueBs, _blockHeight)
}

// VerifyAndExecuteTx is a paid mutator transaction binding the contract method 0xd5ea92e3.
//
// Solidity: function verifyAndExecuteTx(bytes32[] _proof, uint8[] _position, bytes _toMerkleValueBs, uint64 _blockHeight) returns(bool)
func (_EthCrossChainManager *EthCrossChainManagerTransactorSession) VerifyAndExecuteTx(_proof [][32]byte, _position []uint8, _toMerkleValueBs []byte, _blockHeight uint64) (*types.Transaction, error) {
	return _EthCrossChainManager.Contract.VerifyAndExecuteTx(&_EthCrossChainManager.TransactOpts, _proof, _position, _toMerkleValueBs, _blockHeight)
}

// EthCrossChainManagerChangeBookKeeperEventIterator is returned from FilterChangeBookKeeperEvent and is used to iterate over the raw logs and unpacked data for ChangeBookKeeperEvent events raised by the EthCrossChainManager contract.
type EthCrossChainManagerChangeBookKeeperEventIterator struct {
	Event *EthCrossChainManagerChangeBookKeeperEvent // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *EthCrossChainManagerChangeBookKeeperEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EthCrossChainManagerChangeBookKeeperEvent)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(EthCrossChainManagerChangeBookKeeperEvent)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *EthCrossChainManagerChangeBookKeeperEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EthCrossChainManagerChangeBookKeeperEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EthCrossChainManagerChangeBookKeeperEvent represents a ChangeBookKeeperEvent event raised by the EthCrossChainManager contract.
type EthCrossChainManagerChangeBookKeeperEvent struct {
	Height    *big.Int
	RawHeader []byte
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterChangeBookKeeperEvent is a free log retrieval operation binding the contract event 0xe60d33488cba3977bf65766cd2f8ac9617f64bf3b3198aff6240ce5c7d43b690.
//
// Solidity: event ChangeBookKeeperEvent(uint256 height, bytes rawHeader)
func (_EthCrossChainManager *EthCrossChainManagerFilterer) FilterChangeBookKeeperEvent(opts *bind.FilterOpts) (*EthCrossChainManagerChangeBookKeeperEventIterator, error) {

	logs, sub, err := _EthCrossChainManager.contract.FilterLogs(opts, "ChangeBookKeeperEvent")
	if err != nil {
		return nil, err
	}
	return &EthCrossChainManagerChangeBookKeeperEventIterator{contract: _EthCrossChainManager.contract, event: "ChangeBookKeeperEvent", logs: logs, sub: sub}, nil
}

// WatchChangeBookKeeperEvent is a free log subscription operation binding the contract event 0xe60d33488cba3977bf65766cd2f8ac9617f64bf3b3198aff6240ce5c7d43b690.
//
// Solidity: event ChangeBookKeeperEvent(uint256 height, bytes rawHeader)
func (_EthCrossChainManager *EthCrossChainManagerFilterer) WatchChangeBookKeeperEvent(opts *bind.WatchOpts, sink chan<- *EthCrossChainManagerChangeBookKeeperEvent) (event.Subscription, error) {

	logs, sub, err := _EthCrossChainManager.contract.WatchLogs(opts, "ChangeBookKeeperEvent")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EthCrossChainManagerChangeBookKeeperEvent)
				if err := _EthCrossChainManager.contract.UnpackLog(event, "ChangeBookKeeperEvent", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseChangeBookKeeperEvent is a log parse operation binding the contract event 0xe60d33488cba3977bf65766cd2f8ac9617f64bf3b3198aff6240ce5c7d43b690.
//
// Solidity: event ChangeBookKeeperEvent(uint256 height, bytes rawHeader)
func (_EthCrossChainManager *EthCrossChainManagerFilterer) ParseChangeBookKeeperEvent(log types.Log) (*EthCrossChainManagerChangeBookKeeperEvent, error) {
	event := new(EthCrossChainManagerChangeBookKeeperEvent)
	if err := _EthCrossChainManager.contract.UnpackLog(event, "ChangeBookKeeperEvent", log); err != nil {
		return nil, err
	}
	return event, nil
}

// EthCrossChainManagerCrossChainEventIterator is returned from FilterCrossChainEvent and is used to iterate over the raw logs and unpacked data for CrossChainEvent events raised by the EthCrossChainManager contract.
type EthCrossChainManagerCrossChainEventIterator struct {
	Event *EthCrossChainManagerCrossChainEvent // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *EthCrossChainManagerCrossChainEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EthCrossChainManagerCrossChainEvent)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(EthCrossChainManagerCrossChainEvent)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *EthCrossChainManagerCrossChainEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EthCrossChainManagerCrossChainEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EthCrossChainManagerCrossChainEvent represents a CrossChainEvent event raised by the EthCrossChainManager contract.
type EthCrossChainManagerCrossChainEvent struct {
	Sender               common.Address
	TxId                 []byte
	ProxyOrAssetContract common.Address
	ToChainId            uint64
	ToContract           []byte
	Rawdata              []byte
	Raw                  types.Log // Blockchain specific contextual infos
}

// FilterCrossChainEvent is a free log retrieval operation binding the contract event 0x6ad3bf15c1988bc04bc153490cab16db8efb9a3990215bf1c64ea6e28be88483.
//
// Solidity: event CrossChainEvent(address indexed sender, bytes txId, address proxyOrAssetContract, uint64 toChainId, bytes toContract, bytes rawdata)
func (_EthCrossChainManager *EthCrossChainManagerFilterer) FilterCrossChainEvent(opts *bind.FilterOpts, sender []common.Address) (*EthCrossChainManagerCrossChainEventIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _EthCrossChainManager.contract.FilterLogs(opts, "CrossChainEvent", senderRule)
	if err != nil {
		return nil, err
	}
	return &EthCrossChainManagerCrossChainEventIterator{contract: _EthCrossChainManager.contract, event: "CrossChainEvent", logs: logs, sub: sub}, nil
}

// WatchCrossChainEvent is a free log subscription operation binding the contract event 0x6ad3bf15c1988bc04bc153490cab16db8efb9a3990215bf1c64ea6e28be88483.
//
// Solidity: event CrossChainEvent(address indexed sender, bytes txId, address proxyOrAssetContract, uint64 toChainId, bytes toContract, bytes rawdata)
func (_EthCrossChainManager *EthCrossChainManagerFilterer) WatchCrossChainEvent(opts *bind.WatchOpts, sink chan<- *EthCrossChainManagerCrossChainEvent, sender []common.Address) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _EthCrossChainManager.contract.WatchLogs(opts, "CrossChainEvent", senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EthCrossChainManagerCrossChainEvent)
				if err := _EthCrossChainManager.contract.UnpackLog(event, "CrossChainEvent", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseCrossChainEvent is a log parse operation binding the contract event 0x6ad3bf15c1988bc04bc153490cab16db8efb9a3990215bf1c64ea6e28be88483.
//
// Solidity: event CrossChainEvent(address indexed sender, bytes txId, address proxyOrAssetContract, uint64 toChainId, bytes toContract, bytes rawdata)
func (_EthCrossChainManager *EthCrossChainManagerFilterer) ParseCrossChainEvent(log types.Log) (*EthCrossChainManagerCrossChainEvent, error) {
	event := new(EthCrossChainManagerCrossChainEvent)
	if err := _EthCrossChainManager.contract.UnpackLog(event, "CrossChainEvent", log); err != nil {
		return nil, err
	}
	return event, nil
}

// EthCrossChainManagerInitGenesisBlockEventIterator is returned from FilterInitGenesisBlockEvent and is used to iterate over the raw logs and unpacked data for InitGenesisBlockEvent events raised by the EthCrossChainManager contract.
type EthCrossChainManagerInitGenesisBlockEventIterator struct {
	Event *EthCrossChainManagerInitGenesisBlockEvent // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *EthCrossChainManagerInitGenesisBlockEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EthCrossChainManagerInitGenesisBlockEvent)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(EthCrossChainManagerInitGenesisBlockEvent)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *EthCrossChainManagerInitGenesisBlockEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EthCrossChainManagerInitGenesisBlockEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EthCrossChainManagerInitGenesisBlockEvent represents a InitGenesisBlockEvent event raised by the EthCrossChainManager contract.
type EthCrossChainManagerInitGenesisBlockEvent struct {
	Height    *big.Int
	RawHeader []byte
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterInitGenesisBlockEvent is a free log retrieval operation binding the contract event 0xf01968fc3a2655cf1b5144cb32de6dc898f91b9239c103744e8457152ab2fbde.
//
// Solidity: event InitGenesisBlockEvent(uint256 height, bytes rawHeader)
func (_EthCrossChainManager *EthCrossChainManagerFilterer) FilterInitGenesisBlockEvent(opts *bind.FilterOpts) (*EthCrossChainManagerInitGenesisBlockEventIterator, error) {

	logs, sub, err := _EthCrossChainManager.contract.FilterLogs(opts, "InitGenesisBlockEvent")
	if err != nil {
		return nil, err
	}
	return &EthCrossChainManagerInitGenesisBlockEventIterator{contract: _EthCrossChainManager.contract, event: "InitGenesisBlockEvent", logs: logs, sub: sub}, nil
}

// WatchInitGenesisBlockEvent is a free log subscription operation binding the contract event 0xf01968fc3a2655cf1b5144cb32de6dc898f91b9239c103744e8457152ab2fbde.
//
// Solidity: event InitGenesisBlockEvent(uint256 height, bytes rawHeader)
func (_EthCrossChainManager *EthCrossChainManagerFilterer) WatchInitGenesisBlockEvent(opts *bind.WatchOpts, sink chan<- *EthCrossChainManagerInitGenesisBlockEvent) (event.Subscription, error) {

	logs, sub, err := _EthCrossChainManager.contract.WatchLogs(opts, "InitGenesisBlockEvent")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EthCrossChainManagerInitGenesisBlockEvent)
				if err := _EthCrossChainManager.contract.UnpackLog(event, "InitGenesisBlockEvent", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseInitGenesisBlockEvent is a log parse operation binding the contract event 0xf01968fc3a2655cf1b5144cb32de6dc898f91b9239c103744e8457152ab2fbde.
//
// Solidity: event InitGenesisBlockEvent(uint256 height, bytes rawHeader)
func (_EthCrossChainManager *EthCrossChainManagerFilterer) ParseInitGenesisBlockEvent(log types.Log) (*EthCrossChainManagerInitGenesisBlockEvent, error) {
	event := new(EthCrossChainManagerInitGenesisBlockEvent)
	if err := _EthCrossChainManager.contract.UnpackLog(event, "InitGenesisBlockEvent", log); err != nil {
		return nil, err
	}
	return event, nil
}

// EthCrossChainManagerOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the EthCrossChainManager contract.
type EthCrossChainManagerOwnershipTransferredIterator struct {
	Event *EthCrossChainManagerOwnershipTransferred // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *EthCrossChainManagerOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EthCrossChainManagerOwnershipTransferred)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(EthCrossChainManagerOwnershipTransferred)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *EthCrossChainManagerOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EthCrossChainManagerOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EthCrossChainManagerOwnershipTransferred represents a OwnershipTransferred event raised by the EthCrossChainManager contract.
type EthCrossChainManagerOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_EthCrossChainManager *EthCrossChainManagerFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*EthCrossChainManagerOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _EthCrossChainManager.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &EthCrossChainManagerOwnershipTransferredIterator{contract: _EthCrossChainManager.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_EthCrossChainManager *EthCrossChainManagerFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *EthCrossChainManagerOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _EthCrossChainManager.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EthCrossChainManagerOwnershipTransferred)
				if err := _EthCrossChainManager.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseOwnershipTransferred is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_EthCrossChainManager *EthCrossChainManagerFilterer) ParseOwnershipTransferred(log types.Log) (*EthCrossChainManagerOwnershipTransferred, error) {
	event := new(EthCrossChainManagerOwnershipTransferred)
	if err := _EthCrossChainManager.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	return event, nil
}

// EthCrossChainManagerPausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the EthCrossChainManager contract.
type EthCrossChainManagerPausedIterator struct {
	Event *EthCrossChainManagerPaused // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *EthCrossChainManagerPausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EthCrossChainManagerPaused)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(EthCrossChainManagerPaused)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *EthCrossChainManagerPausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EthCrossChainManagerPausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EthCrossChainManagerPaused represents a Paused event raised by the EthCrossChainManager contract.
type EthCrossChainManagerPaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_EthCrossChainManager *EthCrossChainManagerFilterer) FilterPaused(opts *bind.FilterOpts) (*EthCrossChainManagerPausedIterator, error) {

	logs, sub, err := _EthCrossChainManager.contract.FilterLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return &EthCrossChainManagerPausedIterator{contract: _EthCrossChainManager.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_EthCrossChainManager *EthCrossChainManagerFilterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *EthCrossChainManagerPaused) (event.Subscription, error) {

	logs, sub, err := _EthCrossChainManager.contract.WatchLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EthCrossChainManagerPaused)
				if err := _EthCrossChainManager.contract.UnpackLog(event, "Paused", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParsePaused is a log parse operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_EthCrossChainManager *EthCrossChainManagerFilterer) ParsePaused(log types.Log) (*EthCrossChainManagerPaused, error) {
	event := new(EthCrossChainManagerPaused)
	if err := _EthCrossChainManager.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	return event, nil
}

// EthCrossChainManagerSyncBlockHeaderEventIterator is returned from FilterSyncBlockHeaderEvent and is used to iterate over the raw logs and unpacked data for SyncBlockHeaderEvent events raised by the EthCrossChainManager contract.
type EthCrossChainManagerSyncBlockHeaderEventIterator struct {
	Event *EthCrossChainManagerSyncBlockHeaderEvent // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *EthCrossChainManagerSyncBlockHeaderEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EthCrossChainManagerSyncBlockHeaderEvent)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(EthCrossChainManagerSyncBlockHeaderEvent)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *EthCrossChainManagerSyncBlockHeaderEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EthCrossChainManagerSyncBlockHeaderEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EthCrossChainManagerSyncBlockHeaderEvent represents a SyncBlockHeaderEvent event raised by the EthCrossChainManager contract.
type EthCrossChainManagerSyncBlockHeaderEvent struct {
	Height    *big.Int
	RawHeader []byte
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterSyncBlockHeaderEvent is a free log retrieval operation binding the contract event 0x1645cc3eff4fb98fef2ffd7806e7f39bfe18558f6656e69dfc9c8f26e958d4e7.
//
// Solidity: event SyncBlockHeaderEvent(uint256 height, bytes rawHeader)
func (_EthCrossChainManager *EthCrossChainManagerFilterer) FilterSyncBlockHeaderEvent(opts *bind.FilterOpts) (*EthCrossChainManagerSyncBlockHeaderEventIterator, error) {

	logs, sub, err := _EthCrossChainManager.contract.FilterLogs(opts, "SyncBlockHeaderEvent")
	if err != nil {
		return nil, err
	}
	return &EthCrossChainManagerSyncBlockHeaderEventIterator{contract: _EthCrossChainManager.contract, event: "SyncBlockHeaderEvent", logs: logs, sub: sub}, nil
}

// WatchSyncBlockHeaderEvent is a free log subscription operation binding the contract event 0x1645cc3eff4fb98fef2ffd7806e7f39bfe18558f6656e69dfc9c8f26e958d4e7.
//
// Solidity: event SyncBlockHeaderEvent(uint256 height, bytes rawHeader)
func (_EthCrossChainManager *EthCrossChainManagerFilterer) WatchSyncBlockHeaderEvent(opts *bind.WatchOpts, sink chan<- *EthCrossChainManagerSyncBlockHeaderEvent) (event.Subscription, error) {

	logs, sub, err := _EthCrossChainManager.contract.WatchLogs(opts, "SyncBlockHeaderEvent")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EthCrossChainManagerSyncBlockHeaderEvent)
				if err := _EthCrossChainManager.contract.UnpackLog(event, "SyncBlockHeaderEvent", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseSyncBlockHeaderEvent is a log parse operation binding the contract event 0x1645cc3eff4fb98fef2ffd7806e7f39bfe18558f6656e69dfc9c8f26e958d4e7.
//
// Solidity: event SyncBlockHeaderEvent(uint256 height, bytes rawHeader)
func (_EthCrossChainManager *EthCrossChainManagerFilterer) ParseSyncBlockHeaderEvent(log types.Log) (*EthCrossChainManagerSyncBlockHeaderEvent, error) {
	event := new(EthCrossChainManagerSyncBlockHeaderEvent)
	if err := _EthCrossChainManager.contract.UnpackLog(event, "SyncBlockHeaderEvent", log); err != nil {
		return nil, err
	}
	return event, nil
}

// EthCrossChainManagerUnpausedIterator is returned from FilterUnpaused and is used to iterate over the raw logs and unpacked data for Unpaused events raised by the EthCrossChainManager contract.
type EthCrossChainManagerUnpausedIterator struct {
	Event *EthCrossChainManagerUnpaused // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *EthCrossChainManagerUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EthCrossChainManagerUnpaused)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(EthCrossChainManagerUnpaused)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *EthCrossChainManagerUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EthCrossChainManagerUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EthCrossChainManagerUnpaused represents a Unpaused event raised by the EthCrossChainManager contract.
type EthCrossChainManagerUnpaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterUnpaused is a free log retrieval operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_EthCrossChainManager *EthCrossChainManagerFilterer) FilterUnpaused(opts *bind.FilterOpts) (*EthCrossChainManagerUnpausedIterator, error) {

	logs, sub, err := _EthCrossChainManager.contract.FilterLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return &EthCrossChainManagerUnpausedIterator{contract: _EthCrossChainManager.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

// WatchUnpaused is a free log subscription operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_EthCrossChainManager *EthCrossChainManagerFilterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *EthCrossChainManagerUnpaused) (event.Subscription, error) {

	logs, sub, err := _EthCrossChainManager.contract.WatchLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EthCrossChainManagerUnpaused)
				if err := _EthCrossChainManager.contract.UnpackLog(event, "Unpaused", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseUnpaused is a log parse operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_EthCrossChainManager *EthCrossChainManagerFilterer) ParseUnpaused(log types.Log) (*EthCrossChainManagerUnpaused, error) {
	event := new(EthCrossChainManagerUnpaused)
	if err := _EthCrossChainManager.contract.UnpackLog(event, "Unpaused", log); err != nil {
		return nil, err
	}
	return event, nil
}

// EthCrossChainManagerVerifyAndExecuteTxEventIterator is returned from FilterVerifyAndExecuteTxEvent and is used to iterate over the raw logs and unpacked data for VerifyAndExecuteTxEvent events raised by the EthCrossChainManager contract.
type EthCrossChainManagerVerifyAndExecuteTxEventIterator struct {
	Event *EthCrossChainManagerVerifyAndExecuteTxEvent // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *EthCrossChainManagerVerifyAndExecuteTxEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EthCrossChainManagerVerifyAndExecuteTxEvent)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(EthCrossChainManagerVerifyAndExecuteTxEvent)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *EthCrossChainManagerVerifyAndExecuteTxEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EthCrossChainManagerVerifyAndExecuteTxEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EthCrossChainManagerVerifyAndExecuteTxEvent represents a VerifyAndExecuteTxEvent event raised by the EthCrossChainManager contract.
type EthCrossChainManagerVerifyAndExecuteTxEvent struct {
	FromChainID      uint64
	ToContract       []byte
	CrossChainTxHash []byte
	FromChainTxHash  []byte
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterVerifyAndExecuteTxEvent is a free log retrieval operation binding the contract event 0x5944a0908ae763fa254df8e0da54a682cbc62115501336e18be52a7dd774b80b.
//
// Solidity: event VerifyAndExecuteTxEvent(uint64 fromChainID, bytes toContract, bytes crossChainTxHash, bytes fromChainTxHash)
func (_EthCrossChainManager *EthCrossChainManagerFilterer) FilterVerifyAndExecuteTxEvent(opts *bind.FilterOpts) (*EthCrossChainManagerVerifyAndExecuteTxEventIterator, error) {

	logs, sub, err := _EthCrossChainManager.contract.FilterLogs(opts, "VerifyAndExecuteTxEvent")
	if err != nil {
		return nil, err
	}
	return &EthCrossChainManagerVerifyAndExecuteTxEventIterator{contract: _EthCrossChainManager.contract, event: "VerifyAndExecuteTxEvent", logs: logs, sub: sub}, nil
}

// WatchVerifyAndExecuteTxEvent is a free log subscription operation binding the contract event 0x5944a0908ae763fa254df8e0da54a682cbc62115501336e18be52a7dd774b80b.
//
// Solidity: event VerifyAndExecuteTxEvent(uint64 fromChainID, bytes toContract, bytes crossChainTxHash, bytes fromChainTxHash)
func (_EthCrossChainManager *EthCrossChainManagerFilterer) WatchVerifyAndExecuteTxEvent(opts *bind.WatchOpts, sink chan<- *EthCrossChainManagerVerifyAndExecuteTxEvent) (event.Subscription, error) {

	logs, sub, err := _EthCrossChainManager.contract.WatchLogs(opts, "VerifyAndExecuteTxEvent")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EthCrossChainManagerVerifyAndExecuteTxEvent)
				if err := _EthCrossChainManager.contract.UnpackLog(event, "VerifyAndExecuteTxEvent", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseVerifyAndExecuteTxEvent is a log parse operation binding the contract event 0x5944a0908ae763fa254df8e0da54a682cbc62115501336e18be52a7dd774b80b.
//
// Solidity: event VerifyAndExecuteTxEvent(uint64 fromChainID, bytes toContract, bytes crossChainTxHash, bytes fromChainTxHash)
func (_EthCrossChainManager *EthCrossChainManagerFilterer) ParseVerifyAndExecuteTxEvent(log types.Log) (*EthCrossChainManagerVerifyAndExecuteTxEvent, error) {
	event := new(EthCrossChainManagerVerifyAndExecuteTxEvent)
	if err := _EthCrossChainManager.contract.UnpackLog(event, "VerifyAndExecuteTxEvent", log); err != nil {
		return nil, err
	}
	return event, nil
}

// IEthCrossChainDataABI is the input ABI used to generate the binding from.
const IEthCrossChainDataABI = "[{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_crossChainTx\",\"type\":\"bytes32\"}],\"name\":\"getCrossChainTxExist\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_ethTxHashIndex\",\"type\":\"uint256\"}],\"name\":\"getEthTxHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getEthTxHashIndex\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"key1\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"key2\",\"type\":\"bytes32\"}],\"name\":\"getExtraData\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getLatestBookKeeperHeight\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getLatestHeight\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"_height\",\"type\":\"uint64\"}],\"name\":\"getMCHeaderBytes\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getMCKeeperHeight\",\"outputs\":[{\"internalType\":\"uint64[]\",\"name\":\"\",\"type\":\"uint64[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"height\",\"type\":\"uint64\"}],\"name\":\"getMCKeeperPubBytes\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"_height\",\"type\":\"uint64\"}],\"name\":\"getMCKeeperPubKeybytes\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"initGenesisBlock\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"isGenesisBlockInited\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"pause\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_crossChainTx\",\"type\":\"bytes32\"}],\"name\":\"putCrossChainTxExist\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_ethTxHashIndex\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"_ethTxHash\",\"type\":\"bytes32\"}],\"name\":\"putEthTxHash\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"key1\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"key2\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"value\",\"type\":\"bytes\"}],\"name\":\"putExtraData\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"_latestBookKeeperHeight\",\"type\":\"uint64\"}],\"name\":\"putLatestBookKeeperHeight\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"_latestHeight\",\"type\":\"uint64\"}],\"name\":\"putLatestHeight\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"_height\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"_rawHeader\",\"type\":\"bytes\"}],\"name\":\"putMCHeaderBytes\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"height\",\"type\":\"uint64\"}],\"name\":\"putMCKeeperHeight\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"height\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"_MCKeeperPubBytes\",\"type\":\"bytes\"}],\"name\":\"putMCKeeperPubBytes\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"_height\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"_keepersBytes\",\"type\":\"bytes\"}],\"name\":\"putMCKeeperPubKeybytes\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"unpause\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// IEthCrossChainDataFuncSigs maps the 4-byte function signature to its string representation.
var IEthCrossChainDataFuncSigs = map[string]string{
	"cceb218d": "getCrossChainTxExist(bytes32)",
	"29927875": "getEthTxHash(uint256)",
	"ff3d24a1": "getEthTxHashIndex()",
	"40602bb5": "getExtraData(bytes32,bytes32)",
	"71a77750": "getLatestBookKeeperHeight()",
	"4ed1d8cc": "getLatestHeight()",
	"9743e9d6": "getMCHeaderBytes(uint64)",
	"f0b1fcfe": "getMCKeeperHeight()",
	"7d737b8b": "getMCKeeperPubBytes(uint64)",
	"09705412": "getMCKeeperPubKeybytes(uint64)",
	"8b39aa25": "initGenesisBlock()",
	"64f1acb0": "isGenesisBlockInited()",
	"8456cb59": "pause()",
	"5c975abb": "paused()",
	"794b0300": "putCrossChainTxExist(bytes32)",
	"9c59af4f": "putEthTxHash(uint256,bytes32)",
	"1afe374e": "putExtraData(bytes32,bytes32,bytes)",
	"a6c6fe7c": "putLatestBookKeeperHeight(uint64)",
	"6d044440": "putLatestHeight(uint64)",
	"021d70ab": "putMCHeaderBytes(uint64,bytes)",
	"c6dbd0b7": "putMCKeeperHeight(uint64)",
	"191b0ab1": "putMCKeeperPubBytes(uint64,bytes)",
	"73c2f5f6": "putMCKeeperPubKeybytes(uint64,bytes)",
	"f2fde38b": "transferOwnership(address)",
	"3f4ba83a": "unpause()",
}

// IEthCrossChainData is an auto generated Go binding around an Ethereum contract.
type IEthCrossChainData struct {
	IEthCrossChainDataCaller     // Read-only binding to the contract
	IEthCrossChainDataTransactor // Write-only binding to the contract
	IEthCrossChainDataFilterer   // Log filterer for contract events
}

// IEthCrossChainDataCaller is an auto generated read-only Go binding around an Ethereum contract.
type IEthCrossChainDataCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IEthCrossChainDataTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IEthCrossChainDataTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IEthCrossChainDataFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IEthCrossChainDataFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IEthCrossChainDataSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IEthCrossChainDataSession struct {
	Contract     *IEthCrossChainData // Generic contract binding to set the session for
	CallOpts     bind.CallOpts       // Call options to use throughout this session
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// IEthCrossChainDataCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IEthCrossChainDataCallerSession struct {
	Contract *IEthCrossChainDataCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts             // Call options to use throughout this session
}

// IEthCrossChainDataTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IEthCrossChainDataTransactorSession struct {
	Contract     *IEthCrossChainDataTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts             // Transaction auth options to use throughout this session
}

// IEthCrossChainDataRaw is an auto generated low-level Go binding around an Ethereum contract.
type IEthCrossChainDataRaw struct {
	Contract *IEthCrossChainData // Generic contract binding to access the raw methods on
}

// IEthCrossChainDataCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IEthCrossChainDataCallerRaw struct {
	Contract *IEthCrossChainDataCaller // Generic read-only contract binding to access the raw methods on
}

// IEthCrossChainDataTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IEthCrossChainDataTransactorRaw struct {
	Contract *IEthCrossChainDataTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIEthCrossChainData creates a new instance of IEthCrossChainData, bound to a specific deployed contract.
func NewIEthCrossChainData(address common.Address, backend bind.ContractBackend) (*IEthCrossChainData, error) {
	contract, err := bindIEthCrossChainData(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IEthCrossChainData{IEthCrossChainDataCaller: IEthCrossChainDataCaller{contract: contract}, IEthCrossChainDataTransactor: IEthCrossChainDataTransactor{contract: contract}, IEthCrossChainDataFilterer: IEthCrossChainDataFilterer{contract: contract}}, nil
}

// NewIEthCrossChainDataCaller creates a new read-only instance of IEthCrossChainData, bound to a specific deployed contract.
func NewIEthCrossChainDataCaller(address common.Address, caller bind.ContractCaller) (*IEthCrossChainDataCaller, error) {
	contract, err := bindIEthCrossChainData(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IEthCrossChainDataCaller{contract: contract}, nil
}

// NewIEthCrossChainDataTransactor creates a new write-only instance of IEthCrossChainData, bound to a specific deployed contract.
func NewIEthCrossChainDataTransactor(address common.Address, transactor bind.ContractTransactor) (*IEthCrossChainDataTransactor, error) {
	contract, err := bindIEthCrossChainData(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IEthCrossChainDataTransactor{contract: contract}, nil
}

// NewIEthCrossChainDataFilterer creates a new log filterer instance of IEthCrossChainData, bound to a specific deployed contract.
func NewIEthCrossChainDataFilterer(address common.Address, filterer bind.ContractFilterer) (*IEthCrossChainDataFilterer, error) {
	contract, err := bindIEthCrossChainData(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IEthCrossChainDataFilterer{contract: contract}, nil
}

// bindIEthCrossChainData binds a generic wrapper to an already deployed contract.
func bindIEthCrossChainData(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IEthCrossChainDataABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IEthCrossChainData *IEthCrossChainDataRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _IEthCrossChainData.Contract.IEthCrossChainDataCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IEthCrossChainData *IEthCrossChainDataRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IEthCrossChainData.Contract.IEthCrossChainDataTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IEthCrossChainData *IEthCrossChainDataRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IEthCrossChainData.Contract.IEthCrossChainDataTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IEthCrossChainData *IEthCrossChainDataCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _IEthCrossChainData.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IEthCrossChainData *IEthCrossChainDataTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IEthCrossChainData.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IEthCrossChainData *IEthCrossChainDataTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IEthCrossChainData.Contract.contract.Transact(opts, method, params...)
}

// GetCrossChainTxExist is a free data retrieval call binding the contract method 0xcceb218d.
//
// Solidity: function getCrossChainTxExist(bytes32 _crossChainTx) constant returns(bool)
func (_IEthCrossChainData *IEthCrossChainDataCaller) GetCrossChainTxExist(opts *bind.CallOpts, _crossChainTx [32]byte) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _IEthCrossChainData.contract.Call(opts, out, "getCrossChainTxExist", _crossChainTx)
	return *ret0, err
}

// GetCrossChainTxExist is a free data retrieval call binding the contract method 0xcceb218d.
//
// Solidity: function getCrossChainTxExist(bytes32 _crossChainTx) constant returns(bool)
func (_IEthCrossChainData *IEthCrossChainDataSession) GetCrossChainTxExist(_crossChainTx [32]byte) (bool, error) {
	return _IEthCrossChainData.Contract.GetCrossChainTxExist(&_IEthCrossChainData.CallOpts, _crossChainTx)
}

// GetCrossChainTxExist is a free data retrieval call binding the contract method 0xcceb218d.
//
// Solidity: function getCrossChainTxExist(bytes32 _crossChainTx) constant returns(bool)
func (_IEthCrossChainData *IEthCrossChainDataCallerSession) GetCrossChainTxExist(_crossChainTx [32]byte) (bool, error) {
	return _IEthCrossChainData.Contract.GetCrossChainTxExist(&_IEthCrossChainData.CallOpts, _crossChainTx)
}

// GetEthTxHash is a free data retrieval call binding the contract method 0x29927875.
//
// Solidity: function getEthTxHash(uint256 _ethTxHashIndex) constant returns(bytes32)
func (_IEthCrossChainData *IEthCrossChainDataCaller) GetEthTxHash(opts *bind.CallOpts, _ethTxHashIndex *big.Int) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _IEthCrossChainData.contract.Call(opts, out, "getEthTxHash", _ethTxHashIndex)
	return *ret0, err
}

// GetEthTxHash is a free data retrieval call binding the contract method 0x29927875.
//
// Solidity: function getEthTxHash(uint256 _ethTxHashIndex) constant returns(bytes32)
func (_IEthCrossChainData *IEthCrossChainDataSession) GetEthTxHash(_ethTxHashIndex *big.Int) ([32]byte, error) {
	return _IEthCrossChainData.Contract.GetEthTxHash(&_IEthCrossChainData.CallOpts, _ethTxHashIndex)
}

// GetEthTxHash is a free data retrieval call binding the contract method 0x29927875.
//
// Solidity: function getEthTxHash(uint256 _ethTxHashIndex) constant returns(bytes32)
func (_IEthCrossChainData *IEthCrossChainDataCallerSession) GetEthTxHash(_ethTxHashIndex *big.Int) ([32]byte, error) {
	return _IEthCrossChainData.Contract.GetEthTxHash(&_IEthCrossChainData.CallOpts, _ethTxHashIndex)
}

// GetEthTxHashIndex is a free data retrieval call binding the contract method 0xff3d24a1.
//
// Solidity: function getEthTxHashIndex() constant returns(uint256)
func (_IEthCrossChainData *IEthCrossChainDataCaller) GetEthTxHashIndex(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _IEthCrossChainData.contract.Call(opts, out, "getEthTxHashIndex")
	return *ret0, err
}

// GetEthTxHashIndex is a free data retrieval call binding the contract method 0xff3d24a1.
//
// Solidity: function getEthTxHashIndex() constant returns(uint256)
func (_IEthCrossChainData *IEthCrossChainDataSession) GetEthTxHashIndex() (*big.Int, error) {
	return _IEthCrossChainData.Contract.GetEthTxHashIndex(&_IEthCrossChainData.CallOpts)
}

// GetEthTxHashIndex is a free data retrieval call binding the contract method 0xff3d24a1.
//
// Solidity: function getEthTxHashIndex() constant returns(uint256)
func (_IEthCrossChainData *IEthCrossChainDataCallerSession) GetEthTxHashIndex() (*big.Int, error) {
	return _IEthCrossChainData.Contract.GetEthTxHashIndex(&_IEthCrossChainData.CallOpts)
}

// GetExtraData is a free data retrieval call binding the contract method 0x40602bb5.
//
// Solidity: function getExtraData(bytes32 key1, bytes32 key2) constant returns(bytes)
func (_IEthCrossChainData *IEthCrossChainDataCaller) GetExtraData(opts *bind.CallOpts, key1 [32]byte, key2 [32]byte) ([]byte, error) {
	var (
		ret0 = new([]byte)
	)
	out := ret0
	err := _IEthCrossChainData.contract.Call(opts, out, "getExtraData", key1, key2)
	return *ret0, err
}

// GetExtraData is a free data retrieval call binding the contract method 0x40602bb5.
//
// Solidity: function getExtraData(bytes32 key1, bytes32 key2) constant returns(bytes)
func (_IEthCrossChainData *IEthCrossChainDataSession) GetExtraData(key1 [32]byte, key2 [32]byte) ([]byte, error) {
	return _IEthCrossChainData.Contract.GetExtraData(&_IEthCrossChainData.CallOpts, key1, key2)
}

// GetExtraData is a free data retrieval call binding the contract method 0x40602bb5.
//
// Solidity: function getExtraData(bytes32 key1, bytes32 key2) constant returns(bytes)
func (_IEthCrossChainData *IEthCrossChainDataCallerSession) GetExtraData(key1 [32]byte, key2 [32]byte) ([]byte, error) {
	return _IEthCrossChainData.Contract.GetExtraData(&_IEthCrossChainData.CallOpts, key1, key2)
}

// GetLatestBookKeeperHeight is a free data retrieval call binding the contract method 0x71a77750.
//
// Solidity: function getLatestBookKeeperHeight() constant returns(uint64)
func (_IEthCrossChainData *IEthCrossChainDataCaller) GetLatestBookKeeperHeight(opts *bind.CallOpts) (uint64, error) {
	var (
		ret0 = new(uint64)
	)
	out := ret0
	err := _IEthCrossChainData.contract.Call(opts, out, "getLatestBookKeeperHeight")
	return *ret0, err
}

// GetLatestBookKeeperHeight is a free data retrieval call binding the contract method 0x71a77750.
//
// Solidity: function getLatestBookKeeperHeight() constant returns(uint64)
func (_IEthCrossChainData *IEthCrossChainDataSession) GetLatestBookKeeperHeight() (uint64, error) {
	return _IEthCrossChainData.Contract.GetLatestBookKeeperHeight(&_IEthCrossChainData.CallOpts)
}

// GetLatestBookKeeperHeight is a free data retrieval call binding the contract method 0x71a77750.
//
// Solidity: function getLatestBookKeeperHeight() constant returns(uint64)
func (_IEthCrossChainData *IEthCrossChainDataCallerSession) GetLatestBookKeeperHeight() (uint64, error) {
	return _IEthCrossChainData.Contract.GetLatestBookKeeperHeight(&_IEthCrossChainData.CallOpts)
}

// GetLatestHeight is a free data retrieval call binding the contract method 0x4ed1d8cc.
//
// Solidity: function getLatestHeight() constant returns(uint64)
func (_IEthCrossChainData *IEthCrossChainDataCaller) GetLatestHeight(opts *bind.CallOpts) (uint64, error) {
	var (
		ret0 = new(uint64)
	)
	out := ret0
	err := _IEthCrossChainData.contract.Call(opts, out, "getLatestHeight")
	return *ret0, err
}

// GetLatestHeight is a free data retrieval call binding the contract method 0x4ed1d8cc.
//
// Solidity: function getLatestHeight() constant returns(uint64)
func (_IEthCrossChainData *IEthCrossChainDataSession) GetLatestHeight() (uint64, error) {
	return _IEthCrossChainData.Contract.GetLatestHeight(&_IEthCrossChainData.CallOpts)
}

// GetLatestHeight is a free data retrieval call binding the contract method 0x4ed1d8cc.
//
// Solidity: function getLatestHeight() constant returns(uint64)
func (_IEthCrossChainData *IEthCrossChainDataCallerSession) GetLatestHeight() (uint64, error) {
	return _IEthCrossChainData.Contract.GetLatestHeight(&_IEthCrossChainData.CallOpts)
}

// GetMCHeaderBytes is a free data retrieval call binding the contract method 0x9743e9d6.
//
// Solidity: function getMCHeaderBytes(uint64 _height) constant returns(bytes)
func (_IEthCrossChainData *IEthCrossChainDataCaller) GetMCHeaderBytes(opts *bind.CallOpts, _height uint64) ([]byte, error) {
	var (
		ret0 = new([]byte)
	)
	out := ret0
	err := _IEthCrossChainData.contract.Call(opts, out, "getMCHeaderBytes", _height)
	return *ret0, err
}

// GetMCHeaderBytes is a free data retrieval call binding the contract method 0x9743e9d6.
//
// Solidity: function getMCHeaderBytes(uint64 _height) constant returns(bytes)
func (_IEthCrossChainData *IEthCrossChainDataSession) GetMCHeaderBytes(_height uint64) ([]byte, error) {
	return _IEthCrossChainData.Contract.GetMCHeaderBytes(&_IEthCrossChainData.CallOpts, _height)
}

// GetMCHeaderBytes is a free data retrieval call binding the contract method 0x9743e9d6.
//
// Solidity: function getMCHeaderBytes(uint64 _height) constant returns(bytes)
func (_IEthCrossChainData *IEthCrossChainDataCallerSession) GetMCHeaderBytes(_height uint64) ([]byte, error) {
	return _IEthCrossChainData.Contract.GetMCHeaderBytes(&_IEthCrossChainData.CallOpts, _height)
}

// GetMCKeeperHeight is a free data retrieval call binding the contract method 0xf0b1fcfe.
//
// Solidity: function getMCKeeperHeight() constant returns(uint64[])
func (_IEthCrossChainData *IEthCrossChainDataCaller) GetMCKeeperHeight(opts *bind.CallOpts) ([]uint64, error) {
	var (
		ret0 = new([]uint64)
	)
	out := ret0
	err := _IEthCrossChainData.contract.Call(opts, out, "getMCKeeperHeight")
	return *ret0, err
}

// GetMCKeeperHeight is a free data retrieval call binding the contract method 0xf0b1fcfe.
//
// Solidity: function getMCKeeperHeight() constant returns(uint64[])
func (_IEthCrossChainData *IEthCrossChainDataSession) GetMCKeeperHeight() ([]uint64, error) {
	return _IEthCrossChainData.Contract.GetMCKeeperHeight(&_IEthCrossChainData.CallOpts)
}

// GetMCKeeperHeight is a free data retrieval call binding the contract method 0xf0b1fcfe.
//
// Solidity: function getMCKeeperHeight() constant returns(uint64[])
func (_IEthCrossChainData *IEthCrossChainDataCallerSession) GetMCKeeperHeight() ([]uint64, error) {
	return _IEthCrossChainData.Contract.GetMCKeeperHeight(&_IEthCrossChainData.CallOpts)
}

// GetMCKeeperPubBytes is a free data retrieval call binding the contract method 0x7d737b8b.
//
// Solidity: function getMCKeeperPubBytes(uint64 height) constant returns(bytes)
func (_IEthCrossChainData *IEthCrossChainDataCaller) GetMCKeeperPubBytes(opts *bind.CallOpts, height uint64) ([]byte, error) {
	var (
		ret0 = new([]byte)
	)
	out := ret0
	err := _IEthCrossChainData.contract.Call(opts, out, "getMCKeeperPubBytes", height)
	return *ret0, err
}

// GetMCKeeperPubBytes is a free data retrieval call binding the contract method 0x7d737b8b.
//
// Solidity: function getMCKeeperPubBytes(uint64 height) constant returns(bytes)
func (_IEthCrossChainData *IEthCrossChainDataSession) GetMCKeeperPubBytes(height uint64) ([]byte, error) {
	return _IEthCrossChainData.Contract.GetMCKeeperPubBytes(&_IEthCrossChainData.CallOpts, height)
}

// GetMCKeeperPubBytes is a free data retrieval call binding the contract method 0x7d737b8b.
//
// Solidity: function getMCKeeperPubBytes(uint64 height) constant returns(bytes)
func (_IEthCrossChainData *IEthCrossChainDataCallerSession) GetMCKeeperPubBytes(height uint64) ([]byte, error) {
	return _IEthCrossChainData.Contract.GetMCKeeperPubBytes(&_IEthCrossChainData.CallOpts, height)
}

// GetMCKeeperPubKeybytes is a free data retrieval call binding the contract method 0x09705412.
//
// Solidity: function getMCKeeperPubKeybytes(uint64 _height) constant returns(bytes)
func (_IEthCrossChainData *IEthCrossChainDataCaller) GetMCKeeperPubKeybytes(opts *bind.CallOpts, _height uint64) ([]byte, error) {
	var (
		ret0 = new([]byte)
	)
	out := ret0
	err := _IEthCrossChainData.contract.Call(opts, out, "getMCKeeperPubKeybytes", _height)
	return *ret0, err
}

// GetMCKeeperPubKeybytes is a free data retrieval call binding the contract method 0x09705412.
//
// Solidity: function getMCKeeperPubKeybytes(uint64 _height) constant returns(bytes)
func (_IEthCrossChainData *IEthCrossChainDataSession) GetMCKeeperPubKeybytes(_height uint64) ([]byte, error) {
	return _IEthCrossChainData.Contract.GetMCKeeperPubKeybytes(&_IEthCrossChainData.CallOpts, _height)
}

// GetMCKeeperPubKeybytes is a free data retrieval call binding the contract method 0x09705412.
//
// Solidity: function getMCKeeperPubKeybytes(uint64 _height) constant returns(bytes)
func (_IEthCrossChainData *IEthCrossChainDataCallerSession) GetMCKeeperPubKeybytes(_height uint64) ([]byte, error) {
	return _IEthCrossChainData.Contract.GetMCKeeperPubKeybytes(&_IEthCrossChainData.CallOpts, _height)
}

// IsGenesisBlockInited is a free data retrieval call binding the contract method 0x64f1acb0.
//
// Solidity: function isGenesisBlockInited() constant returns(bool)
func (_IEthCrossChainData *IEthCrossChainDataCaller) IsGenesisBlockInited(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _IEthCrossChainData.contract.Call(opts, out, "isGenesisBlockInited")
	return *ret0, err
}

// IsGenesisBlockInited is a free data retrieval call binding the contract method 0x64f1acb0.
//
// Solidity: function isGenesisBlockInited() constant returns(bool)
func (_IEthCrossChainData *IEthCrossChainDataSession) IsGenesisBlockInited() (bool, error) {
	return _IEthCrossChainData.Contract.IsGenesisBlockInited(&_IEthCrossChainData.CallOpts)
}

// IsGenesisBlockInited is a free data retrieval call binding the contract method 0x64f1acb0.
//
// Solidity: function isGenesisBlockInited() constant returns(bool)
func (_IEthCrossChainData *IEthCrossChainDataCallerSession) IsGenesisBlockInited() (bool, error) {
	return _IEthCrossChainData.Contract.IsGenesisBlockInited(&_IEthCrossChainData.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() constant returns(bool)
func (_IEthCrossChainData *IEthCrossChainDataCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _IEthCrossChainData.contract.Call(opts, out, "paused")
	return *ret0, err
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() constant returns(bool)
func (_IEthCrossChainData *IEthCrossChainDataSession) Paused() (bool, error) {
	return _IEthCrossChainData.Contract.Paused(&_IEthCrossChainData.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() constant returns(bool)
func (_IEthCrossChainData *IEthCrossChainDataCallerSession) Paused() (bool, error) {
	return _IEthCrossChainData.Contract.Paused(&_IEthCrossChainData.CallOpts)
}

// InitGenesisBlock is a paid mutator transaction binding the contract method 0x8b39aa25.
//
// Solidity: function initGenesisBlock() returns(bool)
func (_IEthCrossChainData *IEthCrossChainDataTransactor) InitGenesisBlock(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IEthCrossChainData.contract.Transact(opts, "initGenesisBlock")
}

// InitGenesisBlock is a paid mutator transaction binding the contract method 0x8b39aa25.
//
// Solidity: function initGenesisBlock() returns(bool)
func (_IEthCrossChainData *IEthCrossChainDataSession) InitGenesisBlock() (*types.Transaction, error) {
	return _IEthCrossChainData.Contract.InitGenesisBlock(&_IEthCrossChainData.TransactOpts)
}

// InitGenesisBlock is a paid mutator transaction binding the contract method 0x8b39aa25.
//
// Solidity: function initGenesisBlock() returns(bool)
func (_IEthCrossChainData *IEthCrossChainDataTransactorSession) InitGenesisBlock() (*types.Transaction, error) {
	return _IEthCrossChainData.Contract.InitGenesisBlock(&_IEthCrossChainData.TransactOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns(bool)
func (_IEthCrossChainData *IEthCrossChainDataTransactor) Pause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IEthCrossChainData.contract.Transact(opts, "pause")
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns(bool)
func (_IEthCrossChainData *IEthCrossChainDataSession) Pause() (*types.Transaction, error) {
	return _IEthCrossChainData.Contract.Pause(&_IEthCrossChainData.TransactOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns(bool)
func (_IEthCrossChainData *IEthCrossChainDataTransactorSession) Pause() (*types.Transaction, error) {
	return _IEthCrossChainData.Contract.Pause(&_IEthCrossChainData.TransactOpts)
}

// PutCrossChainTxExist is a paid mutator transaction binding the contract method 0x794b0300.
//
// Solidity: function putCrossChainTxExist(bytes32 _crossChainTx) returns(bool)
func (_IEthCrossChainData *IEthCrossChainDataTransactor) PutCrossChainTxExist(opts *bind.TransactOpts, _crossChainTx [32]byte) (*types.Transaction, error) {
	return _IEthCrossChainData.contract.Transact(opts, "putCrossChainTxExist", _crossChainTx)
}

// PutCrossChainTxExist is a paid mutator transaction binding the contract method 0x794b0300.
//
// Solidity: function putCrossChainTxExist(bytes32 _crossChainTx) returns(bool)
func (_IEthCrossChainData *IEthCrossChainDataSession) PutCrossChainTxExist(_crossChainTx [32]byte) (*types.Transaction, error) {
	return _IEthCrossChainData.Contract.PutCrossChainTxExist(&_IEthCrossChainData.TransactOpts, _crossChainTx)
}

// PutCrossChainTxExist is a paid mutator transaction binding the contract method 0x794b0300.
//
// Solidity: function putCrossChainTxExist(bytes32 _crossChainTx) returns(bool)
func (_IEthCrossChainData *IEthCrossChainDataTransactorSession) PutCrossChainTxExist(_crossChainTx [32]byte) (*types.Transaction, error) {
	return _IEthCrossChainData.Contract.PutCrossChainTxExist(&_IEthCrossChainData.TransactOpts, _crossChainTx)
}

// PutEthTxHash is a paid mutator transaction binding the contract method 0x9c59af4f.
//
// Solidity: function putEthTxHash(uint256 _ethTxHashIndex, bytes32 _ethTxHash) returns(bool)
func (_IEthCrossChainData *IEthCrossChainDataTransactor) PutEthTxHash(opts *bind.TransactOpts, _ethTxHashIndex *big.Int, _ethTxHash [32]byte) (*types.Transaction, error) {
	return _IEthCrossChainData.contract.Transact(opts, "putEthTxHash", _ethTxHashIndex, _ethTxHash)
}

// PutEthTxHash is a paid mutator transaction binding the contract method 0x9c59af4f.
//
// Solidity: function putEthTxHash(uint256 _ethTxHashIndex, bytes32 _ethTxHash) returns(bool)
func (_IEthCrossChainData *IEthCrossChainDataSession) PutEthTxHash(_ethTxHashIndex *big.Int, _ethTxHash [32]byte) (*types.Transaction, error) {
	return _IEthCrossChainData.Contract.PutEthTxHash(&_IEthCrossChainData.TransactOpts, _ethTxHashIndex, _ethTxHash)
}

// PutEthTxHash is a paid mutator transaction binding the contract method 0x9c59af4f.
//
// Solidity: function putEthTxHash(uint256 _ethTxHashIndex, bytes32 _ethTxHash) returns(bool)
func (_IEthCrossChainData *IEthCrossChainDataTransactorSession) PutEthTxHash(_ethTxHashIndex *big.Int, _ethTxHash [32]byte) (*types.Transaction, error) {
	return _IEthCrossChainData.Contract.PutEthTxHash(&_IEthCrossChainData.TransactOpts, _ethTxHashIndex, _ethTxHash)
}

// PutExtraData is a paid mutator transaction binding the contract method 0x1afe374e.
//
// Solidity: function putExtraData(bytes32 key1, bytes32 key2, bytes value) returns(bool)
func (_IEthCrossChainData *IEthCrossChainDataTransactor) PutExtraData(opts *bind.TransactOpts, key1 [32]byte, key2 [32]byte, value []byte) (*types.Transaction, error) {
	return _IEthCrossChainData.contract.Transact(opts, "putExtraData", key1, key2, value)
}

// PutExtraData is a paid mutator transaction binding the contract method 0x1afe374e.
//
// Solidity: function putExtraData(bytes32 key1, bytes32 key2, bytes value) returns(bool)
func (_IEthCrossChainData *IEthCrossChainDataSession) PutExtraData(key1 [32]byte, key2 [32]byte, value []byte) (*types.Transaction, error) {
	return _IEthCrossChainData.Contract.PutExtraData(&_IEthCrossChainData.TransactOpts, key1, key2, value)
}

// PutExtraData is a paid mutator transaction binding the contract method 0x1afe374e.
//
// Solidity: function putExtraData(bytes32 key1, bytes32 key2, bytes value) returns(bool)
func (_IEthCrossChainData *IEthCrossChainDataTransactorSession) PutExtraData(key1 [32]byte, key2 [32]byte, value []byte) (*types.Transaction, error) {
	return _IEthCrossChainData.Contract.PutExtraData(&_IEthCrossChainData.TransactOpts, key1, key2, value)
}

// PutLatestBookKeeperHeight is a paid mutator transaction binding the contract method 0xa6c6fe7c.
//
// Solidity: function putLatestBookKeeperHeight(uint64 _latestBookKeeperHeight) returns(bool)
func (_IEthCrossChainData *IEthCrossChainDataTransactor) PutLatestBookKeeperHeight(opts *bind.TransactOpts, _latestBookKeeperHeight uint64) (*types.Transaction, error) {
	return _IEthCrossChainData.contract.Transact(opts, "putLatestBookKeeperHeight", _latestBookKeeperHeight)
}

// PutLatestBookKeeperHeight is a paid mutator transaction binding the contract method 0xa6c6fe7c.
//
// Solidity: function putLatestBookKeeperHeight(uint64 _latestBookKeeperHeight) returns(bool)
func (_IEthCrossChainData *IEthCrossChainDataSession) PutLatestBookKeeperHeight(_latestBookKeeperHeight uint64) (*types.Transaction, error) {
	return _IEthCrossChainData.Contract.PutLatestBookKeeperHeight(&_IEthCrossChainData.TransactOpts, _latestBookKeeperHeight)
}

// PutLatestBookKeeperHeight is a paid mutator transaction binding the contract method 0xa6c6fe7c.
//
// Solidity: function putLatestBookKeeperHeight(uint64 _latestBookKeeperHeight) returns(bool)
func (_IEthCrossChainData *IEthCrossChainDataTransactorSession) PutLatestBookKeeperHeight(_latestBookKeeperHeight uint64) (*types.Transaction, error) {
	return _IEthCrossChainData.Contract.PutLatestBookKeeperHeight(&_IEthCrossChainData.TransactOpts, _latestBookKeeperHeight)
}

// PutLatestHeight is a paid mutator transaction binding the contract method 0x6d044440.
//
// Solidity: function putLatestHeight(uint64 _latestHeight) returns(bool)
func (_IEthCrossChainData *IEthCrossChainDataTransactor) PutLatestHeight(opts *bind.TransactOpts, _latestHeight uint64) (*types.Transaction, error) {
	return _IEthCrossChainData.contract.Transact(opts, "putLatestHeight", _latestHeight)
}

// PutLatestHeight is a paid mutator transaction binding the contract method 0x6d044440.
//
// Solidity: function putLatestHeight(uint64 _latestHeight) returns(bool)
func (_IEthCrossChainData *IEthCrossChainDataSession) PutLatestHeight(_latestHeight uint64) (*types.Transaction, error) {
	return _IEthCrossChainData.Contract.PutLatestHeight(&_IEthCrossChainData.TransactOpts, _latestHeight)
}

// PutLatestHeight is a paid mutator transaction binding the contract method 0x6d044440.
//
// Solidity: function putLatestHeight(uint64 _latestHeight) returns(bool)
func (_IEthCrossChainData *IEthCrossChainDataTransactorSession) PutLatestHeight(_latestHeight uint64) (*types.Transaction, error) {
	return _IEthCrossChainData.Contract.PutLatestHeight(&_IEthCrossChainData.TransactOpts, _latestHeight)
}

// PutMCHeaderBytes is a paid mutator transaction binding the contract method 0x021d70ab.
//
// Solidity: function putMCHeaderBytes(uint64 _height, bytes _rawHeader) returns(bool)
func (_IEthCrossChainData *IEthCrossChainDataTransactor) PutMCHeaderBytes(opts *bind.TransactOpts, _height uint64, _rawHeader []byte) (*types.Transaction, error) {
	return _IEthCrossChainData.contract.Transact(opts, "putMCHeaderBytes", _height, _rawHeader)
}

// PutMCHeaderBytes is a paid mutator transaction binding the contract method 0x021d70ab.
//
// Solidity: function putMCHeaderBytes(uint64 _height, bytes _rawHeader) returns(bool)
func (_IEthCrossChainData *IEthCrossChainDataSession) PutMCHeaderBytes(_height uint64, _rawHeader []byte) (*types.Transaction, error) {
	return _IEthCrossChainData.Contract.PutMCHeaderBytes(&_IEthCrossChainData.TransactOpts, _height, _rawHeader)
}

// PutMCHeaderBytes is a paid mutator transaction binding the contract method 0x021d70ab.
//
// Solidity: function putMCHeaderBytes(uint64 _height, bytes _rawHeader) returns(bool)
func (_IEthCrossChainData *IEthCrossChainDataTransactorSession) PutMCHeaderBytes(_height uint64, _rawHeader []byte) (*types.Transaction, error) {
	return _IEthCrossChainData.Contract.PutMCHeaderBytes(&_IEthCrossChainData.TransactOpts, _height, _rawHeader)
}

// PutMCKeeperHeight is a paid mutator transaction binding the contract method 0xc6dbd0b7.
//
// Solidity: function putMCKeeperHeight(uint64 height) returns(bool)
func (_IEthCrossChainData *IEthCrossChainDataTransactor) PutMCKeeperHeight(opts *bind.TransactOpts, height uint64) (*types.Transaction, error) {
	return _IEthCrossChainData.contract.Transact(opts, "putMCKeeperHeight", height)
}

// PutMCKeeperHeight is a paid mutator transaction binding the contract method 0xc6dbd0b7.
//
// Solidity: function putMCKeeperHeight(uint64 height) returns(bool)
func (_IEthCrossChainData *IEthCrossChainDataSession) PutMCKeeperHeight(height uint64) (*types.Transaction, error) {
	return _IEthCrossChainData.Contract.PutMCKeeperHeight(&_IEthCrossChainData.TransactOpts, height)
}

// PutMCKeeperHeight is a paid mutator transaction binding the contract method 0xc6dbd0b7.
//
// Solidity: function putMCKeeperHeight(uint64 height) returns(bool)
func (_IEthCrossChainData *IEthCrossChainDataTransactorSession) PutMCKeeperHeight(height uint64) (*types.Transaction, error) {
	return _IEthCrossChainData.Contract.PutMCKeeperHeight(&_IEthCrossChainData.TransactOpts, height)
}

// PutMCKeeperPubBytes is a paid mutator transaction binding the contract method 0x191b0ab1.
//
// Solidity: function putMCKeeperPubBytes(uint64 height, bytes _MCKeeperPubBytes) returns(bool)
func (_IEthCrossChainData *IEthCrossChainDataTransactor) PutMCKeeperPubBytes(opts *bind.TransactOpts, height uint64, _MCKeeperPubBytes []byte) (*types.Transaction, error) {
	return _IEthCrossChainData.contract.Transact(opts, "putMCKeeperPubBytes", height, _MCKeeperPubBytes)
}

// PutMCKeeperPubBytes is a paid mutator transaction binding the contract method 0x191b0ab1.
//
// Solidity: function putMCKeeperPubBytes(uint64 height, bytes _MCKeeperPubBytes) returns(bool)
func (_IEthCrossChainData *IEthCrossChainDataSession) PutMCKeeperPubBytes(height uint64, _MCKeeperPubBytes []byte) (*types.Transaction, error) {
	return _IEthCrossChainData.Contract.PutMCKeeperPubBytes(&_IEthCrossChainData.TransactOpts, height, _MCKeeperPubBytes)
}

// PutMCKeeperPubBytes is a paid mutator transaction binding the contract method 0x191b0ab1.
//
// Solidity: function putMCKeeperPubBytes(uint64 height, bytes _MCKeeperPubBytes) returns(bool)
func (_IEthCrossChainData *IEthCrossChainDataTransactorSession) PutMCKeeperPubBytes(height uint64, _MCKeeperPubBytes []byte) (*types.Transaction, error) {
	return _IEthCrossChainData.Contract.PutMCKeeperPubBytes(&_IEthCrossChainData.TransactOpts, height, _MCKeeperPubBytes)
}

// PutMCKeeperPubKeybytes is a paid mutator transaction binding the contract method 0x73c2f5f6.
//
// Solidity: function putMCKeeperPubKeybytes(uint64 _height, bytes _keepersBytes) returns(bool)
func (_IEthCrossChainData *IEthCrossChainDataTransactor) PutMCKeeperPubKeybytes(opts *bind.TransactOpts, _height uint64, _keepersBytes []byte) (*types.Transaction, error) {
	return _IEthCrossChainData.contract.Transact(opts, "putMCKeeperPubKeybytes", _height, _keepersBytes)
}

// PutMCKeeperPubKeybytes is a paid mutator transaction binding the contract method 0x73c2f5f6.
//
// Solidity: function putMCKeeperPubKeybytes(uint64 _height, bytes _keepersBytes) returns(bool)
func (_IEthCrossChainData *IEthCrossChainDataSession) PutMCKeeperPubKeybytes(_height uint64, _keepersBytes []byte) (*types.Transaction, error) {
	return _IEthCrossChainData.Contract.PutMCKeeperPubKeybytes(&_IEthCrossChainData.TransactOpts, _height, _keepersBytes)
}

// PutMCKeeperPubKeybytes is a paid mutator transaction binding the contract method 0x73c2f5f6.
//
// Solidity: function putMCKeeperPubKeybytes(uint64 _height, bytes _keepersBytes) returns(bool)
func (_IEthCrossChainData *IEthCrossChainDataTransactorSession) PutMCKeeperPubKeybytes(_height uint64, _keepersBytes []byte) (*types.Transaction, error) {
	return _IEthCrossChainData.Contract.PutMCKeeperPubKeybytes(&_IEthCrossChainData.TransactOpts, _height, _keepersBytes)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_IEthCrossChainData *IEthCrossChainDataTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _IEthCrossChainData.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_IEthCrossChainData *IEthCrossChainDataSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _IEthCrossChainData.Contract.TransferOwnership(&_IEthCrossChainData.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_IEthCrossChainData *IEthCrossChainDataTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _IEthCrossChainData.Contract.TransferOwnership(&_IEthCrossChainData.TransactOpts, newOwner)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns(bool)
func (_IEthCrossChainData *IEthCrossChainDataTransactor) Unpause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IEthCrossChainData.contract.Transact(opts, "unpause")
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns(bool)
func (_IEthCrossChainData *IEthCrossChainDataSession) Unpause() (*types.Transaction, error) {
	return _IEthCrossChainData.Contract.Unpause(&_IEthCrossChainData.TransactOpts)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns(bool)
func (_IEthCrossChainData *IEthCrossChainDataTransactorSession) Unpause() (*types.Transaction, error) {
	return _IEthCrossChainData.Contract.Unpause(&_IEthCrossChainData.TransactOpts)
}

// IEthCrossChainManagerABI is the input ABI used to generate the binding from.
const IEthCrossChainManagerABI = "[{\"constant\":false,\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"_toChainId\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"_toContract\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"_method\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"_txData\",\"type\":\"bytes\"}],\"name\":\"crossChain\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// IEthCrossChainManagerFuncSigs maps the 4-byte function signature to its string representation.
var IEthCrossChainManagerFuncSigs = map[string]string{
	"bd5cf625": "crossChain(uint64,bytes,bytes,bytes)",
}

// IEthCrossChainManager is an auto generated Go binding around an Ethereum contract.
type IEthCrossChainManager struct {
	IEthCrossChainManagerCaller     // Read-only binding to the contract
	IEthCrossChainManagerTransactor // Write-only binding to the contract
	IEthCrossChainManagerFilterer   // Log filterer for contract events
}

// IEthCrossChainManagerCaller is an auto generated read-only Go binding around an Ethereum contract.
type IEthCrossChainManagerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IEthCrossChainManagerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IEthCrossChainManagerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IEthCrossChainManagerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IEthCrossChainManagerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IEthCrossChainManagerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IEthCrossChainManagerSession struct {
	Contract     *IEthCrossChainManager // Generic contract binding to set the session for
	CallOpts     bind.CallOpts          // Call options to use throughout this session
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// IEthCrossChainManagerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IEthCrossChainManagerCallerSession struct {
	Contract *IEthCrossChainManagerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                // Call options to use throughout this session
}

// IEthCrossChainManagerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IEthCrossChainManagerTransactorSession struct {
	Contract     *IEthCrossChainManagerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                // Transaction auth options to use throughout this session
}

// IEthCrossChainManagerRaw is an auto generated low-level Go binding around an Ethereum contract.
type IEthCrossChainManagerRaw struct {
	Contract *IEthCrossChainManager // Generic contract binding to access the raw methods on
}

// IEthCrossChainManagerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IEthCrossChainManagerCallerRaw struct {
	Contract *IEthCrossChainManagerCaller // Generic read-only contract binding to access the raw methods on
}

// IEthCrossChainManagerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IEthCrossChainManagerTransactorRaw struct {
	Contract *IEthCrossChainManagerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIEthCrossChainManager creates a new instance of IEthCrossChainManager, bound to a specific deployed contract.
func NewIEthCrossChainManager(address common.Address, backend bind.ContractBackend) (*IEthCrossChainManager, error) {
	contract, err := bindIEthCrossChainManager(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IEthCrossChainManager{IEthCrossChainManagerCaller: IEthCrossChainManagerCaller{contract: contract}, IEthCrossChainManagerTransactor: IEthCrossChainManagerTransactor{contract: contract}, IEthCrossChainManagerFilterer: IEthCrossChainManagerFilterer{contract: contract}}, nil
}

// NewIEthCrossChainManagerCaller creates a new read-only instance of IEthCrossChainManager, bound to a specific deployed contract.
func NewIEthCrossChainManagerCaller(address common.Address, caller bind.ContractCaller) (*IEthCrossChainManagerCaller, error) {
	contract, err := bindIEthCrossChainManager(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IEthCrossChainManagerCaller{contract: contract}, nil
}

// NewIEthCrossChainManagerTransactor creates a new write-only instance of IEthCrossChainManager, bound to a specific deployed contract.
func NewIEthCrossChainManagerTransactor(address common.Address, transactor bind.ContractTransactor) (*IEthCrossChainManagerTransactor, error) {
	contract, err := bindIEthCrossChainManager(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IEthCrossChainManagerTransactor{contract: contract}, nil
}

// NewIEthCrossChainManagerFilterer creates a new log filterer instance of IEthCrossChainManager, bound to a specific deployed contract.
func NewIEthCrossChainManagerFilterer(address common.Address, filterer bind.ContractFilterer) (*IEthCrossChainManagerFilterer, error) {
	contract, err := bindIEthCrossChainManager(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IEthCrossChainManagerFilterer{contract: contract}, nil
}

// bindIEthCrossChainManager binds a generic wrapper to an already deployed contract.
func bindIEthCrossChainManager(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IEthCrossChainManagerABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IEthCrossChainManager *IEthCrossChainManagerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _IEthCrossChainManager.Contract.IEthCrossChainManagerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IEthCrossChainManager *IEthCrossChainManagerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IEthCrossChainManager.Contract.IEthCrossChainManagerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IEthCrossChainManager *IEthCrossChainManagerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IEthCrossChainManager.Contract.IEthCrossChainManagerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IEthCrossChainManager *IEthCrossChainManagerCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _IEthCrossChainManager.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IEthCrossChainManager *IEthCrossChainManagerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IEthCrossChainManager.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IEthCrossChainManager *IEthCrossChainManagerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IEthCrossChainManager.Contract.contract.Transact(opts, method, params...)
}

// CrossChain is a paid mutator transaction binding the contract method 0xbd5cf625.
//
// Solidity: function crossChain(uint64 _toChainId, bytes _toContract, bytes _method, bytes _txData) returns(bool)
func (_IEthCrossChainManager *IEthCrossChainManagerTransactor) CrossChain(opts *bind.TransactOpts, _toChainId uint64, _toContract []byte, _method []byte, _txData []byte) (*types.Transaction, error) {
	return _IEthCrossChainManager.contract.Transact(opts, "crossChain", _toChainId, _toContract, _method, _txData)
}

// CrossChain is a paid mutator transaction binding the contract method 0xbd5cf625.
//
// Solidity: function crossChain(uint64 _toChainId, bytes _toContract, bytes _method, bytes _txData) returns(bool)
func (_IEthCrossChainManager *IEthCrossChainManagerSession) CrossChain(_toChainId uint64, _toContract []byte, _method []byte, _txData []byte) (*types.Transaction, error) {
	return _IEthCrossChainManager.Contract.CrossChain(&_IEthCrossChainManager.TransactOpts, _toChainId, _toContract, _method, _txData)
}

// CrossChain is a paid mutator transaction binding the contract method 0xbd5cf625.
//
// Solidity: function crossChain(uint64 _toChainId, bytes _toContract, bytes _method, bytes _txData) returns(bool)
func (_IEthCrossChainManager *IEthCrossChainManagerTransactorSession) CrossChain(_toChainId uint64, _toContract []byte, _method []byte, _txData []byte) (*types.Transaction, error) {
	return _IEthCrossChainManager.Contract.CrossChain(&_IEthCrossChainManager.TransactOpts, _toChainId, _toContract, _method, _txData)
}

// IUpgradableECCMABI is the input ABI used to generate the binding from.
const IUpgradableECCMABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"isOwner\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"pause\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"unpause\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"upgradeToNew\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// IUpgradableECCMFuncSigs maps the 4-byte function signature to its string representation.
var IUpgradableECCMFuncSigs = map[string]string{
	"8f32d59b": "isOwner()",
	"8456cb59": "pause()",
	"5c975abb": "paused()",
	"3f4ba83a": "unpause()",
	"7e724ff3": "upgradeToNew(address)",
}

// IUpgradableECCM is an auto generated Go binding around an Ethereum contract.
type IUpgradableECCM struct {
	IUpgradableECCMCaller     // Read-only binding to the contract
	IUpgradableECCMTransactor // Write-only binding to the contract
	IUpgradableECCMFilterer   // Log filterer for contract events
}

// IUpgradableECCMCaller is an auto generated read-only Go binding around an Ethereum contract.
type IUpgradableECCMCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IUpgradableECCMTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IUpgradableECCMTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IUpgradableECCMFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IUpgradableECCMFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IUpgradableECCMSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IUpgradableECCMSession struct {
	Contract     *IUpgradableECCM  // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IUpgradableECCMCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IUpgradableECCMCallerSession struct {
	Contract *IUpgradableECCMCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts          // Call options to use throughout this session
}

// IUpgradableECCMTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IUpgradableECCMTransactorSession struct {
	Contract     *IUpgradableECCMTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts          // Transaction auth options to use throughout this session
}

// IUpgradableECCMRaw is an auto generated low-level Go binding around an Ethereum contract.
type IUpgradableECCMRaw struct {
	Contract *IUpgradableECCM // Generic contract binding to access the raw methods on
}

// IUpgradableECCMCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IUpgradableECCMCallerRaw struct {
	Contract *IUpgradableECCMCaller // Generic read-only contract binding to access the raw methods on
}

// IUpgradableECCMTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IUpgradableECCMTransactorRaw struct {
	Contract *IUpgradableECCMTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIUpgradableECCM creates a new instance of IUpgradableECCM, bound to a specific deployed contract.
func NewIUpgradableECCM(address common.Address, backend bind.ContractBackend) (*IUpgradableECCM, error) {
	contract, err := bindIUpgradableECCM(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IUpgradableECCM{IUpgradableECCMCaller: IUpgradableECCMCaller{contract: contract}, IUpgradableECCMTransactor: IUpgradableECCMTransactor{contract: contract}, IUpgradableECCMFilterer: IUpgradableECCMFilterer{contract: contract}}, nil
}

// NewIUpgradableECCMCaller creates a new read-only instance of IUpgradableECCM, bound to a specific deployed contract.
func NewIUpgradableECCMCaller(address common.Address, caller bind.ContractCaller) (*IUpgradableECCMCaller, error) {
	contract, err := bindIUpgradableECCM(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IUpgradableECCMCaller{contract: contract}, nil
}

// NewIUpgradableECCMTransactor creates a new write-only instance of IUpgradableECCM, bound to a specific deployed contract.
func NewIUpgradableECCMTransactor(address common.Address, transactor bind.ContractTransactor) (*IUpgradableECCMTransactor, error) {
	contract, err := bindIUpgradableECCM(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IUpgradableECCMTransactor{contract: contract}, nil
}

// NewIUpgradableECCMFilterer creates a new log filterer instance of IUpgradableECCM, bound to a specific deployed contract.
func NewIUpgradableECCMFilterer(address common.Address, filterer bind.ContractFilterer) (*IUpgradableECCMFilterer, error) {
	contract, err := bindIUpgradableECCM(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IUpgradableECCMFilterer{contract: contract}, nil
}

// bindIUpgradableECCM binds a generic wrapper to an already deployed contract.
func bindIUpgradableECCM(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IUpgradableECCMABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IUpgradableECCM *IUpgradableECCMRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _IUpgradableECCM.Contract.IUpgradableECCMCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IUpgradableECCM *IUpgradableECCMRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IUpgradableECCM.Contract.IUpgradableECCMTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IUpgradableECCM *IUpgradableECCMRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IUpgradableECCM.Contract.IUpgradableECCMTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IUpgradableECCM *IUpgradableECCMCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _IUpgradableECCM.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IUpgradableECCM *IUpgradableECCMTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IUpgradableECCM.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IUpgradableECCM *IUpgradableECCMTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IUpgradableECCM.Contract.contract.Transact(opts, method, params...)
}

// IsOwner is a free data retrieval call binding the contract method 0x8f32d59b.
//
// Solidity: function isOwner() constant returns(bool)
func (_IUpgradableECCM *IUpgradableECCMCaller) IsOwner(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _IUpgradableECCM.contract.Call(opts, out, "isOwner")
	return *ret0, err
}

// IsOwner is a free data retrieval call binding the contract method 0x8f32d59b.
//
// Solidity: function isOwner() constant returns(bool)
func (_IUpgradableECCM *IUpgradableECCMSession) IsOwner() (bool, error) {
	return _IUpgradableECCM.Contract.IsOwner(&_IUpgradableECCM.CallOpts)
}

// IsOwner is a free data retrieval call binding the contract method 0x8f32d59b.
//
// Solidity: function isOwner() constant returns(bool)
func (_IUpgradableECCM *IUpgradableECCMCallerSession) IsOwner() (bool, error) {
	return _IUpgradableECCM.Contract.IsOwner(&_IUpgradableECCM.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() constant returns(bool)
func (_IUpgradableECCM *IUpgradableECCMCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _IUpgradableECCM.contract.Call(opts, out, "paused")
	return *ret0, err
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() constant returns(bool)
func (_IUpgradableECCM *IUpgradableECCMSession) Paused() (bool, error) {
	return _IUpgradableECCM.Contract.Paused(&_IUpgradableECCM.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() constant returns(bool)
func (_IUpgradableECCM *IUpgradableECCMCallerSession) Paused() (bool, error) {
	return _IUpgradableECCM.Contract.Paused(&_IUpgradableECCM.CallOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns(bool)
func (_IUpgradableECCM *IUpgradableECCMTransactor) Pause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IUpgradableECCM.contract.Transact(opts, "pause")
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns(bool)
func (_IUpgradableECCM *IUpgradableECCMSession) Pause() (*types.Transaction, error) {
	return _IUpgradableECCM.Contract.Pause(&_IUpgradableECCM.TransactOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns(bool)
func (_IUpgradableECCM *IUpgradableECCMTransactorSession) Pause() (*types.Transaction, error) {
	return _IUpgradableECCM.Contract.Pause(&_IUpgradableECCM.TransactOpts)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns(bool)
func (_IUpgradableECCM *IUpgradableECCMTransactor) Unpause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IUpgradableECCM.contract.Transact(opts, "unpause")
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns(bool)
func (_IUpgradableECCM *IUpgradableECCMSession) Unpause() (*types.Transaction, error) {
	return _IUpgradableECCM.Contract.Unpause(&_IUpgradableECCM.TransactOpts)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns(bool)
func (_IUpgradableECCM *IUpgradableECCMTransactorSession) Unpause() (*types.Transaction, error) {
	return _IUpgradableECCM.Contract.Unpause(&_IUpgradableECCM.TransactOpts)
}

// UpgradeToNew is a paid mutator transaction binding the contract method 0x7e724ff3.
//
// Solidity: function upgradeToNew(address ) returns(bool)
func (_IUpgradableECCM *IUpgradableECCMTransactor) UpgradeToNew(opts *bind.TransactOpts, arg0 common.Address) (*types.Transaction, error) {
	return _IUpgradableECCM.contract.Transact(opts, "upgradeToNew", arg0)
}

// UpgradeToNew is a paid mutator transaction binding the contract method 0x7e724ff3.
//
// Solidity: function upgradeToNew(address ) returns(bool)
func (_IUpgradableECCM *IUpgradableECCMSession) UpgradeToNew(arg0 common.Address) (*types.Transaction, error) {
	return _IUpgradableECCM.Contract.UpgradeToNew(&_IUpgradableECCM.TransactOpts, arg0)
}

// UpgradeToNew is a paid mutator transaction binding the contract method 0x7e724ff3.
//
// Solidity: function upgradeToNew(address ) returns(bool)
func (_IUpgradableECCM *IUpgradableECCMTransactorSession) UpgradeToNew(arg0 common.Address) (*types.Transaction, error) {
	return _IUpgradableECCM.Contract.UpgradeToNew(&_IUpgradableECCM.TransactOpts, arg0)
}

// OwnableABI is the input ABI used to generate the binding from.
const OwnableABI = "[{\"inputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"constant\":true,\"inputs\":[],\"name\":\"isOwner\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// OwnableFuncSigs maps the 4-byte function signature to its string representation.
var OwnableFuncSigs = map[string]string{
	"8f32d59b": "isOwner()",
	"8da5cb5b": "owner()",
	"715018a6": "renounceOwnership()",
	"f2fde38b": "transferOwnership(address)",
}

// Ownable is an auto generated Go binding around an Ethereum contract.
type Ownable struct {
	OwnableCaller     // Read-only binding to the contract
	OwnableTransactor // Write-only binding to the contract
	OwnableFilterer   // Log filterer for contract events
}

// OwnableCaller is an auto generated read-only Go binding around an Ethereum contract.
type OwnableCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OwnableTransactor is an auto generated write-only Go binding around an Ethereum contract.
type OwnableTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OwnableFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type OwnableFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OwnableSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type OwnableSession struct {
	Contract     *Ownable          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// OwnableCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type OwnableCallerSession struct {
	Contract *OwnableCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// OwnableTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type OwnableTransactorSession struct {
	Contract     *OwnableTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// OwnableRaw is an auto generated low-level Go binding around an Ethereum contract.
type OwnableRaw struct {
	Contract *Ownable // Generic contract binding to access the raw methods on
}

// OwnableCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type OwnableCallerRaw struct {
	Contract *OwnableCaller // Generic read-only contract binding to access the raw methods on
}

// OwnableTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type OwnableTransactorRaw struct {
	Contract *OwnableTransactor // Generic write-only contract binding to access the raw methods on
}

// NewOwnable creates a new instance of Ownable, bound to a specific deployed contract.
func NewOwnable(address common.Address, backend bind.ContractBackend) (*Ownable, error) {
	contract, err := bindOwnable(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Ownable{OwnableCaller: OwnableCaller{contract: contract}, OwnableTransactor: OwnableTransactor{contract: contract}, OwnableFilterer: OwnableFilterer{contract: contract}}, nil
}

// NewOwnableCaller creates a new read-only instance of Ownable, bound to a specific deployed contract.
func NewOwnableCaller(address common.Address, caller bind.ContractCaller) (*OwnableCaller, error) {
	contract, err := bindOwnable(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &OwnableCaller{contract: contract}, nil
}

// NewOwnableTransactor creates a new write-only instance of Ownable, bound to a specific deployed contract.
func NewOwnableTransactor(address common.Address, transactor bind.ContractTransactor) (*OwnableTransactor, error) {
	contract, err := bindOwnable(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &OwnableTransactor{contract: contract}, nil
}

// NewOwnableFilterer creates a new log filterer instance of Ownable, bound to a specific deployed contract.
func NewOwnableFilterer(address common.Address, filterer bind.ContractFilterer) (*OwnableFilterer, error) {
	contract, err := bindOwnable(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &OwnableFilterer{contract: contract}, nil
}

// bindOwnable binds a generic wrapper to an already deployed contract.
func bindOwnable(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(OwnableABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Ownable *OwnableRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Ownable.Contract.OwnableCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Ownable *OwnableRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Ownable.Contract.OwnableTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Ownable *OwnableRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Ownable.Contract.OwnableTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Ownable *OwnableCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Ownable.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Ownable *OwnableTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Ownable.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Ownable *OwnableTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Ownable.Contract.contract.Transact(opts, method, params...)
}

// IsOwner is a free data retrieval call binding the contract method 0x8f32d59b.
//
// Solidity: function isOwner() constant returns(bool)
func (_Ownable *OwnableCaller) IsOwner(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Ownable.contract.Call(opts, out, "isOwner")
	return *ret0, err
}

// IsOwner is a free data retrieval call binding the contract method 0x8f32d59b.
//
// Solidity: function isOwner() constant returns(bool)
func (_Ownable *OwnableSession) IsOwner() (bool, error) {
	return _Ownable.Contract.IsOwner(&_Ownable.CallOpts)
}

// IsOwner is a free data retrieval call binding the contract method 0x8f32d59b.
//
// Solidity: function isOwner() constant returns(bool)
func (_Ownable *OwnableCallerSession) IsOwner() (bool, error) {
	return _Ownable.Contract.IsOwner(&_Ownable.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Ownable *OwnableCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Ownable.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Ownable *OwnableSession) Owner() (common.Address, error) {
	return _Ownable.Contract.Owner(&_Ownable.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Ownable *OwnableCallerSession) Owner() (common.Address, error) {
	return _Ownable.Contract.Owner(&_Ownable.CallOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Ownable *OwnableTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Ownable.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Ownable *OwnableSession) RenounceOwnership() (*types.Transaction, error) {
	return _Ownable.Contract.RenounceOwnership(&_Ownable.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Ownable *OwnableTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Ownable.Contract.RenounceOwnership(&_Ownable.TransactOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Ownable *OwnableTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Ownable.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Ownable *OwnableSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Ownable.Contract.TransferOwnership(&_Ownable.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Ownable *OwnableTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Ownable.Contract.TransferOwnership(&_Ownable.TransactOpts, newOwner)
}

// OwnableOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Ownable contract.
type OwnableOwnershipTransferredIterator struct {
	Event *OwnableOwnershipTransferred // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *OwnableOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OwnableOwnershipTransferred)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(OwnableOwnershipTransferred)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *OwnableOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OwnableOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OwnableOwnershipTransferred represents a OwnershipTransferred event raised by the Ownable contract.
type OwnableOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Ownable *OwnableFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*OwnableOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Ownable.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &OwnableOwnershipTransferredIterator{contract: _Ownable.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Ownable *OwnableFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *OwnableOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Ownable.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OwnableOwnershipTransferred)
				if err := _Ownable.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseOwnershipTransferred is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Ownable *OwnableFilterer) ParseOwnershipTransferred(log types.Log) (*OwnableOwnershipTransferred, error) {
	event := new(OwnableOwnershipTransferred)
	if err := _Ownable.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	return event, nil
}

// PausableABI is the input ABI used to generate the binding from.
const PausableABI = "[{\"inputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Paused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Unpaused\",\"type\":\"event\"},{\"constant\":true,\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"}]"

// PausableFuncSigs maps the 4-byte function signature to its string representation.
var PausableFuncSigs = map[string]string{
	"5c975abb": "paused()",
}

// Pausable is an auto generated Go binding around an Ethereum contract.
type Pausable struct {
	PausableCaller     // Read-only binding to the contract
	PausableTransactor // Write-only binding to the contract
	PausableFilterer   // Log filterer for contract events
}

// PausableCaller is an auto generated read-only Go binding around an Ethereum contract.
type PausableCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PausableTransactor is an auto generated write-only Go binding around an Ethereum contract.
type PausableTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PausableFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type PausableFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PausableSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type PausableSession struct {
	Contract     *Pausable         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// PausableCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type PausableCallerSession struct {
	Contract *PausableCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// PausableTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type PausableTransactorSession struct {
	Contract     *PausableTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// PausableRaw is an auto generated low-level Go binding around an Ethereum contract.
type PausableRaw struct {
	Contract *Pausable // Generic contract binding to access the raw methods on
}

// PausableCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type PausableCallerRaw struct {
	Contract *PausableCaller // Generic read-only contract binding to access the raw methods on
}

// PausableTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type PausableTransactorRaw struct {
	Contract *PausableTransactor // Generic write-only contract binding to access the raw methods on
}

// NewPausable creates a new instance of Pausable, bound to a specific deployed contract.
func NewPausable(address common.Address, backend bind.ContractBackend) (*Pausable, error) {
	contract, err := bindPausable(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Pausable{PausableCaller: PausableCaller{contract: contract}, PausableTransactor: PausableTransactor{contract: contract}, PausableFilterer: PausableFilterer{contract: contract}}, nil
}

// NewPausableCaller creates a new read-only instance of Pausable, bound to a specific deployed contract.
func NewPausableCaller(address common.Address, caller bind.ContractCaller) (*PausableCaller, error) {
	contract, err := bindPausable(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &PausableCaller{contract: contract}, nil
}

// NewPausableTransactor creates a new write-only instance of Pausable, bound to a specific deployed contract.
func NewPausableTransactor(address common.Address, transactor bind.ContractTransactor) (*PausableTransactor, error) {
	contract, err := bindPausable(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &PausableTransactor{contract: contract}, nil
}

// NewPausableFilterer creates a new log filterer instance of Pausable, bound to a specific deployed contract.
func NewPausableFilterer(address common.Address, filterer bind.ContractFilterer) (*PausableFilterer, error) {
	contract, err := bindPausable(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &PausableFilterer{contract: contract}, nil
}

// bindPausable binds a generic wrapper to an already deployed contract.
func bindPausable(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(PausableABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Pausable *PausableRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Pausable.Contract.PausableCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Pausable *PausableRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Pausable.Contract.PausableTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Pausable *PausableRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Pausable.Contract.PausableTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Pausable *PausableCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Pausable.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Pausable *PausableTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Pausable.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Pausable *PausableTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Pausable.Contract.contract.Transact(opts, method, params...)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() constant returns(bool)
func (_Pausable *PausableCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Pausable.contract.Call(opts, out, "paused")
	return *ret0, err
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() constant returns(bool)
func (_Pausable *PausableSession) Paused() (bool, error) {
	return _Pausable.Contract.Paused(&_Pausable.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() constant returns(bool)
func (_Pausable *PausableCallerSession) Paused() (bool, error) {
	return _Pausable.Contract.Paused(&_Pausable.CallOpts)
}

// PausablePausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the Pausable contract.
type PausablePausedIterator struct {
	Event *PausablePaused // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *PausablePausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PausablePaused)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(PausablePaused)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *PausablePausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PausablePausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PausablePaused represents a Paused event raised by the Pausable contract.
type PausablePaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_Pausable *PausableFilterer) FilterPaused(opts *bind.FilterOpts) (*PausablePausedIterator, error) {

	logs, sub, err := _Pausable.contract.FilterLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return &PausablePausedIterator{contract: _Pausable.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_Pausable *PausableFilterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *PausablePaused) (event.Subscription, error) {

	logs, sub, err := _Pausable.contract.WatchLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PausablePaused)
				if err := _Pausable.contract.UnpackLog(event, "Paused", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParsePaused is a log parse operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_Pausable *PausableFilterer) ParsePaused(log types.Log) (*PausablePaused, error) {
	event := new(PausablePaused)
	if err := _Pausable.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	return event, nil
}

// PausableUnpausedIterator is returned from FilterUnpaused and is used to iterate over the raw logs and unpacked data for Unpaused events raised by the Pausable contract.
type PausableUnpausedIterator struct {
	Event *PausableUnpaused // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *PausableUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PausableUnpaused)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(PausableUnpaused)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *PausableUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PausableUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PausableUnpaused represents a Unpaused event raised by the Pausable contract.
type PausableUnpaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterUnpaused is a free log retrieval operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_Pausable *PausableFilterer) FilterUnpaused(opts *bind.FilterOpts) (*PausableUnpausedIterator, error) {

	logs, sub, err := _Pausable.contract.FilterLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return &PausableUnpausedIterator{contract: _Pausable.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

// WatchUnpaused is a free log subscription operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_Pausable *PausableFilterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *PausableUnpaused) (event.Subscription, error) {

	logs, sub, err := _Pausable.contract.WatchLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PausableUnpaused)
				if err := _Pausable.contract.UnpackLog(event, "Unpaused", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseUnpaused is a log parse operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_Pausable *PausableFilterer) ParseUnpaused(log types.Log) (*PausableUnpaused, error) {
	event := new(PausableUnpaused)
	if err := _Pausable.contract.UnpackLog(event, "Unpaused", log); err != nil {
		return nil, err
	}
	return event, nil
}

// SafeMathABI is the input ABI used to generate the binding from.
const SafeMathABI = "[]"

// SafeMathBin is the compiled bytecode used for deploying new contracts.
var SafeMathBin = "0x60556023600b82828239805160001a607314601657fe5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea265627a7a7231582067215feb96e2a5cea59332df548b5991ec29be76f39e8f9067e56e428d84bfea64736f6c634300050f0032"

// DeploySafeMath deploys a new Ethereum contract, binding an instance of SafeMath to it.
func DeploySafeMath(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *SafeMath, error) {
	parsed, err := abi.JSON(strings.NewReader(SafeMathABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(SafeMathBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &SafeMath{SafeMathCaller: SafeMathCaller{contract: contract}, SafeMathTransactor: SafeMathTransactor{contract: contract}, SafeMathFilterer: SafeMathFilterer{contract: contract}}, nil
}

// SafeMath is an auto generated Go binding around an Ethereum contract.
type SafeMath struct {
	SafeMathCaller     // Read-only binding to the contract
	SafeMathTransactor // Write-only binding to the contract
	SafeMathFilterer   // Log filterer for contract events
}

// SafeMathCaller is an auto generated read-only Go binding around an Ethereum contract.
type SafeMathCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SafeMathTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SafeMathTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SafeMathFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SafeMathFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SafeMathSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SafeMathSession struct {
	Contract     *SafeMath         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SafeMathCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SafeMathCallerSession struct {
	Contract *SafeMathCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// SafeMathTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SafeMathTransactorSession struct {
	Contract     *SafeMathTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// SafeMathRaw is an auto generated low-level Go binding around an Ethereum contract.
type SafeMathRaw struct {
	Contract *SafeMath // Generic contract binding to access the raw methods on
}

// SafeMathCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SafeMathCallerRaw struct {
	Contract *SafeMathCaller // Generic read-only contract binding to access the raw methods on
}

// SafeMathTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SafeMathTransactorRaw struct {
	Contract *SafeMathTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSafeMath creates a new instance of SafeMath, bound to a specific deployed contract.
func NewSafeMath(address common.Address, backend bind.ContractBackend) (*SafeMath, error) {
	contract, err := bindSafeMath(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &SafeMath{SafeMathCaller: SafeMathCaller{contract: contract}, SafeMathTransactor: SafeMathTransactor{contract: contract}, SafeMathFilterer: SafeMathFilterer{contract: contract}}, nil
}

// NewSafeMathCaller creates a new read-only instance of SafeMath, bound to a specific deployed contract.
func NewSafeMathCaller(address common.Address, caller bind.ContractCaller) (*SafeMathCaller, error) {
	contract, err := bindSafeMath(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SafeMathCaller{contract: contract}, nil
}

// NewSafeMathTransactor creates a new write-only instance of SafeMath, bound to a specific deployed contract.
func NewSafeMathTransactor(address common.Address, transactor bind.ContractTransactor) (*SafeMathTransactor, error) {
	contract, err := bindSafeMath(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SafeMathTransactor{contract: contract}, nil
}

// NewSafeMathFilterer creates a new log filterer instance of SafeMath, bound to a specific deployed contract.
func NewSafeMathFilterer(address common.Address, filterer bind.ContractFilterer) (*SafeMathFilterer, error) {
	contract, err := bindSafeMath(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SafeMathFilterer{contract: contract}, nil
}

// bindSafeMath binds a generic wrapper to an already deployed contract.
func bindSafeMath(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(SafeMathABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SafeMath *SafeMathRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _SafeMath.Contract.SafeMathCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SafeMath *SafeMathRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SafeMath.Contract.SafeMathTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SafeMath *SafeMathRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SafeMath.Contract.SafeMathTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SafeMath *SafeMathCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _SafeMath.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SafeMath *SafeMathTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SafeMath.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SafeMath *SafeMathTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SafeMath.Contract.contract.Transact(opts, method, params...)
}

// UpgradableECCMABI is the input ABI used to generate the binding from.
const UpgradableECCMABI = "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"ethCrossChainDataAddr\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Paused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Unpaused\",\"type\":\"event\"},{\"constant\":true,\"inputs\":[],\"name\":\"EthCrossChainDataAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"isOwner\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"pause\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"unpause\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"newEthCrossChainManagerAddress\",\"type\":\"address\"}],\"name\":\"upgradeToNew\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// UpgradableECCMFuncSigs maps the 4-byte function signature to its string representation.
var UpgradableECCMFuncSigs = map[string]string{
	"00ba1694": "EthCrossChainDataAddress()",
	"8f32d59b": "isOwner()",
	"8da5cb5b": "owner()",
	"8456cb59": "pause()",
	"5c975abb": "paused()",
	"715018a6": "renounceOwnership()",
	"f2fde38b": "transferOwnership(address)",
	"3f4ba83a": "unpause()",
	"7e724ff3": "upgradeToNew(address)",
}

// UpgradableECCMBin is the compiled bytecode used for deploying new contracts.
var UpgradableECCMBin = "0x608060405234801561001057600080fd5b506040516109fe3803806109fe8339818101604052602081101561003357600080fd5b505160006100486001600160e01b036100c416565b600080546001600160a01b0319166001600160a01b0383169081178255604051929350917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0908290a3506000805460ff60a01b19169055600180546001600160a01b0319166001600160a01b03929092169190911790556100c8565b3390565b610927806100d76000396000f3fe608060405234801561001057600080fd5b50600436106100925760003560e01c80637e724ff3116100665780637e724ff3146100e95780638456cb591461010f5780638da5cb5b146101175780638f32d59b1461011f578063f2fde38b1461012757610092565b8062ba1694146100975780633f4ba83a146100bb5780635c975abb146100d7578063715018a6146100df575b600080fd5b61009f61014d565b604080516001600160a01b039092168252519081900360200190f35b6100c361015c565b604080519115158252519081900360200190f35b6100c36102db565b6100e76102eb565b005b6100c3600480360360208110156100ff57600080fd5b50356001600160a01b031661037c565b6100c361048b565b61009f610600565b6100c361060f565b6100e76004803603602081101561013d57600080fd5b50356001600160a01b0316610633565b6001546001600160a01b031681565b600061016661060f565b6101a5576040805162461bcd60e51b815260206004820181905260248201526000805160206108d3833981519152604482015290519081900360640190fd5b6101ad6102db565b156101ba576101ba610686565b60015460408051635c975abb60e01b815290516001600160a01b03909216918291635c975abb916004808301926020929190829003018186803b15801561020057600080fd5b505afa158015610214573d6000803e3d6000fd5b505050506040513d602081101561022a57600080fd5b5051156102d357806001600160a01b0316633f4ba83a6040518163ffffffff1660e01b8152600401602060405180830381600087803b15801561026c57600080fd5b505af1158015610280573d6000803e3d6000fd5b505050506040513d602081101561029657600080fd5b50516102d35760405162461bcd60e51b81526004018080602001828103825260298152602001806108846029913960400191505060405180910390fd5b600191505090565b600054600160a01b900460ff1690565b6102f361060f565b610332576040805162461bcd60e51b815260206004820181905260248201526000805160206108d3833981519152604482015290519081900360640190fd5b600080546040516001600160a01b03909116907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0908390a3600080546001600160a01b0319169055565b60008054600160a01b900460ff166103d2576040805162461bcd60e51b815260206004820152601460248201527314185d5cd8589b194e881b9bdd081c185d5cd95960621b604482015290519081900360640190fd5b6103da61060f565b610419576040805162461bcd60e51b815260206004820181905260248201526000805160206108d3833981519152604482015290519081900360640190fd5b6001546040805163f2fde38b60e01b81526001600160a01b03858116600483015291519190921691829163f2fde38b9160248082019260009290919082900301818387803b15801561046a57600080fd5b505af115801561047e573d6000803e3d6000fd5b5060019695505050505050565b600061049561060f565b6104d4576040805162461bcd60e51b815260206004820181905260248201526000805160206108d3833981519152604482015290519081900360640190fd5b6104dc6102db565b6104e8576104e861072e565b60015460408051635c975abb60e01b815290516001600160a01b03909216918291635c975abb916004808301926020929190829003018186803b15801561052e57600080fd5b505afa158015610542573d6000803e3d6000fd5b505050506040513d602081101561055857600080fd5b50516102d357806001600160a01b0316638456cb596040518163ffffffff1660e01b8152600401602060405180830381600087803b15801561059957600080fd5b505af11580156105ad573d6000803e3d6000fd5b505050506040513d60208110156105c357600080fd5b50516102d35760405162461bcd60e51b815260040180806020018281038252602781526020018061085d6027913960400191505060405180910390fd5b6000546001600160a01b031690565b600080546001600160a01b03166106246107b8565b6001600160a01b031614905090565b61063b61060f565b61067a576040805162461bcd60e51b815260206004820181905260248201526000805160206108d3833981519152604482015290519081900360640190fd5b610683816107bc565b50565b600054600160a01b900460ff166106db576040805162461bcd60e51b815260206004820152601460248201527314185d5cd8589b194e881b9bdd081c185d5cd95960621b604482015290519081900360640190fd5b6000805460ff60a01b191690557f5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa6107116107b8565b604080516001600160a01b039092168252519081900360200190a1565b600054600160a01b900460ff1615610780576040805162461bcd60e51b815260206004820152601060248201526f14185d5cd8589b194e881c185d5cd95960821b604482015290519081900360640190fd5b6000805460ff60a01b1916600160a01b1790557f62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a2586107115b3390565b6001600160a01b0381166108015760405162461bcd60e51b81526004018080602001828103825260268152602001806108ad6026913960400191505060405180910390fd5b600080546040516001600160a01b03808516939216917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e091a3600080546001600160a01b0319166001600160a01b039290921691909117905556fe70617573652045746843726f7373436861696e4461746120636f6e7472616374206661696c6564756e70617573652045746843726f7373436861696e4461746120636f6e7472616374206661696c65644f776e61626c653a206e6577206f776e657220697320746865207a65726f20616464726573734f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e6572a265627a7a723158202e6cc1f2a43ca2ab8ddb4846ced0f2f12bcbbf40cc0a0403d12f2e39324a85b564736f6c634300050f0032"

// DeployUpgradableECCM deploys a new Ethereum contract, binding an instance of UpgradableECCM to it.
func DeployUpgradableECCM(auth *bind.TransactOpts, backend bind.ContractBackend, ethCrossChainDataAddr common.Address) (common.Address, *types.Transaction, *UpgradableECCM, error) {
	parsed, err := abi.JSON(strings.NewReader(UpgradableECCMABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(UpgradableECCMBin), backend, ethCrossChainDataAddr)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &UpgradableECCM{UpgradableECCMCaller: UpgradableECCMCaller{contract: contract}, UpgradableECCMTransactor: UpgradableECCMTransactor{contract: contract}, UpgradableECCMFilterer: UpgradableECCMFilterer{contract: contract}}, nil
}

// UpgradableECCM is an auto generated Go binding around an Ethereum contract.
type UpgradableECCM struct {
	UpgradableECCMCaller     // Read-only binding to the contract
	UpgradableECCMTransactor // Write-only binding to the contract
	UpgradableECCMFilterer   // Log filterer for contract events
}

// UpgradableECCMCaller is an auto generated read-only Go binding around an Ethereum contract.
type UpgradableECCMCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// UpgradableECCMTransactor is an auto generated write-only Go binding around an Ethereum contract.
type UpgradableECCMTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// UpgradableECCMFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type UpgradableECCMFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// UpgradableECCMSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type UpgradableECCMSession struct {
	Contract     *UpgradableECCM   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// UpgradableECCMCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type UpgradableECCMCallerSession struct {
	Contract *UpgradableECCMCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// UpgradableECCMTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type UpgradableECCMTransactorSession struct {
	Contract     *UpgradableECCMTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// UpgradableECCMRaw is an auto generated low-level Go binding around an Ethereum contract.
type UpgradableECCMRaw struct {
	Contract *UpgradableECCM // Generic contract binding to access the raw methods on
}

// UpgradableECCMCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type UpgradableECCMCallerRaw struct {
	Contract *UpgradableECCMCaller // Generic read-only contract binding to access the raw methods on
}

// UpgradableECCMTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type UpgradableECCMTransactorRaw struct {
	Contract *UpgradableECCMTransactor // Generic write-only contract binding to access the raw methods on
}

// NewUpgradableECCM creates a new instance of UpgradableECCM, bound to a specific deployed contract.
func NewUpgradableECCM(address common.Address, backend bind.ContractBackend) (*UpgradableECCM, error) {
	contract, err := bindUpgradableECCM(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &UpgradableECCM{UpgradableECCMCaller: UpgradableECCMCaller{contract: contract}, UpgradableECCMTransactor: UpgradableECCMTransactor{contract: contract}, UpgradableECCMFilterer: UpgradableECCMFilterer{contract: contract}}, nil
}

// NewUpgradableECCMCaller creates a new read-only instance of UpgradableECCM, bound to a specific deployed contract.
func NewUpgradableECCMCaller(address common.Address, caller bind.ContractCaller) (*UpgradableECCMCaller, error) {
	contract, err := bindUpgradableECCM(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &UpgradableECCMCaller{contract: contract}, nil
}

// NewUpgradableECCMTransactor creates a new write-only instance of UpgradableECCM, bound to a specific deployed contract.
func NewUpgradableECCMTransactor(address common.Address, transactor bind.ContractTransactor) (*UpgradableECCMTransactor, error) {
	contract, err := bindUpgradableECCM(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &UpgradableECCMTransactor{contract: contract}, nil
}

// NewUpgradableECCMFilterer creates a new log filterer instance of UpgradableECCM, bound to a specific deployed contract.
func NewUpgradableECCMFilterer(address common.Address, filterer bind.ContractFilterer) (*UpgradableECCMFilterer, error) {
	contract, err := bindUpgradableECCM(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &UpgradableECCMFilterer{contract: contract}, nil
}

// bindUpgradableECCM binds a generic wrapper to an already deployed contract.
func bindUpgradableECCM(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(UpgradableECCMABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_UpgradableECCM *UpgradableECCMRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _UpgradableECCM.Contract.UpgradableECCMCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_UpgradableECCM *UpgradableECCMRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _UpgradableECCM.Contract.UpgradableECCMTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_UpgradableECCM *UpgradableECCMRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _UpgradableECCM.Contract.UpgradableECCMTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_UpgradableECCM *UpgradableECCMCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _UpgradableECCM.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_UpgradableECCM *UpgradableECCMTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _UpgradableECCM.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_UpgradableECCM *UpgradableECCMTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _UpgradableECCM.Contract.contract.Transact(opts, method, params...)
}

// EthCrossChainDataAddress is a free data retrieval call binding the contract method 0x00ba1694.
//
// Solidity: function EthCrossChainDataAddress() constant returns(address)
func (_UpgradableECCM *UpgradableECCMCaller) EthCrossChainDataAddress(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _UpgradableECCM.contract.Call(opts, out, "EthCrossChainDataAddress")
	return *ret0, err
}

// EthCrossChainDataAddress is a free data retrieval call binding the contract method 0x00ba1694.
//
// Solidity: function EthCrossChainDataAddress() constant returns(address)
func (_UpgradableECCM *UpgradableECCMSession) EthCrossChainDataAddress() (common.Address, error) {
	return _UpgradableECCM.Contract.EthCrossChainDataAddress(&_UpgradableECCM.CallOpts)
}

// EthCrossChainDataAddress is a free data retrieval call binding the contract method 0x00ba1694.
//
// Solidity: function EthCrossChainDataAddress() constant returns(address)
func (_UpgradableECCM *UpgradableECCMCallerSession) EthCrossChainDataAddress() (common.Address, error) {
	return _UpgradableECCM.Contract.EthCrossChainDataAddress(&_UpgradableECCM.CallOpts)
}

// IsOwner is a free data retrieval call binding the contract method 0x8f32d59b.
//
// Solidity: function isOwner() constant returns(bool)
func (_UpgradableECCM *UpgradableECCMCaller) IsOwner(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _UpgradableECCM.contract.Call(opts, out, "isOwner")
	return *ret0, err
}

// IsOwner is a free data retrieval call binding the contract method 0x8f32d59b.
//
// Solidity: function isOwner() constant returns(bool)
func (_UpgradableECCM *UpgradableECCMSession) IsOwner() (bool, error) {
	return _UpgradableECCM.Contract.IsOwner(&_UpgradableECCM.CallOpts)
}

// IsOwner is a free data retrieval call binding the contract method 0x8f32d59b.
//
// Solidity: function isOwner() constant returns(bool)
func (_UpgradableECCM *UpgradableECCMCallerSession) IsOwner() (bool, error) {
	return _UpgradableECCM.Contract.IsOwner(&_UpgradableECCM.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_UpgradableECCM *UpgradableECCMCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _UpgradableECCM.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_UpgradableECCM *UpgradableECCMSession) Owner() (common.Address, error) {
	return _UpgradableECCM.Contract.Owner(&_UpgradableECCM.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_UpgradableECCM *UpgradableECCMCallerSession) Owner() (common.Address, error) {
	return _UpgradableECCM.Contract.Owner(&_UpgradableECCM.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() constant returns(bool)
func (_UpgradableECCM *UpgradableECCMCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _UpgradableECCM.contract.Call(opts, out, "paused")
	return *ret0, err
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() constant returns(bool)
func (_UpgradableECCM *UpgradableECCMSession) Paused() (bool, error) {
	return _UpgradableECCM.Contract.Paused(&_UpgradableECCM.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() constant returns(bool)
func (_UpgradableECCM *UpgradableECCMCallerSession) Paused() (bool, error) {
	return _UpgradableECCM.Contract.Paused(&_UpgradableECCM.CallOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns(bool)
func (_UpgradableECCM *UpgradableECCMTransactor) Pause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _UpgradableECCM.contract.Transact(opts, "pause")
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns(bool)
func (_UpgradableECCM *UpgradableECCMSession) Pause() (*types.Transaction, error) {
	return _UpgradableECCM.Contract.Pause(&_UpgradableECCM.TransactOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns(bool)
func (_UpgradableECCM *UpgradableECCMTransactorSession) Pause() (*types.Transaction, error) {
	return _UpgradableECCM.Contract.Pause(&_UpgradableECCM.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_UpgradableECCM *UpgradableECCMTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _UpgradableECCM.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_UpgradableECCM *UpgradableECCMSession) RenounceOwnership() (*types.Transaction, error) {
	return _UpgradableECCM.Contract.RenounceOwnership(&_UpgradableECCM.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_UpgradableECCM *UpgradableECCMTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _UpgradableECCM.Contract.RenounceOwnership(&_UpgradableECCM.TransactOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_UpgradableECCM *UpgradableECCMTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _UpgradableECCM.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_UpgradableECCM *UpgradableECCMSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _UpgradableECCM.Contract.TransferOwnership(&_UpgradableECCM.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_UpgradableECCM *UpgradableECCMTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _UpgradableECCM.Contract.TransferOwnership(&_UpgradableECCM.TransactOpts, newOwner)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns(bool)
func (_UpgradableECCM *UpgradableECCMTransactor) Unpause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _UpgradableECCM.contract.Transact(opts, "unpause")
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns(bool)
func (_UpgradableECCM *UpgradableECCMSession) Unpause() (*types.Transaction, error) {
	return _UpgradableECCM.Contract.Unpause(&_UpgradableECCM.TransactOpts)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns(bool)
func (_UpgradableECCM *UpgradableECCMTransactorSession) Unpause() (*types.Transaction, error) {
	return _UpgradableECCM.Contract.Unpause(&_UpgradableECCM.TransactOpts)
}

// UpgradeToNew is a paid mutator transaction binding the contract method 0x7e724ff3.
//
// Solidity: function upgradeToNew(address newEthCrossChainManagerAddress) returns(bool)
func (_UpgradableECCM *UpgradableECCMTransactor) UpgradeToNew(opts *bind.TransactOpts, newEthCrossChainManagerAddress common.Address) (*types.Transaction, error) {
	return _UpgradableECCM.contract.Transact(opts, "upgradeToNew", newEthCrossChainManagerAddress)
}

// UpgradeToNew is a paid mutator transaction binding the contract method 0x7e724ff3.
//
// Solidity: function upgradeToNew(address newEthCrossChainManagerAddress) returns(bool)
func (_UpgradableECCM *UpgradableECCMSession) UpgradeToNew(newEthCrossChainManagerAddress common.Address) (*types.Transaction, error) {
	return _UpgradableECCM.Contract.UpgradeToNew(&_UpgradableECCM.TransactOpts, newEthCrossChainManagerAddress)
}

// UpgradeToNew is a paid mutator transaction binding the contract method 0x7e724ff3.
//
// Solidity: function upgradeToNew(address newEthCrossChainManagerAddress) returns(bool)
func (_UpgradableECCM *UpgradableECCMTransactorSession) UpgradeToNew(newEthCrossChainManagerAddress common.Address) (*types.Transaction, error) {
	return _UpgradableECCM.Contract.UpgradeToNew(&_UpgradableECCM.TransactOpts, newEthCrossChainManagerAddress)
}

// UpgradableECCMOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the UpgradableECCM contract.
type UpgradableECCMOwnershipTransferredIterator struct {
	Event *UpgradableECCMOwnershipTransferred // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *UpgradableECCMOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(UpgradableECCMOwnershipTransferred)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(UpgradableECCMOwnershipTransferred)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *UpgradableECCMOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *UpgradableECCMOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// UpgradableECCMOwnershipTransferred represents a OwnershipTransferred event raised by the UpgradableECCM contract.
type UpgradableECCMOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_UpgradableECCM *UpgradableECCMFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*UpgradableECCMOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _UpgradableECCM.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &UpgradableECCMOwnershipTransferredIterator{contract: _UpgradableECCM.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_UpgradableECCM *UpgradableECCMFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *UpgradableECCMOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _UpgradableECCM.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(UpgradableECCMOwnershipTransferred)
				if err := _UpgradableECCM.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseOwnershipTransferred is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_UpgradableECCM *UpgradableECCMFilterer) ParseOwnershipTransferred(log types.Log) (*UpgradableECCMOwnershipTransferred, error) {
	event := new(UpgradableECCMOwnershipTransferred)
	if err := _UpgradableECCM.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	return event, nil
}

// UpgradableECCMPausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the UpgradableECCM contract.
type UpgradableECCMPausedIterator struct {
	Event *UpgradableECCMPaused // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *UpgradableECCMPausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(UpgradableECCMPaused)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(UpgradableECCMPaused)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *UpgradableECCMPausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *UpgradableECCMPausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// UpgradableECCMPaused represents a Paused event raised by the UpgradableECCM contract.
type UpgradableECCMPaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_UpgradableECCM *UpgradableECCMFilterer) FilterPaused(opts *bind.FilterOpts) (*UpgradableECCMPausedIterator, error) {

	logs, sub, err := _UpgradableECCM.contract.FilterLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return &UpgradableECCMPausedIterator{contract: _UpgradableECCM.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_UpgradableECCM *UpgradableECCMFilterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *UpgradableECCMPaused) (event.Subscription, error) {

	logs, sub, err := _UpgradableECCM.contract.WatchLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(UpgradableECCMPaused)
				if err := _UpgradableECCM.contract.UnpackLog(event, "Paused", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParsePaused is a log parse operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_UpgradableECCM *UpgradableECCMFilterer) ParsePaused(log types.Log) (*UpgradableECCMPaused, error) {
	event := new(UpgradableECCMPaused)
	if err := _UpgradableECCM.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	return event, nil
}

// UpgradableECCMUnpausedIterator is returned from FilterUnpaused and is used to iterate over the raw logs and unpacked data for Unpaused events raised by the UpgradableECCM contract.
type UpgradableECCMUnpausedIterator struct {
	Event *UpgradableECCMUnpaused // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *UpgradableECCMUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(UpgradableECCMUnpaused)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(UpgradableECCMUnpaused)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *UpgradableECCMUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *UpgradableECCMUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// UpgradableECCMUnpaused represents a Unpaused event raised by the UpgradableECCM contract.
type UpgradableECCMUnpaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterUnpaused is a free log retrieval operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_UpgradableECCM *UpgradableECCMFilterer) FilterUnpaused(opts *bind.FilterOpts) (*UpgradableECCMUnpausedIterator, error) {

	logs, sub, err := _UpgradableECCM.contract.FilterLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return &UpgradableECCMUnpausedIterator{contract: _UpgradableECCM.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

// WatchUnpaused is a free log subscription operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_UpgradableECCM *UpgradableECCMFilterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *UpgradableECCMUnpaused) (event.Subscription, error) {

	logs, sub, err := _UpgradableECCM.contract.WatchLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(UpgradableECCMUnpaused)
				if err := _UpgradableECCM.contract.UnpackLog(event, "Unpaused", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseUnpaused is a log parse operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_UpgradableECCM *UpgradableECCMFilterer) ParseUnpaused(log types.Log) (*UpgradableECCMUnpaused, error) {
	event := new(UpgradableECCMUnpaused)
	if err := _UpgradableECCM.contract.UnpackLog(event, "Unpaused", log); err != nil {
		return nil, err
	}
	return event, nil
}

// UtilsABI is the input ABI used to generate the binding from.
const UtilsABI = "[]"

// UtilsBin is the compiled bytecode used for deploying new contracts.
var UtilsBin = "0x60556023600b82828239805160001a607314601657fe5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea265627a7a723158205ac777cd5dc1521629c36b8c97e6b77192025a5ec6c13156be0d5222cfa8efaf64736f6c634300050f0032"

// DeployUtils deploys a new Ethereum contract, binding an instance of Utils to it.
func DeployUtils(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Utils, error) {
	parsed, err := abi.JSON(strings.NewReader(UtilsABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(UtilsBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Utils{UtilsCaller: UtilsCaller{contract: contract}, UtilsTransactor: UtilsTransactor{contract: contract}, UtilsFilterer: UtilsFilterer{contract: contract}}, nil
}

// Utils is an auto generated Go binding around an Ethereum contract.
type Utils struct {
	UtilsCaller     // Read-only binding to the contract
	UtilsTransactor // Write-only binding to the contract
	UtilsFilterer   // Log filterer for contract events
}

// UtilsCaller is an auto generated read-only Go binding around an Ethereum contract.
type UtilsCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// UtilsTransactor is an auto generated write-only Go binding around an Ethereum contract.
type UtilsTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// UtilsFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type UtilsFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// UtilsSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type UtilsSession struct {
	Contract     *Utils            // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// UtilsCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type UtilsCallerSession struct {
	Contract *UtilsCaller  // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// UtilsTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type UtilsTransactorSession struct {
	Contract     *UtilsTransactor  // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// UtilsRaw is an auto generated low-level Go binding around an Ethereum contract.
type UtilsRaw struct {
	Contract *Utils // Generic contract binding to access the raw methods on
}

// UtilsCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type UtilsCallerRaw struct {
	Contract *UtilsCaller // Generic read-only contract binding to access the raw methods on
}

// UtilsTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type UtilsTransactorRaw struct {
	Contract *UtilsTransactor // Generic write-only contract binding to access the raw methods on
}

// NewUtils creates a new instance of Utils, bound to a specific deployed contract.
func NewUtils(address common.Address, backend bind.ContractBackend) (*Utils, error) {
	contract, err := bindUtils(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Utils{UtilsCaller: UtilsCaller{contract: contract}, UtilsTransactor: UtilsTransactor{contract: contract}, UtilsFilterer: UtilsFilterer{contract: contract}}, nil
}

// NewUtilsCaller creates a new read-only instance of Utils, bound to a specific deployed contract.
func NewUtilsCaller(address common.Address, caller bind.ContractCaller) (*UtilsCaller, error) {
	contract, err := bindUtils(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &UtilsCaller{contract: contract}, nil
}

// NewUtilsTransactor creates a new write-only instance of Utils, bound to a specific deployed contract.
func NewUtilsTransactor(address common.Address, transactor bind.ContractTransactor) (*UtilsTransactor, error) {
	contract, err := bindUtils(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &UtilsTransactor{contract: contract}, nil
}

// NewUtilsFilterer creates a new log filterer instance of Utils, bound to a specific deployed contract.
func NewUtilsFilterer(address common.Address, filterer bind.ContractFilterer) (*UtilsFilterer, error) {
	contract, err := bindUtils(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &UtilsFilterer{contract: contract}, nil
}

// bindUtils binds a generic wrapper to an already deployed contract.
func bindUtils(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(UtilsABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Utils *UtilsRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Utils.Contract.UtilsCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Utils *UtilsRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Utils.Contract.UtilsTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Utils *UtilsRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Utils.Contract.UtilsTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Utils *UtilsCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Utils.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Utils *UtilsTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Utils.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Utils *UtilsTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Utils.Contract.contract.Transact(opts, method, params...)
}

// ZeroCopySinkABI is the input ABI used to generate the binding from.
const ZeroCopySinkABI = "[]"

// ZeroCopySinkBin is the compiled bytecode used for deploying new contracts.
var ZeroCopySinkBin = "0x60556023600b82828239805160001a607314601657fe5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea265627a7a723158208eb43bf0e592a4613bb9a63840a985e6372c70c4140c3ac6270fbc9672836ff064736f6c634300050f0032"

// DeployZeroCopySink deploys a new Ethereum contract, binding an instance of ZeroCopySink to it.
func DeployZeroCopySink(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *ZeroCopySink, error) {
	parsed, err := abi.JSON(strings.NewReader(ZeroCopySinkABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(ZeroCopySinkBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ZeroCopySink{ZeroCopySinkCaller: ZeroCopySinkCaller{contract: contract}, ZeroCopySinkTransactor: ZeroCopySinkTransactor{contract: contract}, ZeroCopySinkFilterer: ZeroCopySinkFilterer{contract: contract}}, nil
}

// ZeroCopySink is an auto generated Go binding around an Ethereum contract.
type ZeroCopySink struct {
	ZeroCopySinkCaller     // Read-only binding to the contract
	ZeroCopySinkTransactor // Write-only binding to the contract
	ZeroCopySinkFilterer   // Log filterer for contract events
}

// ZeroCopySinkCaller is an auto generated read-only Go binding around an Ethereum contract.
type ZeroCopySinkCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ZeroCopySinkTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ZeroCopySinkTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ZeroCopySinkFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ZeroCopySinkFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ZeroCopySinkSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ZeroCopySinkSession struct {
	Contract     *ZeroCopySink     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ZeroCopySinkCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ZeroCopySinkCallerSession struct {
	Contract *ZeroCopySinkCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// ZeroCopySinkTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ZeroCopySinkTransactorSession struct {
	Contract     *ZeroCopySinkTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// ZeroCopySinkRaw is an auto generated low-level Go binding around an Ethereum contract.
type ZeroCopySinkRaw struct {
	Contract *ZeroCopySink // Generic contract binding to access the raw methods on
}

// ZeroCopySinkCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ZeroCopySinkCallerRaw struct {
	Contract *ZeroCopySinkCaller // Generic read-only contract binding to access the raw methods on
}

// ZeroCopySinkTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ZeroCopySinkTransactorRaw struct {
	Contract *ZeroCopySinkTransactor // Generic write-only contract binding to access the raw methods on
}

// NewZeroCopySink creates a new instance of ZeroCopySink, bound to a specific deployed contract.
func NewZeroCopySink(address common.Address, backend bind.ContractBackend) (*ZeroCopySink, error) {
	contract, err := bindZeroCopySink(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ZeroCopySink{ZeroCopySinkCaller: ZeroCopySinkCaller{contract: contract}, ZeroCopySinkTransactor: ZeroCopySinkTransactor{contract: contract}, ZeroCopySinkFilterer: ZeroCopySinkFilterer{contract: contract}}, nil
}

// NewZeroCopySinkCaller creates a new read-only instance of ZeroCopySink, bound to a specific deployed contract.
func NewZeroCopySinkCaller(address common.Address, caller bind.ContractCaller) (*ZeroCopySinkCaller, error) {
	contract, err := bindZeroCopySink(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ZeroCopySinkCaller{contract: contract}, nil
}

// NewZeroCopySinkTransactor creates a new write-only instance of ZeroCopySink, bound to a specific deployed contract.
func NewZeroCopySinkTransactor(address common.Address, transactor bind.ContractTransactor) (*ZeroCopySinkTransactor, error) {
	contract, err := bindZeroCopySink(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ZeroCopySinkTransactor{contract: contract}, nil
}

// NewZeroCopySinkFilterer creates a new log filterer instance of ZeroCopySink, bound to a specific deployed contract.
func NewZeroCopySinkFilterer(address common.Address, filterer bind.ContractFilterer) (*ZeroCopySinkFilterer, error) {
	contract, err := bindZeroCopySink(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ZeroCopySinkFilterer{contract: contract}, nil
}

// bindZeroCopySink binds a generic wrapper to an already deployed contract.
func bindZeroCopySink(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ZeroCopySinkABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ZeroCopySink *ZeroCopySinkRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ZeroCopySink.Contract.ZeroCopySinkCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ZeroCopySink *ZeroCopySinkRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ZeroCopySink.Contract.ZeroCopySinkTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ZeroCopySink *ZeroCopySinkRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ZeroCopySink.Contract.ZeroCopySinkTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ZeroCopySink *ZeroCopySinkCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ZeroCopySink.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ZeroCopySink *ZeroCopySinkTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ZeroCopySink.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ZeroCopySink *ZeroCopySinkTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ZeroCopySink.Contract.contract.Transact(opts, method, params...)
}

// ZeroCopySourceABI is the input ABI used to generate the binding from.
const ZeroCopySourceABI = "[]"

// ZeroCopySourceBin is the compiled bytecode used for deploying new contracts.
var ZeroCopySourceBin = "0x60556023600b82828239805160001a607314601657fe5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea265627a7a7231582086f0b148ea78993fee1adbc0d4b28b84bcd5451020e8bae5a43911ed0f0d55c364736f6c634300050f0032"

// DeployZeroCopySource deploys a new Ethereum contract, binding an instance of ZeroCopySource to it.
func DeployZeroCopySource(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *ZeroCopySource, error) {
	parsed, err := abi.JSON(strings.NewReader(ZeroCopySourceABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(ZeroCopySourceBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ZeroCopySource{ZeroCopySourceCaller: ZeroCopySourceCaller{contract: contract}, ZeroCopySourceTransactor: ZeroCopySourceTransactor{contract: contract}, ZeroCopySourceFilterer: ZeroCopySourceFilterer{contract: contract}}, nil
}

// ZeroCopySource is an auto generated Go binding around an Ethereum contract.
type ZeroCopySource struct {
	ZeroCopySourceCaller     // Read-only binding to the contract
	ZeroCopySourceTransactor // Write-only binding to the contract
	ZeroCopySourceFilterer   // Log filterer for contract events
}

// ZeroCopySourceCaller is an auto generated read-only Go binding around an Ethereum contract.
type ZeroCopySourceCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ZeroCopySourceTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ZeroCopySourceTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ZeroCopySourceFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ZeroCopySourceFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ZeroCopySourceSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ZeroCopySourceSession struct {
	Contract     *ZeroCopySource   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ZeroCopySourceCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ZeroCopySourceCallerSession struct {
	Contract *ZeroCopySourceCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// ZeroCopySourceTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ZeroCopySourceTransactorSession struct {
	Contract     *ZeroCopySourceTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// ZeroCopySourceRaw is an auto generated low-level Go binding around an Ethereum contract.
type ZeroCopySourceRaw struct {
	Contract *ZeroCopySource // Generic contract binding to access the raw methods on
}

// ZeroCopySourceCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ZeroCopySourceCallerRaw struct {
	Contract *ZeroCopySourceCaller // Generic read-only contract binding to access the raw methods on
}

// ZeroCopySourceTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ZeroCopySourceTransactorRaw struct {
	Contract *ZeroCopySourceTransactor // Generic write-only contract binding to access the raw methods on
}

// NewZeroCopySource creates a new instance of ZeroCopySource, bound to a specific deployed contract.
func NewZeroCopySource(address common.Address, backend bind.ContractBackend) (*ZeroCopySource, error) {
	contract, err := bindZeroCopySource(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ZeroCopySource{ZeroCopySourceCaller: ZeroCopySourceCaller{contract: contract}, ZeroCopySourceTransactor: ZeroCopySourceTransactor{contract: contract}, ZeroCopySourceFilterer: ZeroCopySourceFilterer{contract: contract}}, nil
}

// NewZeroCopySourceCaller creates a new read-only instance of ZeroCopySource, bound to a specific deployed contract.
func NewZeroCopySourceCaller(address common.Address, caller bind.ContractCaller) (*ZeroCopySourceCaller, error) {
	contract, err := bindZeroCopySource(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ZeroCopySourceCaller{contract: contract}, nil
}

// NewZeroCopySourceTransactor creates a new write-only instance of ZeroCopySource, bound to a specific deployed contract.
func NewZeroCopySourceTransactor(address common.Address, transactor bind.ContractTransactor) (*ZeroCopySourceTransactor, error) {
	contract, err := bindZeroCopySource(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ZeroCopySourceTransactor{contract: contract}, nil
}

// NewZeroCopySourceFilterer creates a new log filterer instance of ZeroCopySource, bound to a specific deployed contract.
func NewZeroCopySourceFilterer(address common.Address, filterer bind.ContractFilterer) (*ZeroCopySourceFilterer, error) {
	contract, err := bindZeroCopySource(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ZeroCopySourceFilterer{contract: contract}, nil
}

// bindZeroCopySource binds a generic wrapper to an already deployed contract.
func bindZeroCopySource(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ZeroCopySourceABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ZeroCopySource *ZeroCopySourceRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ZeroCopySource.Contract.ZeroCopySourceCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ZeroCopySource *ZeroCopySourceRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ZeroCopySource.Contract.ZeroCopySourceTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ZeroCopySource *ZeroCopySourceRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ZeroCopySource.Contract.ZeroCopySourceTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ZeroCopySource *ZeroCopySourceCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ZeroCopySource.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ZeroCopySource *ZeroCopySourceTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ZeroCopySource.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ZeroCopySource *ZeroCopySourceTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ZeroCopySource.Contract.contract.Transact(opts, method, params...)
}

