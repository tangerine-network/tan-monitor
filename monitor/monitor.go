package monitor

import (
	"log"
	"math"
	"math/big"
	"strconv"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

type node struct {
	owner          common.Address
	email          string
	fined          *big.Int
	name           string
	publicKey      []byte
	nodeKeyAddress common.Address
}

// NetworkConfig represents the network config.
type NetworkConfig struct {
	WSEndpoint      string
	HTTPEndpoint    string
	EthHTTPEndpoint string
	GovAddress      common.Address
	Network         string
}

// GovAddress is governance address.
var GovAddress = common.HexToAddress("0x246FcDE58581e2754f215A523C0718C4BFc8041F")

// Fifty is 50 * 10^18
var Fifty *big.Int

// OneHundred is 100 * 10^18
var OneHundred *big.Int

func init() {
	Fifty, _ = big.NewInt(0).SetString("50000000000000000000", 10)
	OneHundred, _ = big.NewInt(0).SetString("100000000000000000000", 10)
}

// NetworkConfigMap represents the system network config mapping.
var NetworkConfigMap = map[int64]NetworkConfig{
	411: {
		WSEndpoint:      "wss://mainnet-rpc.tangerine-network.io/ws",
		HTTPEndpoint:    "https://mainnet-rpc.tangerine-network.io",
		EthHTTPEndpoint: "https://mainnet.infura.io/v3/4eb07139b29d41c59b352f21c4c9f526",
		Network:         "Mainnet",
	},
	374: {
		WSEndpoint:      "wss://testnet-rpc.tangerine-network.io/ws",
		HTTPEndpoint:    "https://testnet-rpc.tangerine-network.io",
		EthHTTPEndpoint: "https://rinkeby.infura.io/v3/4eb07139b29d41c59b352f21c4c9f526",
		Network:         "Testnet",
	},
}

// Monitor is the object to monitor tangerine network.
type Monitor struct {
	networkID            int
	notifiers            []Notifier
	backend              backendIntf
	tanBalancesCache     map[common.Address]*big.Int
	ethBalancesCache     map[common.Address]*big.Int
	ethThreshold         *big.Int // in wei
	ethThresholdString   string   // in ETH
	checkBalanceDuration time.Duration
}

// NewMonitor is the constructor for monitor object.
func NewMonitor(networkID int, backend backendIntf, threshold string) *Monitor {
	tanBalance := make(map[common.Address]*big.Int)
	ethBalance := make(map[common.Address]*big.Int)
	t, err := strconv.ParseFloat(threshold, 64)
	if err != nil {
		panic(err)
	}
	t = t * math.Pow(10, 18)
	ethThreshold := new(big.Int)
	ethThreshold.SetString(
		strconv.FormatFloat(t, 'f', 0, 64),
		10,
	)
	m := Monitor{
		networkID:            networkID,
		backend:              backend,
		tanBalancesCache:     tanBalance,
		ethBalancesCache:     ethBalance,
		checkBalanceDuration: 10 * time.Minute,
		ethThresholdString:   threshold,
		ethThreshold:         ethThreshold,
	}
	return &m
}

// Run is the entry point for running monitor.
func (m *Monitor) Run() {
	done := make(chan bool)
	finedNodeChan := make(chan node)
	go m.fetchFinedNodes(finedNodeChan)
	go m.sendNotifications(finedNodeChan, FINED)
	go m.checkTanBalance()
	go m.checkEthBalance()
	<-done
}

func (m *Monitor) checkEthBalance() {
	ethNodeChan := make(chan node)
	go m.sendNotifications(ethNodeChan, INSUFFICIENT_ETH)
	for {
		nc := NetworkConfigMap[int64(m.networkID)]
		conn, err := ethclient.Dial(nc.EthHTTPEndpoint)
		if err != nil {
			log.Println("Get Tan balance fail at")
			continue
		}
		nodes := m.backend.NodeSet()
		for i := range nodes {
			n := nodes[i]
			address := n.nodeKeyAddress
			balance := m.backend.BalanceFromAddress(conn, address)
			if balance.Cmp(m.ethThreshold) < 0 {
				if cacheBalance, exist := m.ethBalancesCache[address]; exist {
					if cacheBalance.Cmp(m.ethThreshold) >= 0 {
						ethNodeChan <- n
					}
				} else {
					ethNodeChan <- n
				}
			}
			m.ethBalancesCache[address] = balance
		}
		time.Sleep(m.checkBalanceDuration)
	}
}

func (m *Monitor) checkTanBalance() {
	tanNodeChan := make(chan node)
	go m.sendNotifications(tanNodeChan, INSUFFICIENT_TAN)
	for {
		nodes := m.backend.NodeSet()
		nc := NetworkConfigMap[int64(m.networkID)]
		conn, err := ethclient.Dial(nc.HTTPEndpoint)
		if err != nil {
			log.Println("Get TAN balance fail")
			continue
		}
		for i := range nodes {
			n := nodes[i]
			balance := m.backend.BalanceFromAddress(conn, n.nodeKeyAddress)
			if balance.Cmp(Fifty) < 0 {
				if cacheBalance, exist := m.tanBalancesCache[n.nodeKeyAddress]; exist {
					if cacheBalance.Cmp(Fifty) >= 0 {
						tanNodeChan <- n
					}
				} else {
					tanNodeChan <- n
				}
			} else if balance.Cmp(OneHundred) < 0 {
				if cacheBalance, exist := m.tanBalancesCache[n.nodeKeyAddress]; exist {
					if cacheBalance.Cmp(OneHundred) >= 0 {
						tanNodeChan <- n
					}
				} else {
					tanNodeChan <- n
				}
			}
			m.tanBalancesCache[n.nodeKeyAddress] = balance
		}
		time.Sleep(m.checkBalanceDuration)
	}
}

// Register registries the notifiers.
func (m *Monitor) Register(n Notifier) {
	m.notifiers = append(m.notifiers, n)
}

func (m *Monitor) fetchFinedNodes(nodeChan chan node) {
	m.backend.FetchFinedNodes(nodeChan)
}

func (m *Monitor) sendNotifications(nodeChan chan node, notifyType uint) {
	nc := NetworkConfigMap[int64(m.networkID)]
	for {
		node, open := <-nodeChan
		if !open {
			return
		}
		for _, notifier := range m.notifiers {
			notifier.notify(node, nc.Network,
				notifyType, m.ethThresholdString)
		}
	}
}
