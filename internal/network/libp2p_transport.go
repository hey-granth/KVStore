package network

import (
	"context"
	"io"
	"log"

	libp2p "github.com/libp2p/go-libp2p"
	"github.com/libp2p/go-libp2p/core/host"
	"github.com/libp2p/go-libp2p/core/network"
	"github.com/libp2p/go-libp2p/core/peer"
)

const ProtocolID = "/kvstore/1.0.0"

type Libp2pTransport struct {
	ctx     context.Context
	host    host.Host
	handler MessageHandler
	peers   []peer.ID
}

// NewLibp2pTransport initializes a new Libp2pTransport instance. It creates a new libp2p host with default options, which includes random port assignment and default transports. The method also sets up a stream handler for the defined protocol, allowing the transport to handle incoming messages according to the specified protocol ID. If any error occurs during host creation, it returns the error for proper handling by the caller.
func NewLibp2pTransport(ctx context.Context) (*Libp2pTransport, error) {
	h, err := libp2p.New() // host with default options (random port, default transports, etc.)
	if err != nil {
		return nil, err
	}

	t := &Libp2pTransport{
		ctx:  ctx,
		host: h,
	}

	h.SetStreamHandler(ProtocolID, t.handleStream) // set the stream handler for our protocol

	return t, nil
}

// Start logs the node ID and the addresses the host is listening on. This method is called to initialize the transport layer and prepare it for communication. It does not perform any network operations but provides visibility into the node's identity and its listening addresses, which can be useful for debugging and monitoring purposes.
func (t *Libp2pTransport) Start() error {
	log.Println("Node ID:", t.host.ID().String())
	for _, addr := range t.host.Addrs() {
		log.Println("Listening on:", addr.String())
	}
	return nil
}

// Stop gracefully shuts down the transport layer by closing the libp2p host. This method ensures that all open connections are terminated and resources are cleaned up properly. It returns any error encountered during the shutdown process, allowing the caller to handle it as needed.
func (t *Libp2pTransport) Stop() error {
	return t.host.Close()
}

// RegisterHandler sets the message handler for incoming messages. This method allows the application to define custom logic for processing messages received from peers. The provided handler will be invoked with the sender's identifier and the message data whenever a message is received, enabling flexible handling of incoming communication based on the sender and content.
func (t *Libp2pTransport) RegisterHandler(handler MessageHandler) {
	t.handler = handler
}

// Send sends the given data to a specific peer identified by its peer ID. It first decodes the peer ID string into a peer.ID type, then creates a new stream to the target peer using the defined protocol ID. The method writes the data to the stream and ensures that the stream is properly closed after sending the message. If any error occurs during decoding, stream creation, or writing, it returns the error for proper handling by the caller.
func (t *Libp2pTransport) Send(peerID string, data []byte) error {
	pid, err := peer.Decode(peerID)
	if err != nil {
		return err
	}

	stream, err := t.host.NewStream(t.ctx, pid, ProtocolID)
	if err != nil {
		return err
	}
	defer func(stream network.Stream) {
		err := stream.Close()
		if err != nil {
			log.Fatalf("Failed to close stream: %v", err)
		}
	}(stream)

	_, err = stream.Write(data)
	return err
}

// Broadcast sends the given data to all connected peers. It iterates through the list of known peers and invokes the Send method for each peer, allowing the message to be disseminated across the network. This method does not return an error if sending to any individual peer fails, as it attempts to send to all peers regardless of individual failures.
func (t *Libp2pTransport) Broadcast(data []byte) error {
	for _, p := range t.peers {
		_ = t.Send(p.String(), data)
	}
	return nil
}

// handleStream is the stream handler for incoming connections. It reads the data from the stream and invokes the registered message handler with the sender's peer ID and the message data. After processing, it ensures that the stream is properly closed to free up resources.
func (t *Libp2pTransport) handleStream(s network.Stream) {
	defer func(s network.Stream) {
		err := s.Close()
		if err != nil {
			log.Fatalf("Failed to close stream: %v", err)
		}
	}(s)

	data, err := io.ReadAll(s)
	if err != nil {
		return
	}

	if t.handler != nil {
		t.handler(s.Conn().RemotePeer().String(), data)
	}
}
