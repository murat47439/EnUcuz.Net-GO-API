package user

import (
	"Store-Dio/models"
	"Store-Dio/services/chat"
	"context"
	"encoding/json"
	"net/http"
	"strconv"
	"time"

	"github.com/go-chi/chi/v5"
)

type ChatController struct {
	ChatService *chat.ChatService
}

func NewChatController(service *chat.ChatService) *ChatController {
	return &ChatController{ChatService: service}
}

func (cc *ChatController) CheckChat(w http.ResponseWriter, r *http.Request) {
	userID, _, ok := GetUserIDFromContext(r)
	if !ok {
		RespondWithError(w, http.StatusUnauthorized, "Unauthorized")
		return
	}
	prod_id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		RespondWithError(w, http.StatusBadRequest, "Invalid Product")
		return
	}

	ctx := r.Context()
	ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	exists, err := cc.ChatService.CheckChat(ctx, userID, prod_id)
	if err != nil {
		RespondWithError(w, http.StatusBadRequest, "Invalid Product")
		return
	}
	RespondWithJSON(w, http.StatusOK, map[string]interface{}{
		"status": exists,
	})
}
func (cc *ChatController) NewChat(w http.ResponseWriter, r *http.Request) {
	userID, _, ok := GetUserIDFromContext(r)
	if !ok {
		RespondWithError(w, http.StatusUnauthorized, "Unauthorized")
		return
	}
	var model models.NewChat
	err := json.NewDecoder(r.Body).Decode(&model)
	if err != nil {
		RespondWithError(w, http.StatusBadRequest, "Invalid data")
		return
	}
	defer r.Body.Close()
	model.Chat.Sender = userID
	ctx := r.Context()
	ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	data, mes, err := cc.ChatService.NewChat(ctx, &model.Chat, model.Message)
	if err != nil {
		RespondWithError(w, http.StatusBadRequest, err.Error())
		return
	}
	RespondWithJSON(w, http.StatusOK, map[string]interface{}{
		"message": mes,
		"chat":    data,
	})

}
func (cc *ChatController) NewMessage(w http.ResponseWriter, r *http.Request) {
	userID, _, ok := GetUserIDFromContext(r)
	if !ok {
		RespondWithError(w, http.StatusUnauthorized, "Unauthorized")
		return
	}
	var model models.Message
	err := json.NewDecoder(r.Body).Decode(&model)
	if err != nil {
		RespondWithError(w, http.StatusBadRequest, "Invalid data")
		return
	}
	defer r.Body.Close()
	model.Sender = userID
	ctx := r.Context()
	ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	mes, err := cc.ChatService.NewMessage(ctx, &model)
	if err != nil {
		RespondWithError(w, http.StatusBadRequest, err.Error())
		return
	}
	RespondWithJSON(w, http.StatusOK, map[string]interface{}{
		"message": mes,
	})

}
func (cc *ChatController) GetChat(w http.ResponseWriter, r *http.Request) {
	userID, _, ok := GetUserIDFromContext(r)
	if !ok {
		RespondWithError(w, http.StatusUnauthorized, "Unauthorized")
		return
	}
	chatID, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		RespondWithError(w, http.StatusBadRequest, "Invalid data")
		return
	}
	ctx := r.Context()
	ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()
	messages, err := cc.ChatService.GetChat(ctx, chatID, userID)
	if err != nil {
		RespondWithError(w, http.StatusBadRequest, err.Error())
		return
	}
	RespondWithJSON(w, http.StatusOK, map[string]interface{}{
		"messages": messages,
	})
}
func (cc *ChatController) GetChats(w http.ResponseWriter, r *http.Request) {
	userID, _, ok := GetUserIDFromContext(r)
	query := r.URL.Query()
	if !ok {
		RespondWithError(w, http.StatusUnauthorized, "Unauthorized")
		return
	}
	page, err := strconv.Atoi(query.Get("page"))
	if err != nil {
		page = 1
	}
	ctx := r.Context()
	ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	data, err := cc.ChatService.GetChats(ctx, userID, page)
	if err != nil {
		RespondWithError(w, http.StatusBadRequest, err.Error())
		return
	}
	RespondWithJSON(w, http.StatusOK, map[string]interface{}{
		"chats": data,
	})
}
