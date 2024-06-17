// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package secureDataTransfer

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

// SecureDataTransferMetaData contains all meta data concerning the SecureDataTransfer contract.
var SecureDataTransferMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timeCreated\",\"type\":\"uint256\"}],\"name\":\"DataSent\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"getData\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"encryptedData\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"dataHash\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"signedHash\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"timeCreated\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_receiver\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"_encryptedData\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_dataHash\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_signedHash\",\"type\":\"string\"}],\"name\":\"sendData\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b5033600460006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055506000600560146101000a81548160ff021916908315150217905550610d7b8061007c6000396000f3fe608060405234801561001057600080fd5b50600436106100365760003560e01c80633bc5de301461003b578063658055e01461005c575b600080fd5b610043610078565b6040516100539493929190610600565b60405180910390f35b61007660048036038101906100719190610801565b61035a565b005b60608060606000600460009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614806101285750600560009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16145b610167576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161015e9061092e565b60405180910390fd5b6000806040518060800160405290816000820180546101859061097d565b80601f01602080910402602001604051908101604052809291908181526020018280546101b19061097d565b80156101fe5780601f106101d3576101008083540402835291602001916101fe565b820191906000526020600020905b8154815290600101906020018083116101e157829003601f168201915b505050505081526020016001820180546102179061097d565b80601f01602080910402602001604051908101604052809291908181526020018280546102439061097d565b80156102905780601f1061026557610100808354040283529160200191610290565b820191906000526020600020905b81548152906001019060200180831161027357829003601f168201915b505050505081526020016002820180546102a99061097d565b80601f01602080910402602001604051908101604052809291908181526020018280546102d59061097d565b80156103225780601f106102f757610100808354040283529160200191610322565b820191906000526020600020905b81548152906001019060200180831161030557829003601f168201915b505050505081526020016003820154815250509050806000015181602001518260400151836060015194509450945094505090919293565b600560149054906101000a900460ff16156103aa576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016103a190610a20565b60405180910390fd5b600073ffffffffffffffffffffffffffffffffffffffff168473ffffffffffffffffffffffffffffffffffffffff1603610419576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161041090610a8c565b60405180910390fd5b60006040518060800160405280858152602001848152602001838152602001428152509050806000808201518160000190816104559190610c58565b50602082015181600101908161046b9190610c58565b5060408201518160020190816104819190610c58565b506060820151816003015590505084600560006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055506001600560146101000a81548160ff0219169083151502179055508473ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff167fa66e03c9001bca5fd5a6b12c7ff0565f91adab37b440198b3d413712a3741f36426040516105489190610d2a565b60405180910390a35050505050565b600081519050919050565b600082825260208201905092915050565b60005b83811015610591578082015181840152602081019050610576565b60008484015250505050565b6000601f19601f8301169050919050565b60006105b982610557565b6105c38185610562565b93506105d3818560208601610573565b6105dc8161059d565b840191505092915050565b6000819050919050565b6105fa816105e7565b82525050565b6000608082019050818103600083015261061a81876105ae565b9050818103602083015261062e81866105ae565b9050818103604083015261064281856105ae565b905061065160608301846105f1565b95945050505050565b6000604051905090565b600080fd5b600080fd5b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b60006106998261066e565b9050919050565b6106a98161068e565b81146106b457600080fd5b50565b6000813590506106c6816106a0565b92915050565b600080fd5b600080fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b61070e8261059d565b810181811067ffffffffffffffff8211171561072d5761072c6106d6565b5b80604052505050565b600061074061065a565b905061074c8282610705565b919050565b600067ffffffffffffffff82111561076c5761076b6106d6565b5b6107758261059d565b9050602081019050919050565b82818337600083830152505050565b60006107a461079f84610751565b610736565b9050828152602081018484840111156107c0576107bf6106d1565b5b6107cb848285610782565b509392505050565b600082601f8301126107e8576107e76106cc565b5b81356107f8848260208601610791565b91505092915050565b6000806000806080858703121561081b5761081a610664565b5b6000610829878288016106b7565b945050602085013567ffffffffffffffff81111561084a57610849610669565b5b610856878288016107d3565b935050604085013567ffffffffffffffff81111561087757610876610669565b5b610883878288016107d3565b925050606085013567ffffffffffffffff8111156108a4576108a3610669565b5b6108b0878288016107d3565b91505092959194509250565b7f4f6e6c79207468652073656e646572206f722072656365697665722063616e2060008201527f616363657373207468697320646174612e000000000000000000000000000000602082015250565b6000610918603183610562565b9150610923826108bc565b604082019050919050565b600060208201905081810360008301526109478161090b565b9050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b6000600282049050600182168061099557607f821691505b6020821081036109a8576109a761094e565b5b50919050565b7f4964656e746974792068617320616c7265616479206265656e2063726561746560008201527f642e000000000000000000000000000000000000000000000000000000000000602082015250565b6000610a0a602283610562565b9150610a15826109ae565b604082019050919050565b60006020820190508181036000830152610a39816109fd565b9050919050565b7f496e76616c696420726563656976657220616464726573732e00000000000000600082015250565b6000610a76601983610562565b9150610a8182610a40565b602082019050919050565b60006020820190508181036000830152610aa581610a69565b9050919050565b60008190508160005260206000209050919050565b60006020601f8301049050919050565b600082821b905092915050565b600060088302610b0e7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82610ad1565b610b188683610ad1565b95508019841693508086168417925050509392505050565b6000819050919050565b6000610b55610b50610b4b846105e7565b610b30565b6105e7565b9050919050565b6000819050919050565b610b6f83610b3a565b610b83610b7b82610b5c565b848454610ade565b825550505050565b600090565b610b98610b8b565b610ba3818484610b66565b505050565b5b81811015610bc757610bbc600082610b90565b600181019050610ba9565b5050565b601f821115610c0c57610bdd81610aac565b610be684610ac1565b81016020851015610bf5578190505b610c09610c0185610ac1565b830182610ba8565b50505b505050565b600082821c905092915050565b6000610c2f60001984600802610c11565b1980831691505092915050565b6000610c488383610c1e565b9150826002028217905092915050565b610c6182610557565b67ffffffffffffffff811115610c7a57610c796106d6565b5b610c84825461097d565b610c8f828285610bcb565b600060209050601f831160018114610cc25760008415610cb0578287015190505b610cba8582610c3c565b865550610d22565b601f198416610cd086610aac565b60005b82811015610cf857848901518255600182019150602085019450602081019050610cd3565b86831015610d155784890151610d11601f891682610c1e565b8355505b6001600288020188555050505b505050505050565b6000602082019050610d3f60008301846105f1565b9291505056fea2646970667358221220abfdd9093ec5154f5416a6c4798f28858f903a045b097809df00e5165f2f3b8164736f6c63430008120033",
}

// SecureDataTransferABI is the input ABI used to generate the binding from.
// Deprecated: Use SecureDataTransferMetaData.ABI instead.
var SecureDataTransferABI = SecureDataTransferMetaData.ABI

// SecureDataTransferBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use SecureDataTransferMetaData.Bin instead.
var SecureDataTransferBin = SecureDataTransferMetaData.Bin

// DeploySecureDataTransfer deploys a new Ethereum contract, binding an instance of SecureDataTransfer to it.
func DeploySecureDataTransfer(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *SecureDataTransfer, error) {
	parsed, err := SecureDataTransferMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(SecureDataTransferBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &SecureDataTransfer{SecureDataTransferCaller: SecureDataTransferCaller{contract: contract}, SecureDataTransferTransactor: SecureDataTransferTransactor{contract: contract}, SecureDataTransferFilterer: SecureDataTransferFilterer{contract: contract}}, nil
}

// SecureDataTransfer is an auto generated Go binding around an Ethereum contract.
type SecureDataTransfer struct {
	SecureDataTransferCaller     // Read-only binding to the contract
	SecureDataTransferTransactor // Write-only binding to the contract
	SecureDataTransferFilterer   // Log filterer for contract events
}

// SecureDataTransferCaller is an auto generated read-only Go binding around an Ethereum contract.
type SecureDataTransferCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SecureDataTransferTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SecureDataTransferTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SecureDataTransferFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SecureDataTransferFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SecureDataTransferSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SecureDataTransferSession struct {
	Contract     *SecureDataTransfer // Generic contract binding to set the session for
	CallOpts     bind.CallOpts       // Call options to use throughout this session
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// SecureDataTransferCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SecureDataTransferCallerSession struct {
	Contract *SecureDataTransferCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts             // Call options to use throughout this session
}

// SecureDataTransferTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SecureDataTransferTransactorSession struct {
	Contract     *SecureDataTransferTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts             // Transaction auth options to use throughout this session
}

// SecureDataTransferRaw is an auto generated low-level Go binding around an Ethereum contract.
type SecureDataTransferRaw struct {
	Contract *SecureDataTransfer // Generic contract binding to access the raw methods on
}

// SecureDataTransferCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SecureDataTransferCallerRaw struct {
	Contract *SecureDataTransferCaller // Generic read-only contract binding to access the raw methods on
}

// SecureDataTransferTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SecureDataTransferTransactorRaw struct {
	Contract *SecureDataTransferTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSecureDataTransfer creates a new instance of SecureDataTransfer, bound to a specific deployed contract.
func NewSecureDataTransfer(address common.Address, backend bind.ContractBackend) (*SecureDataTransfer, error) {
	contract, err := bindSecureDataTransfer(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &SecureDataTransfer{SecureDataTransferCaller: SecureDataTransferCaller{contract: contract}, SecureDataTransferTransactor: SecureDataTransferTransactor{contract: contract}, SecureDataTransferFilterer: SecureDataTransferFilterer{contract: contract}}, nil
}

// NewSecureDataTransferCaller creates a new read-only instance of SecureDataTransfer, bound to a specific deployed contract.
func NewSecureDataTransferCaller(address common.Address, caller bind.ContractCaller) (*SecureDataTransferCaller, error) {
	contract, err := bindSecureDataTransfer(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SecureDataTransferCaller{contract: contract}, nil
}

// NewSecureDataTransferTransactor creates a new write-only instance of SecureDataTransfer, bound to a specific deployed contract.
func NewSecureDataTransferTransactor(address common.Address, transactor bind.ContractTransactor) (*SecureDataTransferTransactor, error) {
	contract, err := bindSecureDataTransfer(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SecureDataTransferTransactor{contract: contract}, nil
}

// NewSecureDataTransferFilterer creates a new log filterer instance of SecureDataTransfer, bound to a specific deployed contract.
func NewSecureDataTransferFilterer(address common.Address, filterer bind.ContractFilterer) (*SecureDataTransferFilterer, error) {
	contract, err := bindSecureDataTransfer(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SecureDataTransferFilterer{contract: contract}, nil
}

// bindSecureDataTransfer binds a generic wrapper to an already deployed contract.
func bindSecureDataTransfer(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := SecureDataTransferMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SecureDataTransfer *SecureDataTransferRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SecureDataTransfer.Contract.SecureDataTransferCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SecureDataTransfer *SecureDataTransferRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SecureDataTransfer.Contract.SecureDataTransferTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SecureDataTransfer *SecureDataTransferRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SecureDataTransfer.Contract.SecureDataTransferTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SecureDataTransfer *SecureDataTransferCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SecureDataTransfer.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SecureDataTransfer *SecureDataTransferTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SecureDataTransfer.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SecureDataTransfer *SecureDataTransferTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SecureDataTransfer.Contract.contract.Transact(opts, method, params...)
}

// GetData is a free data retrieval call binding the contract method 0x3bc5de30.
//
// Solidity: function getData() view returns(string encryptedData, string dataHash, string signedHash, uint256 timeCreated)
func (_SecureDataTransfer *SecureDataTransferCaller) GetData(opts *bind.CallOpts) (struct {
	EncryptedData string
	DataHash      string
	SignedHash    string
	TimeCreated   *big.Int
}, error) {
	var out []interface{}
	err := _SecureDataTransfer.contract.Call(opts, &out, "getData")

	outstruct := new(struct {
		EncryptedData string
		DataHash      string
		SignedHash    string
		TimeCreated   *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.EncryptedData = *abi.ConvertType(out[0], new(string)).(*string)
	outstruct.DataHash = *abi.ConvertType(out[1], new(string)).(*string)
	outstruct.SignedHash = *abi.ConvertType(out[2], new(string)).(*string)
	outstruct.TimeCreated = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// GetData is a free data retrieval call binding the contract method 0x3bc5de30.
//
// Solidity: function getData() view returns(string encryptedData, string dataHash, string signedHash, uint256 timeCreated)
func (_SecureDataTransfer *SecureDataTransferSession) GetData() (struct {
	EncryptedData string
	DataHash      string
	SignedHash    string
	TimeCreated   *big.Int
}, error) {
	return _SecureDataTransfer.Contract.GetData(&_SecureDataTransfer.CallOpts)
}

// GetData is a free data retrieval call binding the contract method 0x3bc5de30.
//
// Solidity: function getData() view returns(string encryptedData, string dataHash, string signedHash, uint256 timeCreated)
func (_SecureDataTransfer *SecureDataTransferCallerSession) GetData() (struct {
	EncryptedData string
	DataHash      string
	SignedHash    string
	TimeCreated   *big.Int
}, error) {
	return _SecureDataTransfer.Contract.GetData(&_SecureDataTransfer.CallOpts)
}

// SendData is a paid mutator transaction binding the contract method 0x658055e0.
//
// Solidity: function sendData(address _receiver, string _encryptedData, string _dataHash, string _signedHash) returns()
func (_SecureDataTransfer *SecureDataTransferTransactor) SendData(opts *bind.TransactOpts, _receiver common.Address, _encryptedData string, _dataHash string, _signedHash string) (*types.Transaction, error) {
	return _SecureDataTransfer.contract.Transact(opts, "sendData", _receiver, _encryptedData, _dataHash, _signedHash)
}

// SendData is a paid mutator transaction binding the contract method 0x658055e0.
//
// Solidity: function sendData(address _receiver, string _encryptedData, string _dataHash, string _signedHash) returns()
func (_SecureDataTransfer *SecureDataTransferSession) SendData(_receiver common.Address, _encryptedData string, _dataHash string, _signedHash string) (*types.Transaction, error) {
	return _SecureDataTransfer.Contract.SendData(&_SecureDataTransfer.TransactOpts, _receiver, _encryptedData, _dataHash, _signedHash)
}

// SendData is a paid mutator transaction binding the contract method 0x658055e0.
//
// Solidity: function sendData(address _receiver, string _encryptedData, string _dataHash, string _signedHash) returns()
func (_SecureDataTransfer *SecureDataTransferTransactorSession) SendData(_receiver common.Address, _encryptedData string, _dataHash string, _signedHash string) (*types.Transaction, error) {
	return _SecureDataTransfer.Contract.SendData(&_SecureDataTransfer.TransactOpts, _receiver, _encryptedData, _dataHash, _signedHash)
}

// SecureDataTransferDataSentIterator is returned from FilterDataSent and is used to iterate over the raw logs and unpacked data for DataSent events raised by the SecureDataTransfer contract.
type SecureDataTransferDataSentIterator struct {
	Event *SecureDataTransferDataSent // Event containing the contract specifics and raw log

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
func (it *SecureDataTransferDataSentIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SecureDataTransferDataSent)
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
		it.Event = new(SecureDataTransferDataSent)
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
func (it *SecureDataTransferDataSentIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SecureDataTransferDataSentIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SecureDataTransferDataSent represents a DataSent event raised by the SecureDataTransfer contract.
type SecureDataTransferDataSent struct {
	Sender      common.Address
	Receiver    common.Address
	TimeCreated *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterDataSent is a free log retrieval operation binding the contract event 0xa66e03c9001bca5fd5a6b12c7ff0565f91adab37b440198b3d413712a3741f36.
//
// Solidity: event DataSent(address indexed sender, address indexed receiver, uint256 timeCreated)
func (_SecureDataTransfer *SecureDataTransferFilterer) FilterDataSent(opts *bind.FilterOpts, sender []common.Address, receiver []common.Address) (*SecureDataTransferDataSentIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var receiverRule []interface{}
	for _, receiverItem := range receiver {
		receiverRule = append(receiverRule, receiverItem)
	}

	logs, sub, err := _SecureDataTransfer.contract.FilterLogs(opts, "DataSent", senderRule, receiverRule)
	if err != nil {
		return nil, err
	}
	return &SecureDataTransferDataSentIterator{contract: _SecureDataTransfer.contract, event: "DataSent", logs: logs, sub: sub}, nil
}

// WatchDataSent is a free log subscription operation binding the contract event 0xa66e03c9001bca5fd5a6b12c7ff0565f91adab37b440198b3d413712a3741f36.
//
// Solidity: event DataSent(address indexed sender, address indexed receiver, uint256 timeCreated)
func (_SecureDataTransfer *SecureDataTransferFilterer) WatchDataSent(opts *bind.WatchOpts, sink chan<- *SecureDataTransferDataSent, sender []common.Address, receiver []common.Address) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var receiverRule []interface{}
	for _, receiverItem := range receiver {
		receiverRule = append(receiverRule, receiverItem)
	}

	logs, sub, err := _SecureDataTransfer.contract.WatchLogs(opts, "DataSent", senderRule, receiverRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SecureDataTransferDataSent)
				if err := _SecureDataTransfer.contract.UnpackLog(event, "DataSent", log); err != nil {
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

// ParseDataSent is a log parse operation binding the contract event 0xa66e03c9001bca5fd5a6b12c7ff0565f91adab37b440198b3d413712a3741f36.
//
// Solidity: event DataSent(address indexed sender, address indexed receiver, uint256 timeCreated)
func (_SecureDataTransfer *SecureDataTransferFilterer) ParseDataSent(log types.Log) (*SecureDataTransferDataSent, error) {
	event := new(SecureDataTransferDataSent)
	if err := _SecureDataTransfer.contract.UnpackLog(event, "DataSent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
