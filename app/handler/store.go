package handler

import (
	"encoding/json"
	"net/http"
	"fmt"
	"github.com/gorilla/mux"
	"strings"
)

var memory map[string]string

func init() {
	memory = make(map[string] string)
}

type Store struct {
	Key string `json:"key"`
	Value string  `json:"value"`
}

func CreatePair(w http.ResponseWriter, r *http.Request) {
	store := Store{}

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&store); err != nil {
		respondError(w, http.StatusBadRequest, err.Error())
		return
	}

	fmt.Println(store.Key)
	
	defer r.Body.Close()

	memory[store.Key] = store.Value

	respondJSON(w, http.StatusCreated, "pair-created")
}

func GetValue(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	key := vars["key"]	
	val := memory[key] 

	respondJSON(w, http.StatusOK, val)
}

func SearchPrefixValue(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	prefix := vars["prefix"]

	keys := make([]string, len(memory))
	i := 0

	for key, _ := range memory {

		fmt.Printf("Key: %s Value: %s\n", key, prefix)

		if strings.HasPrefix(key, prefix) == true {
			keys[i] = key
			i = i + 1
		}
	}

	respondJSON(w, http.StatusOK, keys[:i])
}

func SearchSuffixValue(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	suffix := vars["suffix"]

	keys := make([]string, len(memory))

	i := 0

	for key, _ := range memory {

		fmt.Printf("Key: %s Value: %s\n", key, suffix)

		if strings.HasSuffix(key, suffix) == true {
			keys[i] = key
			i = i + 1
		}
	}

	respondJSON(w, http.StatusOK, keys[:i])
}