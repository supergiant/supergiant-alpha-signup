package main

import (
	"bytes"
	"crypto/tls"
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/urfave/cli"
	// _ "github.com/lib/pq"
	// _ "github.com/mattn/go-sqlite3"
	_ "github.com/lib/pq"
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

const (
	// http://stackoverflow.com/questions/22892120/how-to-generate-a-random-string-of-a-fixed-length-in-golang
	letterBytes   = "abcdefghijklmnopqrstuvwxyz0123456789"
	letterIdxBits = 6                    // 6 bits to represent a letter index
	letterIdxMask = 1<<letterIdxBits - 1 // All 1-bits, as many as letterIdxBits
	letterIdxMax  = 63 / letterIdxBits   // # of letter indices fitting in 63 bits
)

func RandomString(n int) string {
	b := make([]byte, n)
	src := rand.NewSource(time.Now().UnixNano())
	// A src.Int63() generates 63 random bits, enough for letterIdxMax characters!
	for i, cache, remain := n-1, src.Int63(), letterIdxMax; i >= 0; {
		if remain == 0 {
			cache, remain = src.Int63(), letterIdxMax
		}
		if idx := int(cache & letterIdxMask); idx < len(letterBytes) {
			b[i] = letterBytes[idx]
			i--
		}
		cache >>= letterIdxBits
		remain--
	}
	return string(b)
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
		//}
		customer := RandomString(16)
		helmJSON := []byte(`{
		  "chart_name": "supergiant",
		  "chart_version": "0.1.0",
		  "config": {
		    "api": {
		      "enabled": true,
		      "image": {
		        "pullPolicy": "Always",
		        "repository": "supergiant/supergiant-api",
		        "tag": "unstable-linux-x64"
		      },
		      "name": "supergiant-api",
		      "resources": {},
		      "service": {
		        "externalPort": 80,
		        "internalPort": 8080
		      },
		      "support": {
		        "enabled": true,
		        "password": "cheese1234"
		      }
		    },
		    "ingress": {
		      "annotations": {
		        "traefik.frontend.rule.type": "PathPrefixStrip"
		      },
		      "enabled": true,
		      "name": "supergiant"
		    },
		    "persistence": {
		      "accessMode": "ReadWriteOnce",
		      "enabled": true,
		      "size": "8Gi",
		      "storageClass": "generic"
		    },
		    "ui": {
		      "enabled": true,
		      "image": {
		        "pullPolicy": "Always",
		        "repository": "supergiant/supergiant-ui",
		        "tag": "unstable-linux-x64"
		      },
		      "name": "supergiant-ui",
		      "replicaCount": 1,
		      "resources": {},
		      "service": {
		        "externalPort": 80,
		        "internalPort": 3001
		      }
		    },
		    "uniqueurl": "` + customer + `"
		  },
		  "kube_name": "sgalpha1",
		  "name": "` + customer + `",
		  "repo_name": "supergiant",
		  "namespace": "` + customer + `"
		}`)
		tr := &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		}
		client := &http.Client{Transport: tr}

		_, err = client.Get("https://admin.alpha.supergiant.io/api/v0/helm_releases")
		if err != nil {
			log.Error(err)
			log.Error("Failed to launch")
		}

		log.Debug(string(helmJSON))

		req, err := http.NewRequest("POST", "https://admin.alpha.supergiant.io/api/v0/helm_releases", bytes.NewBuffer(helmJSON))
		if err != nil {
			log.Error(err)
			log.Error("Failed to launch")
		}
		req.Header.Add("Authorization", `SGAPI token=""`)
		//req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
		req.Header.Add("Content-Type", `application/json`)
		log.Debug(req)
		log.Debug("------")
		log.Debug(req.Body)
		resp, err := client.Do(req)
		log.Debug(resp)
		bs, _ := ioutil.ReadAll(resp.Body)
		log.Debug(string(bs))
		if err != nil {
			log.Error(err)
			log.Error("Failed to launch")
		}
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
func (a *App) Run() {
	log.Debug("Starting API Server")
	headersOk := handlers.AllowedHeaders([]string{"Access-Control-Request-Headers", "Authorization"})
	methodsOk := handlers.AllowedMethods([]string{"GET", "PUT", "UPDATE", "POST", "DELETE"})
	log.Fatal(http.ListenAndServe(":8080", handlers.CORS(headersOk, methodsOk)(a.Router)))
}

type core struct {
	PGUser   string
	PGPass   string
	PGHost   string
	PGPort   string
	PGDB     string
	APIToken string
}

func main() {
	cr := new(core)
	app := cli.NewApp()
	app.Name = "supergiant-server"
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:        "pguser",
			Usage:       "Postgres User",
			Destination: &cr.PGUser,
		},
		cli.StringFlag{
			Name:        "pgpass",
			Usage:       "Postgres Password",
			Destination: &cr.PGPass,
		},
		cli.StringFlag{
			Name:        "pghost",
			Usage:       "Postgres Host",
			Destination: &cr.PGHost,
		},
		cli.StringFlag{
			Name:        "pgport",
			Usage:       "Postgres Port",
			Destination: &cr.PGPort,
		},
		cli.StringFlag{
			Name:        "pgdb",
			Usage:       "Postgres DB",
			Destination: &cr.PGDB,
		},
		cli.StringFlag{
			Name:        "apitoken",
			Usage:       "SG API Token",
			Destination: &cr.APIToken,
		},
	}
	a := App{}

	// db, err := sql.Open("postgres", dbinfo)
	// db, err := sql.Open("sqlite3", "./local.db")
	// if err != nil {
	// 	log.Fatal("Failed to init DB")
	// }
	// a.DB = db
	a.Router = mux.NewRouter()
	a.Router.HandleFunc("/claim", a.useInvite).Methods("GET")
	// a.Router.HandleFunc("/request", a.createProduct).Methods("POST")
	app.Action = func(c *cli.Context) error {
		dbinfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
			cr.PGHost, cr.PGPort, cr.PGUser, cr.PGPass, cr.PGDB)
		log.Debug(dbinfo)

		db, err := sql.Open("postgres", dbinfo)
		if err != nil {
			log.Fatal("Failed to init DB")
		}

		err = db.Ping()
		log.Debug(err)
		if err != nil {
			log.Fatal("Failed to init DB")
		}
		a.DB = db
		a.Run()
		return nil
	}
	app.Run(os.Args)
}
