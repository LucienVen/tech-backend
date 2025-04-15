package data_gen

import (
	"github.com/LucienVen/tech-backend/bootstrap"
	"github.com/LucienVen/tech-backend/manager/log"
	"testing"
)

func TestMock(t *testing.T) {
	app := bootstrap.App()
	app.StartHeartbeat()
	defer app.CloseApplication()

	log.InitLogger(app.Env)

	err := Mock()
	if err != nil {
		t.Error(err)
	}

	t.Log("Mock data generated successfully")
}
