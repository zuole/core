package p2p

import (
	"context"
	"errors"
	"fmt"
	"net"
	"os"
	"runtime"

	"github.com/DOSNetwork/core/log"
	"github.com/DOSNetwork/core/p2p/network"
	"github.com/DOSNetwork/core/suites"

	"github.com/golang/protobuf/proto"
	"github.com/golang/protobuf/ptypes"
)

var logger log.Logger

const (
	NONE = iota // 0
	SWIM
)

func bToMb(b uint64) uint64 {
	return b / 1024 / 1024
}
func PrintMemUsage() {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	// For info on each, see: https://golang.org/pkg/runtime/#MemStats
	fmt.Printf("Alloc = %v MiB", bToMb(m.Alloc))
	fmt.Printf("\tTotalAlloc = %v MiB", bToMb(m.TotalAlloc))
	fmt.Printf("\tSys = %v MiB", bToMb(m.Sys))
	fmt.Printf("\tNumGC = %v\n", m.NumGC)
}

func CreateP2PNetwork(id []byte, port string, netType int) (P2PInterface, error) {
	suite := suites.MustFind("bn256")
	logger = log.New("module", "p2p")
	p := &Server{
		suite:     suite,
		messages:  make(chan P2PMessage, 100),
		subscribe: make(chan Subscription),
		unscribe:  make(chan Subscription),
		port:      port,
	}
	p.ctx, p.cancel = context.WithCancel(context.Background())
	p.secKey = suite.Scalar().Pick(suite.RandomStream())
	p.pubKey = suite.Point().Mul(p.secKey, nil)
	p.id = id
	ip, err := GetPublicIP()
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	p.addr = ip

	switch netType {
	case SWIM:
		p.network, err = network.NewSerfNet(ip, id)
	default:

	}
	return p, nil
}

func GetPublicIP() (ip net.IP, err error) {
	ipString := os.Getenv("PUBLICIP")
	if ipString != "" {
		ip = net.ParseIP(ipString)
	} else {
		err = errors.New("No Public IP")
	}
	return
	/*
		//FOR DOCKER AWS TESTING

		response, err := http.Get("http://ipconfig.me")
		if err != nil {
			return
		}

		ipBytes, err := ioutil.ReadAll(response.Body)
		if err != nil {
			return
		}
		ipString := string(ipBytes)
		ip = net.ParseIP(ipString)
	*/
	//fmt.Println(ip)
	//////////////////////////////
	/*
		var addrs []net.Addr

		if addrs, err = net.InterfaceAddrs(); err != nil {
			return nil, err
		}

		for _, a := range addrs {
			if ipnet, ok := a.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
				if ipnet.IP.To4() != nil {
					return ipnet.IP, nil
				}
			}
		}*/
	return
}

type P2PMessage struct {
	Msg          ptypes.DynamicAny
	Sender       []byte
	RequestNonce uint64
}

type P2PInterface interface {
	GetIP() net.IP
	GetID() []byte
	SetPort(port string)
	Listen() error
	Join(bootstrapIp []string) (num int, err error)
	ConnectTo(ip string, id []byte) ([]byte, error)
	DisConnectTo(id []byte) error
	Leave()
	Request(id []byte, m proto.Message) (msg P2PMessage, err error)
	Reply(id []byte, nonce uint64, m proto.Message) (err error)
	SubscribeEvent(chanBuffer int, messages ...interface{}) (outch chan P2PMessage, err error)
	UnSubscribeEvent(messages ...interface{})
	Members() int
	ConnectToAll() (memNum, connNum int)
}
