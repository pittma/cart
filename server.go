package httprouter

import (
    "net/http"
)

type Server struct {
    router Router
}


func NewRouter(notFoundHandler RouterCallback) *Server {
    routes := make([]*Route, 0)
    r := Router{routes, notFoundHandler}
    return &Server{r}
}

func (s *Server) Get(path string, callback RouterCallback) {
    s.router.AddToRoutes(path, callback, "GET")
}

func (s *Server) Post(path string, callback RouterCallback) {
    s.router.AddToRoutes(path, callback, "POST")
}

func (s *Server) Start() {
    http.HandleFunc("/", s.router.RouteRequest)
    http.ListenAndServe(":8080", nil)
}
