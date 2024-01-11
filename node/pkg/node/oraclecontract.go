// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package node

import (
	"errors"
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
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
	_ = abi.ConvertType
)

// OracleContractOracleNode is an auto generated low-level Go binding around an user-defined struct.
type OracleContractOracleNode struct {
	Addr       common.Address
	IpAddr     string
	PubKeys    [][2]*big.Int
	BlsPubKeys [][4]*big.Int
	Stake      *big.Int
	Rank       *big.Int
	Index      *big.Int
}

// OracleContractMetaData contains all meta data concerning the OracleContract contract.
var OracleContractMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RegisterOracleNode\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"enumOracleContract.ValidationType\",\"name\":\"typ\",\"type\":\"uint8\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"hash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"size\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"minRank\",\"type\":\"uint256\"}],\"name\":\"ValidationRequest\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"AGGREGATE_FEE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"BASE_FEE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MIN_STAKE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"PUBKEY_LENGTH\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"countOracleNodes\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"deleteNode\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_addr\",\"type\":\"address\"}],\"name\":\"findOracleNodeByAddress\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"ipAddr\",\"type\":\"string\"},{\"internalType\":\"uint256[2][]\",\"name\":\"pubKeys\",\"type\":\"uint256[2][]\"},{\"internalType\":\"uint256[4][]\",\"name\":\"blsPubKeys\",\"type\":\"uint256[4][]\"},{\"internalType\":\"uint256\",\"name\":\"stake\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"rank\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"internalType\":\"structOracleContract.OracleNode\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_index\",\"type\":\"uint256\"}],\"name\":\"findOracleNodeByIndex\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"ipAddr\",\"type\":\"string\"},{\"internalType\":\"uint256[2][]\",\"name\":\"pubKeys\",\"type\":\"uint256[2][]\"},{\"internalType\":\"uint256[4][]\",\"name\":\"blsPubKeys\",\"type\":\"uint256[4][]\"},{\"internalType\":\"uint256\",\"name\":\"stake\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"rank\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"internalType\":\"structOracleContract.OracleNode\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getAggregator\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getAggregatorIP\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"getNodeBLSPublicKeys\",\"outputs\":[{\"internalType\":\"uint256[4][]\",\"name\":\"\",\"type\":\"uint256[4][]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getNodeBLSPublicKeysSum\",\"outputs\":[{\"internalType\":\"uint256[4]\",\"name\":\"\",\"type\":\"uint256[4]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"getNodePublicKeys\",\"outputs\":[{\"internalType\":\"uint256[2][]\",\"name\":\"\",\"type\":\"uint256[2][]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"getNodeRank\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_addr\",\"type\":\"address\"}],\"name\":\"isAggregator\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_addr\",\"type\":\"address\"}],\"name\":\"oracleNodeIsRegistered\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"oracleNodes\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"ipAddr\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"stake\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"rank\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_ipAddr\",\"type\":\"string\"},{\"internalType\":\"uint256[2][]\",\"name\":\"_pubKey\",\"type\":\"uint256[2][]\"},{\"internalType\":\"uint256\",\"name\":\"rank\",\"type\":\"uint256\"}],\"name\":\"registerOracleNode\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"_result\",\"type\":\"bool\"},{\"internalType\":\"bytes32\",\"name\":\"message\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"signature\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"rx\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"ry\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_hash\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"validators\",\"type\":\"address[]\"}],\"name\":\"submitTransactionValidationResult\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"enumOracleContract.ValidationType\",\"name\":\"_typ\",\"type\":\"uint8\"},{\"internalType\":\"bool\",\"name\":\"_result\",\"type\":\"bool\"},{\"internalType\":\"bytes32\",\"name\":\"message\",\"type\":\"bytes32\"},{\"internalType\":\"uint256[4]\",\"name\":\"publicKey\",\"type\":\"uint256[4]\"},{\"internalType\":\"uint256[2]\",\"name\":\"_signature\",\"type\":\"uint256[2]\"},{\"internalType\":\"uint256[2]\",\"name\":\"_hash\",\"type\":\"uint256[2]\"},{\"internalType\":\"address[]\",\"name\":\"validators\",\"type\":\"address[]\"}],\"name\":\"submitValidationResultBLS\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"size\",\"type\":\"uint256\"}],\"name\":\"totalFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"unregisterAddr\",\"type\":\"address\"}],\"name\":\"unregister\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_message\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"size\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minRank\",\"type\":\"uint256\"}],\"name\":\"validateTransaction\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"}]",
}

// OracleContractABI is the input ABI used to generate the binding from.
// Deprecated: Use OracleContractMetaData.ABI instead.
var OracleContractABI = OracleContractMetaData.ABI

// OracleContract is an auto generated Go binding around an Ethereum contract.
type OracleContract struct {
	OracleContractCaller     // Read-only binding to the contract
	OracleContractTransactor // Write-only binding to the contract
	OracleContractFilterer   // Log filterer for contract events
}

// OracleContractCaller is an auto generated read-only Go binding around an Ethereum contract.
type OracleContractCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OracleContractTransactor is an auto generated write-only Go binding around an Ethereum contract.
type OracleContractTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OracleContractFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type OracleContractFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OracleContractSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type OracleContractSession struct {
	Contract     *OracleContract   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// OracleContractCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type OracleContractCallerSession struct {
	Contract *OracleContractCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// OracleContractTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type OracleContractTransactorSession struct {
	Contract     *OracleContractTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// OracleContractRaw is an auto generated low-level Go binding around an Ethereum contract.
type OracleContractRaw struct {
	Contract *OracleContract // Generic contract binding to access the raw methods on
}

// OracleContractCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type OracleContractCallerRaw struct {
	Contract *OracleContractCaller // Generic read-only contract binding to access the raw methods on
}

// OracleContractTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type OracleContractTransactorRaw struct {
	Contract *OracleContractTransactor // Generic write-only contract binding to access the raw methods on
}

// NewOracleContract creates a new instance of OracleContract, bound to a specific deployed contract.
func NewOracleContract(address common.Address, backend bind.ContractBackend) (*OracleContract, error) {
	contract, err := bindOracleContract(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &OracleContract{OracleContractCaller: OracleContractCaller{contract: contract}, OracleContractTransactor: OracleContractTransactor{contract: contract}, OracleContractFilterer: OracleContractFilterer{contract: contract}}, nil
}

// NewOracleContractCaller creates a new read-only instance of OracleContract, bound to a specific deployed contract.
func NewOracleContractCaller(address common.Address, caller bind.ContractCaller) (*OracleContractCaller, error) {
	contract, err := bindOracleContract(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &OracleContractCaller{contract: contract}, nil
}

// NewOracleContractTransactor creates a new write-only instance of OracleContract, bound to a specific deployed contract.
func NewOracleContractTransactor(address common.Address, transactor bind.ContractTransactor) (*OracleContractTransactor, error) {
	contract, err := bindOracleContract(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &OracleContractTransactor{contract: contract}, nil
}

// NewOracleContractFilterer creates a new log filterer instance of OracleContract, bound to a specific deployed contract.
func NewOracleContractFilterer(address common.Address, filterer bind.ContractFilterer) (*OracleContractFilterer, error) {
	contract, err := bindOracleContract(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &OracleContractFilterer{contract: contract}, nil
}

// bindOracleContract binds a generic wrapper to an already deployed contract.
func bindOracleContract(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := OracleContractMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_OracleContract *OracleContractRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _OracleContract.Contract.OracleContractCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_OracleContract *OracleContractRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OracleContract.Contract.OracleContractTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_OracleContract *OracleContractRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _OracleContract.Contract.OracleContractTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_OracleContract *OracleContractCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _OracleContract.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_OracleContract *OracleContractTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OracleContract.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_OracleContract *OracleContractTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _OracleContract.Contract.contract.Transact(opts, method, params...)
}

// AGGREGATEFEE is a free data retrieval call binding the contract method 0x1471866c.
//
// Solidity: function AGGREGATE_FEE() view returns(uint256)
func (_OracleContract *OracleContractCaller) AGGREGATEFEE(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _OracleContract.contract.Call(opts, &out, "AGGREGATE_FEE")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// AGGREGATEFEE is a free data retrieval call binding the contract method 0x1471866c.
//
// Solidity: function AGGREGATE_FEE() view returns(uint256)
func (_OracleContract *OracleContractSession) AGGREGATEFEE() (*big.Int, error) {
	return _OracleContract.Contract.AGGREGATEFEE(&_OracleContract.CallOpts)
}

// AGGREGATEFEE is a free data retrieval call binding the contract method 0x1471866c.
//
// Solidity: function AGGREGATE_FEE() view returns(uint256)
func (_OracleContract *OracleContractCallerSession) AGGREGATEFEE() (*big.Int, error) {
	return _OracleContract.Contract.AGGREGATEFEE(&_OracleContract.CallOpts)
}

// BASEFEE is a free data retrieval call binding the contract method 0x3d18651e.
//
// Solidity: function BASE_FEE() view returns(uint256)
func (_OracleContract *OracleContractCaller) BASEFEE(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _OracleContract.contract.Call(opts, &out, "BASE_FEE")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BASEFEE is a free data retrieval call binding the contract method 0x3d18651e.
//
// Solidity: function BASE_FEE() view returns(uint256)
func (_OracleContract *OracleContractSession) BASEFEE() (*big.Int, error) {
	return _OracleContract.Contract.BASEFEE(&_OracleContract.CallOpts)
}

// BASEFEE is a free data retrieval call binding the contract method 0x3d18651e.
//
// Solidity: function BASE_FEE() view returns(uint256)
func (_OracleContract *OracleContractCallerSession) BASEFEE() (*big.Int, error) {
	return _OracleContract.Contract.BASEFEE(&_OracleContract.CallOpts)
}

// MINSTAKE is a free data retrieval call binding the contract method 0xcb1c2b5c.
//
// Solidity: function MIN_STAKE() view returns(uint256)
func (_OracleContract *OracleContractCaller) MINSTAKE(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _OracleContract.contract.Call(opts, &out, "MIN_STAKE")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MINSTAKE is a free data retrieval call binding the contract method 0xcb1c2b5c.
//
// Solidity: function MIN_STAKE() view returns(uint256)
func (_OracleContract *OracleContractSession) MINSTAKE() (*big.Int, error) {
	return _OracleContract.Contract.MINSTAKE(&_OracleContract.CallOpts)
}

// MINSTAKE is a free data retrieval call binding the contract method 0xcb1c2b5c.
//
// Solidity: function MIN_STAKE() view returns(uint256)
func (_OracleContract *OracleContractCallerSession) MINSTAKE() (*big.Int, error) {
	return _OracleContract.Contract.MINSTAKE(&_OracleContract.CallOpts)
}

// PUBKEYLENGTH is a free data retrieval call binding the contract method 0xa4d55d1d.
//
// Solidity: function PUBKEY_LENGTH() view returns(uint256)
func (_OracleContract *OracleContractCaller) PUBKEYLENGTH(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _OracleContract.contract.Call(opts, &out, "PUBKEY_LENGTH")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PUBKEYLENGTH is a free data retrieval call binding the contract method 0xa4d55d1d.
//
// Solidity: function PUBKEY_LENGTH() view returns(uint256)
func (_OracleContract *OracleContractSession) PUBKEYLENGTH() (*big.Int, error) {
	return _OracleContract.Contract.PUBKEYLENGTH(&_OracleContract.CallOpts)
}

// PUBKEYLENGTH is a free data retrieval call binding the contract method 0xa4d55d1d.
//
// Solidity: function PUBKEY_LENGTH() view returns(uint256)
func (_OracleContract *OracleContractCallerSession) PUBKEYLENGTH() (*big.Int, error) {
	return _OracleContract.Contract.PUBKEYLENGTH(&_OracleContract.CallOpts)
}

// CountOracleNodes is a free data retrieval call binding the contract method 0x836f187a.
//
// Solidity: function countOracleNodes() view returns(uint256)
func (_OracleContract *OracleContractCaller) CountOracleNodes(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _OracleContract.contract.Call(opts, &out, "countOracleNodes")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CountOracleNodes is a free data retrieval call binding the contract method 0x836f187a.
//
// Solidity: function countOracleNodes() view returns(uint256)
func (_OracleContract *OracleContractSession) CountOracleNodes() (*big.Int, error) {
	return _OracleContract.Contract.CountOracleNodes(&_OracleContract.CallOpts)
}

// CountOracleNodes is a free data retrieval call binding the contract method 0x836f187a.
//
// Solidity: function countOracleNodes() view returns(uint256)
func (_OracleContract *OracleContractCallerSession) CountOracleNodes() (*big.Int, error) {
	return _OracleContract.Contract.CountOracleNodes(&_OracleContract.CallOpts)
}

// FindOracleNodeByAddress is a free data retrieval call binding the contract method 0x655a102f.
//
// Solidity: function findOracleNodeByAddress(address _addr) view returns((address,string,uint256[2][],uint256[4][],uint256,uint256,uint256))
func (_OracleContract *OracleContractCaller) FindOracleNodeByAddress(opts *bind.CallOpts, _addr common.Address) (OracleContractOracleNode, error) {
	var out []interface{}
	err := _OracleContract.contract.Call(opts, &out, "findOracleNodeByAddress", _addr)

	if err != nil {
		return *new(OracleContractOracleNode), err
	}

	out0 := *abi.ConvertType(out[0], new(OracleContractOracleNode)).(*OracleContractOracleNode)

	return out0, err

}

// FindOracleNodeByAddress is a free data retrieval call binding the contract method 0x655a102f.
//
// Solidity: function findOracleNodeByAddress(address _addr) view returns((address,string,uint256[2][],uint256[4][],uint256,uint256,uint256))
func (_OracleContract *OracleContractSession) FindOracleNodeByAddress(_addr common.Address) (OracleContractOracleNode, error) {
	return _OracleContract.Contract.FindOracleNodeByAddress(&_OracleContract.CallOpts, _addr)
}

// FindOracleNodeByAddress is a free data retrieval call binding the contract method 0x655a102f.
//
// Solidity: function findOracleNodeByAddress(address _addr) view returns((address,string,uint256[2][],uint256[4][],uint256,uint256,uint256))
func (_OracleContract *OracleContractCallerSession) FindOracleNodeByAddress(_addr common.Address) (OracleContractOracleNode, error) {
	return _OracleContract.Contract.FindOracleNodeByAddress(&_OracleContract.CallOpts, _addr)
}

// FindOracleNodeByIndex is a free data retrieval call binding the contract method 0x272132e9.
//
// Solidity: function findOracleNodeByIndex(uint256 _index) view returns((address,string,uint256[2][],uint256[4][],uint256,uint256,uint256))
func (_OracleContract *OracleContractCaller) FindOracleNodeByIndex(opts *bind.CallOpts, _index *big.Int) (OracleContractOracleNode, error) {
	var out []interface{}
	err := _OracleContract.contract.Call(opts, &out, "findOracleNodeByIndex", _index)

	if err != nil {
		return *new(OracleContractOracleNode), err
	}

	out0 := *abi.ConvertType(out[0], new(OracleContractOracleNode)).(*OracleContractOracleNode)

	return out0, err

}

// FindOracleNodeByIndex is a free data retrieval call binding the contract method 0x272132e9.
//
// Solidity: function findOracleNodeByIndex(uint256 _index) view returns((address,string,uint256[2][],uint256[4][],uint256,uint256,uint256))
func (_OracleContract *OracleContractSession) FindOracleNodeByIndex(_index *big.Int) (OracleContractOracleNode, error) {
	return _OracleContract.Contract.FindOracleNodeByIndex(&_OracleContract.CallOpts, _index)
}

// FindOracleNodeByIndex is a free data retrieval call binding the contract method 0x272132e9.
//
// Solidity: function findOracleNodeByIndex(uint256 _index) view returns((address,string,uint256[2][],uint256[4][],uint256,uint256,uint256))
func (_OracleContract *OracleContractCallerSession) FindOracleNodeByIndex(_index *big.Int) (OracleContractOracleNode, error) {
	return _OracleContract.Contract.FindOracleNodeByIndex(&_OracleContract.CallOpts, _index)
}

// GetAggregator is a free data retrieval call binding the contract method 0x3ad59dbc.
//
// Solidity: function getAggregator() view returns(address)
func (_OracleContract *OracleContractCaller) GetAggregator(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _OracleContract.contract.Call(opts, &out, "getAggregator")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetAggregator is a free data retrieval call binding the contract method 0x3ad59dbc.
//
// Solidity: function getAggregator() view returns(address)
func (_OracleContract *OracleContractSession) GetAggregator() (common.Address, error) {
	return _OracleContract.Contract.GetAggregator(&_OracleContract.CallOpts)
}

// GetAggregator is a free data retrieval call binding the contract method 0x3ad59dbc.
//
// Solidity: function getAggregator() view returns(address)
func (_OracleContract *OracleContractCallerSession) GetAggregator() (common.Address, error) {
	return _OracleContract.Contract.GetAggregator(&_OracleContract.CallOpts)
}

// GetAggregatorIP is a free data retrieval call binding the contract method 0x5b217907.
//
// Solidity: function getAggregatorIP() view returns(string)
func (_OracleContract *OracleContractCaller) GetAggregatorIP(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _OracleContract.contract.Call(opts, &out, "getAggregatorIP")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// GetAggregatorIP is a free data retrieval call binding the contract method 0x5b217907.
//
// Solidity: function getAggregatorIP() view returns(string)
func (_OracleContract *OracleContractSession) GetAggregatorIP() (string, error) {
	return _OracleContract.Contract.GetAggregatorIP(&_OracleContract.CallOpts)
}

// GetAggregatorIP is a free data retrieval call binding the contract method 0x5b217907.
//
// Solidity: function getAggregatorIP() view returns(string)
func (_OracleContract *OracleContractCallerSession) GetAggregatorIP() (string, error) {
	return _OracleContract.Contract.GetAggregatorIP(&_OracleContract.CallOpts)
}

// GetNodeBLSPublicKeys is a free data retrieval call binding the contract method 0xd2758032.
//
// Solidity: function getNodeBLSPublicKeys(address addr) view returns(uint256[4][])
func (_OracleContract *OracleContractCaller) GetNodeBLSPublicKeys(opts *bind.CallOpts, addr common.Address) ([][4]*big.Int, error) {
	var out []interface{}
	err := _OracleContract.contract.Call(opts, &out, "getNodeBLSPublicKeys", addr)

	if err != nil {
		return *new([][4]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([][4]*big.Int)).(*[][4]*big.Int)

	return out0, err

}

// GetNodeBLSPublicKeys is a free data retrieval call binding the contract method 0xd2758032.
//
// Solidity: function getNodeBLSPublicKeys(address addr) view returns(uint256[4][])
func (_OracleContract *OracleContractSession) GetNodeBLSPublicKeys(addr common.Address) ([][4]*big.Int, error) {
	return _OracleContract.Contract.GetNodeBLSPublicKeys(&_OracleContract.CallOpts, addr)
}

// GetNodeBLSPublicKeys is a free data retrieval call binding the contract method 0xd2758032.
//
// Solidity: function getNodeBLSPublicKeys(address addr) view returns(uint256[4][])
func (_OracleContract *OracleContractCallerSession) GetNodeBLSPublicKeys(addr common.Address) ([][4]*big.Int, error) {
	return _OracleContract.Contract.GetNodeBLSPublicKeys(&_OracleContract.CallOpts, addr)
}

// GetNodeBLSPublicKeysSum is a free data retrieval call binding the contract method 0x7745e505.
//
// Solidity: function getNodeBLSPublicKeysSum() view returns(uint256[4])
func (_OracleContract *OracleContractCaller) GetNodeBLSPublicKeysSum(opts *bind.CallOpts) ([4]*big.Int, error) {
	var out []interface{}
	err := _OracleContract.contract.Call(opts, &out, "getNodeBLSPublicKeysSum")

	if err != nil {
		return *new([4]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([4]*big.Int)).(*[4]*big.Int)

	return out0, err

}

// GetNodeBLSPublicKeysSum is a free data retrieval call binding the contract method 0x7745e505.
//
// Solidity: function getNodeBLSPublicKeysSum() view returns(uint256[4])
func (_OracleContract *OracleContractSession) GetNodeBLSPublicKeysSum() ([4]*big.Int, error) {
	return _OracleContract.Contract.GetNodeBLSPublicKeysSum(&_OracleContract.CallOpts)
}

// GetNodeBLSPublicKeysSum is a free data retrieval call binding the contract method 0x7745e505.
//
// Solidity: function getNodeBLSPublicKeysSum() view returns(uint256[4])
func (_OracleContract *OracleContractCallerSession) GetNodeBLSPublicKeysSum() ([4]*big.Int, error) {
	return _OracleContract.Contract.GetNodeBLSPublicKeysSum(&_OracleContract.CallOpts)
}

// GetNodePublicKeys is a free data retrieval call binding the contract method 0xfb9fb18a.
//
// Solidity: function getNodePublicKeys(address addr) view returns(uint256[2][])
func (_OracleContract *OracleContractCaller) GetNodePublicKeys(opts *bind.CallOpts, addr common.Address) ([][2]*big.Int, error) {
	var out []interface{}
	err := _OracleContract.contract.Call(opts, &out, "getNodePublicKeys", addr)

	if err != nil {
		return *new([][2]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([][2]*big.Int)).(*[][2]*big.Int)

	return out0, err

}

// GetNodePublicKeys is a free data retrieval call binding the contract method 0xfb9fb18a.
//
// Solidity: function getNodePublicKeys(address addr) view returns(uint256[2][])
func (_OracleContract *OracleContractSession) GetNodePublicKeys(addr common.Address) ([][2]*big.Int, error) {
	return _OracleContract.Contract.GetNodePublicKeys(&_OracleContract.CallOpts, addr)
}

// GetNodePublicKeys is a free data retrieval call binding the contract method 0xfb9fb18a.
//
// Solidity: function getNodePublicKeys(address addr) view returns(uint256[2][])
func (_OracleContract *OracleContractCallerSession) GetNodePublicKeys(addr common.Address) ([][2]*big.Int, error) {
	return _OracleContract.Contract.GetNodePublicKeys(&_OracleContract.CallOpts, addr)
}

// GetNodeRank is a free data retrieval call binding the contract method 0xbc66e9cb.
//
// Solidity: function getNodeRank(address addr) view returns(uint256)
func (_OracleContract *OracleContractCaller) GetNodeRank(opts *bind.CallOpts, addr common.Address) (*big.Int, error) {
	var out []interface{}
	err := _OracleContract.contract.Call(opts, &out, "getNodeRank", addr)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetNodeRank is a free data retrieval call binding the contract method 0xbc66e9cb.
//
// Solidity: function getNodeRank(address addr) view returns(uint256)
func (_OracleContract *OracleContractSession) GetNodeRank(addr common.Address) (*big.Int, error) {
	return _OracleContract.Contract.GetNodeRank(&_OracleContract.CallOpts, addr)
}

// GetNodeRank is a free data retrieval call binding the contract method 0xbc66e9cb.
//
// Solidity: function getNodeRank(address addr) view returns(uint256)
func (_OracleContract *OracleContractCallerSession) GetNodeRank(addr common.Address) (*big.Int, error) {
	return _OracleContract.Contract.GetNodeRank(&_OracleContract.CallOpts, addr)
}

// IsAggregator is a free data retrieval call binding the contract method 0x1e8f3c95.
//
// Solidity: function isAggregator(address _addr) view returns(bool)
func (_OracleContract *OracleContractCaller) IsAggregator(opts *bind.CallOpts, _addr common.Address) (bool, error) {
	var out []interface{}
	err := _OracleContract.contract.Call(opts, &out, "isAggregator", _addr)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsAggregator is a free data retrieval call binding the contract method 0x1e8f3c95.
//
// Solidity: function isAggregator(address _addr) view returns(bool)
func (_OracleContract *OracleContractSession) IsAggregator(_addr common.Address) (bool, error) {
	return _OracleContract.Contract.IsAggregator(&_OracleContract.CallOpts, _addr)
}

// IsAggregator is a free data retrieval call binding the contract method 0x1e8f3c95.
//
// Solidity: function isAggregator(address _addr) view returns(bool)
func (_OracleContract *OracleContractCallerSession) IsAggregator(_addr common.Address) (bool, error) {
	return _OracleContract.Contract.IsAggregator(&_OracleContract.CallOpts, _addr)
}

// OracleNodeIsRegistered is a free data retrieval call binding the contract method 0x140f3daa.
//
// Solidity: function oracleNodeIsRegistered(address _addr) view returns(bool)
func (_OracleContract *OracleContractCaller) OracleNodeIsRegistered(opts *bind.CallOpts, _addr common.Address) (bool, error) {
	var out []interface{}
	err := _OracleContract.contract.Call(opts, &out, "oracleNodeIsRegistered", _addr)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// OracleNodeIsRegistered is a free data retrieval call binding the contract method 0x140f3daa.
//
// Solidity: function oracleNodeIsRegistered(address _addr) view returns(bool)
func (_OracleContract *OracleContractSession) OracleNodeIsRegistered(_addr common.Address) (bool, error) {
	return _OracleContract.Contract.OracleNodeIsRegistered(&_OracleContract.CallOpts, _addr)
}

// OracleNodeIsRegistered is a free data retrieval call binding the contract method 0x140f3daa.
//
// Solidity: function oracleNodeIsRegistered(address _addr) view returns(bool)
func (_OracleContract *OracleContractCallerSession) OracleNodeIsRegistered(_addr common.Address) (bool, error) {
	return _OracleContract.Contract.OracleNodeIsRegistered(&_OracleContract.CallOpts, _addr)
}

// OracleNodes is a free data retrieval call binding the contract method 0x2849463e.
//
// Solidity: function oracleNodes(address ) view returns(address addr, string ipAddr, uint256 stake, uint256 rank, uint256 index)
func (_OracleContract *OracleContractCaller) OracleNodes(opts *bind.CallOpts, arg0 common.Address) (struct {
	Addr   common.Address
	IpAddr string
	Stake  *big.Int
	Rank   *big.Int
	Index  *big.Int
}, error) {
	var out []interface{}
	err := _OracleContract.contract.Call(opts, &out, "oracleNodes", arg0)

	outstruct := new(struct {
		Addr   common.Address
		IpAddr string
		Stake  *big.Int
		Rank   *big.Int
		Index  *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Addr = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.IpAddr = *abi.ConvertType(out[1], new(string)).(*string)
	outstruct.Stake = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.Rank = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.Index = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// OracleNodes is a free data retrieval call binding the contract method 0x2849463e.
//
// Solidity: function oracleNodes(address ) view returns(address addr, string ipAddr, uint256 stake, uint256 rank, uint256 index)
func (_OracleContract *OracleContractSession) OracleNodes(arg0 common.Address) (struct {
	Addr   common.Address
	IpAddr string
	Stake  *big.Int
	Rank   *big.Int
	Index  *big.Int
}, error) {
	return _OracleContract.Contract.OracleNodes(&_OracleContract.CallOpts, arg0)
}

// OracleNodes is a free data retrieval call binding the contract method 0x2849463e.
//
// Solidity: function oracleNodes(address ) view returns(address addr, string ipAddr, uint256 stake, uint256 rank, uint256 index)
func (_OracleContract *OracleContractCallerSession) OracleNodes(arg0 common.Address) (struct {
	Addr   common.Address
	IpAddr string
	Stake  *big.Int
	Rank   *big.Int
	Index  *big.Int
}, error) {
	return _OracleContract.Contract.OracleNodes(&_OracleContract.CallOpts, arg0)
}

// TotalFee is a free data retrieval call binding the contract method 0xcea74edf.
//
// Solidity: function totalFee(uint256 size) pure returns(uint256)
func (_OracleContract *OracleContractCaller) TotalFee(opts *bind.CallOpts, size *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _OracleContract.contract.Call(opts, &out, "totalFee", size)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalFee is a free data retrieval call binding the contract method 0xcea74edf.
//
// Solidity: function totalFee(uint256 size) pure returns(uint256)
func (_OracleContract *OracleContractSession) TotalFee(size *big.Int) (*big.Int, error) {
	return _OracleContract.Contract.TotalFee(&_OracleContract.CallOpts, size)
}

// TotalFee is a free data retrieval call binding the contract method 0xcea74edf.
//
// Solidity: function totalFee(uint256 size) pure returns(uint256)
func (_OracleContract *OracleContractCallerSession) TotalFee(size *big.Int) (*big.Int, error) {
	return _OracleContract.Contract.TotalFee(&_OracleContract.CallOpts, size)
}

// DeleteNode is a paid mutator transaction binding the contract method 0x2d4ede93.
//
// Solidity: function deleteNode(address addr) returns()
func (_OracleContract *OracleContractTransactor) DeleteNode(opts *bind.TransactOpts, addr common.Address) (*types.Transaction, error) {
	return _OracleContract.contract.Transact(opts, "deleteNode", addr)
}

// DeleteNode is a paid mutator transaction binding the contract method 0x2d4ede93.
//
// Solidity: function deleteNode(address addr) returns()
func (_OracleContract *OracleContractSession) DeleteNode(addr common.Address) (*types.Transaction, error) {
	return _OracleContract.Contract.DeleteNode(&_OracleContract.TransactOpts, addr)
}

// DeleteNode is a paid mutator transaction binding the contract method 0x2d4ede93.
//
// Solidity: function deleteNode(address addr) returns()
func (_OracleContract *OracleContractTransactorSession) DeleteNode(addr common.Address) (*types.Transaction, error) {
	return _OracleContract.Contract.DeleteNode(&_OracleContract.TransactOpts, addr)
}

// RegisterOracleNode is a paid mutator transaction binding the contract method 0x16951440.
//
// Solidity: function registerOracleNode(string _ipAddr, uint256[2][] _pubKey, uint256 rank) payable returns()
func (_OracleContract *OracleContractTransactor) RegisterOracleNode(opts *bind.TransactOpts, _ipAddr string, _pubKey [][2]*big.Int, rank *big.Int) (*types.Transaction, error) {
	return _OracleContract.contract.Transact(opts, "registerOracleNode", _ipAddr, _pubKey, rank)
}

// RegisterOracleNode is a paid mutator transaction binding the contract method 0x16951440.
//
// Solidity: function registerOracleNode(string _ipAddr, uint256[2][] _pubKey, uint256 rank) payable returns()
func (_OracleContract *OracleContractSession) RegisterOracleNode(_ipAddr string, _pubKey [][2]*big.Int, rank *big.Int) (*types.Transaction, error) {
	return _OracleContract.Contract.RegisterOracleNode(&_OracleContract.TransactOpts, _ipAddr, _pubKey, rank)
}

// RegisterOracleNode is a paid mutator transaction binding the contract method 0x16951440.
//
// Solidity: function registerOracleNode(string _ipAddr, uint256[2][] _pubKey, uint256 rank) payable returns()
func (_OracleContract *OracleContractTransactorSession) RegisterOracleNode(_ipAddr string, _pubKey [][2]*big.Int, rank *big.Int) (*types.Transaction, error) {
	return _OracleContract.Contract.RegisterOracleNode(&_OracleContract.TransactOpts, _ipAddr, _pubKey, rank)
}

// SubmitTransactionValidationResult is a paid mutator transaction binding the contract method 0x9459142f.
//
// Solidity: function submitTransactionValidationResult(bool _result, bytes32 message, uint256 signature, uint256 rx, uint256 ry, uint256 _hash, address[] validators) returns()
func (_OracleContract *OracleContractTransactor) SubmitTransactionValidationResult(opts *bind.TransactOpts, _result bool, message [32]byte, signature *big.Int, rx *big.Int, ry *big.Int, _hash *big.Int, validators []common.Address) (*types.Transaction, error) {
	return _OracleContract.contract.Transact(opts, "submitTransactionValidationResult", _result, message, signature, rx, ry, _hash, validators)
}

// SubmitTransactionValidationResult is a paid mutator transaction binding the contract method 0x9459142f.
//
// Solidity: function submitTransactionValidationResult(bool _result, bytes32 message, uint256 signature, uint256 rx, uint256 ry, uint256 _hash, address[] validators) returns()
func (_OracleContract *OracleContractSession) SubmitTransactionValidationResult(_result bool, message [32]byte, signature *big.Int, rx *big.Int, ry *big.Int, _hash *big.Int, validators []common.Address) (*types.Transaction, error) {
	return _OracleContract.Contract.SubmitTransactionValidationResult(&_OracleContract.TransactOpts, _result, message, signature, rx, ry, _hash, validators)
}

// SubmitTransactionValidationResult is a paid mutator transaction binding the contract method 0x9459142f.
//
// Solidity: function submitTransactionValidationResult(bool _result, bytes32 message, uint256 signature, uint256 rx, uint256 ry, uint256 _hash, address[] validators) returns()
func (_OracleContract *OracleContractTransactorSession) SubmitTransactionValidationResult(_result bool, message [32]byte, signature *big.Int, rx *big.Int, ry *big.Int, _hash *big.Int, validators []common.Address) (*types.Transaction, error) {
	return _OracleContract.Contract.SubmitTransactionValidationResult(&_OracleContract.TransactOpts, _result, message, signature, rx, ry, _hash, validators)
}

// SubmitValidationResultBLS is a paid mutator transaction binding the contract method 0x0f14ce5e.
//
// Solidity: function submitValidationResultBLS(uint8 _typ, bool _result, bytes32 message, uint256[4] publicKey, uint256[2] _signature, uint256[2] _hash, address[] validators) returns()
func (_OracleContract *OracleContractTransactor) SubmitValidationResultBLS(opts *bind.TransactOpts, _typ uint8, _result bool, message [32]byte, publicKey [4]*big.Int, _signature [2]*big.Int, _hash [2]*big.Int, validators []common.Address) (*types.Transaction, error) {
	return _OracleContract.contract.Transact(opts, "submitValidationResultBLS", _typ, _result, message, publicKey, _signature, _hash, validators)
}

// SubmitValidationResultBLS is a paid mutator transaction binding the contract method 0x0f14ce5e.
//
// Solidity: function submitValidationResultBLS(uint8 _typ, bool _result, bytes32 message, uint256[4] publicKey, uint256[2] _signature, uint256[2] _hash, address[] validators) returns()
func (_OracleContract *OracleContractSession) SubmitValidationResultBLS(_typ uint8, _result bool, message [32]byte, publicKey [4]*big.Int, _signature [2]*big.Int, _hash [2]*big.Int, validators []common.Address) (*types.Transaction, error) {
	return _OracleContract.Contract.SubmitValidationResultBLS(&_OracleContract.TransactOpts, _typ, _result, message, publicKey, _signature, _hash, validators)
}

// SubmitValidationResultBLS is a paid mutator transaction binding the contract method 0x0f14ce5e.
//
// Solidity: function submitValidationResultBLS(uint8 _typ, bool _result, bytes32 message, uint256[4] publicKey, uint256[2] _signature, uint256[2] _hash, address[] validators) returns()
func (_OracleContract *OracleContractTransactorSession) SubmitValidationResultBLS(_typ uint8, _result bool, message [32]byte, publicKey [4]*big.Int, _signature [2]*big.Int, _hash [2]*big.Int, validators []common.Address) (*types.Transaction, error) {
	return _OracleContract.Contract.SubmitValidationResultBLS(&_OracleContract.TransactOpts, _typ, _result, message, publicKey, _signature, _hash, validators)
}

// Unregister is a paid mutator transaction binding the contract method 0x2ec2c246.
//
// Solidity: function unregister(address unregisterAddr) returns()
func (_OracleContract *OracleContractTransactor) Unregister(opts *bind.TransactOpts, unregisterAddr common.Address) (*types.Transaction, error) {
	return _OracleContract.contract.Transact(opts, "unregister", unregisterAddr)
}

// Unregister is a paid mutator transaction binding the contract method 0x2ec2c246.
//
// Solidity: function unregister(address unregisterAddr) returns()
func (_OracleContract *OracleContractSession) Unregister(unregisterAddr common.Address) (*types.Transaction, error) {
	return _OracleContract.Contract.Unregister(&_OracleContract.TransactOpts, unregisterAddr)
}

// Unregister is a paid mutator transaction binding the contract method 0x2ec2c246.
//
// Solidity: function unregister(address unregisterAddr) returns()
func (_OracleContract *OracleContractTransactorSession) Unregister(unregisterAddr common.Address) (*types.Transaction, error) {
	return _OracleContract.Contract.Unregister(&_OracleContract.TransactOpts, unregisterAddr)
}

// ValidateTransaction is a paid mutator transaction binding the contract method 0x3e170dd9.
//
// Solidity: function validateTransaction(bytes32 _message, uint256 size, uint256 minRank) payable returns()
func (_OracleContract *OracleContractTransactor) ValidateTransaction(opts *bind.TransactOpts, _message [32]byte, size *big.Int, minRank *big.Int) (*types.Transaction, error) {
	return _OracleContract.contract.Transact(opts, "validateTransaction", _message, size, minRank)
}

// ValidateTransaction is a paid mutator transaction binding the contract method 0x3e170dd9.
//
// Solidity: function validateTransaction(bytes32 _message, uint256 size, uint256 minRank) payable returns()
func (_OracleContract *OracleContractSession) ValidateTransaction(_message [32]byte, size *big.Int, minRank *big.Int) (*types.Transaction, error) {
	return _OracleContract.Contract.ValidateTransaction(&_OracleContract.TransactOpts, _message, size, minRank)
}

// ValidateTransaction is a paid mutator transaction binding the contract method 0x3e170dd9.
//
// Solidity: function validateTransaction(bytes32 _message, uint256 size, uint256 minRank) payable returns()
func (_OracleContract *OracleContractTransactorSession) ValidateTransaction(_message [32]byte, size *big.Int, minRank *big.Int) (*types.Transaction, error) {
	return _OracleContract.Contract.ValidateTransaction(&_OracleContract.TransactOpts, _message, size, minRank)
}

// OracleContractRegisterOracleNodeIterator is returned from FilterRegisterOracleNode and is used to iterate over the raw logs and unpacked data for RegisterOracleNode events raised by the OracleContract contract.
type OracleContractRegisterOracleNodeIterator struct {
	Event *OracleContractRegisterOracleNode // Event containing the contract specifics and raw log

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
func (it *OracleContractRegisterOracleNodeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OracleContractRegisterOracleNode)
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
		it.Event = new(OracleContractRegisterOracleNode)
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
func (it *OracleContractRegisterOracleNodeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OracleContractRegisterOracleNodeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OracleContractRegisterOracleNode represents a RegisterOracleNode event raised by the OracleContract contract.
type OracleContractRegisterOracleNode struct {
	Sender common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterRegisterOracleNode is a free log retrieval operation binding the contract event 0x463624ffd45713d944420ab743c635b5714b8dbabe9c3ae0045ba085e71fada0.
//
// Solidity: event RegisterOracleNode(address indexed sender)
func (_OracleContract *OracleContractFilterer) FilterRegisterOracleNode(opts *bind.FilterOpts, sender []common.Address) (*OracleContractRegisterOracleNodeIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _OracleContract.contract.FilterLogs(opts, "RegisterOracleNode", senderRule)
	if err != nil {
		return nil, err
	}
	return &OracleContractRegisterOracleNodeIterator{contract: _OracleContract.contract, event: "RegisterOracleNode", logs: logs, sub: sub}, nil
}

// WatchRegisterOracleNode is a free log subscription operation binding the contract event 0x463624ffd45713d944420ab743c635b5714b8dbabe9c3ae0045ba085e71fada0.
//
// Solidity: event RegisterOracleNode(address indexed sender)
func (_OracleContract *OracleContractFilterer) WatchRegisterOracleNode(opts *bind.WatchOpts, sink chan<- *OracleContractRegisterOracleNode, sender []common.Address) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _OracleContract.contract.WatchLogs(opts, "RegisterOracleNode", senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OracleContractRegisterOracleNode)
				if err := _OracleContract.contract.UnpackLog(event, "RegisterOracleNode", log); err != nil {
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

// ParseRegisterOracleNode is a log parse operation binding the contract event 0x463624ffd45713d944420ab743c635b5714b8dbabe9c3ae0045ba085e71fada0.
//
// Solidity: event RegisterOracleNode(address indexed sender)
func (_OracleContract *OracleContractFilterer) ParseRegisterOracleNode(log types.Log) (*OracleContractRegisterOracleNode, error) {
	event := new(OracleContractRegisterOracleNode)
	if err := _OracleContract.contract.UnpackLog(event, "RegisterOracleNode", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OracleContractValidationRequestIterator is returned from FilterValidationRequest and is used to iterate over the raw logs and unpacked data for ValidationRequest events raised by the OracleContract contract.
type OracleContractValidationRequestIterator struct {
	Event *OracleContractValidationRequest // Event containing the contract specifics and raw log

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
func (it *OracleContractValidationRequestIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OracleContractValidationRequest)
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
		it.Event = new(OracleContractValidationRequest)
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
func (it *OracleContractValidationRequestIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OracleContractValidationRequestIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OracleContractValidationRequest represents a ValidationRequest event raised by the OracleContract contract.
type OracleContractValidationRequest struct {
	Typ     uint8
	From    common.Address
	Hash    [32]byte
	Size    *big.Int
	MinRank *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterValidationRequest is a free log retrieval operation binding the contract event 0xa2a630edc48d1fc0a90b61e08bbc34ce73957ce9d59455207cd0591006bd4d4b.
//
// Solidity: event ValidationRequest(uint8 typ, address indexed from, bytes32 hash, uint256 size, uint256 minRank)
func (_OracleContract *OracleContractFilterer) FilterValidationRequest(opts *bind.FilterOpts, from []common.Address) (*OracleContractValidationRequestIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}

	logs, sub, err := _OracleContract.contract.FilterLogs(opts, "ValidationRequest", fromRule)
	if err != nil {
		return nil, err
	}
	return &OracleContractValidationRequestIterator{contract: _OracleContract.contract, event: "ValidationRequest", logs: logs, sub: sub}, nil
}

// WatchValidationRequest is a free log subscription operation binding the contract event 0xa2a630edc48d1fc0a90b61e08bbc34ce73957ce9d59455207cd0591006bd4d4b.
//
// Solidity: event ValidationRequest(uint8 typ, address indexed from, bytes32 hash, uint256 size, uint256 minRank)
func (_OracleContract *OracleContractFilterer) WatchValidationRequest(opts *bind.WatchOpts, sink chan<- *OracleContractValidationRequest, from []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}

	logs, sub, err := _OracleContract.contract.WatchLogs(opts, "ValidationRequest", fromRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OracleContractValidationRequest)
				if err := _OracleContract.contract.UnpackLog(event, "ValidationRequest", log); err != nil {
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

// ParseValidationRequest is a log parse operation binding the contract event 0xa2a630edc48d1fc0a90b61e08bbc34ce73957ce9d59455207cd0591006bd4d4b.
//
// Solidity: event ValidationRequest(uint8 typ, address indexed from, bytes32 hash, uint256 size, uint256 minRank)
func (_OracleContract *OracleContractFilterer) ParseValidationRequest(log types.Log) (*OracleContractValidationRequest, error) {
	event := new(OracleContractValidationRequest)
	if err := _OracleContract.contract.UnpackLog(event, "ValidationRequest", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
