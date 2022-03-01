package main

import (
	"fmt"
	"io/ioutil"

	"github.com/Yapcheekian/protobuf-practice/src/simple"
	"github.com/gogo/protobuf/jsonpb"
	"google.golang.org/protobuf/proto"
)

func main() {
	sm := simple.SimpleMessage{
		Id:         1234,
		IsSimple:   true,
		Name:       "yap",
		SampleList: []int32{1},
	}

	marshaller := jsonpb.Marshaler{}
	text, err := marshaller.MarshalToString(&sm)
	if err != nil {
		panic(err)
	}
	fmt.Println(text)

	var newMsg simple.SimpleMessage
	if err := jsonpb.UnmarshalString(text, &newMsg); err != nil {
		panic(err)
	}

	fmt.Println(newMsg.Name)

	out, err := proto.Marshal(&sm)
	if err != nil {
		panic(err)
	}

	ioutil.WriteFile("message", out, 0666)

	in, err := ioutil.ReadFile("message")
	if err != nil {
		panic(err)
	}

	var new simple.SimpleMessage
	if err := proto.Unmarshal(in, &new); err != nil {
		panic(err)
	}

	fmt.Println(new.Name)
}
