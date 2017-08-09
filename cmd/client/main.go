package main

import (
	"bytes"
	"fmt"
	"github.com/golang/protobuf/proto"
	"github.com/itsankoff/protobuffs/def"
	"net/http"
)

func main() {
	client := def.Client{
		Id:      555,
		Name:    "John Doe",
		Email:   "johndoe@example.com",
		Country: "BG",
	}

	inbox := make([]*def.Client_Mail, 0, 20)
	inbox = append(inbox, &def.Client_Mail{RemoteEmail: "janedoe@example.com", Body: "Hello John"}, &def.Client_Mail{RemoteEmail: "nickdoe@example.com", Body: "Hi John, Greetings from Nick!"})

	client.Inbox = inbox

	data, err := proto.Marshal(&client)
	if err != nil {
		fmt.Println(err)
		return
	}

	_, err = http.Post("http://localhost:3000", "", bytes.NewBuffer(data))
	if err != nil {
		fmt.Println(err)
		return
	}
}
