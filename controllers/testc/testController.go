package testc

import (
	"Store-Dio/controllers/admin"
	"Store-Dio/models/testm"
	"Store-Dio/services/tests"

	"encoding/json"
	"net/http"
)

type TestController struct {
	TestService *tests.DataService
}

func NewTestController(test *tests.DataService) *TestController {
	return &TestController{
		TestService: test,
	}
}

func (tc *TestController) InsertData(w http.ResponseWriter, r *http.Request) {
	var data *testm.Items

	err := json.NewDecoder(r.Body).Decode(&data)

	if err != nil {
		admin.RespondWithError(w, http.StatusBadRequest, err.Error())
		return
	}
	err = tc.TestService.InsertData(data)
	if err != nil {
		admin.RespondWithError(w, http.StatusBadRequest, err.Error())
		return
	}
	admin.RespondWithJSON(w, http.StatusOK, map[string]string{
		"message": "Successfully",
	})
}
