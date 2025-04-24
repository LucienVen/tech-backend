package entity

import (
	"github.com/LucienVen/tech-backend/internal/utils"
	"testing"
)

func TestGetAllSubject(t *testing.T) {
	utils.FirstInit()

	subjects, _ := GetAllSubject()
	t.Log(subjects)
}

func TestGetSubjectByName(t *testing.T) {
	cleanFunc := utils.FirstInit()
	defer cleanFunc()

	subjects, err := GetSubjectByName("语文", "数学")
	if err != nil {
		t.Error(err)
	}

	if len(subjects) == 0 {
		t.Error("no subject")
		return
	}

	utils.StructPrintf(subjects)
}
