package main

import (
	"encoding/json"
	"fmt"
	"log"
//	"os"
	"time"

	"code.google.com/p/goprotobuf/proto"
	"labix.org/v2/mgo/bson"
)

type InternalMessage struct {
	Name             string `json:"name,omitempty"`
	CreatedOn        int64  `json:"createdOn,omitempty"`
	Number           uint64 `json:"number,omitempty"`
}

type BenchSuite struct {
	N    int
	data []Message
	internalData []InternalMessage
}

func NewBenchSuite(size int) BenchSuite {
	data := SetUpMessages(size)
	internalData := SetUpInternalMessages(size)
	return BenchSuite{size, data, internalData}
}

func SetUpMessages(size int) []Message {
	fmt.Printf("Setting up test data with %d entries\n", size)
	data := make([]Message, size, size)
	for i := 0; i < size; i++ {
		data[i] = Message{Name: proto.String("A Name"), CreatedOn: proto.Int64(time.Now().Unix()), Number: proto.Uint64(1234567890)}
	}
	return data
}

func SetUpInternalMessages(size int) []InternalMessage {
	fmt.Printf("Setting up test data with %d entries\n", size)
	data := make([]InternalMessage, size, size)
	for i := 0; i < size; i++ {
		data[i] = InternalMessage{"A Name", time.Now().Unix(), 1234567890}
	}
	return data
}

func (s *BenchSuite) BenchmarkJson() {
	for i := 0; i < s.N; i++ {
		var result Message
		encoding, err := json.Marshal(&s.data[i])
		if err != nil {
			log.Fatal(err)
		}
		err = json.Unmarshal(encoding, &result)
		if err != nil {
			log.Fatal(err)
		}
	}
}

func (s *BenchSuite) BenchmarkBson() {
	for i := 0; i < s.N; i++ {
		var result Message
		encoding, err := bson.Marshal(&s.data[i])
		if err != nil {
			log.Fatal(err)
		}
		err = bson.Unmarshal(encoding, &result)
		if err != nil {
			log.Fatal(err)
		}
	}
}

func (s *BenchSuite) BenchmarkProtoBuf() {
	for i := 0; i < s.N; i++ {
		var result Message
		encoding, err := proto.Marshal(&s.data[i])
		if err != nil {
			log.Fatal(err)
		}
		err = proto.Unmarshal(encoding, &result)
		if err != nil {
			log.Fatal(err)
		}
	}
}

func (s *BenchSuite) BenchmarkInternalJson() {
	for i := 0; i < s.N; i++ {
		var result Message
		encoding, err := json.Marshal(&s.internalData[i])
		if err != nil {
			log.Fatal(err)
		}
		err = json.Unmarshal(encoding, &result)
		if err != nil {
			log.Fatal(err)
		}
	}
}

func (s *BenchSuite) BenchmarkInternalBson() {
	for i := 0; i < s.N; i++ {
		var result Message
		encoding, err := bson.Marshal(&s.internalData[i])
		if err != nil {
			log.Fatal(err)
		}
		err = bson.Unmarshal(encoding, &result)
		if err != nil {
			log.Fatal(err)
		}
	}
}

func runBench() {
	sizes := [...]int{100000, 1000000, 5000000}
	for _, size := range sizes {
		bench := NewBenchSuite(size)
		startTime := time.Now()
		bench.BenchmarkInternalJson()
		fmt.Printf("Internal Json:\t%v\n", time.Since(startTime))
		startTime = time.Now()
		bench.BenchmarkJson()
		fmt.Printf("Json:\t\t%v\n", time.Since(startTime))
		startTime = time.Now()
		bench.BenchmarkInternalBson()
		fmt.Printf("Internal Bson:\t%v\n", time.Since(startTime))
		startTime = time.Now()
		bench.BenchmarkBson()
		fmt.Printf("Bson:\t\t%v\n", time.Since(startTime))
		startTime = time.Now()
		bench.BenchmarkProtoBuf()
		fmt.Printf("ProtoBuf:\t%v\n", time.Since(startTime))
	}
}

func main() {
	fmt.Printf("Starting Benchmark\n")

/*
	outWriter := json.NewEncoder(os.Stdout)
	message := Message{Name: proto.String("A Name"), CreatedOn: proto.Int64(time.Now().Unix()), Number: proto.Uint64(1234567890)}
	outWriter.Encode(message)
	encoding, err := bson.Marshal(message)
	if err != nil {
		log.Fatal(err)
	}
	os.Stdout.Write(encoding)
*/
	runBench()
	fmt.Printf("Finished Benchmark\n")
}
