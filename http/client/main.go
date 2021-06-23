package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type Response struct {
	Foo string `json:"foo,omitempty"`
	Bar int    `json:"mybar,omitempty"`
}


func main() {
	fmt.Println("HEY from client")

	resp := &Response{
		Foo: "ok",
		Bar: 234,
	}

	rw, err := json.Marshal(resp)
	if err != nil {
		log.Fatalf("unexpected error marshalling response %v", err)
	}

	buf := bytes.NewBuffer(rw)
	req, err := http.NewRequest(http.MethodPost, "http://localhost:9999/bar", buf)
	if err != nil {
		log.Fatalf("unexpected error creating request  %v", err)
	}

	rp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatalf("unexpected error creating request  %v", err)
	}

	raw, err := ioutil.ReadAll(rp.Body)
	defer rp.Body.Close()
	if err != nil {
		log.Fatalf("unexpected error creating request  %v", err)
	}

	log.Printf("result: %s", string(raw))
}
