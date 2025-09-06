package user

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

func (rc *ReviewController) AddReview(w http.ResponseWriter, r *http.Request) {
	userID, _, ok := GetUserIDFromContext(r)

	if !ok {
		RespondWithError(w, http.StatusUnauthorized, "Unauthorized")
		return
	}
	var review *models.Review

	err := json.NewDecoder(r.Body).Decode(&review)

	if err != nil {
		RespondWithError(w, http.StatusBadRequest, "Invalid data")
		return
	}
	review.UserID = userID
	err = rc.ReviewService.AddReview(review)

	if err != nil {
		RespondWithError(w, http.StatusBadRequest, err.Error())
		return
	}
	RespondWithJSON(w, http.StatusOK, map[string]string{
		"Message": "Successfully",
	})
}
func (rc *ReviewController) UpdateReview(w http.ResponseWriter, r *http.Request) {
	userID, _, ok := GetUserIDFromContext(r)

	if !ok {
		RespondWithError(w, http.StatusUnauthorized, "Unauthorized")
		return
	}
	var review *models.Review

	err := json.NewDecoder(r.Body).Decode(&review)

	if err != nil {
		RespondWithError(w, http.StatusBadRequest, "Invalid data")
		return
	}
	review.UserID = userID
	err = rc.ReviewService.UpdateReview(review)

	if err != nil {
		RespondWithError(w, http.StatusBadRequest, err.Error())
		return
	}
	RespondWithJSON(w, http.StatusOK, map[string]string{
		"Message": "Successfully",
	})
}
func (rc *ReviewController) RemoveReview(w http.ResponseWriter, r *http.Request) {

	userID, _, ok := GetUserIDFromContext(r)

	if !ok {
		RespondWithError(w, http.StatusUnauthorized, "Unauthorized")
		return
	}

	id, err := strconv.Atoi(chi.URLParam(r, "id"))

	if err != nil {
		RespondWithError(w, http.StatusBadRequest, "Invalid data")
		return
	}

	err = rc.ReviewService.RemoveReview(id, userID)

	if err != nil {
		RespondWithError(w, http.StatusBadRequest, err.Error())
		return
	}
	RespondWithJSON(w, http.StatusOK, map[string]string{
		"Message": "Successfully",
	})
}
func (rc *ReviewController) GetReview(w http.ResponseWriter, r *http.Request) {
	userID, userRole, ok := GetUserIDFromContext(r)

	if !ok {
		RespondWithError(w, http.StatusUnauthorized, "Unauthorized")
		return
	}
	id, err := strconv.Atoi(chi.URLParam(r, "id"))

	if err != nil {
		RespondWithError(w, http.StatusBadRequest, "Invalid data")
		return
	}
	review, err := rc.ReviewService.GetReview(id, userRole, userID)

	if err != nil {
		RespondWithError(w, http.StatusBadRequest, err.Error())
		return
	}
	RespondWithJSON(w, http.StatusOK, map[string]interface{}{
		"Message": "Successfully",
		"Review":  review,
	})
}
func (rc *ReviewController) GetReviews(w http.ResponseWriter, r *http.Request) {

	prodID, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		RespondWithError(w, http.StatusBadRequest, "Invalid data")
		return
	}
	page, err := strconv.Atoi(r.URL.Query().Get("page"))

	if err != nil {
		page = 1
	}
	reviews, err := rc.ReviewService.GetReviews(page, prodID)
	if err != nil {
		RespondWithError(w, http.StatusBadRequest, err.Error())
		return
	}
	RespondWithJSON(w, http.StatusOK, map[string]interface{}{
		"Message": "Successfully",
		"Reviews": reviews,
	})

}
func (rc *ReviewController) GetUserReviews(w http.ResponseWriter, r *http.Request) {

	userID, _, ok := GetUserIDFromContext(r)

	if !ok {
		RespondWithError(w, http.StatusUnauthorized, "Unauthorized")
		return
	}

	reviews, err := rc.ReviewService.GetUserReviews(userID)

	if err != nil {
		RespondWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	RespondWithJSON(w, http.StatusOK, map[string]interface{}{
		"Message": "Successfully",
		"Reviews": reviews,
	})
}
