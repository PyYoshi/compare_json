package main

import (
	"encoding/json"
	"fmt"
	"os"
	"time"
)

func perror(err error) {
	if err != nil {
		panic(err)
		return
	}
}

func loadJson() []byte {
	file, err := os.Open("./public_timeline.json")
	perror(err)
	defer file.Close()
	stat, err := file.Stat()
	perror(err)
	bs := make([]byte, stat.Size())
	_, err = file.Read(bs)
	perror(err)
	return bs
}

func decode(jsonString []byte, loop int) {
	fmt.Println("=== json decode(interface{}) ===")
	var totalTime float64
	for i := 0; i < loop; i++ {
		var obj interface{}
		start := time.Now()
		json.Unmarshal(jsonString, &obj)
		c := time.Now().Sub(start)
		totalTime += c.Seconds()
	}
	fmt.Printf("%f call(s)/s\n", float64(1/(totalTime/float64(loop))))
}

func encode(obj interface{}, loop int) {
	fmt.Println("=== json encode(interface{}) ===")
	var totalTime float64
	for i := 0; i < loop; i++ {
		start := time.Now()
		json.Marshal(obj)
		c := time.Now().Sub(start)
		totalTime += c.Seconds()
	}
	fmt.Printf("%f call(s)/s\n", float64(1/(totalTime/float64(loop))))
}

func main() {
	loop := 10000
	jsonString := loadJson()
	var jsonObj interface{}
	json.Unmarshal(jsonString, &jsonObj)
	decode(jsonString, loop)
	encode(jsonObj, loop)
}
