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

func GetUserWallet(req *http.Request, db database.IDatabase, params martini.Params, r render.Render) {
	userId, err := strconv.Atoi(params["id"])
	if err != nil {
		r.Status(http.StatusBadRequest)
		return
	}

	wallet, err := services.GetUserWallet(db, userId)
	if err != nil {
		r.Status(http.StatusInternalServerError)
		return
	}

	if wallet.UserID == 0 {
		r.Status(http.StatusNotFound)
		return
	}

	r.JSON(http.StatusOK, wallet)
}

func ChargeWallet(req *http.Request, db database.IDatabase, params martini.Params, model models.ChargeWalletParam, r render.Render) {
	userId, err := strconv.Atoi(params["id"])
	if err != nil {
		r.Status(http.StatusBadRequest)
		return
	}

	err = services.ChargeWallet(db, userId, model.Amount)
	if err != nil {
		r.Status(http.StatusBadRequest)
		return
	}

	r.Status(http.StatusOK)
}
