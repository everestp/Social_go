package main

import (
	"log"
	"net/http"
)


func (app *application)  internalServerError (w http.ResponseWriter , r *http.Request ,err error){
	log.Printf(" internnall error :  %s  path : %s  error : %s",r.Method , r.URL.Path ,err)
	writeJSONError(w, http.StatusInternalServerError, "The server encountered a problem")
}

func (app *application)  badRequestResponse (w http.ResponseWriter , r *http.Request ,err error){
	log.Printf(" bad request error :  %s  path : %s  error : %s",r.Method , r.URL.Path ,err)
	writeJSONError(w, http.StatusInternalServerError, err.Error())
}

func (app *application)  notFoundResponse (w http.ResponseWriter , r *http.Request ,err error){
	log.Printf("not found error :  %s  path : %s  error : %s",r.Method , r.URL.Path ,err)
	writeJSONError(w, http.StatusNotFound, "Resource not found")
}