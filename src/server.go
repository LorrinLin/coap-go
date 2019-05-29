package main

import (
	"net"
	coap "github.com/lorrin/go-coap"
)

func main(){
	coap.ListenAndServe("udp",":5683",
		coap.FuncHandler(
			func(l *net.UDPConn, a *net.UDPAddr, m *coap.Message) *coap.Message{
				
				res := &coap.Message{
				Type:      coap.Acknowledgement,
				Code:      coap.Content,
				MessageID: m.MessageID,
				Token:     m.Token,
				Payload:   []byte("hello to you!"),
			}
				res.SetOption(coap.ContentFormat, coap.TextPlain)
				return res
				
			}))
}
