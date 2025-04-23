package entity

import (
	"github.com/LucienVen/tech-backend/bootstrap"
	"github.com/LucienVen/tech-backend/internal/utils"
	"testing"
)

func TestGetAllTeacher(t *testing.T) {
	utils.FirstInit()
	res, err := GetAllTeacher()
	if err != nil {
		t.Error(err)
	}

	//t.Log(res)
	utils.StructPrintf(res)
}

func TestTeacherEntity_UpdateSubjectId(t *testing.T) {
	utils.FirstInit()
	handler := "lxt"

	res, _ := GetAllTeacher()

	if len(res) == 0 {
		t.Error("no teacher")
		return
	}

	temp := res[0]

	instance := NewTeacherEntity(bootstrap.Run().Mysql, temp)
	err := instance.UpdateSubjectId("xxxxx", handler)
	if err != nil {
		t.Error(err)
		return
	}

	t.Log("UpdateSubjectId success")
}
