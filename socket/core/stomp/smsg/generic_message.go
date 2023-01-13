package smsg

import (
	"mem-ws/socket/core/header"
)

type GenericMessage[P interface{}] struct {
	Payload P
	Headers *header.Headers
}

func (msg *GenericMessage[P]) GetPayload() P {
	return msg.Payload
}

func (msg *GenericMessage[P]) GetMessageHeaders() *header.Headers {
	return msg.Headers
}
