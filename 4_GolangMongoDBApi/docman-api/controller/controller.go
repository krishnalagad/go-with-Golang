package controller

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/krishnalagad/docman-api/model"
	"github.com/krishnalagad/docman-api/payload"
)

// controllers
func ServerHome(res http.ResponseWriter, req *http.Request) {
	res.Write([]byte("<h1>Welcome to DocmanMongoDB API by Krishna</h1>"))
}

func CreateDocument(res http.ResponseWriter, req *http.Request) {

	res.Header().Set("Content-Type", "application/x-www-form-urlencoded")
	res.Header().Set("Allow-Control-Allow-Methods", "POST")

	var document model.Document
	_ = json.NewDecoder(req.Body).Decode(&document)
	data := payload.InsertOneDocument(document)
	json.NewEncoder(res).Encode(data)
	return
}

func GetOneDocument(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "application/x-www-form-urlencoded")
	res.Header().Set("Allow-Control-Allow-Methods", "GET")

	params := mux.Vars(req)
	document := payload.GetOneDocument(params["id"])

	json.NewEncoder(res).Encode(document)
	return
}

func UpdateOneDocument(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "application/x-www-form-urlencoded")
	res.Header().Set("Allow-Control-Allow-Methods", "PUT")

	// params := mux.Vars(req)
	
}