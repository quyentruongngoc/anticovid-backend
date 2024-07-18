package main

import (
	"anti-corona-backend/internal"
	"anti-corona-backend/package/repository"
	accountHandler "anti-corona-backend/package/rest-api/account"
	statusHandler "anti-corona-backend/package/rest-api/status"
	accountService "anti-corona-backend/package/service/account"
	statusService "anti-corona-backend/package/service/status"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/gorilla/mux"
)

const (
	apiPort = 20000
	dbHost  = "localhost"
	dbPort  = "3306"
	dbType  = "mysql"
	dbName  = "medical_data"
	// dbUser     = "quyen"
	// dbPassword = "NGOC@quyendb123"
	dbUser     = "medical"
	dbPassword = "CORONAvirus!@#123"
)

var (
	envVariables map[string]interface{}
)

func loadEnvVariables() {
	var env string
	var found bool

	envVariables = map[string]interface{}{
		"API_PORT": apiPort,    // Specify port for software rest API
		"DB_HOST":  dbHost,     // Specify db host
		"DB_PORT":  dbPort,     // Specify db port
		"DB_TYPE":  dbType,     // Specify db type
		"DB_NAME":  dbName,     // Specify db name
		"DB_USER":  dbUser,     // Specify db user name
		"DB_PASS":  dbPassword, // Specify db password
	}

	for k := range envVariables {
		if env, found = os.LookupEnv(k); !found {
			continue
		}

		env = strings.ReplaceAll(env, "\"", "")
		if k == "DB_PORT" || k == "API_PORT" {
			port, err := strconv.Atoi(env)
			if err == nil {
				envVariables[k] = port
			}
		} else {
			envVariables[k] = env
		}
	}
}

func enableApi(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		allowMethods := strings.Join([]string{
			http.MethodGet,
			http.MethodPost,
			http.MethodPut,
			http.MethodDelete,
			http.MethodOptions,
			http.MethodHead,
		}, ",")
		allowHeaders := strings.Join([]string{
			"Accept",
			"Accept-Language",
			"Accept-Encoding",
			"Content-Type",
			"Content-Length",
			"X-CSRF-Token",
			"X-Requested-With",
			"Origin",
			"Authorization",
		}, ",")

		// fmt.Println("MUX MiddlewareFunc allow methods:", allowMethods, ", headers:", allowHeaders)

		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", allowMethods)
		w.Header().Set("Access-Control-Allow-Headers", allowHeaders)

		// fmt.Println("MUX MiddlewareFunc received request", r.URL, ", method:", r.Method)

		if r.Method == http.MethodOptions {
			fmt.Println("MUX MiddlewareFunc return status", http.StatusOK)
			w.WriteHeader(http.StatusOK)
			return
		}

		// fmt.Println("MUX MiddlewareFunc call next handler")

		// Call the next handler, which can be another middleware in the chain, or the final handler.
		next.ServeHTTP(w, r)
	})
}

func main() {
	loadEnvVariables()
	internal.InitInternal()

	c := repository.Config{
		Host:   envVariables["DB_HOST"].(string),
		Port:   envVariables["DB_PORT"].(string),
		Type:   envVariables["DB_TYPE"].(string),
		Name:   envVariables["DB_NAME"].(string),
		User:   envVariables["DB_USER"].(string),
		Passwd: envVariables["DB_PASS"].(string),
	}

	var err error
	if err = repository.InitStorage(c); err != nil {
		log.Fatalf("Failed to init storage, config: %v, err: %v", c, err)
	}

	var repo *repository.Storage
	if repo, err = repository.NewStorage(); err != nil {
		log.Fatalf("Failed to create new storage, err: %v", err)
	}

	accountService := accountService.New(repo)
	// patientService := patientService.New(repo)
	statusService := statusService.New(repo)

	// init rest api handler
	router := mux.NewRouter()
	accountHandler.Install(router, accountService)
	// patientHandler.Install(router, patientService)
	statusHandler.Install(router, statusService)

	http.Handle("/", enableApi(router))
	listener := fmt.Sprintf("0.0.0.0:%d", envVariables["API_PORT"].(int))
	log.Fatal(string("Serving at ")+listener,
		http.ListenAndServe(listener, nil))
}
