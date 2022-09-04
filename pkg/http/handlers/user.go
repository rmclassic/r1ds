package handlers

import (
	"net/http"
	"wallet/database"
	"wallet/models"
	"wallet/services"
	"wallet/util"

	"github.com/martini-contrib/render"
)

func AddUser(req *http.Request, param models.AddUserParam, db database.IDatabase, r render.Render) {
	user := &models.User{
		PhoneNumber: param.PhoneNumber,
	}

	if err := services.AddUser(db, user); err != nil {
		r.Status(http.StatusInternalServerError)
		return
	}

	r.JSON(http.StatusOK, user)
}

func GetUsers(req *http.Request, db database.IDatabase, r render.Render) {
	phoneFilter := util.QueryParamOrDefault(req.URL.Query(), "phone", "").(string)
	if phoneFilter == "" {
		r.Status(http.StatusBadRequest)
		return
	}

	user, err := services.GetUserByPhoneNumber(db, phoneFilter)
	if err != nil {
		r.Status(http.StatusInternalServerError)
		return
	}

	if user.ID == 0 {
		r.Status(http.StatusNotFound)
		return
	}

	r.JSON(http.StatusOK, user)
}
