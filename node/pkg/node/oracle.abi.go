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

// OracleMetaData contains all meta data concerning the Oracle contract.
var OracleMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_registryContract\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_dkgContract\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"hash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"needEnroll\",\"type\":\"bool\"}],\"name\":\"ValidationRequest\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"AGGREGATE_FEE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"BASE_FEE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"res\",\"type\":\"bool\"},{\"internalType\":\"bytes32\",\"name\":\"message\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"signature\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"rx\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"ry\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_hash\",\"type\":\"uint256\"}],\"name\":\"submit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_message\",\"type\":\"bytes32\"}],\"name\":\"validateTransaction\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"}]",
}

// OracleABI is the input ABI used to generate the binding from.
// Deprecated: Use OracleMetaData.ABI instead.
var OracleABI = OracleMetaData.ABI

// Oracle is an auto generated Go binding around an Ethereum contract.
type Oracle struct {
	OracleCaller     // Read-only binding to the contract
	OracleTransactor // Write-only binding to the contract
	OracleFilterer   // Log filterer for contract events
}

// OracleCaller is an auto generated read-only Go binding around an Ethereum contract.
type OracleCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OracleTransactor is an auto generated write-only Go binding around an Ethereum contract.
type OracleTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OracleFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type OracleFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OracleSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type OracleSession struct {
	Contract     *Oracle           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// OracleCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type OracleCallerSession struct {
	Contract *OracleCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// OracleTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type OracleTransactorSession struct {
	Contract     *OracleTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// OracleRaw is an auto generated low-level Go binding around an Ethereum contract.
type OracleRaw struct {
	Contract *Oracle // Generic contract binding to access the raw methods on
}

// OracleCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type OracleCallerRaw struct {
	Contract *OracleCaller // Generic read-only contract binding to access the raw methods on
}

// OracleTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type OracleTransactorRaw struct {
	Contract *OracleTransactor // Generic write-only contract binding to access the raw methods on
}

// NewOracle creates a new instance of Oracle, bound to a specific deployed contract.
func NewOracle(address common.Address, backend bind.ContractBackend) (*Oracle, error) {
	contract, err := bindOracle(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Oracle{OracleCaller: OracleCaller{contract: contract}, OracleTransactor: OracleTransactor{contract: contract}, OracleFilterer: OracleFilterer{contract: contract}}, nil
}

// NewOracleCaller creates a new read-only instance of Oracle, bound to a specific deployed contract.
func NewOracleCaller(address common.Address, caller bind.ContractCaller) (*OracleCaller, error) {
	contract, err := bindOracle(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &OracleCaller{contract: contract}, nil
}

// NewOracleTransactor creates a new write-only instance of Oracle, bound to a specific deployed contract.
func NewOracleTransactor(address common.Address, transactor bind.ContractTransactor) (*OracleTransactor, error) {
	contract, err := bindOracle(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &OracleTransactor{contract: contract}, nil
}

// NewOracleFilterer creates a new log filterer instance of Oracle, bound to a specific deployed contract.
func NewOracleFilterer(address common.Address, filterer bind.ContractFilterer) (*OracleFilterer, error) {
	contract, err := bindOracle(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &OracleFilterer{contract: contract}, nil
}

// bindOracle binds a generic wrapper to an already deployed contract.
func bindOracle(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := OracleMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Oracle *OracleRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Oracle.Contract.OracleCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Oracle *OracleRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Oracle.Contract.OracleTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Oracle *OracleRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Oracle.Contract.OracleTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Oracle *OracleCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Oracle.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Oracle *OracleTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Oracle.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Oracle *OracleTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Oracle.Contract.contract.Transact(opts, method, params...)
}

// AGGREGATEFEE is a free data retrieval call binding the contract method 0x1471866c.
//
// Solidity: function AGGREGATE_FEE() view returns(uint256)
func (_Oracle *OracleCaller) AGGREGATEFEE(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Oracle.contract.Call(opts, &out, "AGGREGATE_FEE")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// AGGREGATEFEE is a free data retrieval call binding the contract method 0x1471866c.
//
// Solidity: function AGGREGATE_FEE() view returns(uint256)
func (_Oracle *OracleSession) AGGREGATEFEE() (*big.Int, error) {
	return _Oracle.Contract.AGGREGATEFEE(&_Oracle.CallOpts)
}

// AGGREGATEFEE is a free data retrieval call binding the contract method 0x1471866c.
//
// Solidity: function AGGREGATE_FEE() view returns(uint256)
func (_Oracle *OracleCallerSession) AGGREGATEFEE() (*big.Int, error) {
	return _Oracle.Contract.AGGREGATEFEE(&_Oracle.CallOpts)
}

// BASEFEE is a free data retrieval call binding the contract method 0x3d18651e.
//
// Solidity: function BASE_FEE() view returns(uint256)
func (_Oracle *OracleCaller) BASEFEE(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Oracle.contract.Call(opts, &out, "BASE_FEE")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BASEFEE is a free data retrieval call binding the contract method 0x3d18651e.
//
// Solidity: function BASE_FEE() view returns(uint256)
func (_Oracle *OracleSession) BASEFEE() (*big.Int, error) {
	return _Oracle.Contract.BASEFEE(&_Oracle.CallOpts)
}

// BASEFEE is a free data retrieval call binding the contract method 0x3d18651e.
//
// Solidity: function BASE_FEE() view returns(uint256)
func (_Oracle *OracleCallerSession) BASEFEE() (*big.Int, error) {
	return _Oracle.Contract.BASEFEE(&_Oracle.CallOpts)
}

// TotalFee is a free data retrieval call binding the contract method 0x1df4ccfc.
//
// Solidity: function totalFee() view returns(uint256)
func (_Oracle *OracleCaller) TotalFee(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Oracle.contract.Call(opts, &out, "totalFee")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalFee is a free data retrieval call binding the contract method 0x1df4ccfc.
//
// Solidity: function totalFee() view returns(uint256)
func (_Oracle *OracleSession) TotalFee() (*big.Int, error) {
	return _Oracle.Contract.TotalFee(&_Oracle.CallOpts)
}

// TotalFee is a free data retrieval call binding the contract method 0x1df4ccfc.
//
// Solidity: function totalFee() view returns(uint256)
func (_Oracle *OracleCallerSession) TotalFee() (*big.Int, error) {
	return _Oracle.Contract.TotalFee(&_Oracle.CallOpts)
}

// Submit is a paid mutator transaction binding the contract method 0x6fe72c1d.
//
// Solidity: function submit(bool res, bytes32 message, uint256 signature, uint256 rx, uint256 ry, uint256 _hash) returns()
func (_Oracle *OracleTransactor) Submit(opts *bind.TransactOpts, res bool, message [32]byte, signature *big.Int, rx *big.Int, ry *big.Int, _hash *big.Int) (*types.Transaction, error) {
	return _Oracle.contract.Transact(opts, "submit", res, message, signature, rx, ry, _hash)
}

// Submit is a paid mutator transaction binding the contract method 0x6fe72c1d.
//
// Solidity: function submit(bool res, bytes32 message, uint256 signature, uint256 rx, uint256 ry, uint256 _hash) returns()
func (_Oracle *OracleSession) Submit(res bool, message [32]byte, signature *big.Int, rx *big.Int, ry *big.Int, _hash *big.Int) (*types.Transaction, error) {
	return _Oracle.Contract.Submit(&_Oracle.TransactOpts, res, message, signature, rx, ry, _hash)
}

// Submit is a paid mutator transaction binding the contract method 0x6fe72c1d.
//
// Solidity: function submit(bool res, bytes32 message, uint256 signature, uint256 rx, uint256 ry, uint256 _hash) returns()
func (_Oracle *OracleTransactorSession) Submit(res bool, message [32]byte, signature *big.Int, rx *big.Int, ry *big.Int, _hash *big.Int) (*types.Transaction, error) {
	return _Oracle.Contract.Submit(&_Oracle.TransactOpts, res, message, signature, rx, ry, _hash)
}

// ValidateTransaction is a paid mutator transaction binding the contract method 0x344829c8.
//
// Solidity: function validateTransaction(bytes32 _message) payable returns()
func (_Oracle *OracleTransactor) ValidateTransaction(opts *bind.TransactOpts, _message [32]byte) (*types.Transaction, error) {
	return _Oracle.contract.Transact(opts, "validateTransaction", _message)
}

// ValidateTransaction is a paid mutator transaction binding the contract method 0x344829c8.
//
// Solidity: function validateTransaction(bytes32 _message) payable returns()
func (_Oracle *OracleSession) ValidateTransaction(_message [32]byte) (*types.Transaction, error) {
	return _Oracle.Contract.ValidateTransaction(&_Oracle.TransactOpts, _message)
}

// ValidateTransaction is a paid mutator transaction binding the contract method 0x344829c8.
//
// Solidity: function validateTransaction(bytes32 _message) payable returns()
func (_Oracle *OracleTransactorSession) ValidateTransaction(_message [32]byte) (*types.Transaction, error) {
	return _Oracle.Contract.ValidateTransaction(&_Oracle.TransactOpts, _message)
}

// OracleValidationRequestIterator is returned from FilterValidationRequest and is used to iterate over the raw logs and unpacked data for ValidationRequest events raised by the Oracle contract.
type OracleValidationRequestIterator struct {
	Event *OracleValidationRequest // Event containing the contract specifics and raw log

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
func (it *OracleValidationRequestIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OracleValidationRequest)
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
		it.Event = new(OracleValidationRequest)
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
func (it *OracleValidationRequestIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OracleValidationRequestIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OracleValidationRequest represents a ValidationRequest event raised by the Oracle contract.
type OracleValidationRequest struct {
	From       common.Address
	Hash       [32]byte
	NeedEnroll bool
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterValidationRequest is a free log retrieval operation binding the contract event 0x0ab6f7cba28f303ff389d1006ea6fa3ffb8b673da1052b131df34e5477a45b22.
//
// Solidity: event ValidationRequest(address indexed from, bytes32 hash, bool needEnroll)
func (_Oracle *OracleFilterer) FilterValidationRequest(opts *bind.FilterOpts, from []common.Address) (*OracleValidationRequestIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}

	logs, sub, err := _Oracle.contract.FilterLogs(opts, "ValidationRequest", fromRule)
	if err != nil {
		return nil, err
	}
	return &OracleValidationRequestIterator{contract: _Oracle.contract, event: "ValidationRequest", logs: logs, sub: sub}, nil
}

// WatchValidationRequest is a free log subscription operation binding the contract event 0x0ab6f7cba28f303ff389d1006ea6fa3ffb8b673da1052b131df34e5477a45b22.
//
// Solidity: event ValidationRequest(address indexed from, bytes32 hash, bool needEnroll)
func (_Oracle *OracleFilterer) WatchValidationRequest(opts *bind.WatchOpts, sink chan<- *OracleValidationRequest, from []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}

	logs, sub, err := _Oracle.contract.WatchLogs(opts, "ValidationRequest", fromRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OracleValidationRequest)
				if err := _Oracle.contract.UnpackLog(event, "ValidationRequest", log); err != nil {
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

// ParseValidationRequest is a log parse operation binding the contract event 0x0ab6f7cba28f303ff389d1006ea6fa3ffb8b673da1052b131df34e5477a45b22.
//
// Solidity: event ValidationRequest(address indexed from, bytes32 hash, bool needEnroll)
func (_Oracle *OracleFilterer) ParseValidationRequest(log types.Log) (*OracleValidationRequest, error) {
	event := new(OracleValidationRequest)
	if err := _Oracle.contract.UnpackLog(event, "ValidationRequest", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
