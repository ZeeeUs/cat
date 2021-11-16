package main

import (
	// "bytes"
	// "fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"

	// "github.com/nats-io/nats.go"
	""
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

	// str := "{2021-10-30 14:03:05.398714165 +0000 UTC развивайка математика 2  { _t0=121&_t1=3969&_t2=17072merger }}"
	// str := "{2021-10-30 14:03:05.398714165 +0000 UTC развивайка математика 2  { Hellomerger }}"
	// arr := strings.Split(str[1:len(str)-1], "{")
	// forSub := arr[1][1:len(arr[1])-1]
	// // fmt.Println(forSub)
	// containsMerger := strings.Contains("one", "o")
	// fmt.Println(containsMerger)

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
