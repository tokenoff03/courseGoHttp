package handler

import (
	"encoding/json"
	"hotel/internal/models"
	"hotel/internal/repositories"
	"net/http"
	"strconv"
	"strings"
)

type Handler struct {
	storage *repositories.Storage
}

func NewHandler(stor *repositories.Storage) *Handler {
	return &Handler{
		storage: stor,
	}
}

func (h *Handler) RoomHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		if strings.HasPrefix(r.URL.Path, "/room/") {
			h.GetRoomById(w, r)
		} else {
			h.GetAllRooms(w, r)
		}

	case http.MethodPost:
		h.CreateRoom(w, r)
	case http.MethodDelete:

	case http.MethodPut:

	}
}
func (h *Handler) GetRoomById(w http.ResponseWriter, r *http.Request) {
	idStr := strings.TrimPrefix(r.URL.Path, "/room/")
	if idStr == "" {

		jsonResponse(w, http.StatusBadRequest, map[string]interface{}{
			"Error": "Missing room ID",
		})
		return

	}
	id, err := strconv.Atoi(idStr)
	if err != nil {
		jsonResponse(w, http.StatusBadRequest, map[string]interface{}{
			"Error": "Invalid room ID",
		})
		return
	}
	room := h.storage.GetRoomById(id)
	jsonResponse(w, http.StatusOK, map[string]interface{}{
		"room": room,
	})
}

func (h *Handler) CreateRoom(w http.ResponseWriter, r *http.Request) {
	var room models.Room
	err := json.NewDecoder(r.Body).Decode(&room)
	if err != nil {
		jsonResponse(w, http.StatusBadRequest, map[string]interface{}{
			"Error": "Invalid request body",
		})
		return
	}

	id, ErrStr := h.storage.CreateRoom(room)
	if ErrStr != "" {
		jsonResponse(w, http.StatusBadRequest, map[string]interface{}{
			"Error": "Room already exist",
		})
		return
	}

	jsonResponse(w, http.StatusOK, map[string]interface{}{
		"id": id,
	})

}

func (h *Handler) GetAllRooms(w http.ResponseWriter, r *http.Request) {
	rooms := h.storage.GetAllRooms()

	jsonResponse(w, http.StatusOK, map[string]interface{}{
		"rooms": rooms,
	})
}