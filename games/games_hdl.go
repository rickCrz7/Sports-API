package games

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rickCrz7/Sports-API/utils"
)

type Handler struct {
	dao Dao
}

func NewHandler(dao Dao, r *mux.Router) *Handler {
	h := &Handler{dao: dao}
	r.HandleFunc("/api/v1/games", h.getGames).Methods("GET")
	r.HandleFunc("/api/v1/games/{id}", h.getGameByID).Methods("GET")
	r.HandleFunc("/api/v1/games/team/{team_id}", h.getGamesByTeamID).Methods("GET")
	r.HandleFunc("/api/v1/games", h.createGame).Methods("POST")
	r.HandleFunc("/api/v1/games/{id}", h.updateGame).Methods("PUT")
	r.HandleFunc("/api/v1/games/{id}", h.deleteGame).Methods("DELETE")
	return h
}

func (h *Handler) getGames(w http.ResponseWriter, r *http.Request) {
	games, err := h.dao.getGames()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(games)
}

func (h *Handler) getGameByID(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	game, err := h.dao.getGameByID(utils.ConvertToInt(id))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(game)
}

func (h *Handler) getGamesByTeamID(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["team_id"]
	games, err := h.dao.getGamesByTeamID(utils.ConvertToInt(id))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(games)
}

func (h *Handler) createGame(w http.ResponseWriter, r *http.Request) {
	game := &utils.Games{}
	err := json.NewDecoder(r.Body).Decode(game)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	err = h.dao.CreateGame(game)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
}

func (h *Handler) updateGame(w http.ResponseWriter, r *http.Request) {
	game := &utils.Games{}
	err := json.NewDecoder(r.Body).Decode(game)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	err = h.dao.UpdateGame(game)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func (h *Handler) deleteGame(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	err := h.dao.DeleteGame(utils.ConvertToInt(id))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}
