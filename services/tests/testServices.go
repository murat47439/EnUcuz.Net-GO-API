package tests

import (
	"Store-Dio/models/testm"
	"Store-Dio/repo/test"
)

type DataService struct {
	DataRepo *test.DataRepo
}

func NewTestServices(dataRepo *test.DataRepo) *DataService {
	return &DataService{DataRepo: dataRepo}
}
func (ds *DataService) InsertData(data *testm.Items) error {
	err := ds.DataRepo.InsertData(data)

	if err != nil {
		return err
	}
	return nil
}
