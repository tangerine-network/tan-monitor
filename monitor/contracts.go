// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package monitor

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

// GovernanceABI is the input ABI used to generate the binding from.
const GovernanceABI = "[{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"}],\"name\":\"dkgSuccesses\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"notarySetSize\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"MinStake\",\"type\":\"uint256\"},{\"name\":\"LockupPeriod\",\"type\":\"uint256\"},{\"name\":\"MinGasPrice\",\"type\":\"uint256\"},{\"name\":\"BlockGasLimit\",\"type\":\"uint256\"},{\"name\":\"LambdaBA\",\"type\":\"uint256\"},{\"name\":\"LambdaDKG\",\"type\":\"uint256\"},{\"name\":\"NotaryParamAlpha\",\"type\":\"uint256\"},{\"name\":\"NotaryParamBeta\",\"type\":\"uint256\"},{\"name\":\"RoundLength\",\"type\":\"uint256\"},{\"name\":\"MinBlockInterval\",\"type\":\"uint256\"},{\"name\":\"FineValues\",\"type\":\"uint256[]\"}],\"name\":\"updateConfiguration\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"PublicKey\",\"type\":\"bytes\"}],\"name\":\"addDKGMasterPublicKey\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"nodes\",\"outputs\":[{\"name\":\"owner\",\"type\":\"address\"},{\"name\":\"publicKey\",\"type\":\"bytes\"},{\"name\":\"staked\",\"type\":\"uint256\"},{\"name\":\"fined\",\"type\":\"uint256\"},{\"name\":\"name\",\"type\":\"string\"},{\"name\":\"email\",\"type\":\"string\"},{\"name\":\"location\",\"type\":\"string\"},{\"name\":\"url\",\"type\":\"string\"},{\"name\":\"unstaked\",\"type\":\"uint256\"},{\"name\":\"unstakedAt\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"notaryParamBeta\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"MPKReady\",\"type\":\"bytes\"}],\"name\":\"addDKGMPKReady\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"miningVelocity\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"Finalize\",\"type\":\"bytes\"}],\"name\":\"addDKGFinalize\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"lambdaBA\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"Amount\",\"type\":\"uint256\"}],\"name\":\"unstake\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"Type\",\"type\":\"uint256\"},{\"name\":\"Arg1\",\"type\":\"bytes\"},{\"name\":\"Arg2\",\"type\":\"bytes\"}],\"name\":\"report\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"minStake\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"stake\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"withdraw\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"NodeAddress\",\"type\":\"address\"}],\"name\":\"payFine\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"NewPublicKey\",\"type\":\"bytes\"}],\"name\":\"replaceNodePublicKey\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"NewOwner\",\"type\":\"address\"}],\"name\":\"transferNodeOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"crsRound\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"notaryParamAlpha\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"dkgSuccessesCount\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"Success\",\"type\":\"bytes\"}],\"name\":\"addDKGSuccess\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"withdrawable\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"NewSignedCRS\",\"type\":\"bytes\"}],\"name\":\"resetDKG\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"}],\"name\":\"dkgFinalizeds\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"blockGasLimit\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"dkgRound\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"OldOwner\",\"type\":\"address\"},{\"name\":\"NewOwner\",\"type\":\"address\"}],\"name\":\"transferNodeOwnershipByFoundation\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"totalStaked\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"}],\"name\":\"nodesOffsetByAddress\",\"outputs\":[{\"name\":\"\",\"type\":\"int256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"crs\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"roundLength\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"nextHalvingSupply\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"dkgComplaints\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"dkgMasterPublicKeyOffset\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"}],\"name\":\"dkgMPKReadys\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"lastHalvedAmount\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"finedRecords\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"lambdaDKG\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"Complaint\",\"type\":\"bytes\"}],\"name\":\"addDKGComplaint\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"fineValues\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"roundHeight\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"dkgMPKReadysCount\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"minBlockInterval\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"Round\",\"type\":\"uint256\"},{\"name\":\"SignedCRS\",\"type\":\"bytes\"}],\"name\":\"proposeCRS\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"Name\",\"type\":\"string\"},{\"name\":\"Email\",\"type\":\"string\"},{\"name\":\"Location\",\"type\":\"string\"},{\"name\":\"Url\",\"type\":\"string\"}],\"name\":\"updateNodeInfo\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"dkgMasterPublicKeys\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"PublicKey\",\"type\":\"bytes\"},{\"name\":\"Name\",\"type\":\"string\"},{\"name\":\"Email\",\"type\":\"string\"},{\"name\":\"Location\",\"type\":\"string\"},{\"name\":\"Url\",\"type\":\"string\"}],\"name\":\"register\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"}],\"name\":\"lastProposedHeight\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"minGasPrice\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"dkgFinalizedsCount\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"dkgComplaintsProposed\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"}],\"name\":\"nodesOffsetByNodeKeyAddress\",\"outputs\":[{\"name\":\"\",\"type\":\"int256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"lockupPeriod\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"dkgResetCount\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"NewOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"nodesLength\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"ConfigurationChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"Round\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"CRS\",\"type\":\"bytes32\"}],\"name\":\"CRSProposed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"NodeAddress\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"NewOwnerAddress\",\"type\":\"address\"}],\"name\":\"NodeOwnershipTransfered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"NodeAddress\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"PublicKey\",\"type\":\"bytes\"}],\"name\":\"NodePublicKeyReplaced\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"NodeAddress\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"Amount\",\"type\":\"uint256\"}],\"name\":\"Staked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"NodeAddress\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"Amount\",\"type\":\"uint256\"}],\"name\":\"Unstaked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"NodeAddress\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"Amount\",\"type\":\"uint256\"}],\"name\":\"Withdrawn\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"NodeAddress\",\"type\":\"address\"}],\"name\":\"NodeAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"NodeAddress\",\"type\":\"address\"}],\"name\":\"NodeRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"NodeAddress\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"Type\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"Arg1\",\"type\":\"bytes\"},{\"indexed\":false,\"name\":\"Arg2\",\"type\":\"bytes\"}],\"name\":\"Reported\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"NodeAddress\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"Amount\",\"type\":\"uint256\"}],\"name\":\"Fined\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"NodeAddress\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"Amount\",\"type\":\"uint256\"}],\"name\":\"FinePaid\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"Round\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"BlockHeight\",\"type\":\"uint256\"}],\"name\":\"DKGReset\",\"type\":\"event\"}]"

// GovernanceFuncSigs maps the 4-byte function signature to its string representation.
var GovernanceFuncSigs = map[string]string{

	"ab6a4013": "addDKGComplaint(bytes)",

	"29891e74": "addDKGFinalize(bytes)",

	"22f8a889": "addDKGMPKReady(bytes)",

	"19cf10bc": "addDKGMasterPublicKey(bytes)",

	"4ccaa59d": "addDKGSuccess(bytes)",

	"7877a797": "blockGasLimit()",

	"8a591369": "crs()",

	"47731325": "crsRound()",

	"8d90a159": "dkgComplaints(uint256)",

	"e746714d": "dkgComplaintsProposed(bytes32)",

	"70f3ac54": "dkgFinalizeds(address)",

	"dd2083d0": "dkgFinalizedsCount()",

	"9bc9f489": "dkgMPKReadys(address)",

	"b03e75e4": "dkgMPKReadysCount()",

	"99aadb9a": "dkgMasterPublicKeyOffset(bytes32)",

	"d21abb78": "dkgMasterPublicKeys(uint256)",

	"f0771659": "dkgResetCount(uint256)",

	"7a8c1c1e": "dkgRound()",

	"072bf11a": "dkgSuccesses(address)",

	"4cb38c7a": "dkgSuccessesCount()",

	"ae1f289d": "fineValues(uint256)",

	"a1e460eb": "finedRecords(bytes32)",

	"2deb3316": "lambdaBA()",

	"a9601a8d": "lambdaDKG()",

	"a0488e55": "lastHalvedAmount()",

	"d767fdc9": "lastProposedHeight(address)",

	"ee947a7c": "lockupPeriod()",

	"b3606b56": "minBlockInterval()",

	"d96ed505": "minGasPrice()",

	"375b3c0a": "minStake()",

	"2357e72a": "miningVelocity()",

	"8ca55162": "nextHalvingSupply()",

	"1c53c280": "nodes(uint256)",

	"f33de6c0": "nodesLength()",

	"85031123": "nodesOffsetByAddress(address)",

	"ed3fbdf9": "nodesOffsetByNodeKeyAddress(address)",

	"4c523214": "notaryParamAlpha()",

	"1f29f5c8": "notaryParamBeta()",

	"0b3441dc": "notarySetSize()",

	"8da5cb5b": "owner()",

	"3edfa229": "payFine(address)",

	"c448af34": "proposeCRS(uint256,bytes)",

	"d242f4cc": "register(bytes,string,string,string,string)",

	"3f007c20": "replaceNodePublicKey(bytes)",

	"320c0826": "report(uint256,bytes,bytes)",

	"57316157": "resetDKG(bytes)",

	"af58ef06": "roundHeight(uint256)",

	"8b649b94": "roundLength()",

	"3a4b66f1": "stake()",

	"817b1cd2": "totalStaked()",

	"18160ddd": "totalSupply()",

	"4264ecf9": "transferNodeOwnership(address)",

	"7ecfdeca": "transferNodeOwnershipByFoundation(address,address)",

	"f2fde38b": "transferOwnership(address)",

	"2e17de78": "unstake(uint256)",

	"11c7c321": "updateConfiguration(uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256[])",

	"c5ea6ea1": "updateNodeInfo(string,string,string,string)",

	"3ccfd60b": "withdraw()",

	"50188301": "withdrawable()",
}

// GovernanceBin is the compiled bytecode used for deploying new contracts.
const GovernanceBin = `0x608060405234801561001057600080fd5b50611abc806100206000396000f3fe60806040526004361061038c5760003560e01c8063817b1cd2116101dc578063b03e75e411610102578063d96ed505116100a0578063ee947a7c1161006f578063ee947a7c1461144c578063f077165914611461578063f2fde38b1461148b578063f33de6c0146114be5761038c565b8063d96ed505146113c5578063dd2083d0146113da578063e746714d146113ef578063ed3fbdf9146114195761038c565b8063c5ea6ea1116100dc578063c5ea6ea114610e70578063d21abb78146110b0578063d242f4cc146110da578063d767fdc9146113925761038c565b8063b03e75e414610d8e578063b3606b5614610da3578063c448af3414610db85761038c565b806399aadb9a1161017a578063a9601a8d11610149578063a9601a8d14610d25578063ab6a401314610505578063ae1f289d14610d3a578063af58ef0614610d645761038c565b806399aadb9a14610c895780639bc9f48914610cb3578063a0488e5514610ce6578063a1e460eb14610cfb5761038c565b80638b649b94116101b65780638b649b9414610b8f5780638ca5516214610ba45780638d90a15914610bb95780638da5cb5b14610c585761038c565b8063817b1cd214610b325780638503112314610b475780638a59136914610b7a5761038c565b80633a4b66f1116102c15780634cb38c7a1161025f57806370f3ac541161022e57806370f3ac5414610a9a5780637877a79714610acd5780637a8c1c1e14610ae25780637ecfdeca14610af75761038c565b80634cb38c7a14610a705780634ccaa59d146105055780635018830114610a8557806357316157146105055761038c565b80633f007c201161029b5780633f007c20146105055780634264ecf914610a235780634773132514610a465780634c52321414610a5b5761038c565b80633a4b66f1146109e05780633ccfd60b146109e85780633edfa229146109fd5761038c565b806322f8a8891161032e5780632deb3316116103085780632deb33161461084f5780632e17de7814610864578063320c08261461088e578063375b3c0a146109cb5761038c565b806322f8a889146105055780632357e72a1461083a57806329891e74146105055761038c565b806318160ddd1161036a57806318160ddd146104f057806319cf10bc146105055780631c53c280146105b65780631f29f5c8146108255761038c565b8063072bf11a146103915780630b3441dc146103d857806311c7c321146103ff575b600080fd5b34801561039d57600080fd5b506103c4600480360360208110156103b457600080fd5b50356001600160a01b03166114d3565b604080519115158252519081900360200190f35b3480156103e457600080fd5b506103ed6114e8565b60408051918252519081900360200190f35b34801561040b57600080fd5b506104ee600480360361016081101561042357600080fd5b81359160208101359160408201359160608101359160808201359160a08101359160c08201359160e08101359161010082013591610120810135918101906101608101610140820135600160201b81111561047d57600080fd5b82018360208201111561048f57600080fd5b803590602001918460208302840111600160201b831117156104b057600080fd5b9190808060200260200160405190810160405280939291908181526020018383602002808284376000920191909152509295506114ee945050505050565b005b3480156104fc57600080fd5b506103ed611512565b34801561051157600080fd5b506104ee6004803603602081101561052857600080fd5b810190602081018135600160201b81111561054257600080fd5b82018360208201111561055457600080fd5b803590602001918460018302840111600160201b8311171561057557600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250929550611518945050505050565b3480156105c257600080fd5b506105e0600480360360208110156105d957600080fd5b503561151b565b604051808b6001600160a01b03166001600160a01b03168152602001806020018a81526020018981526020018060200180602001806020018060200188815260200187815260200186810386528f818151815260200191508051906020019080838360005b8381101561065d578181015183820152602001610645565b50505050905090810190601f16801561068a5780820380516001836020036101000a031916815260200191505b5086810385528c5181528c516020918201918e019080838360005b838110156106bd5781810151838201526020016106a5565b50505050905090810190601f1680156106ea5780820380516001836020036101000a031916815260200191505b5086810384528b5181528b516020918201918d019080838360005b8381101561071d578181015183820152602001610705565b50505050905090810190601f16801561074a5780820380516001836020036101000a031916815260200191505b5086810383528a5181528a516020918201918c019080838360005b8381101561077d578181015183820152602001610765565b50505050905090810190601f1680156107aa5780820380516001836020036101000a031916815260200191505b5086810382528951815289516020918201918b019080838360005b838110156107dd5781810151838201526020016107c5565b50505050905090810190601f16801561080a5780820380516001836020036101000a031916815260200191505b509f5050505050505050505050505050505060405180910390f35b34801561083157600080fd5b506103ed611845565b34801561084657600080fd5b506103ed61184b565b34801561085b57600080fd5b506103ed611851565b34801561087057600080fd5b506104ee6004803603602081101561088757600080fd5b5035611518565b34801561089a57600080fd5b506104ee600480360360608110156108b157600080fd5b81359190810190604081016020820135600160201b8111156108d257600080fd5b8201836020820111156108e457600080fd5b803590602001918460018302840111600160201b8311171561090557600080fd5b91908080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152509295949360208101935035915050600160201b81111561095757600080fd5b82018360208201111561096957600080fd5b803590602001918460018302840111600160201b8311171561098a57600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250929550611857945050505050565b3480156109d757600080fd5b506103ed61185c565b6104ee611862565b3480156109f457600080fd5b506104ee611862565b6104ee60048036036020811015610a1357600080fd5b50356001600160a01b0316611518565b348015610a2f57600080fd5b506104ee60048036036020811015610a1357600080fd5b348015610a5257600080fd5b506103ed611864565b348015610a6757600080fd5b506103ed61186a565b348015610a7c57600080fd5b506103ed611870565b348015610a9157600080fd5b506103c4611876565b348015610aa657600080fd5b506103c460048036036020811015610abd57600080fd5b50356001600160a01b031661187b565b348015610ad957600080fd5b506103ed611890565b348015610aee57600080fd5b506103ed611896565b348015610b0357600080fd5b506104ee60048036036040811015610b1a57600080fd5b506001600160a01b038135811691602001351661189c565b348015610b3e57600080fd5b506103ed6118a0565b348015610b5357600080fd5b506103ed60048036036020811015610b6a57600080fd5b50356001600160a01b03166118a6565b348015610b8657600080fd5b506103ed6118b8565b348015610b9b57600080fd5b506103ed6118be565b348015610bb057600080fd5b506103ed6118c4565b348015610bc557600080fd5b50610be360048036036020811015610bdc57600080fd5b50356118ca565b6040805160208082528351818301528351919283929083019185019080838360005b83811015610c1d578181015183820152602001610c05565b50505050905090810190601f168015610c4a5780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b348015610c6457600080fd5b50610c6d611970565b604080516001600160a01b039092168252519081900360200190f35b348015610c9557600080fd5b506103ed60048036036020811015610cac57600080fd5b503561197f565b348015610cbf57600080fd5b506103c460048036036020811015610cd657600080fd5b50356001600160a01b0316611991565b348015610cf257600080fd5b506103ed6119a6565b348015610d0757600080fd5b506103c460048036036020811015610d1e57600080fd5b50356119ac565b348015610d3157600080fd5b506103ed6119c1565b348015610d4657600080fd5b506103ed60048036036020811015610d5d57600080fd5b50356119c7565b348015610d7057600080fd5b506103ed60048036036020811015610d8757600080fd5b50356119e5565b348015610d9a57600080fd5b506103ed6119f2565b348015610daf57600080fd5b506103ed6119f8565b348015610dc457600080fd5b506104ee60048036036040811015610ddb57600080fd5b81359190810190604081016020820135600160201b811115610dfc57600080fd5b820183602082011115610e0e57600080fd5b803590602001918460018302840111600160201b83111715610e2f57600080fd5b91908080601f01602080910402602001604051908101604052809392919081815260200183838082843760009201919091525092955061189c945050505050565b348015610e7c57600080fd5b506104ee60048036036080811015610e9357600080fd5b810190602081018135600160201b811115610ead57600080fd5b820183602082011115610ebf57600080fd5b803590602001918460018302840111600160201b83111715610ee057600080fd5b91908080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152509295949360208101935035915050600160201b811115610f3257600080fd5b820183602082011115610f4457600080fd5b803590602001918460018302840111600160201b83111715610f6557600080fd5b91908080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152509295949360208101935035915050600160201b811115610fb757600080fd5b820183602082011115610fc957600080fd5b803590602001918460018302840111600160201b83111715610fea57600080fd5b91908080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152509295949360208101935035915050600160201b81111561103c57600080fd5b82018360208201111561104e57600080fd5b803590602001918460018302840111600160201b8311171561106f57600080fd5b91908080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152509295506119fe945050505050565b3480156110bc57600080fd5b50610be3600480360360208110156110d357600080fd5b5035611a04565b6104ee600480360360a08110156110f057600080fd5b810190602081018135600160201b81111561110a57600080fd5b82018360208201111561111c57600080fd5b803590602001918460018302840111600160201b8311171561113d57600080fd5b91908080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152509295949360208101935035915050600160201b81111561118f57600080fd5b8201836020820111156111a157600080fd5b803590602001918460018302840111600160201b831117156111c257600080fd5b91908080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152509295949360208101935035915050600160201b81111561121457600080fd5b82018360208201111561122657600080fd5b803590602001918460018302840111600160201b8311171561124757600080fd5b91908080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152509295949360208101935035915050600160201b81111561129957600080fd5b8201836020820111156112ab57600080fd5b803590602001918460018302840111600160201b831117156112cc57600080fd5b91908080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152509295949360208101935035915050600160201b81111561131e57600080fd5b82018360208201111561133057600080fd5b803590602001918460018302840111600160201b8311171561135157600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250929550611a11945050505050565b34801561139e57600080fd5b506103ed600480360360208110156113b557600080fd5b50356001600160a01b0316611a18565b3480156113d157600080fd5b506103ed611a2a565b3480156113e657600080fd5b506103ed611a30565b3480156113fb57600080fd5b506103c46004803603602081101561141257600080fd5b5035611a36565b34801561142557600080fd5b506103ed6004803603602081101561143c57600080fd5b50356001600160a01b0316611a4b565b34801561145857600080fd5b506103ed611a5d565b34801561146d57600080fd5b506103ed6004803603602081101561148457600080fd5b5035611a63565b34801561149757600080fd5b506104ee600480360360208110156114ae57600080fd5b50356001600160a01b0316611a70565b3480156114ca57600080fd5b506103ed611876565b60136020526000908152604090205460ff1681565b601f5481565b6015546001600160a01b0316331461150557600080fd5b5050505050505050505050565b60015481565b50565b6003818154811061152857fe5b6000918252602091829020600a9091020180546001808301805460408051601f60026000199685161561010002969096019093169490940491820187900487028401870190528083526001600160a01b0390931695509293909291908301828280156115d55780601f106115aa576101008083540402835291602001916115d5565b820191906000526020600020905b8154815290600101906020018083116115b857829003601f168201915b505050505090806002015490806003015490806004018054600181600116156101000203166002900480601f01602080910402602001604051908101604052809291908181526020018280546001816001161561010002031660029004801561167f5780601f106116545761010080835404028352916020019161167f565b820191906000526020600020905b81548152906001019060200180831161166257829003601f168201915b5050505060058301805460408051602060026001851615610100026000190190941693909304601f810184900484028201840190925281815294959493509083018282801561170f5780601f106116e45761010080835404028352916020019161170f565b820191906000526020600020905b8154815290600101906020018083116116f257829003601f168201915b5050505060068301805460408051602060026001851615610100026000190190941693909304601f810184900484028201840190925281815294959493509083018282801561179f5780601f106117745761010080835404028352916020019161179f565b820191906000526020600020905b81548152906001019060200180831161178257829003601f168201915b5050505060078301805460408051602060026001851615610100026000190190941693909304601f810184900484028201840190925281815294959493509083018282801561182f5780601f106118045761010080835404028352916020019161182f565b820191906000526020600020905b81548152906001019060200180831161181257829003601f168201915b505050505090806008015490806009015490508a565b60215481565b60185481565b601d5481565b505050565b60165481565b565b60075481565b60205481565b60145481565b600090565b60116020526000908152604090205460ff1681565b601c5481565b60095481565b5050565b60025481565b60046020526000908152604090205481565b60085481565b60225481565b60195481565b600d81815481106118d757fe5b600091825260209182902001805460408051601f60026000196101006001871615020190941693909304928301859004850281018501909152818152935090918301828280156119685780601f1061193d57610100808354040283529160200191611968565b820191906000526020600020905b81548152906001019060200180831161194b57829003601f168201915b505050505081565b6015546001600160a01b031681565b600c6020526000908152604090205481565b600f6020526000908152604090205460ff1681565b601a5481565b60256020526000908152604090205460ff1681565b601e5481565b602481815481106119d457fe5b600091825260209091200154905081565b600081815481106119d457fe5b60105481565b60235481565b50505050565b600b81815481106118d757fe5b5050505050565b60066020526000908152604090205481565b601b5481565b60125481565b600e6020526000908152604090205460ff1681565b60056020526000908152604090205481565b60175481565b600a81815481106119d457fe5b6015546001600160a01b0316331461151857600080fdfea265627a7a723058209455fe1d223ee2f67e258f5526b5c272bb000f775541bcb469c3def1b40ad0cc64736f6c634300050a0032`

// DeployGovernance deploys a new Ethereum contract, binding an instance of Governance to it.
func DeployGovernance(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Governance, error) {
	parsed, err := abi.JSON(strings.NewReader(GovernanceABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(GovernanceBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Governance{GovernanceCaller: GovernanceCaller{contract: contract}, GovernanceTransactor: GovernanceTransactor{contract: contract}, GovernanceFilterer: GovernanceFilterer{contract: contract}}, nil
}

// Governance is an auto generated Go binding around an Ethereum contract.
type Governance struct {
	GovernanceCaller     // Read-only binding to the contract
	GovernanceTransactor // Write-only binding to the contract
	GovernanceFilterer   // Log filterer for contract events
}

// GovernanceCaller is an auto generated read-only Go binding around an Ethereum contract.
type GovernanceCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GovernanceTransactor is an auto generated write-only Go binding around an Ethereum contract.
type GovernanceTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GovernanceFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type GovernanceFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GovernanceSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type GovernanceSession struct {
	Contract     *Governance       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// GovernanceCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type GovernanceCallerSession struct {
	Contract *GovernanceCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// GovernanceTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type GovernanceTransactorSession struct {
	Contract     *GovernanceTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// GovernanceRaw is an auto generated low-level Go binding around an Ethereum contract.
type GovernanceRaw struct {
	Contract *Governance // Generic contract binding to access the raw methods on
}

// GovernanceCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type GovernanceCallerRaw struct {
	Contract *GovernanceCaller // Generic read-only contract binding to access the raw methods on
}

// GovernanceTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type GovernanceTransactorRaw struct {
	Contract *GovernanceTransactor // Generic write-only contract binding to access the raw methods on
}

// NewGovernance creates a new instance of Governance, bound to a specific deployed contract.
func NewGovernance(address common.Address, backend bind.ContractBackend) (*Governance, error) {
	contract, err := bindGovernance(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Governance{GovernanceCaller: GovernanceCaller{contract: contract}, GovernanceTransactor: GovernanceTransactor{contract: contract}, GovernanceFilterer: GovernanceFilterer{contract: contract}}, nil
}

// NewGovernanceCaller creates a new read-only instance of Governance, bound to a specific deployed contract.
func NewGovernanceCaller(address common.Address, caller bind.ContractCaller) (*GovernanceCaller, error) {
	contract, err := bindGovernance(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &GovernanceCaller{contract: contract}, nil
}

// NewGovernanceTransactor creates a new write-only instance of Governance, bound to a specific deployed contract.
func NewGovernanceTransactor(address common.Address, transactor bind.ContractTransactor) (*GovernanceTransactor, error) {
	contract, err := bindGovernance(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &GovernanceTransactor{contract: contract}, nil
}

// NewGovernanceFilterer creates a new log filterer instance of Governance, bound to a specific deployed contract.
func NewGovernanceFilterer(address common.Address, filterer bind.ContractFilterer) (*GovernanceFilterer, error) {
	contract, err := bindGovernance(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &GovernanceFilterer{contract: contract}, nil
}

// bindGovernance binds a generic wrapper to an already deployed contract.
func bindGovernance(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(GovernanceABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Governance *GovernanceRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Governance.Contract.GovernanceCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Governance *GovernanceRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Governance.Contract.GovernanceTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Governance *GovernanceRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Governance.Contract.GovernanceTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Governance *GovernanceCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Governance.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Governance *GovernanceTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Governance.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Governance *GovernanceTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Governance.Contract.contract.Transact(opts, method, params...)
}

// BlockGasLimit is a free data retrieval call binding the contract method 0x7877a797.
//
// Solidity: function blockGasLimit() constant returns(uint256)
func (_Governance *GovernanceCaller) BlockGasLimit(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Governance.contract.Call(opts, out, "blockGasLimit")
	return *ret0, err
}

// BlockGasLimit is a free data retrieval call binding the contract method 0x7877a797.
//
// Solidity: function blockGasLimit() constant returns(uint256)
func (_Governance *GovernanceSession) BlockGasLimit() (*big.Int, error) {
	return _Governance.Contract.BlockGasLimit(&_Governance.CallOpts)
}

// BlockGasLimit is a free data retrieval call binding the contract method 0x7877a797.
//
// Solidity: function blockGasLimit() constant returns(uint256)
func (_Governance *GovernanceCallerSession) BlockGasLimit() (*big.Int, error) {
	return _Governance.Contract.BlockGasLimit(&_Governance.CallOpts)
}

// Crs is a free data retrieval call binding the contract method 0x8a591369.
//
// Solidity: function crs() constant returns(bytes32)
func (_Governance *GovernanceCaller) Crs(opts *bind.CallOpts) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _Governance.contract.Call(opts, out, "crs")
	return *ret0, err
}

// Crs is a free data retrieval call binding the contract method 0x8a591369.
//
// Solidity: function crs() constant returns(bytes32)
func (_Governance *GovernanceSession) Crs() ([32]byte, error) {
	return _Governance.Contract.Crs(&_Governance.CallOpts)
}

// Crs is a free data retrieval call binding the contract method 0x8a591369.
//
// Solidity: function crs() constant returns(bytes32)
func (_Governance *GovernanceCallerSession) Crs() ([32]byte, error) {
	return _Governance.Contract.Crs(&_Governance.CallOpts)
}

// CrsRound is a free data retrieval call binding the contract method 0x47731325.
//
// Solidity: function crsRound() constant returns(uint256)
func (_Governance *GovernanceCaller) CrsRound(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Governance.contract.Call(opts, out, "crsRound")
	return *ret0, err
}

// CrsRound is a free data retrieval call binding the contract method 0x47731325.
//
// Solidity: function crsRound() constant returns(uint256)
func (_Governance *GovernanceSession) CrsRound() (*big.Int, error) {
	return _Governance.Contract.CrsRound(&_Governance.CallOpts)
}

// CrsRound is a free data retrieval call binding the contract method 0x47731325.
//
// Solidity: function crsRound() constant returns(uint256)
func (_Governance *GovernanceCallerSession) CrsRound() (*big.Int, error) {
	return _Governance.Contract.CrsRound(&_Governance.CallOpts)
}

// DkgComplaints is a free data retrieval call binding the contract method 0x8d90a159.
//
// Solidity: function dkgComplaints(uint256 ) constant returns(bytes)
func (_Governance *GovernanceCaller) DkgComplaints(opts *bind.CallOpts, arg0 *big.Int) ([]byte, error) {
	var (
		ret0 = new([]byte)
	)
	out := ret0
	err := _Governance.contract.Call(opts, out, "dkgComplaints", arg0)
	return *ret0, err
}

// DkgComplaints is a free data retrieval call binding the contract method 0x8d90a159.
//
// Solidity: function dkgComplaints(uint256 ) constant returns(bytes)
func (_Governance *GovernanceSession) DkgComplaints(arg0 *big.Int) ([]byte, error) {
	return _Governance.Contract.DkgComplaints(&_Governance.CallOpts, arg0)
}

// DkgComplaints is a free data retrieval call binding the contract method 0x8d90a159.
//
// Solidity: function dkgComplaints(uint256 ) constant returns(bytes)
func (_Governance *GovernanceCallerSession) DkgComplaints(arg0 *big.Int) ([]byte, error) {
	return _Governance.Contract.DkgComplaints(&_Governance.CallOpts, arg0)
}

// DkgComplaintsProposed is a free data retrieval call binding the contract method 0xe746714d.
//
// Solidity: function dkgComplaintsProposed(bytes32 ) constant returns(bool)
func (_Governance *GovernanceCaller) DkgComplaintsProposed(opts *bind.CallOpts, arg0 [32]byte) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Governance.contract.Call(opts, out, "dkgComplaintsProposed", arg0)
	return *ret0, err
}

// DkgComplaintsProposed is a free data retrieval call binding the contract method 0xe746714d.
//
// Solidity: function dkgComplaintsProposed(bytes32 ) constant returns(bool)
func (_Governance *GovernanceSession) DkgComplaintsProposed(arg0 [32]byte) (bool, error) {
	return _Governance.Contract.DkgComplaintsProposed(&_Governance.CallOpts, arg0)
}

// DkgComplaintsProposed is a free data retrieval call binding the contract method 0xe746714d.
//
// Solidity: function dkgComplaintsProposed(bytes32 ) constant returns(bool)
func (_Governance *GovernanceCallerSession) DkgComplaintsProposed(arg0 [32]byte) (bool, error) {
	return _Governance.Contract.DkgComplaintsProposed(&_Governance.CallOpts, arg0)
}

// DkgFinalizeds is a free data retrieval call binding the contract method 0x70f3ac54.
//
// Solidity: function dkgFinalizeds(address ) constant returns(bool)
func (_Governance *GovernanceCaller) DkgFinalizeds(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Governance.contract.Call(opts, out, "dkgFinalizeds", arg0)
	return *ret0, err
}

// DkgFinalizeds is a free data retrieval call binding the contract method 0x70f3ac54.
//
// Solidity: function dkgFinalizeds(address ) constant returns(bool)
func (_Governance *GovernanceSession) DkgFinalizeds(arg0 common.Address) (bool, error) {
	return _Governance.Contract.DkgFinalizeds(&_Governance.CallOpts, arg0)
}

// DkgFinalizeds is a free data retrieval call binding the contract method 0x70f3ac54.
//
// Solidity: function dkgFinalizeds(address ) constant returns(bool)
func (_Governance *GovernanceCallerSession) DkgFinalizeds(arg0 common.Address) (bool, error) {
	return _Governance.Contract.DkgFinalizeds(&_Governance.CallOpts, arg0)
}

// DkgFinalizedsCount is a free data retrieval call binding the contract method 0xdd2083d0.
//
// Solidity: function dkgFinalizedsCount() constant returns(uint256)
func (_Governance *GovernanceCaller) DkgFinalizedsCount(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Governance.contract.Call(opts, out, "dkgFinalizedsCount")
	return *ret0, err
}

// DkgFinalizedsCount is a free data retrieval call binding the contract method 0xdd2083d0.
//
// Solidity: function dkgFinalizedsCount() constant returns(uint256)
func (_Governance *GovernanceSession) DkgFinalizedsCount() (*big.Int, error) {
	return _Governance.Contract.DkgFinalizedsCount(&_Governance.CallOpts)
}

// DkgFinalizedsCount is a free data retrieval call binding the contract method 0xdd2083d0.
//
// Solidity: function dkgFinalizedsCount() constant returns(uint256)
func (_Governance *GovernanceCallerSession) DkgFinalizedsCount() (*big.Int, error) {
	return _Governance.Contract.DkgFinalizedsCount(&_Governance.CallOpts)
}

// DkgMPKReadys is a free data retrieval call binding the contract method 0x9bc9f489.
//
// Solidity: function dkgMPKReadys(address ) constant returns(bool)
func (_Governance *GovernanceCaller) DkgMPKReadys(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Governance.contract.Call(opts, out, "dkgMPKReadys", arg0)
	return *ret0, err
}

// DkgMPKReadys is a free data retrieval call binding the contract method 0x9bc9f489.
//
// Solidity: function dkgMPKReadys(address ) constant returns(bool)
func (_Governance *GovernanceSession) DkgMPKReadys(arg0 common.Address) (bool, error) {
	return _Governance.Contract.DkgMPKReadys(&_Governance.CallOpts, arg0)
}

// DkgMPKReadys is a free data retrieval call binding the contract method 0x9bc9f489.
//
// Solidity: function dkgMPKReadys(address ) constant returns(bool)
func (_Governance *GovernanceCallerSession) DkgMPKReadys(arg0 common.Address) (bool, error) {
	return _Governance.Contract.DkgMPKReadys(&_Governance.CallOpts, arg0)
}

// DkgMPKReadysCount is a free data retrieval call binding the contract method 0xb03e75e4.
//
// Solidity: function dkgMPKReadysCount() constant returns(uint256)
func (_Governance *GovernanceCaller) DkgMPKReadysCount(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Governance.contract.Call(opts, out, "dkgMPKReadysCount")
	return *ret0, err
}

// DkgMPKReadysCount is a free data retrieval call binding the contract method 0xb03e75e4.
//
// Solidity: function dkgMPKReadysCount() constant returns(uint256)
func (_Governance *GovernanceSession) DkgMPKReadysCount() (*big.Int, error) {
	return _Governance.Contract.DkgMPKReadysCount(&_Governance.CallOpts)
}

// DkgMPKReadysCount is a free data retrieval call binding the contract method 0xb03e75e4.
//
// Solidity: function dkgMPKReadysCount() constant returns(uint256)
func (_Governance *GovernanceCallerSession) DkgMPKReadysCount() (*big.Int, error) {
	return _Governance.Contract.DkgMPKReadysCount(&_Governance.CallOpts)
}

// DkgMasterPublicKeyOffset is a free data retrieval call binding the contract method 0x99aadb9a.
//
// Solidity: function dkgMasterPublicKeyOffset(bytes32 ) constant returns(uint256)
func (_Governance *GovernanceCaller) DkgMasterPublicKeyOffset(opts *bind.CallOpts, arg0 [32]byte) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Governance.contract.Call(opts, out, "dkgMasterPublicKeyOffset", arg0)
	return *ret0, err
}

// DkgMasterPublicKeyOffset is a free data retrieval call binding the contract method 0x99aadb9a.
//
// Solidity: function dkgMasterPublicKeyOffset(bytes32 ) constant returns(uint256)
func (_Governance *GovernanceSession) DkgMasterPublicKeyOffset(arg0 [32]byte) (*big.Int, error) {
	return _Governance.Contract.DkgMasterPublicKeyOffset(&_Governance.CallOpts, arg0)
}

// DkgMasterPublicKeyOffset is a free data retrieval call binding the contract method 0x99aadb9a.
//
// Solidity: function dkgMasterPublicKeyOffset(bytes32 ) constant returns(uint256)
func (_Governance *GovernanceCallerSession) DkgMasterPublicKeyOffset(arg0 [32]byte) (*big.Int, error) {
	return _Governance.Contract.DkgMasterPublicKeyOffset(&_Governance.CallOpts, arg0)
}

// DkgMasterPublicKeys is a free data retrieval call binding the contract method 0xd21abb78.
//
// Solidity: function dkgMasterPublicKeys(uint256 ) constant returns(bytes)
func (_Governance *GovernanceCaller) DkgMasterPublicKeys(opts *bind.CallOpts, arg0 *big.Int) ([]byte, error) {
	var (
		ret0 = new([]byte)
	)
	out := ret0
	err := _Governance.contract.Call(opts, out, "dkgMasterPublicKeys", arg0)
	return *ret0, err
}

// DkgMasterPublicKeys is a free data retrieval call binding the contract method 0xd21abb78.
//
// Solidity: function dkgMasterPublicKeys(uint256 ) constant returns(bytes)
func (_Governance *GovernanceSession) DkgMasterPublicKeys(arg0 *big.Int) ([]byte, error) {
	return _Governance.Contract.DkgMasterPublicKeys(&_Governance.CallOpts, arg0)
}

// DkgMasterPublicKeys is a free data retrieval call binding the contract method 0xd21abb78.
//
// Solidity: function dkgMasterPublicKeys(uint256 ) constant returns(bytes)
func (_Governance *GovernanceCallerSession) DkgMasterPublicKeys(arg0 *big.Int) ([]byte, error) {
	return _Governance.Contract.DkgMasterPublicKeys(&_Governance.CallOpts, arg0)
}

// DkgResetCount is a free data retrieval call binding the contract method 0xf0771659.
//
// Solidity: function dkgResetCount(uint256 ) constant returns(uint256)
func (_Governance *GovernanceCaller) DkgResetCount(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Governance.contract.Call(opts, out, "dkgResetCount", arg0)
	return *ret0, err
}

// DkgResetCount is a free data retrieval call binding the contract method 0xf0771659.
//
// Solidity: function dkgResetCount(uint256 ) constant returns(uint256)
func (_Governance *GovernanceSession) DkgResetCount(arg0 *big.Int) (*big.Int, error) {
	return _Governance.Contract.DkgResetCount(&_Governance.CallOpts, arg0)
}

// DkgResetCount is a free data retrieval call binding the contract method 0xf0771659.
//
// Solidity: function dkgResetCount(uint256 ) constant returns(uint256)
func (_Governance *GovernanceCallerSession) DkgResetCount(arg0 *big.Int) (*big.Int, error) {
	return _Governance.Contract.DkgResetCount(&_Governance.CallOpts, arg0)
}

// DkgRound is a free data retrieval call binding the contract method 0x7a8c1c1e.
//
// Solidity: function dkgRound() constant returns(uint256)
func (_Governance *GovernanceCaller) DkgRound(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Governance.contract.Call(opts, out, "dkgRound")
	return *ret0, err
}

// DkgRound is a free data retrieval call binding the contract method 0x7a8c1c1e.
//
// Solidity: function dkgRound() constant returns(uint256)
func (_Governance *GovernanceSession) DkgRound() (*big.Int, error) {
	return _Governance.Contract.DkgRound(&_Governance.CallOpts)
}

// DkgRound is a free data retrieval call binding the contract method 0x7a8c1c1e.
//
// Solidity: function dkgRound() constant returns(uint256)
func (_Governance *GovernanceCallerSession) DkgRound() (*big.Int, error) {
	return _Governance.Contract.DkgRound(&_Governance.CallOpts)
}

// DkgSuccesses is a free data retrieval call binding the contract method 0x072bf11a.
//
// Solidity: function dkgSuccesses(address ) constant returns(bool)
func (_Governance *GovernanceCaller) DkgSuccesses(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Governance.contract.Call(opts, out, "dkgSuccesses", arg0)
	return *ret0, err
}

// DkgSuccesses is a free data retrieval call binding the contract method 0x072bf11a.
//
// Solidity: function dkgSuccesses(address ) constant returns(bool)
func (_Governance *GovernanceSession) DkgSuccesses(arg0 common.Address) (bool, error) {
	return _Governance.Contract.DkgSuccesses(&_Governance.CallOpts, arg0)
}

// DkgSuccesses is a free data retrieval call binding the contract method 0x072bf11a.
//
// Solidity: function dkgSuccesses(address ) constant returns(bool)
func (_Governance *GovernanceCallerSession) DkgSuccesses(arg0 common.Address) (bool, error) {
	return _Governance.Contract.DkgSuccesses(&_Governance.CallOpts, arg0)
}

// DkgSuccessesCount is a free data retrieval call binding the contract method 0x4cb38c7a.
//
// Solidity: function dkgSuccessesCount() constant returns(uint256)
func (_Governance *GovernanceCaller) DkgSuccessesCount(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Governance.contract.Call(opts, out, "dkgSuccessesCount")
	return *ret0, err
}

// DkgSuccessesCount is a free data retrieval call binding the contract method 0x4cb38c7a.
//
// Solidity: function dkgSuccessesCount() constant returns(uint256)
func (_Governance *GovernanceSession) DkgSuccessesCount() (*big.Int, error) {
	return _Governance.Contract.DkgSuccessesCount(&_Governance.CallOpts)
}

// DkgSuccessesCount is a free data retrieval call binding the contract method 0x4cb38c7a.
//
// Solidity: function dkgSuccessesCount() constant returns(uint256)
func (_Governance *GovernanceCallerSession) DkgSuccessesCount() (*big.Int, error) {
	return _Governance.Contract.DkgSuccessesCount(&_Governance.CallOpts)
}

// FineValues is a free data retrieval call binding the contract method 0xae1f289d.
//
// Solidity: function fineValues(uint256 ) constant returns(uint256)
func (_Governance *GovernanceCaller) FineValues(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Governance.contract.Call(opts, out, "fineValues", arg0)
	return *ret0, err
}

// FineValues is a free data retrieval call binding the contract method 0xae1f289d.
//
// Solidity: function fineValues(uint256 ) constant returns(uint256)
func (_Governance *GovernanceSession) FineValues(arg0 *big.Int) (*big.Int, error) {
	return _Governance.Contract.FineValues(&_Governance.CallOpts, arg0)
}

// FineValues is a free data retrieval call binding the contract method 0xae1f289d.
//
// Solidity: function fineValues(uint256 ) constant returns(uint256)
func (_Governance *GovernanceCallerSession) FineValues(arg0 *big.Int) (*big.Int, error) {
	return _Governance.Contract.FineValues(&_Governance.CallOpts, arg0)
}

// FinedRecords is a free data retrieval call binding the contract method 0xa1e460eb.
//
// Solidity: function finedRecords(bytes32 ) constant returns(bool)
func (_Governance *GovernanceCaller) FinedRecords(opts *bind.CallOpts, arg0 [32]byte) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Governance.contract.Call(opts, out, "finedRecords", arg0)
	return *ret0, err
}

// FinedRecords is a free data retrieval call binding the contract method 0xa1e460eb.
//
// Solidity: function finedRecords(bytes32 ) constant returns(bool)
func (_Governance *GovernanceSession) FinedRecords(arg0 [32]byte) (bool, error) {
	return _Governance.Contract.FinedRecords(&_Governance.CallOpts, arg0)
}

// FinedRecords is a free data retrieval call binding the contract method 0xa1e460eb.
//
// Solidity: function finedRecords(bytes32 ) constant returns(bool)
func (_Governance *GovernanceCallerSession) FinedRecords(arg0 [32]byte) (bool, error) {
	return _Governance.Contract.FinedRecords(&_Governance.CallOpts, arg0)
}

// LambdaBA is a free data retrieval call binding the contract method 0x2deb3316.
//
// Solidity: function lambdaBA() constant returns(uint256)
func (_Governance *GovernanceCaller) LambdaBA(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Governance.contract.Call(opts, out, "lambdaBA")
	return *ret0, err
}

// LambdaBA is a free data retrieval call binding the contract method 0x2deb3316.
//
// Solidity: function lambdaBA() constant returns(uint256)
func (_Governance *GovernanceSession) LambdaBA() (*big.Int, error) {
	return _Governance.Contract.LambdaBA(&_Governance.CallOpts)
}

// LambdaBA is a free data retrieval call binding the contract method 0x2deb3316.
//
// Solidity: function lambdaBA() constant returns(uint256)
func (_Governance *GovernanceCallerSession) LambdaBA() (*big.Int, error) {
	return _Governance.Contract.LambdaBA(&_Governance.CallOpts)
}

// LambdaDKG is a free data retrieval call binding the contract method 0xa9601a8d.
//
// Solidity: function lambdaDKG() constant returns(uint256)
func (_Governance *GovernanceCaller) LambdaDKG(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Governance.contract.Call(opts, out, "lambdaDKG")
	return *ret0, err
}

// LambdaDKG is a free data retrieval call binding the contract method 0xa9601a8d.
//
// Solidity: function lambdaDKG() constant returns(uint256)
func (_Governance *GovernanceSession) LambdaDKG() (*big.Int, error) {
	return _Governance.Contract.LambdaDKG(&_Governance.CallOpts)
}

// LambdaDKG is a free data retrieval call binding the contract method 0xa9601a8d.
//
// Solidity: function lambdaDKG() constant returns(uint256)
func (_Governance *GovernanceCallerSession) LambdaDKG() (*big.Int, error) {
	return _Governance.Contract.LambdaDKG(&_Governance.CallOpts)
}

// LastHalvedAmount is a free data retrieval call binding the contract method 0xa0488e55.
//
// Solidity: function lastHalvedAmount() constant returns(uint256)
func (_Governance *GovernanceCaller) LastHalvedAmount(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Governance.contract.Call(opts, out, "lastHalvedAmount")
	return *ret0, err
}

// LastHalvedAmount is a free data retrieval call binding the contract method 0xa0488e55.
//
// Solidity: function lastHalvedAmount() constant returns(uint256)
func (_Governance *GovernanceSession) LastHalvedAmount() (*big.Int, error) {
	return _Governance.Contract.LastHalvedAmount(&_Governance.CallOpts)
}

// LastHalvedAmount is a free data retrieval call binding the contract method 0xa0488e55.
//
// Solidity: function lastHalvedAmount() constant returns(uint256)
func (_Governance *GovernanceCallerSession) LastHalvedAmount() (*big.Int, error) {
	return _Governance.Contract.LastHalvedAmount(&_Governance.CallOpts)
}

// LastProposedHeight is a free data retrieval call binding the contract method 0xd767fdc9.
//
// Solidity: function lastProposedHeight(address ) constant returns(uint256)
func (_Governance *GovernanceCaller) LastProposedHeight(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Governance.contract.Call(opts, out, "lastProposedHeight", arg0)
	return *ret0, err
}

// LastProposedHeight is a free data retrieval call binding the contract method 0xd767fdc9.
//
// Solidity: function lastProposedHeight(address ) constant returns(uint256)
func (_Governance *GovernanceSession) LastProposedHeight(arg0 common.Address) (*big.Int, error) {
	return _Governance.Contract.LastProposedHeight(&_Governance.CallOpts, arg0)
}

// LastProposedHeight is a free data retrieval call binding the contract method 0xd767fdc9.
//
// Solidity: function lastProposedHeight(address ) constant returns(uint256)
func (_Governance *GovernanceCallerSession) LastProposedHeight(arg0 common.Address) (*big.Int, error) {
	return _Governance.Contract.LastProposedHeight(&_Governance.CallOpts, arg0)
}

// LockupPeriod is a free data retrieval call binding the contract method 0xee947a7c.
//
// Solidity: function lockupPeriod() constant returns(uint256)
func (_Governance *GovernanceCaller) LockupPeriod(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Governance.contract.Call(opts, out, "lockupPeriod")
	return *ret0, err
}

// LockupPeriod is a free data retrieval call binding the contract method 0xee947a7c.
//
// Solidity: function lockupPeriod() constant returns(uint256)
func (_Governance *GovernanceSession) LockupPeriod() (*big.Int, error) {
	return _Governance.Contract.LockupPeriod(&_Governance.CallOpts)
}

// LockupPeriod is a free data retrieval call binding the contract method 0xee947a7c.
//
// Solidity: function lockupPeriod() constant returns(uint256)
func (_Governance *GovernanceCallerSession) LockupPeriod() (*big.Int, error) {
	return _Governance.Contract.LockupPeriod(&_Governance.CallOpts)
}

// MinBlockInterval is a free data retrieval call binding the contract method 0xb3606b56.
//
// Solidity: function minBlockInterval() constant returns(uint256)
func (_Governance *GovernanceCaller) MinBlockInterval(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Governance.contract.Call(opts, out, "minBlockInterval")
	return *ret0, err
}

// MinBlockInterval is a free data retrieval call binding the contract method 0xb3606b56.
//
// Solidity: function minBlockInterval() constant returns(uint256)
func (_Governance *GovernanceSession) MinBlockInterval() (*big.Int, error) {
	return _Governance.Contract.MinBlockInterval(&_Governance.CallOpts)
}

// MinBlockInterval is a free data retrieval call binding the contract method 0xb3606b56.
//
// Solidity: function minBlockInterval() constant returns(uint256)
func (_Governance *GovernanceCallerSession) MinBlockInterval() (*big.Int, error) {
	return _Governance.Contract.MinBlockInterval(&_Governance.CallOpts)
}

// MinGasPrice is a free data retrieval call binding the contract method 0xd96ed505.
//
// Solidity: function minGasPrice() constant returns(uint256)
func (_Governance *GovernanceCaller) MinGasPrice(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Governance.contract.Call(opts, out, "minGasPrice")
	return *ret0, err
}

// MinGasPrice is a free data retrieval call binding the contract method 0xd96ed505.
//
// Solidity: function minGasPrice() constant returns(uint256)
func (_Governance *GovernanceSession) MinGasPrice() (*big.Int, error) {
	return _Governance.Contract.MinGasPrice(&_Governance.CallOpts)
}

// MinGasPrice is a free data retrieval call binding the contract method 0xd96ed505.
//
// Solidity: function minGasPrice() constant returns(uint256)
func (_Governance *GovernanceCallerSession) MinGasPrice() (*big.Int, error) {
	return _Governance.Contract.MinGasPrice(&_Governance.CallOpts)
}

// MinStake is a free data retrieval call binding the contract method 0x375b3c0a.
//
// Solidity: function minStake() constant returns(uint256)
func (_Governance *GovernanceCaller) MinStake(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Governance.contract.Call(opts, out, "minStake")
	return *ret0, err
}

// MinStake is a free data retrieval call binding the contract method 0x375b3c0a.
//
// Solidity: function minStake() constant returns(uint256)
func (_Governance *GovernanceSession) MinStake() (*big.Int, error) {
	return _Governance.Contract.MinStake(&_Governance.CallOpts)
}

// MinStake is a free data retrieval call binding the contract method 0x375b3c0a.
//
// Solidity: function minStake() constant returns(uint256)
func (_Governance *GovernanceCallerSession) MinStake() (*big.Int, error) {
	return _Governance.Contract.MinStake(&_Governance.CallOpts)
}

// MiningVelocity is a free data retrieval call binding the contract method 0x2357e72a.
//
// Solidity: function miningVelocity() constant returns(uint256)
func (_Governance *GovernanceCaller) MiningVelocity(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Governance.contract.Call(opts, out, "miningVelocity")
	return *ret0, err
}

// MiningVelocity is a free data retrieval call binding the contract method 0x2357e72a.
//
// Solidity: function miningVelocity() constant returns(uint256)
func (_Governance *GovernanceSession) MiningVelocity() (*big.Int, error) {
	return _Governance.Contract.MiningVelocity(&_Governance.CallOpts)
}

// MiningVelocity is a free data retrieval call binding the contract method 0x2357e72a.
//
// Solidity: function miningVelocity() constant returns(uint256)
func (_Governance *GovernanceCallerSession) MiningVelocity() (*big.Int, error) {
	return _Governance.Contract.MiningVelocity(&_Governance.CallOpts)
}

// NextHalvingSupply is a free data retrieval call binding the contract method 0x8ca55162.
//
// Solidity: function nextHalvingSupply() constant returns(uint256)
func (_Governance *GovernanceCaller) NextHalvingSupply(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Governance.contract.Call(opts, out, "nextHalvingSupply")
	return *ret0, err
}

// NextHalvingSupply is a free data retrieval call binding the contract method 0x8ca55162.
//
// Solidity: function nextHalvingSupply() constant returns(uint256)
func (_Governance *GovernanceSession) NextHalvingSupply() (*big.Int, error) {
	return _Governance.Contract.NextHalvingSupply(&_Governance.CallOpts)
}

// NextHalvingSupply is a free data retrieval call binding the contract method 0x8ca55162.
//
// Solidity: function nextHalvingSupply() constant returns(uint256)
func (_Governance *GovernanceCallerSession) NextHalvingSupply() (*big.Int, error) {
	return _Governance.Contract.NextHalvingSupply(&_Governance.CallOpts)
}

// Nodes is a free data retrieval call binding the contract method 0x1c53c280.
//
// Solidity: function nodes(uint256 ) constant returns(address owner, bytes publicKey, uint256 staked, uint256 fined, string name, string email, string location, string url, uint256 unstaked, uint256 unstakedAt)
func (_Governance *GovernanceCaller) Nodes(opts *bind.CallOpts, arg0 *big.Int) (struct {
	Owner      common.Address
	PublicKey  []byte
	Staked     *big.Int
	Fined      *big.Int
	Name       string
	Email      string
	Location   string
	Url        string
	Unstaked   *big.Int
	UnstakedAt *big.Int
}, error) {
	ret := new(struct {
		Owner      common.Address
		PublicKey  []byte
		Staked     *big.Int
		Fined      *big.Int
		Name       string
		Email      string
		Location   string
		Url        string
		Unstaked   *big.Int
		UnstakedAt *big.Int
	})
	out := ret
	err := _Governance.contract.Call(opts, out, "nodes", arg0)
	return *ret, err
}

// Nodes is a free data retrieval call binding the contract method 0x1c53c280.
//
// Solidity: function nodes(uint256 ) constant returns(address owner, bytes publicKey, uint256 staked, uint256 fined, string name, string email, string location, string url, uint256 unstaked, uint256 unstakedAt)
func (_Governance *GovernanceSession) Nodes(arg0 *big.Int) (struct {
	Owner      common.Address
	PublicKey  []byte
	Staked     *big.Int
	Fined      *big.Int
	Name       string
	Email      string
	Location   string
	Url        string
	Unstaked   *big.Int
	UnstakedAt *big.Int
}, error) {
	return _Governance.Contract.Nodes(&_Governance.CallOpts, arg0)
}

// Nodes is a free data retrieval call binding the contract method 0x1c53c280.
//
// Solidity: function nodes(uint256 ) constant returns(address owner, bytes publicKey, uint256 staked, uint256 fined, string name, string email, string location, string url, uint256 unstaked, uint256 unstakedAt)
func (_Governance *GovernanceCallerSession) Nodes(arg0 *big.Int) (struct {
	Owner      common.Address
	PublicKey  []byte
	Staked     *big.Int
	Fined      *big.Int
	Name       string
	Email      string
	Location   string
	Url        string
	Unstaked   *big.Int
	UnstakedAt *big.Int
}, error) {
	return _Governance.Contract.Nodes(&_Governance.CallOpts, arg0)
}

// NodesLength is a free data retrieval call binding the contract method 0xf33de6c0.
//
// Solidity: function nodesLength() constant returns(uint256)
func (_Governance *GovernanceCaller) NodesLength(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Governance.contract.Call(opts, out, "nodesLength")
	return *ret0, err
}

// NodesLength is a free data retrieval call binding the contract method 0xf33de6c0.
//
// Solidity: function nodesLength() constant returns(uint256)
func (_Governance *GovernanceSession) NodesLength() (*big.Int, error) {
	return _Governance.Contract.NodesLength(&_Governance.CallOpts)
}

// NodesLength is a free data retrieval call binding the contract method 0xf33de6c0.
//
// Solidity: function nodesLength() constant returns(uint256)
func (_Governance *GovernanceCallerSession) NodesLength() (*big.Int, error) {
	return _Governance.Contract.NodesLength(&_Governance.CallOpts)
}

// NodesOffsetByAddress is a free data retrieval call binding the contract method 0x85031123.
//
// Solidity: function nodesOffsetByAddress(address ) constant returns(int256)
func (_Governance *GovernanceCaller) NodesOffsetByAddress(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Governance.contract.Call(opts, out, "nodesOffsetByAddress", arg0)
	return *ret0, err
}

// NodesOffsetByAddress is a free data retrieval call binding the contract method 0x85031123.
//
// Solidity: function nodesOffsetByAddress(address ) constant returns(int256)
func (_Governance *GovernanceSession) NodesOffsetByAddress(arg0 common.Address) (*big.Int, error) {
	return _Governance.Contract.NodesOffsetByAddress(&_Governance.CallOpts, arg0)
}

// NodesOffsetByAddress is a free data retrieval call binding the contract method 0x85031123.
//
// Solidity: function nodesOffsetByAddress(address ) constant returns(int256)
func (_Governance *GovernanceCallerSession) NodesOffsetByAddress(arg0 common.Address) (*big.Int, error) {
	return _Governance.Contract.NodesOffsetByAddress(&_Governance.CallOpts, arg0)
}

// NodesOffsetByNodeKeyAddress is a free data retrieval call binding the contract method 0xed3fbdf9.
//
// Solidity: function nodesOffsetByNodeKeyAddress(address ) constant returns(int256)
func (_Governance *GovernanceCaller) NodesOffsetByNodeKeyAddress(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Governance.contract.Call(opts, out, "nodesOffsetByNodeKeyAddress", arg0)
	return *ret0, err
}

// NodesOffsetByNodeKeyAddress is a free data retrieval call binding the contract method 0xed3fbdf9.
//
// Solidity: function nodesOffsetByNodeKeyAddress(address ) constant returns(int256)
func (_Governance *GovernanceSession) NodesOffsetByNodeKeyAddress(arg0 common.Address) (*big.Int, error) {
	return _Governance.Contract.NodesOffsetByNodeKeyAddress(&_Governance.CallOpts, arg0)
}

// NodesOffsetByNodeKeyAddress is a free data retrieval call binding the contract method 0xed3fbdf9.
//
// Solidity: function nodesOffsetByNodeKeyAddress(address ) constant returns(int256)
func (_Governance *GovernanceCallerSession) NodesOffsetByNodeKeyAddress(arg0 common.Address) (*big.Int, error) {
	return _Governance.Contract.NodesOffsetByNodeKeyAddress(&_Governance.CallOpts, arg0)
}

// NotaryParamAlpha is a free data retrieval call binding the contract method 0x4c523214.
//
// Solidity: function notaryParamAlpha() constant returns(uint256)
func (_Governance *GovernanceCaller) NotaryParamAlpha(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Governance.contract.Call(opts, out, "notaryParamAlpha")
	return *ret0, err
}

// NotaryParamAlpha is a free data retrieval call binding the contract method 0x4c523214.
//
// Solidity: function notaryParamAlpha() constant returns(uint256)
func (_Governance *GovernanceSession) NotaryParamAlpha() (*big.Int, error) {
	return _Governance.Contract.NotaryParamAlpha(&_Governance.CallOpts)
}

// NotaryParamAlpha is a free data retrieval call binding the contract method 0x4c523214.
//
// Solidity: function notaryParamAlpha() constant returns(uint256)
func (_Governance *GovernanceCallerSession) NotaryParamAlpha() (*big.Int, error) {
	return _Governance.Contract.NotaryParamAlpha(&_Governance.CallOpts)
}

// NotaryParamBeta is a free data retrieval call binding the contract method 0x1f29f5c8.
//
// Solidity: function notaryParamBeta() constant returns(uint256)
func (_Governance *GovernanceCaller) NotaryParamBeta(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Governance.contract.Call(opts, out, "notaryParamBeta")
	return *ret0, err
}

// NotaryParamBeta is a free data retrieval call binding the contract method 0x1f29f5c8.
//
// Solidity: function notaryParamBeta() constant returns(uint256)
func (_Governance *GovernanceSession) NotaryParamBeta() (*big.Int, error) {
	return _Governance.Contract.NotaryParamBeta(&_Governance.CallOpts)
}

// NotaryParamBeta is a free data retrieval call binding the contract method 0x1f29f5c8.
//
// Solidity: function notaryParamBeta() constant returns(uint256)
func (_Governance *GovernanceCallerSession) NotaryParamBeta() (*big.Int, error) {
	return _Governance.Contract.NotaryParamBeta(&_Governance.CallOpts)
}

// NotarySetSize is a free data retrieval call binding the contract method 0x0b3441dc.
//
// Solidity: function notarySetSize() constant returns(uint256)
func (_Governance *GovernanceCaller) NotarySetSize(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Governance.contract.Call(opts, out, "notarySetSize")
	return *ret0, err
}

// NotarySetSize is a free data retrieval call binding the contract method 0x0b3441dc.
//
// Solidity: function notarySetSize() constant returns(uint256)
func (_Governance *GovernanceSession) NotarySetSize() (*big.Int, error) {
	return _Governance.Contract.NotarySetSize(&_Governance.CallOpts)
}

// NotarySetSize is a free data retrieval call binding the contract method 0x0b3441dc.
//
// Solidity: function notarySetSize() constant returns(uint256)
func (_Governance *GovernanceCallerSession) NotarySetSize() (*big.Int, error) {
	return _Governance.Contract.NotarySetSize(&_Governance.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Governance *GovernanceCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Governance.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Governance *GovernanceSession) Owner() (common.Address, error) {
	return _Governance.Contract.Owner(&_Governance.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Governance *GovernanceCallerSession) Owner() (common.Address, error) {
	return _Governance.Contract.Owner(&_Governance.CallOpts)
}

// RoundHeight is a free data retrieval call binding the contract method 0xaf58ef06.
//
// Solidity: function roundHeight(uint256 ) constant returns(uint256)
func (_Governance *GovernanceCaller) RoundHeight(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Governance.contract.Call(opts, out, "roundHeight", arg0)
	return *ret0, err
}

// RoundHeight is a free data retrieval call binding the contract method 0xaf58ef06.
//
// Solidity: function roundHeight(uint256 ) constant returns(uint256)
func (_Governance *GovernanceSession) RoundHeight(arg0 *big.Int) (*big.Int, error) {
	return _Governance.Contract.RoundHeight(&_Governance.CallOpts, arg0)
}

// RoundHeight is a free data retrieval call binding the contract method 0xaf58ef06.
//
// Solidity: function roundHeight(uint256 ) constant returns(uint256)
func (_Governance *GovernanceCallerSession) RoundHeight(arg0 *big.Int) (*big.Int, error) {
	return _Governance.Contract.RoundHeight(&_Governance.CallOpts, arg0)
}

// RoundLength is a free data retrieval call binding the contract method 0x8b649b94.
//
// Solidity: function roundLength() constant returns(uint256)
func (_Governance *GovernanceCaller) RoundLength(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Governance.contract.Call(opts, out, "roundLength")
	return *ret0, err
}

// RoundLength is a free data retrieval call binding the contract method 0x8b649b94.
//
// Solidity: function roundLength() constant returns(uint256)
func (_Governance *GovernanceSession) RoundLength() (*big.Int, error) {
	return _Governance.Contract.RoundLength(&_Governance.CallOpts)
}

// RoundLength is a free data retrieval call binding the contract method 0x8b649b94.
//
// Solidity: function roundLength() constant returns(uint256)
func (_Governance *GovernanceCallerSession) RoundLength() (*big.Int, error) {
	return _Governance.Contract.RoundLength(&_Governance.CallOpts)
}

// TotalStaked is a free data retrieval call binding the contract method 0x817b1cd2.
//
// Solidity: function totalStaked() constant returns(uint256)
func (_Governance *GovernanceCaller) TotalStaked(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Governance.contract.Call(opts, out, "totalStaked")
	return *ret0, err
}

// TotalStaked is a free data retrieval call binding the contract method 0x817b1cd2.
//
// Solidity: function totalStaked() constant returns(uint256)
func (_Governance *GovernanceSession) TotalStaked() (*big.Int, error) {
	return _Governance.Contract.TotalStaked(&_Governance.CallOpts)
}

// TotalStaked is a free data retrieval call binding the contract method 0x817b1cd2.
//
// Solidity: function totalStaked() constant returns(uint256)
func (_Governance *GovernanceCallerSession) TotalStaked() (*big.Int, error) {
	return _Governance.Contract.TotalStaked(&_Governance.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_Governance *GovernanceCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Governance.contract.Call(opts, out, "totalSupply")
	return *ret0, err
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_Governance *GovernanceSession) TotalSupply() (*big.Int, error) {
	return _Governance.Contract.TotalSupply(&_Governance.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_Governance *GovernanceCallerSession) TotalSupply() (*big.Int, error) {
	return _Governance.Contract.TotalSupply(&_Governance.CallOpts)
}

// Withdrawable is a free data retrieval call binding the contract method 0x50188301.
//
// Solidity: function withdrawable() constant returns(bool)
func (_Governance *GovernanceCaller) Withdrawable(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Governance.contract.Call(opts, out, "withdrawable")
	return *ret0, err
}

// Withdrawable is a free data retrieval call binding the contract method 0x50188301.
//
// Solidity: function withdrawable() constant returns(bool)
func (_Governance *GovernanceSession) Withdrawable() (bool, error) {
	return _Governance.Contract.Withdrawable(&_Governance.CallOpts)
}

// Withdrawable is a free data retrieval call binding the contract method 0x50188301.
//
// Solidity: function withdrawable() constant returns(bool)
func (_Governance *GovernanceCallerSession) Withdrawable() (bool, error) {
	return _Governance.Contract.Withdrawable(&_Governance.CallOpts)
}

// AddDKGComplaint is a paid mutator transaction binding the contract method 0xab6a4013.
//
// Solidity: function addDKGComplaint(bytes Complaint) returns()
func (_Governance *GovernanceTransactor) AddDKGComplaint(opts *bind.TransactOpts, Complaint []byte) (*types.Transaction, error) {
	return _Governance.contract.Transact(opts, "addDKGComplaint", Complaint)
}

// AddDKGComplaint is a paid mutator transaction binding the contract method 0xab6a4013.
//
// Solidity: function addDKGComplaint(bytes Complaint) returns()
func (_Governance *GovernanceSession) AddDKGComplaint(Complaint []byte) (*types.Transaction, error) {
	return _Governance.Contract.AddDKGComplaint(&_Governance.TransactOpts, Complaint)
}

// AddDKGComplaint is a paid mutator transaction binding the contract method 0xab6a4013.
//
// Solidity: function addDKGComplaint(bytes Complaint) returns()
func (_Governance *GovernanceTransactorSession) AddDKGComplaint(Complaint []byte) (*types.Transaction, error) {
	return _Governance.Contract.AddDKGComplaint(&_Governance.TransactOpts, Complaint)
}

// AddDKGFinalize is a paid mutator transaction binding the contract method 0x29891e74.
//
// Solidity: function addDKGFinalize(bytes Finalize) returns()
func (_Governance *GovernanceTransactor) AddDKGFinalize(opts *bind.TransactOpts, Finalize []byte) (*types.Transaction, error) {
	return _Governance.contract.Transact(opts, "addDKGFinalize", Finalize)
}

// AddDKGFinalize is a paid mutator transaction binding the contract method 0x29891e74.
//
// Solidity: function addDKGFinalize(bytes Finalize) returns()
func (_Governance *GovernanceSession) AddDKGFinalize(Finalize []byte) (*types.Transaction, error) {
	return _Governance.Contract.AddDKGFinalize(&_Governance.TransactOpts, Finalize)
}

// AddDKGFinalize is a paid mutator transaction binding the contract method 0x29891e74.
//
// Solidity: function addDKGFinalize(bytes Finalize) returns()
func (_Governance *GovernanceTransactorSession) AddDKGFinalize(Finalize []byte) (*types.Transaction, error) {
	return _Governance.Contract.AddDKGFinalize(&_Governance.TransactOpts, Finalize)
}

// AddDKGMPKReady is a paid mutator transaction binding the contract method 0x22f8a889.
//
// Solidity: function addDKGMPKReady(bytes MPKReady) returns()
func (_Governance *GovernanceTransactor) AddDKGMPKReady(opts *bind.TransactOpts, MPKReady []byte) (*types.Transaction, error) {
	return _Governance.contract.Transact(opts, "addDKGMPKReady", MPKReady)
}

// AddDKGMPKReady is a paid mutator transaction binding the contract method 0x22f8a889.
//
// Solidity: function addDKGMPKReady(bytes MPKReady) returns()
func (_Governance *GovernanceSession) AddDKGMPKReady(MPKReady []byte) (*types.Transaction, error) {
	return _Governance.Contract.AddDKGMPKReady(&_Governance.TransactOpts, MPKReady)
}

// AddDKGMPKReady is a paid mutator transaction binding the contract method 0x22f8a889.
//
// Solidity: function addDKGMPKReady(bytes MPKReady) returns()
func (_Governance *GovernanceTransactorSession) AddDKGMPKReady(MPKReady []byte) (*types.Transaction, error) {
	return _Governance.Contract.AddDKGMPKReady(&_Governance.TransactOpts, MPKReady)
}

// AddDKGMasterPublicKey is a paid mutator transaction binding the contract method 0x19cf10bc.
//
// Solidity: function addDKGMasterPublicKey(bytes PublicKey) returns()
func (_Governance *GovernanceTransactor) AddDKGMasterPublicKey(opts *bind.TransactOpts, PublicKey []byte) (*types.Transaction, error) {
	return _Governance.contract.Transact(opts, "addDKGMasterPublicKey", PublicKey)
}

// AddDKGMasterPublicKey is a paid mutator transaction binding the contract method 0x19cf10bc.
//
// Solidity: function addDKGMasterPublicKey(bytes PublicKey) returns()
func (_Governance *GovernanceSession) AddDKGMasterPublicKey(PublicKey []byte) (*types.Transaction, error) {
	return _Governance.Contract.AddDKGMasterPublicKey(&_Governance.TransactOpts, PublicKey)
}

// AddDKGMasterPublicKey is a paid mutator transaction binding the contract method 0x19cf10bc.
//
// Solidity: function addDKGMasterPublicKey(bytes PublicKey) returns()
func (_Governance *GovernanceTransactorSession) AddDKGMasterPublicKey(PublicKey []byte) (*types.Transaction, error) {
	return _Governance.Contract.AddDKGMasterPublicKey(&_Governance.TransactOpts, PublicKey)
}

// AddDKGSuccess is a paid mutator transaction binding the contract method 0x4ccaa59d.
//
// Solidity: function addDKGSuccess(bytes Success) returns()
func (_Governance *GovernanceTransactor) AddDKGSuccess(opts *bind.TransactOpts, Success []byte) (*types.Transaction, error) {
	return _Governance.contract.Transact(opts, "addDKGSuccess", Success)
}

// AddDKGSuccess is a paid mutator transaction binding the contract method 0x4ccaa59d.
//
// Solidity: function addDKGSuccess(bytes Success) returns()
func (_Governance *GovernanceSession) AddDKGSuccess(Success []byte) (*types.Transaction, error) {
	return _Governance.Contract.AddDKGSuccess(&_Governance.TransactOpts, Success)
}

// AddDKGSuccess is a paid mutator transaction binding the contract method 0x4ccaa59d.
//
// Solidity: function addDKGSuccess(bytes Success) returns()
func (_Governance *GovernanceTransactorSession) AddDKGSuccess(Success []byte) (*types.Transaction, error) {
	return _Governance.Contract.AddDKGSuccess(&_Governance.TransactOpts, Success)
}

// PayFine is a paid mutator transaction binding the contract method 0x3edfa229.
//
// Solidity: function payFine(address NodeAddress) returns()
func (_Governance *GovernanceTransactor) PayFine(opts *bind.TransactOpts, NodeAddress common.Address) (*types.Transaction, error) {
	return _Governance.contract.Transact(opts, "payFine", NodeAddress)
}

// PayFine is a paid mutator transaction binding the contract method 0x3edfa229.
//
// Solidity: function payFine(address NodeAddress) returns()
func (_Governance *GovernanceSession) PayFine(NodeAddress common.Address) (*types.Transaction, error) {
	return _Governance.Contract.PayFine(&_Governance.TransactOpts, NodeAddress)
}

// PayFine is a paid mutator transaction binding the contract method 0x3edfa229.
//
// Solidity: function payFine(address NodeAddress) returns()
func (_Governance *GovernanceTransactorSession) PayFine(NodeAddress common.Address) (*types.Transaction, error) {
	return _Governance.Contract.PayFine(&_Governance.TransactOpts, NodeAddress)
}

// ProposeCRS is a paid mutator transaction binding the contract method 0xc448af34.
//
// Solidity: function proposeCRS(uint256 Round, bytes SignedCRS) returns()
func (_Governance *GovernanceTransactor) ProposeCRS(opts *bind.TransactOpts, Round *big.Int, SignedCRS []byte) (*types.Transaction, error) {
	return _Governance.contract.Transact(opts, "proposeCRS", Round, SignedCRS)
}

// ProposeCRS is a paid mutator transaction binding the contract method 0xc448af34.
//
// Solidity: function proposeCRS(uint256 Round, bytes SignedCRS) returns()
func (_Governance *GovernanceSession) ProposeCRS(Round *big.Int, SignedCRS []byte) (*types.Transaction, error) {
	return _Governance.Contract.ProposeCRS(&_Governance.TransactOpts, Round, SignedCRS)
}

// ProposeCRS is a paid mutator transaction binding the contract method 0xc448af34.
//
// Solidity: function proposeCRS(uint256 Round, bytes SignedCRS) returns()
func (_Governance *GovernanceTransactorSession) ProposeCRS(Round *big.Int, SignedCRS []byte) (*types.Transaction, error) {
	return _Governance.Contract.ProposeCRS(&_Governance.TransactOpts, Round, SignedCRS)
}

// Register is a paid mutator transaction binding the contract method 0xd242f4cc.
//
// Solidity: function register(bytes PublicKey, string Name, string Email, string Location, string Url) returns()
func (_Governance *GovernanceTransactor) Register(opts *bind.TransactOpts, PublicKey []byte, Name string, Email string, Location string, Url string) (*types.Transaction, error) {
	return _Governance.contract.Transact(opts, "register", PublicKey, Name, Email, Location, Url)
}

// Register is a paid mutator transaction binding the contract method 0xd242f4cc.
//
// Solidity: function register(bytes PublicKey, string Name, string Email, string Location, string Url) returns()
func (_Governance *GovernanceSession) Register(PublicKey []byte, Name string, Email string, Location string, Url string) (*types.Transaction, error) {
	return _Governance.Contract.Register(&_Governance.TransactOpts, PublicKey, Name, Email, Location, Url)
}

// Register is a paid mutator transaction binding the contract method 0xd242f4cc.
//
// Solidity: function register(bytes PublicKey, string Name, string Email, string Location, string Url) returns()
func (_Governance *GovernanceTransactorSession) Register(PublicKey []byte, Name string, Email string, Location string, Url string) (*types.Transaction, error) {
	return _Governance.Contract.Register(&_Governance.TransactOpts, PublicKey, Name, Email, Location, Url)
}

// ReplaceNodePublicKey is a paid mutator transaction binding the contract method 0x3f007c20.
//
// Solidity: function replaceNodePublicKey(bytes NewPublicKey) returns()
func (_Governance *GovernanceTransactor) ReplaceNodePublicKey(opts *bind.TransactOpts, NewPublicKey []byte) (*types.Transaction, error) {
	return _Governance.contract.Transact(opts, "replaceNodePublicKey", NewPublicKey)
}

// ReplaceNodePublicKey is a paid mutator transaction binding the contract method 0x3f007c20.
//
// Solidity: function replaceNodePublicKey(bytes NewPublicKey) returns()
func (_Governance *GovernanceSession) ReplaceNodePublicKey(NewPublicKey []byte) (*types.Transaction, error) {
	return _Governance.Contract.ReplaceNodePublicKey(&_Governance.TransactOpts, NewPublicKey)
}

// ReplaceNodePublicKey is a paid mutator transaction binding the contract method 0x3f007c20.
//
// Solidity: function replaceNodePublicKey(bytes NewPublicKey) returns()
func (_Governance *GovernanceTransactorSession) ReplaceNodePublicKey(NewPublicKey []byte) (*types.Transaction, error) {
	return _Governance.Contract.ReplaceNodePublicKey(&_Governance.TransactOpts, NewPublicKey)
}

// Report is a paid mutator transaction binding the contract method 0x320c0826.
//
// Solidity: function report(uint256 Type, bytes Arg1, bytes Arg2) returns()
func (_Governance *GovernanceTransactor) Report(opts *bind.TransactOpts, Type *big.Int, Arg1 []byte, Arg2 []byte) (*types.Transaction, error) {
	return _Governance.contract.Transact(opts, "report", Type, Arg1, Arg2)
}

// Report is a paid mutator transaction binding the contract method 0x320c0826.
//
// Solidity: function report(uint256 Type, bytes Arg1, bytes Arg2) returns()
func (_Governance *GovernanceSession) Report(Type *big.Int, Arg1 []byte, Arg2 []byte) (*types.Transaction, error) {
	return _Governance.Contract.Report(&_Governance.TransactOpts, Type, Arg1, Arg2)
}

// Report is a paid mutator transaction binding the contract method 0x320c0826.
//
// Solidity: function report(uint256 Type, bytes Arg1, bytes Arg2) returns()
func (_Governance *GovernanceTransactorSession) Report(Type *big.Int, Arg1 []byte, Arg2 []byte) (*types.Transaction, error) {
	return _Governance.Contract.Report(&_Governance.TransactOpts, Type, Arg1, Arg2)
}

// ResetDKG is a paid mutator transaction binding the contract method 0x57316157.
//
// Solidity: function resetDKG(bytes NewSignedCRS) returns()
func (_Governance *GovernanceTransactor) ResetDKG(opts *bind.TransactOpts, NewSignedCRS []byte) (*types.Transaction, error) {
	return _Governance.contract.Transact(opts, "resetDKG", NewSignedCRS)
}

// ResetDKG is a paid mutator transaction binding the contract method 0x57316157.
//
// Solidity: function resetDKG(bytes NewSignedCRS) returns()
func (_Governance *GovernanceSession) ResetDKG(NewSignedCRS []byte) (*types.Transaction, error) {
	return _Governance.Contract.ResetDKG(&_Governance.TransactOpts, NewSignedCRS)
}

// ResetDKG is a paid mutator transaction binding the contract method 0x57316157.
//
// Solidity: function resetDKG(bytes NewSignedCRS) returns()
func (_Governance *GovernanceTransactorSession) ResetDKG(NewSignedCRS []byte) (*types.Transaction, error) {
	return _Governance.Contract.ResetDKG(&_Governance.TransactOpts, NewSignedCRS)
}

// Stake is a paid mutator transaction binding the contract method 0x3a4b66f1.
//
// Solidity: function stake() returns()
func (_Governance *GovernanceTransactor) Stake(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Governance.contract.Transact(opts, "stake")
}

// Stake is a paid mutator transaction binding the contract method 0x3a4b66f1.
//
// Solidity: function stake() returns()
func (_Governance *GovernanceSession) Stake() (*types.Transaction, error) {
	return _Governance.Contract.Stake(&_Governance.TransactOpts)
}

// Stake is a paid mutator transaction binding the contract method 0x3a4b66f1.
//
// Solidity: function stake() returns()
func (_Governance *GovernanceTransactorSession) Stake() (*types.Transaction, error) {
	return _Governance.Contract.Stake(&_Governance.TransactOpts)
}

// TransferNodeOwnership is a paid mutator transaction binding the contract method 0x4264ecf9.
//
// Solidity: function transferNodeOwnership(address NewOwner) returns()
func (_Governance *GovernanceTransactor) TransferNodeOwnership(opts *bind.TransactOpts, NewOwner common.Address) (*types.Transaction, error) {
	return _Governance.contract.Transact(opts, "transferNodeOwnership", NewOwner)
}

// TransferNodeOwnership is a paid mutator transaction binding the contract method 0x4264ecf9.
//
// Solidity: function transferNodeOwnership(address NewOwner) returns()
func (_Governance *GovernanceSession) TransferNodeOwnership(NewOwner common.Address) (*types.Transaction, error) {
	return _Governance.Contract.TransferNodeOwnership(&_Governance.TransactOpts, NewOwner)
}

// TransferNodeOwnership is a paid mutator transaction binding the contract method 0x4264ecf9.
//
// Solidity: function transferNodeOwnership(address NewOwner) returns()
func (_Governance *GovernanceTransactorSession) TransferNodeOwnership(NewOwner common.Address) (*types.Transaction, error) {
	return _Governance.Contract.TransferNodeOwnership(&_Governance.TransactOpts, NewOwner)
}

// TransferNodeOwnershipByFoundation is a paid mutator transaction binding the contract method 0x7ecfdeca.
//
// Solidity: function transferNodeOwnershipByFoundation(address OldOwner, address NewOwner) returns()
func (_Governance *GovernanceTransactor) TransferNodeOwnershipByFoundation(opts *bind.TransactOpts, OldOwner common.Address, NewOwner common.Address) (*types.Transaction, error) {
	return _Governance.contract.Transact(opts, "transferNodeOwnershipByFoundation", OldOwner, NewOwner)
}

// TransferNodeOwnershipByFoundation is a paid mutator transaction binding the contract method 0x7ecfdeca.
//
// Solidity: function transferNodeOwnershipByFoundation(address OldOwner, address NewOwner) returns()
func (_Governance *GovernanceSession) TransferNodeOwnershipByFoundation(OldOwner common.Address, NewOwner common.Address) (*types.Transaction, error) {
	return _Governance.Contract.TransferNodeOwnershipByFoundation(&_Governance.TransactOpts, OldOwner, NewOwner)
}

// TransferNodeOwnershipByFoundation is a paid mutator transaction binding the contract method 0x7ecfdeca.
//
// Solidity: function transferNodeOwnershipByFoundation(address OldOwner, address NewOwner) returns()
func (_Governance *GovernanceTransactorSession) TransferNodeOwnershipByFoundation(OldOwner common.Address, NewOwner common.Address) (*types.Transaction, error) {
	return _Governance.Contract.TransferNodeOwnershipByFoundation(&_Governance.TransactOpts, OldOwner, NewOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address NewOwner) returns()
func (_Governance *GovernanceTransactor) TransferOwnership(opts *bind.TransactOpts, NewOwner common.Address) (*types.Transaction, error) {
	return _Governance.contract.Transact(opts, "transferOwnership", NewOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address NewOwner) returns()
func (_Governance *GovernanceSession) TransferOwnership(NewOwner common.Address) (*types.Transaction, error) {
	return _Governance.Contract.TransferOwnership(&_Governance.TransactOpts, NewOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address NewOwner) returns()
func (_Governance *GovernanceTransactorSession) TransferOwnership(NewOwner common.Address) (*types.Transaction, error) {
	return _Governance.Contract.TransferOwnership(&_Governance.TransactOpts, NewOwner)
}

// Unstake is a paid mutator transaction binding the contract method 0x2e17de78.
//
// Solidity: function unstake(uint256 Amount) returns()
func (_Governance *GovernanceTransactor) Unstake(opts *bind.TransactOpts, Amount *big.Int) (*types.Transaction, error) {
	return _Governance.contract.Transact(opts, "unstake", Amount)
}

// Unstake is a paid mutator transaction binding the contract method 0x2e17de78.
//
// Solidity: function unstake(uint256 Amount) returns()
func (_Governance *GovernanceSession) Unstake(Amount *big.Int) (*types.Transaction, error) {
	return _Governance.Contract.Unstake(&_Governance.TransactOpts, Amount)
}

// Unstake is a paid mutator transaction binding the contract method 0x2e17de78.
//
// Solidity: function unstake(uint256 Amount) returns()
func (_Governance *GovernanceTransactorSession) Unstake(Amount *big.Int) (*types.Transaction, error) {
	return _Governance.Contract.Unstake(&_Governance.TransactOpts, Amount)
}

// UpdateConfiguration is a paid mutator transaction binding the contract method 0x11c7c321.
//
// Solidity: function updateConfiguration(uint256 MinStake, uint256 LockupPeriod, uint256 MinGasPrice, uint256 BlockGasLimit, uint256 LambdaBA, uint256 LambdaDKG, uint256 NotaryParamAlpha, uint256 NotaryParamBeta, uint256 RoundLength, uint256 MinBlockInterval, uint256[] FineValues) returns()
func (_Governance *GovernanceTransactor) UpdateConfiguration(opts *bind.TransactOpts, MinStake *big.Int, LockupPeriod *big.Int, MinGasPrice *big.Int, BlockGasLimit *big.Int, LambdaBA *big.Int, LambdaDKG *big.Int, NotaryParamAlpha *big.Int, NotaryParamBeta *big.Int, RoundLength *big.Int, MinBlockInterval *big.Int, FineValues []*big.Int) (*types.Transaction, error) {
	return _Governance.contract.Transact(opts, "updateConfiguration", MinStake, LockupPeriod, MinGasPrice, BlockGasLimit, LambdaBA, LambdaDKG, NotaryParamAlpha, NotaryParamBeta, RoundLength, MinBlockInterval, FineValues)
}

// UpdateConfiguration is a paid mutator transaction binding the contract method 0x11c7c321.
//
// Solidity: function updateConfiguration(uint256 MinStake, uint256 LockupPeriod, uint256 MinGasPrice, uint256 BlockGasLimit, uint256 LambdaBA, uint256 LambdaDKG, uint256 NotaryParamAlpha, uint256 NotaryParamBeta, uint256 RoundLength, uint256 MinBlockInterval, uint256[] FineValues) returns()
func (_Governance *GovernanceSession) UpdateConfiguration(MinStake *big.Int, LockupPeriod *big.Int, MinGasPrice *big.Int, BlockGasLimit *big.Int, LambdaBA *big.Int, LambdaDKG *big.Int, NotaryParamAlpha *big.Int, NotaryParamBeta *big.Int, RoundLength *big.Int, MinBlockInterval *big.Int, FineValues []*big.Int) (*types.Transaction, error) {
	return _Governance.Contract.UpdateConfiguration(&_Governance.TransactOpts, MinStake, LockupPeriod, MinGasPrice, BlockGasLimit, LambdaBA, LambdaDKG, NotaryParamAlpha, NotaryParamBeta, RoundLength, MinBlockInterval, FineValues)
}

// UpdateConfiguration is a paid mutator transaction binding the contract method 0x11c7c321.
//
// Solidity: function updateConfiguration(uint256 MinStake, uint256 LockupPeriod, uint256 MinGasPrice, uint256 BlockGasLimit, uint256 LambdaBA, uint256 LambdaDKG, uint256 NotaryParamAlpha, uint256 NotaryParamBeta, uint256 RoundLength, uint256 MinBlockInterval, uint256[] FineValues) returns()
func (_Governance *GovernanceTransactorSession) UpdateConfiguration(MinStake *big.Int, LockupPeriod *big.Int, MinGasPrice *big.Int, BlockGasLimit *big.Int, LambdaBA *big.Int, LambdaDKG *big.Int, NotaryParamAlpha *big.Int, NotaryParamBeta *big.Int, RoundLength *big.Int, MinBlockInterval *big.Int, FineValues []*big.Int) (*types.Transaction, error) {
	return _Governance.Contract.UpdateConfiguration(&_Governance.TransactOpts, MinStake, LockupPeriod, MinGasPrice, BlockGasLimit, LambdaBA, LambdaDKG, NotaryParamAlpha, NotaryParamBeta, RoundLength, MinBlockInterval, FineValues)
}

// UpdateNodeInfo is a paid mutator transaction binding the contract method 0xc5ea6ea1.
//
// Solidity: function updateNodeInfo(string Name, string Email, string Location, string Url) returns()
func (_Governance *GovernanceTransactor) UpdateNodeInfo(opts *bind.TransactOpts, Name string, Email string, Location string, Url string) (*types.Transaction, error) {
	return _Governance.contract.Transact(opts, "updateNodeInfo", Name, Email, Location, Url)
}

// UpdateNodeInfo is a paid mutator transaction binding the contract method 0xc5ea6ea1.
//
// Solidity: function updateNodeInfo(string Name, string Email, string Location, string Url) returns()
func (_Governance *GovernanceSession) UpdateNodeInfo(Name string, Email string, Location string, Url string) (*types.Transaction, error) {
	return _Governance.Contract.UpdateNodeInfo(&_Governance.TransactOpts, Name, Email, Location, Url)
}

// UpdateNodeInfo is a paid mutator transaction binding the contract method 0xc5ea6ea1.
//
// Solidity: function updateNodeInfo(string Name, string Email, string Location, string Url) returns()
func (_Governance *GovernanceTransactorSession) UpdateNodeInfo(Name string, Email string, Location string, Url string) (*types.Transaction, error) {
	return _Governance.Contract.UpdateNodeInfo(&_Governance.TransactOpts, Name, Email, Location, Url)
}

// Withdraw is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() returns()
func (_Governance *GovernanceTransactor) Withdraw(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Governance.contract.Transact(opts, "withdraw")
}

// Withdraw is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() returns()
func (_Governance *GovernanceSession) Withdraw() (*types.Transaction, error) {
	return _Governance.Contract.Withdraw(&_Governance.TransactOpts)
}

// Withdraw is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() returns()
func (_Governance *GovernanceTransactorSession) Withdraw() (*types.Transaction, error) {
	return _Governance.Contract.Withdraw(&_Governance.TransactOpts)
}

// GovernanceCRSProposedIterator is returned from FilterCRSProposed and is used to iterate over the raw logs and unpacked data for CRSProposed events raised by the Governance contract.
type GovernanceCRSProposedIterator struct {
	Event *GovernanceCRSProposed // Event containing the contract specifics and raw log

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
func (it *GovernanceCRSProposedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GovernanceCRSProposed)
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
		it.Event = new(GovernanceCRSProposed)
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
func (it *GovernanceCRSProposedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GovernanceCRSProposedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GovernanceCRSProposed represents a CRSProposed event raised by the Governance contract.
type GovernanceCRSProposed struct {
	Round *big.Int
	CRS   [32]byte
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterCRSProposed is a free log retrieval operation binding the contract event 0xd9b8f7e45cd2897523743eb19fd2021ba458d2962753a38bc31223e9b05ad408.
//
// Solidity: event CRSProposed(uint256 indexed Round, bytes32 CRS)
func (_Governance *GovernanceFilterer) FilterCRSProposed(opts *bind.FilterOpts, Round []*big.Int) (*GovernanceCRSProposedIterator, error) {

	var RoundRule []interface{}
	for _, RoundItem := range Round {
		RoundRule = append(RoundRule, RoundItem)
	}

	logs, sub, err := _Governance.contract.FilterLogs(opts, "CRSProposed", RoundRule)
	if err != nil {
		return nil, err
	}
	return &GovernanceCRSProposedIterator{contract: _Governance.contract, event: "CRSProposed", logs: logs, sub: sub}, nil
}

// WatchCRSProposed is a free log subscription operation binding the contract event 0xd9b8f7e45cd2897523743eb19fd2021ba458d2962753a38bc31223e9b05ad408.
//
// Solidity: event CRSProposed(uint256 indexed Round, bytes32 CRS)
func (_Governance *GovernanceFilterer) WatchCRSProposed(opts *bind.WatchOpts, sink chan<- *GovernanceCRSProposed, Round []*big.Int) (event.Subscription, error) {

	var RoundRule []interface{}
	for _, RoundItem := range Round {
		RoundRule = append(RoundRule, RoundItem)
	}

	logs, sub, err := _Governance.contract.WatchLogs(opts, "CRSProposed", RoundRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GovernanceCRSProposed)
				if err := _Governance.contract.UnpackLog(event, "CRSProposed", log); err != nil {
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

// ParseCRSProposed is a log parse operation binding the contract event 0xd9b8f7e45cd2897523743eb19fd2021ba458d2962753a38bc31223e9b05ad408.
//
// Solidity: event CRSProposed(uint256 indexed Round, bytes32 CRS)
func (_Governance *GovernanceFilterer) ParseCRSProposed(log types.Log) (*GovernanceCRSProposed, error) {
	event := new(GovernanceCRSProposed)
	if err := _Governance.contract.UnpackLog(event, "CRSProposed", log); err != nil {
		return nil, err
	}
	return event, nil
}

// GovernanceConfigurationChangedIterator is returned from FilterConfigurationChanged and is used to iterate over the raw logs and unpacked data for ConfigurationChanged events raised by the Governance contract.
type GovernanceConfigurationChangedIterator struct {
	Event *GovernanceConfigurationChanged // Event containing the contract specifics and raw log

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
func (it *GovernanceConfigurationChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GovernanceConfigurationChanged)
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
		it.Event = new(GovernanceConfigurationChanged)
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
func (it *GovernanceConfigurationChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GovernanceConfigurationChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GovernanceConfigurationChanged represents a ConfigurationChanged event raised by the Governance contract.
type GovernanceConfigurationChanged struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterConfigurationChanged is a free log retrieval operation binding the contract event 0xb6aa5c479f96dcbdbe1d8b70883ec95e8799f8e602ee347ed972a22c974b14c0.
//
// Solidity: event ConfigurationChanged()
func (_Governance *GovernanceFilterer) FilterConfigurationChanged(opts *bind.FilterOpts) (*GovernanceConfigurationChangedIterator, error) {

	logs, sub, err := _Governance.contract.FilterLogs(opts, "ConfigurationChanged")
	if err != nil {
		return nil, err
	}
	return &GovernanceConfigurationChangedIterator{contract: _Governance.contract, event: "ConfigurationChanged", logs: logs, sub: sub}, nil
}

// WatchConfigurationChanged is a free log subscription operation binding the contract event 0xb6aa5c479f96dcbdbe1d8b70883ec95e8799f8e602ee347ed972a22c974b14c0.
//
// Solidity: event ConfigurationChanged()
func (_Governance *GovernanceFilterer) WatchConfigurationChanged(opts *bind.WatchOpts, sink chan<- *GovernanceConfigurationChanged) (event.Subscription, error) {

	logs, sub, err := _Governance.contract.WatchLogs(opts, "ConfigurationChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GovernanceConfigurationChanged)
				if err := _Governance.contract.UnpackLog(event, "ConfigurationChanged", log); err != nil {
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

// ParseConfigurationChanged is a log parse operation binding the contract event 0xb6aa5c479f96dcbdbe1d8b70883ec95e8799f8e602ee347ed972a22c974b14c0.
//
// Solidity: event ConfigurationChanged()
func (_Governance *GovernanceFilterer) ParseConfigurationChanged(log types.Log) (*GovernanceConfigurationChanged, error) {
	event := new(GovernanceConfigurationChanged)
	if err := _Governance.contract.UnpackLog(event, "ConfigurationChanged", log); err != nil {
		return nil, err
	}
	return event, nil
}

// GovernanceDKGResetIterator is returned from FilterDKGReset and is used to iterate over the raw logs and unpacked data for DKGReset events raised by the Governance contract.
type GovernanceDKGResetIterator struct {
	Event *GovernanceDKGReset // Event containing the contract specifics and raw log

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
func (it *GovernanceDKGResetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GovernanceDKGReset)
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
		it.Event = new(GovernanceDKGReset)
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
func (it *GovernanceDKGResetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GovernanceDKGResetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GovernanceDKGReset represents a DKGReset event raised by the Governance contract.
type GovernanceDKGReset struct {
	Round       *big.Int
	BlockHeight *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterDKGReset is a free log retrieval operation binding the contract event 0x1ef9b10077100d06aefe929eb17ff481feebce38b8a770e5b7d744febfab9f50.
//
// Solidity: event DKGReset(uint256 indexed Round, uint256 BlockHeight)
func (_Governance *GovernanceFilterer) FilterDKGReset(opts *bind.FilterOpts, Round []*big.Int) (*GovernanceDKGResetIterator, error) {

	var RoundRule []interface{}
	for _, RoundItem := range Round {
		RoundRule = append(RoundRule, RoundItem)
	}

	logs, sub, err := _Governance.contract.FilterLogs(opts, "DKGReset", RoundRule)
	if err != nil {
		return nil, err
	}
	return &GovernanceDKGResetIterator{contract: _Governance.contract, event: "DKGReset", logs: logs, sub: sub}, nil
}

// WatchDKGReset is a free log subscription operation binding the contract event 0x1ef9b10077100d06aefe929eb17ff481feebce38b8a770e5b7d744febfab9f50.
//
// Solidity: event DKGReset(uint256 indexed Round, uint256 BlockHeight)
func (_Governance *GovernanceFilterer) WatchDKGReset(opts *bind.WatchOpts, sink chan<- *GovernanceDKGReset, Round []*big.Int) (event.Subscription, error) {

	var RoundRule []interface{}
	for _, RoundItem := range Round {
		RoundRule = append(RoundRule, RoundItem)
	}

	logs, sub, err := _Governance.contract.WatchLogs(opts, "DKGReset", RoundRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GovernanceDKGReset)
				if err := _Governance.contract.UnpackLog(event, "DKGReset", log); err != nil {
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

// ParseDKGReset is a log parse operation binding the contract event 0x1ef9b10077100d06aefe929eb17ff481feebce38b8a770e5b7d744febfab9f50.
//
// Solidity: event DKGReset(uint256 indexed Round, uint256 BlockHeight)
func (_Governance *GovernanceFilterer) ParseDKGReset(log types.Log) (*GovernanceDKGReset, error) {
	event := new(GovernanceDKGReset)
	if err := _Governance.contract.UnpackLog(event, "DKGReset", log); err != nil {
		return nil, err
	}
	return event, nil
}

// GovernanceFinePaidIterator is returned from FilterFinePaid and is used to iterate over the raw logs and unpacked data for FinePaid events raised by the Governance contract.
type GovernanceFinePaidIterator struct {
	Event *GovernanceFinePaid // Event containing the contract specifics and raw log

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
func (it *GovernanceFinePaidIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GovernanceFinePaid)
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
		it.Event = new(GovernanceFinePaid)
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
func (it *GovernanceFinePaidIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GovernanceFinePaidIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GovernanceFinePaid represents a FinePaid event raised by the Governance contract.
type GovernanceFinePaid struct {
	NodeAddress common.Address
	Amount      *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterFinePaid is a free log retrieval operation binding the contract event 0x34af8d99a30dafb3157ed162369d603f544a74be3858f8c7ac9e159829589c25.
//
// Solidity: event FinePaid(address indexed NodeAddress, uint256 Amount)
func (_Governance *GovernanceFilterer) FilterFinePaid(opts *bind.FilterOpts, NodeAddress []common.Address) (*GovernanceFinePaidIterator, error) {

	var NodeAddressRule []interface{}
	for _, NodeAddressItem := range NodeAddress {
		NodeAddressRule = append(NodeAddressRule, NodeAddressItem)
	}

	logs, sub, err := _Governance.contract.FilterLogs(opts, "FinePaid", NodeAddressRule)
	if err != nil {
		return nil, err
	}
	return &GovernanceFinePaidIterator{contract: _Governance.contract, event: "FinePaid", logs: logs, sub: sub}, nil
}

// WatchFinePaid is a free log subscription operation binding the contract event 0x34af8d99a30dafb3157ed162369d603f544a74be3858f8c7ac9e159829589c25.
//
// Solidity: event FinePaid(address indexed NodeAddress, uint256 Amount)
func (_Governance *GovernanceFilterer) WatchFinePaid(opts *bind.WatchOpts, sink chan<- *GovernanceFinePaid, NodeAddress []common.Address) (event.Subscription, error) {

	var NodeAddressRule []interface{}
	for _, NodeAddressItem := range NodeAddress {
		NodeAddressRule = append(NodeAddressRule, NodeAddressItem)
	}

	logs, sub, err := _Governance.contract.WatchLogs(opts, "FinePaid", NodeAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GovernanceFinePaid)
				if err := _Governance.contract.UnpackLog(event, "FinePaid", log); err != nil {
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

// ParseFinePaid is a log parse operation binding the contract event 0x34af8d99a30dafb3157ed162369d603f544a74be3858f8c7ac9e159829589c25.
//
// Solidity: event FinePaid(address indexed NodeAddress, uint256 Amount)
func (_Governance *GovernanceFilterer) ParseFinePaid(log types.Log) (*GovernanceFinePaid, error) {
	event := new(GovernanceFinePaid)
	if err := _Governance.contract.UnpackLog(event, "FinePaid", log); err != nil {
		return nil, err
	}
	return event, nil
}

// GovernanceFinedIterator is returned from FilterFined and is used to iterate over the raw logs and unpacked data for Fined events raised by the Governance contract.
type GovernanceFinedIterator struct {
	Event *GovernanceFined // Event containing the contract specifics and raw log

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
func (it *GovernanceFinedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GovernanceFined)
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
		it.Event = new(GovernanceFined)
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
func (it *GovernanceFinedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GovernanceFinedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GovernanceFined represents a Fined event raised by the Governance contract.
type GovernanceFined struct {
	NodeAddress common.Address
	Amount      *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterFined is a free log retrieval operation binding the contract event 0x00913d46aef0f0d115d70ea1c7c23198505f577d1d1916cc60710ca2204ae6ae.
//
// Solidity: event Fined(address indexed NodeAddress, uint256 Amount)
func (_Governance *GovernanceFilterer) FilterFined(opts *bind.FilterOpts, NodeAddress []common.Address) (*GovernanceFinedIterator, error) {

	var NodeAddressRule []interface{}
	for _, NodeAddressItem := range NodeAddress {
		NodeAddressRule = append(NodeAddressRule, NodeAddressItem)
	}

	logs, sub, err := _Governance.contract.FilterLogs(opts, "Fined", NodeAddressRule)
	if err != nil {
		return nil, err
	}
	return &GovernanceFinedIterator{contract: _Governance.contract, event: "Fined", logs: logs, sub: sub}, nil
}

// WatchFined is a free log subscription operation binding the contract event 0x00913d46aef0f0d115d70ea1c7c23198505f577d1d1916cc60710ca2204ae6ae.
//
// Solidity: event Fined(address indexed NodeAddress, uint256 Amount)
func (_Governance *GovernanceFilterer) WatchFined(opts *bind.WatchOpts, sink chan<- *GovernanceFined, NodeAddress []common.Address) (event.Subscription, error) {

	var NodeAddressRule []interface{}
	for _, NodeAddressItem := range NodeAddress {
		NodeAddressRule = append(NodeAddressRule, NodeAddressItem)
	}

	logs, sub, err := _Governance.contract.WatchLogs(opts, "Fined", NodeAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GovernanceFined)
				if err := _Governance.contract.UnpackLog(event, "Fined", log); err != nil {
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

// ParseFined is a log parse operation binding the contract event 0x00913d46aef0f0d115d70ea1c7c23198505f577d1d1916cc60710ca2204ae6ae.
//
// Solidity: event Fined(address indexed NodeAddress, uint256 Amount)
func (_Governance *GovernanceFilterer) ParseFined(log types.Log) (*GovernanceFined, error) {
	event := new(GovernanceFined)
	if err := _Governance.contract.UnpackLog(event, "Fined", log); err != nil {
		return nil, err
	}
	return event, nil
}

// GovernanceNodeAddedIterator is returned from FilterNodeAdded and is used to iterate over the raw logs and unpacked data for NodeAdded events raised by the Governance contract.
type GovernanceNodeAddedIterator struct {
	Event *GovernanceNodeAdded // Event containing the contract specifics and raw log

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
func (it *GovernanceNodeAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GovernanceNodeAdded)
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
		it.Event = new(GovernanceNodeAdded)
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
func (it *GovernanceNodeAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GovernanceNodeAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GovernanceNodeAdded represents a NodeAdded event raised by the Governance contract.
type GovernanceNodeAdded struct {
	NodeAddress common.Address
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterNodeAdded is a free log retrieval operation binding the contract event 0xb25d03aaf308d7291709be1ea28b800463cf3a9a4c4a5555d7333a964c1dfebd.
//
// Solidity: event NodeAdded(address indexed NodeAddress)
func (_Governance *GovernanceFilterer) FilterNodeAdded(opts *bind.FilterOpts, NodeAddress []common.Address) (*GovernanceNodeAddedIterator, error) {

	var NodeAddressRule []interface{}
	for _, NodeAddressItem := range NodeAddress {
		NodeAddressRule = append(NodeAddressRule, NodeAddressItem)
	}

	logs, sub, err := _Governance.contract.FilterLogs(opts, "NodeAdded", NodeAddressRule)
	if err != nil {
		return nil, err
	}
	return &GovernanceNodeAddedIterator{contract: _Governance.contract, event: "NodeAdded", logs: logs, sub: sub}, nil
}

// WatchNodeAdded is a free log subscription operation binding the contract event 0xb25d03aaf308d7291709be1ea28b800463cf3a9a4c4a5555d7333a964c1dfebd.
//
// Solidity: event NodeAdded(address indexed NodeAddress)
func (_Governance *GovernanceFilterer) WatchNodeAdded(opts *bind.WatchOpts, sink chan<- *GovernanceNodeAdded, NodeAddress []common.Address) (event.Subscription, error) {

	var NodeAddressRule []interface{}
	for _, NodeAddressItem := range NodeAddress {
		NodeAddressRule = append(NodeAddressRule, NodeAddressItem)
	}

	logs, sub, err := _Governance.contract.WatchLogs(opts, "NodeAdded", NodeAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GovernanceNodeAdded)
				if err := _Governance.contract.UnpackLog(event, "NodeAdded", log); err != nil {
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

// ParseNodeAdded is a log parse operation binding the contract event 0xb25d03aaf308d7291709be1ea28b800463cf3a9a4c4a5555d7333a964c1dfebd.
//
// Solidity: event NodeAdded(address indexed NodeAddress)
func (_Governance *GovernanceFilterer) ParseNodeAdded(log types.Log) (*GovernanceNodeAdded, error) {
	event := new(GovernanceNodeAdded)
	if err := _Governance.contract.UnpackLog(event, "NodeAdded", log); err != nil {
		return nil, err
	}
	return event, nil
}

// GovernanceNodeOwnershipTransferedIterator is returned from FilterNodeOwnershipTransfered and is used to iterate over the raw logs and unpacked data for NodeOwnershipTransfered events raised by the Governance contract.
type GovernanceNodeOwnershipTransferedIterator struct {
	Event *GovernanceNodeOwnershipTransfered // Event containing the contract specifics and raw log

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
func (it *GovernanceNodeOwnershipTransferedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GovernanceNodeOwnershipTransfered)
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
		it.Event = new(GovernanceNodeOwnershipTransfered)
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
func (it *GovernanceNodeOwnershipTransferedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GovernanceNodeOwnershipTransferedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GovernanceNodeOwnershipTransfered represents a NodeOwnershipTransfered event raised by the Governance contract.
type GovernanceNodeOwnershipTransfered struct {
	NodeAddress     common.Address
	NewOwnerAddress common.Address
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterNodeOwnershipTransfered is a free log retrieval operation binding the contract event 0x63616cf4caeee172ed91719b5e000da424d56855ee726aedb0a1d74fb3f1db2a.
//
// Solidity: event NodeOwnershipTransfered(address indexed NodeAddress, address indexed NewOwnerAddress)
func (_Governance *GovernanceFilterer) FilterNodeOwnershipTransfered(opts *bind.FilterOpts, NodeAddress []common.Address, NewOwnerAddress []common.Address) (*GovernanceNodeOwnershipTransferedIterator, error) {

	var NodeAddressRule []interface{}
	for _, NodeAddressItem := range NodeAddress {
		NodeAddressRule = append(NodeAddressRule, NodeAddressItem)
	}
	var NewOwnerAddressRule []interface{}
	for _, NewOwnerAddressItem := range NewOwnerAddress {
		NewOwnerAddressRule = append(NewOwnerAddressRule, NewOwnerAddressItem)
	}

	logs, sub, err := _Governance.contract.FilterLogs(opts, "NodeOwnershipTransfered", NodeAddressRule, NewOwnerAddressRule)
	if err != nil {
		return nil, err
	}
	return &GovernanceNodeOwnershipTransferedIterator{contract: _Governance.contract, event: "NodeOwnershipTransfered", logs: logs, sub: sub}, nil
}

// WatchNodeOwnershipTransfered is a free log subscription operation binding the contract event 0x63616cf4caeee172ed91719b5e000da424d56855ee726aedb0a1d74fb3f1db2a.
//
// Solidity: event NodeOwnershipTransfered(address indexed NodeAddress, address indexed NewOwnerAddress)
func (_Governance *GovernanceFilterer) WatchNodeOwnershipTransfered(opts *bind.WatchOpts, sink chan<- *GovernanceNodeOwnershipTransfered, NodeAddress []common.Address, NewOwnerAddress []common.Address) (event.Subscription, error) {

	var NodeAddressRule []interface{}
	for _, NodeAddressItem := range NodeAddress {
		NodeAddressRule = append(NodeAddressRule, NodeAddressItem)
	}
	var NewOwnerAddressRule []interface{}
	for _, NewOwnerAddressItem := range NewOwnerAddress {
		NewOwnerAddressRule = append(NewOwnerAddressRule, NewOwnerAddressItem)
	}

	logs, sub, err := _Governance.contract.WatchLogs(opts, "NodeOwnershipTransfered", NodeAddressRule, NewOwnerAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GovernanceNodeOwnershipTransfered)
				if err := _Governance.contract.UnpackLog(event, "NodeOwnershipTransfered", log); err != nil {
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

// ParseNodeOwnershipTransfered is a log parse operation binding the contract event 0x63616cf4caeee172ed91719b5e000da424d56855ee726aedb0a1d74fb3f1db2a.
//
// Solidity: event NodeOwnershipTransfered(address indexed NodeAddress, address indexed NewOwnerAddress)
func (_Governance *GovernanceFilterer) ParseNodeOwnershipTransfered(log types.Log) (*GovernanceNodeOwnershipTransfered, error) {
	event := new(GovernanceNodeOwnershipTransfered)
	if err := _Governance.contract.UnpackLog(event, "NodeOwnershipTransfered", log); err != nil {
		return nil, err
	}
	return event, nil
}

// GovernanceNodePublicKeyReplacedIterator is returned from FilterNodePublicKeyReplaced and is used to iterate over the raw logs and unpacked data for NodePublicKeyReplaced events raised by the Governance contract.
type GovernanceNodePublicKeyReplacedIterator struct {
	Event *GovernanceNodePublicKeyReplaced // Event containing the contract specifics and raw log

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
func (it *GovernanceNodePublicKeyReplacedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GovernanceNodePublicKeyReplaced)
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
		it.Event = new(GovernanceNodePublicKeyReplaced)
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
func (it *GovernanceNodePublicKeyReplacedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GovernanceNodePublicKeyReplacedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GovernanceNodePublicKeyReplaced represents a NodePublicKeyReplaced event raised by the Governance contract.
type GovernanceNodePublicKeyReplaced struct {
	NodeAddress common.Address
	PublicKey   []byte
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterNodePublicKeyReplaced is a free log retrieval operation binding the contract event 0x8aa688c270fa0768c37a1f455f711b6ea1a925f04db9ef433e01dcba78b80432.
//
// Solidity: event NodePublicKeyReplaced(address indexed NodeAddress, bytes PublicKey)
func (_Governance *GovernanceFilterer) FilterNodePublicKeyReplaced(opts *bind.FilterOpts, NodeAddress []common.Address) (*GovernanceNodePublicKeyReplacedIterator, error) {

	var NodeAddressRule []interface{}
	for _, NodeAddressItem := range NodeAddress {
		NodeAddressRule = append(NodeAddressRule, NodeAddressItem)
	}

	logs, sub, err := _Governance.contract.FilterLogs(opts, "NodePublicKeyReplaced", NodeAddressRule)
	if err != nil {
		return nil, err
	}
	return &GovernanceNodePublicKeyReplacedIterator{contract: _Governance.contract, event: "NodePublicKeyReplaced", logs: logs, sub: sub}, nil
}

// WatchNodePublicKeyReplaced is a free log subscription operation binding the contract event 0x8aa688c270fa0768c37a1f455f711b6ea1a925f04db9ef433e01dcba78b80432.
//
// Solidity: event NodePublicKeyReplaced(address indexed NodeAddress, bytes PublicKey)
func (_Governance *GovernanceFilterer) WatchNodePublicKeyReplaced(opts *bind.WatchOpts, sink chan<- *GovernanceNodePublicKeyReplaced, NodeAddress []common.Address) (event.Subscription, error) {

	var NodeAddressRule []interface{}
	for _, NodeAddressItem := range NodeAddress {
		NodeAddressRule = append(NodeAddressRule, NodeAddressItem)
	}

	logs, sub, err := _Governance.contract.WatchLogs(opts, "NodePublicKeyReplaced", NodeAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GovernanceNodePublicKeyReplaced)
				if err := _Governance.contract.UnpackLog(event, "NodePublicKeyReplaced", log); err != nil {
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

// ParseNodePublicKeyReplaced is a log parse operation binding the contract event 0x8aa688c270fa0768c37a1f455f711b6ea1a925f04db9ef433e01dcba78b80432.
//
// Solidity: event NodePublicKeyReplaced(address indexed NodeAddress, bytes PublicKey)
func (_Governance *GovernanceFilterer) ParseNodePublicKeyReplaced(log types.Log) (*GovernanceNodePublicKeyReplaced, error) {
	event := new(GovernanceNodePublicKeyReplaced)
	if err := _Governance.contract.UnpackLog(event, "NodePublicKeyReplaced", log); err != nil {
		return nil, err
	}
	return event, nil
}

// GovernanceNodeRemovedIterator is returned from FilterNodeRemoved and is used to iterate over the raw logs and unpacked data for NodeRemoved events raised by the Governance contract.
type GovernanceNodeRemovedIterator struct {
	Event *GovernanceNodeRemoved // Event containing the contract specifics and raw log

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
func (it *GovernanceNodeRemovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GovernanceNodeRemoved)
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
		it.Event = new(GovernanceNodeRemoved)
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
func (it *GovernanceNodeRemovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GovernanceNodeRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GovernanceNodeRemoved represents a NodeRemoved event raised by the Governance contract.
type GovernanceNodeRemoved struct {
	NodeAddress common.Address
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterNodeRemoved is a free log retrieval operation binding the contract event 0xcfc24166db4bb677e857cacabd1541fb2b30645021b27c5130419589b84db52b.
//
// Solidity: event NodeRemoved(address indexed NodeAddress)
func (_Governance *GovernanceFilterer) FilterNodeRemoved(opts *bind.FilterOpts, NodeAddress []common.Address) (*GovernanceNodeRemovedIterator, error) {

	var NodeAddressRule []interface{}
	for _, NodeAddressItem := range NodeAddress {
		NodeAddressRule = append(NodeAddressRule, NodeAddressItem)
	}

	logs, sub, err := _Governance.contract.FilterLogs(opts, "NodeRemoved", NodeAddressRule)
	if err != nil {
		return nil, err
	}
	return &GovernanceNodeRemovedIterator{contract: _Governance.contract, event: "NodeRemoved", logs: logs, sub: sub}, nil
}

// WatchNodeRemoved is a free log subscription operation binding the contract event 0xcfc24166db4bb677e857cacabd1541fb2b30645021b27c5130419589b84db52b.
//
// Solidity: event NodeRemoved(address indexed NodeAddress)
func (_Governance *GovernanceFilterer) WatchNodeRemoved(opts *bind.WatchOpts, sink chan<- *GovernanceNodeRemoved, NodeAddress []common.Address) (event.Subscription, error) {

	var NodeAddressRule []interface{}
	for _, NodeAddressItem := range NodeAddress {
		NodeAddressRule = append(NodeAddressRule, NodeAddressItem)
	}

	logs, sub, err := _Governance.contract.WatchLogs(opts, "NodeRemoved", NodeAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GovernanceNodeRemoved)
				if err := _Governance.contract.UnpackLog(event, "NodeRemoved", log); err != nil {
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

// ParseNodeRemoved is a log parse operation binding the contract event 0xcfc24166db4bb677e857cacabd1541fb2b30645021b27c5130419589b84db52b.
//
// Solidity: event NodeRemoved(address indexed NodeAddress)
func (_Governance *GovernanceFilterer) ParseNodeRemoved(log types.Log) (*GovernanceNodeRemoved, error) {
	event := new(GovernanceNodeRemoved)
	if err := _Governance.contract.UnpackLog(event, "NodeRemoved", log); err != nil {
		return nil, err
	}
	return event, nil
}

// GovernanceReportedIterator is returned from FilterReported and is used to iterate over the raw logs and unpacked data for Reported events raised by the Governance contract.
type GovernanceReportedIterator struct {
	Event *GovernanceReported // Event containing the contract specifics and raw log

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
func (it *GovernanceReportedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GovernanceReported)
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
		it.Event = new(GovernanceReported)
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
func (it *GovernanceReportedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GovernanceReportedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GovernanceReported represents a Reported event raised by the Governance contract.
type GovernanceReported struct {
	NodeAddress common.Address
	Type        *big.Int
	Arg1        []byte
	Arg2        []byte
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterReported is a free log retrieval operation binding the contract event 0x0b4ded2614643cbc6be3ac9a1ed3188f2667636de75ea5250e190ad6af4b43d9.
//
// Solidity: event Reported(address indexed NodeAddress, uint256 Type, bytes Arg1, bytes Arg2)
func (_Governance *GovernanceFilterer) FilterReported(opts *bind.FilterOpts, NodeAddress []common.Address) (*GovernanceReportedIterator, error) {

	var NodeAddressRule []interface{}
	for _, NodeAddressItem := range NodeAddress {
		NodeAddressRule = append(NodeAddressRule, NodeAddressItem)
	}

	logs, sub, err := _Governance.contract.FilterLogs(opts, "Reported", NodeAddressRule)
	if err != nil {
		return nil, err
	}
	return &GovernanceReportedIterator{contract: _Governance.contract, event: "Reported", logs: logs, sub: sub}, nil
}

// WatchReported is a free log subscription operation binding the contract event 0x0b4ded2614643cbc6be3ac9a1ed3188f2667636de75ea5250e190ad6af4b43d9.
//
// Solidity: event Reported(address indexed NodeAddress, uint256 Type, bytes Arg1, bytes Arg2)
func (_Governance *GovernanceFilterer) WatchReported(opts *bind.WatchOpts, sink chan<- *GovernanceReported, NodeAddress []common.Address) (event.Subscription, error) {

	var NodeAddressRule []interface{}
	for _, NodeAddressItem := range NodeAddress {
		NodeAddressRule = append(NodeAddressRule, NodeAddressItem)
	}

	logs, sub, err := _Governance.contract.WatchLogs(opts, "Reported", NodeAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GovernanceReported)
				if err := _Governance.contract.UnpackLog(event, "Reported", log); err != nil {
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

// ParseReported is a log parse operation binding the contract event 0x0b4ded2614643cbc6be3ac9a1ed3188f2667636de75ea5250e190ad6af4b43d9.
//
// Solidity: event Reported(address indexed NodeAddress, uint256 Type, bytes Arg1, bytes Arg2)
func (_Governance *GovernanceFilterer) ParseReported(log types.Log) (*GovernanceReported, error) {
	event := new(GovernanceReported)
	if err := _Governance.contract.UnpackLog(event, "Reported", log); err != nil {
		return nil, err
	}
	return event, nil
}

// GovernanceStakedIterator is returned from FilterStaked and is used to iterate over the raw logs and unpacked data for Staked events raised by the Governance contract.
type GovernanceStakedIterator struct {
	Event *GovernanceStaked // Event containing the contract specifics and raw log

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
func (it *GovernanceStakedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GovernanceStaked)
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
		it.Event = new(GovernanceStaked)
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
func (it *GovernanceStakedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GovernanceStakedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GovernanceStaked represents a Staked event raised by the Governance contract.
type GovernanceStaked struct {
	NodeAddress common.Address
	Amount      *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterStaked is a free log retrieval operation binding the contract event 0x9e71bc8eea02a63969f509818f2dafb9254532904319f9dbda79b67bd34a5f3d.
//
// Solidity: event Staked(address indexed NodeAddress, uint256 Amount)
func (_Governance *GovernanceFilterer) FilterStaked(opts *bind.FilterOpts, NodeAddress []common.Address) (*GovernanceStakedIterator, error) {

	var NodeAddressRule []interface{}
	for _, NodeAddressItem := range NodeAddress {
		NodeAddressRule = append(NodeAddressRule, NodeAddressItem)
	}

	logs, sub, err := _Governance.contract.FilterLogs(opts, "Staked", NodeAddressRule)
	if err != nil {
		return nil, err
	}
	return &GovernanceStakedIterator{contract: _Governance.contract, event: "Staked", logs: logs, sub: sub}, nil
}

// WatchStaked is a free log subscription operation binding the contract event 0x9e71bc8eea02a63969f509818f2dafb9254532904319f9dbda79b67bd34a5f3d.
//
// Solidity: event Staked(address indexed NodeAddress, uint256 Amount)
func (_Governance *GovernanceFilterer) WatchStaked(opts *bind.WatchOpts, sink chan<- *GovernanceStaked, NodeAddress []common.Address) (event.Subscription, error) {

	var NodeAddressRule []interface{}
	for _, NodeAddressItem := range NodeAddress {
		NodeAddressRule = append(NodeAddressRule, NodeAddressItem)
	}

	logs, sub, err := _Governance.contract.WatchLogs(opts, "Staked", NodeAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GovernanceStaked)
				if err := _Governance.contract.UnpackLog(event, "Staked", log); err != nil {
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

// ParseStaked is a log parse operation binding the contract event 0x9e71bc8eea02a63969f509818f2dafb9254532904319f9dbda79b67bd34a5f3d.
//
// Solidity: event Staked(address indexed NodeAddress, uint256 Amount)
func (_Governance *GovernanceFilterer) ParseStaked(log types.Log) (*GovernanceStaked, error) {
	event := new(GovernanceStaked)
	if err := _Governance.contract.UnpackLog(event, "Staked", log); err != nil {
		return nil, err
	}
	return event, nil
}

// GovernanceUnstakedIterator is returned from FilterUnstaked and is used to iterate over the raw logs and unpacked data for Unstaked events raised by the Governance contract.
type GovernanceUnstakedIterator struct {
	Event *GovernanceUnstaked // Event containing the contract specifics and raw log

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
func (it *GovernanceUnstakedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GovernanceUnstaked)
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
		it.Event = new(GovernanceUnstaked)
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
func (it *GovernanceUnstakedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GovernanceUnstakedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GovernanceUnstaked represents a Unstaked event raised by the Governance contract.
type GovernanceUnstaked struct {
	NodeAddress common.Address
	Amount      *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterUnstaked is a free log retrieval operation binding the contract event 0x0f5bb82176feb1b5e747e28471aa92156a04d9f3ab9f45f28e2d704232b93f75.
//
// Solidity: event Unstaked(address indexed NodeAddress, uint256 Amount)
func (_Governance *GovernanceFilterer) FilterUnstaked(opts *bind.FilterOpts, NodeAddress []common.Address) (*GovernanceUnstakedIterator, error) {

	var NodeAddressRule []interface{}
	for _, NodeAddressItem := range NodeAddress {
		NodeAddressRule = append(NodeAddressRule, NodeAddressItem)
	}

	logs, sub, err := _Governance.contract.FilterLogs(opts, "Unstaked", NodeAddressRule)
	if err != nil {
		return nil, err
	}
	return &GovernanceUnstakedIterator{contract: _Governance.contract, event: "Unstaked", logs: logs, sub: sub}, nil
}

// WatchUnstaked is a free log subscription operation binding the contract event 0x0f5bb82176feb1b5e747e28471aa92156a04d9f3ab9f45f28e2d704232b93f75.
//
// Solidity: event Unstaked(address indexed NodeAddress, uint256 Amount)
func (_Governance *GovernanceFilterer) WatchUnstaked(opts *bind.WatchOpts, sink chan<- *GovernanceUnstaked, NodeAddress []common.Address) (event.Subscription, error) {

	var NodeAddressRule []interface{}
	for _, NodeAddressItem := range NodeAddress {
		NodeAddressRule = append(NodeAddressRule, NodeAddressItem)
	}

	logs, sub, err := _Governance.contract.WatchLogs(opts, "Unstaked", NodeAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GovernanceUnstaked)
				if err := _Governance.contract.UnpackLog(event, "Unstaked", log); err != nil {
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

// ParseUnstaked is a log parse operation binding the contract event 0x0f5bb82176feb1b5e747e28471aa92156a04d9f3ab9f45f28e2d704232b93f75.
//
// Solidity: event Unstaked(address indexed NodeAddress, uint256 Amount)
func (_Governance *GovernanceFilterer) ParseUnstaked(log types.Log) (*GovernanceUnstaked, error) {
	event := new(GovernanceUnstaked)
	if err := _Governance.contract.UnpackLog(event, "Unstaked", log); err != nil {
		return nil, err
	}
	return event, nil
}

// GovernanceWithdrawnIterator is returned from FilterWithdrawn and is used to iterate over the raw logs and unpacked data for Withdrawn events raised by the Governance contract.
type GovernanceWithdrawnIterator struct {
	Event *GovernanceWithdrawn // Event containing the contract specifics and raw log

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
func (it *GovernanceWithdrawnIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GovernanceWithdrawn)
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
		it.Event = new(GovernanceWithdrawn)
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
func (it *GovernanceWithdrawnIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GovernanceWithdrawnIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GovernanceWithdrawn represents a Withdrawn event raised by the Governance contract.
type GovernanceWithdrawn struct {
	NodeAddress common.Address
	Amount      *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterWithdrawn is a free log retrieval operation binding the contract event 0x7084f5476618d8e60b11ef0d7d3f06914655adb8793e28ff7f018d4c76d505d5.
//
// Solidity: event Withdrawn(address indexed NodeAddress, uint256 Amount)
func (_Governance *GovernanceFilterer) FilterWithdrawn(opts *bind.FilterOpts, NodeAddress []common.Address) (*GovernanceWithdrawnIterator, error) {

	var NodeAddressRule []interface{}
	for _, NodeAddressItem := range NodeAddress {
		NodeAddressRule = append(NodeAddressRule, NodeAddressItem)
	}

	logs, sub, err := _Governance.contract.FilterLogs(opts, "Withdrawn", NodeAddressRule)
	if err != nil {
		return nil, err
	}
	return &GovernanceWithdrawnIterator{contract: _Governance.contract, event: "Withdrawn", logs: logs, sub: sub}, nil
}

// WatchWithdrawn is a free log subscription operation binding the contract event 0x7084f5476618d8e60b11ef0d7d3f06914655adb8793e28ff7f018d4c76d505d5.
//
// Solidity: event Withdrawn(address indexed NodeAddress, uint256 Amount)
func (_Governance *GovernanceFilterer) WatchWithdrawn(opts *bind.WatchOpts, sink chan<- *GovernanceWithdrawn, NodeAddress []common.Address) (event.Subscription, error) {

	var NodeAddressRule []interface{}
	for _, NodeAddressItem := range NodeAddress {
		NodeAddressRule = append(NodeAddressRule, NodeAddressItem)
	}

	logs, sub, err := _Governance.contract.WatchLogs(opts, "Withdrawn", NodeAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GovernanceWithdrawn)
				if err := _Governance.contract.UnpackLog(event, "Withdrawn", log); err != nil {
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

// ParseWithdrawn is a log parse operation binding the contract event 0x7084f5476618d8e60b11ef0d7d3f06914655adb8793e28ff7f018d4c76d505d5.
//
// Solidity: event Withdrawn(address indexed NodeAddress, uint256 Amount)
func (_Governance *GovernanceFilterer) ParseWithdrawn(log types.Log) (*GovernanceWithdrawn, error) {
	event := new(GovernanceWithdrawn)
	if err := _Governance.contract.UnpackLog(event, "Withdrawn", log); err != nil {
		return nil, err
	}
	return event, nil
}
