// Copyright (c) 2013-2015 The btcsuite developers
// Use of this source code is governed by an ISC
// license that can be found in the LICENSE file.

package wire

import (
	"fmt"
	"io"

	"github.com/btcboost/copernicus/util"
)

// MsgPong implements the Message interface and represents a bitcoin pong
// message which is used primarily to confirm that a connection is still valid
// in response to a bitcoin ping message (MsgPing).
//
// This message was not added until protocol versions AFTER BIP0031Version.
type MsgPong struct {
	// Unique value associated with message that is used to identify
	// specific ping message.
	Nonce uint64
}

// Decode decodes r using the bitcoin protocol encoding into the receiver.
// This is part of the Message interface implementation.
func (msg *MsgPong) Decode(r io.Reader, pver uint32, enc MessageEncoding) error {
	// NOTE: <= is not a mistake here.  The BIP0031 was defined as AFTER
	// the version unlike most others.
	if pver <= BIP0031Version {
		str := fmt.Sprintf("pong message invalid for protocol "+
			"version %d", pver)
		return messageError("MsgPong.Decode", str)
	}

	return util.ReadElements(r, &msg.Nonce)
}

// Encode encodes the receiver to w using the bitcoin protocol encoding.
// This is part of the Message interface implementation.
func (msg *MsgPong) Encode(w io.Writer, pver uint32, enc MessageEncoding) error {
	// NOTE: <= is not a mistake here.  The BIP0031 was defined as AFTER
	// the version unlike most others.
	if pver <= BIP0031Version {
		str := fmt.Sprintf("pong message invalid for protocol "+
			"version %d", pver)
		return messageError("MsgPong.Encode", str)
	}

	return util.WriteElements(w, msg.Nonce)
}

// Command returns the protocol command string for the message.  This is part
// of the Message interface implementation.
func (msg *MsgPong) Command() string {
	return CmdPong
}

// MaxPayloadLength returns the maximum length the payload can be for the
// receiver.  This is part of the Message interface implementation.
func (msg *MsgPong) MaxPayloadLength(pver uint32) uint32 {
	plen := uint32(0)
	// The pong message did not exist for BIP0031Version and earlier.
	// NOTE: > is not a mistake here.  The BIP0031 was defined as AFTER
	// the version unlike most others.
	if pver > BIP0031Version {
		// Nonce 8 bytes.
		plen += 8
	}

	return plen
}

// NewMsgPong returns a new bitcoin pong message that conforms to the Message
// interface.  See MsgPong for details.
func NewMsgPong(nonce uint64) *MsgPong {
	return &MsgPong{
		Nonce: nonce,
	}
}
