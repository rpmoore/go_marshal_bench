package main

import (
	"encoding/json"
	"fmt"
	"log"
	"time"

	"labix.org/v2/mgo/bson"
)

type TestStruct struct {
	Name      string
	CreatedOn time.Time
	Number    uint64
}

type BenchSuite struct {
	N    int
	data []TestStruct
}

func NewBenchSuite(size int) BenchSuite {
	data := SetUpSuite(size)
	return BenchSuite{size, data}
}

func SetUpSuite(size int) []TestStruct {
	fmt.Printf("Setting up test data with %d entries\n", size)
	data := make([]TestStruct, size, size)
	for i := 0; i < size; i++ {
		data[i] = TestStruct{"A Name", time.Now(), 1234567890}
	}
	return data
}

func (s *BenchSuite) BenchmarkJson() {
	for i := 0; i < s.N; i++ {
		var result TestStruct
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
		var result TestStruct
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

func runBench() {
	sizes := [...]int{100000, 1000000, 5000000}
	for _, size := range sizes {
		bench := NewBenchSuite(size)
		startTime := time.Now()
		bench.BenchmarkJson()
		fmt.Printf("Json: %v\n", time.Since(startTime))
		startTime = time.Now()
		bench.BenchmarkBson()
		fmt.Printf("Bson: %v\n", time.Since(startTime))
	}
}

func main() {
	fmt.Printf("Starting Benchmark\n")
	runBench()
	fmt.Printf("Finished Benchmark\n")
}
