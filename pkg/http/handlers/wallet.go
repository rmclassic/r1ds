package handlers

import (
	"net/http"
	"strconv"
	"wallet/database"
	"wallet/models"
	"wallet/services"

	"github.com/go-martini/martini"
	"github.com/martini-contrib/render"
)

func GetWalletBalance(req *http.Request, db database.IDatabase, params martini.Params, r render.Render) {
	walletId, err := strconv.Atoi(params["id"])
	if err != nil {
		r.Status(http.StatusBadRequest)
		return
	}

	wallet, err := services.GetWalletBalance(db, walletId)
	if err != nil {
		r.Error(http.StatusInternalServerError)
		return
	}

	r.JSON(http.StatusOK, wallet)
}

func AddUser(req *http.Request, param models.AddUserParam, db database.IDatabase, r render.Render) {
	if err := services.AddUser(db, param.PhoneNumber); err != nil {
		r.Status(http.StatusInternalServerError)
		return
	}

	r.Status(http.StatusOK)
}

func ChargeWallet(req *http.Request, param models.ChargeWalletParam, r render.Render) {
	services.ChargeWallet(param.PhoneNumber, param.Amount)
}
