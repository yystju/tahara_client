// +build ut

package main

import (
	"encoding/xml"
	"io/ioutil"
	"path/filepath"
	"testing"
)

func TestMessage_10000_request(t *testing.T) {
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

func TestMessage_10006_request(t *testing.T) {
	b, err := ioutil.ReadFile(filepath.Join(".", "data", "10006_request.xml"))

	if err != nil {
		panic(err)
	}

	message := new(Message)

	err = xml.Unmarshal(b, message)

	if err != nil {
		panic(err)
	}

	t.Logf("message.Body.ModelInfo.ModelNo : %s", message.Body.ModelInfo.ModelNo)

	_, err = xml.Marshal(message)

	if err != nil {
		panic(err)
	}
}

func TestMessage_10006_response(t *testing.T) {
	b, err := ioutil.ReadFile(filepath.Join(".", "data", "10006_response.xml"))

	if err != nil {
		panic(err)
	}

	message := new(Message)

	err = xml.Unmarshal(b, message)

	if err != nil {
		panic(err)
	}

	t.Logf("len(message.Body.Resource.Comps.List) : %d", len(message.Body.Resource.Comps.List))

	_, err = xml.Marshal(message)

	if err != nil {
		panic(err)
	}
}

// func TestMessage_501(t *testing.T) {
// 	f, err := os.OpenFile(filepath.Join(".", "data", "501.xml"), os.O_RDONLY, 0666)
// 	checkerr(err)
// 	defer f.Close()

// 	b, err := ioutil.ReadAll(f)
// 	checkerr(err)

// 	message := new(Message)

// 	err = xml.Unmarshal(b, message)
// 	checkerr(err)

// 	message.Body.Pcb.Barcode = "501501501501501501501501501501"

// 	m, err := xml.Marshal(message)

// 	checkerr(err)

// 	fmt.Println(string(m))
// }

// func TestMessage_550(t *testing.T) {
// 	f, err := os.OpenFile(filepath.Join(".", "data", "550.xml"), os.O_RDONLY, 0666)
// 	checkerr(err)
// 	defer f.Close()

// 	b, err := ioutil.ReadAll(f)
// 	checkerr(err)

// 	message := new(Message)

// 	err = xml.Unmarshal(b, message)
// 	checkerr(err)

// 	message.Body.Panel.PcbID = "550550550550550550550550550550"

// 	m, err := xml.Marshal(message)

// 	checkerr(err)

// 	fmt.Println(string(m))
// }

// func TestMessage_580(t *testing.T) {
// 	f, err := os.OpenFile(filepath.Join(".", "data", "580.xml"), os.O_RDONLY, 0666)
// 	checkerr(err)
// 	defer f.Close()

// 	b, err := ioutil.ReadAll(f)
// 	checkerr(err)

// 	message := new(Message)

// 	err = xml.Unmarshal(b, message)
// 	checkerr(err)

// 	message.Body.ModelInfo.ModelNo = "580580580580580580580580580580"

// 	m, err := xml.Marshal(message)

// 	checkerr(err)

// 	fmt.Println(string(m))
// }
