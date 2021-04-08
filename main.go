package main

import (
	"fmt"
	"log"
	"os"

	// PROTO PACKAGES
	"./book"
	"./enums"
	"./person"

	// EXTERNAL PACKAGES
	"github.com/golang/protobuf/jsonpb"
	"google.golang.org/protobuf/proto"
)

func main() {

	// personReadsBook()
	// readWritePerson()
	// jsonToFrom()

	doEnum()

}

func doEnum() {
	es := &enums.Student{}

	fmt.Println(es)
}

func jsonToFrom() {
	s := toJson()
	fmt.Println("pb as a string:", s)
	b := &book.Book{}
	fmt.Println("before unmarshal of json string", b)
	fromJson(s, b)
	fmt.Println("after unmarshal of json string", b)
}

func fromJson(s string, b *book.Book) {

	err := jsonpb.UnmarshalString(s, b)

	if err != nil {
		log.Fatalln(err)
	}
}

func toJson() string {
	b := &book.Book{
		Title: "The Sevi Islands",
		Pages: 93,
		Author: &book.Author{
			Name: "Ronald Magento",
			Age:  34,
		},
	}

	m := jsonpb.Marshaler{}
	s, err := m.MarshalToString(b)

	if err != nil {
		log.Fatalln(err)
	}
	return s
}

func personReadsBook() {
	b := &book.Book{
		Title: "The Wonder Years",
		Pages: 194,
		Author: &book.Author{
			Name: "Adam Vincent",
			Age:  30,
		},
	}

	fmt.Println(b.GetAuthor())

	bs, err := proto.Marshal(b)

	if err != nil {
		log.Fatal("Oh My..", err)
	}

	fmt.Println(bs)

	p := &person.Person{
		Fname: "Nikolai",
		Lname: "Ifrim",
		Age:   32,
	}
	readBook(p, b)
}

func readWritePerson() {
	p2 := makePerson()

	fmt.Println(p2)

	writeToFile("person.bin", p2)

	p3 := &person.Person{}

	readFromFile("person.bin", p3)

	fmt.Println(p3)
}

func readBook(reader *person.Person, book *book.Book) {
	fmt.Println(reader.GetFname(), "Is reading", book.GetTitle())
}

func makePerson() *person.Person {
	n := &person.Person{
		Fname: "Adam",
		Lname: "Vincent",
		Age:   30,
	}

	return n
}

func writeToFile(filename string, p proto.Message) error {

	bs, err := proto.Marshal(p)

	fmt.Println(bs)

	if err != nil {
		log.Fatal(err)
		return err
	}
	error := os.WriteFile(filename, bs, 0644)
	return error
}

func readFromFile(filename string, pb proto.Message) error {
	bs, err := os.ReadFile(filename)
	if err != nil {
		log.Fatalln("Uh oh!!", err)
		return err
	}

	error := proto.Unmarshal(bs, pb)

	if error != nil {
		log.Fatalln(error)
		return error
	}

	return nil

}
