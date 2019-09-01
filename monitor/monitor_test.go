package monitor

import (
	"io/ioutil"
	"math/big"
	"os"
	"testing"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/stretchr/testify/suite"
)

type monitorTestSuite struct {
	suite.Suite
}

type mockBackend struct{}

func (m *mockBackend) FetchFinedNodes(nodeChan chan node) {
	nodeChan <- node{
		owner: common.Address{},
		email: "jm@byzantine-lab.io",
		fined: big.NewInt(1),
		name:  "unit test",
	}
	close(nodeChan)
}

func (m *mockBackend) NodeSet() []node {
	n := []node{
		{
			owner:          common.Address{},
			email:          "jm@byzantine-lab.io",
			fined:          big.NewInt(1),
			name:           "unit test 100, should not be fined",
			nodeKeyAddress: common.BytesToAddress([]byte{1}),
		},
		{
			owner:          common.Address{},
			email:          "jm@byzantine-lab.io",
			fined:          Fifty,
			name:           "unit test 50",
			publicKey:      []byte{2},
			nodeKeyAddress: common.BytesToAddress([]byte{2}),
		},
	}
	return n
}

func (m *mockBackend) BalanceFromAddress(conn *ethclient.Client,
	address common.Address) *big.Int {
	if address == common.BytesToAddress([]byte{1}) {
		return OneHundred
	}
	lessThanThreshold, _ := big.NewInt(0).SetString("99999999999999999", 10)
	return lessThanThreshold
}

func (s *monitorTestSuite) TestSendNotification() {
	backend := mockBackend{}
	m := NewMonitor(374, &backend, "0.1")
	nodeChan := make(chan node)
	go m.fetchFinedNodes(nodeChan)
	password := os.Getenv("TEST_EMAIL_PASSWORD")
	ccList, _ := ioutil.ReadFile("./cc-list.txt")
	email := NewEmail(
		"no-reply@byzantine-lab.io",
		password,
		"smtp-relay.gmail.com",
		string(ccList),
	)
	m.Register(email)
	m.sendNotifications(nodeChan, FINED)
	m.tanBalancesCache[common.BytesToAddress([]byte{1})] = OneHundred
	m.tanBalancesCache[common.BytesToAddress([]byte{2})] = Fifty
	m.ethBalancesCache[common.BytesToAddress([]byte{1})], _ = big.NewInt(0).SetString("99999999999999999", 10)
	m.ethBalancesCache[common.BytesToAddress([]byte{2})], _ = big.NewInt(0).SetString("100000000000000000", 10)

	go m.checkTanBalance()
	go m.checkEthBalance()
	time.Sleep(5000 * time.Millisecond)
}

func TestMonitor(t *testing.T) {
	suite.Run(t, new(monitorTestSuite))
}
