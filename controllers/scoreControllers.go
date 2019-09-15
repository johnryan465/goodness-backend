package controllers

import (
	"net/http"
	"goodness/models"
	u "goodness/utils"
)


var GetScore = func(w http.ResponseWriter, r *http.Request) {
	user := r.Context().Value("user") . (uint64)
	resp,_ := models.Scores(user)
	u.Respond(w, resp)
}
