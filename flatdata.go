// Package flatdata use flatbuffer to serialize some data
package flatdata

import (
	flatbuffers "github.com/google/flatbuffers/go"
	"github.com/thanhnl/flatdata/model"
)

type Phone struct {
	PhoneType string
	Number    string
}

type Contact struct {
	Id          string
	FirstName   string
	LastName    string
	Description string
	Phones      []Phone
}

type Message struct {
	Id        string
	Contacts  []Contact
	Receivers []string
}

// flatPhone return Phone offsets slice.
func flatPhone(builder *flatbuffers.Builder, phones []Phone) []flatbuffers.UOffsetT {
	var phoneOffsets = []flatbuffers.UOffsetT{}
	for _, phone := range phones {
		numberPos := builder.CreateString(phone.Number)
		typePos := builder.CreateString(phone.PhoneType)
		model.PhoneStart(builder)
		model.PhoneAddNumber(builder, numberPos)
		model.PhoneAddPhoneType(builder, typePos)
		phoneOffset := model.PhoneEnd(builder)
		phoneOffsets = append(phoneOffsets, phoneOffset)
	}
	return phoneOffsets
}

// flatContact return Contact offsets slice.
func flatContact(builder *flatbuffers.Builder, contacts []Contact) []flatbuffers.UOffsetT {
	var contactOffsets = []flatbuffers.UOffsetT{}
	for _, contact := range contacts {
		// flat phone vector
		phoneOffsets := flatPhone(builder, contact.Phones)
		model.ContactStartPhonesVector(builder, len(phoneOffsets))
		for i := len(phoneOffsets) - 1; i >= 0; i-- {
			builder.PrependUOffsetT(phoneOffsets[i])
		}
		phoneVectorOffset := builder.EndVector(len(phoneOffsets))
		idPos := builder.CreateString(contact.Id)
		descPos := builder.CreateString(contact.Description)
		firstnamePos := builder.CreateString(contact.FirstName)
		lastnamePos := builder.CreateString(contact.LastName)
		model.ContactStart(builder)
		model.ContactAddId(builder, idPos)
		model.ContactAddDescription(builder, descPos)
		model.ContactAddFirstName(builder, firstnamePos)
		model.ContactAddLastName(builder, lastnamePos)
		model.ContactAddPhones(builder, phoneVectorOffset)
		contactOffset := model.ContactEnd(builder)
		contactOffsets = append(contactOffsets, contactOffset)
	}
	return contactOffsets
}

// FlatMsg return flatbuffer byte slice
func FlatMsg(msg Message) []byte {
	builder := flatbuffers.NewBuilder(0)

	// flat contact vector
	contactOffsets := flatContact(builder, msg.Contacts)
	model.MessageStartContactsVector(builder, len(contactOffsets))
	for i := len(contactOffsets) - 1; i >= 0; i-- {
		builder.PrependUOffsetT(contactOffsets[i])
	}
	contactVectorOffset := builder.EndVector(len(contactOffsets))

	// flat receiver vector
	var rcvOffsets = []flatbuffers.UOffsetT{}
	for _, rcv := range msg.Receivers {
		rcvOffsets = append(rcvOffsets, builder.CreateString(rcv))
	}
	model.MessageStartReceiversVector(builder, len(msg.Receivers))
	for i := len(rcvOffsets) - 1; i >= 0; i-- {
		builder.PrependUOffsetT(rcvOffsets[i])
	}
	rcvVectorOffset := builder.EndVector(len(rcvOffsets))

	// flat message
	idPos := builder.CreateString(msg.Id)
	model.MessageStart(builder)
	model.MessageAddId(builder, idPos)
	model.MessageAddContacts(builder, contactVectorOffset)
	model.MessageAddReceivers(builder, rcvVectorOffset)
	msgOffset := model.MessageEnd(builder)
	builder.Finish(msgOffset)
	return builder.FinishedBytes()
}

// unFlatPhone return Phone slice
func unFlatPhone(contactBuf *model.Contact) []Phone {
	var phones = []Phone{}
	var phone = Phone{}
	phoneBuf := new(model.Phone)
	for i := 0; i < contactBuf.PhonesLength(); i++ {
		if contactBuf.Phones(phoneBuf, i) {
			phone.Number = string(phoneBuf.Number())
			phone.PhoneType = string(phoneBuf.PhoneType())
			phones = append(phones, phone)
		}
	}
	return phones
}

// unFlatContact return Contact slice
func unFlatContact(msgBuf *model.Message) []Contact {
	var contacts = []Contact{}
	var contact = Contact{}
	contactBuf := new(model.Contact)
	for i := 0; i < msgBuf.ContactsLength(); i++ {
		if msgBuf.Contacts(contactBuf, i) {
			contact.Id = string(contactBuf.Id())
			contact.Description = string(contactBuf.Description())
			contact.FirstName = string(contactBuf.FirstName())
			contact.LastName = string(contactBuf.LastName())
			if contactBuf.PhonesLength() > 0 {
				contact.Phones = append(contact.Phones, unFlatPhone(contactBuf)...)
			}
			contacts = append(contacts, contact)
		}
	}
	return contacts
}

// UnFlatMsg return Message
func UnFlatMsg(buf []byte) Message {
	var msg = Message{}
	msgBuf := model.GetRootAsMessage(buf, 0)
	msg.Id = string(msgBuf.Id())
	if msgBuf.ContactsLength() > 0 {
		msg.Contacts = append(msg.Contacts, unFlatContact(msgBuf)...)
	}
	if msgBuf.ReceiversLength() > 0 {
		for i := 0; i < msgBuf.ReceiversLength(); i++ {
			msg.Receivers = append(msg.Receivers, string(msgBuf.Receivers(i)))
		}
	}
	return msg
}
