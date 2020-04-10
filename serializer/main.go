package main

import (
	"fmt"
	"reflect"
)

func main() {
	port := 8080
	Iterator(Conf{
		IP:   "127.0.0.1",
		Host: "localhost",
		Port: &port,
		Aliases: []string{
			"A1",
			"B2",
		},
	})
}

type Conf struct {
	IP      string `json:"ip" yaml:"dfds"`
	Host    string
	Port    *int
	Aliases []string
}

func (conf Conf) GetIP() string {
	return conf.IP
}

func Iterator(s interface{}) string {
	sT := reflect.TypeOf(s)
	sV := reflect.ValueOf(s)

	for i := 0; i < sT.NumField(); i++ {
		fieldT := sT.Field(i)
		fieldV := sV.Field(i)

		//fmt.Println(fieldT.Name, fieldT.Type, fieldV.String(), fieldT.Tag.Get("json"))
		switch fieldT.Type.Kind() {
		case reflect.Int:
			fmt.Println(fieldV.Int())
		case reflect.Ptr:
			fmt.Println(fieldV.Elem().Int())
		case reflect.Slice:
			switch fieldT.Type.Elem().Kind() {
			case reflect.String:
				for j := 0; j < fieldV.Len(); j++ {
					fmt.Println(j, fieldV.Index(j))
				}
			}
		default:
			fmt.Println(fieldV.String())
		}
	}

	getIP := sV.MethodByName("GetIP")
	results := getIP.Call([]reflect.Value{})
	fmt.Println(results)

	return ""
}
