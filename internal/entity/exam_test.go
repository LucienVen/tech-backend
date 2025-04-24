package entity

import (
	"github.com/LucienVen/tech-backend/internal/utils"
	"testing"
)

func TestGetAllExams(t *testing.T) {
	cleanFunc := utils.FirstInit()
	defer cleanFunc()

	exams, err := GetAllExams()
	if err != nil {
		t.Error(err)
	}

	utils.StructPrintf(exams)

}
