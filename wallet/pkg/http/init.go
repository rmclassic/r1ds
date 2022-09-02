package http

import (
	"net/http"
	"wallet/pkg/http/handlers"

	"github.com/go-martini/martini"
	"github.com/martini-contrib/binding"
)

type Route struct {
	Path     string
	Method   string
	Handlers []martini.Handler
	Param    interface{}
}

var routes []Route = []Route{
	{"/wallet", http.MethodGet, []martini.Handler{handlers.GetWallet}, nil},
}

func Init() {
	c := martini.Classic()
	for _, r := range routes {
		ParamHandlerConcat := make([]martini.Handler, 0)
		if r.Param != nil {
			ParamHandlerConcat = append(ParamHandlerConcat, binding.Bind(r.Param))
		}
		ParamHandlerConcat = append(ParamHandlerConcat, c.Handlers)

		c.AddRoute(r.Path, r.Path, ParamHandlerConcat...)
	}

	c.Run()
}
