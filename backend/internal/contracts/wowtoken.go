// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contracts

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

// WoWTokenMetaData contains all meta data concerning the WoWToken contract.
var WoWTokenMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"creator\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"first_mint_amount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"reserve\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"length\",\"type\":\"uint256\"}],\"name\":\"StringsInsufficientHexLength\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"burner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"createdAt\",\"type\":\"uint256\"}],\"name\":\"BurnToken\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"id_event\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"isSucceeded\",\"type\":\"bool\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"createdAt\",\"type\":\"uint256\"}],\"name\":\"ExecuteAddCoOwner\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"id_event\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"isSucceeded\",\"type\":\"bool\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"createdAt\",\"type\":\"uint256\"}],\"name\":\"ExecuteBurnToken\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"id_event\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"isSucceeded\",\"type\":\"bool\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"createdAt\",\"type\":\"uint256\"}],\"name\":\"ExecuteMintToken\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"maker\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"new_co_owner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"id_event\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"createdAt\",\"type\":\"uint256\"}],\"name\":\"GenerateAddCoOwner\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"maker\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"id_event\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"createdAt\",\"type\":\"uint256\"}],\"name\":\"GenerateBurnToken\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"maker\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"id_event\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"createdAt\",\"type\":\"uint256\"}],\"name\":\"GenerateMintToken\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"createdAt\",\"type\":\"uint256\"}],\"name\":\"Pause\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"signer\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"id_event\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"createdAt\",\"type\":\"uint256\"}],\"name\":\"SignEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"createdAt\",\"type\":\"uint256\"}],\"name\":\"TransferToken\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"createdAt\",\"type\":\"uint256\"}],\"name\":\"Unpause\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"LIMIT_TOKEN\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"new_co_owner\",\"type\":\"address\"}],\"name\":\"addCoOwner\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"balances\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"bank_reserve\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"burnToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"co_owner_count\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"co_token_owner\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"current_total_token\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"eventSigners\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"name\":\"event_requireMultiSignature\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"id_event\",\"type\":\"string\"},{\"internalType\":\"bool\",\"name\":\"isCompleted\",\"type\":\"bool\"},{\"internalType\":\"string\",\"name\":\"event_name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"description\",\"type\":\"string\"},{\"internalType\":\"uint8\",\"name\":\"signature_count\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"createdAt\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"id_event\",\"type\":\"string\"}],\"name\":\"executeAddCoOwnerEvent\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"id_event\",\"type\":\"string\"}],\"name\":\"executeBurnTokenEvent\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"id_event\",\"type\":\"string\"}],\"name\":\"executeMintTokenEvent\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"id_event\",\"type\":\"string\"}],\"name\":\"executePauseEvent\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"id_event\",\"type\":\"string\"}],\"name\":\"executeUnpauseEvent\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"generateBurnTokenEvent\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"generatePauseEvent\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"generateUnpauseEvent\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"mintToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"id_event\",\"type\":\"string\"},{\"internalType\":\"bool\",\"name\":\"_approve\",\"type\":\"bool\"}],\"name\":\"signEvent\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"token_owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transferToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// WoWTokenABI is the input ABI used to generate the binding from.
// Deprecated: Use WoWTokenMetaData.ABI instead.
var WoWTokenABI = WoWTokenMetaData.ABI

// WoWToken is an auto generated Go binding around an Ethereum contract.
type WoWToken struct {
	WoWTokenCaller     // Read-only binding to the contract
	WoWTokenTransactor // Write-only binding to the contract
	WoWTokenFilterer   // Log filterer for contract events
}

// WoWTokenCaller is an auto generated read-only Go binding around an Ethereum contract.
type WoWTokenCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// WoWTokenTransactor is an auto generated write-only Go binding around an Ethereum contract.
type WoWTokenTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// WoWTokenFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type WoWTokenFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// WoWTokenSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type WoWTokenSession struct {
	Contract     *WoWToken         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// WoWTokenCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type WoWTokenCallerSession struct {
	Contract *WoWTokenCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// WoWTokenTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type WoWTokenTransactorSession struct {
	Contract     *WoWTokenTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// WoWTokenRaw is an auto generated low-level Go binding around an Ethereum contract.
type WoWTokenRaw struct {
	Contract *WoWToken // Generic contract binding to access the raw methods on
}

// WoWTokenCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type WoWTokenCallerRaw struct {
	Contract *WoWTokenCaller // Generic read-only contract binding to access the raw methods on
}

// WoWTokenTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type WoWTokenTransactorRaw struct {
	Contract *WoWTokenTransactor // Generic write-only contract binding to access the raw methods on
}

// NewWoWToken creates a new instance of WoWToken, bound to a specific deployed contract.
func NewWoWToken(address common.Address, backend bind.ContractBackend) (*WoWToken, error) {
	contract, err := bindWoWToken(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &WoWToken{WoWTokenCaller: WoWTokenCaller{contract: contract}, WoWTokenTransactor: WoWTokenTransactor{contract: contract}, WoWTokenFilterer: WoWTokenFilterer{contract: contract}}, nil
}

// NewWoWTokenCaller creates a new read-only instance of WoWToken, bound to a specific deployed contract.
func NewWoWTokenCaller(address common.Address, caller bind.ContractCaller) (*WoWTokenCaller, error) {
	contract, err := bindWoWToken(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &WoWTokenCaller{contract: contract}, nil
}

// NewWoWTokenTransactor creates a new write-only instance of WoWToken, bound to a specific deployed contract.
func NewWoWTokenTransactor(address common.Address, transactor bind.ContractTransactor) (*WoWTokenTransactor, error) {
	contract, err := bindWoWToken(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &WoWTokenTransactor{contract: contract}, nil
}

// NewWoWTokenFilterer creates a new log filterer instance of WoWToken, bound to a specific deployed contract.
func NewWoWTokenFilterer(address common.Address, filterer bind.ContractFilterer) (*WoWTokenFilterer, error) {
	contract, err := bindWoWToken(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &WoWTokenFilterer{contract: contract}, nil
}

// bindWoWToken binds a generic wrapper to an already deployed contract.
func bindWoWToken(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := WoWTokenMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_WoWToken *WoWTokenRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _WoWToken.Contract.WoWTokenCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_WoWToken *WoWTokenRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _WoWToken.Contract.WoWTokenTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_WoWToken *WoWTokenRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _WoWToken.Contract.WoWTokenTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_WoWToken *WoWTokenCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _WoWToken.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_WoWToken *WoWTokenTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _WoWToken.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_WoWToken *WoWTokenTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _WoWToken.Contract.contract.Transact(opts, method, params...)
}

// LIMITTOKEN is a free data retrieval call binding the contract method 0xc8eaba7e.
//
// Solidity: function LIMIT_TOKEN() view returns(uint256)
func (_WoWToken *WoWTokenCaller) LIMITTOKEN(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _WoWToken.contract.Call(opts, &out, "LIMIT_TOKEN")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LIMITTOKEN is a free data retrieval call binding the contract method 0xc8eaba7e.
//
// Solidity: function LIMIT_TOKEN() view returns(uint256)
func (_WoWToken *WoWTokenSession) LIMITTOKEN() (*big.Int, error) {
	return _WoWToken.Contract.LIMITTOKEN(&_WoWToken.CallOpts)
}

// LIMITTOKEN is a free data retrieval call binding the contract method 0xc8eaba7e.
//
// Solidity: function LIMIT_TOKEN() view returns(uint256)
func (_WoWToken *WoWTokenCallerSession) LIMITTOKEN() (*big.Int, error) {
	return _WoWToken.Contract.LIMITTOKEN(&_WoWToken.CallOpts)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address , address ) view returns(uint256)
func (_WoWToken *WoWTokenCaller) Allowance(opts *bind.CallOpts, arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _WoWToken.contract.Call(opts, &out, "allowance", arg0, arg1)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address , address ) view returns(uint256)
func (_WoWToken *WoWTokenSession) Allowance(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _WoWToken.Contract.Allowance(&_WoWToken.CallOpts, arg0, arg1)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address , address ) view returns(uint256)
func (_WoWToken *WoWTokenCallerSession) Allowance(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _WoWToken.Contract.Allowance(&_WoWToken.CallOpts, arg0, arg1)
}

// Balances is a free data retrieval call binding the contract method 0x27e235e3.
//
// Solidity: function balances(address ) view returns(uint256)
func (_WoWToken *WoWTokenCaller) Balances(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _WoWToken.contract.Call(opts, &out, "balances", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Balances is a free data retrieval call binding the contract method 0x27e235e3.
//
// Solidity: function balances(address ) view returns(uint256)
func (_WoWToken *WoWTokenSession) Balances(arg0 common.Address) (*big.Int, error) {
	return _WoWToken.Contract.Balances(&_WoWToken.CallOpts, arg0)
}

// Balances is a free data retrieval call binding the contract method 0x27e235e3.
//
// Solidity: function balances(address ) view returns(uint256)
func (_WoWToken *WoWTokenCallerSession) Balances(arg0 common.Address) (*big.Int, error) {
	return _WoWToken.Contract.Balances(&_WoWToken.CallOpts, arg0)
}

// BankReserve is a free data retrieval call binding the contract method 0xda468a1b.
//
// Solidity: function bank_reserve() view returns(address)
func (_WoWToken *WoWTokenCaller) BankReserve(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _WoWToken.contract.Call(opts, &out, "bank_reserve")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// BankReserve is a free data retrieval call binding the contract method 0xda468a1b.
//
// Solidity: function bank_reserve() view returns(address)
func (_WoWToken *WoWTokenSession) BankReserve() (common.Address, error) {
	return _WoWToken.Contract.BankReserve(&_WoWToken.CallOpts)
}

// BankReserve is a free data retrieval call binding the contract method 0xda468a1b.
//
// Solidity: function bank_reserve() view returns(address)
func (_WoWToken *WoWTokenCallerSession) BankReserve() (common.Address, error) {
	return _WoWToken.Contract.BankReserve(&_WoWToken.CallOpts)
}

// CoOwnerCount is a free data retrieval call binding the contract method 0x716a48d0.
//
// Solidity: function co_owner_count() view returns(uint8)
func (_WoWToken *WoWTokenCaller) CoOwnerCount(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _WoWToken.contract.Call(opts, &out, "co_owner_count")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// CoOwnerCount is a free data retrieval call binding the contract method 0x716a48d0.
//
// Solidity: function co_owner_count() view returns(uint8)
func (_WoWToken *WoWTokenSession) CoOwnerCount() (uint8, error) {
	return _WoWToken.Contract.CoOwnerCount(&_WoWToken.CallOpts)
}

// CoOwnerCount is a free data retrieval call binding the contract method 0x716a48d0.
//
// Solidity: function co_owner_count() view returns(uint8)
func (_WoWToken *WoWTokenCallerSession) CoOwnerCount() (uint8, error) {
	return _WoWToken.Contract.CoOwnerCount(&_WoWToken.CallOpts)
}

// CoTokenOwner is a free data retrieval call binding the contract method 0xfb121028.
//
// Solidity: function co_token_owner(address ) view returns(bool)
func (_WoWToken *WoWTokenCaller) CoTokenOwner(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _WoWToken.contract.Call(opts, &out, "co_token_owner", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// CoTokenOwner is a free data retrieval call binding the contract method 0xfb121028.
//
// Solidity: function co_token_owner(address ) view returns(bool)
func (_WoWToken *WoWTokenSession) CoTokenOwner(arg0 common.Address) (bool, error) {
	return _WoWToken.Contract.CoTokenOwner(&_WoWToken.CallOpts, arg0)
}

// CoTokenOwner is a free data retrieval call binding the contract method 0xfb121028.
//
// Solidity: function co_token_owner(address ) view returns(bool)
func (_WoWToken *WoWTokenCallerSession) CoTokenOwner(arg0 common.Address) (bool, error) {
	return _WoWToken.Contract.CoTokenOwner(&_WoWToken.CallOpts, arg0)
}

// CurrentTotalToken is a free data retrieval call binding the contract method 0xcc0b2ea4.
//
// Solidity: function current_total_token() view returns(uint256)
func (_WoWToken *WoWTokenCaller) CurrentTotalToken(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _WoWToken.contract.Call(opts, &out, "current_total_token")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CurrentTotalToken is a free data retrieval call binding the contract method 0xcc0b2ea4.
//
// Solidity: function current_total_token() view returns(uint256)
func (_WoWToken *WoWTokenSession) CurrentTotalToken() (*big.Int, error) {
	return _WoWToken.Contract.CurrentTotalToken(&_WoWToken.CallOpts)
}

// CurrentTotalToken is a free data retrieval call binding the contract method 0xcc0b2ea4.
//
// Solidity: function current_total_token() view returns(uint256)
func (_WoWToken *WoWTokenCallerSession) CurrentTotalToken() (*big.Int, error) {
	return _WoWToken.Contract.CurrentTotalToken(&_WoWToken.CallOpts)
}

// EventSigners is a free data retrieval call binding the contract method 0x8f374935.
//
// Solidity: function eventSigners(string , address ) view returns(bool)
func (_WoWToken *WoWTokenCaller) EventSigners(opts *bind.CallOpts, arg0 string, arg1 common.Address) (bool, error) {
	var out []interface{}
	err := _WoWToken.contract.Call(opts, &out, "eventSigners", arg0, arg1)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// EventSigners is a free data retrieval call binding the contract method 0x8f374935.
//
// Solidity: function eventSigners(string , address ) view returns(bool)
func (_WoWToken *WoWTokenSession) EventSigners(arg0 string, arg1 common.Address) (bool, error) {
	return _WoWToken.Contract.EventSigners(&_WoWToken.CallOpts, arg0, arg1)
}

// EventSigners is a free data retrieval call binding the contract method 0x8f374935.
//
// Solidity: function eventSigners(string , address ) view returns(bool)
func (_WoWToken *WoWTokenCallerSession) EventSigners(arg0 string, arg1 common.Address) (bool, error) {
	return _WoWToken.Contract.EventSigners(&_WoWToken.CallOpts, arg0, arg1)
}

// EventRequireMultiSignature is a free data retrieval call binding the contract method 0x0512b3ed.
//
// Solidity: function event_requireMultiSignature(string ) view returns(string id_event, bool isCompleted, string event_name, string description, uint8 signature_count, uint256 createdAt)
func (_WoWToken *WoWTokenCaller) EventRequireMultiSignature(opts *bind.CallOpts, arg0 string) (struct {
	IdEvent        string
	IsCompleted    bool
	EventName      string
	Description    string
	SignatureCount uint8
	CreatedAt      *big.Int
}, error) {
	var out []interface{}
	err := _WoWToken.contract.Call(opts, &out, "event_requireMultiSignature", arg0)

	outstruct := new(struct {
		IdEvent        string
		IsCompleted    bool
		EventName      string
		Description    string
		SignatureCount uint8
		CreatedAt      *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.IdEvent = *abi.ConvertType(out[0], new(string)).(*string)
	outstruct.IsCompleted = *abi.ConvertType(out[1], new(bool)).(*bool)
	outstruct.EventName = *abi.ConvertType(out[2], new(string)).(*string)
	outstruct.Description = *abi.ConvertType(out[3], new(string)).(*string)
	outstruct.SignatureCount = *abi.ConvertType(out[4], new(uint8)).(*uint8)
	outstruct.CreatedAt = *abi.ConvertType(out[5], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// EventRequireMultiSignature is a free data retrieval call binding the contract method 0x0512b3ed.
//
// Solidity: function event_requireMultiSignature(string ) view returns(string id_event, bool isCompleted, string event_name, string description, uint8 signature_count, uint256 createdAt)
func (_WoWToken *WoWTokenSession) EventRequireMultiSignature(arg0 string) (struct {
	IdEvent        string
	IsCompleted    bool
	EventName      string
	Description    string
	SignatureCount uint8
	CreatedAt      *big.Int
}, error) {
	return _WoWToken.Contract.EventRequireMultiSignature(&_WoWToken.CallOpts, arg0)
}

// EventRequireMultiSignature is a free data retrieval call binding the contract method 0x0512b3ed.
//
// Solidity: function event_requireMultiSignature(string ) view returns(string id_event, bool isCompleted, string event_name, string description, uint8 signature_count, uint256 createdAt)
func (_WoWToken *WoWTokenCallerSession) EventRequireMultiSignature(arg0 string) (struct {
	IdEvent        string
	IsCompleted    bool
	EventName      string
	Description    string
	SignatureCount uint8
	CreatedAt      *big.Int
}, error) {
	return _WoWToken.Contract.EventRequireMultiSignature(&_WoWToken.CallOpts, arg0)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_WoWToken *WoWTokenCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _WoWToken.contract.Call(opts, &out, "paused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_WoWToken *WoWTokenSession) Paused() (bool, error) {
	return _WoWToken.Contract.Paused(&_WoWToken.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_WoWToken *WoWTokenCallerSession) Paused() (bool, error) {
	return _WoWToken.Contract.Paused(&_WoWToken.CallOpts)
}

// TokenOwner is a free data retrieval call binding the contract method 0x6832bfac.
//
// Solidity: function token_owner() view returns(address)
func (_WoWToken *WoWTokenCaller) TokenOwner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _WoWToken.contract.Call(opts, &out, "token_owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// TokenOwner is a free data retrieval call binding the contract method 0x6832bfac.
//
// Solidity: function token_owner() view returns(address)
func (_WoWToken *WoWTokenSession) TokenOwner() (common.Address, error) {
	return _WoWToken.Contract.TokenOwner(&_WoWToken.CallOpts)
}

// TokenOwner is a free data retrieval call binding the contract method 0x6832bfac.
//
// Solidity: function token_owner() view returns(address)
func (_WoWToken *WoWTokenCallerSession) TokenOwner() (common.Address, error) {
	return _WoWToken.Contract.TokenOwner(&_WoWToken.CallOpts)
}

// AddCoOwner is a paid mutator transaction binding the contract method 0x563a4024.
//
// Solidity: function addCoOwner(address new_co_owner) returns()
func (_WoWToken *WoWTokenTransactor) AddCoOwner(opts *bind.TransactOpts, new_co_owner common.Address) (*types.Transaction, error) {
	return _WoWToken.contract.Transact(opts, "addCoOwner", new_co_owner)
}

// AddCoOwner is a paid mutator transaction binding the contract method 0x563a4024.
//
// Solidity: function addCoOwner(address new_co_owner) returns()
func (_WoWToken *WoWTokenSession) AddCoOwner(new_co_owner common.Address) (*types.Transaction, error) {
	return _WoWToken.Contract.AddCoOwner(&_WoWToken.TransactOpts, new_co_owner)
}

// AddCoOwner is a paid mutator transaction binding the contract method 0x563a4024.
//
// Solidity: function addCoOwner(address new_co_owner) returns()
func (_WoWToken *WoWTokenTransactorSession) AddCoOwner(new_co_owner common.Address) (*types.Transaction, error) {
	return _WoWToken.Contract.AddCoOwner(&_WoWToken.TransactOpts, new_co_owner)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_WoWToken *WoWTokenTransactor) Approve(opts *bind.TransactOpts, spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _WoWToken.contract.Transact(opts, "approve", spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_WoWToken *WoWTokenSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _WoWToken.Contract.Approve(&_WoWToken.TransactOpts, spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_WoWToken *WoWTokenTransactorSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _WoWToken.Contract.Approve(&_WoWToken.TransactOpts, spender, amount)
}

// BurnToken is a paid mutator transaction binding the contract method 0x7b47ec1a.
//
// Solidity: function burnToken(uint256 amount) returns()
func (_WoWToken *WoWTokenTransactor) BurnToken(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _WoWToken.contract.Transact(opts, "burnToken", amount)
}

// BurnToken is a paid mutator transaction binding the contract method 0x7b47ec1a.
//
// Solidity: function burnToken(uint256 amount) returns()
func (_WoWToken *WoWTokenSession) BurnToken(amount *big.Int) (*types.Transaction, error) {
	return _WoWToken.Contract.BurnToken(&_WoWToken.TransactOpts, amount)
}

// BurnToken is a paid mutator transaction binding the contract method 0x7b47ec1a.
//
// Solidity: function burnToken(uint256 amount) returns()
func (_WoWToken *WoWTokenTransactorSession) BurnToken(amount *big.Int) (*types.Transaction, error) {
	return _WoWToken.Contract.BurnToken(&_WoWToken.TransactOpts, amount)
}

// ExecuteAddCoOwnerEvent is a paid mutator transaction binding the contract method 0xa927b6fe.
//
// Solidity: function executeAddCoOwnerEvent(string id_event) returns()
func (_WoWToken *WoWTokenTransactor) ExecuteAddCoOwnerEvent(opts *bind.TransactOpts, id_event string) (*types.Transaction, error) {
	return _WoWToken.contract.Transact(opts, "executeAddCoOwnerEvent", id_event)
}

// ExecuteAddCoOwnerEvent is a paid mutator transaction binding the contract method 0xa927b6fe.
//
// Solidity: function executeAddCoOwnerEvent(string id_event) returns()
func (_WoWToken *WoWTokenSession) ExecuteAddCoOwnerEvent(id_event string) (*types.Transaction, error) {
	return _WoWToken.Contract.ExecuteAddCoOwnerEvent(&_WoWToken.TransactOpts, id_event)
}

// ExecuteAddCoOwnerEvent is a paid mutator transaction binding the contract method 0xa927b6fe.
//
// Solidity: function executeAddCoOwnerEvent(string id_event) returns()
func (_WoWToken *WoWTokenTransactorSession) ExecuteAddCoOwnerEvent(id_event string) (*types.Transaction, error) {
	return _WoWToken.Contract.ExecuteAddCoOwnerEvent(&_WoWToken.TransactOpts, id_event)
}

// ExecuteBurnTokenEvent is a paid mutator transaction binding the contract method 0x4f36adf1.
//
// Solidity: function executeBurnTokenEvent(string id_event) returns()
func (_WoWToken *WoWTokenTransactor) ExecuteBurnTokenEvent(opts *bind.TransactOpts, id_event string) (*types.Transaction, error) {
	return _WoWToken.contract.Transact(opts, "executeBurnTokenEvent", id_event)
}

// ExecuteBurnTokenEvent is a paid mutator transaction binding the contract method 0x4f36adf1.
//
// Solidity: function executeBurnTokenEvent(string id_event) returns()
func (_WoWToken *WoWTokenSession) ExecuteBurnTokenEvent(id_event string) (*types.Transaction, error) {
	return _WoWToken.Contract.ExecuteBurnTokenEvent(&_WoWToken.TransactOpts, id_event)
}

// ExecuteBurnTokenEvent is a paid mutator transaction binding the contract method 0x4f36adf1.
//
// Solidity: function executeBurnTokenEvent(string id_event) returns()
func (_WoWToken *WoWTokenTransactorSession) ExecuteBurnTokenEvent(id_event string) (*types.Transaction, error) {
	return _WoWToken.Contract.ExecuteBurnTokenEvent(&_WoWToken.TransactOpts, id_event)
}

// ExecuteMintTokenEvent is a paid mutator transaction binding the contract method 0x151b38e0.
//
// Solidity: function executeMintTokenEvent(string id_event) returns()
func (_WoWToken *WoWTokenTransactor) ExecuteMintTokenEvent(opts *bind.TransactOpts, id_event string) (*types.Transaction, error) {
	return _WoWToken.contract.Transact(opts, "executeMintTokenEvent", id_event)
}

// ExecuteMintTokenEvent is a paid mutator transaction binding the contract method 0x151b38e0.
//
// Solidity: function executeMintTokenEvent(string id_event) returns()
func (_WoWToken *WoWTokenSession) ExecuteMintTokenEvent(id_event string) (*types.Transaction, error) {
	return _WoWToken.Contract.ExecuteMintTokenEvent(&_WoWToken.TransactOpts, id_event)
}

// ExecuteMintTokenEvent is a paid mutator transaction binding the contract method 0x151b38e0.
//
// Solidity: function executeMintTokenEvent(string id_event) returns()
func (_WoWToken *WoWTokenTransactorSession) ExecuteMintTokenEvent(id_event string) (*types.Transaction, error) {
	return _WoWToken.Contract.ExecuteMintTokenEvent(&_WoWToken.TransactOpts, id_event)
}

// ExecutePauseEvent is a paid mutator transaction binding the contract method 0xd61013ea.
//
// Solidity: function executePauseEvent(string id_event) returns()
func (_WoWToken *WoWTokenTransactor) ExecutePauseEvent(opts *bind.TransactOpts, id_event string) (*types.Transaction, error) {
	return _WoWToken.contract.Transact(opts, "executePauseEvent", id_event)
}

// ExecutePauseEvent is a paid mutator transaction binding the contract method 0xd61013ea.
//
// Solidity: function executePauseEvent(string id_event) returns()
func (_WoWToken *WoWTokenSession) ExecutePauseEvent(id_event string) (*types.Transaction, error) {
	return _WoWToken.Contract.ExecutePauseEvent(&_WoWToken.TransactOpts, id_event)
}

// ExecutePauseEvent is a paid mutator transaction binding the contract method 0xd61013ea.
//
// Solidity: function executePauseEvent(string id_event) returns()
func (_WoWToken *WoWTokenTransactorSession) ExecutePauseEvent(id_event string) (*types.Transaction, error) {
	return _WoWToken.Contract.ExecutePauseEvent(&_WoWToken.TransactOpts, id_event)
}

// ExecuteUnpauseEvent is a paid mutator transaction binding the contract method 0x8b04bdd6.
//
// Solidity: function executeUnpauseEvent(string id_event) returns()
func (_WoWToken *WoWTokenTransactor) ExecuteUnpauseEvent(opts *bind.TransactOpts, id_event string) (*types.Transaction, error) {
	return _WoWToken.contract.Transact(opts, "executeUnpauseEvent", id_event)
}

// ExecuteUnpauseEvent is a paid mutator transaction binding the contract method 0x8b04bdd6.
//
// Solidity: function executeUnpauseEvent(string id_event) returns()
func (_WoWToken *WoWTokenSession) ExecuteUnpauseEvent(id_event string) (*types.Transaction, error) {
	return _WoWToken.Contract.ExecuteUnpauseEvent(&_WoWToken.TransactOpts, id_event)
}

// ExecuteUnpauseEvent is a paid mutator transaction binding the contract method 0x8b04bdd6.
//
// Solidity: function executeUnpauseEvent(string id_event) returns()
func (_WoWToken *WoWTokenTransactorSession) ExecuteUnpauseEvent(id_event string) (*types.Transaction, error) {
	return _WoWToken.Contract.ExecuteUnpauseEvent(&_WoWToken.TransactOpts, id_event)
}

// GenerateBurnTokenEvent is a paid mutator transaction binding the contract method 0x1ae02ad4.
//
// Solidity: function generateBurnTokenEvent(uint256 amount) returns()
func (_WoWToken *WoWTokenTransactor) GenerateBurnTokenEvent(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _WoWToken.contract.Transact(opts, "generateBurnTokenEvent", amount)
}

// GenerateBurnTokenEvent is a paid mutator transaction binding the contract method 0x1ae02ad4.
//
// Solidity: function generateBurnTokenEvent(uint256 amount) returns()
func (_WoWToken *WoWTokenSession) GenerateBurnTokenEvent(amount *big.Int) (*types.Transaction, error) {
	return _WoWToken.Contract.GenerateBurnTokenEvent(&_WoWToken.TransactOpts, amount)
}

// GenerateBurnTokenEvent is a paid mutator transaction binding the contract method 0x1ae02ad4.
//
// Solidity: function generateBurnTokenEvent(uint256 amount) returns()
func (_WoWToken *WoWTokenTransactorSession) GenerateBurnTokenEvent(amount *big.Int) (*types.Transaction, error) {
	return _WoWToken.Contract.GenerateBurnTokenEvent(&_WoWToken.TransactOpts, amount)
}

// GeneratePauseEvent is a paid mutator transaction binding the contract method 0xc0ebf322.
//
// Solidity: function generatePauseEvent() returns()
func (_WoWToken *WoWTokenTransactor) GeneratePauseEvent(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _WoWToken.contract.Transact(opts, "generatePauseEvent")
}

// GeneratePauseEvent is a paid mutator transaction binding the contract method 0xc0ebf322.
//
// Solidity: function generatePauseEvent() returns()
func (_WoWToken *WoWTokenSession) GeneratePauseEvent() (*types.Transaction, error) {
	return _WoWToken.Contract.GeneratePauseEvent(&_WoWToken.TransactOpts)
}

// GeneratePauseEvent is a paid mutator transaction binding the contract method 0xc0ebf322.
//
// Solidity: function generatePauseEvent() returns()
func (_WoWToken *WoWTokenTransactorSession) GeneratePauseEvent() (*types.Transaction, error) {
	return _WoWToken.Contract.GeneratePauseEvent(&_WoWToken.TransactOpts)
}

// GenerateUnpauseEvent is a paid mutator transaction binding the contract method 0xfb8632e1.
//
// Solidity: function generateUnpauseEvent() returns()
func (_WoWToken *WoWTokenTransactor) GenerateUnpauseEvent(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _WoWToken.contract.Transact(opts, "generateUnpauseEvent")
}

// GenerateUnpauseEvent is a paid mutator transaction binding the contract method 0xfb8632e1.
//
// Solidity: function generateUnpauseEvent() returns()
func (_WoWToken *WoWTokenSession) GenerateUnpauseEvent() (*types.Transaction, error) {
	return _WoWToken.Contract.GenerateUnpauseEvent(&_WoWToken.TransactOpts)
}

// GenerateUnpauseEvent is a paid mutator transaction binding the contract method 0xfb8632e1.
//
// Solidity: function generateUnpauseEvent() returns()
func (_WoWToken *WoWTokenTransactorSession) GenerateUnpauseEvent() (*types.Transaction, error) {
	return _WoWToken.Contract.GenerateUnpauseEvent(&_WoWToken.TransactOpts)
}

// MintToken is a paid mutator transaction binding the contract method 0xc634d032.
//
// Solidity: function mintToken(uint256 amount) returns()
func (_WoWToken *WoWTokenTransactor) MintToken(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _WoWToken.contract.Transact(opts, "mintToken", amount)
}

// MintToken is a paid mutator transaction binding the contract method 0xc634d032.
//
// Solidity: function mintToken(uint256 amount) returns()
func (_WoWToken *WoWTokenSession) MintToken(amount *big.Int) (*types.Transaction, error) {
	return _WoWToken.Contract.MintToken(&_WoWToken.TransactOpts, amount)
}

// MintToken is a paid mutator transaction binding the contract method 0xc634d032.
//
// Solidity: function mintToken(uint256 amount) returns()
func (_WoWToken *WoWTokenTransactorSession) MintToken(amount *big.Int) (*types.Transaction, error) {
	return _WoWToken.Contract.MintToken(&_WoWToken.TransactOpts, amount)
}

// SignEvent is a paid mutator transaction binding the contract method 0x0f6e2d29.
//
// Solidity: function signEvent(string id_event, bool _approve) returns()
func (_WoWToken *WoWTokenTransactor) SignEvent(opts *bind.TransactOpts, id_event string, _approve bool) (*types.Transaction, error) {
	return _WoWToken.contract.Transact(opts, "signEvent", id_event, _approve)
}

// SignEvent is a paid mutator transaction binding the contract method 0x0f6e2d29.
//
// Solidity: function signEvent(string id_event, bool _approve) returns()
func (_WoWToken *WoWTokenSession) SignEvent(id_event string, _approve bool) (*types.Transaction, error) {
	return _WoWToken.Contract.SignEvent(&_WoWToken.TransactOpts, id_event, _approve)
}

// SignEvent is a paid mutator transaction binding the contract method 0x0f6e2d29.
//
// Solidity: function signEvent(string id_event, bool _approve) returns()
func (_WoWToken *WoWTokenTransactorSession) SignEvent(id_event string, _approve bool) (*types.Transaction, error) {
	return _WoWToken.Contract.SignEvent(&_WoWToken.TransactOpts, id_event, _approve)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 amount) returns(bool)
func (_WoWToken *WoWTokenTransactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _WoWToken.contract.Transact(opts, "transferFrom", from, to, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 amount) returns(bool)
func (_WoWToken *WoWTokenSession) TransferFrom(from common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _WoWToken.Contract.TransferFrom(&_WoWToken.TransactOpts, from, to, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 amount) returns(bool)
func (_WoWToken *WoWTokenTransactorSession) TransferFrom(from common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _WoWToken.Contract.TransferFrom(&_WoWToken.TransactOpts, from, to, amount)
}

// TransferToken is a paid mutator transaction binding the contract method 0x1072cbea.
//
// Solidity: function transferToken(address to, uint256 amount) returns()
func (_WoWToken *WoWTokenTransactor) TransferToken(opts *bind.TransactOpts, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _WoWToken.contract.Transact(opts, "transferToken", to, amount)
}

// TransferToken is a paid mutator transaction binding the contract method 0x1072cbea.
//
// Solidity: function transferToken(address to, uint256 amount) returns()
func (_WoWToken *WoWTokenSession) TransferToken(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _WoWToken.Contract.TransferToken(&_WoWToken.TransactOpts, to, amount)
}

// TransferToken is a paid mutator transaction binding the contract method 0x1072cbea.
//
// Solidity: function transferToken(address to, uint256 amount) returns()
func (_WoWToken *WoWTokenTransactorSession) TransferToken(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _WoWToken.Contract.TransferToken(&_WoWToken.TransactOpts, to, amount)
}

// WoWTokenApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the WoWToken contract.
type WoWTokenApprovalIterator struct {
	Event *WoWTokenApproval // Event containing the contract specifics and raw log

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
func (it *WoWTokenApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WoWTokenApproval)
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
		it.Event = new(WoWTokenApproval)
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
func (it *WoWTokenApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WoWTokenApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WoWTokenApproval represents a Approval event raised by the WoWToken contract.
type WoWTokenApproval struct {
	Owner   common.Address
	Spender common.Address
	Amount  *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 amount)
func (_WoWToken *WoWTokenFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*WoWTokenApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _WoWToken.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &WoWTokenApprovalIterator{contract: _WoWToken.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 amount)
func (_WoWToken *WoWTokenFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *WoWTokenApproval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _WoWToken.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WoWTokenApproval)
				if err := _WoWToken.contract.UnpackLog(event, "Approval", log); err != nil {
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
func (_WoWToken *WoWTokenFilterer) ParseApproval(log types.Log) (*WoWTokenApproval, error) {
	event := new(WoWTokenApproval)
	if err := _WoWToken.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// WoWTokenBurnTokenIterator is returned from FilterBurnToken and is used to iterate over the raw logs and unpacked data for BurnToken events raised by the WoWToken contract.
type WoWTokenBurnTokenIterator struct {
	Event *WoWTokenBurnToken // Event containing the contract specifics and raw log

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
func (it *WoWTokenBurnTokenIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WoWTokenBurnToken)
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
		it.Event = new(WoWTokenBurnToken)
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
func (it *WoWTokenBurnTokenIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WoWTokenBurnTokenIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WoWTokenBurnToken represents a BurnToken event raised by the WoWToken contract.
type WoWTokenBurnToken struct {
	Burner    common.Address
	Amount    *big.Int
	CreatedAt *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterBurnToken is a free log retrieval operation binding the contract event 0x9101ffdf7f446a8cd01ffe1fa15674f7fae32d7e5c8df3b3a5f0612b724f3a80.
//
// Solidity: event BurnToken(address indexed burner, uint256 amount, uint256 createdAt)
func (_WoWToken *WoWTokenFilterer) FilterBurnToken(opts *bind.FilterOpts, burner []common.Address) (*WoWTokenBurnTokenIterator, error) {

	var burnerRule []interface{}
	for _, burnerItem := range burner {
		burnerRule = append(burnerRule, burnerItem)
	}

	logs, sub, err := _WoWToken.contract.FilterLogs(opts, "BurnToken", burnerRule)
	if err != nil {
		return nil, err
	}
	return &WoWTokenBurnTokenIterator{contract: _WoWToken.contract, event: "BurnToken", logs: logs, sub: sub}, nil
}

// WatchBurnToken is a free log subscription operation binding the contract event 0x9101ffdf7f446a8cd01ffe1fa15674f7fae32d7e5c8df3b3a5f0612b724f3a80.
//
// Solidity: event BurnToken(address indexed burner, uint256 amount, uint256 createdAt)
func (_WoWToken *WoWTokenFilterer) WatchBurnToken(opts *bind.WatchOpts, sink chan<- *WoWTokenBurnToken, burner []common.Address) (event.Subscription, error) {

	var burnerRule []interface{}
	for _, burnerItem := range burner {
		burnerRule = append(burnerRule, burnerItem)
	}

	logs, sub, err := _WoWToken.contract.WatchLogs(opts, "BurnToken", burnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WoWTokenBurnToken)
				if err := _WoWToken.contract.UnpackLog(event, "BurnToken", log); err != nil {
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

// ParseBurnToken is a log parse operation binding the contract event 0x9101ffdf7f446a8cd01ffe1fa15674f7fae32d7e5c8df3b3a5f0612b724f3a80.
//
// Solidity: event BurnToken(address indexed burner, uint256 amount, uint256 createdAt)
func (_WoWToken *WoWTokenFilterer) ParseBurnToken(log types.Log) (*WoWTokenBurnToken, error) {
	event := new(WoWTokenBurnToken)
	if err := _WoWToken.contract.UnpackLog(event, "BurnToken", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// WoWTokenExecuteAddCoOwnerIterator is returned from FilterExecuteAddCoOwner and is used to iterate over the raw logs and unpacked data for ExecuteAddCoOwner events raised by the WoWToken contract.
type WoWTokenExecuteAddCoOwnerIterator struct {
	Event *WoWTokenExecuteAddCoOwner // Event containing the contract specifics and raw log

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
func (it *WoWTokenExecuteAddCoOwnerIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WoWTokenExecuteAddCoOwner)
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
		it.Event = new(WoWTokenExecuteAddCoOwner)
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
func (it *WoWTokenExecuteAddCoOwnerIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WoWTokenExecuteAddCoOwnerIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WoWTokenExecuteAddCoOwner represents a ExecuteAddCoOwner event raised by the WoWToken contract.
type WoWTokenExecuteAddCoOwner struct {
	IdEvent     string
	IsSucceeded bool
	CreatedAt   *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterExecuteAddCoOwner is a free log retrieval operation binding the contract event 0x71114ec6bb771939fcfa6679decb7554d976cfba8c7b420b9614d8864424e31b.
//
// Solidity: event ExecuteAddCoOwner(string id_event, bool isSucceeded, uint256 createdAt)
func (_WoWToken *WoWTokenFilterer) FilterExecuteAddCoOwner(opts *bind.FilterOpts) (*WoWTokenExecuteAddCoOwnerIterator, error) {

	logs, sub, err := _WoWToken.contract.FilterLogs(opts, "ExecuteAddCoOwner")
	if err != nil {
		return nil, err
	}
	return &WoWTokenExecuteAddCoOwnerIterator{contract: _WoWToken.contract, event: "ExecuteAddCoOwner", logs: logs, sub: sub}, nil
}

// WatchExecuteAddCoOwner is a free log subscription operation binding the contract event 0x71114ec6bb771939fcfa6679decb7554d976cfba8c7b420b9614d8864424e31b.
//
// Solidity: event ExecuteAddCoOwner(string id_event, bool isSucceeded, uint256 createdAt)
func (_WoWToken *WoWTokenFilterer) WatchExecuteAddCoOwner(opts *bind.WatchOpts, sink chan<- *WoWTokenExecuteAddCoOwner) (event.Subscription, error) {

	logs, sub, err := _WoWToken.contract.WatchLogs(opts, "ExecuteAddCoOwner")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WoWTokenExecuteAddCoOwner)
				if err := _WoWToken.contract.UnpackLog(event, "ExecuteAddCoOwner", log); err != nil {
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

// ParseExecuteAddCoOwner is a log parse operation binding the contract event 0x71114ec6bb771939fcfa6679decb7554d976cfba8c7b420b9614d8864424e31b.
//
// Solidity: event ExecuteAddCoOwner(string id_event, bool isSucceeded, uint256 createdAt)
func (_WoWToken *WoWTokenFilterer) ParseExecuteAddCoOwner(log types.Log) (*WoWTokenExecuteAddCoOwner, error) {
	event := new(WoWTokenExecuteAddCoOwner)
	if err := _WoWToken.contract.UnpackLog(event, "ExecuteAddCoOwner", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// WoWTokenExecuteBurnTokenIterator is returned from FilterExecuteBurnToken and is used to iterate over the raw logs and unpacked data for ExecuteBurnToken events raised by the WoWToken contract.
type WoWTokenExecuteBurnTokenIterator struct {
	Event *WoWTokenExecuteBurnToken // Event containing the contract specifics and raw log

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
func (it *WoWTokenExecuteBurnTokenIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WoWTokenExecuteBurnToken)
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
		it.Event = new(WoWTokenExecuteBurnToken)
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
func (it *WoWTokenExecuteBurnTokenIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WoWTokenExecuteBurnTokenIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WoWTokenExecuteBurnToken represents a ExecuteBurnToken event raised by the WoWToken contract.
type WoWTokenExecuteBurnToken struct {
	IdEvent     string
	Amount      *big.Int
	IsSucceeded bool
	CreatedAt   *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterExecuteBurnToken is a free log retrieval operation binding the contract event 0x30b07a2cb65e926e0ff9d6b29c7575d1da894476921a1d08596337df827e5fec.
//
// Solidity: event ExecuteBurnToken(string id_event, uint256 amount, bool isSucceeded, uint256 createdAt)
func (_WoWToken *WoWTokenFilterer) FilterExecuteBurnToken(opts *bind.FilterOpts) (*WoWTokenExecuteBurnTokenIterator, error) {

	logs, sub, err := _WoWToken.contract.FilterLogs(opts, "ExecuteBurnToken")
	if err != nil {
		return nil, err
	}
	return &WoWTokenExecuteBurnTokenIterator{contract: _WoWToken.contract, event: "ExecuteBurnToken", logs: logs, sub: sub}, nil
}

// WatchExecuteBurnToken is a free log subscription operation binding the contract event 0x30b07a2cb65e926e0ff9d6b29c7575d1da894476921a1d08596337df827e5fec.
//
// Solidity: event ExecuteBurnToken(string id_event, uint256 amount, bool isSucceeded, uint256 createdAt)
func (_WoWToken *WoWTokenFilterer) WatchExecuteBurnToken(opts *bind.WatchOpts, sink chan<- *WoWTokenExecuteBurnToken) (event.Subscription, error) {

	logs, sub, err := _WoWToken.contract.WatchLogs(opts, "ExecuteBurnToken")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WoWTokenExecuteBurnToken)
				if err := _WoWToken.contract.UnpackLog(event, "ExecuteBurnToken", log); err != nil {
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

// ParseExecuteBurnToken is a log parse operation binding the contract event 0x30b07a2cb65e926e0ff9d6b29c7575d1da894476921a1d08596337df827e5fec.
//
// Solidity: event ExecuteBurnToken(string id_event, uint256 amount, bool isSucceeded, uint256 createdAt)
func (_WoWToken *WoWTokenFilterer) ParseExecuteBurnToken(log types.Log) (*WoWTokenExecuteBurnToken, error) {
	event := new(WoWTokenExecuteBurnToken)
	if err := _WoWToken.contract.UnpackLog(event, "ExecuteBurnToken", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// WoWTokenExecuteMintTokenIterator is returned from FilterExecuteMintToken and is used to iterate over the raw logs and unpacked data for ExecuteMintToken events raised by the WoWToken contract.
type WoWTokenExecuteMintTokenIterator struct {
	Event *WoWTokenExecuteMintToken // Event containing the contract specifics and raw log

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
func (it *WoWTokenExecuteMintTokenIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WoWTokenExecuteMintToken)
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
		it.Event = new(WoWTokenExecuteMintToken)
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
func (it *WoWTokenExecuteMintTokenIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WoWTokenExecuteMintTokenIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WoWTokenExecuteMintToken represents a ExecuteMintToken event raised by the WoWToken contract.
type WoWTokenExecuteMintToken struct {
	IdEvent     string
	Amount      *big.Int
	IsSucceeded bool
	CreatedAt   *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterExecuteMintToken is a free log retrieval operation binding the contract event 0x3a666ebe4e48ccfbf0e942a4c528cf7b700c1e22d53cbeb9f05a542b4d6025e0.
//
// Solidity: event ExecuteMintToken(string id_event, uint256 amount, bool isSucceeded, uint256 createdAt)
func (_WoWToken *WoWTokenFilterer) FilterExecuteMintToken(opts *bind.FilterOpts) (*WoWTokenExecuteMintTokenIterator, error) {

	logs, sub, err := _WoWToken.contract.FilterLogs(opts, "ExecuteMintToken")
	if err != nil {
		return nil, err
	}
	return &WoWTokenExecuteMintTokenIterator{contract: _WoWToken.contract, event: "ExecuteMintToken", logs: logs, sub: sub}, nil
}

// WatchExecuteMintToken is a free log subscription operation binding the contract event 0x3a666ebe4e48ccfbf0e942a4c528cf7b700c1e22d53cbeb9f05a542b4d6025e0.
//
// Solidity: event ExecuteMintToken(string id_event, uint256 amount, bool isSucceeded, uint256 createdAt)
func (_WoWToken *WoWTokenFilterer) WatchExecuteMintToken(opts *bind.WatchOpts, sink chan<- *WoWTokenExecuteMintToken) (event.Subscription, error) {

	logs, sub, err := _WoWToken.contract.WatchLogs(opts, "ExecuteMintToken")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WoWTokenExecuteMintToken)
				if err := _WoWToken.contract.UnpackLog(event, "ExecuteMintToken", log); err != nil {
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

// ParseExecuteMintToken is a log parse operation binding the contract event 0x3a666ebe4e48ccfbf0e942a4c528cf7b700c1e22d53cbeb9f05a542b4d6025e0.
//
// Solidity: event ExecuteMintToken(string id_event, uint256 amount, bool isSucceeded, uint256 createdAt)
func (_WoWToken *WoWTokenFilterer) ParseExecuteMintToken(log types.Log) (*WoWTokenExecuteMintToken, error) {
	event := new(WoWTokenExecuteMintToken)
	if err := _WoWToken.contract.UnpackLog(event, "ExecuteMintToken", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// WoWTokenGenerateAddCoOwnerIterator is returned from FilterGenerateAddCoOwner and is used to iterate over the raw logs and unpacked data for GenerateAddCoOwner events raised by the WoWToken contract.
type WoWTokenGenerateAddCoOwnerIterator struct {
	Event *WoWTokenGenerateAddCoOwner // Event containing the contract specifics and raw log

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
func (it *WoWTokenGenerateAddCoOwnerIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WoWTokenGenerateAddCoOwner)
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
		it.Event = new(WoWTokenGenerateAddCoOwner)
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
func (it *WoWTokenGenerateAddCoOwnerIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WoWTokenGenerateAddCoOwnerIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WoWTokenGenerateAddCoOwner represents a GenerateAddCoOwner event raised by the WoWToken contract.
type WoWTokenGenerateAddCoOwner struct {
	Maker      common.Address
	NewCoOwner common.Address
	IdEvent    string
	CreatedAt  *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterGenerateAddCoOwner is a free log retrieval operation binding the contract event 0xdc19a66d04a37cdffe1c0fa71cdfc2c874536beab83e57b1d55ab8e631e55e86.
//
// Solidity: event GenerateAddCoOwner(address indexed maker, address indexed new_co_owner, string id_event, uint256 createdAt)
func (_WoWToken *WoWTokenFilterer) FilterGenerateAddCoOwner(opts *bind.FilterOpts, maker []common.Address, new_co_owner []common.Address) (*WoWTokenGenerateAddCoOwnerIterator, error) {

	var makerRule []interface{}
	for _, makerItem := range maker {
		makerRule = append(makerRule, makerItem)
	}
	var new_co_ownerRule []interface{}
	for _, new_co_ownerItem := range new_co_owner {
		new_co_ownerRule = append(new_co_ownerRule, new_co_ownerItem)
	}

	logs, sub, err := _WoWToken.contract.FilterLogs(opts, "GenerateAddCoOwner", makerRule, new_co_ownerRule)
	if err != nil {
		return nil, err
	}
	return &WoWTokenGenerateAddCoOwnerIterator{contract: _WoWToken.contract, event: "GenerateAddCoOwner", logs: logs, sub: sub}, nil
}

// WatchGenerateAddCoOwner is a free log subscription operation binding the contract event 0xdc19a66d04a37cdffe1c0fa71cdfc2c874536beab83e57b1d55ab8e631e55e86.
//
// Solidity: event GenerateAddCoOwner(address indexed maker, address indexed new_co_owner, string id_event, uint256 createdAt)
func (_WoWToken *WoWTokenFilterer) WatchGenerateAddCoOwner(opts *bind.WatchOpts, sink chan<- *WoWTokenGenerateAddCoOwner, maker []common.Address, new_co_owner []common.Address) (event.Subscription, error) {

	var makerRule []interface{}
	for _, makerItem := range maker {
		makerRule = append(makerRule, makerItem)
	}
	var new_co_ownerRule []interface{}
	for _, new_co_ownerItem := range new_co_owner {
		new_co_ownerRule = append(new_co_ownerRule, new_co_ownerItem)
	}

	logs, sub, err := _WoWToken.contract.WatchLogs(opts, "GenerateAddCoOwner", makerRule, new_co_ownerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WoWTokenGenerateAddCoOwner)
				if err := _WoWToken.contract.UnpackLog(event, "GenerateAddCoOwner", log); err != nil {
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

// ParseGenerateAddCoOwner is a log parse operation binding the contract event 0xdc19a66d04a37cdffe1c0fa71cdfc2c874536beab83e57b1d55ab8e631e55e86.
//
// Solidity: event GenerateAddCoOwner(address indexed maker, address indexed new_co_owner, string id_event, uint256 createdAt)
func (_WoWToken *WoWTokenFilterer) ParseGenerateAddCoOwner(log types.Log) (*WoWTokenGenerateAddCoOwner, error) {
	event := new(WoWTokenGenerateAddCoOwner)
	if err := _WoWToken.contract.UnpackLog(event, "GenerateAddCoOwner", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// WoWTokenGenerateBurnTokenIterator is returned from FilterGenerateBurnToken and is used to iterate over the raw logs and unpacked data for GenerateBurnToken events raised by the WoWToken contract.
type WoWTokenGenerateBurnTokenIterator struct {
	Event *WoWTokenGenerateBurnToken // Event containing the contract specifics and raw log

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
func (it *WoWTokenGenerateBurnTokenIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WoWTokenGenerateBurnToken)
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
		it.Event = new(WoWTokenGenerateBurnToken)
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
func (it *WoWTokenGenerateBurnTokenIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WoWTokenGenerateBurnTokenIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WoWTokenGenerateBurnToken represents a GenerateBurnToken event raised by the WoWToken contract.
type WoWTokenGenerateBurnToken struct {
	Maker     common.Address
	Amount    *big.Int
	IdEvent   string
	CreatedAt *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterGenerateBurnToken is a free log retrieval operation binding the contract event 0x63bb70f731ecb47b463abd26aabc2a7aefa07ee16848948071e295081f6bfa09.
//
// Solidity: event GenerateBurnToken(address indexed maker, uint256 amount, string id_event, uint256 createdAt)
func (_WoWToken *WoWTokenFilterer) FilterGenerateBurnToken(opts *bind.FilterOpts, maker []common.Address) (*WoWTokenGenerateBurnTokenIterator, error) {

	var makerRule []interface{}
	for _, makerItem := range maker {
		makerRule = append(makerRule, makerItem)
	}

	logs, sub, err := _WoWToken.contract.FilterLogs(opts, "GenerateBurnToken", makerRule)
	if err != nil {
		return nil, err
	}
	return &WoWTokenGenerateBurnTokenIterator{contract: _WoWToken.contract, event: "GenerateBurnToken", logs: logs, sub: sub}, nil
}

// WatchGenerateBurnToken is a free log subscription operation binding the contract event 0x63bb70f731ecb47b463abd26aabc2a7aefa07ee16848948071e295081f6bfa09.
//
// Solidity: event GenerateBurnToken(address indexed maker, uint256 amount, string id_event, uint256 createdAt)
func (_WoWToken *WoWTokenFilterer) WatchGenerateBurnToken(opts *bind.WatchOpts, sink chan<- *WoWTokenGenerateBurnToken, maker []common.Address) (event.Subscription, error) {

	var makerRule []interface{}
	for _, makerItem := range maker {
		makerRule = append(makerRule, makerItem)
	}

	logs, sub, err := _WoWToken.contract.WatchLogs(opts, "GenerateBurnToken", makerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WoWTokenGenerateBurnToken)
				if err := _WoWToken.contract.UnpackLog(event, "GenerateBurnToken", log); err != nil {
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

// ParseGenerateBurnToken is a log parse operation binding the contract event 0x63bb70f731ecb47b463abd26aabc2a7aefa07ee16848948071e295081f6bfa09.
//
// Solidity: event GenerateBurnToken(address indexed maker, uint256 amount, string id_event, uint256 createdAt)
func (_WoWToken *WoWTokenFilterer) ParseGenerateBurnToken(log types.Log) (*WoWTokenGenerateBurnToken, error) {
	event := new(WoWTokenGenerateBurnToken)
	if err := _WoWToken.contract.UnpackLog(event, "GenerateBurnToken", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// WoWTokenGenerateMintTokenIterator is returned from FilterGenerateMintToken and is used to iterate over the raw logs and unpacked data for GenerateMintToken events raised by the WoWToken contract.
type WoWTokenGenerateMintTokenIterator struct {
	Event *WoWTokenGenerateMintToken // Event containing the contract specifics and raw log

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
func (it *WoWTokenGenerateMintTokenIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WoWTokenGenerateMintToken)
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
		it.Event = new(WoWTokenGenerateMintToken)
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
func (it *WoWTokenGenerateMintTokenIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WoWTokenGenerateMintTokenIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WoWTokenGenerateMintToken represents a GenerateMintToken event raised by the WoWToken contract.
type WoWTokenGenerateMintToken struct {
	Maker     common.Address
	Amount    *big.Int
	IdEvent   string
	CreatedAt *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterGenerateMintToken is a free log retrieval operation binding the contract event 0x025fabde05ada16c0121e69fbad092a7376f87135074bebc04001c1aace26090.
//
// Solidity: event GenerateMintToken(address indexed maker, uint256 amount, string id_event, uint256 createdAt)
func (_WoWToken *WoWTokenFilterer) FilterGenerateMintToken(opts *bind.FilterOpts, maker []common.Address) (*WoWTokenGenerateMintTokenIterator, error) {

	var makerRule []interface{}
	for _, makerItem := range maker {
		makerRule = append(makerRule, makerItem)
	}

	logs, sub, err := _WoWToken.contract.FilterLogs(opts, "GenerateMintToken", makerRule)
	if err != nil {
		return nil, err
	}
	return &WoWTokenGenerateMintTokenIterator{contract: _WoWToken.contract, event: "GenerateMintToken", logs: logs, sub: sub}, nil
}

// WatchGenerateMintToken is a free log subscription operation binding the contract event 0x025fabde05ada16c0121e69fbad092a7376f87135074bebc04001c1aace26090.
//
// Solidity: event GenerateMintToken(address indexed maker, uint256 amount, string id_event, uint256 createdAt)
func (_WoWToken *WoWTokenFilterer) WatchGenerateMintToken(opts *bind.WatchOpts, sink chan<- *WoWTokenGenerateMintToken, maker []common.Address) (event.Subscription, error) {

	var makerRule []interface{}
	for _, makerItem := range maker {
		makerRule = append(makerRule, makerItem)
	}

	logs, sub, err := _WoWToken.contract.WatchLogs(opts, "GenerateMintToken", makerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WoWTokenGenerateMintToken)
				if err := _WoWToken.contract.UnpackLog(event, "GenerateMintToken", log); err != nil {
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

// ParseGenerateMintToken is a log parse operation binding the contract event 0x025fabde05ada16c0121e69fbad092a7376f87135074bebc04001c1aace26090.
//
// Solidity: event GenerateMintToken(address indexed maker, uint256 amount, string id_event, uint256 createdAt)
func (_WoWToken *WoWTokenFilterer) ParseGenerateMintToken(log types.Log) (*WoWTokenGenerateMintToken, error) {
	event := new(WoWTokenGenerateMintToken)
	if err := _WoWToken.contract.UnpackLog(event, "GenerateMintToken", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// WoWTokenPauseIterator is returned from FilterPause and is used to iterate over the raw logs and unpacked data for Pause events raised by the WoWToken contract.
type WoWTokenPauseIterator struct {
	Event *WoWTokenPause // Event containing the contract specifics and raw log

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
func (it *WoWTokenPauseIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WoWTokenPause)
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
		it.Event = new(WoWTokenPause)
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
func (it *WoWTokenPauseIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WoWTokenPauseIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WoWTokenPause represents a Pause event raised by the WoWToken contract.
type WoWTokenPause struct {
	CreatedAt *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterPause is a free log retrieval operation binding the contract event 0x68b095021b1f40fe513109f513c66692f0b3219aee674a69f4efc57badb8201d.
//
// Solidity: event Pause(uint256 createdAt)
func (_WoWToken *WoWTokenFilterer) FilterPause(opts *bind.FilterOpts) (*WoWTokenPauseIterator, error) {

	logs, sub, err := _WoWToken.contract.FilterLogs(opts, "Pause")
	if err != nil {
		return nil, err
	}
	return &WoWTokenPauseIterator{contract: _WoWToken.contract, event: "Pause", logs: logs, sub: sub}, nil
}

// WatchPause is a free log subscription operation binding the contract event 0x68b095021b1f40fe513109f513c66692f0b3219aee674a69f4efc57badb8201d.
//
// Solidity: event Pause(uint256 createdAt)
func (_WoWToken *WoWTokenFilterer) WatchPause(opts *bind.WatchOpts, sink chan<- *WoWTokenPause) (event.Subscription, error) {

	logs, sub, err := _WoWToken.contract.WatchLogs(opts, "Pause")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WoWTokenPause)
				if err := _WoWToken.contract.UnpackLog(event, "Pause", log); err != nil {
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

// ParsePause is a log parse operation binding the contract event 0x68b095021b1f40fe513109f513c66692f0b3219aee674a69f4efc57badb8201d.
//
// Solidity: event Pause(uint256 createdAt)
func (_WoWToken *WoWTokenFilterer) ParsePause(log types.Log) (*WoWTokenPause, error) {
	event := new(WoWTokenPause)
	if err := _WoWToken.contract.UnpackLog(event, "Pause", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// WoWTokenSignEventIterator is returned from FilterSignEvent and is used to iterate over the raw logs and unpacked data for SignEvent events raised by the WoWToken contract.
type WoWTokenSignEventIterator struct {
	Event *WoWTokenSignEvent // Event containing the contract specifics and raw log

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
func (it *WoWTokenSignEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WoWTokenSignEvent)
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
		it.Event = new(WoWTokenSignEvent)
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
func (it *WoWTokenSignEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WoWTokenSignEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WoWTokenSignEvent represents a SignEvent event raised by the WoWToken contract.
type WoWTokenSignEvent struct {
	Signer    common.Address
	IdEvent   string
	Approved  bool
	CreatedAt *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterSignEvent is a free log retrieval operation binding the contract event 0x6817d5b86863f7cf8eee0dbc2ea7ebfc06c03e30011ce957cfcc9e5a40ad358e.
//
// Solidity: event SignEvent(address indexed signer, string id_event, bool approved, uint256 createdAt)
func (_WoWToken *WoWTokenFilterer) FilterSignEvent(opts *bind.FilterOpts, signer []common.Address) (*WoWTokenSignEventIterator, error) {

	var signerRule []interface{}
	for _, signerItem := range signer {
		signerRule = append(signerRule, signerItem)
	}

	logs, sub, err := _WoWToken.contract.FilterLogs(opts, "SignEvent", signerRule)
	if err != nil {
		return nil, err
	}
	return &WoWTokenSignEventIterator{contract: _WoWToken.contract, event: "SignEvent", logs: logs, sub: sub}, nil
}

// WatchSignEvent is a free log subscription operation binding the contract event 0x6817d5b86863f7cf8eee0dbc2ea7ebfc06c03e30011ce957cfcc9e5a40ad358e.
//
// Solidity: event SignEvent(address indexed signer, string id_event, bool approved, uint256 createdAt)
func (_WoWToken *WoWTokenFilterer) WatchSignEvent(opts *bind.WatchOpts, sink chan<- *WoWTokenSignEvent, signer []common.Address) (event.Subscription, error) {

	var signerRule []interface{}
	for _, signerItem := range signer {
		signerRule = append(signerRule, signerItem)
	}

	logs, sub, err := _WoWToken.contract.WatchLogs(opts, "SignEvent", signerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WoWTokenSignEvent)
				if err := _WoWToken.contract.UnpackLog(event, "SignEvent", log); err != nil {
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

// ParseSignEvent is a log parse operation binding the contract event 0x6817d5b86863f7cf8eee0dbc2ea7ebfc06c03e30011ce957cfcc9e5a40ad358e.
//
// Solidity: event SignEvent(address indexed signer, string id_event, bool approved, uint256 createdAt)
func (_WoWToken *WoWTokenFilterer) ParseSignEvent(log types.Log) (*WoWTokenSignEvent, error) {
	event := new(WoWTokenSignEvent)
	if err := _WoWToken.contract.UnpackLog(event, "SignEvent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// WoWTokenTransferTokenIterator is returned from FilterTransferToken and is used to iterate over the raw logs and unpacked data for TransferToken events raised by the WoWToken contract.
type WoWTokenTransferTokenIterator struct {
	Event *WoWTokenTransferToken // Event containing the contract specifics and raw log

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
func (it *WoWTokenTransferTokenIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WoWTokenTransferToken)
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
		it.Event = new(WoWTokenTransferToken)
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
func (it *WoWTokenTransferTokenIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WoWTokenTransferTokenIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WoWTokenTransferToken represents a TransferToken event raised by the WoWToken contract.
type WoWTokenTransferToken struct {
	From      common.Address
	To        common.Address
	Amount    *big.Int
	CreatedAt *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterTransferToken is a free log retrieval operation binding the contract event 0x59158d23d62750f16cd67dd3dce5fcfeb193d88db7c46f085a2eac8fa5d8b950.
//
// Solidity: event TransferToken(address indexed from, address indexed to, uint256 amount, uint256 createdAt)
func (_WoWToken *WoWTokenFilterer) FilterTransferToken(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*WoWTokenTransferTokenIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _WoWToken.contract.FilterLogs(opts, "TransferToken", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &WoWTokenTransferTokenIterator{contract: _WoWToken.contract, event: "TransferToken", logs: logs, sub: sub}, nil
}

// WatchTransferToken is a free log subscription operation binding the contract event 0x59158d23d62750f16cd67dd3dce5fcfeb193d88db7c46f085a2eac8fa5d8b950.
//
// Solidity: event TransferToken(address indexed from, address indexed to, uint256 amount, uint256 createdAt)
func (_WoWToken *WoWTokenFilterer) WatchTransferToken(opts *bind.WatchOpts, sink chan<- *WoWTokenTransferToken, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _WoWToken.contract.WatchLogs(opts, "TransferToken", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WoWTokenTransferToken)
				if err := _WoWToken.contract.UnpackLog(event, "TransferToken", log); err != nil {
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

// ParseTransferToken is a log parse operation binding the contract event 0x59158d23d62750f16cd67dd3dce5fcfeb193d88db7c46f085a2eac8fa5d8b950.
//
// Solidity: event TransferToken(address indexed from, address indexed to, uint256 amount, uint256 createdAt)
func (_WoWToken *WoWTokenFilterer) ParseTransferToken(log types.Log) (*WoWTokenTransferToken, error) {
	event := new(WoWTokenTransferToken)
	if err := _WoWToken.contract.UnpackLog(event, "TransferToken", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// WoWTokenUnpauseIterator is returned from FilterUnpause and is used to iterate over the raw logs and unpacked data for Unpause events raised by the WoWToken contract.
type WoWTokenUnpauseIterator struct {
	Event *WoWTokenUnpause // Event containing the contract specifics and raw log

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
func (it *WoWTokenUnpauseIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WoWTokenUnpause)
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
		it.Event = new(WoWTokenUnpause)
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
func (it *WoWTokenUnpauseIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WoWTokenUnpauseIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WoWTokenUnpause represents a Unpause event raised by the WoWToken contract.
type WoWTokenUnpause struct {
	CreatedAt *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterUnpause is a free log retrieval operation binding the contract event 0xaaa520fdd7d2c83061d632fa017b0432407e798818af63ea908589fceda39ab7.
//
// Solidity: event Unpause(uint256 createdAt)
func (_WoWToken *WoWTokenFilterer) FilterUnpause(opts *bind.FilterOpts) (*WoWTokenUnpauseIterator, error) {

	logs, sub, err := _WoWToken.contract.FilterLogs(opts, "Unpause")
	if err != nil {
		return nil, err
	}
	return &WoWTokenUnpauseIterator{contract: _WoWToken.contract, event: "Unpause", logs: logs, sub: sub}, nil
}

// WatchUnpause is a free log subscription operation binding the contract event 0xaaa520fdd7d2c83061d632fa017b0432407e798818af63ea908589fceda39ab7.
//
// Solidity: event Unpause(uint256 createdAt)
func (_WoWToken *WoWTokenFilterer) WatchUnpause(opts *bind.WatchOpts, sink chan<- *WoWTokenUnpause) (event.Subscription, error) {

	logs, sub, err := _WoWToken.contract.WatchLogs(opts, "Unpause")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WoWTokenUnpause)
				if err := _WoWToken.contract.UnpackLog(event, "Unpause", log); err != nil {
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

// ParseUnpause is a log parse operation binding the contract event 0xaaa520fdd7d2c83061d632fa017b0432407e798818af63ea908589fceda39ab7.
//
// Solidity: event Unpause(uint256 createdAt)
func (_WoWToken *WoWTokenFilterer) ParseUnpause(log types.Log) (*WoWTokenUnpause, error) {
	event := new(WoWTokenUnpause)
	if err := _WoWToken.contract.UnpackLog(event, "Unpause", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
