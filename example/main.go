package main

import (
	"fmt"

	"github.com/thanhnl/flatdata"
)

var (
	Phone1   = flatdata.Phone{PhoneType: "Work", Number: "123"}
	Phone2a  = flatdata.Phone{PhoneType: "Home", Number: "456"}
	Phone2b  = flatdata.Phone{PhoneType: "Work", Number: "789"}
	Contact1 = flatdata.Contact{Id: "1", FirstName: "Foo", LastName: "Bar",
		Description: "Unknown", Phones: []flatdata.Phone{Phone1}}
	Contact2 = flatdata.Contact{Id: "2", FirstName: "John", LastName: "Smith",
		Description: "Friend", Phones: []flatdata.Phone{Phone2a, Phone2b}}
	Message1 = flatdata.Message{Id: "1", Contacts: []flatdata.Contact{Contact1},
		Receivers: []string{"sms"}}
	Message2 = flatdata.Message{Id: "2", Contacts: []flatdata.Contact{Contact1, Contact2},
		Receivers: []string{"sms", "im"}}
	Message3 = flatdata.Message{Id: "3", Contacts: []flatdata.Contact{Contact1}}
	Messages = []flatdata.Message{Message1, Message2, Message3}
)

func main() {
	var buf []byte
	var msg flatdata.Message
	for i := 0; i < len(Messages); i++ {
		buf = flatdata.FlatMsg(Messages[i])
		fmt.Println("Msg   :", Messages[i])
		msg = flatdata.UnFlatMsg(buf)
		fmt.Printf("Unflat: %v \n\n", msg)
	}
}
