package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Greeting!")
}

func widgetList(w http.ResponseWriter, r *http.Request) {
	widgets := Widgets{
		Widget{ID: "1", Name: "Item A"},
		Widget{ID: "2", Name: "Item B"},
	}

	if err := json.NewEncoder(w).Encode(widgets); err != nil {
		panic(err)
	}
}

func widgetShow(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	ID := vars["widgetID"]
	fmt.Fprintln(w, "Widget show:", ID)
}
