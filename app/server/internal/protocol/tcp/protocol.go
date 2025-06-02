package tcp

import (
	"encoding/binary"
	"github.com/panjf2000/gnet/v2"
	"os"
)

const (
	MsgTypeHandshakeReq  = 0x01
	MsgTypeHandshakeResp = 0x02
	MsgTypeUploadReq     = 0x03
	MsgTypeUploadResp    = 0x04
	MsgTypeHeartbeat     = 0x05

	ProtocolVersion = 1
	AuthToken       = "secret-token"
)

// 连接状态
type ConnState struct {
	Authenticated bool
	FileOffset    int64
	Buffer        []byte
}

func decodeMessage(buf []byte) (msg *Message, remaining []byte, ok bool) {
	if len(buf) < 6 {
		return nil, buf, false
	}
	totalLen := int(binary.BigEndian.Uint32(buf[0:4]))
	if len(buf) < totalLen {
		return nil, buf, false
	}
	msg = &Message{
		Version: buf[4],
		Type:    buf[5],
		Payload: buf[6:totalLen],
	}
	return msg, buf[totalLen:], true
}

type Message struct {
	Version byte
	Type    byte
	Payload []byte
}

func (s *server) handleMessage(c gnet.Conn, state *ConnState, msg *Message) *gnet.Action {
	switch msg.Type {
	case MsgTypeHandshakeReq:
		return handleHandshake(c, state, msg)
	case MsgTypeUploadReq:
		return handleUpload(c, state, msg)
	case MsgTypeHeartbeat:
		err := c.AsyncWrite(buildMessage(ProtocolVersion, MsgTypeHeartbeat, nil), func(c gnet.Conn, err error) error {
			return nil
		})
		if err != nil {
			return nil
		}
	default:
		act := gnet.Close
		return &act
	}
	return nil
}

func handleHandshake(c gnet.Conn, state *ConnState, msg *Message) *gnet.Action {
	token := string(msg.Payload)
	if token == AuthToken {
		state.Authenticated = true
		err := c.AsyncWrite(buildMessage(ProtocolVersion, MsgTypeHandshakeResp, []byte{1}), func(c gnet.Conn, err error) error {
			return nil
		})
		if err != nil {
			return nil
		}
	} else {
		err := c.AsyncWrite(buildMessage(ProtocolVersion, MsgTypeHandshakeResp, []byte{0}), func(c gnet.Conn, err error) error {
			return nil
		})
		if err != nil {
			return nil
		}
		act := gnet.Close
		return &act
	}
	return nil
}

func handleUpload(c gnet.Conn, state *ConnState, msg *Message) *gnet.Action {
	if !state.Authenticated || len(msg.Payload) < 8 {
		act := gnet.Close
		return &act
	}
	offset := int64(binary.BigEndian.Uint64(msg.Payload[:8]))
	fileData := msg.Payload[8:]
	f, err := os.OpenFile("uploaded_file.data", os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		act := gnet.Close
		return &act
	}
	defer f.Close()
	if _, err := f.Seek(offset, 0); err != nil {
		act := gnet.Close
		return &act
	}
	n, err := f.Write(fileData)
	if err != nil {
		act := gnet.Close
		return &act
	}
	state.FileOffset = offset + int64(n)
	offsetBuf := make([]byte, 8)
	binary.BigEndian.PutUint64(offsetBuf, uint64(state.FileOffset))
	err = c.AsyncWrite(
		buildMessage(ProtocolVersion, MsgTypeUploadResp, offsetBuf),
		func(c gnet.Conn, err error) error {
			return nil
		},
	)
	if err != nil {
		return nil
	}
	return nil
}

func buildMessage(version byte, msgType byte, payload []byte) []byte {
	payloadLen := 2 // version + type
	if payload != nil {
		payloadLen += len(payload)
	}

	buf := make([]byte, 4+payloadLen)
	binary.BigEndian.PutUint32(buf[:4], uint32(payloadLen))
	buf[4] = version
	buf[5] = msgType
	if payload != nil {
		copy(buf[6:], payload)
	}
	return buf
}
