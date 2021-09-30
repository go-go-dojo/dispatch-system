package repositories

import (
	"daitan-dispatch-system/cmd/app/models"
	"testing"
)

func TestDriverRepository_ProcessDriverInfo(t *testing.T) {

	d := models.DriverInfo{
		Uuid:     "717995b2-978b-4351-9050-873be05e014c",
		Name:     "Luiz Henrique",
		Ranking:  0.5,
		Trips:    1,
		Car:      models.Car{},
		Status:   0,
		Location: models.Location{},
	}

	mapa := make(map[string]*models.DriverInfo)
	mapa["717995b2-978b-4351-9050-873be05e014c"] = &d

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
		{
			name: "processDriverInfoSuccessTest",
			fields: fields{
				drivers:    mapa,
				requestCh:  nil,
				ResponseCh: nil,
			},
			args: args{},
		},
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
