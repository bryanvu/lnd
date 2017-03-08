package lnwire

import (
	"io"
)

// Shutdown is sent by either side in order to initiate the cooperative closure
// of a channel. This message is sparse as both sides implicitly have the
// information necessary to construct a transaction that will send the settled
// funds of both parties to the final delivery addresses negotiated during the
// funding workflow.
//
// NOTE: The requester is able to only send a signature to initiate the
// cooperative channel closure as all transactions are assembled observing BIP
// 69, which defines a canonical ordering for inputs/outputs. Therefore, both
// sides are able to arrive at an identical closure transaction as they know
// the order of the inputs/outputs.
type Shutdown struct {
	// ChannelID serves to identify which channel is to be closed.
	ChannelID ChannelID

	Address Address
}

// Address -- TODO(bvu): add comment
type Address []byte

// NewShutdown creates a new Shutdown message.
func NewShutdown(cid ChannelID, addr Address) *Shutdown {
	return &Shutdown{
		ChannelID: cid,
		Address:   addr,
	}
}

// A compile-time check to ensure Shutdown implements the lnwire.Message
// interface.
var _ Message = (*Shutdown)(nil)

// Decode deserializes a serialized Shutdown stored in the passed io.Reader
// observing the specified protocol version.
//
// This is part of the lnwire.Message interface.
func (s *Shutdown) Decode(r io.Reader, pver uint32) error {
	return readElements(r,
		&s.ChannelID,
		&s.Address)
}

// Encode serializes the target Shutdown into the passed io.Writer observing
// the protocol version specified.
//
// This is part of the lnwire.Message interface.
func (s *Shutdown) Encode(w io.Writer, pver uint32) error {
	return writeElements(w,
		s.ChannelID,
		s.Address)
}

// MsgType returns the integer uniquely identifying this message type on the
// wire.
//
// This is part of the lnwire.Message interface.
func (s *Shutdown) MsgType() MessageType {
	return MsgShutdown
}

// MaxPayloadLength returns the maximum allowed payload size for this message
// observing the specified protocol version.
//
// This is part of the lnwire.Message interface.
func (s *Shutdown) MaxPayloadLength(pver uint32) uint32 {
	var length uint32

	// ChannelID (outpoint) - 36 bytes
	// TODO(bvu): Adjust when new ChannelID is added.
	length += 36

	// Len - 2 bytes
	length += 2

	// ScriptPubKey - 32 bytes for pay to witness script hash
	length += 32

	// NOTE: pay to pubkey hash, pay to script hash and pay to witness pubkey
	// signatures are 20 bytes in length.

	return length
}

// Validate examines each populated field within the Shutdown message for field
// sanity. For example, signature fields must not be nil.
//
// This is part of the lnwire.Message interface.
func (s *Shutdown) Validate() error {

	// We're good!
	return nil
}
