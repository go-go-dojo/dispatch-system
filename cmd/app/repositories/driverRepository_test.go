package repositories

import (
	"daitan-dispatch-system/cmd/app/models"
	"testing"
)

func TestDriverRepository_ProcessDriverInfo(t *testing.T) {

		d := models.DriverInfo{
			Uuid:     "1",
			Name:     "Luiz Henrique",
			Ranking:  0.5,
			Trips:    1,
			Car:      models.Car{},
			Status:   0,
			Location: models.Location{},
		}






	type fields struct {
		drivers    map[string]*models.DriverInfo
		requestCh  chan interface{}
		ResponseCh chan interface{}
	}
	type args struct {
		newDriver *models.DriverInfo
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &DriverRepository{
				drivers:    tt.fields.drivers,
				requestCh:  tt.fields.requestCh,
				ResponseCh: tt.fields.ResponseCh,
			}
		})
	}
}
