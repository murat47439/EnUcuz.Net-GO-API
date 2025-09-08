package admin

import (
	"Store-Dio/models"
	"Store-Dio/services/reviews"
	"encoding/json"
	"net/http"
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

	var review *models.Review
	err := json.NewDecoder(r.Body).Decode(&review)

	if err != nil {
		RespondWithError(w, http.StatusBadRequest, "Invalid data")
		return
	}
	err = rc.ReviewService.ReviewStatusUpdate(review)
	if err != nil {
		RespondWithError(w, http.StatusBadRequest, err.Error())

	}
	RespondWithJSON(w, http.StatusOK, map[string]string{
		"message": "Successfully",
	})
}
