package leagues

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
	r.HandleFunc("/api/v1/leagues", h.getAllLeagues).Methods("GET")
	r.HandleFunc("/api/v1/leagues/sport/{id}", h.getLeagueBySportID).Methods("GET")
	r.HandleFunc("/api/v1/leagues/{id}", h.getLeagueByID).Methods("GET")
	r.HandleFunc("/api/v1/leagues", h.createLeague).Methods("POST")
	r.HandleFunc("/api/v1/leagues/{id}", h.updateLeague).Methods("PUT")
	r.HandleFunc("/api/v1/leagues/{id}", h.deleteLeague).Methods("DELETE")
	return h
}

func (h *Handler) getAllLeagues(w http.ResponseWriter, r *http.Request) {
	leagues, err := h.dao.getAllLeagues()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(leagues)
}

func (h *Handler) getLeagueBySportID(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	leagues, err := h.dao.getLeagueBySportID(utils.ConvertToInt(id))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(leagues)
}



func (h *Handler) getLeagueByID(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	league, err := h.dao.getLeagueByID(utils.ConvertToInt(id))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(league)
}

func (h *Handler) createLeague(w http.ResponseWriter, r *http.Request) {
	league := &utils.Leagues{}
	err := json.NewDecoder(r.Body).Decode(league)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	err = h.dao.CreateLeague(league)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
}

func (h *Handler) updateLeague(w http.ResponseWriter, r *http.Request) {
	league := &utils.Leagues{}
	err := json.NewDecoder(r.Body).Decode(league)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	err = h.dao.UpdateLeague(league)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func (h *Handler) deleteLeague(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	err := h.dao.DeleteLeague(utils.ConvertToInt(id))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}
