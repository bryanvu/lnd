package lnwire

import (
	"bytes"
	"reflect"
	"testing"
)

func TestFundingLockedWire(t *testing.T) {
	// First create a new FundingLocked message.
	fl := NewFundingLocked(outpoint1, someSig, someSig, pubKey)

	// Next encode the FundingLocked message into an empty bytes buffer.
	var b bytes.Buffer
	if err := fl.Encode(&b, 0); err != nil {
		t.Fatalf("unable to encode FundingLocked: %v", err)
	}

	// Deserialize the encoded FundingLocked message into an empty struct.
	fl2 := &FundingLocked{}
	if err := fl2.Decode(&b, 0); err != nil {
		t.Fatalf("unable to decode FundingLocked: %v", err)
	}

	// Assert equality of the two instances
	if !reflect.DeepEqual(fl, fl2) {
		t.Fatalf("encode/decode error messages don't match %#v vs %#v", fl, fl2)
	}
}
