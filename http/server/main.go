package main

import (
	"encoding/json"
	"fmt"
	"github.com/davecgh/go-spew/spew"
	log "github.com/sirupsen/logrus"
	"io/ioutil"
	"net/http"
)

func main() {

	http.HandleFunc("/hello", hello)
	http.HandleFunc("/headers", headers)
	http.HandleFunc("/foo", foo)
	http.HandleFunc("/bar", bar)
	http.HandleFunc("/resp", resp)
	err := http.ListenAndServe(":9999", nil)
	if err != nil {
		log.Fatalf("unexpected server error %v", err)
	}
}

func hello(w http.ResponseWriter, req *http.Request) {
	spew.Dump(req)
	fmt.Fprintf(w, "hello\n")
}

type Foo struct {
	Email string
	Pass  string
}

func headers(w http.ResponseWriter, req *http.Request) {
	if req.Method != http.MethodPost {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	for name, headers := range req.Header {
		for _, h := range headers {
			fmt.Fprintf(w, "%v: %v\n", name, h)
		}
	}
}

func foo(w http.ResponseWriter, req *http.Request) {
	if req.Method != http.MethodPost {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	raw, err := ioutil.ReadAll(req.Body)
	defer req.Body.Close()

	if err != nil {
		log.Errorf("unexpected error reading body %v", err)
		w.WriteHeader(http.StatusExpectationFailed)
		return
	}

	spew.Dump(string(raw))

	res := make(map[string]string)
	err = json.Unmarshal(raw, &res)
	if err != nil {
		log.Errorf("unexpected error reading body %v", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	spew.Dump(res)
}

type Response struct {
	Foo string `json:"foo,omitempty"`
	Bar int    `json:"mybar,omitempty"`
}

func bar(w http.ResponseWriter, req *http.Request) {
	if req.Method != http.MethodPost {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	raw, err := ioutil.ReadAll(req.Body)
	defer req.Body.Close()
	if err != nil {
		log.Errorf("unexpected error reading body %v", err)
		w.WriteHeader(http.StatusExpectationFailed)
		return
	}

	spew.Dump(string(raw))

	res := Foo{}
	err = json.Unmarshal(raw, &res)
	if err != nil {
		log.Errorf("unexpected error reading body %v", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	resp := &Response{
		Foo: "ok",
		Bar: 234,
	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(resp)   // @TODO: Reply using responseWriter  EXPLAIN interfaces
	if err != nil {
		log.Errorf("Unexpected error Marshalling user profile, error %v", err)
	}
}

func resp(w http.ResponseWriter, req *http.Request) {
	resp := &Response{
		Foo: "ok",
		Bar: 234,
	}

	rw, err := json.Marshal(resp)
	if err != nil {
		log.Errorf("unexpected error marshalling response %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	n, err := w.Write(rw)
	if err != nil {
		log.Errorf("Unexpected error Marshalling user profile, error %v", err)
	}

	log.Infof("total writen: %d", n)


}
