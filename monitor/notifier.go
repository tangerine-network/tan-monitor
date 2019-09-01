package monitor

import (
	"crypto/tls"
	"fmt"
	"log"
	"net"
	"net/mail"
	"net/smtp"
	"strings"

	"github.com/ethereum/go-ethereum/common"
)

const (
	// FINED notes the node is fined
	FINED uint = iota
	// INSUFFICIENT_TAN notes the node key with balance < 100 TAN
	INSUFFICIENT_TAN
	// INSUFFICIENT_ETH notes the node key with balance < threshold
	INSUFFICIENT_ETH
)

// Notifier is the interface for notifier interfaces.
type Notifier interface {
	notify(node, string, uint, string)
}

// Email is the email notifier object.
type Email struct {
	sender     string
	password   string
	smtpServer string
	ccList     string
}

// NewEmail is the Email constructor.
func NewEmail(sender, password, server, ccList string) Notifier {
	e := Email{
		sender:     sender,
		password:   password,
		smtpServer: server,
		ccList:     ccList,
	}
	return &e
}

func (e *Email) notify(n node, network string, notifyType uint, threshold string) {
	var subj string
	var body string
	switch notifyType {
	case FINED:
		subj = "[Notification] Your Tangerine " + network + " Network Full Node is Fined"
		body = e.finedBody(n.name)
		go e.send(n.email, network, subj, body)
	case INSUFFICIENT_TAN:
		subj = "[Notification] Tangerine " + network + " Insufficient Balance in Node Key"
		body = e.insufficientBalanceBody(n.name, n.nodeKeyAddress)
		go e.send(n.email, network, subj, body)
	case INSUFFICIENT_ETH:
		subj = "[Notification] Tangerine " + network + " Insufficient ETH in Node Key"
		body = e.insufficientETHBalanceBody(n.name,
			n.nodeKeyAddress, threshold, network)
		go e.send(n.email, network, subj, body)
	}
}

func (e *Email) finedBody(name string) string {
	body := `
Dear Node Operator:

	Your full node "` + name + `" is fined.
	Please pay the fine, and check your node status.
	Also, please check https://tangerine-network.github.io/wiki/#/Rules-for-the-node-set?id=penalty for the penalty condition.

	Thanks for your cooperation.

-- Tangerine Network
`
	return body
}

func (e *Email) insufficientBalanceBody(name string, address common.Address) string {
	body := `
Dear Node Operator:

	There is no sufficient balacne in your full node "` + name + `" key address "` + address.Hex() + `".
	Please transfer some tokens to the node key address and keep the balance at least 200 TAN.

	Thanks for your cooperation.

-- Tangerine Network
`
	return body
}

func (e *Email) insufficientETHBalanceBody(
	name string,
	address common.Address,
	threshold string,
	network string) string {
	ethNet := "Rinkeby"
	if network == "Mainnet" {
		ethNet = "Mainnet"
	}
	body := `
Dear Node Operator:

	There is no sufficient ETH in your full node "` + name + `" key address "` + address.Hex() + `".
	Please transfer some ETH (` + ethNet + `) to the node key address and keep the balance at least ` + threshold + ` ETH.

	Thanks for your cooperation.

-- Tangerine Network
`
	return body
}

func (e *Email) send(toAddress, network, subj, body string) {
	from := mail.Address{Name: "Tangerine Notifier", Address: e.sender}
	to := mail.Address{Address: toAddress}
	// Setup headers
	headers := make(map[string]string)
	headers["From"] = from.String()
	headers["To"] = to.String()
	headers["Subject"] = subj
	if e.ccList != "" {
		headers["Cc"] = strings.TrimSuffix(e.ccList, "\n")
	}

	// Setup message
	message := ""
	for k, v := range headers {
		message += fmt.Sprintf("%s: %s\r\n", k, v)
	}
	message += "\r\n" + body

	// Connect to the SMTP Server
	servername := e.smtpServer + ":465"

	host, _, _ := net.SplitHostPort(servername)

	auth := smtp.PlainAuth("", e.sender, e.password, host)

	// TLS config
	tlsconfig := &tls.Config{
		InsecureSkipVerify: false,
		ServerName:         host,
	}

	// Here is the key, you need to call tls.Dial instead of smtp.Dial
	// for smtp servers running on 465 that require an ssl connection
	// from the very beginning (no starttls)
	conn, err := tls.Dial("tcp", servername, tlsconfig)
	if err != nil {
		log.Panic(err)
	}

	c, err := smtp.NewClient(conn, host)
	if err != nil {
		log.Panic(err)
	}

	// Auth
	if err = c.Auth(auth); err != nil {
		log.Panic(err)
	}

	// To && From
	if err = c.Mail(from.Address); err != nil {
		log.Panic(err)
	}

	if err = c.Rcpt(to.Address); err != nil {
		log.Panic(err)
	}

	if e.ccList != "" {
		cc := strings.Split(e.ccList, ",")
		for _, address := range cc {
			address = strings.TrimSuffix(address, "\n")
			if err = c.Rcpt(address); err != nil {
				log.Println("add cc fail", address, err)
			}
		}
	}

	// Data
	w, err := c.Data()
	if err != nil {
		log.Panic(err)
	}

	_, err = w.Write([]byte(message))
	if err != nil {
		log.Panic(err)
	}

	err = w.Close()
	if err != nil {
		log.Panic(err)
	}

	c.Quit()
}
