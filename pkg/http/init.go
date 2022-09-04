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

var httpRoutes []Route = []Route{
	{"/user", http.MethodPut, []martini.Handler{handlers.AddUser}, models.AddUserParam{}},
	{"/user", http.MethodGet, []martini.Handler{handlers.GetUsers}, nil},
	{"/user/:id/wallet", http.MethodGet, []martini.Handler{handlers.GetUserWallet}, nil},
	{"/user/:id/wallet/charge", http.MethodPost, []martini.Handler{handlers.ChargeWallet}, models.ChargeWalletParam{}},
}

func Init() {
	db := database.IDatabase(&mysql.MysqlDatabase{})
	c := martini.Classic()
	c.Use(render.Renderer())
	c.Map(db)

	for _, r := range httpRoutes {
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
