package main

import (
	"bufio"
	"encoding/xml"
	"fmt"
	"os"
	"reflect"
)

type Component struct {
	XMLName  xml.Name `xml:"component"`
	Id       string   `xml:"id,attr"`
	Children []string `xml:"children"`
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	var comp = &Component{Id: "Hey boi"}
	comp.Children = []string{"hey", "boi", "what"}
	marshalledComponent, err := xml.MarshalIndent(comp, "", "     ")
	check(err)

	f, err := os.Create("structure.txt")
	check(err)

	defer f.Close()

	w := bufio.NewWriter(f)
	_, err = w.WriteString(string(marshalledComponent) + "\n")
	check(err)

	w.Flush()

	var t Component
	err = xml.Unmarshal(marshalledComponent, &t)
	check(err)

	fmt.Println(reflect.TypeOf(t.XMLName.Local))
}
