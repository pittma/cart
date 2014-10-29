package cart

import (
	"net/http"
	"strconv"
)

type Server struct {
	Port   port
	router *router
}

type port int

type RouterCallback func(*http.Request, http.ResponseWriter, map[string]string)

func (p port) String() string {
	return ":" + strconv.Itoa(int(p))
}

func NewServer(prt int, notFoundHandler RouterCallback) *Server {
	routes := make([]*route, 0)
	r := &router{
		Routes: routes,

		notFoundHandler: notFoundHandler,
	}
	return &Server{
		Port: port(prt),

		router: r,
	}
}

func (s *Server) Get(path string, callback RouterCallback) {
	s.router.AddToRoutes(path, callback, "GET")
}

func (s *Server) Post(path string, callback RouterCallback) {
	s.router.AddToRoutes(path, callback, "POST")
}

func (s *Server) Serve() {
	http.ListenAndServe(s.Port.String(), s.router)
}
