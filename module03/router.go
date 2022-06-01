package module03

import "github.com/julienschmidt/httprouter"

func NewRouter(routes Routes) *httprouter.Router {
	router := httprouter.New()
	for _, r := range routes {
		router.Handle(r.Method, r.Path, r.HandlerFunc)
	}
	return router
}
