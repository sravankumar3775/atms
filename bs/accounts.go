package bs

import (
	"atms/ds"
	"atms/models"
	"encoding/json"
	"fmt"
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
	router.HandleFunc("/api/users/{userID}/get", getUserWithIDHandler).Methods(http.MethodGet)
	router.HandleFunc("/api/users/{userID}/update", updateUsersHandler).Methods(http.MethodPut)
	router.HandleFunc("/api/users/{userID}/delete", deleteUsersHandler).Methods(http.MethodDelete)
	router.HandleFunc("/api/login", userLoginHandler).Methods(http.MethodGet)
	log.Fatal(http.ListenAndServe("localhost:1234", router))
}

func getUsersHandler(w http.ResponseWriter, r *http.Request) {
	accounts, err := ds.GetUserAccounts(dbds)
	if err != nil {
		fmt.Println("err")
		panic(err.Error())
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(accounts)
}

func getUserWithIDHandler(w http.ResponseWriter, r *http.Request) {
	accounts, err := ds.GetUserAccounts(dbds)
	if err != nil {
		fmt.Println("err")
		panic(err.Error())
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(accounts)
}

func createUsersHandler(w http.ResponseWriter, r *http.Request) {

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err.Error())
	}

	userPayload := models.Account{}
	err = json.Unmarshal(body, &userPayload)
	if err != nil {
		panic(err.Error())
	}

	err = ds.CreateUserAccount(dbds, userPayload)
	if err != nil {
		fmt.Println("err")
		panic(err.Error())
	}

	accounts, err := ds.GetUserAccounts(dbds)
	if err != nil {
		fmt.Println("err")
		panic(err.Error())
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(accounts)
}

func updateUsersHandler(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err.Error())
	}

	userID, ok := mux.Vars(r)["userID"]
	if !ok {
		panic("error converting user id")
	}

	convUserID, err := strconv.Atoi(userID)
	if err != nil {
		panic(err.Error())
	}
	userPayload := models.Account{}
	err = json.Unmarshal(body, &userPayload)
	if err != nil {
		panic(err.Error())
	}

	err = ds.UpdateUserAccount(dbds, convUserID, userPayload)
	if err != nil {
		fmt.Println("err")
		panic(err.Error())
	}

	accounts, err := ds.GetUserAccounts(dbds)
	if err != nil {
		fmt.Println("err")
		panic(err.Error())
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(accounts)
}

func deleteUsersHandler(w http.ResponseWriter, r *http.Request) {

	userID, ok := mux.Vars(r)["userID"]
	if !ok {
		panic("error converting userid")
	}

	convUserID, err := strconv.Atoi(userID)
	if err != nil {
		panic(err.Error())
	}

	err = ds.DeleteUserAccount(dbds, convUserID)
	if err != nil {
		fmt.Println("err")
		panic(err.Error())
	}

	accounts, err := ds.GetUserAccounts(dbds)
	if err != nil {
		fmt.Println("err")
		panic(err.Error())
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(accounts)
}
