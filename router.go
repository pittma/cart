package httprouter

import (
    "regexp"
    "net/http"
)

type RouterCallback func(*http.Request, http.ResponseWriter)

type Route struct {
    path, method string
    callback RouterCallback
}

type Router struct {
    routes []*Route
    notFoundHandler RouterCallback
}

func (r *Router) RouteRequest(rsp http.ResponseWriter, req *http.Request) {
    for _, rt := range r.routes {
        match, err := regexp.MatchString(rt.path, req.URL.Path)
        if err != nil {
            rsp.WriteHeader(http.StatusInternalServerError)
            return
        }
        if  match && rt.method == req.Method {
            rt.callback(req, rsp)
            return
        }
    }
    rsp.WriteHeader(http.StatusNotFound)
    r.notFoundHandler(req, rsp)
}

func (r *Router) AddToRoutes(path string, callback RouterCallback, method string) {
    r.routes = append(r.routes, &Route{path, method, callback})
}
