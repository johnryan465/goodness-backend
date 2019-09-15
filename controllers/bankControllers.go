package controllers

import (
	"net/http"
	"goodness/models"
	u "goodness/utils"
	"encoding/json"
)


var SetToken = func(w http.ResponseWriter, r *http.Request) {
	user := r.Context().Value("user") . (uint64)
	token := &models.RefreshToken{}

	err := json.NewDecoder(r.Body).Decode(token)
	if err != nil {
		u.Respond(w, u.Message(false, "Error while decoding request body"))
		return
	}

	token.User = user
	resp := token.Save()
	u.Respond(w, resp)
}

var GetData = func(w http.ResponseWriter, r *http.Request) {
	user := r.Context().Value("user") . (uint64)
	resp,_ := models.Scores(user)
	u.Respond(w, resp)
}

var GetTransactions = func(w http.ResponseWriter, r *http.Request) {
	user := r.Context().Value("user") . (uint64)
	resp,_ := models.Scores(user)
	u.Respond(w, resp)
}
