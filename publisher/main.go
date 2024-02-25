package main

import (
	"fmt"
	"log"
	"time"

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

	for i := range 15 {

		err = natsConn.Publish("demo", []byte("Welcome to the jungle..."))

		if err != nil {
			fmt.Println("error while publishing msg: ", err)
		}

		fmt.Println("published msg..", i)
		time.Sleep(1 * time.Second)

	}
}
