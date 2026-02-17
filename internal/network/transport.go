package network

// MessageHandler is a function type that defines the signature for handling incoming messages. It takes the sender's identifier and the message data as a byte slice. This allows for flexible handling of messages, as different implementations can define their own logic for processing incoming data based on the sender and content.
type MessageHandler func(from string, data []byte)

type Transport interface {
	Start() error // For libp2p, this will create host, start listening, set stream handlers
	Stop() error  // this method will be responsible for gracefully shutting down the transport layer, closing any open connections, and cleaning up resources.

	Send(peerID string, data []byte) error  // send msg to a specific peer
	Broadcast(data []byte) error            // send msg to all the connected peers
	RegisterHandler(handler MessageHandler) // register a handler for incoming messages, allowing the transport layer to invoke the provided handler whenever a message is received from any peer. This enables the application to define custom logic for processing incoming messages based on the sender and content.
}
