package main

import "fmt"

type User struct {
	Username   string
	FirstName  string
	SecondName string
	Age        int
}

type AttachmentType string

const (
	TypeDocument AttachmentType = "doc"
	TypeImage    AttachmentType = "img"
)

type Attachment struct {
	Type AttachmentType
	URL  string
	Data []byte
}

type Message struct {
	ID          string `ser:"id"`
	User        *User  `ser:"user"`
	Text        string `ser:"text" check:"not_nil"`
	Attachments []Attachment
}

func (m Message) Test() {
	fmt.Println("HELLO!")
}
