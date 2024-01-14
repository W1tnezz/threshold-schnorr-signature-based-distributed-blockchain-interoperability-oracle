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

// DKGMetaData contains all meta data concerning the DKG contract.
var DKGMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_registryContract\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256[2]\",\"name\":\"pubKey\",\"type\":\"uint256[2]\"}],\"name\":\"DistKey\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"Y\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"enroll\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getPubKey\",\"outputs\":[{\"internalType\":\"uint256[2]\",\"name\":\"\",\"type\":\"uint256[2]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getValidators\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"needEnroll\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"usePubKey\",\"outputs\":[{\"internalType\":\"uint256[2]\",\"name\":\"\",\"type\":\"uint256[2]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// DKGABI is the input ABI used to generate the binding from.
// Deprecated: Use DKGMetaData.ABI instead.
var DKGABI = DKGMetaData.ABI

// DKG is an auto generated Go binding around an Ethereum contract.
type DKG struct {
	DKGCaller     // Read-only binding to the contract
	DKGTransactor // Write-only binding to the contract
	DKGFilterer   // Log filterer for contract events
}

// DKGCaller is an auto generated read-only Go binding around an Ethereum contract.
type DKGCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DKGTransactor is an auto generated write-only Go binding around an Ethereum contract.
type DKGTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DKGFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type DKGFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DKGSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type DKGSession struct {
	Contract     *DKG              // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// DKGCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type DKGCallerSession struct {
	Contract *DKGCaller    // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// DKGTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type DKGTransactorSession struct {
	Contract     *DKGTransactor    // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// DKGRaw is an auto generated low-level Go binding around an Ethereum contract.
type DKGRaw struct {
	Contract *DKG // Generic contract binding to access the raw methods on
}

// DKGCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type DKGCallerRaw struct {
	Contract *DKGCaller // Generic read-only contract binding to access the raw methods on
}

// DKGTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type DKGTransactorRaw struct {
	Contract *DKGTransactor // Generic write-only contract binding to access the raw methods on
}

// NewDKG creates a new instance of DKG, bound to a specific deployed contract.
func NewDKG(address common.Address, backend bind.ContractBackend) (*DKG, error) {
	contract, err := bindDKG(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &DKG{DKGCaller: DKGCaller{contract: contract}, DKGTransactor: DKGTransactor{contract: contract}, DKGFilterer: DKGFilterer{contract: contract}}, nil
}

// NewDKGCaller creates a new read-only instance of DKG, bound to a specific deployed contract.
func NewDKGCaller(address common.Address, caller bind.ContractCaller) (*DKGCaller, error) {
	contract, err := bindDKG(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &DKGCaller{contract: contract}, nil
}

// NewDKGTransactor creates a new write-only instance of DKG, bound to a specific deployed contract.
func NewDKGTransactor(address common.Address, transactor bind.ContractTransactor) (*DKGTransactor, error) {
	contract, err := bindDKG(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &DKGTransactor{contract: contract}, nil
}

// NewDKGFilterer creates a new log filterer instance of DKG, bound to a specific deployed contract.
func NewDKGFilterer(address common.Address, filterer bind.ContractFilterer) (*DKGFilterer, error) {
	contract, err := bindDKG(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &DKGFilterer{contract: contract}, nil
}

// bindDKG binds a generic wrapper to an already deployed contract.
func bindDKG(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := DKGMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_DKG *DKGRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _DKG.Contract.DKGCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_DKG *DKGRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DKG.Contract.DKGTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_DKG *DKGRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _DKG.Contract.DKGTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_DKG *DKGCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _DKG.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_DKG *DKGTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DKG.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_DKG *DKGTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _DKG.Contract.contract.Transact(opts, method, params...)
}

// Y is a free data retrieval call binding the contract method 0xa63f6e5f.
//
// Solidity: function Y(uint256 ) view returns(uint256)
func (_DKG *DKGCaller) Y(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _DKG.contract.Call(opts, &out, "Y", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Y is a free data retrieval call binding the contract method 0xa63f6e5f.
//
// Solidity: function Y(uint256 ) view returns(uint256)
func (_DKG *DKGSession) Y(arg0 *big.Int) (*big.Int, error) {
	return _DKG.Contract.Y(&_DKG.CallOpts, arg0)
}

// Y is a free data retrieval call binding the contract method 0xa63f6e5f.
//
// Solidity: function Y(uint256 ) view returns(uint256)
func (_DKG *DKGCallerSession) Y(arg0 *big.Int) (*big.Int, error) {
	return _DKG.Contract.Y(&_DKG.CallOpts, arg0)
}

// GetPubKey is a free data retrieval call binding the contract method 0x4ad02ef1.
//
// Solidity: function getPubKey() view returns(uint256[2])
func (_DKG *DKGCaller) GetPubKey(opts *bind.CallOpts) ([2]*big.Int, error) {
	var out []interface{}
	err := _DKG.contract.Call(opts, &out, "getPubKey")

	if err != nil {
		return *new([2]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([2]*big.Int)).(*[2]*big.Int)

	return out0, err

}

// GetPubKey is a free data retrieval call binding the contract method 0x4ad02ef1.
//
// Solidity: function getPubKey() view returns(uint256[2])
func (_DKG *DKGSession) GetPubKey() ([2]*big.Int, error) {
	return _DKG.Contract.GetPubKey(&_DKG.CallOpts)
}

// GetPubKey is a free data retrieval call binding the contract method 0x4ad02ef1.
//
// Solidity: function getPubKey() view returns(uint256[2])
func (_DKG *DKGCallerSession) GetPubKey() ([2]*big.Int, error) {
	return _DKG.Contract.GetPubKey(&_DKG.CallOpts)
}

// GetValidators is a free data retrieval call binding the contract method 0xb7ab4db5.
//
// Solidity: function getValidators() view returns(address[])
func (_DKG *DKGCaller) GetValidators(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _DKG.contract.Call(opts, &out, "getValidators")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetValidators is a free data retrieval call binding the contract method 0xb7ab4db5.
//
// Solidity: function getValidators() view returns(address[])
func (_DKG *DKGSession) GetValidators() ([]common.Address, error) {
	return _DKG.Contract.GetValidators(&_DKG.CallOpts)
}

// GetValidators is a free data retrieval call binding the contract method 0xb7ab4db5.
//
// Solidity: function getValidators() view returns(address[])
func (_DKG *DKGCallerSession) GetValidators() ([]common.Address, error) {
	return _DKG.Contract.GetValidators(&_DKG.CallOpts)
}

// NeedEnroll is a free data retrieval call binding the contract method 0x013f7be5.
//
// Solidity: function needEnroll() view returns(bool)
func (_DKG *DKGCaller) NeedEnroll(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _DKG.contract.Call(opts, &out, "needEnroll")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// NeedEnroll is a free data retrieval call binding the contract method 0x013f7be5.
//
// Solidity: function needEnroll() view returns(bool)
func (_DKG *DKGSession) NeedEnroll() (bool, error) {
	return _DKG.Contract.NeedEnroll(&_DKG.CallOpts)
}

// NeedEnroll is a free data retrieval call binding the contract method 0x013f7be5.
//
// Solidity: function needEnroll() view returns(bool)
func (_DKG *DKGCallerSession) NeedEnroll() (bool, error) {
	return _DKG.Contract.NeedEnroll(&_DKG.CallOpts)
}

// Enroll is a paid mutator transaction binding the contract method 0xe65f2a7e.
//
// Solidity: function enroll() payable returns(bool)
func (_DKG *DKGTransactor) Enroll(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DKG.contract.Transact(opts, "enroll")
}

// Enroll is a paid mutator transaction binding the contract method 0xe65f2a7e.
//
// Solidity: function enroll() payable returns(bool)
func (_DKG *DKGSession) Enroll() (*types.Transaction, error) {
	return _DKG.Contract.Enroll(&_DKG.TransactOpts)
}

// Enroll is a paid mutator transaction binding the contract method 0xe65f2a7e.
//
// Solidity: function enroll() payable returns(bool)
func (_DKG *DKGTransactorSession) Enroll() (*types.Transaction, error) {
	return _DKG.Contract.Enroll(&_DKG.TransactOpts)
}

// UsePubKey is a paid mutator transaction binding the contract method 0xba1276c7.
//
// Solidity: function usePubKey() returns(uint256[2])
func (_DKG *DKGTransactor) UsePubKey(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DKG.contract.Transact(opts, "usePubKey")
}

// UsePubKey is a paid mutator transaction binding the contract method 0xba1276c7.
//
// Solidity: function usePubKey() returns(uint256[2])
func (_DKG *DKGSession) UsePubKey() (*types.Transaction, error) {
	return _DKG.Contract.UsePubKey(&_DKG.TransactOpts)
}

// UsePubKey is a paid mutator transaction binding the contract method 0xba1276c7.
//
// Solidity: function usePubKey() returns(uint256[2])
func (_DKG *DKGTransactorSession) UsePubKey() (*types.Transaction, error) {
	return _DKG.Contract.UsePubKey(&_DKG.TransactOpts)
}

// DKGDistKeyIterator is returned from FilterDistKey and is used to iterate over the raw logs and unpacked data for DistKey events raised by the DKG contract.
type DKGDistKeyIterator struct {
	Event *DKGDistKey // Event containing the contract specifics and raw log

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
func (it *DKGDistKeyIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DKGDistKey)
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
		it.Event = new(DKGDistKey)
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
func (it *DKGDistKeyIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DKGDistKeyIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DKGDistKey represents a DistKey event raised by the DKG contract.
type DKGDistKey struct {
	PubKey [2]*big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterDistKey is a free log retrieval operation binding the contract event 0x3039495696db8c6135ec1af78a3dd269dc6a9df4b25b07a806701140e8dec0d5.
//
// Solidity: event DistKey(uint256[2] pubKey)
func (_DKG *DKGFilterer) FilterDistKey(opts *bind.FilterOpts) (*DKGDistKeyIterator, error) {

	logs, sub, err := _DKG.contract.FilterLogs(opts, "DistKey")
	if err != nil {
		return nil, err
	}
	return &DKGDistKeyIterator{contract: _DKG.contract, event: "DistKey", logs: logs, sub: sub}, nil
}

// WatchDistKey is a free log subscription operation binding the contract event 0x3039495696db8c6135ec1af78a3dd269dc6a9df4b25b07a806701140e8dec0d5.
//
// Solidity: event DistKey(uint256[2] pubKey)
func (_DKG *DKGFilterer) WatchDistKey(opts *bind.WatchOpts, sink chan<- *DKGDistKey) (event.Subscription, error) {

	logs, sub, err := _DKG.contract.WatchLogs(opts, "DistKey")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DKGDistKey)
				if err := _DKG.contract.UnpackLog(event, "DistKey", log); err != nil {
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

// ParseDistKey is a log parse operation binding the contract event 0x3039495696db8c6135ec1af78a3dd269dc6a9df4b25b07a806701140e8dec0d5.
//
// Solidity: event DistKey(uint256[2] pubKey)
func (_DKG *DKGFilterer) ParseDistKey(log types.Log) (*DKGDistKey, error) {
	event := new(DKGDistKey)
	if err := _DKG.contract.UnpackLog(event, "DistKey", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
