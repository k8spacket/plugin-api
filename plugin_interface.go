package plugin_api

import (
	"net/http"
)

type Address struct {
	Addr      string
	Port      uint16
	Name      string
	Namespace string
}
type TCPEvent struct {
	Client  Address
	Server  Address
	TxB     uint64
	RxB     uint64
	DeltaUs uint64
}

type TLSEvent struct {
	Client         Address
	Server         Address
	TlsVersions    []uint16
	Ciphers        []uint16
	ServerName     string
	UsedTlsVersion uint16
	UsedCipher     uint16
}

type ConsumerPlugin interface {
	InitPlugin(manager PluginManager)
}

type TCPConsumerPlugin interface {
	ConsumerPlugin
	DistributeTCPEvent(tcpEvent TCPEvent)
}

type TLSConsumerPlugin interface {
	ConsumerPlugin
	DistributeTLSEvent(tlsEvent TLSEvent)
}

type PluginManager interface {
	RegisterTCPPlugin(plugin TCPConsumerPlugin)
	RegisterTLSPlugin(plugin TLSConsumerPlugin)
	RegisterHttpHandler(pattern string, handler func(http.ResponseWriter, *http.Request))
}
