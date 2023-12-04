package sports

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
	r.HandleFunc("/api/v1/sports", h.getAllSports).Methods("GET")
	r.HandleFunc("/api/v1/sports/{id}", h.getSportByID).Methods("GET")
	r.HandleFunc("/api/v1/sports", h.createSport).Methods("POST")
	r.HandleFunc("/api/v1/sports/{id}", h.updateSport).Methods("PUT")
	r.HandleFunc("/api/v1/sports/{id}", h.deleteSport).Methods("DELETE")
	return h
}

func (h *Handler) getAllSports(w http.ResponseWriter, r *http.Request) {
	sports, err := h.dao.getAllSports()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(sports)
}

func (h *Handler) getSportByID(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	sport, err := h.dao.getSportByID(utils.ConvertToInt(id))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(sport)
}

func (h *Handler) createSport(w http.ResponseWriter, r *http.Request) {
	sport := &utils.Sports{}
	err := json.NewDecoder(r.Body).Decode(sport)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	err = h.dao.CreateSport(sport)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
}

func (h *Handler) updateSport(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	sport := new(utils.Sports)
	err := json.NewDecoder(r.Body).Decode(sport)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	sport.ID = utils.ConvertToInt(id)
	err = h.dao.UpdateSport(sport)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func (h *Handler) deleteSport(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	err := h.dao.DeleteSport(utils.ConvertToInt(id))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}

