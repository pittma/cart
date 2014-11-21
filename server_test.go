package cart

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"testing"
	"time"
)

func mockNotFoundHandler(req *http.Request, rsp http.ResponseWriter, params map[string]string) {}
func mockHandler(req *http.Request, rsp http.ResponseWriter, params map[string]string) {
	fmt.Fprintf(rsp, req.URL.Path)
}

func TestFindRoot(t *testing.T) {
	srv := NewServer(8080, mockNotFoundHandler)
	srv.Get("/", mockHandler)
	srv.Get("/hi", mockHandler)
	srv.Get("/hi/dude", mockHandler)
	go srv.Serve()

	time.Sleep(100 * time.Millisecond)
	resp, _ := http.Get("http://localhost:8080/")
	response, _ := ioutil.ReadAll(resp.Body)

	if string(response) != "/" {
		t.Fatalf("Should find root [/] but found [" + string(response) + "]")
	}
}

func TestFindPathWhenMany(t *testing.T) {
	resp, _ := http.Get("http://localhost:8080/hi")
	response, _ := ioutil.ReadAll(resp.Body)

	if string(response) != "/hi" {
		t.Fatalf("Should find path [/hi] but found [" + string(response) + "]")
	}
}

func TestFindRootWhenMany(t *testing.T) {
	resp, _ := http.Get("http://localhost:8080/")
	response, _ := ioutil.ReadAll(resp.Body)

	if string(response) != "/" {
		t.Fatalf("Should find root [/] but found [" + string(response) + "]")
	}
}
