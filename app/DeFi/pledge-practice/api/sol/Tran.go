// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package sol

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

// SolMetaData contains all meta data concerning the Sol contract.
var SolMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"EtherReceived\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"FallbackCalled\",\"type\":\"event\"},{\"stateMutability\":\"payable\",\"type\":\"fallback\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"senders\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
	Bin: "0x6080604052348015600e575f5ffd5b50335f5f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055506103e58061005b5f395ff3fe60806040526004361061002c575f3560e01c80638da5cb5b146101495780639977c78a14610173576100f9565b366100f9573460015f8282546100429190610244565b92505081905550600233908060018154018082558091505060019003905f5260205f20015f9091909190916101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055503373ffffffffffffffffffffffffffffffffffffffff167f1e57e3bb474320be3d2c77138f75b7c3941292d647f5f9634e33a8e94e0e069b346040516100ef9190610286565b60405180910390a2005b3373ffffffffffffffffffffffffffffffffffffffff167faca09dd456ca888dccf8cc966e382e6e3042bb7e4d2d7815015f844edeafce423460405161013f91906102d2565b60405180910390a2005b348015610154575f5ffd5b5061015d6101af565b60405161016a919061033d565b60405180910390f35b34801561017e575f5ffd5b5061019960048036038101906101949190610384565b6101d3565b6040516101a6919061033d565b60405180910390f35b5f5f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b600281815481106101e2575f80fd5b905f5260205f20015f915054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b5f819050919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52601160045260245ffd5b5f61024e8261020e565b91506102598361020e565b925082820190508082111561027157610270610217565b5b92915050565b6102808161020e565b82525050565b5f6020820190506102995f830184610277565b92915050565b5f82825260208201905092915050565b50565b5f6102bd5f8361029f565b91506102c8826102af565b5f82019050919050565b5f6040820190506102e55f830184610277565b81810360208301526102f6816102b2565b905092915050565b5f73ffffffffffffffffffffffffffffffffffffffff82169050919050565b5f610327826102fe565b9050919050565b6103378161031d565b82525050565b5f6020820190506103505f83018461032e565b92915050565b5f5ffd5b6103638161020e565b811461036d575f5ffd5b50565b5f8135905061037e8161035a565b92915050565b5f6020828403121561039957610398610356565b5b5f6103a684828501610370565b9150509291505056fea2646970667358221220ec57b4916b2bbd12e1dddfa692c1f32a60957e6e874017bfa8b8b8d396fb97de64736f6c634300081d0033",
}

// SolABI is the input ABI used to generate the binding from.
// Deprecated: Use SolMetaData.ABI instead.
var SolABI = SolMetaData.ABI

// SolBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use SolMetaData.Bin instead.
var SolBin = SolMetaData.Bin

// DeploySol deploys a new Ethereum contract, binding an instance of Sol to it.
func DeploySol(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Sol, error) {
	parsed, err := SolMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(SolBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Sol{SolCaller: SolCaller{contract: contract}, SolTransactor: SolTransactor{contract: contract}, SolFilterer: SolFilterer{contract: contract}}, nil
}

// Sol is an auto generated Go binding around an Ethereum contract.
type Sol struct {
	SolCaller     // Read-only binding to the contract
	SolTransactor // Write-only binding to the contract
	SolFilterer   // Log filterer for contract events
}

// SolCaller is an auto generated read-only Go binding around an Ethereum contract.
type SolCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SolTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SolTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SolFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SolFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SolSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SolSession struct {
	Contract     *Sol              // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SolCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SolCallerSession struct {
	Contract *SolCaller    // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// SolTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SolTransactorSession struct {
	Contract     *SolTransactor    // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SolRaw is an auto generated low-level Go binding around an Ethereum contract.
type SolRaw struct {
	Contract *Sol // Generic contract binding to access the raw methods on
}

// SolCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SolCallerRaw struct {
	Contract *SolCaller // Generic read-only contract binding to access the raw methods on
}

// SolTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SolTransactorRaw struct {
	Contract *SolTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSol creates a new instance of Sol, bound to a specific deployed contract.
func NewSol(address common.Address, backend bind.ContractBackend) (*Sol, error) {
	contract, err := bindSol(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Sol{SolCaller: SolCaller{contract: contract}, SolTransactor: SolTransactor{contract: contract}, SolFilterer: SolFilterer{contract: contract}}, nil
}

// NewSolCaller creates a new read-only instance of Sol, bound to a specific deployed contract.
func NewSolCaller(address common.Address, caller bind.ContractCaller) (*SolCaller, error) {
	contract, err := bindSol(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SolCaller{contract: contract}, nil
}

// NewSolTransactor creates a new write-only instance of Sol, bound to a specific deployed contract.
func NewSolTransactor(address common.Address, transactor bind.ContractTransactor) (*SolTransactor, error) {
	contract, err := bindSol(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SolTransactor{contract: contract}, nil
}

// NewSolFilterer creates a new log filterer instance of Sol, bound to a specific deployed contract.
func NewSolFilterer(address common.Address, filterer bind.ContractFilterer) (*SolFilterer, error) {
	contract, err := bindSol(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SolFilterer{contract: contract}, nil
}

// bindSol binds a generic wrapper to an already deployed contract.
func bindSol(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := SolMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Sol *SolRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Sol.Contract.SolCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Sol *SolRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Sol.Contract.SolTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Sol *SolRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Sol.Contract.SolTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Sol *SolCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Sol.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Sol *SolTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Sol.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Sol *SolTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Sol.Contract.contract.Transact(opts, method, params...)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Sol *SolCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Sol.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Sol *SolSession) Owner() (common.Address, error) {
	return _Sol.Contract.Owner(&_Sol.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Sol *SolCallerSession) Owner() (common.Address, error) {
	return _Sol.Contract.Owner(&_Sol.CallOpts)
}

// Senders is a free data retrieval call binding the contract method 0x9977c78a.
//
// Solidity: function senders(uint256 ) view returns(address)
func (_Sol *SolCaller) Senders(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Sol.contract.Call(opts, &out, "senders", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Senders is a free data retrieval call binding the contract method 0x9977c78a.
//
// Solidity: function senders(uint256 ) view returns(address)
func (_Sol *SolSession) Senders(arg0 *big.Int) (common.Address, error) {
	return _Sol.Contract.Senders(&_Sol.CallOpts, arg0)
}

// Senders is a free data retrieval call binding the contract method 0x9977c78a.
//
// Solidity: function senders(uint256 ) view returns(address)
func (_Sol *SolCallerSession) Senders(arg0 *big.Int) (common.Address, error) {
	return _Sol.Contract.Senders(&_Sol.CallOpts, arg0)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_Sol *SolTransactor) Fallback(opts *bind.TransactOpts, calldata []byte) (*types.Transaction, error) {
	return _Sol.contract.RawTransact(opts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_Sol *SolSession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _Sol.Contract.Fallback(&_Sol.TransactOpts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_Sol *SolTransactorSession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _Sol.Contract.Fallback(&_Sol.TransactOpts, calldata)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Sol *SolTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Sol.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Sol *SolSession) Receive() (*types.Transaction, error) {
	return _Sol.Contract.Receive(&_Sol.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Sol *SolTransactorSession) Receive() (*types.Transaction, error) {
	return _Sol.Contract.Receive(&_Sol.TransactOpts)
}

// SolEtherReceivedIterator is returned from FilterEtherReceived and is used to iterate over the raw logs and unpacked data for EtherReceived events raised by the Sol contract.
type SolEtherReceivedIterator struct {
	Event *SolEtherReceived // Event containing the contract specifics and raw log

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
func (it *SolEtherReceivedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SolEtherReceived)
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
		it.Event = new(SolEtherReceived)
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
func (it *SolEtherReceivedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SolEtherReceivedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SolEtherReceived represents a EtherReceived event raised by the Sol contract.
type SolEtherReceived struct {
	Sender common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterEtherReceived is a free log retrieval operation binding the contract event 0x1e57e3bb474320be3d2c77138f75b7c3941292d647f5f9634e33a8e94e0e069b.
//
// Solidity: event EtherReceived(address indexed sender, uint256 amount)
func (_Sol *SolFilterer) FilterEtherReceived(opts *bind.FilterOpts, sender []common.Address) (*SolEtherReceivedIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _Sol.contract.FilterLogs(opts, "EtherReceived", senderRule)
	if err != nil {
		return nil, err
	}
	return &SolEtherReceivedIterator{contract: _Sol.contract, event: "EtherReceived", logs: logs, sub: sub}, nil
}

// WatchEtherReceived is a free log subscription operation binding the contract event 0x1e57e3bb474320be3d2c77138f75b7c3941292d647f5f9634e33a8e94e0e069b.
//
// Solidity: event EtherReceived(address indexed sender, uint256 amount)
func (_Sol *SolFilterer) WatchEtherReceived(opts *bind.WatchOpts, sink chan<- *SolEtherReceived, sender []common.Address) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _Sol.contract.WatchLogs(opts, "EtherReceived", senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SolEtherReceived)
				if err := _Sol.contract.UnpackLog(event, "EtherReceived", log); err != nil {
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

// ParseEtherReceived is a log parse operation binding the contract event 0x1e57e3bb474320be3d2c77138f75b7c3941292d647f5f9634e33a8e94e0e069b.
//
// Solidity: event EtherReceived(address indexed sender, uint256 amount)
func (_Sol *SolFilterer) ParseEtherReceived(log types.Log) (*SolEtherReceived, error) {
	event := new(SolEtherReceived)
	if err := _Sol.contract.UnpackLog(event, "EtherReceived", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SolFallbackCalledIterator is returned from FilterFallbackCalled and is used to iterate over the raw logs and unpacked data for FallbackCalled events raised by the Sol contract.
type SolFallbackCalledIterator struct {
	Event *SolFallbackCalled // Event containing the contract specifics and raw log

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
func (it *SolFallbackCalledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SolFallbackCalled)
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
		it.Event = new(SolFallbackCalled)
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
func (it *SolFallbackCalledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SolFallbackCalledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SolFallbackCalled represents a FallbackCalled event raised by the Sol contract.
type SolFallbackCalled struct {
	Sender common.Address
	Amount *big.Int
	Data   []byte
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterFallbackCalled is a free log retrieval operation binding the contract event 0xaca09dd456ca888dccf8cc966e382e6e3042bb7e4d2d7815015f844edeafce42.
//
// Solidity: event FallbackCalled(address indexed sender, uint256 amount, bytes data)
func (_Sol *SolFilterer) FilterFallbackCalled(opts *bind.FilterOpts, sender []common.Address) (*SolFallbackCalledIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _Sol.contract.FilterLogs(opts, "FallbackCalled", senderRule)
	if err != nil {
		return nil, err
	}
	return &SolFallbackCalledIterator{contract: _Sol.contract, event: "FallbackCalled", logs: logs, sub: sub}, nil
}

// WatchFallbackCalled is a free log subscription operation binding the contract event 0xaca09dd456ca888dccf8cc966e382e6e3042bb7e4d2d7815015f844edeafce42.
//
// Solidity: event FallbackCalled(address indexed sender, uint256 amount, bytes data)
func (_Sol *SolFilterer) WatchFallbackCalled(opts *bind.WatchOpts, sink chan<- *SolFallbackCalled, sender []common.Address) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _Sol.contract.WatchLogs(opts, "FallbackCalled", senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SolFallbackCalled)
				if err := _Sol.contract.UnpackLog(event, "FallbackCalled", log); err != nil {
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

// ParseFallbackCalled is a log parse operation binding the contract event 0xaca09dd456ca888dccf8cc966e382e6e3042bb7e4d2d7815015f844edeafce42.
//
// Solidity: event FallbackCalled(address indexed sender, uint256 amount, bytes data)
func (_Sol *SolFilterer) ParseFallbackCalled(log types.Log) (*SolFallbackCalled, error) {
	event := new(SolFallbackCalled)
	if err := _Sol.contract.UnpackLog(event, "FallbackCalled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
