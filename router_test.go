package cart

import (
	"testing"
)

func TestSetRoutePath(t *testing.T) {
	srv := NewServer(80, mockHandler)
	srv.Get("/hello", mockHandler)
	rt := srv.router.routes.Find([]string{"hello"}, "GET")

	if rt.key.path != "/hello" {
		t.Fatalf("Should add route [/hello] added [" + rt.key.path + "]")
	}
}

func TestSetRouteMethod(t *testing.T) {
	srv := NewServer(80, mockHandler)
	srv.Post("/hello", mockHandler)
	rt := srv.router.routes.Find([]string{"hello"}, "POST")

	if rt.key.method != "POST" {
		t.Fatalf("Should set method [POST] set [" + rt.key.method + "]")
	}
}

func TestDispatchCorrectMethod(t *testing.T) {
	srv := NewServer(80, mockHandler)
	srv.Get("/hello", mockHandler)
	srv.Post("/hello", mockHandler)

	rt := srv.router.routes.Find([]string{"hello"}, "POST")
	if rt.key.method != "POST" {
		t.Fatalf("Should set method [POST] set [" + rt.key.method + "]")
	}

	rt = srv.router.routes.Find([]string{"hello"}, "GET")
	if rt.key.method != "GET" {
		t.Fatalf("Should set method [GET] set [" + rt.key.method + "]")
	}
}
