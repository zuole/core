// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package commitreveal

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

// CommitrevealABI is the input ABI used to generate the binding from.
const CommitrevealABI = "[{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"campaigns\",\"outputs\":[{\"name\":\"startBlock\",\"type\":\"uint256\"},{\"name\":\"commitDuration\",\"type\":\"uint256\"},{\"name\":\"revealDuration\",\"type\":\"uint256\"},{\"name\":\"revealThreshold\",\"type\":\"uint256\"},{\"name\":\"commitNum\",\"type\":\"uint256\"},{\"name\":\"revealNum\",\"type\":\"uint256\"},{\"name\":\"generatedRandom\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_addr\",\"type\":\"address\"}],\"name\":\"removeFromWhitelist\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"isOwner\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_cid\",\"type\":\"uint256\"},{\"name\":\"_secret\",\"type\":\"uint256\"}],\"name\":\"reveal\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_startBlock\",\"type\":\"uint256\"},{\"name\":\"_commitDuration\",\"type\":\"uint256\"},{\"name\":\"_revealDuration\",\"type\":\"uint256\"},{\"name\":\"_revealThreshold\",\"type\":\"uint256\"}],\"name\":\"startCommitReveal\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_cid\",\"type\":\"uint256\"}],\"name\":\"getRandom\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"}],\"name\":\"whitelisted\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_addr\",\"type\":\"address\"}],\"name\":\"addToWhitelist\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_cid\",\"type\":\"uint256\"},{\"name\":\"_secretHash\",\"type\":\"bytes32\"}],\"name\":\"commit\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"cid\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"startBlock\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"commitDuration\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"revealDuration\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"revealThreshold\",\"type\":\"uint256\"}],\"name\":\"LogStartCommitReveal\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"cid\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"from\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"commitment\",\"type\":\"bytes32\"}],\"name\":\"LogCommit\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"cid\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"from\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"secret\",\"type\":\"uint256\"}],\"name\":\"LogReveal\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"cid\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"random\",\"type\":\"uint256\"}],\"name\":\"LogRandom\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"cid\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"commitNum\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"revealNum\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"revealThreshold\",\"type\":\"uint256\"}],\"name\":\"LogRandomFailure\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"previousOwner\",\"type\":\"address\"}],\"name\":\"OwnershipRenounced\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"}]"

// CommitrevealBin is the compiled bytecode used for deploying new contracts.
const CommitrevealBin = `608060405234801561001057600080fd5b50600080546001600160a01b0319163317905560018054906100349082810161003a565b506100b5565b81548183558181111561006657600902816009028360005260206000209182019101610066919061006b565b505050565b6100b291905b808211156100ae57600080825560018201819055600282018190556003820181905560048201819055600582018190556006820155600901610071565b5090565b90565b610c2f806100c46000396000f3fe608060405234801561001057600080fd5b50600436106100b45760003560e01c8063b917b5a511610071578063b917b5a5146101a1578063cd4b6914146101e2578063d936547e146101ff578063e43252d714610225578063f2f038771461024b578063f2fde38b1461026e576100b4565b8063141961bc146100b9578063715018a61461010e5780638ab1d681146101185780638da5cb5b1461013e5780638f32d59b146101625780639348cef71461017e575b600080fd5b6100d6600480360360208110156100cf57600080fd5b5035610294565b604080519788526020880196909652868601949094526060860192909252608085015260a084015260c0830152519081900360e00190f35b6101166102e3565b005b6101166004803603602081101561012e57600080fd5b50356001600160a01b031661033c565b61014661036e565b604080516001600160a01b039092168252519081900360200190f35b61016a61037d565b604080519115158252519081900360200190f35b6101166004803603604081101561019457600080fd5b508035906020013561038e565b6101d0600480360360808110156101b757600080fd5b508035906020810135906040810135906060013561054a565b60408051918252519081900360200190f35b6101d0600480360360208110156101f857600080fd5b5035610764565b61016a6004803603602081101561021557600080fd5b50356001600160a01b03166108d3565b6101166004803603602081101561023b57600080fd5b50356001600160a01b03166108e8565b6101166004803603604081101561026157600080fd5b508035906020013561091d565b6101166004803603602081101561028457600080fd5b50356001600160a01b0316610b4b565b600181815481106102a157fe5b90600052602060002090600902016000915090508060000154908060010154908060020154908060030154908060040154908060050154908060060154905087565b6102eb61037d565b6102f457600080fd5b600080546040516001600160a01b03909116917ff8df31144d9c2f0f6b59d69b8b98abd5459d07f2742c4df920b25aae33c6482091a2600080546001600160a01b0319169055565b61034461037d565b61034d57600080fd5b6001600160a01b03166000908152600260205260409020805460ff19169055565b6000546001600160a01b031690565b6000546001600160a01b0316331490565b8160006001828154811061039e57fe5b90600052602060002090600902019050816000141580156103c6575060018101548154014310155b80156103de5750600281015460018201548254010143105b6104325760408051600160e51b62461bcd02815260206004820152601360248201527f4e6f7420696e2072657665616c20706861736500000000000000000000000000604482015290519081900360640190fd5b60006001858154811061044157fe5b6000918252602080832033845260076009909302019182019052604090912060028101549192509060ff1615801561049b575060018101546040805160208082018990528251808303820181529183019092528051910120145b6104d957604051600160e51b62461bcd02815260040180806020018281038252602d815260200180610bd7602d913960400191505060405180910390fd5b84815560028101805460ff191660019081179091556005830180549091019055600682018054861890556040805187815233602082015280820187905290517f9141bfaedbc77aa7b8d9c989cd81909d95bb1677e556e34cfd45e50e0bea29679181900360600190a1505050505050565b3360009081526002602052604081205460ff166105b15760408051600160e51b62461bcd02815260206004820152601060248201527f4e6f742077686974656c69737465642100000000000000000000000000000000604482015290519081900360640190fd5b600180546040805160e081018252888152602080820189815282840189815260608085018a81526000608080880182815260a0808a0184815260c08b018581528d8f018f559d909452985160098c027fb10e2d527612073b26eecdfd717e6a320cf44b4afac2b0732d9fcbe2b7fa0cf681019190915596517fb10e2d527612073b26eecdfd717e6a320cf44b4afac2b0732d9fcbe2b7fa0cf788015594517fb10e2d527612073b26eecdfd717e6a320cf44b4afac2b0732d9fcbe2b7fa0cf887015591517fb10e2d527612073b26eecdfd717e6a320cf44b4afac2b0732d9fcbe2b7fa0cf986015592517fb10e2d527612073b26eecdfd717e6a320cf44b4afac2b0732d9fcbe2b7fa0cfa85015591517fb10e2d527612073b26eecdfd717e6a320cf44b4afac2b0732d9fcbe2b7fa0cfb84015596517fb10e2d527612073b26eecdfd717e6a320cf44b4afac2b0732d9fcbe2b7fa0cfc9092019190915583518581529182018b90528184018a9052948101889052938401869052905191927fbbfccb30e8cf1b5d88802741ceb4d63cf854fa8931eeeaec6b700f31f429dc09929081900390910190a195945050505050565b60008160006001828154811061077657fe5b90600052602060002090600902019050816000141580156107a4575060028101546001820154825401014310155b6107f85760408051600160e51b62461bcd02815260206004820152601e60248201527f436f6d6d69742052657665616c206e6f742066696e6973686564207965740000604482015290519081900360640190fd5b60006001858154811061080757fe5b90600052602060002090600902019050806003015481600501541061087157600681015460408051878152602081019290925280517fa34f42a90fadfe357ee14419d618438a057569fbb63bab0c5fbcca0fc2b11e8d9281900390910190a16006015492506108cc565b600481015460058201546003830154604080518981526020810194909452838101929092526060830152517fe888e7582d0505bce81eef694dfa216179eaaa3c1bd96b7894de8b4370d8543e9181900360800190a160009350505b5050919050565b60026020526000908152604090205460ff1681565b6108f061037d565b6108f957600080fd5b6001600160a01b03166000908152600260205260409020805460ff19166001179055565b818160006001838154811061092e57fe5b9060005260206000209060090201905082600014158015610950575080544310155b80156109625750600181015481540143105b6109b65760408051600160e51b62461bcd02815260206004820152601360248201527f4e6f7420696e20636f6d6d697420706861736500000000000000000000000000604482015290519081900360640190fd5b81610a0b5760408051600160e51b62461bcd02815260206004820152601060248201527f456d70747920636f6d6d69746d656e7400000000000000000000000000000000604482015290519081900360640190fd5b600082815260088201602052604090205460ff1615610a745760408051600160e51b62461bcd02815260206004820152601560248201527f4475706c69636174656420636f6d6d69746d656e740000000000000000000000604482015290519081900360640190fd5b600060018681548110610a8357fe5b60009182526020808320888452600860099093020191820181526040808420805460ff1990811660019081179092558251606080820185528782528186018d815282860189815233808b5260078a01895299879020935184559051838601555160029092018054921515929093169190911790915560048501805490920190915581518b815292830194909452818101899052519193507f918c00c65dd2a8dee4c6985d1d67f04aa8cd2c93e8d427d398a90444c7f7c75e92908290030190a1505050505050565b610b5361037d565b610b5c57600080fd5b610b6581610b68565b50565b6001600160a01b038116610b7b57600080fd5b600080546040516001600160a01b03808516939216917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e091a3600080546001600160a01b0319166001600160a01b039290921691909117905556fe52657665616c65642073656372657420646f65736e2774206d61746368207769746820636f6d6d69746d656e74a165627a7a723058201cf942c3f6e796b43fa7ec1aa6a54995353c20b27795cc5144ccc78618eea79f0029`

// DeployCommitreveal deploys a new Ethereum contract, binding an instance of Commitreveal to it.
func DeployCommitreveal(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Commitreveal, error) {
	parsed, err := abi.JSON(strings.NewReader(CommitrevealABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(CommitrevealBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Commitreveal{CommitrevealCaller: CommitrevealCaller{contract: contract}, CommitrevealTransactor: CommitrevealTransactor{contract: contract}, CommitrevealFilterer: CommitrevealFilterer{contract: contract}}, nil
}

// Commitreveal is an auto generated Go binding around an Ethereum contract.
type Commitreveal struct {
	CommitrevealCaller     // Read-only binding to the contract
	CommitrevealTransactor // Write-only binding to the contract
	CommitrevealFilterer   // Log filterer for contract events
}

// CommitrevealCaller is an auto generated read-only Go binding around an Ethereum contract.
type CommitrevealCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CommitrevealTransactor is an auto generated write-only Go binding around an Ethereum contract.
type CommitrevealTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CommitrevealFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type CommitrevealFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CommitrevealSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type CommitrevealSession struct {
	Contract     *Commitreveal     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// CommitrevealCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type CommitrevealCallerSession struct {
	Contract *CommitrevealCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// CommitrevealTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type CommitrevealTransactorSession struct {
	Contract     *CommitrevealTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// CommitrevealRaw is an auto generated low-level Go binding around an Ethereum contract.
type CommitrevealRaw struct {
	Contract *Commitreveal // Generic contract binding to access the raw methods on
}

// CommitrevealCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type CommitrevealCallerRaw struct {
	Contract *CommitrevealCaller // Generic read-only contract binding to access the raw methods on
}

// CommitrevealTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type CommitrevealTransactorRaw struct {
	Contract *CommitrevealTransactor // Generic write-only contract binding to access the raw methods on
}

// NewCommitreveal creates a new instance of Commitreveal, bound to a specific deployed contract.
func NewCommitreveal(address common.Address, backend bind.ContractBackend) (*Commitreveal, error) {
	contract, err := bindCommitreveal(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Commitreveal{CommitrevealCaller: CommitrevealCaller{contract: contract}, CommitrevealTransactor: CommitrevealTransactor{contract: contract}, CommitrevealFilterer: CommitrevealFilterer{contract: contract}}, nil
}

// NewCommitrevealCaller creates a new read-only instance of Commitreveal, bound to a specific deployed contract.
func NewCommitrevealCaller(address common.Address, caller bind.ContractCaller) (*CommitrevealCaller, error) {
	contract, err := bindCommitreveal(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &CommitrevealCaller{contract: contract}, nil
}

// NewCommitrevealTransactor creates a new write-only instance of Commitreveal, bound to a specific deployed contract.
func NewCommitrevealTransactor(address common.Address, transactor bind.ContractTransactor) (*CommitrevealTransactor, error) {
	contract, err := bindCommitreveal(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &CommitrevealTransactor{contract: contract}, nil
}

// NewCommitrevealFilterer creates a new log filterer instance of Commitreveal, bound to a specific deployed contract.
func NewCommitrevealFilterer(address common.Address, filterer bind.ContractFilterer) (*CommitrevealFilterer, error) {
	contract, err := bindCommitreveal(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &CommitrevealFilterer{contract: contract}, nil
}

// bindCommitreveal binds a generic wrapper to an already deployed contract.
func bindCommitreveal(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(CommitrevealABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Commitreveal *CommitrevealRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Commitreveal.Contract.CommitrevealCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Commitreveal *CommitrevealRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Commitreveal.Contract.CommitrevealTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Commitreveal *CommitrevealRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Commitreveal.Contract.CommitrevealTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Commitreveal *CommitrevealCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Commitreveal.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Commitreveal *CommitrevealTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Commitreveal.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Commitreveal *CommitrevealTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Commitreveal.Contract.contract.Transact(opts, method, params...)
}

// Campaigns is a free data retrieval call binding the contract method 0x141961bc.
//
// Solidity: function campaigns(uint256 ) constant returns(uint256 startBlock, uint256 commitDuration, uint256 revealDuration, uint256 revealThreshold, uint256 commitNum, uint256 revealNum, uint256 generatedRandom)
func (_Commitreveal *CommitrevealCaller) Campaigns(opts *bind.CallOpts, arg0 *big.Int) (struct {
	StartBlock      *big.Int
	CommitDuration  *big.Int
	RevealDuration  *big.Int
	RevealThreshold *big.Int
	CommitNum       *big.Int
	RevealNum       *big.Int
	GeneratedRandom *big.Int
}, error) {
	ret := new(struct {
		StartBlock      *big.Int
		CommitDuration  *big.Int
		RevealDuration  *big.Int
		RevealThreshold *big.Int
		CommitNum       *big.Int
		RevealNum       *big.Int
		GeneratedRandom *big.Int
	})
	out := ret
	err := _Commitreveal.contract.Call(opts, out, "campaigns", arg0)
	return *ret, err
}

// Campaigns is a free data retrieval call binding the contract method 0x141961bc.
//
// Solidity: function campaigns(uint256 ) constant returns(uint256 startBlock, uint256 commitDuration, uint256 revealDuration, uint256 revealThreshold, uint256 commitNum, uint256 revealNum, uint256 generatedRandom)
func (_Commitreveal *CommitrevealSession) Campaigns(arg0 *big.Int) (struct {
	StartBlock      *big.Int
	CommitDuration  *big.Int
	RevealDuration  *big.Int
	RevealThreshold *big.Int
	CommitNum       *big.Int
	RevealNum       *big.Int
	GeneratedRandom *big.Int
}, error) {
	return _Commitreveal.Contract.Campaigns(&_Commitreveal.CallOpts, arg0)
}

// Campaigns is a free data retrieval call binding the contract method 0x141961bc.
//
// Solidity: function campaigns(uint256 ) constant returns(uint256 startBlock, uint256 commitDuration, uint256 revealDuration, uint256 revealThreshold, uint256 commitNum, uint256 revealNum, uint256 generatedRandom)
func (_Commitreveal *CommitrevealCallerSession) Campaigns(arg0 *big.Int) (struct {
	StartBlock      *big.Int
	CommitDuration  *big.Int
	RevealDuration  *big.Int
	RevealThreshold *big.Int
	CommitNum       *big.Int
	RevealNum       *big.Int
	GeneratedRandom *big.Int
}, error) {
	return _Commitreveal.Contract.Campaigns(&_Commitreveal.CallOpts, arg0)
}

// IsOwner is a free data retrieval call binding the contract method 0x8f32d59b.
//
// Solidity: function isOwner() constant returns(bool)
func (_Commitreveal *CommitrevealCaller) IsOwner(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Commitreveal.contract.Call(opts, out, "isOwner")
	return *ret0, err
}

// IsOwner is a free data retrieval call binding the contract method 0x8f32d59b.
//
// Solidity: function isOwner() constant returns(bool)
func (_Commitreveal *CommitrevealSession) IsOwner() (bool, error) {
	return _Commitreveal.Contract.IsOwner(&_Commitreveal.CallOpts)
}

// IsOwner is a free data retrieval call binding the contract method 0x8f32d59b.
//
// Solidity: function isOwner() constant returns(bool)
func (_Commitreveal *CommitrevealCallerSession) IsOwner() (bool, error) {
	return _Commitreveal.Contract.IsOwner(&_Commitreveal.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Commitreveal *CommitrevealCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Commitreveal.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Commitreveal *CommitrevealSession) Owner() (common.Address, error) {
	return _Commitreveal.Contract.Owner(&_Commitreveal.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Commitreveal *CommitrevealCallerSession) Owner() (common.Address, error) {
	return _Commitreveal.Contract.Owner(&_Commitreveal.CallOpts)
}

// Whitelisted is a free data retrieval call binding the contract method 0xd936547e.
//
// Solidity: function whitelisted(address ) constant returns(bool)
func (_Commitreveal *CommitrevealCaller) Whitelisted(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Commitreveal.contract.Call(opts, out, "whitelisted", arg0)
	return *ret0, err
}

// Whitelisted is a free data retrieval call binding the contract method 0xd936547e.
//
// Solidity: function whitelisted(address ) constant returns(bool)
func (_Commitreveal *CommitrevealSession) Whitelisted(arg0 common.Address) (bool, error) {
	return _Commitreveal.Contract.Whitelisted(&_Commitreveal.CallOpts, arg0)
}

// Whitelisted is a free data retrieval call binding the contract method 0xd936547e.
//
// Solidity: function whitelisted(address ) constant returns(bool)
func (_Commitreveal *CommitrevealCallerSession) Whitelisted(arg0 common.Address) (bool, error) {
	return _Commitreveal.Contract.Whitelisted(&_Commitreveal.CallOpts, arg0)
}

// AddToWhitelist is a paid mutator transaction binding the contract method 0xe43252d7.
//
// Solidity: function addToWhitelist(address _addr) returns()
func (_Commitreveal *CommitrevealTransactor) AddToWhitelist(opts *bind.TransactOpts, _addr common.Address) (*types.Transaction, error) {
	return _Commitreveal.contract.Transact(opts, "addToWhitelist", _addr)
}

// AddToWhitelist is a paid mutator transaction binding the contract method 0xe43252d7.
//
// Solidity: function addToWhitelist(address _addr) returns()
func (_Commitreveal *CommitrevealSession) AddToWhitelist(_addr common.Address) (*types.Transaction, error) {
	return _Commitreveal.Contract.AddToWhitelist(&_Commitreveal.TransactOpts, _addr)
}

// AddToWhitelist is a paid mutator transaction binding the contract method 0xe43252d7.
//
// Solidity: function addToWhitelist(address _addr) returns()
func (_Commitreveal *CommitrevealTransactorSession) AddToWhitelist(_addr common.Address) (*types.Transaction, error) {
	return _Commitreveal.Contract.AddToWhitelist(&_Commitreveal.TransactOpts, _addr)
}

// Commit is a paid mutator transaction binding the contract method 0xf2f03877.
//
// Solidity: function commit(uint256 _cid, bytes32 _secretHash) returns()
func (_Commitreveal *CommitrevealTransactor) Commit(opts *bind.TransactOpts, _cid *big.Int, _secretHash [32]byte) (*types.Transaction, error) {
	return _Commitreveal.contract.Transact(opts, "commit", _cid, _secretHash)
}

// Commit is a paid mutator transaction binding the contract method 0xf2f03877.
//
// Solidity: function commit(uint256 _cid, bytes32 _secretHash) returns()
func (_Commitreveal *CommitrevealSession) Commit(_cid *big.Int, _secretHash [32]byte) (*types.Transaction, error) {
	return _Commitreveal.Contract.Commit(&_Commitreveal.TransactOpts, _cid, _secretHash)
}

// Commit is a paid mutator transaction binding the contract method 0xf2f03877.
//
// Solidity: function commit(uint256 _cid, bytes32 _secretHash) returns()
func (_Commitreveal *CommitrevealTransactorSession) Commit(_cid *big.Int, _secretHash [32]byte) (*types.Transaction, error) {
	return _Commitreveal.Contract.Commit(&_Commitreveal.TransactOpts, _cid, _secretHash)
}

// GetRandom is a paid mutator transaction binding the contract method 0xcd4b6914.
//
// Solidity: function getRandom(uint256 _cid) returns(uint256)
func (_Commitreveal *CommitrevealTransactor) GetRandom(opts *bind.TransactOpts, _cid *big.Int) (*types.Transaction, error) {
	return _Commitreveal.contract.Transact(opts, "getRandom", _cid)
}

// GetRandom is a paid mutator transaction binding the contract method 0xcd4b6914.
//
// Solidity: function getRandom(uint256 _cid) returns(uint256)
func (_Commitreveal *CommitrevealSession) GetRandom(_cid *big.Int) (*types.Transaction, error) {
	return _Commitreveal.Contract.GetRandom(&_Commitreveal.TransactOpts, _cid)
}

// GetRandom is a paid mutator transaction binding the contract method 0xcd4b6914.
//
// Solidity: function getRandom(uint256 _cid) returns(uint256)
func (_Commitreveal *CommitrevealTransactorSession) GetRandom(_cid *big.Int) (*types.Transaction, error) {
	return _Commitreveal.Contract.GetRandom(&_Commitreveal.TransactOpts, _cid)
}

// RemoveFromWhitelist is a paid mutator transaction binding the contract method 0x8ab1d681.
//
// Solidity: function removeFromWhitelist(address _addr) returns()
func (_Commitreveal *CommitrevealTransactor) RemoveFromWhitelist(opts *bind.TransactOpts, _addr common.Address) (*types.Transaction, error) {
	return _Commitreveal.contract.Transact(opts, "removeFromWhitelist", _addr)
}

// RemoveFromWhitelist is a paid mutator transaction binding the contract method 0x8ab1d681.
//
// Solidity: function removeFromWhitelist(address _addr) returns()
func (_Commitreveal *CommitrevealSession) RemoveFromWhitelist(_addr common.Address) (*types.Transaction, error) {
	return _Commitreveal.Contract.RemoveFromWhitelist(&_Commitreveal.TransactOpts, _addr)
}

// RemoveFromWhitelist is a paid mutator transaction binding the contract method 0x8ab1d681.
//
// Solidity: function removeFromWhitelist(address _addr) returns()
func (_Commitreveal *CommitrevealTransactorSession) RemoveFromWhitelist(_addr common.Address) (*types.Transaction, error) {
	return _Commitreveal.Contract.RemoveFromWhitelist(&_Commitreveal.TransactOpts, _addr)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Commitreveal *CommitrevealTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Commitreveal.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Commitreveal *CommitrevealSession) RenounceOwnership() (*types.Transaction, error) {
	return _Commitreveal.Contract.RenounceOwnership(&_Commitreveal.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Commitreveal *CommitrevealTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Commitreveal.Contract.RenounceOwnership(&_Commitreveal.TransactOpts)
}

// Reveal is a paid mutator transaction binding the contract method 0x9348cef7.
//
// Solidity: function reveal(uint256 _cid, uint256 _secret) returns()
func (_Commitreveal *CommitrevealTransactor) Reveal(opts *bind.TransactOpts, _cid *big.Int, _secret *big.Int) (*types.Transaction, error) {
	return _Commitreveal.contract.Transact(opts, "reveal", _cid, _secret)
}

// Reveal is a paid mutator transaction binding the contract method 0x9348cef7.
//
// Solidity: function reveal(uint256 _cid, uint256 _secret) returns()
func (_Commitreveal *CommitrevealSession) Reveal(_cid *big.Int, _secret *big.Int) (*types.Transaction, error) {
	return _Commitreveal.Contract.Reveal(&_Commitreveal.TransactOpts, _cid, _secret)
}

// Reveal is a paid mutator transaction binding the contract method 0x9348cef7.
//
// Solidity: function reveal(uint256 _cid, uint256 _secret) returns()
func (_Commitreveal *CommitrevealTransactorSession) Reveal(_cid *big.Int, _secret *big.Int) (*types.Transaction, error) {
	return _Commitreveal.Contract.Reveal(&_Commitreveal.TransactOpts, _cid, _secret)
}

// StartCommitReveal is a paid mutator transaction binding the contract method 0xb917b5a5.
//
// Solidity: function startCommitReveal(uint256 _startBlock, uint256 _commitDuration, uint256 _revealDuration, uint256 _revealThreshold) returns(uint256)
func (_Commitreveal *CommitrevealTransactor) StartCommitReveal(opts *bind.TransactOpts, _startBlock *big.Int, _commitDuration *big.Int, _revealDuration *big.Int, _revealThreshold *big.Int) (*types.Transaction, error) {
	return _Commitreveal.contract.Transact(opts, "startCommitReveal", _startBlock, _commitDuration, _revealDuration, _revealThreshold)
}

// StartCommitReveal is a paid mutator transaction binding the contract method 0xb917b5a5.
//
// Solidity: function startCommitReveal(uint256 _startBlock, uint256 _commitDuration, uint256 _revealDuration, uint256 _revealThreshold) returns(uint256)
func (_Commitreveal *CommitrevealSession) StartCommitReveal(_startBlock *big.Int, _commitDuration *big.Int, _revealDuration *big.Int, _revealThreshold *big.Int) (*types.Transaction, error) {
	return _Commitreveal.Contract.StartCommitReveal(&_Commitreveal.TransactOpts, _startBlock, _commitDuration, _revealDuration, _revealThreshold)
}

// StartCommitReveal is a paid mutator transaction binding the contract method 0xb917b5a5.
//
// Solidity: function startCommitReveal(uint256 _startBlock, uint256 _commitDuration, uint256 _revealDuration, uint256 _revealThreshold) returns(uint256)
func (_Commitreveal *CommitrevealTransactorSession) StartCommitReveal(_startBlock *big.Int, _commitDuration *big.Int, _revealDuration *big.Int, _revealThreshold *big.Int) (*types.Transaction, error) {
	return _Commitreveal.Contract.StartCommitReveal(&_Commitreveal.TransactOpts, _startBlock, _commitDuration, _revealDuration, _revealThreshold)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Commitreveal *CommitrevealTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Commitreveal.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Commitreveal *CommitrevealSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Commitreveal.Contract.TransferOwnership(&_Commitreveal.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Commitreveal *CommitrevealTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Commitreveal.Contract.TransferOwnership(&_Commitreveal.TransactOpts, newOwner)
}

// CommitrevealLogCommitIterator is returned from FilterLogCommit and is used to iterate over the raw logs and unpacked data for LogCommit events raised by the Commitreveal contract.
type CommitrevealLogCommitIterator struct {
	Event *CommitrevealLogCommit // Event containing the contract specifics and raw log

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
func (it *CommitrevealLogCommitIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CommitrevealLogCommit)
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
		it.Event = new(CommitrevealLogCommit)
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
func (it *CommitrevealLogCommitIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CommitrevealLogCommitIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CommitrevealLogCommit represents a LogCommit event raised by the Commitreveal contract.
type CommitrevealLogCommit struct {
	Cid        *big.Int
	From       common.Address
	Commitment [32]byte
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterLogCommit is a free log retrieval operation binding the contract event 0x918c00c65dd2a8dee4c6985d1d67f04aa8cd2c93e8d427d398a90444c7f7c75e.
//
// Solidity: event LogCommit(uint256 cid, address from, bytes32 commitment)
func (_Commitreveal *CommitrevealFilterer) FilterLogCommit(opts *bind.FilterOpts) (*CommitrevealLogCommitIterator, error) {

	logs, sub, err := _Commitreveal.contract.FilterLogs(opts, "LogCommit")
	if err != nil {
		return nil, err
	}
	return &CommitrevealLogCommitIterator{contract: _Commitreveal.contract, event: "LogCommit", logs: logs, sub: sub}, nil
}

// WatchLogCommit is a free log subscription operation binding the contract event 0x918c00c65dd2a8dee4c6985d1d67f04aa8cd2c93e8d427d398a90444c7f7c75e.
//
// Solidity: event LogCommit(uint256 cid, address from, bytes32 commitment)
func (_Commitreveal *CommitrevealFilterer) WatchLogCommit(opts *bind.WatchOpts, sink chan<- *CommitrevealLogCommit) (event.Subscription, error) {

	logs, sub, err := _Commitreveal.contract.WatchLogs(opts, "LogCommit")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CommitrevealLogCommit)
				if err := _Commitreveal.contract.UnpackLog(event, "LogCommit", log); err != nil {
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

// CommitrevealLogRandomIterator is returned from FilterLogRandom and is used to iterate over the raw logs and unpacked data for LogRandom events raised by the Commitreveal contract.
type CommitrevealLogRandomIterator struct {
	Event *CommitrevealLogRandom // Event containing the contract specifics and raw log

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
func (it *CommitrevealLogRandomIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CommitrevealLogRandom)
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
		it.Event = new(CommitrevealLogRandom)
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
func (it *CommitrevealLogRandomIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CommitrevealLogRandomIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CommitrevealLogRandom represents a LogRandom event raised by the Commitreveal contract.
type CommitrevealLogRandom struct {
	Cid    *big.Int
	Random *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterLogRandom is a free log retrieval operation binding the contract event 0xa34f42a90fadfe357ee14419d618438a057569fbb63bab0c5fbcca0fc2b11e8d.
//
// Solidity: event LogRandom(uint256 cid, uint256 random)
func (_Commitreveal *CommitrevealFilterer) FilterLogRandom(opts *bind.FilterOpts) (*CommitrevealLogRandomIterator, error) {

	logs, sub, err := _Commitreveal.contract.FilterLogs(opts, "LogRandom")
	if err != nil {
		return nil, err
	}
	return &CommitrevealLogRandomIterator{contract: _Commitreveal.contract, event: "LogRandom", logs: logs, sub: sub}, nil
}

// WatchLogRandom is a free log subscription operation binding the contract event 0xa34f42a90fadfe357ee14419d618438a057569fbb63bab0c5fbcca0fc2b11e8d.
//
// Solidity: event LogRandom(uint256 cid, uint256 random)
func (_Commitreveal *CommitrevealFilterer) WatchLogRandom(opts *bind.WatchOpts, sink chan<- *CommitrevealLogRandom) (event.Subscription, error) {

	logs, sub, err := _Commitreveal.contract.WatchLogs(opts, "LogRandom")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CommitrevealLogRandom)
				if err := _Commitreveal.contract.UnpackLog(event, "LogRandom", log); err != nil {
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

// CommitrevealLogRandomFailureIterator is returned from FilterLogRandomFailure and is used to iterate over the raw logs and unpacked data for LogRandomFailure events raised by the Commitreveal contract.
type CommitrevealLogRandomFailureIterator struct {
	Event *CommitrevealLogRandomFailure // Event containing the contract specifics and raw log

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
func (it *CommitrevealLogRandomFailureIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CommitrevealLogRandomFailure)
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
		it.Event = new(CommitrevealLogRandomFailure)
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
func (it *CommitrevealLogRandomFailureIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CommitrevealLogRandomFailureIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CommitrevealLogRandomFailure represents a LogRandomFailure event raised by the Commitreveal contract.
type CommitrevealLogRandomFailure struct {
	Cid             *big.Int
	CommitNum       *big.Int
	RevealNum       *big.Int
	RevealThreshold *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterLogRandomFailure is a free log retrieval operation binding the contract event 0xe888e7582d0505bce81eef694dfa216179eaaa3c1bd96b7894de8b4370d8543e.
//
// Solidity: event LogRandomFailure(uint256 cid, uint256 commitNum, uint256 revealNum, uint256 revealThreshold)
func (_Commitreveal *CommitrevealFilterer) FilterLogRandomFailure(opts *bind.FilterOpts) (*CommitrevealLogRandomFailureIterator, error) {

	logs, sub, err := _Commitreveal.contract.FilterLogs(opts, "LogRandomFailure")
	if err != nil {
		return nil, err
	}
	return &CommitrevealLogRandomFailureIterator{contract: _Commitreveal.contract, event: "LogRandomFailure", logs: logs, sub: sub}, nil
}

// WatchLogRandomFailure is a free log subscription operation binding the contract event 0xe888e7582d0505bce81eef694dfa216179eaaa3c1bd96b7894de8b4370d8543e.
//
// Solidity: event LogRandomFailure(uint256 cid, uint256 commitNum, uint256 revealNum, uint256 revealThreshold)
func (_Commitreveal *CommitrevealFilterer) WatchLogRandomFailure(opts *bind.WatchOpts, sink chan<- *CommitrevealLogRandomFailure) (event.Subscription, error) {

	logs, sub, err := _Commitreveal.contract.WatchLogs(opts, "LogRandomFailure")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CommitrevealLogRandomFailure)
				if err := _Commitreveal.contract.UnpackLog(event, "LogRandomFailure", log); err != nil {
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

// CommitrevealLogRevealIterator is returned from FilterLogReveal and is used to iterate over the raw logs and unpacked data for LogReveal events raised by the Commitreveal contract.
type CommitrevealLogRevealIterator struct {
	Event *CommitrevealLogReveal // Event containing the contract specifics and raw log

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
func (it *CommitrevealLogRevealIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CommitrevealLogReveal)
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
		it.Event = new(CommitrevealLogReveal)
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
func (it *CommitrevealLogRevealIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CommitrevealLogRevealIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CommitrevealLogReveal represents a LogReveal event raised by the Commitreveal contract.
type CommitrevealLogReveal struct {
	Cid    *big.Int
	From   common.Address
	Secret *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterLogReveal is a free log retrieval operation binding the contract event 0x9141bfaedbc77aa7b8d9c989cd81909d95bb1677e556e34cfd45e50e0bea2967.
//
// Solidity: event LogReveal(uint256 cid, address from, uint256 secret)
func (_Commitreveal *CommitrevealFilterer) FilterLogReveal(opts *bind.FilterOpts) (*CommitrevealLogRevealIterator, error) {

	logs, sub, err := _Commitreveal.contract.FilterLogs(opts, "LogReveal")
	if err != nil {
		return nil, err
	}
	return &CommitrevealLogRevealIterator{contract: _Commitreveal.contract, event: "LogReveal", logs: logs, sub: sub}, nil
}

// WatchLogReveal is a free log subscription operation binding the contract event 0x9141bfaedbc77aa7b8d9c989cd81909d95bb1677e556e34cfd45e50e0bea2967.
//
// Solidity: event LogReveal(uint256 cid, address from, uint256 secret)
func (_Commitreveal *CommitrevealFilterer) WatchLogReveal(opts *bind.WatchOpts, sink chan<- *CommitrevealLogReveal) (event.Subscription, error) {

	logs, sub, err := _Commitreveal.contract.WatchLogs(opts, "LogReveal")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CommitrevealLogReveal)
				if err := _Commitreveal.contract.UnpackLog(event, "LogReveal", log); err != nil {
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

// CommitrevealLogStartCommitRevealIterator is returned from FilterLogStartCommitReveal and is used to iterate over the raw logs and unpacked data for LogStartCommitReveal events raised by the Commitreveal contract.
type CommitrevealLogStartCommitRevealIterator struct {
	Event *CommitrevealLogStartCommitReveal // Event containing the contract specifics and raw log

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
func (it *CommitrevealLogStartCommitRevealIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CommitrevealLogStartCommitReveal)
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
		it.Event = new(CommitrevealLogStartCommitReveal)
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
func (it *CommitrevealLogStartCommitRevealIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CommitrevealLogStartCommitRevealIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CommitrevealLogStartCommitReveal represents a LogStartCommitReveal event raised by the Commitreveal contract.
type CommitrevealLogStartCommitReveal struct {
	Cid             *big.Int
	StartBlock      *big.Int
	CommitDuration  *big.Int
	RevealDuration  *big.Int
	RevealThreshold *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterLogStartCommitReveal is a free log retrieval operation binding the contract event 0xbbfccb30e8cf1b5d88802741ceb4d63cf854fa8931eeeaec6b700f31f429dc09.
//
// Solidity: event LogStartCommitReveal(uint256 cid, uint256 startBlock, uint256 commitDuration, uint256 revealDuration, uint256 revealThreshold)
func (_Commitreveal *CommitrevealFilterer) FilterLogStartCommitReveal(opts *bind.FilterOpts) (*CommitrevealLogStartCommitRevealIterator, error) {

	logs, sub, err := _Commitreveal.contract.FilterLogs(opts, "LogStartCommitReveal")
	if err != nil {
		return nil, err
	}
	return &CommitrevealLogStartCommitRevealIterator{contract: _Commitreveal.contract, event: "LogStartCommitReveal", logs: logs, sub: sub}, nil
}

// WatchLogStartCommitReveal is a free log subscription operation binding the contract event 0xbbfccb30e8cf1b5d88802741ceb4d63cf854fa8931eeeaec6b700f31f429dc09.
//
// Solidity: event LogStartCommitReveal(uint256 cid, uint256 startBlock, uint256 commitDuration, uint256 revealDuration, uint256 revealThreshold)
func (_Commitreveal *CommitrevealFilterer) WatchLogStartCommitReveal(opts *bind.WatchOpts, sink chan<- *CommitrevealLogStartCommitReveal) (event.Subscription, error) {

	logs, sub, err := _Commitreveal.contract.WatchLogs(opts, "LogStartCommitReveal")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CommitrevealLogStartCommitReveal)
				if err := _Commitreveal.contract.UnpackLog(event, "LogStartCommitReveal", log); err != nil {
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

// CommitrevealOwnershipRenouncedIterator is returned from FilterOwnershipRenounced and is used to iterate over the raw logs and unpacked data for OwnershipRenounced events raised by the Commitreveal contract.
type CommitrevealOwnershipRenouncedIterator struct {
	Event *CommitrevealOwnershipRenounced // Event containing the contract specifics and raw log

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
func (it *CommitrevealOwnershipRenouncedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CommitrevealOwnershipRenounced)
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
		it.Event = new(CommitrevealOwnershipRenounced)
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
func (it *CommitrevealOwnershipRenouncedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CommitrevealOwnershipRenouncedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CommitrevealOwnershipRenounced represents a OwnershipRenounced event raised by the Commitreveal contract.
type CommitrevealOwnershipRenounced struct {
	PreviousOwner common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipRenounced is a free log retrieval operation binding the contract event 0xf8df31144d9c2f0f6b59d69b8b98abd5459d07f2742c4df920b25aae33c64820.
//
// Solidity: event OwnershipRenounced(address indexed previousOwner)
func (_Commitreveal *CommitrevealFilterer) FilterOwnershipRenounced(opts *bind.FilterOpts, previousOwner []common.Address) (*CommitrevealOwnershipRenouncedIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}

	logs, sub, err := _Commitreveal.contract.FilterLogs(opts, "OwnershipRenounced", previousOwnerRule)
	if err != nil {
		return nil, err
	}
	return &CommitrevealOwnershipRenouncedIterator{contract: _Commitreveal.contract, event: "OwnershipRenounced", logs: logs, sub: sub}, nil
}

// WatchOwnershipRenounced is a free log subscription operation binding the contract event 0xf8df31144d9c2f0f6b59d69b8b98abd5459d07f2742c4df920b25aae33c64820.
//
// Solidity: event OwnershipRenounced(address indexed previousOwner)
func (_Commitreveal *CommitrevealFilterer) WatchOwnershipRenounced(opts *bind.WatchOpts, sink chan<- *CommitrevealOwnershipRenounced, previousOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}

	logs, sub, err := _Commitreveal.contract.WatchLogs(opts, "OwnershipRenounced", previousOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CommitrevealOwnershipRenounced)
				if err := _Commitreveal.contract.UnpackLog(event, "OwnershipRenounced", log); err != nil {
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

// CommitrevealOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Commitreveal contract.
type CommitrevealOwnershipTransferredIterator struct {
	Event *CommitrevealOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *CommitrevealOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CommitrevealOwnershipTransferred)
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
		it.Event = new(CommitrevealOwnershipTransferred)
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
func (it *CommitrevealOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CommitrevealOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CommitrevealOwnershipTransferred represents a OwnershipTransferred event raised by the Commitreveal contract.
type CommitrevealOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Commitreveal *CommitrevealFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*CommitrevealOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Commitreveal.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &CommitrevealOwnershipTransferredIterator{contract: _Commitreveal.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Commitreveal *CommitrevealFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *CommitrevealOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Commitreveal.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CommitrevealOwnershipTransferred)
				if err := _Commitreveal.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
