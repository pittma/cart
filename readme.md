cart
====

Map your URIs

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

}
```

```
$ curl http://localhost:8080/hi/cheese/dan
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
