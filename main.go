package main

import (
	"fmt"
	"net"
	"strings"

	"github.com/naman_1402/redis-clone/resp"
)

func main() {

	l, err := net.Listen("tcp", ":6379") // start server, returns listener and error
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Listening on port :6379")

	conn, err := l.Accept()
	if err != nil {
		fmt.Println(err)
		return
	}
	defer conn.Close()
	// close connection once finished

	for {
		// create new deserializer for each request
		deserializer := resp.NewDeserializer(conn)
		// Read RESP formatted input
		value, err := deserializer.Read()
		if err != nil {
			fmt.Println(err)
			return
		}
		// Validation
		if value.Type != "array" || len(value.Array) == 0 {
			fmt.Println("invalid request")
			continue
		}
		// Extraction
		command := strings.ToUpper(value.Array[0].Bulk)
		args := value.Array[1:]

		// Set up writer, and execute handler
		writer := resp.NewWriter(conn)
		// Handler function
		handler, ok := resp.Handlers[command]
		if !ok {
			fmt.Println("invalid command: ", command)
			writer.Write(resp.Value{Type: "string", Str: ""})
			continue
		}

		result := handler(args)
		writer.Write(result)
	}
}
