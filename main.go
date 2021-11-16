package main

import (
	// "bytes"
	// "fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"

	// "github.com/nats-io/nats.go"

)

func filter(data []byte) {
	file, err := os.Create("./new.txt")
	if err != nil {
		log.Println(err)
	}
	defer file.Close()

	arrOfData := strings.Split(string(data), "\n")
	for _, msg := range arrOfData {
		arr := strings.Split(msg[1:len(msg)-1], "{")
		containsMerger := strings.Contains(arr[1], "merger")
		if containsMerger {
			file.WriteString(msg + "\n")
		}
	}
}

func main() {
	data, err := ioutil.ReadFile("/home/zeus/cat/something.txt")
	if err != nil {
		log.Println(err)
	}
	filter(data)

	// nc, err := nats.Connect()
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// defer nc.Close()
}
