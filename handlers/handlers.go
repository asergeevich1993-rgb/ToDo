package handlers

import (
	"encoding/json"
	todo "http/list"
	"net/http"
	"strconv"
	"time"
)

type Handler struct {
	storage *todo.List
}

func NewHandler(td *todo.List) *Handler {
	return &Handler{
		storage: td,
	}
}

func (h *Handler) HandleCreateTask(w http.ResponseWriter, r *http.Request) {

	var t TaskDTO
	if err := json.NewDecoder(r.Body).Decode(&t); err != nil {
		writeErrors(w, http.StatusBadRequest, err.Error())
		return
	}
	if err := t.ValidateForCreate(); err != nil {
		writeErrors(w, http.StatusBadRequest, err.Error())
		return
	}
	task := todo.NewTask(t.Title, t.Descrition)
	if err := h.storage.Add(task); err != nil {
		writeErrors(w, http.StatusConflict, err.Error())
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(201)
	json.NewEncoder(w).Encode(task)
}
func (h *Handler) HandleCompliteTask(w http.ResponseWriter, r *http.Request) {
	str := r.PathValue("title")
	if err := h.storage.Done(str); err != nil {
		writeErrors(w, http.StatusNotFound, err.Error())
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	json.NewEncoder(w).Encode(map[string]string{"status": "complited"})

}
func (h *Handler) HandleUnCompliteTask(w http.ResponseWriter, r *http.Request) {
	str := r.PathValue("title")
	if err := h.storage.UnDone(str); err != nil {
		writeErrors(w, http.StatusNotFound, err.Error())
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	json.NewEncoder(w).Encode(map[string]string{"status": "uncomplited"})

}
func (h *Handler) HandleShowTasks(w http.ResponseWriter, r *http.Request) {
	complite := r.URL.Query().Get("complited")
	if complite == "" {
		list := h.storage.Show()
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(200)
		json.NewEncoder(w).Encode(list)
		return
	}
	cbool, err := strconv.ParseBool(complite)
	if err != nil {
		writeErrors(w, http.StatusBadRequest, err.Error())
		return
	}

	list := h.storage.ShowComplitedUnComplited(cbool)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	json.NewEncoder(w).Encode(list)

}
func (h *Handler) HandleDeleteTask(w http.ResponseWriter, r *http.Request) {
	task := r.PathValue("title")
	if err := h.storage.Del(task); err != nil {
		writeErrors(w, http.StatusNotFound, err.Error())
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	json.NewEncoder(w).Encode(map[string]string{"status": "delete"})

}

func writeErrors(w http.ResponseWriter, status int, err string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(ErrorDTO{
		Err:  err,
		Time: time.Now(),
	})
}
