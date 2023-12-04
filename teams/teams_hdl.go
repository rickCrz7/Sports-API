package teams

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
	r.HandleFunc("/api/v1/teams", h.getAllTeams).Methods("GET")
	r.HandleFunc("/api/v1/teams/sports/{sportsID}", h.getTeamsBySportsID).Methods("GET")
	r.HandleFunc("/api/v1/teams/league/{leagueID}", h.getTeamsByLeagueID).Methods("GET")
	r.HandleFunc("/api/v1/teams/{id}", h.getTeamByID).Methods("GET")
	r.HandleFunc("/api/v1/teams", h.createTeam).Methods("POST")
	r.HandleFunc("/api/v1/teams/{id}", h.updateTeam).Methods("PUT")
	r.HandleFunc("/api/v1/teams/{id}", h.deleteTeam).Methods("DELETE")
	return h
}

func (h *Handler) getAllTeams(w http.ResponseWriter, r *http.Request) {
	teams, err := h.dao.getAllTeams()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(teams)
}

func (h *Handler) getTeamsBySportsID(w http.ResponseWriter, r *http.Request) {
	sportsID := mux.Vars(r)["sportsID"]
	teams, err := h.dao.getTeamsBySportsID(utils.ConvertToInt(sportsID))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(teams)
}

func (h *Handler) getTeamsByLeagueID(w http.ResponseWriter, r *http.Request) {
	leagueID := mux.Vars(r)["leagueID"]
	teams, err := h.dao.getTeamsByLeagueID(utils.ConvertToInt(leagueID))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(teams)
}

func (h *Handler) getTeamByID(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	team, err := h.dao.getTeamByID(utils.ConvertToInt(id))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(team)
}

func (h *Handler) createTeam(w http.ResponseWriter, r *http.Request) {
	team := &utils.Teams{}
	err := json.NewDecoder(r.Body).Decode(team)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	err = h.dao.CreateTeam(team)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
}

func (h *Handler) updateTeam(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	team := new(utils.Teams)
	err := json.NewDecoder(r.Body).Decode(team)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	team.ID = utils.ConvertToInt(id)
	err = h.dao.UpdateTeam(team)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func (h *Handler) deleteTeam(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	err := h.dao.DeleteTeam(utils.ConvertToInt(id))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}