package main

import "net/http"

func (app *app) heathCheckHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("all good"))
}
