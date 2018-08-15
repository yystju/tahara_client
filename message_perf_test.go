// +build perf

package main

import (
	"encoding/xml"
	"io/ioutil"
	"path/filepath"
	"testing"
)

func Benchmark_message_10000_request(b *testing.B) {
	for i := 0; i < b.N; i++ {
		b, err := ioutil.ReadFile(filepath.Join(".", "data", "10000_request.xml"))

		if err != nil {
			panic(err)
		}

		message := new(Message)

		err = xml.Unmarshal(b, message)

		if err != nil {
			panic(err)
		}

		message.Body.Pcb.Barcode = "10000"

		_, err = xml.Marshal(message)

		if err != nil {
			panic(err)
		}
	}
}
