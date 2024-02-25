package main

import (
	"fmt"
	"log"

	"github.com/nats-io/nats.go"
)

func main() {

	// try to connect to nats server
	natsConn, err := nats.Connect(nats.DefaultURL)

	if err != nil {
		log.Fatal(err)
	}

	// close the connection after work is done
	defer natsConn.Close()

	fmt.Println("connection successfull: ", natsConn.Statistics)

	sub, err := natsConn.Subscribe("demo", func(msg *nats.Msg) {
		fmt.Println(string(msg.Data))
	})

	if err != nil {
		fmt.Println(err)
	}

	defer sub.Unsubscribe()

	select {}

}
