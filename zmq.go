package zmq

import (
	"github.com/loadimpact/k6/js/modules"
	zmq "github.com/pebbe/zmq4"
)

func init() {
	modules.Register("k6/x/zmq", new(ZMQ))
}

// ZMQ is the k6 ZeroMQ extension.
type ZMQ struct{}

// NewSocket creates a new ZeroMQ socket
func (*ZMQ) NewSocket(addr string) *zmq.Socket {
	requester, err := zmq.NewSocket(zmq.REQ)
	if err != nil {
		ReportError(err, "Failed to create a new socket")
	}
	requester.Connect(addr)
	return requester
}

// Send sends a message part on a socket.
func (*ZMQ) Send(requester *zmq.Socket, data string) {
	requester.Send(data, 0)
}

// CloseSocket closes the socket. If not called explicitly, the socket will be closed on garbage collection
func (*ZMQ) CloseSocket(requester *zmq.Socket) {
	requester.Close()
}
