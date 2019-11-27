package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/IronCore864/protobuf-example/pb"
	"github.com/golang/protobuf/proto"
)

func read(fname string) {
	in, err := ioutil.ReadFile(fname)
	if err != nil {
		log.Fatalln("Error reading file:", err)
	}
	p := &pb.Person{}
	if err := proto.Unmarshal(in, p); err != nil {
		log.Fatalln("Failed to parse address book:", err)
	}
	fmt.Println(p.Email)
}

func write(fname string) {
	p := pb.Person{
		Id:    1234,
		Name:  "John Doe",
		Email: "jdoe@example.com",
		Phones: []*pb.Person_PhoneNumber{
			{Number: "555-4321", Type: pb.Person_HOME},
		},
	}

	out, err := proto.Marshal(&p)

	if err != nil {
		log.Fatalln("Failed to encode address book:", err)
	}

	if err := ioutil.WriteFile(fname, out, 0644); err != nil {
		log.Fatalln("Failed to write address book:", err)
	}
}

func printUsageAndQuit() {
	fmt.Println("Usage:")
	fmt.Println("	./protobuf-example write")
	fmt.Println("	./protobuf-example read")
	os.Exit(1)
}

func main() {
	argsWithoutProg := os.Args[1:]
	if len(argsWithoutProg) == 0 {
		printUsageAndQuit()
	}

	action := argsWithoutProg[0]

	if action != "read" && action != "write" {
		printUsageAndQuit()
	}

	fname := "protobuf.out"

	switch action {
	case "read":
		read(fname)
	case "write":
		write(fname)
	}
}
