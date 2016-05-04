// automatically generated, do not modify

package model

import (
	flatbuffers "github.com/google/flatbuffers/go"
)
type Phone struct {
	_tab flatbuffers.Table
}

func (rcv *Phone) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *Phone) PhoneType() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

func (rcv *Phone) Number() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

func PhoneStart(builder *flatbuffers.Builder) { builder.StartObject(2) }
func PhoneAddPhoneType(builder *flatbuffers.Builder, phoneType flatbuffers.UOffsetT) { builder.PrependUOffsetTSlot(0, flatbuffers.UOffsetT(phoneType), 0) }
func PhoneAddNumber(builder *flatbuffers.Builder, number flatbuffers.UOffsetT) { builder.PrependUOffsetTSlot(1, flatbuffers.UOffsetT(number), 0) }
func PhoneEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT { return builder.EndObject() }
