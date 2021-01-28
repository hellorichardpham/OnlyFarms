package main

import (
	"encoding/json"
	"net/http"
	"strconv"
)

func (app *application) getItem(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Query().Get(":id"))
	if err != nil || id < 1 {
		app.notFound(w)
		return
	}

	items, err := app.items.Get(id)
	if err != nil {
		app.handleError(w, err)
	} else {
		json.NewEncoder(w).Encode(items)
	}
}

func (app *application) getPackage(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Query().Get(":id"))
	if err != nil || id < 1 {
		app.notFound(w)
		return
	}

	pack, err := app.packages.Get(id)
	if err != nil {
		app.handleError(w, err)
	} else {
		json.NewEncoder(w).Encode(pack)
	}
}

func (app *application) getFarmer(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Query().Get(":id"))
	if err != nil || id < 1 {
		app.notFound(w)
		return
	}

	farmer, err := app.farmers.Get(id)
	if err != nil {
		app.handleError(w, err)
	} else {
		json.NewEncoder(w).Encode(farmer)
	}
}

func (app *application) getPackageItems(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Query().Get(":id"))
	if err != nil || id < 1 {
		app.notFound(w)
		return
	}

	packs, err := app.items.GetItemsByPackageId(id)
	if err != nil {
		app.handleError(w, err)
	} else {
		json.NewEncoder(w).Encode(packs)
	}
}
