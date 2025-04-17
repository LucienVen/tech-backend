package data_gen

import (
	"github.com/LucienVen/tech-backend/bootstrap"
	"github.com/LucienVen/tech-backend/manager/log"
	"testing"
)

func First() {
	app := bootstrap.App()
	app.StartHeartbeat()
	defer app.CloseApplication()

	log.InitLogger(app.Env)
}

func TestMock(t *testing.T) {
	First()

	err := Mock()
	if err != nil {
		t.Error(err)
	}

	t.Log("Mock data generated successfully")
}

func TestInsertTeachers(t *testing.T) {
	First()
	InsertTeachers()
}
