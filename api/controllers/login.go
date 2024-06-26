package controllers

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/carlos1377/devbook/api/authentication"
	"github.com/carlos1377/devbook/api/database"
	"github.com/carlos1377/devbook/api/models"
	"github.com/carlos1377/devbook/api/repositories"
	"github.com/carlos1377/devbook/api/responses"
	"github.com/carlos1377/devbook/api/security"
)

func Login(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		responses.Error(w, http.StatusUnprocessableEntity, err)
		return
	}

	var user models.User
	if err = json.Unmarshal(body, &user); err != nil {
		responses.Error(w, http.StatusBadRequest, err)
		return
	}

	db, err := database.Connect()
	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	repository := repositories.NewUserRepository(db)
	userOnDb, err := repository.SearchByEmail(user.Email)
	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}

	if err = security.VerifyPassword(user.Password, userOnDb.Password); err != nil {
		responses.Error(w, http.StatusUnauthorized, err)
		return
	}

	token, err := authentication.CreateToken(userOnDb.ID)
	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}

	tokenResponse := map[string]string{
		"token": token,
	}

	responses.JSON(w, http.StatusOK, tokenResponse)
}
