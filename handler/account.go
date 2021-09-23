package handler

import (
	"bank/errs"
	"bank/service"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type accountHandler struct {
	accSrv service.AccountService
}

func NewAccountHandler (accSrv service.AccountService) accountHandler{
	return accountHandler{accSrv: accSrv}
}

func (h accountHandler) NewAccount(w http.ResponseWriter, r *http.Request){
	// get customer id
	customerID, _ := strconv.Atoi(mux.Vars(r)["customerID"])
	// check if type is not json
	if r.Header.Get("content-type") != "application/json"{
		handleError(w, errs.NewValidateError("Request body incorrect format, not json")) // error from font-end
		return
	}
	// create variable for store data
	request := service.NewAccountRequest{}
	// decode and append data to request variable
	err := json.NewDecoder(r.Body).Decode(&request)
	// check if decode error
	if err != nil {
		handleError(w, errs.NewValidateError("Request body incorrect format")) // error from font-end or back-end
		return
	}
	response, err := h.accSrv.NewAccount(customerID, request)
	if err != nil {
		handleError(w, err) // error from service business
		return
	}
	// create successfully return 201 status code
	w.WriteHeader(http.StatusCreated)
	w.Header().Set("content-type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func (h accountHandler) GetAccounts(w http.ResponseWriter, r *http.Request){
	customerID, _ := strconv.Atoi(mux.Vars(r)["customerID"])
	response, err := h.accSrv.GetAccounts(customerID)
	if err != nil {
		handleError(w, err) // error from service
		return
	}
	w.Header().Set("content-type", "application/json")
	err = json.NewEncoder(w).Encode(response)
	if err != nil {
		handleError(w, err)
		return
	}
}
