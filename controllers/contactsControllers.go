package controllers

import (
	"net/http"
	"go-contacts/models"
	"encoding/json"
	u "go-contacts/utils"
	"strconv"
	"github.com/gorilla/mux"
)

var CreateContact = func (w http.ResponseWriter, r *http.Request) {
	user := r.Context().Value("user_id").(uint) // Get the id of the user that send the request
	contact := &models.Contact{}
	err := json.NewDecoder(r.Body).Decode(contact)
	
	if err != nil {
		u.Respond(w, u.Message(false, "Error while decoding request body"))
		return
	}

	contact.User_FK = user
	resp := contact.Create()
	u.Respond(w, resp)
}

var GetContact = func(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		u.Respond(w, u.Message(false, "There was an error in your request"))
		return
	}

	data := models.GetContact(uint(id))
	resp := u.Message(true, "success")
	resp["data"] = data
	u.Respond(w, resp)
}

var GetContacts = func(w http.ResponseWriter, r *http.Request) {
	id := r.Context().Value("user_id") . (uint)
	data := models.GetContacts(uint(id))
	resp := u.Message(true, "Success getting contacts")
	resp["data"] = data
	u.Respond(w, resp)
}