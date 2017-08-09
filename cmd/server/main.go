package main

import (
	"fmt"
	"github.com/golang/protobuf/proto"
	"github.com/itsankoff/protobuf/def"
	"io/ioutil"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		client := def.Client{}
		data, err := ioutil.ReadAll(r.Body)
		if err != nil {
			fmt.Println(err)
		}

		if err := proto.Unmarshal(data, &client); err != nil {
			fmt.Println(err)
		}

		fmt.Println("Client:")
		fmt.Printf("\tId: %d\n\tName: %s\n\tEmail:%s\n\tCountry: %s\n",
			client.Id, client.Name, client.Email, client.Country)

		fmt.Println("\tInbox:")
		for _, mail := range client.Inbox {
			fmt.Println("\t\t", mail.RemoteEmail, ": ", mail.Body)
		}
	})

	fmt.Println("Server started")
	http.ListenAndServe(":3000", nil)
}
