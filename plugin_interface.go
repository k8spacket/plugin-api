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
	TlsVersions    [8]uint16
	Ciphers        [100]uint16
	ServerName     string
	UsedTlsVersion uint16
	UsedCipher     uint16
}

type TCPConsumerPlugin interface {
	InitPlugin(manager PluginManager)
	DistributeTCPEvent(tcpEvent TCPEvent)
}

type TLSConsumerPlugin interface {
	InitPlugin(manager PluginManager)
	DistributeTLSEvent(tlsEvent TLSEvent)
}

type PluginManager interface {
	RegisterTCPPlugin(plugin TCPConsumerPlugin)
	RegisterTLSPlugin(plugin TLSConsumerPlugin)
	RegisterHttpHandler(pattern string, handler func(http.ResponseWriter, *http.Request))
}
