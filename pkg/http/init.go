package http

import (
	"net/http"
	"wallet/database"
	"wallet/database/mysql"
	"wallet/models"
	"wallet/pkg/http/handlers"

	"github.com/go-martini/martini"
	"github.com/martini-contrib/binding"
	"github.com/martini-contrib/render"
)

type Route struct {
	Path     string
	Method   string
	Handlers []martini.Handler
	Param    interface{}
}

var routes []Route = []Route{
	{"/wallet/:id", http.MethodGet, []martini.Handler{handlers.GetWalletBalance}, nil},
	{"/wallet", http.MethodPut, []martini.Handler{handlers.AddUser}, models.AddUserParam{}},
}

func Init() {
	db := database.IDatabase(&mysql.MysqlDatabase{})
	c := martini.Classic()
	c.Use(render.Renderer())
	c.Map(db)

	for _, r := range routes {
		ParamHandlerConcat := make([]martini.Handler, 0)
		if r.Param != nil {
			ParamHandlerConcat = append(ParamHandlerConcat, binding.Bind(r.Param))
		}
		ParamHandlerConcat = append(ParamHandlerConcat, r.Handlers...)
		c.AddRoute(r.Method, r.Path, ParamHandlerConcat...)
	}

	db.Init()
	c.Run()
}
