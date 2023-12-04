package games

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rickCrz7/Sports-API/utils"
)

type ScoresHandler struct {
	dao ScoresDao
}

func NewScoresHandler(dao ScoresDao, r *mux.Router) *ScoresHandler {
	h := &ScoresHandler{dao: dao}
	r.HandleFunc("/api/v1/scores", h.createScore).Methods("POST")
	r.HandleFunc("/api/v1/scores/{game_id}", h.updateScore).Methods("PUT")
	r.HandleFunc("/api/v1/scores/{id}", h.deleteScore).Methods("DELETE")
	return h
}

func (h *ScoresHandler) createScore(w http.ResponseWriter, r *http.Request) {
	score := &utils.Scores{}
	err := json.NewDecoder(r.Body).Decode(score)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	err = h.dao.createScore(score)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
}

func (h *ScoresHandler) updateScore(w http.ResponseWriter, r *http.Request) {
	score := &utils.Scores{}
	err := json.NewDecoder(r.Body).Decode(score)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	err = h.dao.updateScore(score)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func (h *ScoresHandler) deleteScore(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	err := h.dao.deleteScore(utils.ConvertToInt(id))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}
