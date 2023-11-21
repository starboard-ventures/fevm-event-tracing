// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package wfil

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

// WfilMetaData contains all meta data concerning the Wfil contract.
var WfilMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"CallFailed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InsufficientFunds\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidAddress\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TimelockActive\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"Unauthorized\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Deposit\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"currentOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"pendingOwner\",\"type\":\"address\"}],\"name\":\"OwnershipPending\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"oldOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Withdrawal\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"acceptOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_a\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"deposit\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_depositor\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"recoverDeposit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"recoveryTimelock\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"revokeOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
}

// WfilABI is the input ABI used to generate the binding from.
// Deprecated: Use WfilMetaData.ABI instead.
var WfilABI = WfilMetaData.ABI

// Wfil is an auto generated Go binding around an Ethereum contract.
type Wfil struct {
	WfilCaller     // Read-only binding to the contract
	WfilTransactor // Write-only binding to the contract
	WfilFilterer   // Log filterer for contract events
}

// WfilCaller is an auto generated read-only Go binding around an Ethereum contract.
type WfilCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// WfilTransactor is an auto generated write-only Go binding around an Ethereum contract.
type WfilTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// WfilFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type WfilFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// WfilSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type WfilSession struct {
	Contract     *Wfil             // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// WfilCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type WfilCallerSession struct {
	Contract *WfilCaller   // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// WfilTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type WfilTransactorSession struct {
	Contract     *WfilTransactor   // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// WfilRaw is an auto generated low-level Go binding around an Ethereum contract.
type WfilRaw struct {
	Contract *Wfil // Generic contract binding to access the raw methods on
}

// WfilCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type WfilCallerRaw struct {
	Contract *WfilCaller // Generic read-only contract binding to access the raw methods on
}

// WfilTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type WfilTransactorRaw struct {
	Contract *WfilTransactor // Generic write-only contract binding to access the raw methods on
}

// NewWfil creates a new instance of Wfil, bound to a specific deployed contract.
func NewWfil(address common.Address, backend bind.ContractBackend) (*Wfil, error) {
	contract, err := bindWfil(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Wfil{WfilCaller: WfilCaller{contract: contract}, WfilTransactor: WfilTransactor{contract: contract}, WfilFilterer: WfilFilterer{contract: contract}}, nil
}

// NewWfilCaller creates a new read-only instance of Wfil, bound to a specific deployed contract.
func NewWfilCaller(address common.Address, caller bind.ContractCaller) (*WfilCaller, error) {
	contract, err := bindWfil(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &WfilCaller{contract: contract}, nil
}

// NewWfilTransactor creates a new write-only instance of Wfil, bound to a specific deployed contract.
func NewWfilTransactor(address common.Address, transactor bind.ContractTransactor) (*WfilTransactor, error) {
	contract, err := bindWfil(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &WfilTransactor{contract: contract}, nil
}

// NewWfilFilterer creates a new log filterer instance of Wfil, bound to a specific deployed contract.
func NewWfilFilterer(address common.Address, filterer bind.ContractFilterer) (*WfilFilterer, error) {
	contract, err := bindWfil(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &WfilFilterer{contract: contract}, nil
}

// bindWfil binds a generic wrapper to an already deployed contract.
func bindWfil(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := WfilMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Wfil *WfilRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Wfil.Contract.WfilCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Wfil *WfilRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Wfil.Contract.WfilTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Wfil *WfilRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Wfil.Contract.WfilTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Wfil *WfilCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Wfil.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Wfil *WfilTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Wfil.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Wfil *WfilTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Wfil.Contract.contract.Transact(opts, method, params...)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address _owner, address _spender) view returns(uint256)
func (_Wfil *WfilCaller) Allowance(opts *bind.CallOpts, _owner common.Address, _spender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Wfil.contract.Call(opts, &out, "allowance", _owner, _spender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address _owner, address _spender) view returns(uint256)
func (_Wfil *WfilSession) Allowance(_owner common.Address, _spender common.Address) (*big.Int, error) {
	return _Wfil.Contract.Allowance(&_Wfil.CallOpts, _owner, _spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address _owner, address _spender) view returns(uint256)
func (_Wfil *WfilCallerSession) Allowance(_owner common.Address, _spender common.Address) (*big.Int, error) {
	return _Wfil.Contract.Allowance(&_Wfil.CallOpts, _owner, _spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address _a) view returns(uint256)
func (_Wfil *WfilCaller) BalanceOf(opts *bind.CallOpts, _a common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Wfil.contract.Call(opts, &out, "balanceOf", _a)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address _a) view returns(uint256)
func (_Wfil *WfilSession) BalanceOf(_a common.Address) (*big.Int, error) {
	return _Wfil.Contract.BalanceOf(&_Wfil.CallOpts, _a)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address _a) view returns(uint256)
func (_Wfil *WfilCallerSession) BalanceOf(_a common.Address) (*big.Int, error) {
	return _Wfil.Contract.BalanceOf(&_Wfil.CallOpts, _a)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_Wfil *WfilCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _Wfil.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_Wfil *WfilSession) Decimals() (uint8, error) {
	return _Wfil.Contract.Decimals(&_Wfil.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_Wfil *WfilCallerSession) Decimals() (uint8, error) {
	return _Wfil.Contract.Decimals(&_Wfil.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Wfil *WfilCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Wfil.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Wfil *WfilSession) Name() (string, error) {
	return _Wfil.Contract.Name(&_Wfil.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Wfil *WfilCallerSession) Name() (string, error) {
	return _Wfil.Contract.Name(&_Wfil.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Wfil *WfilCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Wfil.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Wfil *WfilSession) Owner() (common.Address, error) {
	return _Wfil.Contract.Owner(&_Wfil.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Wfil *WfilCallerSession) Owner() (common.Address, error) {
	return _Wfil.Contract.Owner(&_Wfil.CallOpts)
}

// RecoveryTimelock is a free data retrieval call binding the contract method 0x670a43c4.
//
// Solidity: function recoveryTimelock() view returns(uint256)
func (_Wfil *WfilCaller) RecoveryTimelock(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Wfil.contract.Call(opts, &out, "recoveryTimelock")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// RecoveryTimelock is a free data retrieval call binding the contract method 0x670a43c4.
//
// Solidity: function recoveryTimelock() view returns(uint256)
func (_Wfil *WfilSession) RecoveryTimelock() (*big.Int, error) {
	return _Wfil.Contract.RecoveryTimelock(&_Wfil.CallOpts)
}

// RecoveryTimelock is a free data retrieval call binding the contract method 0x670a43c4.
//
// Solidity: function recoveryTimelock() view returns(uint256)
func (_Wfil *WfilCallerSession) RecoveryTimelock() (*big.Int, error) {
	return _Wfil.Contract.RecoveryTimelock(&_Wfil.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Wfil *WfilCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Wfil.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Wfil *WfilSession) Symbol() (string, error) {
	return _Wfil.Contract.Symbol(&_Wfil.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Wfil *WfilCallerSession) Symbol() (string, error) {
	return _Wfil.Contract.Symbol(&_Wfil.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Wfil *WfilCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Wfil.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Wfil *WfilSession) TotalSupply() (*big.Int, error) {
	return _Wfil.Contract.TotalSupply(&_Wfil.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Wfil *WfilCallerSession) TotalSupply() (*big.Int, error) {
	return _Wfil.Contract.TotalSupply(&_Wfil.CallOpts)
}

// AcceptOwnership is a paid mutator transaction binding the contract method 0x79ba5097.
//
// Solidity: function acceptOwnership() returns()
func (_Wfil *WfilTransactor) AcceptOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Wfil.contract.Transact(opts, "acceptOwnership")
}

// AcceptOwnership is a paid mutator transaction binding the contract method 0x79ba5097.
//
// Solidity: function acceptOwnership() returns()
func (_Wfil *WfilSession) AcceptOwnership() (*types.Transaction, error) {
	return _Wfil.Contract.AcceptOwnership(&_Wfil.TransactOpts)
}

// AcceptOwnership is a paid mutator transaction binding the contract method 0x79ba5097.
//
// Solidity: function acceptOwnership() returns()
func (_Wfil *WfilTransactorSession) AcceptOwnership() (*types.Transaction, error) {
	return _Wfil.Contract.AcceptOwnership(&_Wfil.TransactOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address _spender, uint256 _amount) returns(bool)
func (_Wfil *WfilTransactor) Approve(opts *bind.TransactOpts, _spender common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Wfil.contract.Transact(opts, "approve", _spender, _amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address _spender, uint256 _amount) returns(bool)
func (_Wfil *WfilSession) Approve(_spender common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Wfil.Contract.Approve(&_Wfil.TransactOpts, _spender, _amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address _spender, uint256 _amount) returns(bool)
func (_Wfil *WfilTransactorSession) Approve(_spender common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Wfil.Contract.Approve(&_Wfil.TransactOpts, _spender, _amount)
}

// Deposit is a paid mutator transaction binding the contract method 0xd0e30db0.
//
// Solidity: function deposit() payable returns()
func (_Wfil *WfilTransactor) Deposit(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Wfil.contract.Transact(opts, "deposit")
}

// Deposit is a paid mutator transaction binding the contract method 0xd0e30db0.
//
// Solidity: function deposit() payable returns()
func (_Wfil *WfilSession) Deposit() (*types.Transaction, error) {
	return _Wfil.Contract.Deposit(&_Wfil.TransactOpts)
}

// Deposit is a paid mutator transaction binding the contract method 0xd0e30db0.
//
// Solidity: function deposit() payable returns()
func (_Wfil *WfilTransactorSession) Deposit() (*types.Transaction, error) {
	return _Wfil.Contract.Deposit(&_Wfil.TransactOpts)
}

// RecoverDeposit is a paid mutator transaction binding the contract method 0x436f3b60.
//
// Solidity: function recoverDeposit(address _depositor, uint256 _amount) returns()
func (_Wfil *WfilTransactor) RecoverDeposit(opts *bind.TransactOpts, _depositor common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Wfil.contract.Transact(opts, "recoverDeposit", _depositor, _amount)
}

// RecoverDeposit is a paid mutator transaction binding the contract method 0x436f3b60.
//
// Solidity: function recoverDeposit(address _depositor, uint256 _amount) returns()
func (_Wfil *WfilSession) RecoverDeposit(_depositor common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Wfil.Contract.RecoverDeposit(&_Wfil.TransactOpts, _depositor, _amount)
}

// RecoverDeposit is a paid mutator transaction binding the contract method 0x436f3b60.
//
// Solidity: function recoverDeposit(address _depositor, uint256 _amount) returns()
func (_Wfil *WfilTransactorSession) RecoverDeposit(_depositor common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Wfil.Contract.RecoverDeposit(&_Wfil.TransactOpts, _depositor, _amount)
}

// RevokeOwnership is a paid mutator transaction binding the contract method 0x2b968958.
//
// Solidity: function revokeOwnership() returns()
func (_Wfil *WfilTransactor) RevokeOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Wfil.contract.Transact(opts, "revokeOwnership")
}

// RevokeOwnership is a paid mutator transaction binding the contract method 0x2b968958.
//
// Solidity: function revokeOwnership() returns()
func (_Wfil *WfilSession) RevokeOwnership() (*types.Transaction, error) {
	return _Wfil.Contract.RevokeOwnership(&_Wfil.TransactOpts)
}

// RevokeOwnership is a paid mutator transaction binding the contract method 0x2b968958.
//
// Solidity: function revokeOwnership() returns()
func (_Wfil *WfilTransactorSession) RevokeOwnership() (*types.Transaction, error) {
	return _Wfil.Contract.RevokeOwnership(&_Wfil.TransactOpts)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address _to, uint256 _amount) returns(bool)
func (_Wfil *WfilTransactor) Transfer(opts *bind.TransactOpts, _to common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Wfil.contract.Transact(opts, "transfer", _to, _amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address _to, uint256 _amount) returns(bool)
func (_Wfil *WfilSession) Transfer(_to common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Wfil.Contract.Transfer(&_Wfil.TransactOpts, _to, _amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address _to, uint256 _amount) returns(bool)
func (_Wfil *WfilTransactorSession) Transfer(_to common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Wfil.Contract.Transfer(&_Wfil.TransactOpts, _to, _amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address _owner, address _to, uint256 _amount) returns(bool)
func (_Wfil *WfilTransactor) TransferFrom(opts *bind.TransactOpts, _owner common.Address, _to common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Wfil.contract.Transact(opts, "transferFrom", _owner, _to, _amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address _owner, address _to, uint256 _amount) returns(bool)
func (_Wfil *WfilSession) TransferFrom(_owner common.Address, _to common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Wfil.Contract.TransferFrom(&_Wfil.TransactOpts, _owner, _to, _amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address _owner, address _to, uint256 _amount) returns(bool)
func (_Wfil *WfilTransactorSession) TransferFrom(_owner common.Address, _to common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Wfil.Contract.TransferFrom(&_Wfil.TransactOpts, _owner, _to, _amount)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address _newOwner) returns()
func (_Wfil *WfilTransactor) TransferOwnership(opts *bind.TransactOpts, _newOwner common.Address) (*types.Transaction, error) {
	return _Wfil.contract.Transact(opts, "transferOwnership", _newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address _newOwner) returns()
func (_Wfil *WfilSession) TransferOwnership(_newOwner common.Address) (*types.Transaction, error) {
	return _Wfil.Contract.TransferOwnership(&_Wfil.TransactOpts, _newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address _newOwner) returns()
func (_Wfil *WfilTransactorSession) TransferOwnership(_newOwner common.Address) (*types.Transaction, error) {
	return _Wfil.Contract.TransferOwnership(&_Wfil.TransactOpts, _newOwner)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 _amount) returns()
func (_Wfil *WfilTransactor) Withdraw(opts *bind.TransactOpts, _amount *big.Int) (*types.Transaction, error) {
	return _Wfil.contract.Transact(opts, "withdraw", _amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 _amount) returns()
func (_Wfil *WfilSession) Withdraw(_amount *big.Int) (*types.Transaction, error) {
	return _Wfil.Contract.Withdraw(&_Wfil.TransactOpts, _amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 _amount) returns()
func (_Wfil *WfilTransactorSession) Withdraw(_amount *big.Int) (*types.Transaction, error) {
	return _Wfil.Contract.Withdraw(&_Wfil.TransactOpts, _amount)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Wfil *WfilTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Wfil.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Wfil *WfilSession) Receive() (*types.Transaction, error) {
	return _Wfil.Contract.Receive(&_Wfil.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Wfil *WfilTransactorSession) Receive() (*types.Transaction, error) {
	return _Wfil.Contract.Receive(&_Wfil.TransactOpts)
}

// WfilApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the Wfil contract.
type WfilApprovalIterator struct {
	Event *WfilApproval // Event containing the contract specifics and raw log

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
func (it *WfilApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WfilApproval)
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
		it.Event = new(WfilApproval)
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
func (it *WfilApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WfilApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WfilApproval represents a Approval event raised by the Wfil contract.
type WfilApproval struct {
	Owner   common.Address
	Spender common.Address
	Amount  *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 amount)
func (_Wfil *WfilFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*WfilApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _Wfil.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &WfilApprovalIterator{contract: _Wfil.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 amount)
func (_Wfil *WfilFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *WfilApproval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _Wfil.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WfilApproval)
				if err := _Wfil.contract.UnpackLog(event, "Approval", log); err != nil {
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

// ParseApproval is a log parse operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 amount)
func (_Wfil *WfilFilterer) ParseApproval(log types.Log) (*WfilApproval, error) {
	event := new(WfilApproval)
	if err := _Wfil.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// WfilDepositIterator is returned from FilterDeposit and is used to iterate over the raw logs and unpacked data for Deposit events raised by the Wfil contract.
type WfilDepositIterator struct {
	Event *WfilDeposit // Event containing the contract specifics and raw log

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
func (it *WfilDepositIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WfilDeposit)
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
		it.Event = new(WfilDeposit)
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
func (it *WfilDepositIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WfilDepositIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WfilDeposit represents a Deposit event raised by the Wfil contract.
type WfilDeposit struct {
	From   common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterDeposit is a free log retrieval operation binding the contract event 0xe1fffcc4923d04b559f4d29a8bfc6cda04eb5b0d3c460751c2402c5c5cc9109c.
//
// Solidity: event Deposit(address indexed from, uint256 amount)
func (_Wfil *WfilFilterer) FilterDeposit(opts *bind.FilterOpts, from []common.Address) (*WfilDepositIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}

	logs, sub, err := _Wfil.contract.FilterLogs(opts, "Deposit", fromRule)
	if err != nil {
		return nil, err
	}
	return &WfilDepositIterator{contract: _Wfil.contract, event: "Deposit", logs: logs, sub: sub}, nil
}

// WatchDeposit is a free log subscription operation binding the contract event 0xe1fffcc4923d04b559f4d29a8bfc6cda04eb5b0d3c460751c2402c5c5cc9109c.
//
// Solidity: event Deposit(address indexed from, uint256 amount)
func (_Wfil *WfilFilterer) WatchDeposit(opts *bind.WatchOpts, sink chan<- *WfilDeposit, from []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}

	logs, sub, err := _Wfil.contract.WatchLogs(opts, "Deposit", fromRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WfilDeposit)
				if err := _Wfil.contract.UnpackLog(event, "Deposit", log); err != nil {
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

// ParseDeposit is a log parse operation binding the contract event 0xe1fffcc4923d04b559f4d29a8bfc6cda04eb5b0d3c460751c2402c5c5cc9109c.
//
// Solidity: event Deposit(address indexed from, uint256 amount)
func (_Wfil *WfilFilterer) ParseDeposit(log types.Log) (*WfilDeposit, error) {
	event := new(WfilDeposit)
	if err := _Wfil.contract.UnpackLog(event, "Deposit", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// WfilOwnershipPendingIterator is returned from FilterOwnershipPending and is used to iterate over the raw logs and unpacked data for OwnershipPending events raised by the Wfil contract.
type WfilOwnershipPendingIterator struct {
	Event *WfilOwnershipPending // Event containing the contract specifics and raw log

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
func (it *WfilOwnershipPendingIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WfilOwnershipPending)
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
		it.Event = new(WfilOwnershipPending)
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
func (it *WfilOwnershipPendingIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WfilOwnershipPendingIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WfilOwnershipPending represents a OwnershipPending event raised by the Wfil contract.
type WfilOwnershipPending struct {
	CurrentOwner common.Address
	PendingOwner common.Address
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterOwnershipPending is a free log retrieval operation binding the contract event 0x3c672e6c16e239c29e969deaf4eae15d7002350bf5490175e7481155af04c83a.
//
// Solidity: event OwnershipPending(address indexed currentOwner, address indexed pendingOwner)
func (_Wfil *WfilFilterer) FilterOwnershipPending(opts *bind.FilterOpts, currentOwner []common.Address, pendingOwner []common.Address) (*WfilOwnershipPendingIterator, error) {

	var currentOwnerRule []interface{}
	for _, currentOwnerItem := range currentOwner {
		currentOwnerRule = append(currentOwnerRule, currentOwnerItem)
	}
	var pendingOwnerRule []interface{}
	for _, pendingOwnerItem := range pendingOwner {
		pendingOwnerRule = append(pendingOwnerRule, pendingOwnerItem)
	}

	logs, sub, err := _Wfil.contract.FilterLogs(opts, "OwnershipPending", currentOwnerRule, pendingOwnerRule)
	if err != nil {
		return nil, err
	}
	return &WfilOwnershipPendingIterator{contract: _Wfil.contract, event: "OwnershipPending", logs: logs, sub: sub}, nil
}

// WatchOwnershipPending is a free log subscription operation binding the contract event 0x3c672e6c16e239c29e969deaf4eae15d7002350bf5490175e7481155af04c83a.
//
// Solidity: event OwnershipPending(address indexed currentOwner, address indexed pendingOwner)
func (_Wfil *WfilFilterer) WatchOwnershipPending(opts *bind.WatchOpts, sink chan<- *WfilOwnershipPending, currentOwner []common.Address, pendingOwner []common.Address) (event.Subscription, error) {

	var currentOwnerRule []interface{}
	for _, currentOwnerItem := range currentOwner {
		currentOwnerRule = append(currentOwnerRule, currentOwnerItem)
	}
	var pendingOwnerRule []interface{}
	for _, pendingOwnerItem := range pendingOwner {
		pendingOwnerRule = append(pendingOwnerRule, pendingOwnerItem)
	}

	logs, sub, err := _Wfil.contract.WatchLogs(opts, "OwnershipPending", currentOwnerRule, pendingOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WfilOwnershipPending)
				if err := _Wfil.contract.UnpackLog(event, "OwnershipPending", log); err != nil {
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

// ParseOwnershipPending is a log parse operation binding the contract event 0x3c672e6c16e239c29e969deaf4eae15d7002350bf5490175e7481155af04c83a.
//
// Solidity: event OwnershipPending(address indexed currentOwner, address indexed pendingOwner)
func (_Wfil *WfilFilterer) ParseOwnershipPending(log types.Log) (*WfilOwnershipPending, error) {
	event := new(WfilOwnershipPending)
	if err := _Wfil.contract.UnpackLog(event, "OwnershipPending", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// WfilOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Wfil contract.
type WfilOwnershipTransferredIterator struct {
	Event *WfilOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *WfilOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WfilOwnershipTransferred)
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
		it.Event = new(WfilOwnershipTransferred)
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
func (it *WfilOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WfilOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WfilOwnershipTransferred represents a OwnershipTransferred event raised by the Wfil contract.
type WfilOwnershipTransferred struct {
	OldOwner common.Address
	NewOwner common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed oldOwner, address indexed newOwner)
func (_Wfil *WfilFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, oldOwner []common.Address, newOwner []common.Address) (*WfilOwnershipTransferredIterator, error) {

	var oldOwnerRule []interface{}
	for _, oldOwnerItem := range oldOwner {
		oldOwnerRule = append(oldOwnerRule, oldOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Wfil.contract.FilterLogs(opts, "OwnershipTransferred", oldOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &WfilOwnershipTransferredIterator{contract: _Wfil.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed oldOwner, address indexed newOwner)
func (_Wfil *WfilFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *WfilOwnershipTransferred, oldOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var oldOwnerRule []interface{}
	for _, oldOwnerItem := range oldOwner {
		oldOwnerRule = append(oldOwnerRule, oldOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Wfil.contract.WatchLogs(opts, "OwnershipTransferred", oldOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WfilOwnershipTransferred)
				if err := _Wfil.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
// Solidity: event OwnershipTransferred(address indexed oldOwner, address indexed newOwner)
func (_Wfil *WfilFilterer) ParseOwnershipTransferred(log types.Log) (*WfilOwnershipTransferred, error) {
	event := new(WfilOwnershipTransferred)
	if err := _Wfil.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// WfilTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the Wfil contract.
type WfilTransferIterator struct {
	Event *WfilTransfer // Event containing the contract specifics and raw log

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
func (it *WfilTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WfilTransfer)
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
		it.Event = new(WfilTransfer)
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
func (it *WfilTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WfilTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WfilTransfer represents a Transfer event raised by the Wfil contract.
type WfilTransfer struct {
	From   common.Address
	To     common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 amount)
func (_Wfil *WfilFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*WfilTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Wfil.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &WfilTransferIterator{contract: _Wfil.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 amount)
func (_Wfil *WfilFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *WfilTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Wfil.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WfilTransfer)
				if err := _Wfil.contract.UnpackLog(event, "Transfer", log); err != nil {
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

// ParseTransfer is a log parse operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 amount)
func (_Wfil *WfilFilterer) ParseTransfer(log types.Log) (*WfilTransfer, error) {
	event := new(WfilTransfer)
	if err := _Wfil.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// WfilWithdrawalIterator is returned from FilterWithdrawal and is used to iterate over the raw logs and unpacked data for Withdrawal events raised by the Wfil contract.
type WfilWithdrawalIterator struct {
	Event *WfilWithdrawal // Event containing the contract specifics and raw log

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
func (it *WfilWithdrawalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WfilWithdrawal)
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
		it.Event = new(WfilWithdrawal)
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
func (it *WfilWithdrawalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WfilWithdrawalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WfilWithdrawal represents a Withdrawal event raised by the Wfil contract.
type WfilWithdrawal struct {
	To     common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterWithdrawal is a free log retrieval operation binding the contract event 0x7fcf532c15f0a6db0bd6d0e038bea71d30d808c7d98cb3bf7268a95bf5081b65.
//
// Solidity: event Withdrawal(address indexed to, uint256 amount)
func (_Wfil *WfilFilterer) FilterWithdrawal(opts *bind.FilterOpts, to []common.Address) (*WfilWithdrawalIterator, error) {

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Wfil.contract.FilterLogs(opts, "Withdrawal", toRule)
	if err != nil {
		return nil, err
	}
	return &WfilWithdrawalIterator{contract: _Wfil.contract, event: "Withdrawal", logs: logs, sub: sub}, nil
}

// WatchWithdrawal is a free log subscription operation binding the contract event 0x7fcf532c15f0a6db0bd6d0e038bea71d30d808c7d98cb3bf7268a95bf5081b65.
//
// Solidity: event Withdrawal(address indexed to, uint256 amount)
func (_Wfil *WfilFilterer) WatchWithdrawal(opts *bind.WatchOpts, sink chan<- *WfilWithdrawal, to []common.Address) (event.Subscription, error) {

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Wfil.contract.WatchLogs(opts, "Withdrawal", toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WfilWithdrawal)
				if err := _Wfil.contract.UnpackLog(event, "Withdrawal", log); err != nil {
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

// ParseWithdrawal is a log parse operation binding the contract event 0x7fcf532c15f0a6db0bd6d0e038bea71d30d808c7d98cb3bf7268a95bf5081b65.
//
// Solidity: event Withdrawal(address indexed to, uint256 amount)
func (_Wfil *WfilFilterer) ParseWithdrawal(log types.Log) (*WfilWithdrawal, error) {
	event := new(WfilWithdrawal)
	if err := _Wfil.contract.UnpackLog(event, "Withdrawal", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
