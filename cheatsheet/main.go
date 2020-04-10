package main

import (
	"fmt"
	"reflect"
)

func main() {

	msg := Message{
		ID:          "abc",
		User:        &User{
			Username:   "vapupkin",
			FirstName:  "Vasya",
			SecondName: "Pupkin",
			Age:        25,
		},
		Text:        "test text",
		Attachments: []Attachment{
			{
				Type: TypeDocument,
				URL:  "https://site.com/doc",
				Data: nil,
			},
		},
	}

	msgT := reflect.TypeOf(msg)
	msgV := reflect.ValueOf(msg)

	msgV.Method(0).Call([]reflect.Value{})

	for fieldNum := 0; fieldNum < msgT.NumField(); fieldNum++ {
		fieldT := msgT.Field(fieldNum)
		fieldV := msgV.Field(fieldNum)

		fmt.Println(fieldT.Name, fieldT.Type, fieldT.Tag.Get("ser"))
		if fieldT.Type.Kind() == reflect.Slice {
			attachT := fieldT.Type.Elem()

			for i := 0; i < fieldV.Len(); i++ {
				attachV := fieldV.Index(i)

				for fNum := 0; fNum < attachT.NumField(); fNum++ {
					fmt.Println("attach:", attachV.Field(fNum))
				}

			}

		}
	}

}