package data_gen

import (
	"github.com/LucienVen/tech-backend/bootstrap"
	"github.com/LucienVen/tech-backend/manager/log"
	"testing"
)

func First() {
	bootstrap.Run()
	bootstrap.App.StartHeartbeat()
	log.InitLogger(bootstrap.App.Env)
}

func TestMock(t *testing.T) {
	First()
	defer bootstrap.App.CloseApplication()

	err := Mock()
	if err != nil {
		t.Error(err)
	}

	t.Log("Mock data generated successfully")
}

func TestInsertTeachers(t *testing.T) {
	First()
	defer bootstrap.App.CloseApplication()
	InsertTeachers()
}

func TestInsertClasses(t *testing.T) {
	First()
	defer bootstrap.App.CloseApplication()
	InsertClasses()
}

func TestInsertStudent(t *testing.T) {
	First()
	defer bootstrap.App.CloseApplication()
	InsertStudent()
}

func TestInsertExam(t *testing.T) {
	First()
	defer bootstrap.App.CloseApplication()
	InsertExam()
}

func TestInsertSubjects(t *testing.T) {
	First()
	defer bootstrap.App.CloseApplication()
	InsertSubjects()
}

func TestTeacherSubjectRelation(t *testing.T) {
	First()
	defer bootstrap.App.CloseApplication()
	err := TeacherSubjectRelation()
	if err != nil {
		t.Error(err)
		return
	}

	t.Log("TeacherSubjectRelation success")
}

func TestMainClassTeacherRelation(t *testing.T) {
	First()
	defer bootstrap.App.CloseApplication()
	MainClassTeacherRelation()

	t.Log("MainClassTeacherRelation success")

}
