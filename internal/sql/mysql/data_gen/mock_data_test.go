package data_gen

import (
	"github.com/LucienVen/tech-backend/bootstrap"
	"github.com/LucienVen/tech-backend/manager/log"
	"testing"
)

func First() {
	bootstrap.Run()

	bootstrap.App.StartHeartbeat()
	defer bootstrap.App.CloseApplication()

	log.InitLogger(bootstrap.App.Env)
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

func TestInsertClasses(t *testing.T) {
	First()
	InsertClasses()
}

func TestInsertStudent(t *testing.T) {
	First()
	InsertStudent()
}

func TestInsertExam(t *testing.T) {
	First()
	InsertExam()
}

func TestInsertSubjects(t *testing.T) {
	First()
	InsertSubjects()
}

func TestTeacherSubjectRelation(t *testing.T) {
	First()
	err := TeacherSubjectRelation()
	if err != nil {
		t.Error(err)
		return
	}

	t.Log("TeacherSubjectRelation success")
}
