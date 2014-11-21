cart
====

As in _map_ your URIs.

[![Build Status](https://travis-ci.org/danielscottt/cart.svg)](https://travis-ci.org/danielscottt/cart)

[![GoDoc](https://godoc.org/github.com/danielscottt/cart?status.svg)](https://godoc.org/github.com/danielscottt/cart)

Cart uses a [trie](http://en.wikipedia.org/wiki/Trie) based URI lookup algorithm.  This means that routing performance is independent from the number of routes.  When using a traditional URI dispatcher, worst case scenario time performance is O(n) where n is the total number of routes.  Using a trie, routes are dispatched at worst O(m) where m is the length of the longest URI.

### This app:
```go
package main

import (
	"fmt"
	"net/http"

	"github.com/danielscottt/cart"
)

func notFound(req *http.Request, rsp http.ResponseWriter, params map[string]string) {
	fmt.Fprintf(rsp, "not found")
}

func main() {
	server := cart.NewServer(8080, notFound)

	server.Get("/cheese/:test", func(req *http.Request, rsp http.ResponseWriter, params map[string]string) {
		fmt.Fprintf(rsp, params["test"])
	})

	server.Get("/hi/:you/its/:me", func(req *http.Request, rsp http.ResponseWriter, params map[string]string) {
		fmt.Fprintf(rsp, "hello "+params["you"]+", from "+params["me"])
	})

	server.Serve()

}
```

### Would do this:
```
$ curl http://localhost:8080/cheese/dan
dan

$ curl http://localhost:8080/hi/dan/its/meg
hello dan from meg

$ curl -i http://localhost:8080/nope
HTTP/1.1 404 Not Found
Date: Wed, 29 Oct 2014 16:45:16 GMT
Content-Length: 10
Content-Type: text/plain; charset=utf-8

not found
```
