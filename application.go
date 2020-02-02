package main

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"io/ioutil"
	"log"
	"net/http"
)

/*
	StartServer starts the web application server by initializing the web handlers and middleware.
*/
func StartServer() {
	router := httprouter.New()
	router.GET("/", indexHandler)
	router.GET("/download/*filepath", downloadHandler)

	// todo HTTPS
	// todo parameterize port
	log.Println("Starting webapp listener")
	log.Fatalln(http.ListenAndServe(":8080", logRequest(router)))
}

func downloadHandler(response http.ResponseWriter, _ *http.Request, _ httprouter.Params) {
	response.WriteHeader(http.StatusNotImplemented)
}

func indexHandler(response http.ResponseWriter, _ *http.Request, _ httprouter.Params) {
	fileData, err := ioutil.ReadFile("index.html")
	if err != nil {
		log.Printf("Index read error: %v", err)
		response.WriteHeader(http.StatusInternalServerError)
	} else {
		_, err := fmt.Fprint(response, string(fileData))
		if err != nil {
			log.Printf("Response write error: %v", err)
			response.WriteHeader(http.StatusInternalServerError)
		}
	}
}

func logRequest(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(response http.ResponseWriter, request *http.Request) {
		log.Printf("Processing request [%s] %s", request.Method, request.RequestURI)
		handler.ServeHTTP(response, request)
	})
}