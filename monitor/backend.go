package monitor

import (
	"context"
	"fmt"
	"log"
	"math/big"
	"strings"
	"time"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

type backendIntf interface {
	FetchFinedNodes(chan node)
	NodeSet() []node
	BalanceFromAddress(*ethclient.Client, common.Address) *big.Int
}

// BlockchainBackend is the blockchain backend object for fetching blockchain info.
type BlockchainBackend struct {
	networkID int
}

// NodeSet returns the registered nodes.
func (b *BlockchainBackend) NodeSet() []node {
	nc := NetworkConfigMap[int64(b.networkID)]
	conn, err := ethclient.Dial(nc.HTTPEndpoint)
	if err != nil {
		log.Println("EthClient Dial failed: ", err)
		return []node{}
	}
	g, err := NewGovernance(GovAddress, conn)
	if err != nil {
		panic(err)
	}
	txOps := &bind.CallOpts{}
	nodeLen, err := g.NodesLength(txOps)
	if err != nil {
		panic(err)
	}
	minStake, err := g.MinStake(txOps)
	if err != nil {
		panic(err)
	}
	nodes := []node{}
	for i := uint64(0); i < nodeLen.Uint64(); i++ {
		n, err := g.Nodes(txOps, big.NewInt(int64(i)))
		if n.Staked.Cmp(minStake) < 0 {
			continue
		}
		if err != nil {
			panic(err)
		}
		publicKey, err := crypto.UnmarshalPubkey(n.PublicKey)
		if err != nil {
			log.Println("Unmarshal publicKey fail: ", err)
			continue
		}
		address := crypto.PubkeyToAddress(*publicKey)
		nodes = append(nodes, node{
			owner:          n.Owner,
			email:          n.Email,
			fined:          n.Fined,
			name:           n.Name,
			publicKey:      n.PublicKey,
			nodeKeyAddress: address,
		})
	}
	return nodes
}

// BalanceFromAddress return address's balance.
func (b *BlockchainBackend) BalanceFromAddress(conn *ethclient.Client,
	address common.Address) *big.Int {
	balance, err := conn.BalanceAt(context.Background(),
		address, nil)
	if err != nil {
		log.Println("Get balance fail: ", err)
		return big.NewInt(0)
	}
	return balance
}

func (b *BlockchainBackend) subscribe(
	logs chan types.Log,
	headers chan *types.Header,
) (ethereum.Subscription, *ethclient.Client, error) {

	nc := NetworkConfigMap[int64(b.networkID)]
	conn, err := ethclient.Dial(nc.WSEndpoint)
	if err != nil {
		log.Println("EthClient Dial failed: ", err)
		return nil, conn, err
	}
	abiObject, err := abi.JSON(strings.NewReader(GovernanceABI))
	query := ethereum.FilterQuery{
		Addresses: []common.Address{GovAddress},
		Topics: [][]common.Hash{
			[]common.Hash{
				abiObject.Events["Fined"].Id(),
			},
		},
	}
	sub, err := conn.SubscribeFilterLogs(context.Background(), query, logs)
	if err != nil {
		log.Println("Subscribe logs failed: ", err)
		return nil, conn, err
	}

	// Subscribe chain head to keep ws from timeout
	_, err = conn.SubscribeNewHead(context.Background(), headers)
	if err != nil {
		log.Println("Subscribe chain head failed: ", err)
	}
	return sub, conn, nil
}

// FetchFinedNodes fetches nodes' information from blockchain.
func (b *BlockchainBackend) FetchFinedNodes(nodeChan chan node) {
	logs := make(chan types.Log)
	headers := make(chan *types.Header)
	sub, conn, err := b.subscribe(logs, headers)
	if err != nil {
		panic(err)
	}
	for {
		select {
		case err := <-sub.Err():
			log.Println("sub.Err()", err)
			// retry
			for err != nil {
				sub, conn, err = b.subscribe(logs, headers)
				time.Sleep(30000)
			}
		case vLog := <-logs:
			nodeAddress := common.BytesToAddress(vLog.Topics[1].Bytes())
			log.Println("Notify node: ", nodeAddress.Hex())
			nodeChan <- b.getNodeByAddress(nodeAddress, conn)
		}
	}
}

func (b *BlockchainBackend) getNodeByAddress(
	nodeAddress common.Address, conn *ethclient.Client) node {

	g, err := NewGovernance(GovAddress, conn)
	if err != nil {
		panic(err)
	}
	txOps := &bind.CallOpts{}
	offset, err := g.NodesOffsetByAddress(txOps, nodeAddress)
	if err != nil {
		panic(err)
	}
	n, err := g.Nodes(txOps, offset)
	if err != nil {
		panic(err)
	}
	return node{
		owner: n.Owner,
		email: n.Email,
		fined: n.Fined,
		name:  n.Name,
	}
}

func printNodes(nodes []node) {
	for i := range nodes {
		fmt.Println("Owner", nodes[i].owner.Hex())
		fmt.Println("Email", nodes[i].email)
		fmt.Println("Fined", nodes[i].fined.Uint64())
	}
}

// NewBlockchainBackend is the blockchain backend constructor.
func NewBlockchainBackend(networkID int) *BlockchainBackend {
	return &BlockchainBackend{networkID: networkID}
}
