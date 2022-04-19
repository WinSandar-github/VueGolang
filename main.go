package main

import (
	"net/http"

	api "backend-golang/src/api"

	_ "github.com/go-sql-driver/mysql"

	"github.com/gorilla/mux"
)

func main() {
	Routers()

}

func Routers() {

	router := mux.NewRouter()

	router.HandleFunc("/api/login", api.Login).Methods("POST")
	router.HandleFunc("/api/getUser", api.GetUser).Methods("GET")
	router.HandleFunc("/api/createUser", api.CreateUser).Methods("POST")
	router.HandleFunc("/api/createProduct", api.CreateProduct).Methods("POST")
	router.HandleFunc("/api/getProduct", api.GetProduct).Methods("GET")
	router.HandleFunc("/api/updateProduct", api.UpdateProduct).Methods("PUT")
	router.HandleFunc("/api/deleteProduct/{id}", api.DeleteProduct).Methods("DELETE")
	http.ListenAndServe(":9080",
		&CORSRouterDecorator{router})
}

// CORSRouterDecorator applies CORS headers to a mux.Router
type CORSRouterDecorator struct {
	R *mux.Router
}

func (c *CORSRouterDecorator) ServeHTTP(rw http.ResponseWriter,
	req *http.Request) {
	if origin := req.Header.Get("Origin"); origin != "" {
		rw.Header().Set("Access-Control-Allow-Origin", origin)
		rw.Header().Set("Access-Control-Allow-Methods",
			"POST, GET, OPTIONS, PUT, DELETE")
		rw.Header().Set("Access-Control-Allow-Headers",
			"Accept, Accept-Language,"+
				" Content-Type, YourOwnHeader,X-CSRF-Token, *")
		rw.Header().Set("Access-Control-Allow-Credentials", "true")
	}
	// Stop here if its Preflighted OPTIONS request
	if req.Method == "OPTIONS" {
		return
	}

	c.R.ServeHTTP(rw, req)
}
