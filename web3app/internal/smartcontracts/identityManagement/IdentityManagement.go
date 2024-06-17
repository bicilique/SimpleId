package identityManagement

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

// IdentityManagementMetaData contains all meta data concerning the IdentityManagement contract.
var IdentityManagementMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"nik\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"email\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"street\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"country\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"birthdate\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"status\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"active\",\"type\":\"bool\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"createdAt\",\"type\":\"uint256\"}],\"name\":\"IdentityCreated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"nik\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"email\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"street\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"country\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"birthdate\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"status\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"active\",\"type\":\"bool\"}],\"name\":\"IdentityUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"stakeholder\",\"type\":\"address\"}],\"name\":\"StakeholderAdded\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_stakeholder\",\"type\":\"address\"}],\"name\":\"addStakeholder\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_nik\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_email\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_street\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_country\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_birthdate\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_status\",\"type\":\"string\"},{\"internalType\":\"bool\",\"name\":\"_active\",\"type\":\"bool\"}],\"name\":\"createIdentity\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getIdentity\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"issuer\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"nik\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"email\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"street\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"country\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"birthdate\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"status\",\"type\":\"string\"},{\"internalType\":\"bool\",\"name\":\"active\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"createdAt\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"isStakeholder\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_nik\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_email\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_street\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_country\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_birthdate\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_status\",\"type\":\"string\"},{\"internalType\":\"bool\",\"name\":\"_active\",\"type\":\"bool\"}],\"name\":\"updateIdentity\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b506001600b60003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548160ff0219169083151502179055506001600c60003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548160ff02191690831515021790555060008060006101000a81548160ff0219169083151502179055506117a0806100ea6000396000f3fe608060405234801561001057600080fd5b50600436106100575760003560e01c80632bfb03df1461005c57806336afc6fa146100785780636b114a1b1461009f578063e5c42fd1146100bb578063ef037b90146100d7575b600080fd5b61007660048036038101906100719190610d34565b610107565b005b61008061029a565b6040516100969a99989796959493929190610f96565b60405180910390f35b6100b960048036038101906100b49190610d34565b61078b565b005b6100d560048036038101906100d0919061108f565b6109cb565b005b6100f160048036038101906100ec919061108f565b610b82565b6040516100fe91906110bc565b60405180910390f35b600b60003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff16610193576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161018a90611149565b60405180910390fd5b876001800190816101a49190611375565b5086600160020190816101b79190611375565b5085600160030190816101ca9190611375565b5084600160040190816101dd9190611375565b5083600160050190816101f09190611375565b5082600160060190816102039190611375565b5081600160070190816102169190611375565b5080600160080160006101000a81548160ff0219169083151502179055503373ffffffffffffffffffffffffffffffffffffffff167f042ebdc9ce857f38e573c2e5ea247e764e3a74c8e425d2a166e438009e1f74008989898989898989604051610288989796959493929190611447565b60405180910390a25050505050505050565b60006060806060806060806060600080600c60003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff16610336576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161032d90611568565b60405180910390fd5b600160000160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1660018001600160020160016003016001600401600160050160016006016001600701600160080160009054906101000a900460ff166001600901548880546103a390611198565b80601f01602080910402602001604051908101604052809291908181526020018280546103cf90611198565b801561041c5780601f106103f15761010080835404028352916020019161041c565b820191906000526020600020905b8154815290600101906020018083116103ff57829003601f168201915b5050505050985087805461042f90611198565b80601f016020809104026020016040519081016040528092919081815260200182805461045b90611198565b80156104a85780601f1061047d576101008083540402835291602001916104a8565b820191906000526020600020905b81548152906001019060200180831161048b57829003601f168201915b505050505097508680546104bb90611198565b80601f01602080910402602001604051908101604052809291908181526020018280546104e790611198565b80156105345780601f1061050957610100808354040283529160200191610534565b820191906000526020600020905b81548152906001019060200180831161051757829003601f168201915b5050505050965085805461054790611198565b80601f016020809104026020016040519081016040528092919081815260200182805461057390611198565b80156105c05780601f10610595576101008083540402835291602001916105c0565b820191906000526020600020905b8154815290600101906020018083116105a357829003601f168201915b505050505095508480546105d390611198565b80601f01602080910402602001604051908101604052809291908181526020018280546105ff90611198565b801561064c5780601f106106215761010080835404028352916020019161064c565b820191906000526020600020905b81548152906001019060200180831161062f57829003601f168201915b5050505050945083805461065f90611198565b80601f016020809104026020016040519081016040528092919081815260200182805461068b90611198565b80156106d85780601f106106ad576101008083540402835291602001916106d8565b820191906000526020600020905b8154815290600101906020018083116106bb57829003601f168201915b505050505093508280546106eb90611198565b80601f016020809104026020016040519081016040528092919081815260200182805461071790611198565b80156107645780601f1061073957610100808354040283529160200191610764565b820191906000526020600020905b81548152906001019060200180831161074757829003601f168201915b50505050509250995099509950995099509950995099509950995090919293949596979899565b60008054906101000a900460ff16156107d9576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016107d0906115fa565b60405180910390fd5b6040518061014001604052803373ffffffffffffffffffffffffffffffffffffffff168152602001898152602001888152602001878152602001868152602001858152602001848152602001838152602001821515815260200142815250600160008201518160000160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060208201518160010190816108959190611375565b5060408201518160020190816108ab9190611375565b5060608201518160030190816108c19190611375565b5060808201518160040190816108d79190611375565b5060a08201518160050190816108ed9190611375565b5060c08201518160060190816109039190611375565b5060e08201518160070190816109199190611375565b506101008201518160080160006101000a81548160ff021916908315150217905550610120820151816009015590505060016000806101000a81548160ff0219169083151502179055503373ffffffffffffffffffffffffffffffffffffffff167f75596379ac4172e040c2f0a4d12b72cee72ea3dc485d71cced2d9d6d2e2470298989898989898989426040516109b99998979695949392919061161a565b60405180910390a25050505050505050565b600b60003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff16610a57576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610a4e90611149565b60405180910390fd5b600c60008273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff1615610ae4576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610adb9061174a565b60405180910390fd5b6001600c60008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548160ff0219169083151502179055508073ffffffffffffffffffffffffffffffffffffffff167fd9b802c55a7d0aabbc6554676754b5b4aabc8edf1d540b77812f81f6d4409ea160405160405180910390a250565b600c6020528060005260406000206000915054906101000a900460ff1681565b6000604051905090565b600080fd5b600080fd5b600080fd5b600080fd5b6000601f19601f8301169050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b610c0982610bc0565b810181811067ffffffffffffffff82111715610c2857610c27610bd1565b5b80604052505050565b6000610c3b610ba2565b9050610c478282610c00565b919050565b600067ffffffffffffffff821115610c6757610c66610bd1565b5b610c7082610bc0565b9050602081019050919050565b82818337600083830152505050565b6000610c9f610c9a84610c4c565b610c31565b905082815260208101848484011115610cbb57610cba610bbb565b5b610cc6848285610c7d565b509392505050565b600082601f830112610ce357610ce2610bb6565b5b8135610cf3848260208601610c8c565b91505092915050565b60008115159050919050565b610d1181610cfc565b8114610d1c57600080fd5b50565b600081359050610d2e81610d08565b92915050565b600080600080600080600080610100898b031215610d5557610d54610bac565b5b600089013567ffffffffffffffff811115610d7357610d72610bb1565b5b610d7f8b828c01610cce565b985050602089013567ffffffffffffffff811115610da057610d9f610bb1565b5b610dac8b828c01610cce565b975050604089013567ffffffffffffffff811115610dcd57610dcc610bb1565b5b610dd98b828c01610cce565b965050606089013567ffffffffffffffff811115610dfa57610df9610bb1565b5b610e068b828c01610cce565b955050608089013567ffffffffffffffff811115610e2757610e26610bb1565b5b610e338b828c01610cce565b94505060a089013567ffffffffffffffff811115610e5457610e53610bb1565b5b610e608b828c01610cce565b93505060c089013567ffffffffffffffff811115610e8157610e80610bb1565b5b610e8d8b828c01610cce565b92505060e0610e9e8b828c01610d1f565b9150509295985092959890939650565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000610ed982610eae565b9050919050565b610ee981610ece565b82525050565b600081519050919050565b600082825260208201905092915050565b60005b83811015610f29578082015181840152602081019050610f0e565b60008484015250505050565b6000610f4082610eef565b610f4a8185610efa565b9350610f5a818560208601610f0b565b610f6381610bc0565b840191505092915050565b610f7781610cfc565b82525050565b6000819050919050565b610f9081610f7d565b82525050565b600061014082019050610fac600083018d610ee0565b8181036020830152610fbe818c610f35565b90508181036040830152610fd2818b610f35565b90508181036060830152610fe6818a610f35565b90508181036080830152610ffa8189610f35565b905081810360a083015261100e8188610f35565b905081810360c08301526110228187610f35565b905081810360e08301526110368186610f35565b9050611046610100830185610f6e565b611054610120830184610f87565b9b9a5050505050505050505050565b61106c81610ece565b811461107757600080fd5b50565b60008135905061108981611063565b92915050565b6000602082840312156110a5576110a4610bac565b5b60006110b38482850161107a565b91505092915050565b60006020820190506110d16000830184610f6e565b92915050565b7f4f6e6c7920617574686f72697a656420697373756572732063616e207065726660008201527f6f726d207468697320616374696f6e2e00000000000000000000000000000000602082015250565b6000611133603083610efa565b915061113e826110d7565b604082019050919050565b6000602082019050818103600083015261116281611126565b9050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b600060028204905060018216806111b057607f821691505b6020821081036111c3576111c2611169565b5b50919050565b60008190508160005260206000209050919050565b60006020601f8301049050919050565b600082821b905092915050565b60006008830261122b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff826111ee565b61123586836111ee565b95508019841693508086168417925050509392505050565b6000819050919050565b600061127261126d61126884610f7d565b61124d565b610f7d565b9050919050565b6000819050919050565b61128c83611257565b6112a061129882611279565b8484546111fb565b825550505050565b600090565b6112b56112a8565b6112c0818484611283565b505050565b5b818110156112e4576112d96000826112ad565b6001810190506112c6565b5050565b601f821115611329576112fa816111c9565b611303846111de565b81016020851015611312578190505b61132661131e856111de565b8301826112c5565b50505b505050565b600082821c905092915050565b600061134c6000198460080261132e565b1980831691505092915050565b6000611365838361133b565b9150826002028217905092915050565b61137e82610eef565b67ffffffffffffffff81111561139757611396610bd1565b5b6113a18254611198565b6113ac8282856112e8565b600060209050601f8311600181146113df57600084156113cd578287015190505b6113d78582611359565b86555061143f565b601f1984166113ed866111c9565b60005b82811015611415578489015182556001820191506020850194506020810190506113f0565b86831015611432578489015161142e601f89168261133b565b8355505b6001600288020188555050505b505050505050565b6000610100820190508181036000830152611462818b610f35565b90508181036020830152611476818a610f35565b9050818103604083015261148a8189610f35565b9050818103606083015261149e8188610f35565b905081810360808301526114b28187610f35565b905081810360a08301526114c68186610f35565b905081810360c08301526114da8185610f35565b90506114e960e0830184610f6e565b9998505050505050505050565b7f4f6e6c7920746865206964656e74697479207374616b65686f6c64657220636160008201527f6e20706572666f726d207468697320616374696f6e2e00000000000000000000602082015250565b6000611552603683610efa565b915061155d826114f6565b604082019050919050565b6000602082019050818103600083015261158181611545565b9050919050565b7f4964656e746974792068617320616c7265616479206265656e2063726561746560008201527f642e000000000000000000000000000000000000000000000000000000000000602082015250565b60006115e4602283610efa565b91506115ef82611588565b604082019050919050565b60006020820190508181036000830152611613816115d7565b9050919050565b6000610120820190508181036000830152611635818c610f35565b90508181036020830152611649818b610f35565b9050818103604083015261165d818a610f35565b905081810360608301526116718189610f35565b905081810360808301526116858188610f35565b905081810360a08301526116998187610f35565b905081810360c08301526116ad8186610f35565b90506116bc60e0830185610f6e565b6116ca610100830184610f87565b9a9950505050505050505050565b7f4164647265737320697320616c72656164792061207374616b65686f6c64657260008201527f2e00000000000000000000000000000000000000000000000000000000000000602082015250565b6000611734602183610efa565b915061173f826116d8565b604082019050919050565b6000602082019050818103600083015261176381611727565b905091905056fea2646970667358221220737879ae90bf8a423ca8eae865c7f77462754d1c29c71c58a16c7badd5bc2e2f64736f6c63430008120033",
}

// IdentityManagementABI is the input ABI used to generate the binding from.
// Deprecated: Use IdentityManagementMetaData.ABI instead.
var IdentityManagementABI = IdentityManagementMetaData.ABI

// IdentityManagementBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use IdentityManagementMetaData.Bin instead.
var IdentityManagementBin = IdentityManagementMetaData.Bin

// DeployIdentityManagement deploys a new Ethereum contract, binding an instance of IdentityManagement to it.
func DeployIdentityManagement(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *IdentityManagement, error) {
	parsed, err := IdentityManagementMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(IdentityManagementBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &IdentityManagement{IdentityManagementCaller: IdentityManagementCaller{contract: contract}, IdentityManagementTransactor: IdentityManagementTransactor{contract: contract}, IdentityManagementFilterer: IdentityManagementFilterer{contract: contract}}, nil
}

// IdentityManagement is an auto generated Go binding around an Ethereum contract.
type IdentityManagement struct {
	IdentityManagementCaller     // Read-only binding to the contract
	IdentityManagementTransactor // Write-only binding to the contract
	IdentityManagementFilterer   // Log filterer for contract events
}

// IdentityManagementCaller is an auto generated read-only Go binding around an Ethereum contract.
type IdentityManagementCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IdentityManagementTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IdentityManagementTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IdentityManagementFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IdentityManagementFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IdentityManagementSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IdentityManagementSession struct {
	Contract     *IdentityManagement // Generic contract binding to set the session for
	CallOpts     bind.CallOpts       // Call options to use throughout this session
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// IdentityManagementCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IdentityManagementCallerSession struct {
	Contract *IdentityManagementCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts             // Call options to use throughout this session
}

// IdentityManagementTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IdentityManagementTransactorSession struct {
	Contract     *IdentityManagementTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts             // Transaction auth options to use throughout this session
}

// IdentityManagementRaw is an auto generated low-level Go binding around an Ethereum contract.
type IdentityManagementRaw struct {
	Contract *IdentityManagement // Generic contract binding to access the raw methods on
}

// IdentityManagementCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IdentityManagementCallerRaw struct {
	Contract *IdentityManagementCaller // Generic read-only contract binding to access the raw methods on
}

// IdentityManagementTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IdentityManagementTransactorRaw struct {
	Contract *IdentityManagementTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIdentityManagement creates a new instance of IdentityManagement, bound to a specific deployed contract.
func NewIdentityManagement(address common.Address, backend bind.ContractBackend) (*IdentityManagement, error) {
	contract, err := bindIdentityManagement(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IdentityManagement{IdentityManagementCaller: IdentityManagementCaller{contract: contract}, IdentityManagementTransactor: IdentityManagementTransactor{contract: contract}, IdentityManagementFilterer: IdentityManagementFilterer{contract: contract}}, nil
}

// NewIdentityManagementCaller creates a new read-only instance of IdentityManagement, bound to a specific deployed contract.
func NewIdentityManagementCaller(address common.Address, caller bind.ContractCaller) (*IdentityManagementCaller, error) {
	contract, err := bindIdentityManagement(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IdentityManagementCaller{contract: contract}, nil
}

// NewIdentityManagementTransactor creates a new write-only instance of IdentityManagement, bound to a specific deployed contract.
func NewIdentityManagementTransactor(address common.Address, transactor bind.ContractTransactor) (*IdentityManagementTransactor, error) {
	contract, err := bindIdentityManagement(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IdentityManagementTransactor{contract: contract}, nil
}

// NewIdentityManagementFilterer creates a new log filterer instance of IdentityManagement, bound to a specific deployed contract.
func NewIdentityManagementFilterer(address common.Address, filterer bind.ContractFilterer) (*IdentityManagementFilterer, error) {
	contract, err := bindIdentityManagement(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IdentityManagementFilterer{contract: contract}, nil
}

// bindIdentityManagement binds a generic wrapper to an already deployed contract.
func bindIdentityManagement(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := IdentityManagementMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IdentityManagement *IdentityManagementRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IdentityManagement.Contract.IdentityManagementCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IdentityManagement *IdentityManagementRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IdentityManagement.Contract.IdentityManagementTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IdentityManagement *IdentityManagementRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IdentityManagement.Contract.IdentityManagementTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IdentityManagement *IdentityManagementCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IdentityManagement.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IdentityManagement *IdentityManagementTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IdentityManagement.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IdentityManagement *IdentityManagementTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IdentityManagement.Contract.contract.Transact(opts, method, params...)
}

// GetIdentity is a free data retrieval call binding the contract method 0x36afc6fa.
//
// Solidity: function getIdentity() view returns(address issuer, string name, string nik, string email, string street, string country, string birthdate, string status, bool active, uint256 createdAt)
func (_IdentityManagement *IdentityManagementCaller) GetIdentity(opts *bind.CallOpts) (struct {
	Issuer    common.Address
	Name      string
	Nik       string
	Email     string
	Street    string
	Country   string
	Birthdate string
	Status    string
	Active    bool
	CreatedAt *big.Int
}, error) {
	var out []interface{}
	err := _IdentityManagement.contract.Call(opts, &out, "getIdentity")

	outstruct := new(struct {
		Issuer    common.Address
		Name      string
		Nik       string
		Email     string
		Street    string
		Country   string
		Birthdate string
		Status    string
		Active    bool
		CreatedAt *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Issuer = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.Name = *abi.ConvertType(out[1], new(string)).(*string)
	outstruct.Nik = *abi.ConvertType(out[2], new(string)).(*string)
	outstruct.Email = *abi.ConvertType(out[3], new(string)).(*string)
	outstruct.Street = *abi.ConvertType(out[4], new(string)).(*string)
	outstruct.Country = *abi.ConvertType(out[5], new(string)).(*string)
	outstruct.Birthdate = *abi.ConvertType(out[6], new(string)).(*string)
	outstruct.Status = *abi.ConvertType(out[7], new(string)).(*string)
	outstruct.Active = *abi.ConvertType(out[8], new(bool)).(*bool)
	outstruct.CreatedAt = *abi.ConvertType(out[9], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// GetIdentity is a free data retrieval call binding the contract method 0x36afc6fa.
//
// Solidity: function getIdentity() view returns(address issuer, string name, string nik, string email, string street, string country, string birthdate, string status, bool active, uint256 createdAt)
func (_IdentityManagement *IdentityManagementSession) GetIdentity() (struct {
	Issuer    common.Address
	Name      string
	Nik       string
	Email     string
	Street    string
	Country   string
	Birthdate string
	Status    string
	Active    bool
	CreatedAt *big.Int
}, error) {
	return _IdentityManagement.Contract.GetIdentity(&_IdentityManagement.CallOpts)
}

// GetIdentity is a free data retrieval call binding the contract method 0x36afc6fa.
//
// Solidity: function getIdentity() view returns(address issuer, string name, string nik, string email, string street, string country, string birthdate, string status, bool active, uint256 createdAt)
func (_IdentityManagement *IdentityManagementCallerSession) GetIdentity() (struct {
	Issuer    common.Address
	Name      string
	Nik       string
	Email     string
	Street    string
	Country   string
	Birthdate string
	Status    string
	Active    bool
	CreatedAt *big.Int
}, error) {
	return _IdentityManagement.Contract.GetIdentity(&_IdentityManagement.CallOpts)
}

// IsStakeholder is a free data retrieval call binding the contract method 0xef037b90.
//
// Solidity: function isStakeholder(address ) view returns(bool)
func (_IdentityManagement *IdentityManagementCaller) IsStakeholder(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _IdentityManagement.contract.Call(opts, &out, "isStakeholder", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsStakeholder is a free data retrieval call binding the contract method 0xef037b90.
//
// Solidity: function isStakeholder(address ) view returns(bool)
func (_IdentityManagement *IdentityManagementSession) IsStakeholder(arg0 common.Address) (bool, error) {
	return _IdentityManagement.Contract.IsStakeholder(&_IdentityManagement.CallOpts, arg0)
}

// IsStakeholder is a free data retrieval call binding the contract method 0xef037b90.
//
// Solidity: function isStakeholder(address ) view returns(bool)
func (_IdentityManagement *IdentityManagementCallerSession) IsStakeholder(arg0 common.Address) (bool, error) {
	return _IdentityManagement.Contract.IsStakeholder(&_IdentityManagement.CallOpts, arg0)
}

// AddStakeholder is a paid mutator transaction binding the contract method 0xe5c42fd1.
//
// Solidity: function addStakeholder(address _stakeholder) returns()
func (_IdentityManagement *IdentityManagementTransactor) AddStakeholder(opts *bind.TransactOpts, _stakeholder common.Address) (*types.Transaction, error) {
	return _IdentityManagement.contract.Transact(opts, "addStakeholder", _stakeholder)
}

// AddStakeholder is a paid mutator transaction binding the contract method 0xe5c42fd1.
//
// Solidity: function addStakeholder(address _stakeholder) returns()
func (_IdentityManagement *IdentityManagementSession) AddStakeholder(_stakeholder common.Address) (*types.Transaction, error) {
	return _IdentityManagement.Contract.AddStakeholder(&_IdentityManagement.TransactOpts, _stakeholder)
}

// AddStakeholder is a paid mutator transaction binding the contract method 0xe5c42fd1.
//
// Solidity: function addStakeholder(address _stakeholder) returns()
func (_IdentityManagement *IdentityManagementTransactorSession) AddStakeholder(_stakeholder common.Address) (*types.Transaction, error) {
	return _IdentityManagement.Contract.AddStakeholder(&_IdentityManagement.TransactOpts, _stakeholder)
}

// CreateIdentity is a paid mutator transaction binding the contract method 0x6b114a1b.
//
// Solidity: function createIdentity(string _name, string _nik, string _email, string _street, string _country, string _birthdate, string _status, bool _active) returns()
func (_IdentityManagement *IdentityManagementTransactor) CreateIdentity(opts *bind.TransactOpts, _name string, _nik string, _email string, _street string, _country string, _birthdate string, _status string, _active bool) (*types.Transaction, error) {
	return _IdentityManagement.contract.Transact(opts, "createIdentity", _name, _nik, _email, _street, _country, _birthdate, _status, _active)
}

// CreateIdentity is a paid mutator transaction binding the contract method 0x6b114a1b.
//
// Solidity: function createIdentity(string _name, string _nik, string _email, string _street, string _country, string _birthdate, string _status, bool _active) returns()
func (_IdentityManagement *IdentityManagementSession) CreateIdentity(_name string, _nik string, _email string, _street string, _country string, _birthdate string, _status string, _active bool) (*types.Transaction, error) {
	return _IdentityManagement.Contract.CreateIdentity(&_IdentityManagement.TransactOpts, _name, _nik, _email, _street, _country, _birthdate, _status, _active)
}

// CreateIdentity is a paid mutator transaction binding the contract method 0x6b114a1b.
//
// Solidity: function createIdentity(string _name, string _nik, string _email, string _street, string _country, string _birthdate, string _status, bool _active) returns()
func (_IdentityManagement *IdentityManagementTransactorSession) CreateIdentity(_name string, _nik string, _email string, _street string, _country string, _birthdate string, _status string, _active bool) (*types.Transaction, error) {
	return _IdentityManagement.Contract.CreateIdentity(&_IdentityManagement.TransactOpts, _name, _nik, _email, _street, _country, _birthdate, _status, _active)
}

// UpdateIdentity is a paid mutator transaction binding the contract method 0x2bfb03df.
//
// Solidity: function updateIdentity(string _name, string _nik, string _email, string _street, string _country, string _birthdate, string _status, bool _active) returns()
func (_IdentityManagement *IdentityManagementTransactor) UpdateIdentity(opts *bind.TransactOpts, _name string, _nik string, _email string, _street string, _country string, _birthdate string, _status string, _active bool) (*types.Transaction, error) {
	return _IdentityManagement.contract.Transact(opts, "updateIdentity", _name, _nik, _email, _street, _country, _birthdate, _status, _active)
}

// UpdateIdentity is a paid mutator transaction binding the contract method 0x2bfb03df.
//
// Solidity: function updateIdentity(string _name, string _nik, string _email, string _street, string _country, string _birthdate, string _status, bool _active) returns()
func (_IdentityManagement *IdentityManagementSession) UpdateIdentity(_name string, _nik string, _email string, _street string, _country string, _birthdate string, _status string, _active bool) (*types.Transaction, error) {
	return _IdentityManagement.Contract.UpdateIdentity(&_IdentityManagement.TransactOpts, _name, _nik, _email, _street, _country, _birthdate, _status, _active)
}

// UpdateIdentity is a paid mutator transaction binding the contract method 0x2bfb03df.
//
// Solidity: function updateIdentity(string _name, string _nik, string _email, string _street, string _country, string _birthdate, string _status, bool _active) returns()
func (_IdentityManagement *IdentityManagementTransactorSession) UpdateIdentity(_name string, _nik string, _email string, _street string, _country string, _birthdate string, _status string, _active bool) (*types.Transaction, error) {
	return _IdentityManagement.Contract.UpdateIdentity(&_IdentityManagement.TransactOpts, _name, _nik, _email, _street, _country, _birthdate, _status, _active)
}

// IdentityManagementIdentityCreatedIterator is returned from FilterIdentityCreated and is used to iterate over the raw logs and unpacked data for IdentityCreated events raised by the IdentityManagement contract.
type IdentityManagementIdentityCreatedIterator struct {
	Event *IdentityManagementIdentityCreated // Event containing the contract specifics and raw log

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
func (it *IdentityManagementIdentityCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IdentityManagementIdentityCreated)
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
		it.Event = new(IdentityManagementIdentityCreated)
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
func (it *IdentityManagementIdentityCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IdentityManagementIdentityCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IdentityManagementIdentityCreated represents a IdentityCreated event raised by the IdentityManagement contract.
type IdentityManagementIdentityCreated struct {
	Owner     common.Address
	Name      string
	Nik       string
	Email     string
	Street    string
	Country   string
	Birthdate string
	Status    string
	Active    bool
	CreatedAt *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterIdentityCreated is a free log retrieval operation binding the contract event 0x75596379ac4172e040c2f0a4d12b72cee72ea3dc485d71cced2d9d6d2e247029.
//
// Solidity: event IdentityCreated(address indexed owner, string name, string nik, string email, string street, string country, string birthdate, string status, bool active, uint256 createdAt)
func (_IdentityManagement *IdentityManagementFilterer) FilterIdentityCreated(opts *bind.FilterOpts, owner []common.Address) (*IdentityManagementIdentityCreatedIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _IdentityManagement.contract.FilterLogs(opts, "IdentityCreated", ownerRule)
	if err != nil {
		return nil, err
	}
	return &IdentityManagementIdentityCreatedIterator{contract: _IdentityManagement.contract, event: "IdentityCreated", logs: logs, sub: sub}, nil
}

// WatchIdentityCreated is a free log subscription operation binding the contract event 0x75596379ac4172e040c2f0a4d12b72cee72ea3dc485d71cced2d9d6d2e247029.
//
// Solidity: event IdentityCreated(address indexed owner, string name, string nik, string email, string street, string country, string birthdate, string status, bool active, uint256 createdAt)
func (_IdentityManagement *IdentityManagementFilterer) WatchIdentityCreated(opts *bind.WatchOpts, sink chan<- *IdentityManagementIdentityCreated, owner []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _IdentityManagement.contract.WatchLogs(opts, "IdentityCreated", ownerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IdentityManagementIdentityCreated)
				if err := _IdentityManagement.contract.UnpackLog(event, "IdentityCreated", log); err != nil {
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

// ParseIdentityCreated is a log parse operation binding the contract event 0x75596379ac4172e040c2f0a4d12b72cee72ea3dc485d71cced2d9d6d2e247029.
//
// Solidity: event IdentityCreated(address indexed owner, string name, string nik, string email, string street, string country, string birthdate, string status, bool active, uint256 createdAt)
func (_IdentityManagement *IdentityManagementFilterer) ParseIdentityCreated(log types.Log) (*IdentityManagementIdentityCreated, error) {
	event := new(IdentityManagementIdentityCreated)
	if err := _IdentityManagement.contract.UnpackLog(event, "IdentityCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IdentityManagementIdentityUpdatedIterator is returned from FilterIdentityUpdated and is used to iterate over the raw logs and unpacked data for IdentityUpdated events raised by the IdentityManagement contract.
type IdentityManagementIdentityUpdatedIterator struct {
	Event *IdentityManagementIdentityUpdated // Event containing the contract specifics and raw log

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
func (it *IdentityManagementIdentityUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IdentityManagementIdentityUpdated)
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
		it.Event = new(IdentityManagementIdentityUpdated)
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
func (it *IdentityManagementIdentityUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IdentityManagementIdentityUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IdentityManagementIdentityUpdated represents a IdentityUpdated event raised by the IdentityManagement contract.
type IdentityManagementIdentityUpdated struct {
	Owner     common.Address
	Name      string
	Nik       string
	Email     string
	Street    string
	Country   string
	Birthdate string
	Status    string
	Active    bool
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterIdentityUpdated is a free log retrieval operation binding the contract event 0x042ebdc9ce857f38e573c2e5ea247e764e3a74c8e425d2a166e438009e1f7400.
//
// Solidity: event IdentityUpdated(address indexed owner, string name, string nik, string email, string street, string country, string birthdate, string status, bool active)
func (_IdentityManagement *IdentityManagementFilterer) FilterIdentityUpdated(opts *bind.FilterOpts, owner []common.Address) (*IdentityManagementIdentityUpdatedIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _IdentityManagement.contract.FilterLogs(opts, "IdentityUpdated", ownerRule)
	if err != nil {
		return nil, err
	}
	return &IdentityManagementIdentityUpdatedIterator{contract: _IdentityManagement.contract, event: "IdentityUpdated", logs: logs, sub: sub}, nil
}

// WatchIdentityUpdated is a free log subscription operation binding the contract event 0x042ebdc9ce857f38e573c2e5ea247e764e3a74c8e425d2a166e438009e1f7400.
//
// Solidity: event IdentityUpdated(address indexed owner, string name, string nik, string email, string street, string country, string birthdate, string status, bool active)
func (_IdentityManagement *IdentityManagementFilterer) WatchIdentityUpdated(opts *bind.WatchOpts, sink chan<- *IdentityManagementIdentityUpdated, owner []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _IdentityManagement.contract.WatchLogs(opts, "IdentityUpdated", ownerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IdentityManagementIdentityUpdated)
				if err := _IdentityManagement.contract.UnpackLog(event, "IdentityUpdated", log); err != nil {
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

// ParseIdentityUpdated is a log parse operation binding the contract event 0x042ebdc9ce857f38e573c2e5ea247e764e3a74c8e425d2a166e438009e1f7400.
//
// Solidity: event IdentityUpdated(address indexed owner, string name, string nik, string email, string street, string country, string birthdate, string status, bool active)
func (_IdentityManagement *IdentityManagementFilterer) ParseIdentityUpdated(log types.Log) (*IdentityManagementIdentityUpdated, error) {
	event := new(IdentityManagementIdentityUpdated)
	if err := _IdentityManagement.contract.UnpackLog(event, "IdentityUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IdentityManagementStakeholderAddedIterator is returned from FilterStakeholderAdded and is used to iterate over the raw logs and unpacked data for StakeholderAdded events raised by the IdentityManagement contract.
type IdentityManagementStakeholderAddedIterator struct {
	Event *IdentityManagementStakeholderAdded // Event containing the contract specifics and raw log

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
func (it *IdentityManagementStakeholderAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IdentityManagementStakeholderAdded)
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
		it.Event = new(IdentityManagementStakeholderAdded)
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
func (it *IdentityManagementStakeholderAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IdentityManagementStakeholderAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IdentityManagementStakeholderAdded represents a StakeholderAdded event raised by the IdentityManagement contract.
type IdentityManagementStakeholderAdded struct {
	Stakeholder common.Address
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterStakeholderAdded is a free log retrieval operation binding the contract event 0xd9b802c55a7d0aabbc6554676754b5b4aabc8edf1d540b77812f81f6d4409ea1.
//
// Solidity: event StakeholderAdded(address indexed stakeholder)
func (_IdentityManagement *IdentityManagementFilterer) FilterStakeholderAdded(opts *bind.FilterOpts, stakeholder []common.Address) (*IdentityManagementStakeholderAddedIterator, error) {

	var stakeholderRule []interface{}
	for _, stakeholderItem := range stakeholder {
		stakeholderRule = append(stakeholderRule, stakeholderItem)
	}

	logs, sub, err := _IdentityManagement.contract.FilterLogs(opts, "StakeholderAdded", stakeholderRule)
	if err != nil {
		return nil, err
	}
	return &IdentityManagementStakeholderAddedIterator{contract: _IdentityManagement.contract, event: "StakeholderAdded", logs: logs, sub: sub}, nil
}

// WatchStakeholderAdded is a free log subscription operation binding the contract event 0xd9b802c55a7d0aabbc6554676754b5b4aabc8edf1d540b77812f81f6d4409ea1.
//
// Solidity: event StakeholderAdded(address indexed stakeholder)
func (_IdentityManagement *IdentityManagementFilterer) WatchStakeholderAdded(opts *bind.WatchOpts, sink chan<- *IdentityManagementStakeholderAdded, stakeholder []common.Address) (event.Subscription, error) {

	var stakeholderRule []interface{}
	for _, stakeholderItem := range stakeholder {
		stakeholderRule = append(stakeholderRule, stakeholderItem)
	}

	logs, sub, err := _IdentityManagement.contract.WatchLogs(opts, "StakeholderAdded", stakeholderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IdentityManagementStakeholderAdded)
				if err := _IdentityManagement.contract.UnpackLog(event, "StakeholderAdded", log); err != nil {
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

// ParseStakeholderAdded is a log parse operation binding the contract event 0xd9b802c55a7d0aabbc6554676754b5b4aabc8edf1d540b77812f81f6d4409ea1.
//
// Solidity: event StakeholderAdded(address indexed stakeholder)
func (_IdentityManagement *IdentityManagementFilterer) ParseStakeholderAdded(log types.Log) (*IdentityManagementStakeholderAdded, error) {
	event := new(IdentityManagementStakeholderAdded)
	if err := _IdentityManagement.contract.UnpackLog(event, "StakeholderAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
