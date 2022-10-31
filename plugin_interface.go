package plugin_api

import "net/http"

type ReassembledStream struct {
	StreamId      uint32
	Src           string
	SrcPort       string
	SrcName       string
	SrcNamespace  string
	Dst           string
	DstPort       string
	DstName       string
	DstNamespace  string
	Closed        bool
	BytesSent     float64
	BytesReceived float64
	Duration      float64
}

type TCPPacketPayload struct {
	StreamId uint32
	Payload  []byte
}

type StreamPlugin interface {
	InitPlugin(manager PluginManager)
	DistributeReassembledStream(stream ReassembledStream)
	DistributeTCPPacketPayload(tcpPacket TCPPacketPayload)
}

type PluginManager interface {
	RegisterPlugin(plugin StreamPlugin)
	RegisterHttpHandler(pattern string, handler func(http.ResponseWriter, *http.Request))
}
