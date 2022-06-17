package bs

import (
	"atms/ds"
	"atms/models"
	"database/sql"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func RegisterAccountHandlers() {
	router := mux.NewRouter()
	router.HandleFunc("/api/users/get", getUsersHandler).Methods(http.MethodGet)
	router.HandleFunc("/api/users/create", createUsersHandler).Methods(http.MethodPost)
	router.HandleFunc("/api/users/{username}/get", getUserByNameHandler).Methods(http.MethodGet)
	router.HandleFunc("/api/users/{userID}/update", updateUsersHandler).Methods(http.MethodPut)
	router.HandleFunc("/api/users/{userID}/delete", deleteUsersHandler).Methods(http.MethodDelete)
	router.HandleFunc("/api/login", userLoginHandler).Methods(http.MethodGet)
	// router.HandleFunc("/api/login", userLoginHandler).Methods(http.MethodGet)
	log.Fatal(http.ListenAndServe("localhost:1234", router))
}

func getUsersHandler(w http.ResponseWriter, r *http.Request) {

	authErr := Authenticate(r)
	if authErr != nil {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	accounts, err := ds.GetUserAccounts(dbds)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(accounts)
}

func getUserByNameHandler(w http.ResponseWriter, r *http.Request) {

	authErr := Authenticate(r)
	if authErr != nil {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	user, ok := mux.Vars(r)["username"]
	if !ok {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	accounts, err := ds.GetUserAccountByName(dbds, user)
	if err != nil {
		if err == sql.ErrNoRows {
			w.WriteHeader(http.StatusNotFound)
			return
		}
		log.Fatal(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	accByte, _ := json.Marshal(accounts)

	w.Header().Set("Content-Type", "application/json")
	w.Write(accByte)
}

func createUsersHandler(w http.ResponseWriter, r *http.Request) {

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	userPayload := models.Account{}
	err = json.Unmarshal(body, &userPayload)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	err = ds.CreateUserAccount(dbds, userPayload)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	accounts, err := ds.GetUserAccounts(dbds)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(accounts)
}

func updateUsersHandler(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	userID, ok := mux.Vars(r)["userID"]
	if !ok {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	convUserID, err := strconv.Atoi(userID)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	userPayload := models.Account{}
	err = json.Unmarshal(body, &userPayload)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	err = ds.UpdateUserAccount(dbds, convUserID, userPayload)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	accounts, err := ds.GetUserAccounts(dbds)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(accounts)
}

func deleteUsersHandler(w http.ResponseWriter, r *http.Request) {

	userID, ok := mux.Vars(r)["userID"]
	if !ok {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	convUserID, err := strconv.Atoi(userID)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	err = ds.DeleteUserAccount(dbds, convUserID)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	accounts, err := ds.GetUserAccounts(dbds)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(accounts)
}
