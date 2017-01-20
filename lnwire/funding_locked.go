package lnwire

import (
	"fmt"
	"io"

	"github.com/roasbeef/btcd/btcec"
	"github.com/roasbeef/btcd/wire"
)

type FundingLocked struct {
	ChannelOutpoint *wire.OutPoint

	AnnouncementNodeSignature *btcec.Signature

	AnnouncementBitcoinSignature *btcec.Signature

	NextPerCommitmentPoint *btcec.PublicKey
}

func NewFundingLocked(op *wire.OutPoint, ans, abs *btcec.Signature,
	npcp *btcec.PublicKey) *FundingLocked {
	return &FundingLocked{
		ChannelOutpoint:              op,
		AnnouncementNodeSignature:    ans,
		AnnouncementBitcoinSignature: abs,
		NextPerCommitmentPoint:       npcp,
	}
}

// A compile time check to ensure FundingLocked implements the
// lnwire.Message interface.
var _ Message = (*FundingLocked)(nil)

// Decode deserializes the serialized FundingLocked message stored in the passed
// io.Reader into the target FundingLocked using the deserialization
// rules defined by the passed protocol version.
//
// This is part of the lnwire.Message interface.
func (c *FundingLocked) Decode(r io.Reader, pver uint32) error {
	// ChannelID (36)
	// AnnouncementNodeSignature (64)
	// AnnouncementBitcoinSignature (64)
	// NextPerCommitmentPoint (33)
	err := readElements(r,
		&c.ChannelOutpoint,
		&c.AnnouncementNodeSignature,
		&c.AnnouncementBitcoinSignature,
		&c.NextPerCommitmentPoint)
	if err != nil {
		return err
	}

	return nil
}

// Encode serializes the target FundingLocked message into the passed io.Writer
// implementation. Serialization will observe the rules defined by the passed
// protocol version.
//
// This is part of the lnwire.Message interface.
func (c *FundingLocked) Encode(w io.Writer, pver uint32) error {
	// ChannelOutpoint (36)
	// AnnouncementNodeSignature (64)
	// AnnouncementBitcoinSignature (64)
	// NextPerCommitmentPoint (33)
	err := writeElements(w,
		c.ChannelOutpoint,
		c.AnnouncementNodeSignature,
		c.AnnouncementBitcoinSignature,
		c.NextPerCommitmentPoint)
	if err != nil {
		return err
	}

	return nil
}

// Command returns the uint32 code which uniquely identifies this message as a
// FundingLocked message on the wire.
//
// This is part of the lnwire.Message interface.
func (c *FundingLocked) Command() uint32 {
	return CmdFundingLocked
}

// MaxPayloadLength returns the maximum allowed payload length for a
// FundingLocked message. This is calculated by summing the max length of all
// the fields within a FundingLocked message. The final breakdown is: 8 + 64 + 64 + 33.
//
// This is part of the lnwire.Message interface.
func (c *FundingLocked) MaxPayloadLength(uint32) uint32 {
	return 197
}

// Validate examines each populated field within the FundingLocked message for
// field sanity. For example, signature fields MUST NOT be nil.
//
// This is part of the lnwire.Message interface.
func (c *FundingLocked) Validate() error {
	if c.ChannelOutpoint == nil {
		return fmt.Errorf("the funding outpoint must be non-nil.")
	}

	if c.AnnouncementNodeSignature == nil {
		return fmt.Errorf("The announcement node signature must be non-nil.")
	}

	if c.AnnouncementBitcoinSignature == nil {
		return fmt.Errorf("The announcement bitcoin signature must be non-nil.")
	}

	if c.NextPerCommitmentPoint == nil {
		return fmt.Errorf("The next per commitment point must be non-nil.")
	}

	// We're good!
	return nil
}

// String returns the string representation of the FundingLocked message.
//
// This is part of the lnwire.Message interface.
func (c *FundingLocked) String() string {
	var npcp []byte
	if &c.NextPerCommitmentPoint != nil {
		npcp = c.NextPerCommitmentPoint.SerializeCompressed()
	}

	return fmt.Sprintf("\n--- Begin FundingLocked ---\n") +
		fmt.Sprintf("ChannelOutpoint:\t\t%x\n", c.ChannelOutpoint) +
		fmt.Sprintf("AnnouncementNodeSignature:\t%s\n", c.AnnouncementNodeSignature) +
		fmt.Sprintf("AnnouncementBitcoinSignature:\t%s\n", c.AnnouncementBitcoinSignature) +
		fmt.Sprintf("NextPerCommitmentPoint:\t\t%x\n", npcp) +
		fmt.Sprintf("--- End FundingLocked ---\n")
}
