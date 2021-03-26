package handler

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)


func CreatePair(w http.ResponseWriter, r *http.Request) {
	
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&project); err != nil {
		respondError(w, http.StatusBadRequest, err.Error())
		return
	}
	defer r.Body.Close()

	// if err := db.Save(&project).Error; err != nil {
	// 	respondError(w, http.StatusInternalServerError, err.Error())
	// 	return
	// }
	respondJSON(w, http.StatusCreated, project)
}

func GetValue(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	key := vars["key"]

	value := nil

	respondJSON(w, http.StatusOK, value)
}
