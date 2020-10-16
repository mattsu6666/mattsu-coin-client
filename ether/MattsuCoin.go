// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package ether

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
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// MattsuCoinABI is the input ABI used to generate the binding from.
const MattsuCoinABI = "[{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"initialSupply\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"admin\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"buyer\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"int256\",\"name\":\"ethUsd\",\"type\":\"int256\"}],\"name\":\"BuyEthUsd\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"previousAdminRole\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"newAdminRole\",\"type\":\"bytes32\"}],\"name\":\"RoleAdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleGranted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleRevoked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\",\"constant\":true},{\"inputs\":[],\"name\":\"DEFAULT_ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\",\"constant\":true},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\",\"constant\":true},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\",\"constant\":true},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\",\"constant\":true},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"subtractedValue\",\"type\":\"uint256\"}],\"name\":\"decreaseAllowance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleAdmin\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\",\"constant\":true},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"getRoleMember\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\",\"constant\":true},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleMemberCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\",\"constant\":true},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"grantRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"hasRole\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\",\"constant\":true},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"addedValue\",\"type\":\"uint256\"}],\"name\":\"increaseAllowance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\",\"constant\":true},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"renounceRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"revokeRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\",\"constant\":true},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\",\"constant\":true},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"mint\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"burn\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"buyLatestEthUsd\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// MattsuCoinBin is the compiled bytecode used for deploying new contracts.
var MattsuCoinBin = "0x60806040523480156200001157600080fd5b50604051620028b3380380620028b3833981810160405260408110156200003757600080fd5b8101908080519060200190929190805190602001909291905050506040518060400160405280600681526020017f4d617474737500000000000000000000000000000000000000000000000000008152506040518060400160405280600381526020017f4d415400000000000000000000000000000000000000000000000000000000008152508160039080519060200190620000d6929190620005ce565b508060049080519060200190620000ef929190620005ce565b506012600560006101000a81548160ff021916908360ff1602179055505050620001206003620001ad60201b60201c565b620001323383620001cb60201b60201c565b620001647fa49807205ce4d355092ef5a8a18f56e8913cf4a201fbe287825b095693c2177582620003a960201b60201c565b80600760006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550505062000674565b80600560006101000a81548160ff021916908360ff16021790555050565b600073ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff1614156200026f576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252601f8152602001807f45524332303a206d696e7420746f20746865207a65726f20616464726573730081525060200191505060405180910390fd5b6200028360008383620003bf60201b60201c565b6200029f81600254620003c460201b6200122d1790919060201c565b600281905550620002fd816000808573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002054620003c460201b6200122d1790919060201c565b6000808473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020819055508173ffffffffffffffffffffffffffffffffffffffff16600073ffffffffffffffffffffffffffffffffffffffff167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef836040518082815260200191505060405180910390a35050565b620003bb82826200044d60201b60201c565b5050565b505050565b60008082840190508381101562000443576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252601b8152602001807f536166654d6174683a206164646974696f6e206f766572666c6f77000000000081525060200191505060405180910390fd5b8091505092915050565b6200047c8160066000858152602001908152602001600020600001620004f160201b620012b51790919060201c565b15620004ed57620004926200052960201b60201c565b73ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff16837f2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d60405160405180910390a45b5050565b600062000521836000018373ffffffffffffffffffffffffffffffffffffffff1660001b6200053160201b60201c565b905092915050565b600033905090565b6000620005458383620005ab60201b60201c565b620005a0578260000182908060018154018082558091505060019003906000526020600020016000909190919091505582600001805490508360010160008481526020019081526020016000208190555060019050620005a5565b600090505b92915050565b600080836001016000848152602001908152602001600020541415905092915050565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f106200061157805160ff191683800117855562000642565b8280016001018555821562000642579182015b828111156200064157825182559160200191906001019062000624565b5b50905062000651919062000655565b5090565b5b808211156200067057600081600090555060010162000656565b5090565b61222f80620006846000396000f3fe608060405234801561001057600080fd5b506004361061014d5760003560e01c806375b238fc116100c3578063a457c2d71161007c578063a457c2d7146106b7578063a9059cbb1461071b578063ca15c8731461077f578063d547741f146107c1578063dd62ed3e1461080f578063fc3feb53146108875761014d565b806375b238fc146104e45780639010d07c1461050257806391d148541461056457806395d89b41146105c85780639dc29fac1461064b578063a217fddf146106995761014d565b80632f2ff15d116101155780632f2ff15d1461031d578063313ce5671461036b57806336568abe1461038c57806339509351146103da57806340c10f191461043e57806370a082311461048c5761014d565b806306fdde0314610152578063095ea7b3146101d557806318160ddd1461023957806323b872dd14610257578063248a9ca3146102db575b600080fd5b61015a610891565b6040518080602001828103825283818151815260200191508051906020019080838360005b8381101561019a57808201518184015260208101905061017f565b50505050905090810190601f1680156101c75780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b610221600480360360408110156101eb57600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff16906020019092919080359060200190929190505050610933565b60405180821515815260200191505060405180910390f35b610241610951565b6040518082815260200191505060405180910390f35b6102c36004803603606081101561026d57600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190803573ffffffffffffffffffffffffffffffffffffffff1690602001909291908035906020019092919050505061095b565b60405180821515815260200191505060405180910390f35b610307600480360360208110156102f157600080fd5b8101908080359060200190929190505050610a34565b6040518082815260200191505060405180910390f35b6103696004803603604081101561033357600080fd5b8101908080359060200190929190803573ffffffffffffffffffffffffffffffffffffffff169060200190929190505050610a54565b005b610373610ade565b604051808260ff16815260200191505060405180910390f35b6103d8600480360360408110156103a257600080fd5b8101908080359060200190929190803573ffffffffffffffffffffffffffffffffffffffff169060200190929190505050610af5565b005b610426600480360360408110156103f057600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff16906020019092919080359060200190929190505050610b8e565b60405180821515815260200191505060405180910390f35b61048a6004803603604081101561045457600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff16906020019092919080359060200190929190505050610c41565b005b6104ce600480360360208110156104a257600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190505050610ceb565b6040518082815260200191505060405180910390f35b6104ec610d33565b6040518082815260200191505060405180910390f35b6105386004803603604081101561051857600080fd5b810190808035906020019092919080359060200190929190505050610d57565b604051808273ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b6105b06004803603604081101561057a57600080fd5b8101908080359060200190929190803573ffffffffffffffffffffffffffffffffffffffff169060200190929190505050610d89565b60405180821515815260200191505060405180910390f35b6105d0610dbb565b6040518080602001828103825283818151815260200191508051906020019080838360005b838110156106105780820151818401526020810190506105f5565b50505050905090810190601f16801561063d5780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b6106976004803603604081101561066157600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff16906020019092919080359060200190929190505050610e5d565b005b6106a1610f07565b6040518082815260200191505060405180910390f35b610703600480360360408110156106cd57600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff16906020019092919080359060200190929190505050610f0e565b60405180821515815260200191505060405180910390f35b6107676004803603604081101561073157600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff16906020019092919080359060200190929190505050610fdb565b60405180821515815260200191505060405180910390f35b6107ab6004803603602081101561079557600080fd5b8101908080359060200190929190505050610ff9565b6040518082815260200191505060405180910390f35b61080d600480360360408110156107d757600080fd5b8101908080359060200190929190803573ffffffffffffffffffffffffffffffffffffffff169060200190929190505050611020565b005b6108716004803603604081101561082557600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190803573ffffffffffffffffffffffffffffffffffffffff1690602001909291905050506110aa565b6040518082815260200191505060405180910390f35b61088f611131565b005b606060038054600181600116156101000203166002900480601f0160208091040260200160405190810160405280929190818152602001828054600181600116156101000203166002900480156109295780601f106108fe57610100808354040283529160200191610929565b820191906000526020600020905b81548152906001019060200180831161090c57829003601f168201915b5050505050905090565b60006109476109406112e5565b84846112ed565b6001905092915050565b6000600254905090565b60006109688484846114e4565b610a29846109746112e5565b610a248560405180606001604052806028815260200161211460289139600160008b73ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006109da6112e5565b73ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020546117a59092919063ffffffff16565b6112ed565b600190509392505050565b600060066000838152602001908152602001600020600201549050919050565b610a7b6006600084815260200190815260200160002060020154610a766112e5565b610d89565b610ad0576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252602f81526020018061204b602f913960400191505060405180910390fd5b610ada8282611865565b5050565b6000600560009054906101000a900460ff16905090565b610afd6112e5565b73ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff1614610b80576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252602f8152602001806121cb602f913960400191505060405180910390fd5b610b8a82826118f9565b5050565b6000610c37610b9b6112e5565b84610c328560016000610bac6112e5565b73ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008973ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000205461122d90919063ffffffff16565b6112ed565b6001905092915050565b610c6b7fa49807205ce4d355092ef5a8a18f56e8913cf4a201fbe287825b095693c2177533610d89565b610cdd576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260118152602001807f596f7520617265206e6f742061646d696e00000000000000000000000000000081525060200191505060405180910390fd5b610ce7828261198d565b5050565b60008060008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020549050919050565b7fa49807205ce4d355092ef5a8a18f56e8913cf4a201fbe287825b095693c2177581565b6000610d818260066000868152602001908152602001600020600001611b5490919063ffffffff16565b905092915050565b6000610db38260066000868152602001908152602001600020600001611b6e90919063ffffffff16565b905092915050565b606060048054600181600116156101000203166002900480601f016020809104026020016040519081016040528092919081815260200182805460018160011615610100020316600290048015610e535780601f10610e2857610100808354040283529160200191610e53565b820191906000526020600020905b815481529060010190602001808311610e3657829003601f168201915b5050505050905090565b610e877fa49807205ce4d355092ef5a8a18f56e8913cf4a201fbe287825b095693c2177533610d89565b610ef9576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260118152602001807f596f7520617265206e6f742061646d696e00000000000000000000000000000081525060200191505060405180910390fd5b610f038282611b9e565b5050565b6000801b81565b6000610fd1610f1b6112e5565b84610fcc856040518060600160405280602581526020016121a66025913960016000610f456112e5565b73ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008a73ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020546117a59092919063ffffffff16565b6112ed565b6001905092915050565b6000610fef610fe86112e5565b84846114e4565b6001905092915050565b600061101960066000848152602001908152602001600020600001611d62565b9050919050565b61104760066000848152602001908152602001600020600201546110426112e5565b610d89565b61109c576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260308152602001806120e46030913960400191505060405180910390fd5b6110a682826118f9565b5050565b6000600160008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002054905092915050565b61115f33600760009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1660646114e4565b3373ffffffffffffffffffffffffffffffffffffffff167fb0f4fec3f80714b4385839888dd1c5bc0b9f7892fb7d895bebda9fd03c75db5d73__PriceConsumerLib______________________63bb27e0ce6040518163ffffffff1660e01b815260040160206040518083038186803b1580156111db57600080fd5b505af41580156111ef573d6000803e3d6000fd5b505050506040513d602081101561120557600080fd5b81019080805190602001909291905050506040518082815260200191505060405180910390a2565b6000808284019050838110156112ab576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252601b8152602001807f536166654d6174683a206164646974696f6e206f766572666c6f77000000000081525060200191505060405180910390fd5b8091505092915050565b60006112dd836000018373ffffffffffffffffffffffffffffffffffffffff1660001b611d77565b905092915050565b600033905090565b600073ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff161415611373576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260248152602001806121826024913960400191505060405180910390fd5b600073ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff1614156113f9576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252602281526020018061209c6022913960400191505060405180910390fd5b80600160008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020819055508173ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff167f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925836040518082815260200191505060405180910390a3505050565b600073ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff16141561156a576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252602581526020018061215d6025913960400191505060405180910390fd5b600073ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff1614156115f0576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260238152602001806120286023913960400191505060405180910390fd5b6115fb838383611de7565b611666816040518060600160405280602681526020016120be602691396000808773ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020546117a59092919063ffffffff16565b6000808573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020819055506116f9816000808573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000205461122d90919063ffffffff16565b6000808473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020819055508173ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef836040518082815260200191505060405180910390a3505050565b6000838311158290611852576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825283818151815260200191508051906020019080838360005b838110156118175780820151818401526020810190506117fc565b50505050905090810190601f1680156118445780820380516001836020036101000a031916815260200191505b509250505060405180910390fd5b5060008385039050809150509392505050565b61188d81600660008581526020019081526020016000206000016112b590919063ffffffff16565b156118f55761189a6112e5565b73ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff16837f2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d60405160405180910390a45b5050565b6119218160066000858152602001908152602001600020600001611dec90919063ffffffff16565b156119895761192e6112e5565b73ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff16837ff6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b60405160405180910390a45b5050565b600073ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff161415611a30576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252601f8152602001807f45524332303a206d696e7420746f20746865207a65726f20616464726573730081525060200191505060405180910390fd5b611a3c60008383611de7565b611a518160025461122d90919063ffffffff16565b600281905550611aa8816000808573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000205461122d90919063ffffffff16565b6000808473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020819055508173ffffffffffffffffffffffffffffffffffffffff16600073ffffffffffffffffffffffffffffffffffffffff167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef836040518082815260200191505060405180910390a35050565b6000611b638360000183611e1c565b60001c905092915050565b6000611b96836000018373ffffffffffffffffffffffffffffffffffffffff1660001b611e9f565b905092915050565b600073ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff161415611c24576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252602181526020018061213c6021913960400191505060405180910390fd5b611c3082600083611de7565b611c9b8160405180606001604052806022815260200161207a602291396000808673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020546117a59092919063ffffffff16565b6000808473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002081905550611cf281600254611ec290919063ffffffff16565b600281905550600073ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef836040518082815260200191505060405180910390a35050565b6000611d7082600001611f0c565b9050919050565b6000611d838383611e9f565b611ddc578260000182908060018154018082558091505060019003906000526020600020016000909190919091505582600001805490508360010160008481526020019081526020016000208190555060019050611de1565b600090505b92915050565b505050565b6000611e14836000018373ffffffffffffffffffffffffffffffffffffffff1660001b611f1d565b905092915050565b600081836000018054905011611e7d576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260228152602001806120066022913960400191505060405180910390fd5b826000018281548110611e8c57fe5b9060005260206000200154905092915050565b600080836001016000848152602001908152602001600020541415905092915050565b6000611f0483836040518060400160405280601e81526020017f536166654d6174683a207375627472616374696f6e206f766572666c6f7700008152506117a5565b905092915050565b600081600001805490509050919050565b60008083600101600084815260200190815260200160002054905060008114611ff95760006001820390506000600186600001805490500390506000866000018281548110611f6857fe5b9060005260206000200154905080876000018481548110611f8557fe5b9060005260206000200181905550600183018760010160008381526020019081526020016000208190555086600001805480611fbd57fe5b60019003818190600052602060002001600090559055866001016000878152602001908152602001600020600090556001945050505050611fff565b60009150505b9291505056fe456e756d657261626c655365743a20696e646578206f7574206f6620626f756e647345524332303a207472616e7366657220746f20746865207a65726f2061646472657373416363657373436f6e74726f6c3a2073656e646572206d75737420626520616e2061646d696e20746f206772616e7445524332303a206275726e20616d6f756e7420657863656564732062616c616e636545524332303a20617070726f766520746f20746865207a65726f206164647265737345524332303a207472616e7366657220616d6f756e7420657863656564732062616c616e6365416363657373436f6e74726f6c3a2073656e646572206d75737420626520616e2061646d696e20746f207265766f6b6545524332303a207472616e7366657220616d6f756e74206578636565647320616c6c6f77616e636545524332303a206275726e2066726f6d20746865207a65726f206164647265737345524332303a207472616e736665722066726f6d20746865207a65726f206164647265737345524332303a20617070726f76652066726f6d20746865207a65726f206164647265737345524332303a2064656372656173656420616c6c6f77616e63652062656c6f77207a65726f416363657373436f6e74726f6c3a2063616e206f6e6c792072656e6f756e636520726f6c657320666f722073656c66a2646970667358221220a218e20a276223ce369789b9076e7c789144151db5bc4851697b26fb4426817e64736f6c634300060c0033"

// DeployMattsuCoin deploys a new Ethereum contract, binding an instance of MattsuCoin to it.
func DeployMattsuCoin(auth *bind.TransactOpts, backend bind.ContractBackend, initialSupply *big.Int, admin common.Address) (common.Address, *types.Transaction, *MattsuCoin, error) {
	parsed, err := abi.JSON(strings.NewReader(MattsuCoinABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(MattsuCoinBin), backend, initialSupply, admin)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &MattsuCoin{MattsuCoinCaller: MattsuCoinCaller{contract: contract}, MattsuCoinTransactor: MattsuCoinTransactor{contract: contract}, MattsuCoinFilterer: MattsuCoinFilterer{contract: contract}}, nil
}

// MattsuCoin is an auto generated Go binding around an Ethereum contract.
type MattsuCoin struct {
	MattsuCoinCaller     // Read-only binding to the contract
	MattsuCoinTransactor // Write-only binding to the contract
	MattsuCoinFilterer   // Log filterer for contract events
}

// MattsuCoinCaller is an auto generated read-only Go binding around an Ethereum contract.
type MattsuCoinCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MattsuCoinTransactor is an auto generated write-only Go binding around an Ethereum contract.
type MattsuCoinTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MattsuCoinFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type MattsuCoinFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MattsuCoinSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type MattsuCoinSession struct {
	Contract     *MattsuCoin       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MattsuCoinCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type MattsuCoinCallerSession struct {
	Contract *MattsuCoinCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// MattsuCoinTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type MattsuCoinTransactorSession struct {
	Contract     *MattsuCoinTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// MattsuCoinRaw is an auto generated low-level Go binding around an Ethereum contract.
type MattsuCoinRaw struct {
	Contract *MattsuCoin // Generic contract binding to access the raw methods on
}

// MattsuCoinCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type MattsuCoinCallerRaw struct {
	Contract *MattsuCoinCaller // Generic read-only contract binding to access the raw methods on
}

// MattsuCoinTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type MattsuCoinTransactorRaw struct {
	Contract *MattsuCoinTransactor // Generic write-only contract binding to access the raw methods on
}

// NewMattsuCoin creates a new instance of MattsuCoin, bound to a specific deployed contract.
func NewMattsuCoin(address common.Address, backend bind.ContractBackend) (*MattsuCoin, error) {
	contract, err := bindMattsuCoin(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &MattsuCoin{MattsuCoinCaller: MattsuCoinCaller{contract: contract}, MattsuCoinTransactor: MattsuCoinTransactor{contract: contract}, MattsuCoinFilterer: MattsuCoinFilterer{contract: contract}}, nil
}

// NewMattsuCoinCaller creates a new read-only instance of MattsuCoin, bound to a specific deployed contract.
func NewMattsuCoinCaller(address common.Address, caller bind.ContractCaller) (*MattsuCoinCaller, error) {
	contract, err := bindMattsuCoin(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &MattsuCoinCaller{contract: contract}, nil
}

// NewMattsuCoinTransactor creates a new write-only instance of MattsuCoin, bound to a specific deployed contract.
func NewMattsuCoinTransactor(address common.Address, transactor bind.ContractTransactor) (*MattsuCoinTransactor, error) {
	contract, err := bindMattsuCoin(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &MattsuCoinTransactor{contract: contract}, nil
}

// NewMattsuCoinFilterer creates a new log filterer instance of MattsuCoin, bound to a specific deployed contract.
func NewMattsuCoinFilterer(address common.Address, filterer bind.ContractFilterer) (*MattsuCoinFilterer, error) {
	contract, err := bindMattsuCoin(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &MattsuCoinFilterer{contract: contract}, nil
}

// bindMattsuCoin binds a generic wrapper to an already deployed contract.
func bindMattsuCoin(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(MattsuCoinABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MattsuCoin *MattsuCoinRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _MattsuCoin.Contract.MattsuCoinCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MattsuCoin *MattsuCoinRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MattsuCoin.Contract.MattsuCoinTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MattsuCoin *MattsuCoinRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MattsuCoin.Contract.MattsuCoinTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MattsuCoin *MattsuCoinCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _MattsuCoin.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MattsuCoin *MattsuCoinTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MattsuCoin.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MattsuCoin *MattsuCoinTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MattsuCoin.Contract.contract.Transact(opts, method, params...)
}

// ADMINROLE is a free data retrieval call binding the contract method 0x75b238fc.
//
// Solidity: function ADMIN_ROLE() view returns(bytes32)
func (_MattsuCoin *MattsuCoinCaller) ADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _MattsuCoin.contract.Call(opts, out, "ADMIN_ROLE")
	return *ret0, err
}

// ADMINROLE is a free data retrieval call binding the contract method 0x75b238fc.
//
// Solidity: function ADMIN_ROLE() view returns(bytes32)
func (_MattsuCoin *MattsuCoinSession) ADMINROLE() ([32]byte, error) {
	return _MattsuCoin.Contract.ADMINROLE(&_MattsuCoin.CallOpts)
}

// ADMINROLE is a free data retrieval call binding the contract method 0x75b238fc.
//
// Solidity: function ADMIN_ROLE() view returns(bytes32)
func (_MattsuCoin *MattsuCoinCallerSession) ADMINROLE() ([32]byte, error) {
	return _MattsuCoin.Contract.ADMINROLE(&_MattsuCoin.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_MattsuCoin *MattsuCoinCaller) DEFAULTADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _MattsuCoin.contract.Call(opts, out, "DEFAULT_ADMIN_ROLE")
	return *ret0, err
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_MattsuCoin *MattsuCoinSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _MattsuCoin.Contract.DEFAULTADMINROLE(&_MattsuCoin.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_MattsuCoin *MattsuCoinCallerSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _MattsuCoin.Contract.DEFAULTADMINROLE(&_MattsuCoin.CallOpts)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_MattsuCoin *MattsuCoinCaller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _MattsuCoin.contract.Call(opts, out, "allowance", owner, spender)
	return *ret0, err
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_MattsuCoin *MattsuCoinSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _MattsuCoin.Contract.Allowance(&_MattsuCoin.CallOpts, owner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_MattsuCoin *MattsuCoinCallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _MattsuCoin.Contract.Allowance(&_MattsuCoin.CallOpts, owner, spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_MattsuCoin *MattsuCoinCaller) BalanceOf(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _MattsuCoin.contract.Call(opts, out, "balanceOf", account)
	return *ret0, err
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_MattsuCoin *MattsuCoinSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _MattsuCoin.Contract.BalanceOf(&_MattsuCoin.CallOpts, account)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_MattsuCoin *MattsuCoinCallerSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _MattsuCoin.Contract.BalanceOf(&_MattsuCoin.CallOpts, account)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_MattsuCoin *MattsuCoinCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var (
		ret0 = new(uint8)
	)
	out := ret0
	err := _MattsuCoin.contract.Call(opts, out, "decimals")
	return *ret0, err
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_MattsuCoin *MattsuCoinSession) Decimals() (uint8, error) {
	return _MattsuCoin.Contract.Decimals(&_MattsuCoin.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_MattsuCoin *MattsuCoinCallerSession) Decimals() (uint8, error) {
	return _MattsuCoin.Contract.Decimals(&_MattsuCoin.CallOpts)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_MattsuCoin *MattsuCoinCaller) GetRoleAdmin(opts *bind.CallOpts, role [32]byte) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _MattsuCoin.contract.Call(opts, out, "getRoleAdmin", role)
	return *ret0, err
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_MattsuCoin *MattsuCoinSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _MattsuCoin.Contract.GetRoleAdmin(&_MattsuCoin.CallOpts, role)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_MattsuCoin *MattsuCoinCallerSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _MattsuCoin.Contract.GetRoleAdmin(&_MattsuCoin.CallOpts, role)
}

// GetRoleMember is a free data retrieval call binding the contract method 0x9010d07c.
//
// Solidity: function getRoleMember(bytes32 role, uint256 index) view returns(address)
func (_MattsuCoin *MattsuCoinCaller) GetRoleMember(opts *bind.CallOpts, role [32]byte, index *big.Int) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _MattsuCoin.contract.Call(opts, out, "getRoleMember", role, index)
	return *ret0, err
}

// GetRoleMember is a free data retrieval call binding the contract method 0x9010d07c.
//
// Solidity: function getRoleMember(bytes32 role, uint256 index) view returns(address)
func (_MattsuCoin *MattsuCoinSession) GetRoleMember(role [32]byte, index *big.Int) (common.Address, error) {
	return _MattsuCoin.Contract.GetRoleMember(&_MattsuCoin.CallOpts, role, index)
}

// GetRoleMember is a free data retrieval call binding the contract method 0x9010d07c.
//
// Solidity: function getRoleMember(bytes32 role, uint256 index) view returns(address)
func (_MattsuCoin *MattsuCoinCallerSession) GetRoleMember(role [32]byte, index *big.Int) (common.Address, error) {
	return _MattsuCoin.Contract.GetRoleMember(&_MattsuCoin.CallOpts, role, index)
}

// GetRoleMemberCount is a free data retrieval call binding the contract method 0xca15c873.
//
// Solidity: function getRoleMemberCount(bytes32 role) view returns(uint256)
func (_MattsuCoin *MattsuCoinCaller) GetRoleMemberCount(opts *bind.CallOpts, role [32]byte) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _MattsuCoin.contract.Call(opts, out, "getRoleMemberCount", role)
	return *ret0, err
}

// GetRoleMemberCount is a free data retrieval call binding the contract method 0xca15c873.
//
// Solidity: function getRoleMemberCount(bytes32 role) view returns(uint256)
func (_MattsuCoin *MattsuCoinSession) GetRoleMemberCount(role [32]byte) (*big.Int, error) {
	return _MattsuCoin.Contract.GetRoleMemberCount(&_MattsuCoin.CallOpts, role)
}

// GetRoleMemberCount is a free data retrieval call binding the contract method 0xca15c873.
//
// Solidity: function getRoleMemberCount(bytes32 role) view returns(uint256)
func (_MattsuCoin *MattsuCoinCallerSession) GetRoleMemberCount(role [32]byte) (*big.Int, error) {
	return _MattsuCoin.Contract.GetRoleMemberCount(&_MattsuCoin.CallOpts, role)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_MattsuCoin *MattsuCoinCaller) HasRole(opts *bind.CallOpts, role [32]byte, account common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _MattsuCoin.contract.Call(opts, out, "hasRole", role, account)
	return *ret0, err
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_MattsuCoin *MattsuCoinSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _MattsuCoin.Contract.HasRole(&_MattsuCoin.CallOpts, role, account)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_MattsuCoin *MattsuCoinCallerSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _MattsuCoin.Contract.HasRole(&_MattsuCoin.CallOpts, role, account)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_MattsuCoin *MattsuCoinCaller) Name(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _MattsuCoin.contract.Call(opts, out, "name")
	return *ret0, err
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_MattsuCoin *MattsuCoinSession) Name() (string, error) {
	return _MattsuCoin.Contract.Name(&_MattsuCoin.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_MattsuCoin *MattsuCoinCallerSession) Name() (string, error) {
	return _MattsuCoin.Contract.Name(&_MattsuCoin.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_MattsuCoin *MattsuCoinCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _MattsuCoin.contract.Call(opts, out, "symbol")
	return *ret0, err
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_MattsuCoin *MattsuCoinSession) Symbol() (string, error) {
	return _MattsuCoin.Contract.Symbol(&_MattsuCoin.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_MattsuCoin *MattsuCoinCallerSession) Symbol() (string, error) {
	return _MattsuCoin.Contract.Symbol(&_MattsuCoin.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_MattsuCoin *MattsuCoinCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _MattsuCoin.contract.Call(opts, out, "totalSupply")
	return *ret0, err
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_MattsuCoin *MattsuCoinSession) TotalSupply() (*big.Int, error) {
	return _MattsuCoin.Contract.TotalSupply(&_MattsuCoin.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_MattsuCoin *MattsuCoinCallerSession) TotalSupply() (*big.Int, error) {
	return _MattsuCoin.Contract.TotalSupply(&_MattsuCoin.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_MattsuCoin *MattsuCoinTransactor) Approve(opts *bind.TransactOpts, spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _MattsuCoin.contract.Transact(opts, "approve", spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_MattsuCoin *MattsuCoinSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _MattsuCoin.Contract.Approve(&_MattsuCoin.TransactOpts, spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_MattsuCoin *MattsuCoinTransactorSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _MattsuCoin.Contract.Approve(&_MattsuCoin.TransactOpts, spender, amount)
}

// Burn is a paid mutator transaction binding the contract method 0x9dc29fac.
//
// Solidity: function burn(address to, uint256 amount) returns()
func (_MattsuCoin *MattsuCoinTransactor) Burn(opts *bind.TransactOpts, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _MattsuCoin.contract.Transact(opts, "burn", to, amount)
}

// Burn is a paid mutator transaction binding the contract method 0x9dc29fac.
//
// Solidity: function burn(address to, uint256 amount) returns()
func (_MattsuCoin *MattsuCoinSession) Burn(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _MattsuCoin.Contract.Burn(&_MattsuCoin.TransactOpts, to, amount)
}

// Burn is a paid mutator transaction binding the contract method 0x9dc29fac.
//
// Solidity: function burn(address to, uint256 amount) returns()
func (_MattsuCoin *MattsuCoinTransactorSession) Burn(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _MattsuCoin.Contract.Burn(&_MattsuCoin.TransactOpts, to, amount)
}

// BuyLatestEthUsd is a paid mutator transaction binding the contract method 0xfc3feb53.
//
// Solidity: function buyLatestEthUsd() returns()
func (_MattsuCoin *MattsuCoinTransactor) BuyLatestEthUsd(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MattsuCoin.contract.Transact(opts, "buyLatestEthUsd")
}

// BuyLatestEthUsd is a paid mutator transaction binding the contract method 0xfc3feb53.
//
// Solidity: function buyLatestEthUsd() returns()
func (_MattsuCoin *MattsuCoinSession) BuyLatestEthUsd() (*types.Transaction, error) {
	return _MattsuCoin.Contract.BuyLatestEthUsd(&_MattsuCoin.TransactOpts)
}

// BuyLatestEthUsd is a paid mutator transaction binding the contract method 0xfc3feb53.
//
// Solidity: function buyLatestEthUsd() returns()
func (_MattsuCoin *MattsuCoinTransactorSession) BuyLatestEthUsd() (*types.Transaction, error) {
	return _MattsuCoin.Contract.BuyLatestEthUsd(&_MattsuCoin.TransactOpts)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_MattsuCoin *MattsuCoinTransactor) DecreaseAllowance(opts *bind.TransactOpts, spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _MattsuCoin.contract.Transact(opts, "decreaseAllowance", spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_MattsuCoin *MattsuCoinSession) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _MattsuCoin.Contract.DecreaseAllowance(&_MattsuCoin.TransactOpts, spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_MattsuCoin *MattsuCoinTransactorSession) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _MattsuCoin.Contract.DecreaseAllowance(&_MattsuCoin.TransactOpts, spender, subtractedValue)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_MattsuCoin *MattsuCoinTransactor) GrantRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _MattsuCoin.contract.Transact(opts, "grantRole", role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_MattsuCoin *MattsuCoinSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _MattsuCoin.Contract.GrantRole(&_MattsuCoin.TransactOpts, role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_MattsuCoin *MattsuCoinTransactorSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _MattsuCoin.Contract.GrantRole(&_MattsuCoin.TransactOpts, role, account)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_MattsuCoin *MattsuCoinTransactor) IncreaseAllowance(opts *bind.TransactOpts, spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _MattsuCoin.contract.Transact(opts, "increaseAllowance", spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_MattsuCoin *MattsuCoinSession) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _MattsuCoin.Contract.IncreaseAllowance(&_MattsuCoin.TransactOpts, spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_MattsuCoin *MattsuCoinTransactorSession) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _MattsuCoin.Contract.IncreaseAllowance(&_MattsuCoin.TransactOpts, spender, addedValue)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address to, uint256 amount) returns()
func (_MattsuCoin *MattsuCoinTransactor) Mint(opts *bind.TransactOpts, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _MattsuCoin.contract.Transact(opts, "mint", to, amount)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address to, uint256 amount) returns()
func (_MattsuCoin *MattsuCoinSession) Mint(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _MattsuCoin.Contract.Mint(&_MattsuCoin.TransactOpts, to, amount)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address to, uint256 amount) returns()
func (_MattsuCoin *MattsuCoinTransactorSession) Mint(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _MattsuCoin.Contract.Mint(&_MattsuCoin.TransactOpts, to, amount)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_MattsuCoin *MattsuCoinTransactor) RenounceRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _MattsuCoin.contract.Transact(opts, "renounceRole", role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_MattsuCoin *MattsuCoinSession) RenounceRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _MattsuCoin.Contract.RenounceRole(&_MattsuCoin.TransactOpts, role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_MattsuCoin *MattsuCoinTransactorSession) RenounceRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _MattsuCoin.Contract.RenounceRole(&_MattsuCoin.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_MattsuCoin *MattsuCoinTransactor) RevokeRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _MattsuCoin.contract.Transact(opts, "revokeRole", role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_MattsuCoin *MattsuCoinSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _MattsuCoin.Contract.RevokeRole(&_MattsuCoin.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_MattsuCoin *MattsuCoinTransactorSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _MattsuCoin.Contract.RevokeRole(&_MattsuCoin.TransactOpts, role, account)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_MattsuCoin *MattsuCoinTransactor) Transfer(opts *bind.TransactOpts, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _MattsuCoin.contract.Transact(opts, "transfer", recipient, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_MattsuCoin *MattsuCoinSession) Transfer(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _MattsuCoin.Contract.Transfer(&_MattsuCoin.TransactOpts, recipient, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_MattsuCoin *MattsuCoinTransactorSession) Transfer(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _MattsuCoin.Contract.Transfer(&_MattsuCoin.TransactOpts, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_MattsuCoin *MattsuCoinTransactor) TransferFrom(opts *bind.TransactOpts, sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _MattsuCoin.contract.Transact(opts, "transferFrom", sender, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_MattsuCoin *MattsuCoinSession) TransferFrom(sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _MattsuCoin.Contract.TransferFrom(&_MattsuCoin.TransactOpts, sender, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_MattsuCoin *MattsuCoinTransactorSession) TransferFrom(sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _MattsuCoin.Contract.TransferFrom(&_MattsuCoin.TransactOpts, sender, recipient, amount)
}

// MattsuCoinApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the MattsuCoin contract.
type MattsuCoinApprovalIterator struct {
	Event *MattsuCoinApproval // Event containing the contract specifics and raw log

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
func (it *MattsuCoinApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MattsuCoinApproval)
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
		it.Event = new(MattsuCoinApproval)
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
func (it *MattsuCoinApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MattsuCoinApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MattsuCoinApproval represents a Approval event raised by the MattsuCoin contract.
type MattsuCoinApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_MattsuCoin *MattsuCoinFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*MattsuCoinApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _MattsuCoin.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &MattsuCoinApprovalIterator{contract: _MattsuCoin.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_MattsuCoin *MattsuCoinFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *MattsuCoinApproval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _MattsuCoin.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MattsuCoinApproval)
				if err := _MattsuCoin.contract.UnpackLog(event, "Approval", log); err != nil {
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
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_MattsuCoin *MattsuCoinFilterer) ParseApproval(log types.Log) (*MattsuCoinApproval, error) {
	event := new(MattsuCoinApproval)
	if err := _MattsuCoin.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	return event, nil
}

// MattsuCoinBuyEthUsdIterator is returned from FilterBuyEthUsd and is used to iterate over the raw logs and unpacked data for BuyEthUsd events raised by the MattsuCoin contract.
type MattsuCoinBuyEthUsdIterator struct {
	Event *MattsuCoinBuyEthUsd // Event containing the contract specifics and raw log

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
func (it *MattsuCoinBuyEthUsdIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MattsuCoinBuyEthUsd)
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
		it.Event = new(MattsuCoinBuyEthUsd)
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
func (it *MattsuCoinBuyEthUsdIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MattsuCoinBuyEthUsdIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MattsuCoinBuyEthUsd represents a BuyEthUsd event raised by the MattsuCoin contract.
type MattsuCoinBuyEthUsd struct {
	Buyer  common.Address
	EthUsd *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterBuyEthUsd is a free log retrieval operation binding the contract event 0xb0f4fec3f80714b4385839888dd1c5bc0b9f7892fb7d895bebda9fd03c75db5d.
//
// Solidity: event BuyEthUsd(address indexed buyer, int256 ethUsd)
func (_MattsuCoin *MattsuCoinFilterer) FilterBuyEthUsd(opts *bind.FilterOpts, buyer []common.Address) (*MattsuCoinBuyEthUsdIterator, error) {

	var buyerRule []interface{}
	for _, buyerItem := range buyer {
		buyerRule = append(buyerRule, buyerItem)
	}

	logs, sub, err := _MattsuCoin.contract.FilterLogs(opts, "BuyEthUsd", buyerRule)
	if err != nil {
		return nil, err
	}
	return &MattsuCoinBuyEthUsdIterator{contract: _MattsuCoin.contract, event: "BuyEthUsd", logs: logs, sub: sub}, nil
}

// WatchBuyEthUsd is a free log subscription operation binding the contract event 0xb0f4fec3f80714b4385839888dd1c5bc0b9f7892fb7d895bebda9fd03c75db5d.
//
// Solidity: event BuyEthUsd(address indexed buyer, int256 ethUsd)
func (_MattsuCoin *MattsuCoinFilterer) WatchBuyEthUsd(opts *bind.WatchOpts, sink chan<- *MattsuCoinBuyEthUsd, buyer []common.Address) (event.Subscription, error) {

	var buyerRule []interface{}
	for _, buyerItem := range buyer {
		buyerRule = append(buyerRule, buyerItem)
	}

	logs, sub, err := _MattsuCoin.contract.WatchLogs(opts, "BuyEthUsd", buyerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MattsuCoinBuyEthUsd)
				if err := _MattsuCoin.contract.UnpackLog(event, "BuyEthUsd", log); err != nil {
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

// ParseBuyEthUsd is a log parse operation binding the contract event 0xb0f4fec3f80714b4385839888dd1c5bc0b9f7892fb7d895bebda9fd03c75db5d.
//
// Solidity: event BuyEthUsd(address indexed buyer, int256 ethUsd)
func (_MattsuCoin *MattsuCoinFilterer) ParseBuyEthUsd(log types.Log) (*MattsuCoinBuyEthUsd, error) {
	event := new(MattsuCoinBuyEthUsd)
	if err := _MattsuCoin.contract.UnpackLog(event, "BuyEthUsd", log); err != nil {
		return nil, err
	}
	return event, nil
}

// MattsuCoinRoleAdminChangedIterator is returned from FilterRoleAdminChanged and is used to iterate over the raw logs and unpacked data for RoleAdminChanged events raised by the MattsuCoin contract.
type MattsuCoinRoleAdminChangedIterator struct {
	Event *MattsuCoinRoleAdminChanged // Event containing the contract specifics and raw log

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
func (it *MattsuCoinRoleAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MattsuCoinRoleAdminChanged)
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
		it.Event = new(MattsuCoinRoleAdminChanged)
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
func (it *MattsuCoinRoleAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MattsuCoinRoleAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MattsuCoinRoleAdminChanged represents a RoleAdminChanged event raised by the MattsuCoin contract.
type MattsuCoinRoleAdminChanged struct {
	Role              [32]byte
	PreviousAdminRole [32]byte
	NewAdminRole      [32]byte
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterRoleAdminChanged is a free log retrieval operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_MattsuCoin *MattsuCoinFilterer) FilterRoleAdminChanged(opts *bind.FilterOpts, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (*MattsuCoinRoleAdminChangedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var previousAdminRoleRule []interface{}
	for _, previousAdminRoleItem := range previousAdminRole {
		previousAdminRoleRule = append(previousAdminRoleRule, previousAdminRoleItem)
	}
	var newAdminRoleRule []interface{}
	for _, newAdminRoleItem := range newAdminRole {
		newAdminRoleRule = append(newAdminRoleRule, newAdminRoleItem)
	}

	logs, sub, err := _MattsuCoin.contract.FilterLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return &MattsuCoinRoleAdminChangedIterator{contract: _MattsuCoin.contract, event: "RoleAdminChanged", logs: logs, sub: sub}, nil
}

// WatchRoleAdminChanged is a free log subscription operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_MattsuCoin *MattsuCoinFilterer) WatchRoleAdminChanged(opts *bind.WatchOpts, sink chan<- *MattsuCoinRoleAdminChanged, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var previousAdminRoleRule []interface{}
	for _, previousAdminRoleItem := range previousAdminRole {
		previousAdminRoleRule = append(previousAdminRoleRule, previousAdminRoleItem)
	}
	var newAdminRoleRule []interface{}
	for _, newAdminRoleItem := range newAdminRole {
		newAdminRoleRule = append(newAdminRoleRule, newAdminRoleItem)
	}

	logs, sub, err := _MattsuCoin.contract.WatchLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MattsuCoinRoleAdminChanged)
				if err := _MattsuCoin.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
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

// ParseRoleAdminChanged is a log parse operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_MattsuCoin *MattsuCoinFilterer) ParseRoleAdminChanged(log types.Log) (*MattsuCoinRoleAdminChanged, error) {
	event := new(MattsuCoinRoleAdminChanged)
	if err := _MattsuCoin.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
		return nil, err
	}
	return event, nil
}

// MattsuCoinRoleGrantedIterator is returned from FilterRoleGranted and is used to iterate over the raw logs and unpacked data for RoleGranted events raised by the MattsuCoin contract.
type MattsuCoinRoleGrantedIterator struct {
	Event *MattsuCoinRoleGranted // Event containing the contract specifics and raw log

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
func (it *MattsuCoinRoleGrantedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MattsuCoinRoleGranted)
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
		it.Event = new(MattsuCoinRoleGranted)
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
func (it *MattsuCoinRoleGrantedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MattsuCoinRoleGrantedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MattsuCoinRoleGranted represents a RoleGranted event raised by the MattsuCoin contract.
type MattsuCoinRoleGranted struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleGranted is a free log retrieval operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_MattsuCoin *MattsuCoinFilterer) FilterRoleGranted(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*MattsuCoinRoleGrantedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _MattsuCoin.contract.FilterLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &MattsuCoinRoleGrantedIterator{contract: _MattsuCoin.contract, event: "RoleGranted", logs: logs, sub: sub}, nil
}

// WatchRoleGranted is a free log subscription operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_MattsuCoin *MattsuCoinFilterer) WatchRoleGranted(opts *bind.WatchOpts, sink chan<- *MattsuCoinRoleGranted, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _MattsuCoin.contract.WatchLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MattsuCoinRoleGranted)
				if err := _MattsuCoin.contract.UnpackLog(event, "RoleGranted", log); err != nil {
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

// ParseRoleGranted is a log parse operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_MattsuCoin *MattsuCoinFilterer) ParseRoleGranted(log types.Log) (*MattsuCoinRoleGranted, error) {
	event := new(MattsuCoinRoleGranted)
	if err := _MattsuCoin.contract.UnpackLog(event, "RoleGranted", log); err != nil {
		return nil, err
	}
	return event, nil
}

// MattsuCoinRoleRevokedIterator is returned from FilterRoleRevoked and is used to iterate over the raw logs and unpacked data for RoleRevoked events raised by the MattsuCoin contract.
type MattsuCoinRoleRevokedIterator struct {
	Event *MattsuCoinRoleRevoked // Event containing the contract specifics and raw log

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
func (it *MattsuCoinRoleRevokedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MattsuCoinRoleRevoked)
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
		it.Event = new(MattsuCoinRoleRevoked)
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
func (it *MattsuCoinRoleRevokedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MattsuCoinRoleRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MattsuCoinRoleRevoked represents a RoleRevoked event raised by the MattsuCoin contract.
type MattsuCoinRoleRevoked struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleRevoked is a free log retrieval operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_MattsuCoin *MattsuCoinFilterer) FilterRoleRevoked(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*MattsuCoinRoleRevokedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _MattsuCoin.contract.FilterLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &MattsuCoinRoleRevokedIterator{contract: _MattsuCoin.contract, event: "RoleRevoked", logs: logs, sub: sub}, nil
}

// WatchRoleRevoked is a free log subscription operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_MattsuCoin *MattsuCoinFilterer) WatchRoleRevoked(opts *bind.WatchOpts, sink chan<- *MattsuCoinRoleRevoked, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _MattsuCoin.contract.WatchLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MattsuCoinRoleRevoked)
				if err := _MattsuCoin.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
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

// ParseRoleRevoked is a log parse operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_MattsuCoin *MattsuCoinFilterer) ParseRoleRevoked(log types.Log) (*MattsuCoinRoleRevoked, error) {
	event := new(MattsuCoinRoleRevoked)
	if err := _MattsuCoin.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
		return nil, err
	}
	return event, nil
}

// MattsuCoinTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the MattsuCoin contract.
type MattsuCoinTransferIterator struct {
	Event *MattsuCoinTransfer // Event containing the contract specifics and raw log

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
func (it *MattsuCoinTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MattsuCoinTransfer)
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
		it.Event = new(MattsuCoinTransfer)
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
func (it *MattsuCoinTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MattsuCoinTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MattsuCoinTransfer represents a Transfer event raised by the MattsuCoin contract.
type MattsuCoinTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_MattsuCoin *MattsuCoinFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*MattsuCoinTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _MattsuCoin.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &MattsuCoinTransferIterator{contract: _MattsuCoin.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_MattsuCoin *MattsuCoinFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *MattsuCoinTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _MattsuCoin.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MattsuCoinTransfer)
				if err := _MattsuCoin.contract.UnpackLog(event, "Transfer", log); err != nil {
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
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_MattsuCoin *MattsuCoinFilterer) ParseTransfer(log types.Log) (*MattsuCoinTransfer, error) {
	event := new(MattsuCoinTransfer)
	if err := _MattsuCoin.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	return event, nil
}
