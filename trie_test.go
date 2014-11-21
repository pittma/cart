package cart

import (
	"testing"
)

func TestAddRoute(t *testing.T) {
	srv := NewServer(80, mockHandler)
	srv.Get("/hello", mockHandler)
	rt := srv.router.routes.Find([]string{"hello"}, "GET")

	if rt == nil {
		t.Fatalf("Should add route")
	}
}

func TestRouteNotFound(t *testing.T) {
	srv := NewServer(80, mockHandler)
	srv.Post("/hello", mockHandler)
	rt := srv.router.routes.Find([]string{"goodbye"}, "POST")

	if rt != nil {
		t.Fatalf("Should return nil for route, returned [" + rt.key.path + "]")
	}
}

func TestMethodNotFound(t *testing.T) {
	srv := NewServer(80, mockNotFoundHandler)
	srv.Post("/hello", mockNotFoundHandler)
	rt := srv.router.routes.Find([]string{"hello"}, "GET")

	if rt != nil {
		t.Fatalf("Should return nil for route, returned [" + rt.key.path + "]")
	}
}
