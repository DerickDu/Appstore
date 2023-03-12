package handler

import (
	"appstore/model"
	"appstore/service"
	"encoding/json"
	"fmt"
	"net/http"
)

func uploadHandler(w http.ResponseWriter, r *http.Request) {
	// Parse from body of request to get a json object.
	fmt.Println("Received one upload request")
	decoder := json.NewDecoder(r.Body)
	var app model.App
	if err := decoder.Decode(&app); err != nil {
		panic(err)
	}
	// call service level function to handle this Request

	// return a message
	fmt.Fprintf(w, "Upload request received: %s\n", app.Description)
}

func searchHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Received one search request")

	// get param from request
	w.Header().Set("Content-Type", "application/json")
	title := r.URL.Query().Get("title")
	description := r.URL.Query().Get("description")

	// call service to handle request
	var apps []model.App
	var err error
	apps, err = service.SearchApps(title, description)
	if err != nil {
		http.Error(w, "Failed to read Apps from backend", http.StatusInternalServerError)
		return
	}

	// construct Response
	js, err := json.Marshal(apps)
	if err != nil {
		http.Error(w, "Failed to parse Apps into JSON format", http.StatusInternalServerError)
		return
	}
	w.Write(js) // httpWriter connects to the response, then we don't need return value. HttpWriter is related to the router
}
