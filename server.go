package cart

import (
	"net/http"
	"strconv"
)

// Server is the main http server in cart
type Server struct {
	Port   port
	router *router
}

type port int

// A method signature type that can be used to define a function to be used when a route is dispatched
type RouterCallback func(*http.Request, http.ResponseWriter, map[string]string)

func (p port) String() string {
	return ":" + strconv.Itoa(int(p))
}

// Returns an initialized server, with the routing trie initialized as well
func NewServer(prt int, notFoundHandler RouterCallback) *Server {
	r := &router{
		routes:          newTrie(),
		rootHandlers:    make(map[string]*route),
		notFoundHandler: notFoundHandler,
	}
	return &Server{
		Port: port(prt),

		router: r,
	}
}

// defines a GET path on the server
func (s *Server) Get(path string, callback RouterCallback) {
	s.router.AddToRoutes(path, callback, "GET")
}

// defines a POST path on the server
func (s *Server) Post(path string, callback RouterCallback) {
	s.router.AddToRoutes(path, callback, "POST")
}

// defines a PUT path on the server
func (s *Server) Put(path string, callback RouterCallback) {
	s.router.AddToRoutes(path, callback, "PUT")
}

// defines a DELETE path on the server
func (s *Server) Delete(path string, callback RouterCallback) {
	s.router.AddToRoutes(path, callback, "DELETE")
}

// Starts the http.ListenAndServe server, using the router created by route definitions
func (s *Server) Serve() {
	http.ListenAndServe(s.Port.String(), s.router)
}
