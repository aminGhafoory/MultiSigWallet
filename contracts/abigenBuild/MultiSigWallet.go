// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package MultiSig

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

// MultiSigMetaData contains all meta data concerning the MultiSig contract.
var MultiSigMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"_owners\",\"type\":\"address[]\"},{\"internalType\":\"uint256\",\"name\":\"_numConfirmationsRequired\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"txIndex\",\"type\":\"uint256\"}],\"name\":\"ConfirmTransaction\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"}],\"name\":\"Deposit\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"txIndex\",\"type\":\"uint256\"}],\"name\":\"ExecuteTransaction\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"txIndex\",\"type\":\"uint256\"}],\"name\":\"RevokeConfirmation\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"txIndex\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"SubmitTransaction\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_txIndex\",\"type\":\"uint256\"}],\"name\":\"confirmTransaction\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_txIndex\",\"type\":\"uint256\"}],\"name\":\"executeTransaction\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getOwners\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_txIndex\",\"type\":\"uint256\"}],\"name\":\"getTransaction\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"internalType\":\"bool\",\"name\":\"executed\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"numConfirmations\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getTransactionCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"isConfirmed\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"isOwner\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"numConfirmationsRequired\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"owners\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_txIndex\",\"type\":\"uint256\"}],\"name\":\"revokeConfirmation\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_value\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"_data\",\"type\":\"bytes\"}],\"name\":\"submitTransaction\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"transactions\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"internalType\":\"bool\",\"name\":\"executed\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"numConfirmations\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
	Bin: "0x608060405234801561000f575f80fd5b5060405161246f38038061246f833981810160405281019061003191906104a2565b5f825111610074576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161006b90610556565b60405180910390fd5b5f81118015610084575081518111155b6100c3576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016100ba906105e4565b60405180910390fd5b5f5b82518110156102a5575f8382815181106100e2576100e1610602565b5b602002602001015190505f73ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff160361015a576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161015190610679565b60405180910390fd5b60015f8273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f9054906101000a900460ff16156101e4576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016101db906106e1565b60405180910390fd5b6001805f8373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f6101000a81548160ff0219169083151502179055505f81908060018154018082558091505060019003905f5260205f20015f9091909190916101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055505080806001019150506100c5565b508060028190555050506106ff565b5f604051905090565b5f80fd5b5f80fd5b5f80fd5b5f601f19601f8301169050919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52604160045260245ffd5b61030f826102c9565b810181811067ffffffffffffffff8211171561032e5761032d6102d9565b5b80604052505050565b5f6103406102b4565b905061034c8282610306565b919050565b5f67ffffffffffffffff82111561036b5761036a6102d9565b5b602082029050602081019050919050565b5f80fd5b5f73ffffffffffffffffffffffffffffffffffffffff82169050919050565b5f6103a982610380565b9050919050565b6103b98161039f565b81146103c3575f80fd5b50565b5f815190506103d4816103b0565b92915050565b5f6103ec6103e784610351565b610337565b9050808382526020820190506020840283018581111561040f5761040e61037c565b5b835b81811015610438578061042488826103c6565b845260208401935050602081019050610411565b5050509392505050565b5f82601f830112610456576104556102c5565b5b81516104668482602086016103da565b91505092915050565b5f819050919050565b6104818161046f565b811461048b575f80fd5b50565b5f8151905061049c81610478565b92915050565b5f80604083850312156104b8576104b76102bd565b5b5f83015167ffffffffffffffff8111156104d5576104d46102c1565b5b6104e185828601610442565b92505060206104f28582860161048e565b9150509250929050565b5f82825260208201905092915050565b7f6f776e65727320726571756972656400000000000000000000000000000000005f82015250565b5f610540600f836104fc565b915061054b8261050c565b602082019050919050565b5f6020820190508181035f83015261056d81610534565b9050919050565b7f696e76616c6964206e756d626572206f6620726571756972656420636f6e66695f8201527f726d6174696f6e73000000000000000000000000000000000000000000000000602082015250565b5f6105ce6028836104fc565b91506105d982610574565b604082019050919050565b5f6020820190508181035f8301526105fb816105c2565b9050919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52603260045260245ffd5b7f696e76616c6964206f776e6572000000000000000000000000000000000000005f82015250565b5f610663600d836104fc565b915061066e8261062f565b602082019050919050565b5f6020820190508181035f83015261069081610657565b9050919050565b7f6f776e6572206e6f7420756e69717565000000000000000000000000000000005f82015250565b5f6106cb6010836104fc565b91506106d682610697565b602082019050919050565b5f6020820190508181035f8301526106f8816106bf565b9050919050565b611d638061070c5f395ff3fe6080604052600436106100aa575f3560e01c80639ace38c2116100635780639ace38c21461024b578063a0e67e2b1461028b578063c01a8c84146102b5578063c6427474146102dd578063d0549b8514610305578063ee22610b1461032f57610101565b8063025e7c271461010557806320ea8d86146101415780632e7700f0146101695780632f54bf6e1461019357806333ea3dc8146101cf57806380f59a651461020f57610101565b36610101573373ffffffffffffffffffffffffffffffffffffffff167f90890809c654f11d6e72a28fa60149770a0d11ec6c92319d6ceb2bb0a4ea1a1534476040516100f79291906110f3565b60405180910390a2005b5f80fd5b348015610110575f80fd5b5061012b60048036038101906101269190611155565b610357565b60405161013891906111bf565b60405180910390f35b34801561014c575f80fd5b5061016760048036038101906101629190611155565b610391565b005b348015610174575f80fd5b5061017d610656565b60405161018a91906111d8565b60405180910390f35b34801561019e575f80fd5b506101b960048036038101906101b4919061121b565b610662565b6040516101c69190611260565b60405180910390f35b3480156101da575f80fd5b506101f560048036038101906101f09190611155565b61067f565b6040516102069594939291906112e9565b60405180910390f35b34801561021a575f80fd5b5061023560048036038101906102309190611341565b610788565b6040516102429190611260565b60405180910390f35b348015610256575f80fd5b50610271600480360381019061026c9190611155565b6107b2565b6040516102829594939291906112e9565b60405180910390f35b348015610296575f80fd5b5061029f6108a4565b6040516102ac9190611436565b60405180910390f35b3480156102c0575f80fd5b506102db60048036038101906102d69190611155565b61092e565b005b3480156102e8575f80fd5b5061030360048036038101906102fe9190611582565b610bf7565b005b348015610310575f80fd5b50610319610ded565b60405161032691906111d8565b60405180910390f35b34801561033a575f80fd5b5061035560048036038101906103509190611155565b610df3565b005b5f8181548110610365575f80fd5b905f5260205f20015f915054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b60015f3373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f9054906101000a900460ff1661041a576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161041190611648565b60405180910390fd5b806004805490508110610462576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610459906116b0565b60405180910390fd5b8160048181548110610477576104766116ce565b5b905f5260205f2090600502016003015f9054906101000a900460ff16156104d3576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016104ca90611745565b60405180910390fd5b5f600484815481106104e8576104e76116ce565b5b905f5260205f209060050201905060035f8581526020019081526020015f205f3373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f9054906101000a900460ff1661058e576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610585906117ad565b60405180910390fd5b6001816004015f8282546105a291906117f8565b925050819055505f60035f8681526020019081526020015f205f3373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f6101000a81548160ff021916908315150217905550833373ffffffffffffffffffffffffffffffffffffffff167ff0dca620e2e81f7841d07bcc105e1704fb01475b278a9d4c236e1c62945edd5560405160405180910390a350505050565b5f600480549050905090565b6001602052805f5260405f205f915054906101000a900460ff1681565b5f8060605f805f6004878154811061069a576106996116ce565b5b905f5260205f2090600502019050805f015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff16816001015482600201836003015f9054906101000a900460ff1684600401548280546106f690611858565b80601f016020809104026020016040519081016040528092919081815260200182805461072290611858565b801561076d5780601f106107445761010080835404028352916020019161076d565b820191905f5260205f20905b81548152906001019060200180831161075057829003601f168201915b50505050509250955095509550955095505091939590929450565b6003602052815f5260405f20602052805f5260405f205f915091509054906101000a900460ff1681565b600481815481106107c1575f80fd5b905f5260205f2090600502015f91509050805f015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff169080600101549080600201805461080b90611858565b80601f016020809104026020016040519081016040528092919081815260200182805461083790611858565b80156108825780601f1061085957610100808354040283529160200191610882565b820191905f5260205f20905b81548152906001019060200180831161086557829003601f168201915b505050505090806003015f9054906101000a900460ff16908060040154905085565b60605f80548060200260200160405190810160405280929190818152602001828054801561092457602002820191905f5260205f20905b815f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190600101908083116108db575b5050505050905090565b60015f3373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f9054906101000a900460ff166109b7576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016109ae90611648565b60405180910390fd5b8060048054905081106109ff576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016109f6906116b0565b60405180910390fd5b8160048181548110610a1457610a136116ce565b5b905f5260205f2090600502016003015f9054906101000a900460ff1615610a70576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610a6790611745565b60405180910390fd5b8260035f8281526020019081526020015f205f3373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f9054906101000a900460ff1615610b0a576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610b01906118d2565b60405180910390fd5b5f60048581548110610b1f57610b1e6116ce565b5b905f5260205f20906005020190506001816004015f828254610b4191906118f0565b92505081905550600160035f8781526020019081526020015f205f3373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f6101000a81548160ff021916908315150217905550843373ffffffffffffffffffffffffffffffffffffffff167f5cbe105e36805f7820e291f799d5794ff948af2a5f664e580382defb6339004160405160405180910390a35050505050565b60015f3373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f9054906101000a900460ff16610c80576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610c7790611648565b60405180910390fd5b5f600480549050905060046040518060a001604052808673ffffffffffffffffffffffffffffffffffffffff1681526020018581526020018481526020015f151581526020015f815250908060018154018082558091505060019003905f5260205f2090600502015f909190919091505f820151815f015f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550602082015181600101556040820151816002019081610d539190611ac0565b506060820151816003015f6101000a81548160ff0219169083151502179055506080820151816004015550508373ffffffffffffffffffffffffffffffffffffffff16813373ffffffffffffffffffffffffffffffffffffffff167fd5a05bf70715ad82a09a756320284a1b54c9ff74cd0f8cce6219e79b563fe59d8686604051610ddf929190611b8f565b60405180910390a450505050565b60025481565b60015f3373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f9054906101000a900460ff16610e7c576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610e7390611648565b60405180910390fd5b806004805490508110610ec4576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610ebb906116b0565b60405180910390fd5b8160048181548110610ed957610ed86116ce565b5b905f5260205f2090600502016003015f9054906101000a900460ff1615610f35576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610f2c90611745565b60405180910390fd5b5f60048481548110610f4a57610f496116ce565b5b905f5260205f209060050201905060025481600401541015610fa1576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610f9890611c07565b60405180910390fd5b6001816003015f6101000a81548160ff0219169083151502179055505f815f015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1682600101548360020160405161100d9190611caf565b5f6040518083038185875af1925050503d805f8114611047576040519150601f19603f3d011682016040523d82523d5f602084013e61104c565b606091505b5050905080611090576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161108790611d0f565b60405180910390fd5b843373ffffffffffffffffffffffffffffffffffffffff167f5445f318f4f5fcfb66592e68e0cc5822aa15664039bd5f0ffde24c5a8142b1ac60405160405180910390a35050505050565b5f819050919050565b6110ed816110db565b82525050565b5f6040820190506111065f8301856110e4565b61111360208301846110e4565b9392505050565b5f604051905090565b5f80fd5b5f80fd5b611134816110db565b811461113e575f80fd5b50565b5f8135905061114f8161112b565b92915050565b5f6020828403121561116a57611169611123565b5b5f61117784828501611141565b91505092915050565b5f73ffffffffffffffffffffffffffffffffffffffff82169050919050565b5f6111a982611180565b9050919050565b6111b98161119f565b82525050565b5f6020820190506111d25f8301846111b0565b92915050565b5f6020820190506111eb5f8301846110e4565b92915050565b6111fa8161119f565b8114611204575f80fd5b50565b5f81359050611215816111f1565b92915050565b5f602082840312156112305761122f611123565b5b5f61123d84828501611207565b91505092915050565b5f8115159050919050565b61125a81611246565b82525050565b5f6020820190506112735f830184611251565b92915050565b5f81519050919050565b5f82825260208201905092915050565b8281835e5f83830152505050565b5f601f19601f8301169050919050565b5f6112bb82611279565b6112c58185611283565b93506112d5818560208601611293565b6112de816112a1565b840191505092915050565b5f60a0820190506112fc5f8301886111b0565b61130960208301876110e4565b818103604083015261131b81866112b1565b905061132a6060830185611251565b61133760808301846110e4565b9695505050505050565b5f806040838503121561135757611356611123565b5b5f61136485828601611141565b925050602061137585828601611207565b9150509250929050565b5f81519050919050565b5f82825260208201905092915050565b5f819050602082019050919050565b6113b18161119f565b82525050565b5f6113c283836113a8565b60208301905092915050565b5f602082019050919050565b5f6113e48261137f565b6113ee8185611389565b93506113f983611399565b805f5b8381101561142957815161141088826113b7565b975061141b836113ce565b9250506001810190506113fc565b5085935050505092915050565b5f6020820190508181035f83015261144e81846113da565b905092915050565b5f80fd5b5f80fd5b7f4e487b71000000000000000000000000000000000000000000000000000000005f52604160045260245ffd5b611494826112a1565b810181811067ffffffffffffffff821117156114b3576114b261145e565b5b80604052505050565b5f6114c561111a565b90506114d1828261148b565b919050565b5f67ffffffffffffffff8211156114f0576114ef61145e565b5b6114f9826112a1565b9050602081019050919050565b828183375f83830152505050565b5f611526611521846114d6565b6114bc565b9050828152602081018484840111156115425761154161145a565b5b61154d848285611506565b509392505050565b5f82601f83011261156957611568611456565b5b8135611579848260208601611514565b91505092915050565b5f805f6060848603121561159957611598611123565b5b5f6115a686828701611207565b93505060206115b786828701611141565b925050604084013567ffffffffffffffff8111156115d8576115d7611127565b5b6115e486828701611555565b9150509250925092565b5f82825260208201905092915050565b7f6e6f74206f776e657200000000000000000000000000000000000000000000005f82015250565b5f6116326009836115ee565b915061163d826115fe565b602082019050919050565b5f6020820190508181035f83015261165f81611626565b9050919050565b7f747820646f6573206e6f742065786973740000000000000000000000000000005f82015250565b5f61169a6011836115ee565b91506116a582611666565b602082019050919050565b5f6020820190508181035f8301526116c78161168e565b9050919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52603260045260245ffd5b7f747820616c7265616479206578656375746564000000000000000000000000005f82015250565b5f61172f6013836115ee565b915061173a826116fb565b602082019050919050565b5f6020820190508181035f83015261175c81611723565b9050919050565b7f7478206e6f7420636f6e6669726d6564000000000000000000000000000000005f82015250565b5f6117976010836115ee565b91506117a282611763565b602082019050919050565b5f6020820190508181035f8301526117c48161178b565b9050919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52601160045260245ffd5b5f611802826110db565b915061180d836110db565b9250828203905081811115611825576118246117cb565b5b92915050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52602260045260245ffd5b5f600282049050600182168061186f57607f821691505b6020821081036118825761188161182b565b5b50919050565b7f747820616c726561647920636f6e6669726d65640000000000000000000000005f82015250565b5f6118bc6014836115ee565b91506118c782611888565b602082019050919050565b5f6020820190508181035f8301526118e9816118b0565b9050919050565b5f6118fa826110db565b9150611905836110db565b925082820190508082111561191d5761191c6117cb565b5b92915050565b5f819050815f5260205f209050919050565b5f6020601f8301049050919050565b5f82821b905092915050565b5f6008830261197f7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82611944565b6119898683611944565b95508019841693508086168417925050509392505050565b5f819050919050565b5f6119c46119bf6119ba846110db565b6119a1565b6110db565b9050919050565b5f819050919050565b6119dd836119aa565b6119f16119e9826119cb565b848454611950565b825550505050565b5f90565b611a056119f9565b611a108184846119d4565b505050565b5b81811015611a3357611a285f826119fd565b600181019050611a16565b5050565b601f821115611a7857611a4981611923565b611a5284611935565b81016020851015611a61578190505b611a75611a6d85611935565b830182611a15565b50505b505050565b5f82821c905092915050565b5f611a985f1984600802611a7d565b1980831691505092915050565b5f611ab08383611a89565b9150826002028217905092915050565b611ac982611279565b67ffffffffffffffff811115611ae257611ae161145e565b5b611aec8254611858565b611af7828285611a37565b5f60209050601f831160018114611b28575f8415611b16578287015190505b611b208582611aa5565b865550611b87565b601f198416611b3686611923565b5f5b82811015611b5d57848901518255600182019150602085019450602081019050611b38565b86831015611b7a5784890151611b76601f891682611a89565b8355505b6001600288020188555050505b505050505050565b5f604082019050611ba25f8301856110e4565b8181036020830152611bb481846112b1565b90509392505050565b7f63616e6e6f7420657865637574652074780000000000000000000000000000005f82015250565b5f611bf16011836115ee565b9150611bfc82611bbd565b602082019050919050565b5f6020820190508181035f830152611c1e81611be5565b9050919050565b5f81905092915050565b5f8154611c3b81611858565b611c458186611c25565b9450600182165f8114611c5f5760018114611c7457611ca6565b60ff1983168652811515820286019350611ca6565b611c7d85611923565b5f5b83811015611c9e57815481890152600182019150602081019050611c7f565b838801955050505b50505092915050565b5f611cba8284611c2f565b915081905092915050565b7f7478206661696c656400000000000000000000000000000000000000000000005f82015250565b5f611cf96009836115ee565b9150611d0482611cc5565b602082019050919050565b5f6020820190508181035f830152611d2681611ced565b905091905056fea2646970667358221220f19e388c9c1947e9eef610118b95b4c9b532ad359118f0eb92a747acb19b164a64736f6c634300081a0033",
}

// MultiSigABI is the input ABI used to generate the binding from.
// Deprecated: Use MultiSigMetaData.ABI instead.
var MultiSigABI = MultiSigMetaData.ABI

// MultiSigBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use MultiSigMetaData.Bin instead.
var MultiSigBin = MultiSigMetaData.Bin

// DeployMultiSig deploys a new Ethereum contract, binding an instance of MultiSig to it.
func DeployMultiSig(auth *bind.TransactOpts, backend bind.ContractBackend, _owners []common.Address, _numConfirmationsRequired *big.Int) (common.Address, *types.Transaction, *MultiSig, error) {
	parsed, err := MultiSigMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(MultiSigBin), backend, _owners, _numConfirmationsRequired)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &MultiSig{MultiSigCaller: MultiSigCaller{contract: contract}, MultiSigTransactor: MultiSigTransactor{contract: contract}, MultiSigFilterer: MultiSigFilterer{contract: contract}}, nil
}

// MultiSig is an auto generated Go binding around an Ethereum contract.
type MultiSig struct {
	MultiSigCaller     // Read-only binding to the contract
	MultiSigTransactor // Write-only binding to the contract
	MultiSigFilterer   // Log filterer for contract events
}

// MultiSigCaller is an auto generated read-only Go binding around an Ethereum contract.
type MultiSigCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MultiSigTransactor is an auto generated write-only Go binding around an Ethereum contract.
type MultiSigTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MultiSigFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type MultiSigFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MultiSigSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type MultiSigSession struct {
	Contract     *MultiSig         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MultiSigCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type MultiSigCallerSession struct {
	Contract *MultiSigCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// MultiSigTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type MultiSigTransactorSession struct {
	Contract     *MultiSigTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// MultiSigRaw is an auto generated low-level Go binding around an Ethereum contract.
type MultiSigRaw struct {
	Contract *MultiSig // Generic contract binding to access the raw methods on
}

// MultiSigCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type MultiSigCallerRaw struct {
	Contract *MultiSigCaller // Generic read-only contract binding to access the raw methods on
}

// MultiSigTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type MultiSigTransactorRaw struct {
	Contract *MultiSigTransactor // Generic write-only contract binding to access the raw methods on
}

// NewMultiSig creates a new instance of MultiSig, bound to a specific deployed contract.
func NewMultiSig(address common.Address, backend bind.ContractBackend) (*MultiSig, error) {
	contract, err := bindMultiSig(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &MultiSig{MultiSigCaller: MultiSigCaller{contract: contract}, MultiSigTransactor: MultiSigTransactor{contract: contract}, MultiSigFilterer: MultiSigFilterer{contract: contract}}, nil
}

// NewMultiSigCaller creates a new read-only instance of MultiSig, bound to a specific deployed contract.
func NewMultiSigCaller(address common.Address, caller bind.ContractCaller) (*MultiSigCaller, error) {
	contract, err := bindMultiSig(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &MultiSigCaller{contract: contract}, nil
}

// NewMultiSigTransactor creates a new write-only instance of MultiSig, bound to a specific deployed contract.
func NewMultiSigTransactor(address common.Address, transactor bind.ContractTransactor) (*MultiSigTransactor, error) {
	contract, err := bindMultiSig(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &MultiSigTransactor{contract: contract}, nil
}

// NewMultiSigFilterer creates a new log filterer instance of MultiSig, bound to a specific deployed contract.
func NewMultiSigFilterer(address common.Address, filterer bind.ContractFilterer) (*MultiSigFilterer, error) {
	contract, err := bindMultiSig(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &MultiSigFilterer{contract: contract}, nil
}

// bindMultiSig binds a generic wrapper to an already deployed contract.
func bindMultiSig(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := MultiSigMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MultiSig *MultiSigRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MultiSig.Contract.MultiSigCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MultiSig *MultiSigRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MultiSig.Contract.MultiSigTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MultiSig *MultiSigRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MultiSig.Contract.MultiSigTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MultiSig *MultiSigCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MultiSig.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MultiSig *MultiSigTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MultiSig.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MultiSig *MultiSigTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MultiSig.Contract.contract.Transact(opts, method, params...)
}

// GetOwners is a free data retrieval call binding the contract method 0xa0e67e2b.
//
// Solidity: function getOwners() view returns(address[])
func (_MultiSig *MultiSigCaller) GetOwners(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _MultiSig.contract.Call(opts, &out, "getOwners")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetOwners is a free data retrieval call binding the contract method 0xa0e67e2b.
//
// Solidity: function getOwners() view returns(address[])
func (_MultiSig *MultiSigSession) GetOwners() ([]common.Address, error) {
	return _MultiSig.Contract.GetOwners(&_MultiSig.CallOpts)
}

// GetOwners is a free data retrieval call binding the contract method 0xa0e67e2b.
//
// Solidity: function getOwners() view returns(address[])
func (_MultiSig *MultiSigCallerSession) GetOwners() ([]common.Address, error) {
	return _MultiSig.Contract.GetOwners(&_MultiSig.CallOpts)
}

// GetTransaction is a free data retrieval call binding the contract method 0x33ea3dc8.
//
// Solidity: function getTransaction(uint256 _txIndex) view returns(address to, uint256 value, bytes data, bool executed, uint256 numConfirmations)
func (_MultiSig *MultiSigCaller) GetTransaction(opts *bind.CallOpts, _txIndex *big.Int) (struct {
	To               common.Address
	Value            *big.Int
	Data             []byte
	Executed         bool
	NumConfirmations *big.Int
}, error) {
	var out []interface{}
	err := _MultiSig.contract.Call(opts, &out, "getTransaction", _txIndex)

	outstruct := new(struct {
		To               common.Address
		Value            *big.Int
		Data             []byte
		Executed         bool
		NumConfirmations *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.To = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.Value = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.Data = *abi.ConvertType(out[2], new([]byte)).(*[]byte)
	outstruct.Executed = *abi.ConvertType(out[3], new(bool)).(*bool)
	outstruct.NumConfirmations = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// GetTransaction is a free data retrieval call binding the contract method 0x33ea3dc8.
//
// Solidity: function getTransaction(uint256 _txIndex) view returns(address to, uint256 value, bytes data, bool executed, uint256 numConfirmations)
func (_MultiSig *MultiSigSession) GetTransaction(_txIndex *big.Int) (struct {
	To               common.Address
	Value            *big.Int
	Data             []byte
	Executed         bool
	NumConfirmations *big.Int
}, error) {
	return _MultiSig.Contract.GetTransaction(&_MultiSig.CallOpts, _txIndex)
}

// GetTransaction is a free data retrieval call binding the contract method 0x33ea3dc8.
//
// Solidity: function getTransaction(uint256 _txIndex) view returns(address to, uint256 value, bytes data, bool executed, uint256 numConfirmations)
func (_MultiSig *MultiSigCallerSession) GetTransaction(_txIndex *big.Int) (struct {
	To               common.Address
	Value            *big.Int
	Data             []byte
	Executed         bool
	NumConfirmations *big.Int
}, error) {
	return _MultiSig.Contract.GetTransaction(&_MultiSig.CallOpts, _txIndex)
}

// GetTransactionCount is a free data retrieval call binding the contract method 0x2e7700f0.
//
// Solidity: function getTransactionCount() view returns(uint256)
func (_MultiSig *MultiSigCaller) GetTransactionCount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _MultiSig.contract.Call(opts, &out, "getTransactionCount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetTransactionCount is a free data retrieval call binding the contract method 0x2e7700f0.
//
// Solidity: function getTransactionCount() view returns(uint256)
func (_MultiSig *MultiSigSession) GetTransactionCount() (*big.Int, error) {
	return _MultiSig.Contract.GetTransactionCount(&_MultiSig.CallOpts)
}

// GetTransactionCount is a free data retrieval call binding the contract method 0x2e7700f0.
//
// Solidity: function getTransactionCount() view returns(uint256)
func (_MultiSig *MultiSigCallerSession) GetTransactionCount() (*big.Int, error) {
	return _MultiSig.Contract.GetTransactionCount(&_MultiSig.CallOpts)
}

// IsConfirmed is a free data retrieval call binding the contract method 0x80f59a65.
//
// Solidity: function isConfirmed(uint256 , address ) view returns(bool)
func (_MultiSig *MultiSigCaller) IsConfirmed(opts *bind.CallOpts, arg0 *big.Int, arg1 common.Address) (bool, error) {
	var out []interface{}
	err := _MultiSig.contract.Call(opts, &out, "isConfirmed", arg0, arg1)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsConfirmed is a free data retrieval call binding the contract method 0x80f59a65.
//
// Solidity: function isConfirmed(uint256 , address ) view returns(bool)
func (_MultiSig *MultiSigSession) IsConfirmed(arg0 *big.Int, arg1 common.Address) (bool, error) {
	return _MultiSig.Contract.IsConfirmed(&_MultiSig.CallOpts, arg0, arg1)
}

// IsConfirmed is a free data retrieval call binding the contract method 0x80f59a65.
//
// Solidity: function isConfirmed(uint256 , address ) view returns(bool)
func (_MultiSig *MultiSigCallerSession) IsConfirmed(arg0 *big.Int, arg1 common.Address) (bool, error) {
	return _MultiSig.Contract.IsConfirmed(&_MultiSig.CallOpts, arg0, arg1)
}

// IsOwner is a free data retrieval call binding the contract method 0x2f54bf6e.
//
// Solidity: function isOwner(address ) view returns(bool)
func (_MultiSig *MultiSigCaller) IsOwner(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _MultiSig.contract.Call(opts, &out, "isOwner", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsOwner is a free data retrieval call binding the contract method 0x2f54bf6e.
//
// Solidity: function isOwner(address ) view returns(bool)
func (_MultiSig *MultiSigSession) IsOwner(arg0 common.Address) (bool, error) {
	return _MultiSig.Contract.IsOwner(&_MultiSig.CallOpts, arg0)
}

// IsOwner is a free data retrieval call binding the contract method 0x2f54bf6e.
//
// Solidity: function isOwner(address ) view returns(bool)
func (_MultiSig *MultiSigCallerSession) IsOwner(arg0 common.Address) (bool, error) {
	return _MultiSig.Contract.IsOwner(&_MultiSig.CallOpts, arg0)
}

// NumConfirmationsRequired is a free data retrieval call binding the contract method 0xd0549b85.
//
// Solidity: function numConfirmationsRequired() view returns(uint256)
func (_MultiSig *MultiSigCaller) NumConfirmationsRequired(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _MultiSig.contract.Call(opts, &out, "numConfirmationsRequired")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// NumConfirmationsRequired is a free data retrieval call binding the contract method 0xd0549b85.
//
// Solidity: function numConfirmationsRequired() view returns(uint256)
func (_MultiSig *MultiSigSession) NumConfirmationsRequired() (*big.Int, error) {
	return _MultiSig.Contract.NumConfirmationsRequired(&_MultiSig.CallOpts)
}

// NumConfirmationsRequired is a free data retrieval call binding the contract method 0xd0549b85.
//
// Solidity: function numConfirmationsRequired() view returns(uint256)
func (_MultiSig *MultiSigCallerSession) NumConfirmationsRequired() (*big.Int, error) {
	return _MultiSig.Contract.NumConfirmationsRequired(&_MultiSig.CallOpts)
}

// Owners is a free data retrieval call binding the contract method 0x025e7c27.
//
// Solidity: function owners(uint256 ) view returns(address)
func (_MultiSig *MultiSigCaller) Owners(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _MultiSig.contract.Call(opts, &out, "owners", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owners is a free data retrieval call binding the contract method 0x025e7c27.
//
// Solidity: function owners(uint256 ) view returns(address)
func (_MultiSig *MultiSigSession) Owners(arg0 *big.Int) (common.Address, error) {
	return _MultiSig.Contract.Owners(&_MultiSig.CallOpts, arg0)
}

// Owners is a free data retrieval call binding the contract method 0x025e7c27.
//
// Solidity: function owners(uint256 ) view returns(address)
func (_MultiSig *MultiSigCallerSession) Owners(arg0 *big.Int) (common.Address, error) {
	return _MultiSig.Contract.Owners(&_MultiSig.CallOpts, arg0)
}

// Transactions is a free data retrieval call binding the contract method 0x9ace38c2.
//
// Solidity: function transactions(uint256 ) view returns(address to, uint256 value, bytes data, bool executed, uint256 numConfirmations)
func (_MultiSig *MultiSigCaller) Transactions(opts *bind.CallOpts, arg0 *big.Int) (struct {
	To               common.Address
	Value            *big.Int
	Data             []byte
	Executed         bool
	NumConfirmations *big.Int
}, error) {
	var out []interface{}
	err := _MultiSig.contract.Call(opts, &out, "transactions", arg0)

	outstruct := new(struct {
		To               common.Address
		Value            *big.Int
		Data             []byte
		Executed         bool
		NumConfirmations *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.To = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.Value = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.Data = *abi.ConvertType(out[2], new([]byte)).(*[]byte)
	outstruct.Executed = *abi.ConvertType(out[3], new(bool)).(*bool)
	outstruct.NumConfirmations = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// Transactions is a free data retrieval call binding the contract method 0x9ace38c2.
//
// Solidity: function transactions(uint256 ) view returns(address to, uint256 value, bytes data, bool executed, uint256 numConfirmations)
func (_MultiSig *MultiSigSession) Transactions(arg0 *big.Int) (struct {
	To               common.Address
	Value            *big.Int
	Data             []byte
	Executed         bool
	NumConfirmations *big.Int
}, error) {
	return _MultiSig.Contract.Transactions(&_MultiSig.CallOpts, arg0)
}

// Transactions is a free data retrieval call binding the contract method 0x9ace38c2.
//
// Solidity: function transactions(uint256 ) view returns(address to, uint256 value, bytes data, bool executed, uint256 numConfirmations)
func (_MultiSig *MultiSigCallerSession) Transactions(arg0 *big.Int) (struct {
	To               common.Address
	Value            *big.Int
	Data             []byte
	Executed         bool
	NumConfirmations *big.Int
}, error) {
	return _MultiSig.Contract.Transactions(&_MultiSig.CallOpts, arg0)
}

// ConfirmTransaction is a paid mutator transaction binding the contract method 0xc01a8c84.
//
// Solidity: function confirmTransaction(uint256 _txIndex) returns()
func (_MultiSig *MultiSigTransactor) ConfirmTransaction(opts *bind.TransactOpts, _txIndex *big.Int) (*types.Transaction, error) {
	return _MultiSig.contract.Transact(opts, "confirmTransaction", _txIndex)
}

// ConfirmTransaction is a paid mutator transaction binding the contract method 0xc01a8c84.
//
// Solidity: function confirmTransaction(uint256 _txIndex) returns()
func (_MultiSig *MultiSigSession) ConfirmTransaction(_txIndex *big.Int) (*types.Transaction, error) {
	return _MultiSig.Contract.ConfirmTransaction(&_MultiSig.TransactOpts, _txIndex)
}

// ConfirmTransaction is a paid mutator transaction binding the contract method 0xc01a8c84.
//
// Solidity: function confirmTransaction(uint256 _txIndex) returns()
func (_MultiSig *MultiSigTransactorSession) ConfirmTransaction(_txIndex *big.Int) (*types.Transaction, error) {
	return _MultiSig.Contract.ConfirmTransaction(&_MultiSig.TransactOpts, _txIndex)
}

// ExecuteTransaction is a paid mutator transaction binding the contract method 0xee22610b.
//
// Solidity: function executeTransaction(uint256 _txIndex) returns()
func (_MultiSig *MultiSigTransactor) ExecuteTransaction(opts *bind.TransactOpts, _txIndex *big.Int) (*types.Transaction, error) {
	return _MultiSig.contract.Transact(opts, "executeTransaction", _txIndex)
}

// ExecuteTransaction is a paid mutator transaction binding the contract method 0xee22610b.
//
// Solidity: function executeTransaction(uint256 _txIndex) returns()
func (_MultiSig *MultiSigSession) ExecuteTransaction(_txIndex *big.Int) (*types.Transaction, error) {
	return _MultiSig.Contract.ExecuteTransaction(&_MultiSig.TransactOpts, _txIndex)
}

// ExecuteTransaction is a paid mutator transaction binding the contract method 0xee22610b.
//
// Solidity: function executeTransaction(uint256 _txIndex) returns()
func (_MultiSig *MultiSigTransactorSession) ExecuteTransaction(_txIndex *big.Int) (*types.Transaction, error) {
	return _MultiSig.Contract.ExecuteTransaction(&_MultiSig.TransactOpts, _txIndex)
}

// RevokeConfirmation is a paid mutator transaction binding the contract method 0x20ea8d86.
//
// Solidity: function revokeConfirmation(uint256 _txIndex) returns()
func (_MultiSig *MultiSigTransactor) RevokeConfirmation(opts *bind.TransactOpts, _txIndex *big.Int) (*types.Transaction, error) {
	return _MultiSig.contract.Transact(opts, "revokeConfirmation", _txIndex)
}

// RevokeConfirmation is a paid mutator transaction binding the contract method 0x20ea8d86.
//
// Solidity: function revokeConfirmation(uint256 _txIndex) returns()
func (_MultiSig *MultiSigSession) RevokeConfirmation(_txIndex *big.Int) (*types.Transaction, error) {
	return _MultiSig.Contract.RevokeConfirmation(&_MultiSig.TransactOpts, _txIndex)
}

// RevokeConfirmation is a paid mutator transaction binding the contract method 0x20ea8d86.
//
// Solidity: function revokeConfirmation(uint256 _txIndex) returns()
func (_MultiSig *MultiSigTransactorSession) RevokeConfirmation(_txIndex *big.Int) (*types.Transaction, error) {
	return _MultiSig.Contract.RevokeConfirmation(&_MultiSig.TransactOpts, _txIndex)
}

// SubmitTransaction is a paid mutator transaction binding the contract method 0xc6427474.
//
// Solidity: function submitTransaction(address _to, uint256 _value, bytes _data) returns()
func (_MultiSig *MultiSigTransactor) SubmitTransaction(opts *bind.TransactOpts, _to common.Address, _value *big.Int, _data []byte) (*types.Transaction, error) {
	return _MultiSig.contract.Transact(opts, "submitTransaction", _to, _value, _data)
}

// SubmitTransaction is a paid mutator transaction binding the contract method 0xc6427474.
//
// Solidity: function submitTransaction(address _to, uint256 _value, bytes _data) returns()
func (_MultiSig *MultiSigSession) SubmitTransaction(_to common.Address, _value *big.Int, _data []byte) (*types.Transaction, error) {
	return _MultiSig.Contract.SubmitTransaction(&_MultiSig.TransactOpts, _to, _value, _data)
}

// SubmitTransaction is a paid mutator transaction binding the contract method 0xc6427474.
//
// Solidity: function submitTransaction(address _to, uint256 _value, bytes _data) returns()
func (_MultiSig *MultiSigTransactorSession) SubmitTransaction(_to common.Address, _value *big.Int, _data []byte) (*types.Transaction, error) {
	return _MultiSig.Contract.SubmitTransaction(&_MultiSig.TransactOpts, _to, _value, _data)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_MultiSig *MultiSigTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MultiSig.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_MultiSig *MultiSigSession) Receive() (*types.Transaction, error) {
	return _MultiSig.Contract.Receive(&_MultiSig.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_MultiSig *MultiSigTransactorSession) Receive() (*types.Transaction, error) {
	return _MultiSig.Contract.Receive(&_MultiSig.TransactOpts)
}

// MultiSigConfirmTransactionIterator is returned from FilterConfirmTransaction and is used to iterate over the raw logs and unpacked data for ConfirmTransaction events raised by the MultiSig contract.
type MultiSigConfirmTransactionIterator struct {
	Event *MultiSigConfirmTransaction // Event containing the contract specifics and raw log

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
func (it *MultiSigConfirmTransactionIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MultiSigConfirmTransaction)
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
		it.Event = new(MultiSigConfirmTransaction)
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
func (it *MultiSigConfirmTransactionIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MultiSigConfirmTransactionIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MultiSigConfirmTransaction represents a ConfirmTransaction event raised by the MultiSig contract.
type MultiSigConfirmTransaction struct {
	Owner   common.Address
	TxIndex *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterConfirmTransaction is a free log retrieval operation binding the contract event 0x5cbe105e36805f7820e291f799d5794ff948af2a5f664e580382defb63390041.
//
// Solidity: event ConfirmTransaction(address indexed owner, uint256 indexed txIndex)
func (_MultiSig *MultiSigFilterer) FilterConfirmTransaction(opts *bind.FilterOpts, owner []common.Address, txIndex []*big.Int) (*MultiSigConfirmTransactionIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var txIndexRule []interface{}
	for _, txIndexItem := range txIndex {
		txIndexRule = append(txIndexRule, txIndexItem)
	}

	logs, sub, err := _MultiSig.contract.FilterLogs(opts, "ConfirmTransaction", ownerRule, txIndexRule)
	if err != nil {
		return nil, err
	}
	return &MultiSigConfirmTransactionIterator{contract: _MultiSig.contract, event: "ConfirmTransaction", logs: logs, sub: sub}, nil
}

// WatchConfirmTransaction is a free log subscription operation binding the contract event 0x5cbe105e36805f7820e291f799d5794ff948af2a5f664e580382defb63390041.
//
// Solidity: event ConfirmTransaction(address indexed owner, uint256 indexed txIndex)
func (_MultiSig *MultiSigFilterer) WatchConfirmTransaction(opts *bind.WatchOpts, sink chan<- *MultiSigConfirmTransaction, owner []common.Address, txIndex []*big.Int) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var txIndexRule []interface{}
	for _, txIndexItem := range txIndex {
		txIndexRule = append(txIndexRule, txIndexItem)
	}

	logs, sub, err := _MultiSig.contract.WatchLogs(opts, "ConfirmTransaction", ownerRule, txIndexRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MultiSigConfirmTransaction)
				if err := _MultiSig.contract.UnpackLog(event, "ConfirmTransaction", log); err != nil {
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

// ParseConfirmTransaction is a log parse operation binding the contract event 0x5cbe105e36805f7820e291f799d5794ff948af2a5f664e580382defb63390041.
//
// Solidity: event ConfirmTransaction(address indexed owner, uint256 indexed txIndex)
func (_MultiSig *MultiSigFilterer) ParseConfirmTransaction(log types.Log) (*MultiSigConfirmTransaction, error) {
	event := new(MultiSigConfirmTransaction)
	if err := _MultiSig.contract.UnpackLog(event, "ConfirmTransaction", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MultiSigDepositIterator is returned from FilterDeposit and is used to iterate over the raw logs and unpacked data for Deposit events raised by the MultiSig contract.
type MultiSigDepositIterator struct {
	Event *MultiSigDeposit // Event containing the contract specifics and raw log

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
func (it *MultiSigDepositIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MultiSigDeposit)
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
		it.Event = new(MultiSigDeposit)
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
func (it *MultiSigDepositIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MultiSigDepositIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MultiSigDeposit represents a Deposit event raised by the MultiSig contract.
type MultiSigDeposit struct {
	Sender  common.Address
	Amount  *big.Int
	Balance *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterDeposit is a free log retrieval operation binding the contract event 0x90890809c654f11d6e72a28fa60149770a0d11ec6c92319d6ceb2bb0a4ea1a15.
//
// Solidity: event Deposit(address indexed sender, uint256 amount, uint256 balance)
func (_MultiSig *MultiSigFilterer) FilterDeposit(opts *bind.FilterOpts, sender []common.Address) (*MultiSigDepositIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _MultiSig.contract.FilterLogs(opts, "Deposit", senderRule)
	if err != nil {
		return nil, err
	}
	return &MultiSigDepositIterator{contract: _MultiSig.contract, event: "Deposit", logs: logs, sub: sub}, nil
}

// WatchDeposit is a free log subscription operation binding the contract event 0x90890809c654f11d6e72a28fa60149770a0d11ec6c92319d6ceb2bb0a4ea1a15.
//
// Solidity: event Deposit(address indexed sender, uint256 amount, uint256 balance)
func (_MultiSig *MultiSigFilterer) WatchDeposit(opts *bind.WatchOpts, sink chan<- *MultiSigDeposit, sender []common.Address) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _MultiSig.contract.WatchLogs(opts, "Deposit", senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MultiSigDeposit)
				if err := _MultiSig.contract.UnpackLog(event, "Deposit", log); err != nil {
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

// ParseDeposit is a log parse operation binding the contract event 0x90890809c654f11d6e72a28fa60149770a0d11ec6c92319d6ceb2bb0a4ea1a15.
//
// Solidity: event Deposit(address indexed sender, uint256 amount, uint256 balance)
func (_MultiSig *MultiSigFilterer) ParseDeposit(log types.Log) (*MultiSigDeposit, error) {
	event := new(MultiSigDeposit)
	if err := _MultiSig.contract.UnpackLog(event, "Deposit", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MultiSigExecuteTransactionIterator is returned from FilterExecuteTransaction and is used to iterate over the raw logs and unpacked data for ExecuteTransaction events raised by the MultiSig contract.
type MultiSigExecuteTransactionIterator struct {
	Event *MultiSigExecuteTransaction // Event containing the contract specifics and raw log

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
func (it *MultiSigExecuteTransactionIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MultiSigExecuteTransaction)
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
		it.Event = new(MultiSigExecuteTransaction)
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
func (it *MultiSigExecuteTransactionIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MultiSigExecuteTransactionIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MultiSigExecuteTransaction represents a ExecuteTransaction event raised by the MultiSig contract.
type MultiSigExecuteTransaction struct {
	Owner   common.Address
	TxIndex *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterExecuteTransaction is a free log retrieval operation binding the contract event 0x5445f318f4f5fcfb66592e68e0cc5822aa15664039bd5f0ffde24c5a8142b1ac.
//
// Solidity: event ExecuteTransaction(address indexed owner, uint256 indexed txIndex)
func (_MultiSig *MultiSigFilterer) FilterExecuteTransaction(opts *bind.FilterOpts, owner []common.Address, txIndex []*big.Int) (*MultiSigExecuteTransactionIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var txIndexRule []interface{}
	for _, txIndexItem := range txIndex {
		txIndexRule = append(txIndexRule, txIndexItem)
	}

	logs, sub, err := _MultiSig.contract.FilterLogs(opts, "ExecuteTransaction", ownerRule, txIndexRule)
	if err != nil {
		return nil, err
	}
	return &MultiSigExecuteTransactionIterator{contract: _MultiSig.contract, event: "ExecuteTransaction", logs: logs, sub: sub}, nil
}

// WatchExecuteTransaction is a free log subscription operation binding the contract event 0x5445f318f4f5fcfb66592e68e0cc5822aa15664039bd5f0ffde24c5a8142b1ac.
//
// Solidity: event ExecuteTransaction(address indexed owner, uint256 indexed txIndex)
func (_MultiSig *MultiSigFilterer) WatchExecuteTransaction(opts *bind.WatchOpts, sink chan<- *MultiSigExecuteTransaction, owner []common.Address, txIndex []*big.Int) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var txIndexRule []interface{}
	for _, txIndexItem := range txIndex {
		txIndexRule = append(txIndexRule, txIndexItem)
	}

	logs, sub, err := _MultiSig.contract.WatchLogs(opts, "ExecuteTransaction", ownerRule, txIndexRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MultiSigExecuteTransaction)
				if err := _MultiSig.contract.UnpackLog(event, "ExecuteTransaction", log); err != nil {
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

// ParseExecuteTransaction is a log parse operation binding the contract event 0x5445f318f4f5fcfb66592e68e0cc5822aa15664039bd5f0ffde24c5a8142b1ac.
//
// Solidity: event ExecuteTransaction(address indexed owner, uint256 indexed txIndex)
func (_MultiSig *MultiSigFilterer) ParseExecuteTransaction(log types.Log) (*MultiSigExecuteTransaction, error) {
	event := new(MultiSigExecuteTransaction)
	if err := _MultiSig.contract.UnpackLog(event, "ExecuteTransaction", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MultiSigRevokeConfirmationIterator is returned from FilterRevokeConfirmation and is used to iterate over the raw logs and unpacked data for RevokeConfirmation events raised by the MultiSig contract.
type MultiSigRevokeConfirmationIterator struct {
	Event *MultiSigRevokeConfirmation // Event containing the contract specifics and raw log

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
func (it *MultiSigRevokeConfirmationIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MultiSigRevokeConfirmation)
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
		it.Event = new(MultiSigRevokeConfirmation)
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
func (it *MultiSigRevokeConfirmationIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MultiSigRevokeConfirmationIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MultiSigRevokeConfirmation represents a RevokeConfirmation event raised by the MultiSig contract.
type MultiSigRevokeConfirmation struct {
	Owner   common.Address
	TxIndex *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRevokeConfirmation is a free log retrieval operation binding the contract event 0xf0dca620e2e81f7841d07bcc105e1704fb01475b278a9d4c236e1c62945edd55.
//
// Solidity: event RevokeConfirmation(address indexed owner, uint256 indexed txIndex)
func (_MultiSig *MultiSigFilterer) FilterRevokeConfirmation(opts *bind.FilterOpts, owner []common.Address, txIndex []*big.Int) (*MultiSigRevokeConfirmationIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var txIndexRule []interface{}
	for _, txIndexItem := range txIndex {
		txIndexRule = append(txIndexRule, txIndexItem)
	}

	logs, sub, err := _MultiSig.contract.FilterLogs(opts, "RevokeConfirmation", ownerRule, txIndexRule)
	if err != nil {
		return nil, err
	}
	return &MultiSigRevokeConfirmationIterator{contract: _MultiSig.contract, event: "RevokeConfirmation", logs: logs, sub: sub}, nil
}

// WatchRevokeConfirmation is a free log subscription operation binding the contract event 0xf0dca620e2e81f7841d07bcc105e1704fb01475b278a9d4c236e1c62945edd55.
//
// Solidity: event RevokeConfirmation(address indexed owner, uint256 indexed txIndex)
func (_MultiSig *MultiSigFilterer) WatchRevokeConfirmation(opts *bind.WatchOpts, sink chan<- *MultiSigRevokeConfirmation, owner []common.Address, txIndex []*big.Int) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var txIndexRule []interface{}
	for _, txIndexItem := range txIndex {
		txIndexRule = append(txIndexRule, txIndexItem)
	}

	logs, sub, err := _MultiSig.contract.WatchLogs(opts, "RevokeConfirmation", ownerRule, txIndexRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MultiSigRevokeConfirmation)
				if err := _MultiSig.contract.UnpackLog(event, "RevokeConfirmation", log); err != nil {
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

// ParseRevokeConfirmation is a log parse operation binding the contract event 0xf0dca620e2e81f7841d07bcc105e1704fb01475b278a9d4c236e1c62945edd55.
//
// Solidity: event RevokeConfirmation(address indexed owner, uint256 indexed txIndex)
func (_MultiSig *MultiSigFilterer) ParseRevokeConfirmation(log types.Log) (*MultiSigRevokeConfirmation, error) {
	event := new(MultiSigRevokeConfirmation)
	if err := _MultiSig.contract.UnpackLog(event, "RevokeConfirmation", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MultiSigSubmitTransactionIterator is returned from FilterSubmitTransaction and is used to iterate over the raw logs and unpacked data for SubmitTransaction events raised by the MultiSig contract.
type MultiSigSubmitTransactionIterator struct {
	Event *MultiSigSubmitTransaction // Event containing the contract specifics and raw log

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
func (it *MultiSigSubmitTransactionIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MultiSigSubmitTransaction)
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
		it.Event = new(MultiSigSubmitTransaction)
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
func (it *MultiSigSubmitTransactionIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MultiSigSubmitTransactionIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MultiSigSubmitTransaction represents a SubmitTransaction event raised by the MultiSig contract.
type MultiSigSubmitTransaction struct {
	Owner   common.Address
	TxIndex *big.Int
	To      common.Address
	Value   *big.Int
	Data    []byte
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterSubmitTransaction is a free log retrieval operation binding the contract event 0xd5a05bf70715ad82a09a756320284a1b54c9ff74cd0f8cce6219e79b563fe59d.
//
// Solidity: event SubmitTransaction(address indexed owner, uint256 indexed txIndex, address indexed to, uint256 value, bytes data)
func (_MultiSig *MultiSigFilterer) FilterSubmitTransaction(opts *bind.FilterOpts, owner []common.Address, txIndex []*big.Int, to []common.Address) (*MultiSigSubmitTransactionIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var txIndexRule []interface{}
	for _, txIndexItem := range txIndex {
		txIndexRule = append(txIndexRule, txIndexItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _MultiSig.contract.FilterLogs(opts, "SubmitTransaction", ownerRule, txIndexRule, toRule)
	if err != nil {
		return nil, err
	}
	return &MultiSigSubmitTransactionIterator{contract: _MultiSig.contract, event: "SubmitTransaction", logs: logs, sub: sub}, nil
}

// WatchSubmitTransaction is a free log subscription operation binding the contract event 0xd5a05bf70715ad82a09a756320284a1b54c9ff74cd0f8cce6219e79b563fe59d.
//
// Solidity: event SubmitTransaction(address indexed owner, uint256 indexed txIndex, address indexed to, uint256 value, bytes data)
func (_MultiSig *MultiSigFilterer) WatchSubmitTransaction(opts *bind.WatchOpts, sink chan<- *MultiSigSubmitTransaction, owner []common.Address, txIndex []*big.Int, to []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var txIndexRule []interface{}
	for _, txIndexItem := range txIndex {
		txIndexRule = append(txIndexRule, txIndexItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _MultiSig.contract.WatchLogs(opts, "SubmitTransaction", ownerRule, txIndexRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MultiSigSubmitTransaction)
				if err := _MultiSig.contract.UnpackLog(event, "SubmitTransaction", log); err != nil {
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

// ParseSubmitTransaction is a log parse operation binding the contract event 0xd5a05bf70715ad82a09a756320284a1b54c9ff74cd0f8cce6219e79b563fe59d.
//
// Solidity: event SubmitTransaction(address indexed owner, uint256 indexed txIndex, address indexed to, uint256 value, bytes data)
func (_MultiSig *MultiSigFilterer) ParseSubmitTransaction(log types.Log) (*MultiSigSubmitTransaction, error) {
	event := new(MultiSigSubmitTransaction)
	if err := _MultiSig.contract.UnpackLog(event, "SubmitTransaction", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
