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

// RegistryNode is an auto generated low-level Go binding around an user-defined struct.
type RegistryNode struct {
	Addr   common.Address
	IpAddr string
	PubKey [2]*big.Int
	Index  *big.Int
	Lambda *big.Int
	Stake  *big.Int
}

// RegistryMetaData contains all meta data concerning the Registry contract.
var RegistryMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RegisterOracleNode\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"countOracleNodes\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_index\",\"type\":\"uint256\"}],\"name\":\"findOracleNodeByIndex\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"ipAddr\",\"type\":\"string\"},{\"internalType\":\"uint256[2]\",\"name\":\"pubKey\",\"type\":\"uint256[2]\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"lambda\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"stake\",\"type\":\"uint256\"}],\"internalType\":\"structRegistry.Node\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getAggregator\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"getIndex\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"getLambda\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"getNodeByAddress\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"ipAddr\",\"type\":\"string\"},{\"internalType\":\"uint256[2]\",\"name\":\"pubKey\",\"type\":\"uint256[2]\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"lambda\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"stake\",\"type\":\"uint256\"}],\"internalType\":\"structRegistry.Node\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"isAggregator\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"minStake\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"ipAddr\",\"type\":\"string\"},{\"internalType\":\"uint256[2]\",\"name\":\"pubKey\",\"type\":\"uint256[2]\"}],\"name\":\"register\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"lambda\",\"type\":\"uint256\"}],\"name\":\"setLambda\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"unregister\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"}]",
}

// RegistryABI is the input ABI used to generate the binding from.
// Deprecated: Use RegistryMetaData.ABI instead.
var RegistryABI = RegistryMetaData.ABI

// Registry is an auto generated Go binding around an Ethereum contract.
type Registry struct {
	RegistryCaller     // Read-only binding to the contract
	RegistryTransactor // Write-only binding to the contract
	RegistryFilterer   // Log filterer for contract events
}

// RegistryCaller is an auto generated read-only Go binding around an Ethereum contract.
type RegistryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RegistryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type RegistryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RegistryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type RegistryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RegistrySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type RegistrySession struct {
	Contract     *Registry         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// RegistryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type RegistryCallerSession struct {
	Contract *RegistryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// RegistryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type RegistryTransactorSession struct {
	Contract     *RegistryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// RegistryRaw is an auto generated low-level Go binding around an Ethereum contract.
type RegistryRaw struct {
	Contract *Registry // Generic contract binding to access the raw methods on
}

// RegistryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type RegistryCallerRaw struct {
	Contract *RegistryCaller // Generic read-only contract binding to access the raw methods on
}

// RegistryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type RegistryTransactorRaw struct {
	Contract *RegistryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewRegistry creates a new instance of Registry, bound to a specific deployed contract.
func NewRegistry(address common.Address, backend bind.ContractBackend) (*Registry, error) {
	contract, err := bindRegistry(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Registry{RegistryCaller: RegistryCaller{contract: contract}, RegistryTransactor: RegistryTransactor{contract: contract}, RegistryFilterer: RegistryFilterer{contract: contract}}, nil
}

// NewRegistryCaller creates a new read-only instance of Registry, bound to a specific deployed contract.
func NewRegistryCaller(address common.Address, caller bind.ContractCaller) (*RegistryCaller, error) {
	contract, err := bindRegistry(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &RegistryCaller{contract: contract}, nil
}

// NewRegistryTransactor creates a new write-only instance of Registry, bound to a specific deployed contract.
func NewRegistryTransactor(address common.Address, transactor bind.ContractTransactor) (*RegistryTransactor, error) {
	contract, err := bindRegistry(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &RegistryTransactor{contract: contract}, nil
}

// NewRegistryFilterer creates a new log filterer instance of Registry, bound to a specific deployed contract.
func NewRegistryFilterer(address common.Address, filterer bind.ContractFilterer) (*RegistryFilterer, error) {
	contract, err := bindRegistry(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &RegistryFilterer{contract: contract}, nil
}

// bindRegistry binds a generic wrapper to an already deployed contract.
func bindRegistry(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := RegistryMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Registry *RegistryRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Registry.Contract.RegistryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Registry *RegistryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Registry.Contract.RegistryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Registry *RegistryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Registry.Contract.RegistryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Registry *RegistryCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Registry.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Registry *RegistryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Registry.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Registry *RegistryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Registry.Contract.contract.Transact(opts, method, params...)
}

// CountOracleNodes is a free data retrieval call binding the contract method 0x836f187a.
//
// Solidity: function countOracleNodes() view returns(uint256)
func (_Registry *RegistryCaller) CountOracleNodes(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Registry.contract.Call(opts, &out, "countOracleNodes")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CountOracleNodes is a free data retrieval call binding the contract method 0x836f187a.
//
// Solidity: function countOracleNodes() view returns(uint256)
func (_Registry *RegistrySession) CountOracleNodes() (*big.Int, error) {
	return _Registry.Contract.CountOracleNodes(&_Registry.CallOpts)
}

// CountOracleNodes is a free data retrieval call binding the contract method 0x836f187a.
//
// Solidity: function countOracleNodes() view returns(uint256)
func (_Registry *RegistryCallerSession) CountOracleNodes() (*big.Int, error) {
	return _Registry.Contract.CountOracleNodes(&_Registry.CallOpts)
}

// FindOracleNodeByIndex is a free data retrieval call binding the contract method 0x272132e9.
//
// Solidity: function findOracleNodeByIndex(uint256 _index) view returns((address,string,uint256[2],uint256,uint256,uint256))
func (_Registry *RegistryCaller) FindOracleNodeByIndex(opts *bind.CallOpts, _index *big.Int) (RegistryNode, error) {
	var out []interface{}
	err := _Registry.contract.Call(opts, &out, "findOracleNodeByIndex", _index)

	if err != nil {
		return *new(RegistryNode), err
	}

	out0 := *abi.ConvertType(out[0], new(RegistryNode)).(*RegistryNode)

	return out0, err

}

// FindOracleNodeByIndex is a free data retrieval call binding the contract method 0x272132e9.
//
// Solidity: function findOracleNodeByIndex(uint256 _index) view returns((address,string,uint256[2],uint256,uint256,uint256))
func (_Registry *RegistrySession) FindOracleNodeByIndex(_index *big.Int) (RegistryNode, error) {
	return _Registry.Contract.FindOracleNodeByIndex(&_Registry.CallOpts, _index)
}

// FindOracleNodeByIndex is a free data retrieval call binding the contract method 0x272132e9.
//
// Solidity: function findOracleNodeByIndex(uint256 _index) view returns((address,string,uint256[2],uint256,uint256,uint256))
func (_Registry *RegistryCallerSession) FindOracleNodeByIndex(_index *big.Int) (RegistryNode, error) {
	return _Registry.Contract.FindOracleNodeByIndex(&_Registry.CallOpts, _index)
}

// GetAggregator is a free data retrieval call binding the contract method 0x3ad59dbc.
//
// Solidity: function getAggregator() view returns(address)
func (_Registry *RegistryCaller) GetAggregator(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Registry.contract.Call(opts, &out, "getAggregator")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetAggregator is a free data retrieval call binding the contract method 0x3ad59dbc.
//
// Solidity: function getAggregator() view returns(address)
func (_Registry *RegistrySession) GetAggregator() (common.Address, error) {
	return _Registry.Contract.GetAggregator(&_Registry.CallOpts)
}

// GetAggregator is a free data retrieval call binding the contract method 0x3ad59dbc.
//
// Solidity: function getAggregator() view returns(address)
func (_Registry *RegistryCallerSession) GetAggregator() (common.Address, error) {
	return _Registry.Contract.GetAggregator(&_Registry.CallOpts)
}

// GetIndex is a free data retrieval call binding the contract method 0xb31610db.
//
// Solidity: function getIndex(address addr) view returns(uint256)
func (_Registry *RegistryCaller) GetIndex(opts *bind.CallOpts, addr common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Registry.contract.Call(opts, &out, "getIndex", addr)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetIndex is a free data retrieval call binding the contract method 0xb31610db.
//
// Solidity: function getIndex(address addr) view returns(uint256)
func (_Registry *RegistrySession) GetIndex(addr common.Address) (*big.Int, error) {
	return _Registry.Contract.GetIndex(&_Registry.CallOpts, addr)
}

// GetIndex is a free data retrieval call binding the contract method 0xb31610db.
//
// Solidity: function getIndex(address addr) view returns(uint256)
func (_Registry *RegistryCallerSession) GetIndex(addr common.Address) (*big.Int, error) {
	return _Registry.Contract.GetIndex(&_Registry.CallOpts, addr)
}

// GetLambda is a free data retrieval call binding the contract method 0x12fa1e35.
//
// Solidity: function getLambda(address addr) view returns(uint256)
func (_Registry *RegistryCaller) GetLambda(opts *bind.CallOpts, addr common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Registry.contract.Call(opts, &out, "getLambda", addr)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetLambda is a free data retrieval call binding the contract method 0x12fa1e35.
//
// Solidity: function getLambda(address addr) view returns(uint256)
func (_Registry *RegistrySession) GetLambda(addr common.Address) (*big.Int, error) {
	return _Registry.Contract.GetLambda(&_Registry.CallOpts, addr)
}

// GetLambda is a free data retrieval call binding the contract method 0x12fa1e35.
//
// Solidity: function getLambda(address addr) view returns(uint256)
func (_Registry *RegistryCallerSession) GetLambda(addr common.Address) (*big.Int, error) {
	return _Registry.Contract.GetLambda(&_Registry.CallOpts, addr)
}

// GetNodeByAddress is a free data retrieval call binding the contract method 0xa9860308.
//
// Solidity: function getNodeByAddress(address addr) view returns((address,string,uint256[2],uint256,uint256,uint256))
func (_Registry *RegistryCaller) GetNodeByAddress(opts *bind.CallOpts, addr common.Address) (RegistryNode, error) {
	var out []interface{}
	err := _Registry.contract.Call(opts, &out, "getNodeByAddress", addr)

	if err != nil {
		return *new(RegistryNode), err
	}

	out0 := *abi.ConvertType(out[0], new(RegistryNode)).(*RegistryNode)

	return out0, err

}

// GetNodeByAddress is a free data retrieval call binding the contract method 0xa9860308.
//
// Solidity: function getNodeByAddress(address addr) view returns((address,string,uint256[2],uint256,uint256,uint256))
func (_Registry *RegistrySession) GetNodeByAddress(addr common.Address) (RegistryNode, error) {
	return _Registry.Contract.GetNodeByAddress(&_Registry.CallOpts, addr)
}

// GetNodeByAddress is a free data retrieval call binding the contract method 0xa9860308.
//
// Solidity: function getNodeByAddress(address addr) view returns((address,string,uint256[2],uint256,uint256,uint256))
func (_Registry *RegistryCallerSession) GetNodeByAddress(addr common.Address) (RegistryNode, error) {
	return _Registry.Contract.GetNodeByAddress(&_Registry.CallOpts, addr)
}

// IsAggregator is a free data retrieval call binding the contract method 0x1e8f3c95.
//
// Solidity: function isAggregator(address addr) view returns(bool)
func (_Registry *RegistryCaller) IsAggregator(opts *bind.CallOpts, addr common.Address) (bool, error) {
	var out []interface{}
	err := _Registry.contract.Call(opts, &out, "isAggregator", addr)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsAggregator is a free data retrieval call binding the contract method 0x1e8f3c95.
//
// Solidity: function isAggregator(address addr) view returns(bool)
func (_Registry *RegistrySession) IsAggregator(addr common.Address) (bool, error) {
	return _Registry.Contract.IsAggregator(&_Registry.CallOpts, addr)
}

// IsAggregator is a free data retrieval call binding the contract method 0x1e8f3c95.
//
// Solidity: function isAggregator(address addr) view returns(bool)
func (_Registry *RegistryCallerSession) IsAggregator(addr common.Address) (bool, error) {
	return _Registry.Contract.IsAggregator(&_Registry.CallOpts, addr)
}

// MinStake is a free data retrieval call binding the contract method 0x375b3c0a.
//
// Solidity: function minStake() pure returns(uint256)
func (_Registry *RegistryCaller) MinStake(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Registry.contract.Call(opts, &out, "minStake")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MinStake is a free data retrieval call binding the contract method 0x375b3c0a.
//
// Solidity: function minStake() pure returns(uint256)
func (_Registry *RegistrySession) MinStake() (*big.Int, error) {
	return _Registry.Contract.MinStake(&_Registry.CallOpts)
}

// MinStake is a free data retrieval call binding the contract method 0x375b3c0a.
//
// Solidity: function minStake() pure returns(uint256)
func (_Registry *RegistryCallerSession) MinStake() (*big.Int, error) {
	return _Registry.Contract.MinStake(&_Registry.CallOpts)
}

// Register is a paid mutator transaction binding the contract method 0xb97e80bb.
//
// Solidity: function register(string ipAddr, uint256[2] pubKey) payable returns()
func (_Registry *RegistryTransactor) Register(opts *bind.TransactOpts, ipAddr string, pubKey [2]*big.Int) (*types.Transaction, error) {
	return _Registry.contract.Transact(opts, "register", ipAddr, pubKey)
}

// Register is a paid mutator transaction binding the contract method 0xb97e80bb.
//
// Solidity: function register(string ipAddr, uint256[2] pubKey) payable returns()
func (_Registry *RegistrySession) Register(ipAddr string, pubKey [2]*big.Int) (*types.Transaction, error) {
	return _Registry.Contract.Register(&_Registry.TransactOpts, ipAddr, pubKey)
}

// Register is a paid mutator transaction binding the contract method 0xb97e80bb.
//
// Solidity: function register(string ipAddr, uint256[2] pubKey) payable returns()
func (_Registry *RegistryTransactorSession) Register(ipAddr string, pubKey [2]*big.Int) (*types.Transaction, error) {
	return _Registry.Contract.Register(&_Registry.TransactOpts, ipAddr, pubKey)
}

// SetLambda is a paid mutator transaction binding the contract method 0x9994deb2.
//
// Solidity: function setLambda(address addr, uint256 lambda) returns()
func (_Registry *RegistryTransactor) SetLambda(opts *bind.TransactOpts, addr common.Address, lambda *big.Int) (*types.Transaction, error) {
	return _Registry.contract.Transact(opts, "setLambda", addr, lambda)
}

// SetLambda is a paid mutator transaction binding the contract method 0x9994deb2.
//
// Solidity: function setLambda(address addr, uint256 lambda) returns()
func (_Registry *RegistrySession) SetLambda(addr common.Address, lambda *big.Int) (*types.Transaction, error) {
	return _Registry.Contract.SetLambda(&_Registry.TransactOpts, addr, lambda)
}

// SetLambda is a paid mutator transaction binding the contract method 0x9994deb2.
//
// Solidity: function setLambda(address addr, uint256 lambda) returns()
func (_Registry *RegistryTransactorSession) SetLambda(addr common.Address, lambda *big.Int) (*types.Transaction, error) {
	return _Registry.Contract.SetLambda(&_Registry.TransactOpts, addr, lambda)
}

// Unregister is a paid mutator transaction binding the contract method 0xe79a198f.
//
// Solidity: function unregister() payable returns()
func (_Registry *RegistryTransactor) Unregister(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Registry.contract.Transact(opts, "unregister")
}

// Unregister is a paid mutator transaction binding the contract method 0xe79a198f.
//
// Solidity: function unregister() payable returns()
func (_Registry *RegistrySession) Unregister() (*types.Transaction, error) {
	return _Registry.Contract.Unregister(&_Registry.TransactOpts)
}

// Unregister is a paid mutator transaction binding the contract method 0xe79a198f.
//
// Solidity: function unregister() payable returns()
func (_Registry *RegistryTransactorSession) Unregister() (*types.Transaction, error) {
	return _Registry.Contract.Unregister(&_Registry.TransactOpts)
}

// RegistryRegisterOracleNodeIterator is returned from FilterRegisterOracleNode and is used to iterate over the raw logs and unpacked data for RegisterOracleNode events raised by the Registry contract.
type RegistryRegisterOracleNodeIterator struct {
	Event *RegistryRegisterOracleNode // Event containing the contract specifics and raw log

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
func (it *RegistryRegisterOracleNodeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RegistryRegisterOracleNode)
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
		it.Event = new(RegistryRegisterOracleNode)
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
func (it *RegistryRegisterOracleNodeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RegistryRegisterOracleNodeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RegistryRegisterOracleNode represents a RegisterOracleNode event raised by the Registry contract.
type RegistryRegisterOracleNode struct {
	Sender common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterRegisterOracleNode is a free log retrieval operation binding the contract event 0x463624ffd45713d944420ab743c635b5714b8dbabe9c3ae0045ba085e71fada0.
//
// Solidity: event RegisterOracleNode(address indexed sender)
func (_Registry *RegistryFilterer) FilterRegisterOracleNode(opts *bind.FilterOpts, sender []common.Address) (*RegistryRegisterOracleNodeIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _Registry.contract.FilterLogs(opts, "RegisterOracleNode", senderRule)
	if err != nil {
		return nil, err
	}
	return &RegistryRegisterOracleNodeIterator{contract: _Registry.contract, event: "RegisterOracleNode", logs: logs, sub: sub}, nil
}

// WatchRegisterOracleNode is a free log subscription operation binding the contract event 0x463624ffd45713d944420ab743c635b5714b8dbabe9c3ae0045ba085e71fada0.
//
// Solidity: event RegisterOracleNode(address indexed sender)
func (_Registry *RegistryFilterer) WatchRegisterOracleNode(opts *bind.WatchOpts, sink chan<- *RegistryRegisterOracleNode, sender []common.Address) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _Registry.contract.WatchLogs(opts, "RegisterOracleNode", senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RegistryRegisterOracleNode)
				if err := _Registry.contract.UnpackLog(event, "RegisterOracleNode", log); err != nil {
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
func (_Registry *RegistryFilterer) ParseRegisterOracleNode(log types.Log) (*RegistryRegisterOracleNode, error) {
	event := new(RegistryRegisterOracleNode)
	if err := _Registry.contract.UnpackLog(event, "RegisterOracleNode", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
