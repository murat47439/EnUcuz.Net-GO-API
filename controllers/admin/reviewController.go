package admin

import (
	"Store-Dio/models"
	"Store-Dio/services/reviews"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

type ReviewController struct {
	ReviewService *reviews.ReviewService
}

func NewReviewController(service *reviews.ReviewService) *ReviewController {
	return &ReviewController{
		ReviewService: service,
	}
}
func (rc *ReviewController) ReviewStatusUpdate(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		RespondWithError(w, http.StatusBadRequest, "Invalid data")
		return
	}
	var review *models.Review
	err = json.NewDecoder(r.Body).Decode(&review)

	if err != nil {
		RespondWithError(w, http.StatusBadRequest, "Invalid data")
		return
	}
	err = rc.ReviewService.ReviewStatusUpdate(id, review.Status)
	if err != nil {
		RespondWithError(w, http.StatusBadRequest, err.Error())

	}
	RespondWithJSON(w, http.StatusOK, map[string]string{
		"message": "Successfully",
	})
}
