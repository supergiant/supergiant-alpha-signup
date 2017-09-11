package main

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"strconv"
	"strings"

	"github.com/gorilla/mux"
	// _ "github.com/lib/pq"
	_ "github.com/mattn/go-sqlite3"
	"github.com/op/go-logging"
)

var log = logging.MustGetLogger("alpha")

type Invite struct {
	ID       int    `json:"id"`
	Email    string `json:"email"`
	Invite   string `json:"invite"`
	URL      string `json:"url"`
	Password string `json:"password"`
	Claimed  bool   `json:"claimed"`
}

type Reset struct {
	Email   string `json:"email"`
	Token   string `json:"token"`
	NewPass string `json:"pw"`
}

// App basic app struct
type App struct {
	Router *mux.Router
	DB     *sql.DB
}

// func (a *App) pwReset(w http.ResponseWriter, r *http.Request) {
// 	var rp Reset
// 	decoder := json.NewDecoder(r.Body)
// 	if err := decoder.Decode(&rp); err != nil {
// 		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
// 		return
// 	}
// 	defer r.Body.Close()
//
// 	_, err := a.DB.Exec("DELETE FROM products WHERE id=$1", rp.Token)
// 	if err != nil {
// 		respondWithError(w, http.StatusInternalServerError, err.Error())
// 		return
// 	}
//
// 	respondWithJSON(w, http.StatusCreated, rp)
// }

func (a *App) useInvite(w http.ResponseWriter, r *http.Request) {
	vars := r.URL.Query()
	log.Debug(vars)
	invite, ok := vars["invite"]
	if !ok {
		respondWithError(w, http.StatusBadRequest, "Invalid Invite")
		return
	}
	email, ok := vars["email"]
	if !ok {
		respondWithError(w, http.StatusBadRequest, "Invalid Email")
		return
	}
	_ = email
	row, err := a.DB.Query("SELECT COUNT(*) FROM invites WHERE invite=$1 AND claimed=0 LIMIT 1", strings.Join(invite, ""))
	//row, err := a.DB.Query("SELECT * FROM invites WHERE invite='$1' AND claimed=123 LIMIT 1", strings.Join(invite, ""))
	// I am doing something wrong here becaose ErrNoRows is never caught, I am somehow always getting a nil err value
	log.Debug(row)
	log.Debug(err)
	if err != nil {
		switch err {
		case sql.ErrNoRows:
			respondWithError(w, http.StatusNotFound, "Invalid Invite Code")
		default:
			respondWithError(w, http.StatusInternalServerError, err.Error())
		}
		return
	}
	defer row.Close()
	var count string
	for row.Next() {
		if err := row.Scan(&count); err != nil {
			log.Debug(err)
			log.Debug("Row scan failed")
			respondWithError(w, http.StatusNotFound, "Invalid Invite Code")
		}
		log.Debug(count)
	}

	if count == "0" {
		respondWithError(w, http.StatusNotFound, "Invalid Invite Code")
	}
	// this is a terrible hack because of the no rows error not bubbling up
	row, err = a.DB.Query("SELECT * FROM invites WHERE invite=$1 AND claimed=0 LIMIT 1", strings.Join(invite, ""))
	if err != nil {
		switch err {
		case sql.ErrNoRows:
			respondWithError(w, http.StatusNotFound, "Invalid Invite Code")
		default:
			respondWithError(w, http.StatusInternalServerError, err.Error())
		}
		return
	}

	if row.Next() {
		//for row.Next() {
		var i Invite
		// byte arrays to work around null fields
		var col2, col3, col4, col5 []byte
		if err := row.Scan(&i.ID, &col2, &col3, &col4, &col5); err != nil {
			log.Debug(err)
			log.Debug("Row scan failed")
			respondWithError(w, http.StatusNotFound, "Invalid Invite Code")
		}
		i.Email = string(col2)
		i.Invite = string(col3)
		i.URL = string(col4)
		i.Claimed, err = strconv.ParseBool(string(col5))
		log.Debug(i)
		// TODO: call function to launch helm chart
		//}
	}
	respondWithJSON(w, http.StatusOK, invite)
}

func respondWithError(w http.ResponseWriter, code int, message string) {
	respondWithJSON(w, code, map[string]string{"error": message})
}

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

// Run main app forever
func (a *App) Run(addr string) {
	log.Debug("Starting API Server")
	log.Fatal(http.ListenAndServe(":8080", a.Router))
}

func main() {
	a := App{}
	db, err := sql.Open("sqlite3", "./local.db")
	if err != nil {
		log.Fatal("Failed to init DB")
	}
	a.DB = db
	a.Router = mux.NewRouter()
	a.Router.HandleFunc("/claim", a.useInvite).Methods("GET")
	//a.Router.HandleFunc("/request", a.createProduct).Methods("POST")
	a.Run(":8080")
}
