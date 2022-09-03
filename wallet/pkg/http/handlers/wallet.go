package handlers

import (
	"net/http"
	"strconv"
	"wallet/database"
	"wallet/services"

	"github.com/go-martini/martini"
	"github.com/martini-contrib/render"
)

func GetWalletBalance(req *http.Request, db database.IDatabase, params martini.Params, r render.Render) {
	walletId, err := strconv.Atoi(params["id"])
	if err != nil {
		r.Status(http.StatusInternalServerError)
		return
	}

	wallet, err := services.GetWalletBalance(db, walletId)
	if err != nil {
		r.Error(http.StatusInternalServerError)
		return
	}

	r.JSON(http.StatusOK, wallet)
}
