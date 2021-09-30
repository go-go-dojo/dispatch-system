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

	type fields struct {
		drivers map[string]*models.DriverInfo
	}
	type args struct {
		newDriver *models.DriverInfo
	}

	tests := struct {
		name   string
		fields fields
		args   args
	}{
		name: "processDriverInfoSuccessTest",
		fields: fields{
			drivers: mapa,
		},
		args: args{},
	}

	s := &DriverRepository{
		drivers: tests.fields.drivers,
	}
	s.ProcessDriverInfo(&d)
	if got, ok := s.drivers["717995b2-978b-4351-9050-873be05e014c"]; ok {
		if *got != d {
			t.Errorf("Drivers are not the same")
		}
	} else {
		t.Errorf("Driver not in map")
	}
}
