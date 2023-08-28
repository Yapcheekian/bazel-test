package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net"

	"github.com/Yapcheekian/bazel-test/client-server-go/model"
	"github.com/google/uuid"
)

func main() {
	id := uuid.New()
	fmt.Println(id)
	log.Println("Spinning up the Echo Server in Go...")

	listen, err := net.Listen("tcp", ":1234")
	if err != nil {
		log.Panicln("Unable to listen: " + err.Error())
	}

	defer listen.Close()
	connection, err := listen.Accept()
	if err != nil {
		log.Panicln("Cannot accept a connection! Error: " + err.Error())
	}

	log.Println("Receiving on a new connection")
	defer connection.Close()
	defer log.Println("Connection now closed.")

	buffer := make([]byte, 2048)
	size, err := connection.Read(buffer)
	if err != nil {
		log.Println("Cannot read from the buffer! Error: " + err.Error())
	}

	data := buffer[:size]
	var input model.Input
	if err := json.Unmarshal(data, &input); err != nil {
		log.Panicln("Unable to unmarshal the buffer! Error: " + err.Error())
	}

	log.Println("Message = " + input.Message)
	log.Println("Value = " + fmt.Sprintf("%f", input.Value))
	input.Message = "Echoed from Go: " + input.Message
	input.Value = 2 * input.Value
	message, err := json.Marshal(input)
	if err != nil {
		log.Panicln(
			"Unable to marshall the object! Error: " + err.Error())
	}

	connection.Write(message)
}
