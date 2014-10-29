package cart

import (
	"net/http"
	"regexp"
	"strings"
)

type route struct {
	path     string
	method   string
	callback RouterCallback
	params   map[string]string
	regex    *regexp.Regexp
	urlVars  []string
}

type router struct {
	Routes []*route

	notFoundHandler RouterCallback
}

func (r *router) ServeHTTP(rsp http.ResponseWriter, req *http.Request) {
	for _, rt := range r.Routes {
		match := (*rt).regex.FindAllStringSubmatch(req.URL.Path, -1)
		if match != nil && (*rt).method == req.Method {
			for i, m := range match[0][1:] {
				(*rt).params[(*rt).urlVars[i]] = m
			}
			(*rt).callback(req, rsp, (*rt).params)
			return
		}
	}
	rsp.WriteHeader(http.StatusNotFound)
	r.notFoundHandler(req, rsp, map[string]string{})
}

func (r *router) AddToRoutes(path string, callback RouterCallback, method string) {
	rt := &route{
		path:     path,
		method:   method,
		callback: callback,
		params:   make(map[string]string),
		urlVars:  make([]string, 0),
	}
	sPath := strings.Split(string(path[1:]), "/")
	for i, n := range sPath {
		if string(n[0]) == ":" {
			rt.urlVars = append(rt.urlVars, string(n[1:]))
			sPath[i] = "([a-zA-Z0-9]+)"
		}
	}
	rt.regex = regexp.MustCompile("^/" + strings.Join(sPath, "/") + "$")
	r.Routes = append(r.Routes, rt)
}
